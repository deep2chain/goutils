package goutils

import (
	"fmt"
	"os/exec"
)

func RunScripts(script string) (output string, err error) {
	var cmd []byte
	if FileExists(script) {
		cmd, err = exec.Command("/bin/sh", script).Output()
	} else {
		cmd, err = exec.Command("/bin/sh", "-c", script).Output()
	}
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output = string(cmd)
	return
}
