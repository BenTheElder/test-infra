package main

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"os"

	"k8s.io/apimachinery/pkg/util/sets"

	junit "github.com/joshdk/go-junit"
)

// GetPassed collects the set of passed tests from all suites by the test name
func GetPassed(suites []junit.Suite) sets.String {
	passed := sets.NewString()
	for _, suite := range suites {
		for _, test := range suite.Tests {
			if test.Status == junit.StatusPassed {
				passed.Insert(test.Name)
			}
		}
	}
	return passed
}

// GetExpected returns the list of expected tests given the raw contexts
// of a conformance.txt file
// See EG: https://github.com/kubernetes/kubernetes/blob/6f897af2daffd8cd3539f1ef279b9f9bc280d98f/test/conformance/testdata/conformance.txt
func GetExpected(raw []byte) sets.String {
	buff := bytes.NewBuffer(raw)
	scanner := bufio.NewScanner(buff)
	expected := sets.NewString()
	for scanner.Scan() {
		expected.Insert(scanner.Text())
	}
	return expected
}

// GetExpectedFromFile wraps GetExpected, first reading in the file at path
func GetExpectedFromFile(path string) (sets.String, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return GetExpected(raw), nil
}

func main() {
	junitPath := flag.String("junit-path", "", "path to the junit results file")
	txtPath := flag.String("txt-path", "", "path to conformance.txt, the list of expected tests")
	flag.Parse()

	expected, err := GetExpectedFromFile(*txtPath)
	if err != nil {
		panic(err)
	}

	suites, err := junit.IngestFile(*junitPath)
	if err != nil {
		panic(err)
	}
	passed := GetPassed(suites)

	diff := expected.Difference(passed)
	if diff.Len() != 0 {
		println("Expected set of tests does not match passed tests!")
		println("The following tests were expected but not found: ")
		l := diff.List()
		for _, t := range l {
			println(l)
		os.Exit(-1)
	}

	println("Success")
}
