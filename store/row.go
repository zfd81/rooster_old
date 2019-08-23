package store

import (
	"bufio"
	"bytes"
)

const (
	defaultRowCapacity = 256
)

type Row []Field

func (r *Row) Set(bytes []byte, args ...byte) {
	if bytes != nil {
		*r = bytes2Fields(bytes, args...)
	}
}

func (r *Row) Add(field Field) *Row {
	if field != nil {
		*r = append(*r, field)
	}
	return r
}

func (r *Row) GetField(index int) *Field {
	return &(*r)[index]
}

func (r *Row) SetField(index int, field Field) {
	(*r)[index] = field
}

func (r *Row) Length() int {
	return len(*r)
}

func (r *Row) Size() int {
	size := 0
	for _, field := range *r {
		size = size + field.Size()
	}
	return size
}

func NewRow() *Row {
	return NewRowByCapacity(defaultRowCapacity)
}

func NewRowByCapacity(cap int) *Row {
	var row Row = make([]Field, 0, cap)
	return &row
}

func NewRowByFields(fields []Field) *Row {
	if fields == nil {
		return NewRow()
	}
	var row Row = fields
	return &row
}

func NewRowByBytes(bytes []byte, args ...byte) *Row {
	if bytes == nil {
		return NewRow()
	}
	return NewRowByFields(bytes2Fields(bytes, args...))
}

func bytes2Fields(data []byte, args ...byte) []Field {
	if data != nil {
		if len(args) == 0 {
			return []Field{NewField(data)}
		} else {
			scanner := bufio.NewScanner(bytes.NewReader(data))
			//scanner.Split(bufio.ScanWords)
			scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}
				if i := bytes.IndexByte(data, args[0]); i >= 0 {
					return i + 1, data[0:i], nil
				}
				if atEOF {
					return len(data), data, nil
				}
				return 0, nil, nil
			})
			fields := make([]Field, 0, defaultRowCapacity)
			for scanner.Scan() {
				fields = append(fields, NewField(scanner.Bytes()))
			}
			return fields
		}
	}
	return nil
}
