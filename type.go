package goutils

import "reflect"

func Type(value interface{}) reflect.Kind {
	// return fmt.Sprintf("%T", x)
	// reflect.TypeOf(value)
	return reflect.ValueOf(value).Kind()
}

func IsInstance(value interface{}, kind reflect.Kind) bool {
	return Type(value) == kind
}

func IsString(value interface{}) bool {
	return IsInstance(value, reflect.String)
}

func IsSlice(value interface{}) bool {
	return Type(value) == reflect.Slice || Type(value) == reflect.Array
}
