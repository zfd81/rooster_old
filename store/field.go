package store

import (
	"strconv"
)

type Field []byte

func (f *Field) Set(bytes []byte) {
	*f = bytes
}

func (f *Field) Get() []byte {
	return *f
}

func (f *Field) Size() int {
	return len(*f)
}

func (f *Field) ToString() string {
	return string(*f)
}

func (f *Field) ToInt() int {
	//return int64(binary.BigEndian.Uint64(*f))
	val, err := strconv.Atoi(string(*f))
	if err != nil {
		return 0
	}
	return val
}

func NewField(bytes []byte) Field {
	return bytes
}
