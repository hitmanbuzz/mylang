package utils

import (
	"errors"
	"os"
)

func IsFileExist(fname string) bool {
	_, err := os.Stat(fname)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// byte -> []byte
func ToByteArr(bs ...byte) []byte {
	return bs
}

func IsNum(b byte) bool {
	return b >= '0' && b <= '9'
}
