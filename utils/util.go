package utils

import (
	"errors"
	"os"
)

// byte representation of number from 0-9 as string
var NUM_ARR = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

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

// check if `b` is in `arrByte`
func IsByteContain(arrByte []byte, b byte) bool {
	low := 0
	high := len(arrByte) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arrByte[mid] == b {
			return true
		} else if arrByte[mid] < b {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
