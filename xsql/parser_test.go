package xsql

import (
	"testing"
)

func Test_bindParam(t *testing.T) {
	type User struct {
		Name string
		Pwd  string
	}
	str := "select name from tbale where name=:Name {and pwd =:PWD} and{} {age>1}"
	countryCapitalMap := make(map[string]interface{})
	countryCapitalMap["NAME"] = "Paris"
	countryCapitalMap["pwd"] = "Rome"
	param := NewParam()
	param.Add("NAME", "hello")
	param.Add("pwd", 123)
	// param1, err := NewMapParam(countryCapitalMap)
	// if err != nil {
	// 	t.Log(err.Error())
	// }
	// str, params, err := bindMap(str, param1)
	// t.Log(str)
	// t.Log(params)
	// t.Log(err)
	// var pa User
	pa1 := &User{"zfd", "456"}

	param1, err1 := NewStructParam(pa1)
	if err1 != nil {
		t.Log(err1.Error())
		return
	}
	str, params, err := bindParam(str, param1)
	t.Log(str)
	t.Log(params)
	t.Log(err)
	// if err != nil {
	// 	t.Log(err.Error())
	// }
}
