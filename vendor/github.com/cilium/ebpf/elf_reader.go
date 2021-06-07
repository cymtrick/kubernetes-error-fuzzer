package ebpf

import (
	"bufio"
	"bytes"
	"debug/elf"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/btf"
	"github.com/cilium/ebpf/internal/unix"
)

// elfCode is a convenience to reduce the amount of arguments that have to
// be passed around explicitly. You should treat it's contents as immutable.
type elfCode struct {
	*internal.SafeELFFile
	sections map[elf.SectionIndex]*elfSection
	license  string
	version  uint32
	btf      *btf.Spec
}

// LoadCollectionSpec parses an ELF file into a CollectionSpec.
func LoadCollectionSpec(file string) (*CollectionSpec, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	spec, err := LoadCollectionSpecFromReader(f)
	if err != nil {
		return nil, fmt.Errorf("file %s: %w", file, err)
	}
	return spec, nil
}

// LoadCollectionSpecFromReader parses an ELF file into a CollectionSpec.
func LoadCollectionSpecFromReader(rd io.ReaderAt) (*CollectionSpec, error) {
	f, err := internal.NewSafeELFFile(rd)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		licenseSection *elf.Section
		versionSection *elf.Section
		sections       = make(map[elf.SectionIndex]*elfSection)
		relSections    = make(map[elf.SectionIndex]*elf.Section)
	)

	// This is the target of relocations generated by inline assembly.
	sections[elf.SHN_UNDEF] = newElfSection(new(elf.Section), undefSection)

	// Collect all the sections we're interested in. This includes relocations
	// which we parse later.
	for i, sec := range f.Sections {
		idx := elf.SectionIndex(i)

		switch {
		case strings.HasPrefix(sec.Name, "license"):
			licenseSection = sec
		case strings.HasPrefix(sec.Name, "version"):
			versionSection = sec
		case strings.HasPrefix(sec.Name, "maps"):
			sections[idx] = newElfSection(sec, mapSection)
		case sec.Name == ".maps":
			sections[idx] = newElfSection(sec, btfMapSection)
		case sec.Name == ".bss" || sec.Name == ".data" || strings.HasPrefix(sec.Name, ".rodata"):
			sections[idx] = newElfSection(sec, dataSection)
		case sec.Type == elf.SHT_REL:
			// Store relocations under the section index of the target
			relSections[elf.SectionIndex(sec.Info)] = sec
		case sec.Type == elf.SHT_PROGBITS && (sec.Flags&elf.SHF_EXECINSTR) != 0 && sec.Size > 0:
			sections[idx] = newElfSection(sec, programSection)
		}
	}

	license, err := loadLicense(licenseSection)
	if err != nil {
		return nil, fmt.Errorf("load license: %w", err)
	}

	version, err := loadVersion(versionSection, f.ByteOrder)
	if err != nil {
		return nil, fmt.Errorf("load version: %w", err)
	}

	btfSpec, err := btf.LoadSpecFromReader(rd)
	if err != nil && !errors.Is(err, btf.ErrNotFound) {
		return nil, fmt.Errorf("load BTF: %w", err)
	}

	// Assign symbols to all the sections we're interested in.
	symbols, err := f.Symbols()
	if err != nil {
		return nil, fmt.Errorf("load symbols: %v", err)
	}

	for _, symbol := range symbols {
		idx := symbol.Section
		symType := elf.ST_TYPE(symbol.Info)

		section := sections[idx]
		if section == nil {
			continue
		}

		// Older versions of LLVM don't tag symbols correctly, so keep
		// all NOTYPE ones.
		keep := symType == elf.STT_NOTYPE
		switch section.kind {
		case mapSection, btfMapSection, dataSection:
			keep = keep || symType == elf.STT_OBJECT
		case programSection:
			keep = keep || symType == elf.STT_FUNC
		}
		if !keep || symbol.Name == "" {
			continue
		}

		section.symbols[symbol.Value] = symbol
	}

	ec := &elfCode{
		SafeELFFile: f,
		sections:    sections,
		license:     license,
		version:     version,
		btf:         btfSpec,
	}

	// Go through relocation sections, and parse the ones for sections we're
	// interested in. Make sure that relocations point at valid sections.
	for idx, relSection := range relSections {
		section := sections[idx]
		if section == nil {
			continue
		}

		rels, err := ec.loadRelocations(relSection, symbols)
		if err != nil {
			return nil, fmt.Errorf("relocation for section %q: %w", section.Name, err)
		}

		for _, rel := range rels {
			target := sections[rel.Section]
			if target == nil {
				return nil, fmt.Errorf("section %q: reference to %q in section %s: %w", section.Name, rel.Name, rel.Section, ErrNotSupported)
			}

			if target.Flags&elf.SHF_STRINGS > 0 {
				return nil, fmt.Errorf("section %q: string is not stack allocated: %w", section.Name, ErrNotSupported)
			}

			target.references++
		}

		section.relocations = rels
	}

	// Collect all the various ways to define maps.
	maps := make(map[string]*MapSpec)
	if err := ec.loadMaps(maps); err != nil {
		return nil, fmt.Errorf("load maps: %w", err)
	}

	if err := ec.loadBTFMaps(maps); err != nil {
		return nil, fmt.Errorf("load BTF maps: %w", err)
	}

	if err := ec.loadDataSections(maps); err != nil {
		return nil, fmt.Errorf("load data sections: %w", err)
	}

	// Finally, collect programs and link them.
	progs, err := ec.loadPrograms()
	if err != nil {
		return nil, fmt.Errorf("load programs: %w", err)
	}

	return &CollectionSpec{maps, progs}, nil
}

