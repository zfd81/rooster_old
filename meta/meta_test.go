package meta

import (
	"fmt"
	"testing"
)

func TestWriteFile(t *testing.T) {
	ins := &Instance{Name: "zfd", Text: "ccc"}
	WriteFile("abc", ins)
}

func TestReadFile(t *testing.T) {
	ins, _ := ReadFile("abc")
	fmt.Println(*ins)
}
