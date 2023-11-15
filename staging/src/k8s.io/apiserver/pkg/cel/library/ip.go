/*
Copyright 2023 The Kubernetes Authors.

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

package library

import (
	"fmt"
	"net/netip"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"

	apiservercel "k8s.io/apiserver/pkg/cel"
)

// IP provides a CEL function library extension of IP address parsing functions.
//
// ip
//
// Converts a string to an IP address or results in an error if the string is not a valid IP address.
// The IP address must be an IPv4 or IPv6 address.
// IPv4-mapped IPv6 addresses (e.g. ::ffff:1.2.3.4) are not allowed.
// IP addresses with zones (e.g. fe80::1%eth0) are not allowed.
// Leading zeros in IPv4 address octets are not allowed.
//
//	ip(<string>) <IPAddr>
//
// Examples:
//
//	ip('127.0.0.1') // returns an IPv4 address
//	ip('::1') // returns an IPv6 address
//	ip('127.0.0.256') // error
//	ip(':::1') // error
//
// isIP
//
// Returns true if a string is a valid IP address.
// The IP address must be an IPv4 or IPv6 address.
// IPv4-mapped IPv6 addresses (e.g. ::ffff:1.2.3.4) are not allowed.
// IP addresses with zones (e.g. fe80::1%eth0) are not allowed.
// Leading zeros in IPv4 address octets are not allowed.
//
//	isIP(<string>) <bool>
//
// Examples:
//
//	isIP('127.0.0.1') // returns true
//	isIP('::1') // returns true
//	isIP('127.0.0.256') // returns false
//	isIP(':::1') // returns false
//
// ip.isCanonical
//
// Returns true if the IP address is in its canonical form.
// There is exactly one canonical form for every IP address, so fields containing
// IPs in canonical form can just be treated as strings when checking for equality or uniqueness.
//
//	ip.isCanonical(<string>) <bool>
//
// Examples:
//
//	ip.isCanonical('127.0.0.1') // returns true; all valid IPv4 addresses are canonical
//	ip.isCanonical('2001:db8::abcd') // returns true
//	ip.isCanonical('2001:DB8::ABCD') // returns false
//	ip.isCanonical('2001:db8::0:0:0:abcd') // returns false
//
// family
//
// - family: returns the IP addresses' family (IPv4 or IPv6) as an integer, either '4' or '6'.
//
// Examples:
//
// ip('127.0.0.1').family() // returns '4”
// ip('::1').family() // returns '6'
// ip('127.0.0.1').family() == 4 // returns true
// ip('::1').family() == 6 // returns true
func IP() cel.EnvOption {
	return cel.Lib(ipLib)
}

var ipLib = &ip{}

type ip struct{}

func (*ip) LibraryName() string {
	return "net.ip"
}

var ipLibraryDecls = map[string][]cel.FunctionOpt{
	"ip": {
		cel.Overload("string_to_ip", []*cel.Type{cel.StringType}, apiservercel.IPType,
			cel.UnaryBinding(stringToIP)),
	},
	"family": {
		cel.MemberOverload("ip_family", []*cel.Type{apiservercel.IPType}, cel.IntType,
			cel.UnaryBinding(family)),
	},
	"ip.isCanonical": {
		cel.Overload("ip_is_canonical", []*cel.Type{cel.StringType}, cel.BoolType,
			cel.UnaryBinding(ipIsCanonical)),
	},
	"isIP": {
		cel.Overload("is_ip", []*cel.Type{cel.StringType}, cel.BoolType,
			cel.UnaryBinding(isIP)),
	},
	"string": {
		cel.Overload("ip_to_string", []*cel.Type{apiservercel.IPType}, cel.StringType,
			cel.UnaryBinding(ipToString)),
	},
}

func (*ip) CompileOptions() []cel.EnvOption {
	options := []cel.EnvOption{cel.Types(apiservercel.IPType),
		cel.Variable(apiservercel.IPType.TypeName(), types.NewTypeTypeWithParam(apiservercel.IPType)),
	}
	for name, overloads := range ipLibraryDecls {
		options = append(options, cel.Function(name, overloads...))
	}
	return options
}

func (*ip) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{}
}

func stringToIP(arg ref.Val) ref.Val {
	s, ok := arg.Value().(string)
	if !ok {
		return types.MaybeNoSuchOverloadErr(arg)
	}

	addr, err := parseIPAddr(s)
	if err != nil {
		// Don't add context, we control the error message already.
		return types.NewErr("%v", err)
	}

	return apiservercel.IP{
		Addr: addr,
	}
}

func ipToString(arg ref.Val) ref.Val {
	ip, ok := arg.(apiservercel.IP)
	if !ok {
		return types.MaybeNoSuchOverloadErr(arg)
	}

	return types.String(ip.Addr.String())
}

func family(arg ref.Val) ref.Val {
	ip, ok := arg.(apiservercel.IP)
	if !ok {
		return types.MaybeNoSuchOverloadErr(arg)
	}

	switch {
	case ip.Addr.Is4():
		return types.Int(4)
	case ip.Addr.Is6():
		return types.Int(6)
	default:
		return types.NewErr("IP address %q is not an IPv4 or IPv6 address", ip.Addr.String())
	}
}

func ipIsCanonical(arg ref.Val) ref.Val {
	s, ok := arg.Value().(string)
	if !ok {
		return types.MaybeNoSuchOverloadErr(arg)
	}

	addr, err := parseIPAddr(s)
	if err != nil {
		// Don't add context, we control the error message already.
		return types.NewErr("%v", err)
	}

	// Addr.String() always returns the canonical form of the IP address.
	// Therefore comparing this with the original string representation
	// will tell us if the IP address is in its canonical form.
	return types.Bool(addr.String() == s)
}

func isIP(arg ref.Val) ref.Val {
	s, ok := arg.Value().(string)
	if !ok {
		return types.MaybeNoSuchOverloadErr(arg)
	}

	_, err := parseIPAddr(s)
	return types.Bool(err == nil)
}

// parseIPAddr parses a string into an IP address.
// We use this function to parse IP addresses in the CEL library
// so that we can share the common logic of rejecting IP addresses
// that contain zones or are IPv4-mapped IPv6 addresses.
func parseIPAddr(raw string) (netip.Addr, error) {
	addr, err := netip.ParseAddr(raw)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("IP Address %q parse error during conversion from string: %v", raw, err)
	}

	if addr.Zone() != "" {
		return netip.Addr{}, fmt.Errorf("IP address %q with zone value is not allowed", raw)
	}

	if addr.Is4In6() {
		return netip.Addr{}, fmt.Errorf("IPv4-mapped IPv6 address %q is not allowed", raw)
	}

	return addr, nil
}