func loadLicense(sec *elf.Section) (string, error) {
	if sec == nil {
		return "", nil
	}

	data, err := sec.Data()
	if err != nil {
		return "", fmt.Errorf("section %s: %v", sec.Name, err)
	}
	return string(bytes.TrimRight(data, "\000")), nil
}

func loadVersion(sec *elf.Section, bo binary.ByteOrder) (uint32, error) {
	if sec == nil {
		return 0, nil
	}

	var version uint32
	if err := binary.Read(sec.Open(), bo, &version); err != nil {
		return 0, fmt.Errorf("section %s: %v", sec.Name, err)
	}
	return version, nil
}

type elfSectionKind int

const (
	undefSection elfSectionKind = iota
	mapSection
	btfMapSection
	programSection
	dataSection
)

type elfSection struct {
	*elf.Section
	kind elfSectionKind
	// Offset from the start of the section to a symbol
	symbols map[uint64]elf.Symbol
	// Offset from the start of the section to a relocation, which points at
	// a symbol in another section.
	relocations map[uint64]elf.Symbol
	// The number of relocations pointing at this section.
	references int
}

func newElfSection(section *elf.Section, kind elfSectionKind) *elfSection {
	return &elfSection{
		section,
		kind,
		make(map[uint64]elf.Symbol),
		make(map[uint64]elf.Symbol),
		0,
	}
}

func (ec *elfCode) loadPrograms() (map[string]*ProgramSpec, error) {
	var (
		progs []*ProgramSpec
		libs  []*ProgramSpec
	)

	for _, sec := range ec.sections {
		if sec.kind != programSection {
			continue
		}

		if len(sec.symbols) == 0 {
			return nil, fmt.Errorf("section %v: missing symbols", sec.Name)
		}

		funcSym, ok := sec.symbols[0]
		if !ok {
			return nil, fmt.Errorf("section %v: no label at start", sec.Name)
		}

		insns, length, err := ec.loadInstructions(sec)
		if err != nil {
			return nil, fmt.Errorf("program %s: %w", funcSym.Name, err)
		}

		progType, attachType, progFlags, attachTo := getProgType(sec.Name)

		spec := &ProgramSpec{
			Name:          funcSym.Name,
			Type:          progType,
			Flags:         progFlags,
			AttachType:    attachType,
			AttachTo:      attachTo,
			License:       ec.license,
			KernelVersion: ec.version,
			Instructions:  insns,
			ByteOrder:     ec.ByteOrder,
		}

		if ec.btf != nil {
			spec.BTF, err = ec.btf.Program(sec.Name, length)
			if err != nil && !errors.Is(err, btf.ErrNoExtendedInfo) {
				return nil, fmt.Errorf("program %s: %w", funcSym.Name, err)
			}
		}

		if spec.Type == UnspecifiedProgram {
			// There is no single name we can use for "library" sections,
			// since they may contain multiple functions. We'll decode the
			// labels they contain later on, and then link sections that way.
			libs = append(libs, spec)
		} else {
			progs = append(progs, spec)
		}
	}

	res := make(map[string]*ProgramSpec, len(progs))
	for _, prog := range progs {
		err := link(prog, libs)
		if err != nil {
			return nil, fmt.Errorf("program %s: %w", prog.Name, err)
		}
		res[prog.Name] = prog
	}

	return res, nil
}

