/*
Copyright 2014 The Kubernetes Authors.

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

package validation

import (
	"strings"
	"testing"
)

func TestIsDNS1123Label(t *testing.T) {
	goodValues := []string{
		"a", "ab", "abc", "a1", "a-1", "a--1--2--b",
		"0", "01", "012", "1a", "1-a", "1--a--b--2",
		strings.Repeat("a", 63),
	}
	for _, val := range goodValues {
		if msgs := IsDNS1123Label(val); len(msgs) != 0 {
			t.Errorf("expected true for '%s': %v", val, msgs)
		}
	}

	badValues := []string{
		"", "A", "ABC", "aBc", "A1", "A-1", "1-A",
		"-", "a-", "-a", "1-", "-1",
		"_", "a_", "_a", "a_b", "1_", "_1", "1_2",
		".", "a.", ".a", "a.b", "1.", ".1", "1.2",
		" ", "a ", " a", "a b", "1 ", " 1", "1 2",
		strings.Repeat("a", 64),
	}
	for _, val := range badValues {
		if msgs := IsDNS1123Label(val); len(msgs) == 0 {
			t.Errorf("expected false for '%s'", val)
		}
	}
}

func TestIsDNS1123Subdomain(t *testing.T) {
	goodValues := []string{
		"a", "ab", "abc", "a1", "a-1", "a--1--2--b",
		"0", "01", "012", "1a", "1-a", "1--a--b--2",
		"a.a", "ab.a", "abc.a", "a1.a", "a-1.a", "a--1--2--b.a",
		"a.1", "ab.1", "abc.1", "a1.1", "a-1.1", "a--1--2--b.1",
		"0.a", "01.a", "012.a", "1a.a", "1-a.a", "1--a--b--2",
		"0.1", "01.1", "012.1", "1a.1", "1-a.1", "1--a--b--2.1",
		"a.b.c.d.e", "aa.bb.cc.dd.ee", "1.2.3.4.5", "11.22.33.44.55",
		strings.Repeat("a", 253),
	}
	for _, val := range goodValues {
		if msgs := IsDNS1123Subdomain(val); len(msgs) != 0 {
			t.Errorf("expected true for '%s': %v", val, msgs)
		}
	}

	badValues := []string{
		"", "A", "ABC", "aBc", "A1", "A-1", "1-A",
		"-", "a-", "-a", "1-", "-1",
		"_", "a_", "_a", "a_b", "1_", "_1", "1_2",
		".", "a.", ".a", "a..b", "1.", ".1", "1..2",
		" ", "a ", " a", "a b", "1 ", " 1", "1 2",
		"A.a", "aB.a", "ab.A", "A1.a", "a1.A",
		"A.1", "aB.1", "A1.1", "1A.1",
		"0.A", "01.A", "012.A", "1A.a", "1a.A",
		"A.B.C.D.E", "AA.BB.CC.DD.EE", "a.B.c.d.e", "aa.bB.cc.dd.ee",
		"a@b", "a,b", "a_b", "a;b",
		"a:b", "a%b", "a?b", "a$b",
		strings.Repeat("a", 254),
	}
	for _, val := range badValues {
		if msgs := IsDNS1123Subdomain(val); len(msgs) == 0 {
			t.Errorf("expected false for '%s'", val)
		}
	}
}

func TestIsDNS952Label(t *testing.T) {
	goodValues := []string{
		"a", "ab", "abc", "a1", "a-1", "a--1--2--b",
		strings.Repeat("a", 24),
	}
	for _, val := range goodValues {
		if msgs := IsDNS952Label(val); len(msgs) != 0 {
			t.Errorf("expected true for '%s': %v", val, msgs)
		}
	}

	badValues := []string{
		"0", "01", "012", "1a", "1-a", "1--a--b--2",
		"", "A", "ABC", "aBc", "A1", "A-1", "1-A",
		"-", "a-", "-a", "1-", "-1",
		"_", "a_", "_a", "a_b", "1_", "_1", "1_2",
		".", "a.", ".a", "a.b", "1.", ".1", "1.2",
		" ", "a ", " a", "a b", "1 ", " 1", "1 2",
		strings.Repeat("a", 25),
	}
	for _, val := range badValues {
		if msgs := IsDNS952Label(val); len(msgs) == 0 {
			t.Errorf("expected false for '%s'", val)
		}
	}
}

func TestIsCIdentifier(t *testing.T) {
	goodValues := []string{
		"a", "ab", "abc", "a1", "_a", "a_", "a_b", "a_1", "a__1__2__b", "__abc_123",
		"A", "AB", "AbC", "A1", "_A", "A_", "A_B", "A_1", "A__1__2__B", "__123_ABC",
	}
	for _, val := range goodValues {
		if msgs := IsCIdentifier(val); len(msgs) != 0 {
			t.Errorf("expected true for '%s': %v", val, msgs)
		}
	}

	badValues := []string{
		"", "1", "123", "1a",
		"-", "a-", "-a", "1-", "-1", "1_", "1_2",
		".", "a.", ".a", "a.b", "1.", ".1", "1.2",
		" ", "a ", " a", "a b", "1 ", " 1", "1 2",
		"#a#",
	}
	for _, val := range badValues {
		if msgs := IsCIdentifier(val); len(msgs) == 0 {
			t.Errorf("expected false for '%s'", val)
		}
	}
}

func TestIsValidPortNum(t *testing.T) {
	goodValues := []int{1, 2, 1000, 16384, 32768, 65535}
	for _, val := range goodValues {
		if msgs := IsValidPortNum(val); len(msgs) != 0 {
			t.Errorf("expected true for %d, got %v", val, msgs)
		}
	}

	badValues := []int{0, -1, 65536, 100000}
	for _, val := range badValues {
		if msgs := IsValidPortNum(val); len(msgs) == 0 {
			t.Errorf("expected false for %d", val)
		}
	}
}

func TestIsValidGroupId(t *testing.T) {
	goodValues := []int64{0, 1, 1000, 65535, 2147483647}
	for _, val := range goodValues {
		if !IsValidGroupId(val) {
			t.Errorf("expected true for '%d'", val)
		}
	}

	badValues := []int64{-1, -1003, 2147483648, 4147483647}
	for _, val := range badValues {
		if IsValidGroupId(val) {
			t.Errorf("expected false for '%d'", val)
		}
	}
}

func TestIsValidUserId(t *testing.T) {
	goodValues := []int64{0, 1, 1000, 65535, 2147483647}
	for _, val := range goodValues {
		if !IsValidUserId(val) {
			t.Errorf("expected true for '%d'", val)
		}
	}

	badValues := []int64{-1, -1003, 2147483648, 4147483647}
	for _, val := range badValues {
		if IsValidUserId(val) {
			t.Errorf("expected false for '%d'", val)
		}
	}
}

func TestIsValidPortName(t *testing.T) {
	goodValues := []string{"telnet", "re-mail-ck", "pop3", "a", "a-1", "1-a", "a-1-b-2-c", "1-a-2-b-3"}
	for _, val := range goodValues {
		if msgs := IsValidPortName(val); len(msgs) != 0 {
			t.Errorf("expected true for %q: %v", val, msgs)
		}
	}

	badValues := []string{"longerthan15characters", "", strings.Repeat("a", 16), "12345", "1-2-3-4", "-begin", "end-", "two--hyphens", "whois++"}
	for _, val := range badValues {
		if msgs := IsValidPortName(val); len(msgs) == 0 {
			t.Errorf("expected false for %q", val)
		}
	}
}

func TestIsQualifiedName(t *testing.T) {
	successCases := []string{
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"1234",
		"simple/simple",
		"now-with-dashes/simple",
		"now-with-dashes/now-with-dashes",
		"now.with.dots/simple",
		"now-with.dashes-and.dots/simple",
		"1-num.2-num/3-num",
		"1234/5678",
		"1.2.3.4/5678",
		"Uppercase_Is_OK_123",
		"example.com/Uppercase_Is_OK_123",
		strings.Repeat("a", 63),
		strings.Repeat("a", 253) + "/" + strings.Repeat("b", 63),
	}
	for i := range successCases {
		if errs := IsQualifiedName(successCases[i]); len(errs) != 0 {
			t.Errorf("case[%d]: %q: expected success: %v", i, successCases[i], errs)
		}
	}

	errorCases := []string{
		"nospecialchars%^=@",
		"cantendwithadash-",
		"-cantstartwithadash-",
		"only/one/slash",
		"Example.com/abc",
		"example_com/abc",
		"example.com/",
		"/simple",
		strings.Repeat("a", 64),
		strings.Repeat("a", 254) + "/abc",
	}
	for i := range errorCases {
		if errs := IsQualifiedName(errorCases[i]); len(errs) == 0 {
			t.Errorf("case[%d]: %q: expected failure", i, errorCases[i])
		}
	}
}

func TestIsValidLabelValue(t *testing.T) {
	successCases := []string{
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"end-with-num-1",
		"1234",                  // only num
		strings.Repeat("a", 63), // to the limit
		"", // empty value
	}
	for i := range successCases {
		if errs := IsValidLabelValue(successCases[i]); len(errs) != 0 {
			t.Errorf("case %s expected success: %v", successCases[i], errs)
		}
	}

	errorCases := []string{
		"nospecialchars%^=@",
		"Tama-nui-te-rā.is.Māori.sun",
		"\\backslashes\\are\\bad",
		"-starts-with-dash",
		"ends-with-dash-",
		".starts.with.dot",
		"ends.with.dot.",
		strings.Repeat("a", 64), // over the limit
	}
	for i := range errorCases {
		if errs := IsValidLabelValue(errorCases[i]); len(errs) == 0 {
			t.Errorf("case[%d] expected failure", i)
		}
	}
}

func TestIsValidIP(t *testing.T) {
	goodValues := []string{
		"::1",
		"2a00:79e0:2:0:f1c3:e797:93c1:df80",
		"::",
		"2001:4860:4860::8888",
		"::fff:1.1.1.1",
		"1.1.1.1",
		"1.1.1.01",
		"255.0.0.1",
		"1.0.0.0",
		"0.0.0.0",
	}
	for _, val := range goodValues {
		if msgs := IsValidIP(val); len(msgs) != 0 {
			t.Errorf("expected true for %q: %v", val, msgs)
		}
	}

	badValues := []string{
		"[2001:db8:0:1]:80",
		"myhost.mydomain",
		"-1.0.0.0",
		"[2001:db8:0:1]",
		"a",
	}
	for _, val := range badValues {
		if msgs := IsValidIP(val); len(msgs) == 0 {
			t.Errorf("expected false for %q", val)
		}
	}
}

func TestIsHTTPHeaderName(t *testing.T) {
	goodValues := []string{
		// Common ones
		"Accept-Encoding", "Host", "If-Modified-Since", "X-Forwarded-For",
		// Weirdo, but still conforming names
		"a", "ab", "abc", "a1", "-a", "a-", "a-b", "a-1", "a--1--2--b", "--abc-123",
		"A", "AB", "AbC", "A1", "-A", "A-", "A-B", "A-1", "A--1--2--B", "--123-ABC",
	}
	for _, val := range goodValues {
		if !IsHTTPHeaderName(val) {
			t.Errorf("expected true for '%s'", val)
		}
	}

	badValues := []string{
		"Host:", "X-Forwarded-For:", "X-@Home",
		"", "_", "a_", "_a", "1_", "1_2", ".", "a.", ".a", "a.b", "1.", ".1", "1.2",
		" ", "a ", " a", "a b", "1 ", " 1", "1 2", "#a#", "^", ",", ";", "=", "<",
		"?", "@", "{",
	}
	for _, val := range badValues {
		if IsHTTPHeaderName(val) {
			t.Errorf("expected false for '%s'", val)
		}
	}
}
