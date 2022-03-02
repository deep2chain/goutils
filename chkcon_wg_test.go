package goutils

import (
	"fmt"
	"testing"
)

func TestCheckConnectivity(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://baidu.com",
		"https://github.com",
		"https://msn.com",
	}
	connected := CheckConnectivity(websites)
	fmt.Printf("connected cnt %d\n", connected)
}