func (ec *elfCode) loadInstructions(section *elfSection) (asm.Instructions, uint64, error) {
	var (
		r      = bufio.NewReader(section.Open())
		insns  asm.Instructions
		offset uint64
	)
	for {
		var ins asm.Instruction
		n, err := ins.Unmarshal(r, ec.ByteOrder)
		if err == io.EOF {
			return insns, offset, nil
		}
		if err != nil {
			return nil, 0, fmt.Errorf("offset %d: %w", offset, err)
		}

		ins.Symbol = section.symbols[offset].Name

		if rel, ok := section.relocations[offset]; ok {
			if err = ec.relocateInstruction(&ins, rel); err != nil {
				return nil, 0, fmt.Errorf("offset %d: relocate instruction: %w", offset, err)
			}
		}

		insns = append(insns, ins)
		offset += n
	}
}

func (ec *elfCode) relocateInstruction(ins *asm.Instruction, rel elf.Symbol) error {
	var (
		typ  = elf.ST_TYPE(rel.Info)
		bind = elf.ST_BIND(rel.Info)
		name = rel.Name
	)

	target := ec.sections[rel.Section]

	switch target.kind {
	case mapSection, btfMapSection:
		if bind != elf.STB_GLOBAL {
			return fmt.Errorf("possible erroneous static qualifier on map definition: found reference to %q", name)
		}

		if typ != elf.STT_OBJECT && typ != elf.STT_NOTYPE {
			// STT_NOTYPE is generated on clang < 8 which doesn't tag
			// relocations appropriately.
			return fmt.Errorf("map load: incorrect relocation type %v", typ)
		}

		ins.Src = asm.PseudoMapFD

		// Mark the instruction as needing an update when creating the
		// collection.
		if err := ins.RewriteMapPtr(-1); err != nil {
			return err
		}

	case dataSection:
		var offset uint32
		switch typ {
		case elf.STT_SECTION:
			if bind != elf.STB_LOCAL {
				return fmt.Errorf("direct load: %s: unsupported relocation %s", name, bind)
			}

			// This is really a reference to a static symbol, which clang doesn't
			// emit a symbol table entry for. Instead it encodes the offset in
			// the instruction itself.
			offset = uint32(uint64(ins.Constant))

		case elf.STT_OBJECT:
			if bind != elf.STB_GLOBAL {
				return fmt.Errorf("direct load: %s: unsupported relocation %s", name, bind)
			}

			offset = uint32(rel.Value)

		default:
			return fmt.Errorf("incorrect relocation type %v for direct map load", typ)
		}

		// We rely on using the name of the data section as the reference. It
		// would be nicer to keep the real name in case of an STT_OBJECT, but
		// it's not clear how to encode that into Instruction.
		name = target.Name

		// The kernel expects the offset in the second basic BPF instruction.
		ins.Constant = int64(uint64(offset) << 32)
		ins.Src = asm.PseudoMapValue

		// Mark the instruction as needing an update when creating the
		// collection.
		if err := ins.RewriteMapPtr(-1); err != nil {
			return err
		}

	case programSection:
		if ins.OpCode.JumpOp() != asm.Call {
			return fmt.Errorf("not a call instruction: %s", ins)
		}

		if ins.Src != asm.PseudoCall {
			return fmt.Errorf("call: %s: incorrect source register", name)
		}

		switch typ {
		case elf.STT_NOTYPE, elf.STT_FUNC:
			if bind != elf.STB_GLOBAL {
				return fmt.Errorf("call: %s: unsupported binding: %s", name, bind)
			}

		case elf.STT_SECTION:
			if bind != elf.STB_LOCAL {
				return fmt.Errorf("call: %s: unsupported binding: %s", name, bind)
			}

			// The function we want to call is in the indicated section,
			// at the offset encoded in the instruction itself. Reverse
			// the calculation to find the real function we're looking for.
			// A value of -1 references the first instruction in the section.
			offset := int64(int32(ins.Constant)+1) * asm.InstructionSize
			if offset < 0 {
				return fmt.Errorf("call: %s: invalid offset %d", name, offset)
			}

			sym, ok := target.symbols[uint64(offset)]
			if !ok {
				return fmt.Errorf("call: %s: no symbol at offset %d", name, offset)
			}

			ins.Constant = -1
			name = sym.Name

		default:
			return fmt.Errorf("call: %s: invalid symbol type %s", name, typ)
		}

	case undefSection:
		if bind != elf.STB_GLOBAL {
			return fmt.Errorf("asm relocation: %s: unsupported binding: %s", name, bind)
		}

		if typ != elf.STT_NOTYPE {
			return fmt.Errorf("asm relocation: %s: unsupported type %s", name, typ)
		}

		// There is nothing to do here but set ins.Reference.

	default:
		return fmt.Errorf("relocation to %q: %w", target.Name, ErrNotSupported)
	}

	ins.Reference = name
	return nil
}

