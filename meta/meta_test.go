package meta

import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	//ins := &Instance{Name: "zfd", Text: "ccc"}
	//writeFile("abc", ins)
}

func TestReadFile(t *testing.T) {
	//ins, _ := readFile("abc")
	//fmt.Println(*ins)
}

func Test_getMetaName(t *testing.T) {
	path := "/rooster/meta/ins_zfd"
	t.Log(getMetaName(path))
}

func Test_getName(t *testing.T) {
	mname := "ins_zfd"
	t.Log(mname)
	t.Log(getName(mname))
	mname = "ins_zfd_cc"
	t.Log(mname)
	t.Log(getName(mname))
}
