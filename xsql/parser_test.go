package xsql

import (
	"testing"
)

func Test_bindParams(t *testing.T) {
	type User struct {
		Name string
		Pwd  string
	}
	str := "select name from tbale where name=:Name {and pwd =:PWD} and{} {age>1}"

	//mapParam := make(map[string]interface{})
	//mapParam["NAME"] = "Paris"
	//mapParam["pwd"] = "Rome"
	//param, err := NewMapParams(mapParam)
	//str, params, err := bindParams(str, param)
	//t.Log(str)
	//t.Log(params)
	//t.Log(err)

	//param := NewParams()
	//param.Add("NAME", "hello")
	//param.Add("pwd", 123)
	//str, params, err := bindParams(str, param)
	//t.Log(str)
	//t.Log(params)
	//t.Log(err)

	user := &User{"zfd", "456"}
	param := NewStructParams(user)
	str, params, err := bindParams(str, param)
	t.Log(str)
	t.Log(params)
	t.Log(err)
	t.Log(param.Size())
	t.Log(param.Names())
	t.Log(len(param.Names()))
}

func Test_insert(t *testing.T) {
	type User struct {
		Name string
		Pwd  string
	}
	user := &User{"zfd", "4568"}
	str, params, err := insert("userInfo", NewStructParams(user))
	t.Log(str)
	t.Log(params)
	t.Log(len(params))
	t.Log(err)

	countryCapitalMap := make(map[string]interface{})
	countryCapitalMap["NAME"] = "Paris"
	countryCapitalMap["pwd"] = "Rome"
	param := NewMapParams(countryCapitalMap)
	str, params, err = insert("UserInfo2", param)
	t.Log(str)
	t.Log(params)
	t.Log(err)
}
