# goutils
A collection of handy go utilities. This is go version of pyutils(https://github.com/deep2essence/pyutils)
## One-Line Functions
Function|Usage
--------|-----
*array2slice*|array[:],<br> array[0:len(array)]
*slice2array*|copy(array[:],slice[0:len(array)])
*list2string*|strings.Join(s []string, sep string)
*string2list*|strings.Split(s string,sep string)
## Diff
Q|A
-|--------
*array vs. slice*| array has len(length) &cap(capacity), but slice has no such restriction. what len(slice) returns? what cap(slice) returns?
