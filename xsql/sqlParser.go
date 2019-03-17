package xsql

import (
	"fmt"

	"github.com/zfd81/rooster/util"
)

func bindMap(sql string, arg map[string]interface{}) (string, []interface{}, error) {
	arglist := make([]interface{}, 0, 20)
	newSql, err := util.ReplaceByKeyword(sql, ':', func(index int, start int, end int, content string) (string, error) {
		val, ok := arg[content]
		if !ok {
			return "?", fmt.Errorf("could not find name %s in %#v", content, arg)
		}
		arglist = append(arglist, val)
		return "?", nil
	})
	return newSql, arglist, err
}