func (ec *elfCode) loadMaps(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != mapSection {
			continue
		}

		nSym := len(sec.symbols)
		if nSym == 0 {
			return fmt.Errorf("section %v: no symbols", sec.Name)
		}

		if sec.Size%uint64(nSym) != 0 {
			return fmt.Errorf("section %v: map descriptors are not of equal size", sec.Name)
		}

		var (
			r    = bufio.NewReader(sec.Open())
			size = sec.Size / uint64(nSym)
		)
		for i, offset := 0, uint64(0); i < nSym; i, offset = i+1, offset+size {
			mapSym, ok := sec.symbols[offset]
			if !ok {
				return fmt.Errorf("section %s: missing symbol for map at offset %d", sec.Name, offset)
			}

			mapName := mapSym.Name
			if maps[mapName] != nil {
				return fmt.Errorf("section %v: map %v already exists", sec.Name, mapSym)
			}

			lr := io.LimitReader(r, int64(size))

			spec := MapSpec{
				Name: SanitizeName(mapName, -1),
			}
			switch {
			case binary.Read(lr, ec.ByteOrder, &spec.Type) != nil:
				return fmt.Errorf("map %s: missing type", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.KeySize) != nil:
				return fmt.Errorf("map %s: missing key size", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.ValueSize) != nil:
				return fmt.Errorf("map %s: missing value size", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.MaxEntries) != nil:
				return fmt.Errorf("map %s: missing max entries", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.Flags) != nil:
				return fmt.Errorf("map %s: missing flags", mapName)
			}

			if _, err := io.Copy(internal.DiscardZeroes{}, lr); err != nil {
				return fmt.Errorf("map %s: unknown and non-zero fields in definition", mapName)
			}

			if err := spec.clampPerfEventArraySize(); err != nil {
				return fmt.Errorf("map %s: %w", mapName, err)
			}

			maps[mapName] = &spec
		}
	}

	return nil
}

func (ec *elfCode) loadBTFMaps(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != btfMapSection {
			continue
		}

		if ec.btf == nil {
			return fmt.Errorf("missing BTF")
		}

		_, err := io.Copy(internal.DiscardZeroes{}, bufio.NewReader(sec.Open()))
		if err != nil {
			return fmt.Errorf("section %v: initializing BTF map definitions: %w", sec.Name, internal.ErrNotSupported)
		}

		var ds btf.Datasec
		if err := ec.btf.FindType(sec.Name, &ds); err != nil {
			return fmt.Errorf("cannot find section '%s' in BTF: %w", sec.Name, err)
		}

		for _, vs := range ds.Vars {
			v, ok := vs.Type.(*btf.Var)
			if !ok {
				return fmt.Errorf("section %v: unexpected type %s", sec.Name, vs.Type)
			}
			name := string(v.Name)

			if maps[name] != nil {
				return fmt.Errorf("section %v: map %s already exists", sec.Name, name)
			}

			mapStruct, ok := v.Type.(*btf.Struct)
			if !ok {
				return fmt.Errorf("expected struct, got %s", v.Type)
			}

			mapSpec, err := mapSpecFromBTF(name, mapStruct, false, ec.btf)
			if err != nil {
				return fmt.Errorf("map %v: %w", name, err)
			}

			if err := mapSpec.clampPerfEventArraySize(); err != nil {
				return fmt.Errorf("map %v: %w", name, err)
			}

			maps[name] = mapSpec
		}
	}

	return nil
}

