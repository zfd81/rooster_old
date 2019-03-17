package xsql

import (
	"testing"
)

func Test_variable(t *testing.T) {
	str := "select name from tbale where name=:name and pwd =:pwd and age>1"
	for i := 0; i < 10000; i++ {
		// t.Log(variable(str))
		variable(str)
	}
}
