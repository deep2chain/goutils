package goutils

import (
	"testing"
)

func TestFile2List(t *testing.T) {
	var filepath string = "test/abc.lst"
	expected := []string{"A", "B", "C"}
	var err error
	var results []string
	if results, err = File2List(filepath); err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if !Equals(ToInterfaceSlice(results), ToInterfaceSlice(expected)) {
		t.Errorf("The result is not identical with expected result.")
	}
}

func TestList2File(t *testing.T) {
	var filepath string = "test/bcd.lst"
	input := []string{"B", "C", "D"}
	List2File(filepath, ToInterfaceSlice(input))
	var err error
	var output []string
	if output, err = File2List(filepath); err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if !Equals(ToInterfaceSlice(input), ToInterfaceSlice(output)) {
		t.Errorf("The input is expected to be identical with output.")
	}
}
