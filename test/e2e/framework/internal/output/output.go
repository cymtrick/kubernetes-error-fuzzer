/*
Copyright 2022 The Kubernetes Authors.

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

package output

import (
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"

	"k8s.io/kubernetes/test/e2e/framework"
)

func TestGinkgoOutput(t *testing.T, runTests func(t *testing.T, reporter ginkgo.Reporter), expected SuiteResults) {
	// Run the Ginkgo suite with output collected by a custom
	// reporter in adddition to the default one. To see what the full
	// Ginkgo report looks like, run this test with "go test -v".
	config.DefaultReporterConfig.FullTrace = true
	gomega.RegisterFailHandler(framework.Fail)
	fakeT := &testing.T{}
	reporter := reporters.NewFakeReporter()
	runTests(fakeT, reporter)

	// Now check the output.
	actual := normalizeReport(*reporter)

	// assert.Equal prints a useful diff if the slices are not
	// equal. However, the diff does not show changes inside the
	// strings. Therefore we also compare the individual fields.
	if !assert.Equal(t, expected, actual) {
		for i := 0; i < len(expected) && i < len(actual); i++ {
			assert.Equal(t, expected[i].Output, actual[i].Output, "output from test #%d (%s)", i, expected[i].Name)
			assert.Equal(t, expected[i].Stack, actual[i].Stack, "stack from test #%d (%s)", i, expected[i].Name)
			assert.Equal(t, expected[i].Failure, actual[i].Failure, "failure from test #%d (%s)", i, expected[i].Name)
		}
	}
}

// TestResult is the outcome of one It spec.
type TestResult struct {
	// Name is the full string for a Ginkgo It, including the "[Top Level]" prefix.
	Name string
	// Output written to GinkgoWriter during test.
	Output string
	// Failure is SpecSummary.Failure.Message with varying parts stripped.
	Failure string
	// Stack is a normalized version (just file names, function parametes stripped) of
	// Ginkgo's FullStackTrace of a failure. Empty if no failure.
	Stack string
}

type SuiteResults []TestResult

func normalizeReport(report reporters.FakeReporter) SuiteResults {
	var results SuiteResults
	for _, spec := range report.SpecSummaries {
		results = append(results, TestResult{
			Name:    strings.Join(spec.ComponentTexts, " "),
			Output:  normalizeLocation(stripAddresses(stripTimes(spec.CapturedOutput))),
			Failure: stripAddresses(stripTimes(spec.Failure.Message)),
			Stack:   normalizeLocation(spec.Failure.Location.FullStackTrace),
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return strings.Compare(results[i].Name, results[j].Name) < 0
	})
	return results
}

// timePrefix matches "Jul 17 08:08:25.950: " at the beginning of each line.
var timePrefix = regexp.MustCompile(`(?m)^[[:alpha:]]{3} +[[:digit:]]{1,2} +[[:digit:]]{2}:[[:digit:]]{2}:[[:digit:]]{2}.[[:digit:]]{3}: `)

func stripTimes(in string) string {
	out := timePrefix.ReplaceAllString(in, "")
	return out
}

// instanceAddr matches " | 0xc0003dec60>"
var instanceAddr = regexp.MustCompile(` \| 0x[0-9a-fA-F]+>`)

func stripAddresses(in string) string {
	return instanceAddr.ReplaceAllString(in, ">")
}

// stackLocation matches "<some path>/<file>.go:75 +0x1f1" after a slash (built
// locally) or one of a few relative paths (built in the Kubernetes CI).
var stackLocation = regexp.MustCompile(`(?:/|vendor/|test/|GOROOT/).*/([[:^space:]]+.go:[[:digit:]]+)( \+0x[0-9a-fA-F]+)?`)

// functionArgs matches "<function name>(...)".
var functionArgs = regexp.MustCompile(`([[:alpha:]]+)\(.*\)`)

// testFailureOutput matches TestFailureOutput() and its source followed by additional stack entries:
//
// k8s.io/kubernetes/test/e2e/framework/pod/pod_test.TestFailureOutput(0xc000558800)
//	/nvme/gopath/src/k8s.io/kubernetes/test/e2e/framework/pod/wait_test.go:73 +0x1c9
// testing.tRunner(0xc000558800, 0x1af2848)
// 	/nvme/gopath/go/src/testing/testing.go:865 +0xc0
// created by testing.(*T).Run
//	/nvme/gopath/go/src/testing/testing.go:916 +0x35a
var testFailureOutput = regexp.MustCompile(`(?m)^k8s.io/kubernetes/test/e2e/framework/internal/output\.TestGinkgoOutput\(.*\n\t.*(\n.*\n\t.*)*`)

// normalizeLocation removes path prefix and function parameters and certain stack entries
// that we don't care about.
func normalizeLocation(in string) string {
	out := in
	out = stackLocation.ReplaceAllString(out, "$1")
	out = functionArgs.ReplaceAllString(out, "$1()")
	out = testFailureOutput.ReplaceAllString(out, "")
	return out
}