// mapSpecFromBTF produces a MapSpec based on a btf.Struct def representing
// a BTF map definition. The name and spec arguments will be copied to the
// resulting MapSpec, and inner must be true on any resursive invocations.
func mapSpecFromBTF(name string, def *btf.Struct, inner bool, spec *btf.Spec) (*MapSpec, error) {

	var (
		key, value                 btf.Type
		keySize, valueSize         uint32
		mapType, flags, maxEntries uint32
		pinType                    PinType
		innerMapSpec               *MapSpec
		err                        error
	)

	for i, member := range def.Members {
		switch member.Name {
		case "type":
			mapType, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get type: %w", err)
			}

		case "map_flags":
			flags, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF map flags: %w", err)
			}

		case "max_entries":
			maxEntries, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF map max entries: %w", err)
			}

		case "key":
			if keySize != 0 {
				return nil, errors.New("both key and key_size given")
			}

			pk, ok := member.Type.(*btf.Pointer)
			if !ok {
				return nil, fmt.Errorf("key type is not a pointer: %T", member.Type)
			}

			key = pk.Target

			size, err := btf.Sizeof(pk.Target)
			if err != nil {
				return nil, fmt.Errorf("can't get size of BTF key: %w", err)
			}

			keySize = uint32(size)

		case "value":
			if valueSize != 0 {
				return nil, errors.New("both value and value_size given")
			}

			vk, ok := member.Type.(*btf.Pointer)
			if !ok {
				return nil, fmt.Errorf("value type is not a pointer: %T", member.Type)
			}

			value = vk.Target

			size, err := btf.Sizeof(vk.Target)
			if err != nil {
				return nil, fmt.Errorf("can't get size of BTF value: %w", err)
			}

			valueSize = uint32(size)

		case "key_size":
			// Key needs to be nil and keySize needs to be 0 for key_size to be
			// considered a valid member.
			if key != nil || keySize != 0 {
				return nil, errors.New("both key and key_size given")
			}

			keySize, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF key size: %w", err)
			}

		case "value_size":
			// Value needs to be nil and valueSize needs to be 0 for value_size to be
			// considered a valid member.
			if value != nil || valueSize != 0 {
				return nil, errors.New("both value and value_size given")
			}

			valueSize, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF value size: %w", err)
			}

		case "pinning":
			if inner {
				return nil, errors.New("inner maps can't be pinned")
			}

			pinning, err := uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get pinning: %w", err)
			}

			pinType = PinType(pinning)

		case "values":
			// The 'values' field in BTF map definitions is used for declaring map
			// value types that are references to other BPF objects, like other maps
			// or programs. It is always expected to be an array of pointers.
			if i != len(def.Members)-1 {
				return nil, errors.New("'values' must be the last member in a BTF map definition")
			}

			if valueSize != 0 && valueSize != 4 {
				return nil, errors.New("value_size must be 0 or 4")
			}
			valueSize = 4

			valueType, err := resolveBTFArrayMacro(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't resolve type of member 'values': %w", err)
			}

			switch t := valueType.(type) {
			case *btf.Struct:
				// The values member pointing to an array of structs means we're expecting
				// a map-in-map declaration.
				if MapType(mapType) != ArrayOfMaps && MapType(mapType) != HashOfMaps {
					return nil, errors.New("outer map needs to be an array or a hash of maps")
				}
				if inner {
					return nil, fmt.Errorf("nested inner maps are not supported")
				}

				// This inner map spec is used as a map template, but it needs to be
				// created as a traditional map before it can be used to do so.
				// libbpf names the inner map template '<outer_name>.inner', but we
				// opted for _inner to simplify validation logic. (dots only supported
				// on kernels 5.2 and up)
				// Pass the BTF spec from the parent object, since both parent and
				// child must be created from the same BTF blob (on kernels that support BTF).
				innerMapSpec, err = mapSpecFromBTF(name+"_inner", t, true, spec)
				if err != nil {
					return nil, fmt.Errorf("can't parse BTF map definition of inner map: %w", err)
				}

			default:
				return nil, fmt.Errorf("unsupported value type %q in 'values' field", t)
			}

		default:
			return nil, fmt.Errorf("unrecognized field %s in BTF map definition", member.Name)
		}
	}

	bm := btf.NewMap(spec, key, value)

	return &MapSpec{
		Name:       SanitizeName(name, -1),
		Type:       MapType(mapType),
		KeySize:    keySize,
		ValueSize:  valueSize,
		MaxEntries: maxEntries,
		Flags:      flags,
		BTF:        &bm,
		Pinning:    pinType,
		InnerMap:   innerMapSpec,
	}, nil
}

