/*
Copyright 2015 The Kubernetes Authors.

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

package e2e

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/ginkgo/v2/types"
	"gopkg.in/yaml.v2"

	// Never, ever remove the line with "/ginkgo". Without it,
	// the ginkgo test runner will not detect that this
	// directory contains a Ginkgo test suite.
	// See https://github.com/kubernetes/kubernetes/issues/74827
	// "github.com/onsi/ginkgo/v2"

	"k8s.io/component-base/version"
	"k8s.io/klog/v2"
	conformancetestdata "k8s.io/kubernetes/test/conformance/testdata"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"
	"k8s.io/kubernetes/test/e2e/framework/testfiles"
	e2etestingmanifests "k8s.io/kubernetes/test/e2e/testing-manifests"
	testfixtures "k8s.io/kubernetes/test/fixtures"
	"k8s.io/kubernetes/test/utils/image"

	// test sources
	_ "k8s.io/kubernetes/test/e2e/apimachinery"
	_ "k8s.io/kubernetes/test/e2e/apps"
	_ "k8s.io/kubernetes/test/e2e/architecture"
	_ "k8s.io/kubernetes/test/e2e/auth"
	_ "k8s.io/kubernetes/test/e2e/autoscaling"
	_ "k8s.io/kubernetes/test/e2e/cloud"
	_ "k8s.io/kubernetes/test/e2e/common"
	_ "k8s.io/kubernetes/test/e2e/instrumentation"
	_ "k8s.io/kubernetes/test/e2e/kubectl"
	_ "k8s.io/kubernetes/test/e2e/lifecycle"
	_ "k8s.io/kubernetes/test/e2e/lifecycle/bootstrap"
	_ "k8s.io/kubernetes/test/e2e/network"
	_ "k8s.io/kubernetes/test/e2e/node"
	_ "k8s.io/kubernetes/test/e2e/scheduling"
	_ "k8s.io/kubernetes/test/e2e/storage"
	_ "k8s.io/kubernetes/test/e2e/storage/external"
	_ "k8s.io/kubernetes/test/e2e/windows"
)

// handleFlags sets up all flags and parses the command line.
func handleFlags() {
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
}

func TestMain(m *testing.M) {
	var versionFlag bool
	flag.CommandLine.BoolVar(&versionFlag, "version", false, "Displays version information.")

	// Register test flags, then parse flags.
	handleFlags()

	if framework.TestContext.ListImages {
		for _, v := range image.GetImageConfigs() {
			fmt.Println(v.GetE2EImage())
		}
		os.Exit(0)
	}
	if versionFlag {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}

	// Enable embedded FS file lookup as fallback
	testfiles.AddFileSource(e2etestingmanifests.GetE2ETestingManifestsFS())
	testfiles.AddFileSource(testfixtures.GetTestFixturesFS())
	testfiles.AddFileSource(conformancetestdata.GetConformanceTestdataFS())

	if framework.TestContext.ListConformanceTests {
		var tests []struct {
			Testname    string `yaml:"testname"`
			Codename    string `yaml:"codename"`
			Description string `yaml:"description"`
			Release     string `yaml:"release"`
			File        string `yaml:"file"`
		}

		data, err := testfiles.Read("test/conformance/testdata/conformance.yaml")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := yaml.Unmarshal(data, &tests); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := yaml.NewEncoder(os.Stdout).Encode(tests); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	framework.AfterReadingAllFlags(&framework.TestContext)

	// TODO: Deprecating repo-root over time... instead just use gobindata_util.go , see #23987.
	// Right now it is still needed, for example by
	// test/e2e/framework/ingress/ingress_utils.go
	// for providing the optional secret.yaml file and by
	// test/e2e/framework/util.go for cluster/log-dump.
	if framework.TestContext.RepoRoot != "" {
		testfiles.AddFileSource(testfiles.RootFileSource{Root: framework.TestContext.RepoRoot})
	}

	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}

var _ = ginkgo.ReportAfterEach(func(report ginkgo.SpecReport) {
	progressReporter.ProcessSpecReport(report)
})

// writeSpecSummaryOutput will output details about every test (name, files, lines, etc) which helps
// when documenting our tests.
func writeSpecSummaryOutput(report ginkgo.Report) {
	if len(framework.TestContext.SpecSummaryOutput) <= 0 {
		return
	}
	absPath, err := filepath.Abs(framework.TestContext.SpecSummaryOutput)
	if err != nil {
		klog.Errorf("%#v\n", err)
		panic(err)
	}
	f, err := os.Create(absPath)
	if err != nil {
		klog.Errorf("%#v\n", err)
		panic(err)
	}

	defer f.Close()

	for _, specReport := range report.SpecReports {
		b, err := specReport.MarshalJSON()
		if err != nil {
			klog.Errorf("Error in detail reporter: %v", err)
			return
		}
		if _, err = f.Write(b); err != nil {
			klog.Errorf("Error saving test details in detail reporter: %v", err)
			return
		}
		// Printing newline between records for easier viewing in various tools.
		if _, err = fmt.Fprintln(f, ""); err != nil {
			klog.Errorf("Error saving test details in detail reporter: %v", err)
			return
		}
	}
}

// writeJUnitReport generates a JUnit file in the e2e report directory that is
// shorter than the one normally written by `ginkgo --junit-report`. This is
// needed because the full report can become too large for tools like Spyglass
// (https://github.com/kubernetes/kubernetes/issues/111510).
//
// Users who want the full report can use `--junit-report`.
func writeJUnitReport(report ginkgo.Report) {
	if framework.TestContext.ReportDir == "" {
		return
	}

	trimmedReport := report
	trimmedReport.SpecReports = nil
	for _, specReport := range report.SpecReports {
		// Remove details for any spec that hasn't failed. In Prow,
		// the test output captured in build-log.txt has all of this
		// information, so we don't need it in the XML.
		if specReport.State != types.SpecStateFailed {
			// strip the "system-error" if the testcase is not failed.
			specReport.CapturedGinkgoWriterOutput = ""
			// strip the "system-out" if the testcase is not failed.
			specReport.CapturedStdOutErr = ""
			// strip some details for tracing each steps executed by Ginkgo, this is used to build the "system-out"
			// while "system-out" is not shown by Spyglass if the testcase is not failed.
			specReport.ReportEntries = nil

		}

		trimmedReport.SpecReports = append(trimmedReport.SpecReports, specReport)
	}

	// With Ginkgo v1, we used to write one file per parallel node. Now
	// Ginkgo v2 automatically merges all results into a report for us. The
	// 01 suffix is kept in case that users expect files to be called
	// "junit_<prefix><number>.xml".
	junitReport := path.Join(framework.TestContext.ReportDir, "junit_"+framework.TestContext.ReportPrefix+"01.xml")
	reporters.GenerateJUnitReport(trimmedReport, junitReport)
}

var _ = ginkgo.ReportAfterSuite("Kubernetes e2e suite report", func(report ginkgo.Report) {
	writeSpecSummaryOutput(report)
	writeJUnitReport(report)
})
