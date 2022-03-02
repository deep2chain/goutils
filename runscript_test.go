package goutils

import (
	"fmt"
	"strings"
	"testing"
)

func TestRunScripts(t *testing.T) {
	hello := "hello,world"
	result, err := RunScripts("echo " + hello)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Trim(result, "\n") != hello {
		t.Errorf("result expected %v", hello)
	}
}