// uintFromBTF resolves the __uint macro, which is a pointer to a sized
// array, e.g. for int (*foo)[10], this function will return 10.
func uintFromBTF(typ btf.Type) (uint32, error) {
	ptr, ok := typ.(*btf.Pointer)
	if !ok {
		return 0, fmt.Errorf("not a pointer: %v", typ)
	}

	arr, ok := ptr.Target.(*btf.Array)
	if !ok {
		return 0, fmt.Errorf("not a pointer to array: %v", typ)
	}

	return arr.Nelems, nil
}

// resolveBTFArrayMacro resolves the __array macro, which declares an array
// of pointers to a given type. This function returns the target Type of
// the pointers in the array.
func resolveBTFArrayMacro(typ btf.Type) (btf.Type, error) {
	arr, ok := typ.(*btf.Array)
	if !ok {
		return nil, fmt.Errorf("not an array: %v", typ)
	}

	ptr, ok := arr.Type.(*btf.Pointer)
	if !ok {
		return nil, fmt.Errorf("not an array of pointers: %v", typ)
	}

	return ptr.Target, nil
}

func (ec *elfCode) loadDataSections(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != dataSection {
			continue
		}

		if sec.references == 0 {
			// Prune data sections which are not referenced by any
			// instructions.
			continue
		}

		if ec.btf == nil {
			return errors.New("data sections require BTF, make sure all consts are marked as static")
		}

		btfMap, err := ec.btf.Datasec(sec.Name)
		if err != nil {
			return err
		}

		data, err := sec.Data()
		if err != nil {
			return fmt.Errorf("data section %s: can't get contents: %w", sec.Name, err)
		}

		if uint64(len(data)) > math.MaxUint32 {
			return fmt.Errorf("data section %s: contents exceed maximum size", sec.Name)
		}

		mapSpec := &MapSpec{
			Name:       SanitizeName(sec.Name, -1),
			Type:       Array,
			KeySize:    4,
			ValueSize:  uint32(len(data)),
			MaxEntries: 1,
			Contents:   []MapKV{{uint32(0), data}},
			BTF:        btfMap,
		}

		switch sec.Name {
		case ".rodata":
			mapSpec.Flags = unix.BPF_F_RDONLY_PROG
			mapSpec.Freeze = true
		case ".bss":
			// The kernel already zero-initializes the map
			mapSpec.Contents = nil
		}

		maps[sec.Name] = mapSpec
	}
	return nil
}

