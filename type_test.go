package goutils

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	a := 123
	b := "123"
	c := []int{1, 2, 3}
	d := func() string { return "123" }
	if Type(a) != reflect.Int {
		t.Errorf("expected reflect.Int")
	}
	if !IsInstance(b, reflect.String) {
		t.Errorf("expected true")
	}
	if !IsSlice(c) {
		t.Errorf("expected true")
	}
	if Type(d) != reflect.Func {
		t.Errorf("expected reflect.Func")
	}
	if !IsString(d()) {
		t.Errorf("expected true")
	}
}
