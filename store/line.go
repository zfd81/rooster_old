package store

import (
	"bufio"
	"bytes"
)

const (
	defaultRowCapacity = 256
)

type Line []Field

func (l *Line) Set(bytes []byte, args ...byte) {
	if bytes != nil {
		*l = bytes2Fields(bytes, args...)
	}
}

func (l *Line) Add(field Field) *Line {
	if field != nil {
		*l = append(*l, field)
	}
	return l
}

func (l *Line) GetField(index int) *Field {
	return &(*l)[index]
}

func (l *Line) SetField(index int, field Field) {
	(*l)[index] = field
}

func (l *Line) Length() int {
	return len(*l)
}

func (l *Line) Size() int {
	size := 0
	for _, field := range *l {
		size = size + field.Size()
	}
	return size
}

func NewRow() *Line {
	return NewRowByCapacity(defaultRowCapacity)
}

func NewRowByCapacity(cap int) *Line {
	var line Line = make([]Field, 0, cap)
	return &line
}

func NewRowByFields(fields []Field) *Line {
	if fields == nil {
		return NewRow()
	}
	var line Line = fields
	return &line
}

func NewRowByBytes(bytes []byte, args ...byte) *Line {
	if bytes == nil {
		return NewRow()
	}
	return NewRowByFields(bytes2Fields(bytes, args...))
}

func bytes2Fields(data []byte, split ...byte) []Field {
	if data != nil {
		if len(split) == 0 {
			return []Field{NewField(data)}
		} else {
			scanner := bufio.NewScanner(bytes.NewReader(data))
			//scanner.Split(bufio.ScanWords)
			scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}
				if i := bytes.IndexByte(data, split[0]); i >= 0 {
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
