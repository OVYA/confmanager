package util

import (
	"os"
)

// test if the direcroy exist
func IsExistDir(dirname string) bool {

	f, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// test if the file exist
func IsExistFile(filename string) bool {

	f, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !f.IsDir()
}
