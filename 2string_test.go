package goutils

import (
	"strconv"
	"testing"
)

var (
	a = 123
	b = []int{1, 2, 3}
	c = 1.23
	// d = 0x111
)

func Test2StringEx(t *testing.T) {

	if !IsString(ToStringEx(a)) {
		t.Errorf("%T:%v)stringified failure\n", a, a)
	}
	if !IsString(ToStringEx(b)) {
		t.Errorf("%T:%v)stringified failure\n", a, a)
	}
	if !IsString(ToStringEx(c)) {
		t.Errorf("%T:%v)stringified failure\n", a, a)
	}
}

func Test2StringBy(t *testing.T) {
	var x int64 = 100
	y := []byte{1, 2, 3}
	if !IsString(strconv.FormatInt(x, 10)) {
		t.Errorf("%T:%v)stringified failure\n", x, x)
	}
	if !IsString(string(y)) {
		t.Errorf("%T:%v)stringified failure\n", y, y)
	}
}