func getProgType(sectionName string) (ProgramType, AttachType, uint32, string) {
	types := map[string]struct {
		progType   ProgramType
		attachType AttachType
		progFlags  uint32
	}{
		// From https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/tools/lib/bpf/libbpf.c
		"socket":                {SocketFilter, AttachNone, 0},
		"seccomp":               {SocketFilter, AttachNone, 0},
		"kprobe/":               {Kprobe, AttachNone, 0},
		"uprobe/":               {Kprobe, AttachNone, 0},
		"kretprobe/":            {Kprobe, AttachNone, 0},
		"uretprobe/":            {Kprobe, AttachNone, 0},
		"tracepoint/":           {TracePoint, AttachNone, 0},
		"raw_tracepoint/":       {RawTracepoint, AttachNone, 0},
		"raw_tp/":               {RawTracepoint, AttachNone, 0},
		"tp_btf/":               {Tracing, AttachTraceRawTp, 0},
		"xdp":                   {XDP, AttachNone, 0},
		"perf_event":            {PerfEvent, AttachNone, 0},
		"lwt_in":                {LWTIn, AttachNone, 0},
		"lwt_out":               {LWTOut, AttachNone, 0},
		"lwt_xmit":              {LWTXmit, AttachNone, 0},
		"lwt_seg6local":         {LWTSeg6Local, AttachNone, 0},
		"sockops":               {SockOps, AttachCGroupSockOps, 0},
		"sk_skb/stream_parser":  {SkSKB, AttachSkSKBStreamParser, 0},
		"sk_skb/stream_verdict": {SkSKB, AttachSkSKBStreamParser, 0},
		"sk_msg":                {SkMsg, AttachSkSKBStreamVerdict, 0},
		"lirc_mode2":            {LircMode2, AttachLircMode2, 0},
		"flow_dissector":        {FlowDissector, AttachFlowDissector, 0},
		"iter/":                 {Tracing, AttachTraceIter, 0},
		"fentry.s/":             {Tracing, AttachTraceFEntry, unix.BPF_F_SLEEPABLE},
		"fmod_ret.s/":           {Tracing, AttachModifyReturn, unix.BPF_F_SLEEPABLE},
		"fexit.s/":              {Tracing, AttachTraceFExit, unix.BPF_F_SLEEPABLE},
		"sk_lookup/":            {SkLookup, AttachSkLookup, 0},
		"lsm/":                  {LSM, AttachLSMMac, 0},
		"lsm.s/":                {LSM, AttachLSMMac, unix.BPF_F_SLEEPABLE},

		"cgroup_skb/ingress": {CGroupSKB, AttachCGroupInetIngress, 0},
		"cgroup_skb/egress":  {CGroupSKB, AttachCGroupInetEgress, 0},
		"cgroup/dev":         {CGroupDevice, AttachCGroupDevice, 0},
		"cgroup/skb":         {CGroupSKB, AttachNone, 0},
		"cgroup/sock":        {CGroupSock, AttachCGroupInetSockCreate, 0},
		"cgroup/post_bind4":  {CGroupSock, AttachCGroupInet4PostBind, 0},
		"cgroup/post_bind6":  {CGroupSock, AttachCGroupInet6PostBind, 0},
		"cgroup/bind4":       {CGroupSockAddr, AttachCGroupInet4Bind, 0},
		"cgroup/bind6":       {CGroupSockAddr, AttachCGroupInet6Bind, 0},
		"cgroup/connect4":    {CGroupSockAddr, AttachCGroupInet4Connect, 0},
		"cgroup/connect6":    {CGroupSockAddr, AttachCGroupInet6Connect, 0},
		"cgroup/sendmsg4":    {CGroupSockAddr, AttachCGroupUDP4Sendmsg, 0},
		"cgroup/sendmsg6":    {CGroupSockAddr, AttachCGroupUDP6Sendmsg, 0},
		"cgroup/recvmsg4":    {CGroupSockAddr, AttachCGroupUDP4Recvmsg, 0},
		"cgroup/recvmsg6":    {CGroupSockAddr, AttachCGroupUDP6Recvmsg, 0},
		"cgroup/sysctl":      {CGroupSysctl, AttachCGroupSysctl, 0},
		"cgroup/getsockopt":  {CGroupSockopt, AttachCGroupGetsockopt, 0},
		"cgroup/setsockopt":  {CGroupSockopt, AttachCGroupSetsockopt, 0},
		"classifier":         {SchedCLS, AttachNone, 0},
		"action":             {SchedACT, AttachNone, 0},
	}

	for prefix, t := range types {
		if !strings.HasPrefix(sectionName, prefix) {
			continue
		}

		if !strings.HasSuffix(prefix, "/") {
			return t.progType, t.attachType, t.progFlags, ""
		}

		return t.progType, t.attachType, t.progFlags, sectionName[len(prefix):]
	}

	return UnspecifiedProgram, AttachNone, 0, ""
}

func (ec *elfCode) loadRelocations(sec *elf.Section, symbols []elf.Symbol) (map[uint64]elf.Symbol, error) {
	rels := make(map[uint64]elf.Symbol)

	if sec.Entsize < 16 {
		return nil, fmt.Errorf("section %s: relocations are less than 16 bytes", sec.Name)
	}

	r := bufio.NewReader(sec.Open())
	for off := uint64(0); off < sec.Size; off += sec.Entsize {
		ent := io.LimitReader(r, int64(sec.Entsize))

		var rel elf.Rel64
		if binary.Read(ent, ec.ByteOrder, &rel) != nil {
			return nil, fmt.Errorf("can't parse relocation at offset %v", off)
		}

		symNo := int(elf.R_SYM64(rel.Info) - 1)
		if symNo >= len(symbols) {
			return nil, fmt.Errorf("offset %d: symbol %d doesn't exist", off, symNo)
		}

		symbol := symbols[symNo]
		rels[rel.Off] = symbol
	}

	return rels, nil
}
