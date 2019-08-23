package handler

import (
	"encoding/binary"
	"testing"
)

var b1 = []byte("abc")
var b2 = []byte("abc")
var b3 = []byte("abd")
var b4 = Int64ToBytes(10)
var b5 = Int64ToBytes(10)
var b6 = Int64ToBytes(20)
var b7 = Int64ToBytes(10)

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func TestEqual(t *testing.T) {
	t.Log("==========查看地址")
	t.Log(Equal(b1, b2))
	t.Log(NotEqual(b1, b2))
	t.Log(Equal(b1, b3))
	t.Log("==========比较字符串")
	t.Log(Equal(b1, b2)())
	t.Log(Equal(b1, b3)())
	t.Log("==========比较整数")
	t.Log(Equal(b4, b5)())
	t.Log(Equal(b4, b6)())
	t.Log(Equal(b5, b6)())
}

func TestNotEqual(t *testing.T) {
	t.Log("==========比较字符串")
	t.Log(NotEqual(b1, b2)())
	t.Log(NotEqual(b1, b3)())
	t.Log("==========比较整数")
	t.Log(NotEqual(b4, b5)())
	t.Log(NotEqual(b4, b6)())
	t.Log(NotEqual(b5, b6)())
}

func TestGreater(t *testing.T) {
	t.Log("==========比较字符串")
	t.Log(Greater(b1, b2)())
	t.Log(Greater(b1, b3)())
	t.Log(Greater(b3, b1)())
	t.Log("==========比较整数")
	t.Log(Greater(b4, b5)())
	t.Log(Greater(b4, b6)())
	t.Log(Greater(b5, b6)())
	t.Log(Greater(b6, b5)())
}

func TestGreaterOrEqual(t *testing.T) {
	t.Log("==========比较字符串")
	t.Log(GreaterOrEqual(b1, b2)())
	t.Log(GreaterOrEqual(b1, b3)())
	t.Log(GreaterOrEqual(b3, b1)())
	t.Log("==========比较整数")
	t.Log(GreaterOrEqual(b4, b5)())
	t.Log(GreaterOrEqual(b4, b6)())
	t.Log(GreaterOrEqual(b5, b6)())
	t.Log(GreaterOrEqual(b6, b5)())
}

func TestLessOrEqual(t *testing.T) {
	t.Log("==========比较字符串")
	t.Log(LessOrEqual(b1, b2)())
	t.Log(LessOrEqual(b1, b3)())
	t.Log(LessOrEqual(b3, b1)())
	t.Log("==========比较整数")
	t.Log(LessOrEqual(b4, b5)())
	t.Log(LessOrEqual(b4, b6)())
	t.Log(LessOrEqual(b5, b6)())
	t.Log(LessOrEqual(b6, b5)())
}

func TestHasPrefix(t *testing.T) {
	t.Log(HasPrefix(b1, []byte("a"))())
	t.Log(HasPrefix(b1, []byte("ab"))())
	t.Log(HasPrefix(b1, []byte("c"))())
}

func TestHasSuffix(t *testing.T) {
	t.Log(HasSuffix(b1, []byte("a"))())
	t.Log(HasSuffix(b1, []byte("ab"))())
	t.Log(HasSuffix(b1, []byte("bc"))())
	t.Log(HasSuffix(b1, []byte("c"))())
}
