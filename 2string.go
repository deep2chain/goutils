package goutils

import (
	"fmt"
)

func ToStringArray(values []interface{}) []string {
	var results []string
	for _, v := range values {
		results = append(results, ToString(v))
	}
	return results
}

func ToString(v interface{}) string {
	return v.(string)
}

func ToStringEx(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// func ToStringBy(v interface{}) string {
// 	if IsInstance(v, reflect.Int) {
// 		return strconv.FormatInt(v, 10)
// 	}
// 	if IsInstance(v, reflect.Array) {
// 		return string(v)
// 	}
// }
