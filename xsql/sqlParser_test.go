package xsql

import (
	"testing"
)

func Test_bindMap(t *testing.T) {
	str := "select name from tbale where name=:name {and pwd =:pwd} and{} {age>1}"
	countryCapitalMap := make(map[string]interface{})
	countryCapitalMap["name"] = "Paris"
	countryCapitalMap["pwd"] = "Rome"
	str, params, err := bindMap(str, countryCapitalMap)
	t.Log(str)
	t.Log(params)
	t.Log(err)
}
