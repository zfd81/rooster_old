package xsql

import (
	"reflect"
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
	param, err := NewStructParams(user)
	str, params, err := bindParams(str, param)
	t.Log(str)
	t.Log(params)
	t.Log(err)
}

func Test_insertByMap(t *testing.T) {
	countryCapitalMap := make(map[string]interface{})
	countryCapitalMap["NAME"] = "Paris"
	countryCapitalMap["pwd"] = "Rome"
	str, params, err := insertByMap("aaa", countryCapitalMap)
	t.Log(str)
	t.Log(params)
	t.Log(err)
}

func Test_insertByStruct(t *testing.T) {
	type User struct {
		Name string
		Pwd  string
	}
	pa1 := &User{"zfd", "4568"}
	var p User
	//p.Name = ""
	//p.Pwd = ""
	//str, params, err := insertByStruct("aaa", p)
	//t.Log(str)
	//t.Log(len(params))
	//t.Log(err)
	countryCapitalMap := make(map[string]interface{})
	t.Log(reflect.ValueOf(p).Kind())
	t.Log(reflect.ValueOf(pa1).Kind())
	t.Log(reflect.ValueOf(countryCapitalMap).Kind())
}
