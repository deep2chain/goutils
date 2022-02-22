package goutils

import (
	"fmt"
	"testing"
)

func TestRunScripts(t *testing.T) {
	hello := "hello,world"
	result, err := RunScripts("echo " + hello)
	if err != nil {
		fmt.Println(err)
	}
	if result != hello {
		t.Errorf("expected true")
	}
}
