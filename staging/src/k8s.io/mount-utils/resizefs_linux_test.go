//go:build linux
// +build linux

/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mount

import (
	"fmt"
	"testing"

	"k8s.io/utils/exec"
	fakeexec "k8s.io/utils/exec/testing"
)

func TestGetFileSystemSize(t *testing.T) {
	cmdOutputSuccessXfs :=
		`
	statfs.f_bsize = 4096
	statfs.f_blocks = 1832448
	statfs.f_bavail = 1822366
	statfs.f_files = 3670016
	statfs.f_ffree = 3670012
	statfs.f_flags = 0x1020
	geom.bsize = 4096
	geom.agcount = 4
	geom.agblocks = 458752
	geom.datablocks = 1835008
	geom.rtblocks = 0
	geom.rtextents = 0
	geom.rtextsize = 1
	geom.sunit = 0
	geom.swidth = 0
	counts.freedata = 1822372
	counts.freertx = 0
	counts.freeino = 61
	counts.allocino = 64
`
	cmdOutputNoDataXfs :=
		`
	statfs.f_bsize = 4096
	statfs.f_blocks = 1832448
	statfs.f_bavail = 1822366
	statfs.f_files = 3670016
	statfs.f_ffree = 3670012
	statfs.f_flags = 0x1020
	geom.agcount = 4
	geom.agblocks = 458752
	geom.rtblocks = 0
	geom.rtextents = 0
	geom.rtextsize = 1
	geom.sunit = 0
	geom.swidth = 0
	counts.freedata = 1822372
	counts.freertx = 0
	counts.freeino = 61
	counts.allocino = 64
`
	cmdOutputSuccessExt4 :=
		`
Filesystem volume name:   cloudimg-rootfs
Last mounted on:          /
Filesystem UUID:          testUUID
Filesystem magic number:  0xEF53
Filesystem revision #:    1 (dynamic)
Filesystem features:      has_journal ext_attr resize_inode dir_index filetype needs_recovery extent 64bit
Default mount options:    user_xattr acl
Filesystem state:         clean
Errors behavior:          Continue
Filesystem OS type:       Linux
Inode count:              3840000
Block count:              5242880
Reserved block count:     0
Free blocks:              5514413
Free inodes:              3677492
First block:              0
Block size:               4096
Fragment size:            4096
Group descriptor size:    64
Reserved GDT blocks:      252
Blocks per group:         32768
Fragments per group:      32768
Inodes per group:         16000
Inode blocks per group:   1000
Flex block group size:    16
Mount count:              2
Maximum mount count:      -1
Check interval:           0 (<none>)
Lifetime writes:          180 GB
Reserved blocks uid:      0 (user root)
Reserved blocks gid:      0 (group root)
First inode:              11
Inode size:	              256
Required extra isize:     32
Desired extra isize:      32
Journal inode:            8
Default directory hash:   half_md4
Directory Hash Seed:      Test Hashing
Journal backup:           inode blocks
Checksum type:            crc32c
Checksum:                 0x57705f62
Journal features:         journal_incompat_revoke journal_64bit journal_checksum_v3
Journal size:             64M
Journal length:           16384
Journal sequence:         0x00037109
Journal start:            1
Journal checksum type:    crc32c
Journal checksum:         0xb7df3c6e
`
	cmdOutputNoDataExt4 :=
		`Filesystem volume name:   cloudimg-rootfs
Last mounted on:          /
Filesystem UUID:          testUUID
Filesystem magic number:  0xEF53
Filesystem revision #:    1 (dynamic)
Filesystem features:      has_journal ext_attr resize_inode dir_index filetype needs_recovery extent 64bit
Default mount options:    user_xattr acl
Filesystem state:         clean
Errors behavior:          Continue
Filesystem OS type:       Linux
Inode count:              3840000
Reserved block count:     0
Free blocks:              5514413
Free inodes:              3677492
First block:              0
Fragment size:            4096
Group descriptor size:    64
Reserved GDT blocks:      252
Blocks per group:         32768
Fragments per group:      32768
Inodes per group:         16000
Inode blocks per group:   1000
Flex block group size:    16
Mount count:              2
Maximum mount count:      -1
Check interval:           0 (<none>)
Lifetime writes:          180 GB
Reserved blocks uid:      0 (user root)
Reserved blocks gid:      0 (group root)
First inode:              11
Inode size:	              256
Required extra isize:     32
Desired extra isize:      32
Journal inode:            8
Default directory hash:   half_md4
Directory Hash Seed:      Test Hashing
Journal backup:           inode blocks
Checksum type:            crc32c
Checksum:                 0x57705f62
Journal features:         journal_incompat_revoke journal_64bit journal_checksum_v3
Journal size:             64M
Journal length:           16384
Journal sequence:         0x00037109
Journal start:            1
Journal checksum type:    crc32c
Journal checksum:         0xb7df3c6e
`
	cmdOutputSuccessBtrfs :=
		`superblock: bytenr=65536, device=/dev/loop0
---------------------------------------------------------
csum_type               0 (crc32c)
csum_size               4
csum                    0x31693b11 [match]
bytenr                  65536
flags                   0x1
                        ( WRITTEN )
magic                   _BHRfS_M [match]
fsid                    3f53c8f7-3c57-4185-bf1d-b305b42cce97
metadata_uuid           3f53c8f7-3c57-4185-bf1d-b305b42cce97
label
generation              7
root                    30441472
sys_array_size          129
chunk_root_generation   6
root_level              0
chunk_root              22036480
chunk_root_level        0
log_root                0
log_root_transid        0
log_root_level          0
total_bytes             1048576000
bytes_used              147456
sectorsize              4096
nodesize                16384
leafsize (deprecated)   16384
stripesize              4096
root_dir                6
num_devices             1
compat_flags            0x0
compat_ro_flags         0x3
                        ( FREE_SPACE_TREE |
                          FREE_SPACE_TREE_VALID )
incompat_flags          0x341
                        ( MIXED_BACKREF |
                          EXTENDED_IREF |
                          SKINNY_METADATA |
                          NO_HOLES )
cache_generation        0
uuid_tree_generation    7
dev_item.uuid           987c8423-fba3-4168-9892-560a116feb81
dev_item.fsid           3f53c8f7-3c57-4185-bf1d-b305b42cce97 [match]
dev_item.type           0
dev_item.total_bytes    1048576000
dev_item.bytes_used     130023424
dev_item.io_align       4096
dev_item.io_width       4096
dev_item.sector_size    4096
dev_item.devid          1
dev_item.dev_group      0
dev_item.seek_speed     0
dev_item.bandwidth      0
dev_item.generation     0
sys_chunk_array[2048]:
        item 0 key (FIRST_CHUNK_TREE CHUNK_ITEM 22020096)
                length 8388608 owner 2 stripe_len 65536 type SYSTEM|DUP
                io_align 65536 io_width 65536 sector_size 4096
                num_stripes 2 sub_stripes 1
                        stripe 0 devid 1 offset 22020096
                        dev_uuid 987c8423-fba3-4168-9892-560a116feb81
                        stripe 1 devid 1 offset 30408704
                        dev_uuid 987c8423-fba3-4168-9892-560a116feb81
backup_roots[4]:
        backup 0:
                backup_tree_root:       30441472        gen: 5  level: 0
                backup_chunk_root:      22020096        gen: 5  level: 0
                backup_extent_root:     30474240        gen: 5  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30457856        gen: 5  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 1:
                backup_tree_root:       30588928        gen: 6  level: 0
                backup_chunk_root:      22036480        gen: 6  level: 0
                backup_extent_root:     30408704        gen: 6  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30556160        gen: 6  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 2:
                backup_tree_root:       30441472        gen: 7  level: 0
                backup_chunk_root:      22036480        gen: 6  level: 0
                backup_extent_root:     30474240        gen: 7  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30457856        gen: 7  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 3:
                backup_tree_root:       30408704        gen: 4  level: 0
                backup_chunk_root:      1064960 gen: 4  level: 0
                backup_extent_root:     5341184 gen: 4  level: 0
                backup_fs_root:         5324800 gen: 3  level: 0
                backup_dev_root:        5242880 gen: 4  level: 0
                backup_csum_root:       1130496 gen: 1  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      114688
                backup_num_devices:     1

`
	cmdOutputNoDataBtrfs :=
		`superblock: bytenr=65536, device=/dev/loop0
---------------------------------------------------------
csum_type               0 (crc32c)
csum_size               4
csum                    0x31693b11 [match]
bytenr                  65536
flags                   0x1
                        ( WRITTEN )
magic                   _BHRfS_M [match]
fsid                    3f53c8f7-3c57-4185-bf1d-b305b42cce97
metadata_uuid           3f53c8f7-3c57-4185-bf1d-b305b42cce97
label
generation              7
root                    30441472
sys_array_size          129
chunk_root_generation   6
root_level              0
chunk_root              22036480
chunk_root_level        0
log_root                0
log_root_transid        0
log_root_level          0
bytes_used              147456
nodesize                16384
leafsize (deprecated)   16384
stripesize              4096
root_dir                6
num_devices             1
compat_flags            0x0
compat_ro_flags         0x3
                        ( FREE_SPACE_TREE |
                          FREE_SPACE_TREE_VALID )
incompat_flags          0x341
                        ( MIXED_BACKREF |
                          EXTENDED_IREF |
                          SKINNY_METADATA |
                          NO_HOLES )
cache_generation        0
uuid_tree_generation    7
dev_item.uuid           987c8423-fba3-4168-9892-560a116feb81
dev_item.fsid           3f53c8f7-3c57-4185-bf1d-b305b42cce97 [match]
dev_item.type           0
dev_item.total_bytes    1048576000
dev_item.bytes_used     130023424
dev_item.io_align       4096
dev_item.io_width       4096
dev_item.sector_size    4096
dev_item.devid          1
dev_item.dev_group      0
dev_item.seek_speed     0
dev_item.bandwidth      0
dev_item.generation     0
sys_chunk_array[2048]:
        item 0 key (FIRST_CHUNK_TREE CHUNK_ITEM 22020096)
                length 8388608 owner 2 stripe_len 65536 type SYSTEM|DUP
                io_align 65536 io_width 65536 sector_size 4096
                num_stripes 2 sub_stripes 1
                        stripe 0 devid 1 offset 22020096
                        dev_uuid 987c8423-fba3-4168-9892-560a116feb81
                        stripe 1 devid 1 offset 30408704
                        dev_uuid 987c8423-fba3-4168-9892-560a116feb81
backup_roots[4]:
        backup 0:
                backup_tree_root:       30441472        gen: 5  level: 0
                backup_chunk_root:      22020096        gen: 5  level: 0
                backup_extent_root:     30474240        gen: 5  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30457856        gen: 5  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 1:
                backup_tree_root:       30588928        gen: 6  level: 0
                backup_chunk_root:      22036480        gen: 6  level: 0
                backup_extent_root:     30408704        gen: 6  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30556160        gen: 6  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 2:
                backup_tree_root:       30441472        gen: 7  level: 0
                backup_chunk_root:      22036480        gen: 6  level: 0
                backup_extent_root:     30474240        gen: 7  level: 0
                backup_fs_root:         30425088        gen: 5  level: 0
                backup_dev_root:        30457856        gen: 7  level: 0
                backup_csum_root:       30490624        gen: 5  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      147456
                backup_num_devices:     1

        backup 3:
                backup_tree_root:       30408704        gen: 4  level: 0
                backup_chunk_root:      1064960 gen: 4  level: 0
                backup_extent_root:     5341184 gen: 4  level: 0
                backup_fs_root:         5324800 gen: 3  level: 0
                backup_dev_root:        5242880 gen: 4  level: 0
                backup_csum_root:       1130496 gen: 1  level: 0
                backup_total_bytes:     1048576000
                backup_bytes_used:      114688
                backup_num_devices:     1

`
	testcases := []struct {
		name        string
		devicePath  string
		blocksize   uint64
		blockCount  uint64
		cmdOutput   string
		expectError bool
		fsType      string
	}{
		{
			name:        "success parse xfs info",
			devicePath:  "/dev/test1",
			blocksize:   4096,
			blockCount:  1835008,
			cmdOutput:   cmdOutputSuccessXfs,
			expectError: false,
			fsType:      "xfs",
		},
		{
			name:        "block size not present - xfs",
			devicePath:  "/dev/test1",
			blocksize:   0,
			blockCount:  0,
			cmdOutput:   cmdOutputNoDataXfs,
			expectError: true,
			fsType:      "xfs",
		},
		{
			name:        "success parse ext info",
			devicePath:  "/dev/test1",
			blocksize:   4096,
			blockCount:  5242880,
			cmdOutput:   cmdOutputSuccessExt4,
			expectError: false,
			fsType:      "ext4",
		},
		{
			name:        "block size not present - ext4",
			devicePath:  "/dev/test1",
			blocksize:   0,
			blockCount:  0,
			cmdOutput:   cmdOutputNoDataExt4,
			expectError: true,
			fsType:      "ext4",
		},
		{
			name:        "success parse btrfs info",
			devicePath:  "/dev/test1",
			blocksize:   4096,
			blockCount:  256000,
			cmdOutput:   cmdOutputSuccessBtrfs,
			expectError: false,
			fsType:      "btrfs",
		},
		{
			name:        "block size not present - btrfs",
			devicePath:  "/dev/test1",
			blocksize:   0,
			blockCount:  0,
			cmdOutput:   cmdOutputNoDataBtrfs,
			expectError: true,
			fsType:      "btrfs",
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			fcmd := fakeexec.FakeCmd{
				CombinedOutputScript: []fakeexec.FakeAction{
					func() ([]byte, []byte, error) { return []byte(test.cmdOutput), nil, nil },
				},
			}
			fexec := fakeexec.FakeExec{
				CommandScript: []fakeexec.FakeCommandAction{
					func(cmd string, args ...string) exec.Cmd {
						return fakeexec.InitFakeCmd(&fcmd, cmd, args...)
					},
				},
			}
			resizefs := ResizeFs{exec: &fexec}

			var blockSize uint64
			var fsSize uint64
			var err error
			switch test.fsType {
			case "xfs":
				blockSize, fsSize, err = resizefs.getXFSSize(test.devicePath)
			case "ext4":
				blockSize, fsSize, err = resizefs.getExtSize(test.devicePath)
			case "btrfs":
				blockSize, fsSize, err = resizefs.getBtrfsSize(test.devicePath)
			}

			if blockSize != test.blocksize {
				t.Fatalf("Parse wrong block size value, expect %d, but got %d", test.blocksize, blockSize)
			}
			if fsSize != test.blocksize*test.blockCount {
				t.Fatalf("Parse wrong fs size value, expect %d, but got %d", test.blocksize*test.blockCount, fsSize)
			}
			if !test.expectError && err != nil {
				t.Fatalf("Expect no error but got %v", err)
			}
		})
	}
}

