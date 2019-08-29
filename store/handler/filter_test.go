package handler

import (
	"bytes"
	"encoding/binary"
	"github.com/zfd81/rooster/store"
	"testing"
	"unsafe"
)

var str string = "第一列1w单独,第一列,第三列,第四例,54,11,33,44,55"

var row *store.Line = store.NewRowByBytes([]byte(str), ',')

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Int2Byte(data int) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

func TestEqual(t *testing.T) {
	t.Log("==========查看地址")
	t.Log(Equal(NewVar(1), NewVar(2)))
	t.Log(NotEqual(NewVar(1), NewVar(2)))
	t.Log(Equal(NewVar(1), NewVar(3)))
	t.Log("==========比较字符串")
	t.Log(Equal(NewVar(0), NewVar(1))(row))
	fun := Equal(NewVar(1), NewConst([]byte("第一列")))
	t.Log(fun.Filter(row))
	t.Log(Equal(NewVar(0), NewVar(2))(row))
	t.Log(Equal(NewVar(0), NewConst([]byte("第一列")))(row))
	t.Log("==========比较整数")
	t.Log(Equal(NewVar(4), NewVar(5))(row))
	t.Log(Equal(NewVar(4), NewConst(12))(row))
	t.Log(Equal(NewVar(4), NewVar(6))(row))

}

//
//func TestNotEqual(t *testing.T) {
//	t.Log("==========比较字符串")
//	t.Log(NotEqual(b1, b2)())
//	t.Log(NotEqual(b1, b3)())
//	t.Log("==========比较整数")
//	t.Log(NotEqual(b4, b5)())
//	t.Log(NotEqual(b4, b6)())
//	t.Log(NotEqual(b5, b6)())
//}
//
func TestGreater(t *testing.T) {
	t.Log("==========比较字符串")
	//t.Log(Greater(b1, b2)())
	t.Log("==========比较整数")
	t.Log(Greater(NewConst(351), NewConst(35))(row))
	rs := bytes.Runes(*row.GetField(0))
	t.Log(len(rs))
}

//
//func TestGreaterOrEqual(t *testing.T) {
//	t.Log("==========比较字符串")
//	t.Log(GreaterOrEqual(b1, b2)())
//	t.Log(GreaterOrEqual(b1, b3)())
//	t.Log(GreaterOrEqual(b3, b1)())
//	t.Log("==========比较整数")
//	t.Log(GreaterOrEqual(b4, b5)())
//	t.Log(GreaterOrEqual(b4, b6)())
//	t.Log(GreaterOrEqual(b5, b6)())
//	t.Log(GreaterOrEqual(b6, b5)())
//}
//
//func TestLessOrEqual(t *testing.T) {
//	t.Log("==========比较字符串")
//	t.Log(LessOrEqual(b1, b2)())
//	t.Log(LessOrEqual(b1, b3)())
//	t.Log(LessOrEqual(b3, b1)())
//	t.Log("==========比较整数")
//	t.Log(LessOrEqual(b4, b5)())
//	t.Log(LessOrEqual(b4, b6)())
//	t.Log(LessOrEqual(b5, b6)())
//	t.Log(LessOrEqual(b6, b5)())
//}
//
//func TestHasPrefix(t *testing.T) {
//	t.Log(HasPrefix(b1, []byte("a"))())
//	t.Log(HasPrefix(b1, []byte("ab"))())
//	t.Log(HasPrefix(b1, []byte("c"))())
//}
//
//func TestHasSuffix(t *testing.T) {
//	t.Log(HasSuffix(b1, []byte("a"))())
//	t.Log(HasSuffix(b1, []byte("ab"))())
//	t.Log(HasSuffix(b1, []byte("bc"))())
//	t.Log(HasSuffix(b1, []byte("c"))())
//}
//
//func TestAND(t *testing.T) {
//	t.Log(AND(HasSuffix(b1, []byte("c")), HasSuffix(b1, []byte("abc")), HasSuffix(b1, []byte("bc")))())
//}
//
//func TestOR(t *testing.T) {
//	fun := OR(HasSuffix(b1, []byte("a")), HasSuffix(b1, []byte("abc")), HasSuffix(b1, []byte("bc")))
//	t.Log(fun())
//}
