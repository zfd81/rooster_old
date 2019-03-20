package xsql

import (
	"fmt"

	"github.com/zfd81/rooster/util"
)

// func bindMap(sql string, arg map[string]interface{}) (string, []interface{}, error) {
// 	arglist := make([]interface{}, 0, 20)
// 	newSql, err := util.ReplaceByKeyword(sql, ':', func(index int, start int, end int, content string) (string, error) {
// 		val, ok := arg[content]
// 		if !ok {
// 			return "?", fmt.Errorf("could not find name %s in %#v", content, arg)
// 		}
// 		arglist = append(arglist, val)
// 		return "?", nil
// 	})
// 	return newSql, arglist, err
// }

func bindMap(sql string, arg Paramer) (string, []interface{}, error) {
	newSql, err := util.ReplaceBetween(sql, "{", "}", func(index int, start int, end int, content string) (string, error) {
		ignore := false
		fragment, err := util.ReplaceByKeyword(content, ':', func(i int, s int, e int, c string) (string, error) {
			_, ok := arg.Get(c)
			if !ok {
				ignore = true
				return "", nil
			}
			return fmt.Sprintf(":%s", c), nil
		})
		if ignore {
			return "", err
		}
		return fragment, err
	})
	if err != nil {
		return "", nil, err
	}
	params := make([]interface{}, 0, 20)
	newSql, err = util.ReplaceByKeyword(newSql, ':', func(index int, start int, end int, content string) (string, error) {
		val, ok := arg.Get(content)
		if !ok {
			return "?", fmt.Errorf("could not find name %s in %#v", content, arg)
		}
		params = append(params, val)
		return "?", nil
	})
	return newSql, params, err
}
