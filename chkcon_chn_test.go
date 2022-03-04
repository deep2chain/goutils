package goutils

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://baidu.com",
		"https://github.com",
		"https://msn.com",
	}
	connected := Run(websites)
	fmt.Printf("connected cnt %d\n", connected)
}
