package util

import "bytes"

func Equal(a, b []byte) bool {
	return bytes.Equal(a, b)
}

func NotEqual(a, b []byte) bool {
	return !bytes.Equal(a, b)
}

func Greater(a, b []byte) bool {
	return bytes.Compare(a, b) == 1
}

func GreaterOrEqual(a, b []byte) bool {
	return bytes.Compare(a, b) >= 0
}

func Less(a, b []byte) bool {
	return bytes.Compare(a, b) == -1
}

func LessOrEqual(a, b []byte) bool {
	return bytes.Compare(a, b) <= 0
}

func In(a []byte, b [][]byte) bool {
	for _, v := range b {
		if bytes.Equal(a, v) {
			return true
		}
	}
	return false
}

func NotIn(a []byte, b [][]byte) bool {
	for _, v := range b {
		if bytes.Equal(a, v) {
			return false
		}
	}
	return true
}

func Empty(val []byte) bool {
	if val == nil || len(val) == 0 {
		return true
	}
	return false
}

func NotEmpty(val []byte) bool {
	if val != nil && len(val) != 0 {
		return true
	}
	return false
}
