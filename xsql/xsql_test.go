package xsql

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB_XExec(t *testing.T) {
	db, err := Open("mysql", "root:123456@tcp(localhost:3306)/tcm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Log(err.Error())
	}
	sql := "insert into test1 values (:key,:pwd,:ip,:tim)"
	param := NewParam()
	param.Add("Key", "zfd1").Add("ip", "1234567").Add("pwd", "890").Add("Tim", time.Now())
	// num, err := db.XExec(sql, param)

	// param.Add("pwd", "hello")

	sql = "update test1 set tim=:tim1 where pwd=:pwd"
	num, err := db.XExec(sql, param)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(num)
}
