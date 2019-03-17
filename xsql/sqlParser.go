package xsql

import (
	"github.com/zfd81/rooster/util"
)

func variable(str string) (string, []string, error) {
	return util.ReplaceByKeyword(str, ':', func(index int, start int, end int, content string) (string, error) {
		// str := fmt.Sprintf("%d:%d:%d", index, start, end)
		return "?", nil
	})
}
