package goutils

import "testing"

func TestIsEnvExists(t *testing.T) {
	if !isEnvExist("GOPATH") {
		t.Errorf("Use `echo $GOPATH` check if $GOPATH exists. if exists, this function is in trouble.")
	}
}
