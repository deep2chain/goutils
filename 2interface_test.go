package goutils

import (
	"testing"
)

func Test2InterfaceSlice(t *testing.T) {
	var array = []string{"a", "b", "c"}
	var iarray = ToInterfaceSlice(array)
	if !IsSlice(iarray) {
		t.Errorf("%T should be reflect.Arrary, but %v", iarray, Type(iarray))
	}
}

func Test2Interface(t *testing.T) {
	b := "123"
	if !IsString(ToInterface(b).(string)) {
		t.Errorf("%T:%v)stringified failure\n", b, b)
	}
}
