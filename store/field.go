package store

import "encoding/binary"

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

func (f *Field) ToInt() int64 {
	return int64(binary.BigEndian.Uint64(*f))
}

func NewField(bytes []byte) Field {
	return bytes
}
