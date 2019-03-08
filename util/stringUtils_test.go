package util

import "testing"

func TestIndexOf(t *testing.T) {
	str := "123456789"
	t.Log(IndexOf(str, "1", 5))
	t.Log(IndexOf(str, "3", 0))
}
