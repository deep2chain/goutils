package goutils

import "fmt"

// check if []string - A,B,C,... Identical
func Equals(params ...[]interface{}) bool {

	var size = len(params[0])
	if size == 1 {
		fmt.Println("No Comparable. At least 2 params needed.")
		return true
	}
	// check if the sizes are same
	for _, param := range params {
		if len(param) != size {
			return false
		}
	}
	// check if item with same index is all same.
	for i, v := range params[0] {
		for _, param := range params {
			if param[i] != v {
				return false
			}
		}
	}

	return true
}
