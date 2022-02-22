package goutils

import "os"

func SetEnv(key string, value string) {

}

func GetEnv(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return ""
}

func isEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}
