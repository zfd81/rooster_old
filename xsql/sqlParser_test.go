package xsql

import (
	"testing"
)

func Test_bindMap(t *testing.T) {
	// str := "select name from tbale where name=:name {and pwd =:pwd} and{} {age>1}"
	countryCapitalMap := make(map[string]interface{})
	countryCapitalMap["name"] = "Paris"
	countryCapitalMap["pwd"] = "Rome"
	param := NewParam()
	param.Add("name", "hello")
	param.Add("pwd1", 123)
	// param1, err := NewMapParam(emap)
	// if err != nil {
	// 	t.Log(err.Error())
	// }
	// str, params, err := bindMap(str, param1)
	// t.Log(str)
	// t.Log(params)
	// t.Log(err)
	var pa MapParam

	_, err1 := NewStructParam(&pa)
	if err1 != nil {
		t.Log(err1.Error())
	}
}