func TestNeedResize(t *testing.T) {
	testcases := []struct {
		name            string
		devicePath      string
		deviceMountPath string
		readonly        string
		deviceSize      string
		extSize         string
		cmdOutputFsType string
		expectError     bool
		expectResult    bool
	}{
		{
			name:            "True",
			devicePath:      "/dev/test1",
			deviceMountPath: "/mnt/test1",
			readonly:        "0",
			deviceSize:      "2048",
			cmdOutputFsType: "TYPE=ext3",
			extSize:         "20",
			expectError:     false,
			expectResult:    true,
		},
		{
			name:            "False - needed by size but fs is readonly",
			devicePath:      "/dev/test1",
			deviceMountPath: "/mnt/test1",
			readonly:        "1",
			deviceSize:      "2048",
			cmdOutputFsType: "TYPE=ext3",
			extSize:         "20",
			expectError:     false,
			expectResult:    false,
		},
		{
			name:            "False - Not needed by size",
			devicePath:      "/dev/test1",
			deviceMountPath: "/mnt/test1",
			readonly:        "0",
			deviceSize:      "20",
			cmdOutputFsType: "TYPE=ext3",
			extSize:         "2048",
			expectError:     false,
			expectResult:    false,
		},
		{
			name:            "False - Unsupported fs type",
			devicePath:      "/dev/test1",
			deviceMountPath: "/mnt/test1",
			readonly:        "0",
			deviceSize:      "2048",
			extSize:         "1",
			cmdOutputFsType: "TYPE=ntfs",
			expectError:     true,
			expectResult:    false,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			fcmd := fakeexec.FakeCmd{
				CombinedOutputScript: []fakeexec.FakeAction{
					func() ([]byte, []byte, error) { return []byte(test.readonly), nil, nil },
					func() ([]byte, []byte, error) { return []byte(test.deviceSize), nil, nil },
					func() ([]byte, []byte, error) { return []byte(test.cmdOutputFsType), nil, nil },
					func() ([]byte, []byte, error) {
						return []byte(fmt.Sprintf("block size: %s\nblock count: 1", test.extSize)), nil, nil
					},
				},
			}
			fexec := fakeexec.FakeExec{
				CommandScript: []fakeexec.FakeCommandAction{
					func(cmd string, args ...string) exec.Cmd { return fakeexec.InitFakeCmd(&fcmd, cmd, args...) },
					func(cmd string, args ...string) exec.Cmd { return fakeexec.InitFakeCmd(&fcmd, cmd, args...) },
					func(cmd string, args ...string) exec.Cmd { return fakeexec.InitFakeCmd(&fcmd, cmd, args...) },
					func(cmd string, args ...string) exec.Cmd { return fakeexec.InitFakeCmd(&fcmd, cmd, args...) },
				},
			}
			resizefs := ResizeFs{exec: &fexec}

			needResize, err := resizefs.NeedResize(test.devicePath, test.deviceMountPath)
			if !test.expectError && err != nil {
				t.Fatalf("Expect no error but got %v", err)
			}
			if needResize != test.expectResult {
				t.Fatalf("Expect result is %v but got %v", test.expectResult, needResize)
			}
		})
	}
}
