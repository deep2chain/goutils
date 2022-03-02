package goutils

import (
	"errors"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Open(path)
	return !errors.Is(err, os.ErrNotExist)
}

func FileOpen(path string) (*os.File, error) {
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("file not exists error")
	}
	return file, nil
}
