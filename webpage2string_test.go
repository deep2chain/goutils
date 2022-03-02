package goutils

import (
	"testing"
)

func TestGetContent(t *testing.T) {
	content, err := getWebSite("https://google.com")
	if content == nil || err != nil {
		t.Errorf("content is nil or err is not nil %v\n", err)
	}
}
