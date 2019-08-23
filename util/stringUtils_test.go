package util

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	type args struct {
		str       string
		substr    string
		fromIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add socket cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.str, tt.args.substr, tt.args.fromIndex); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceBetween(t *testing.T) {

	//str := "select name, age, sex from table where id=#id# and age>#aaa#"
	//var newStr string
	//for i := 0; i < 1000000; i++ {
	//	newStr, _ = ReplaceBetween(str, "#", "#", func(index int, start int, end int, content string) (string, error) {
	//		return "?", nil
	//	})
	//}
	//t.Log(newStr)
	var item string
	//arr := [7]string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}
	//arr := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}
	//for i := 0; i < 100000000; i++ {
	//	item = arr[5]
	//}
	var ma = make(map[string]string)
	ma["aa"] = "aa"
	ma["bb"] = "bb"
	ma["cc"] = "cc"
	ma["dd"] = "dd"
	ma["ee"] = "ee"
	ma["ff"] = "ff"
	ma["gg"] = "gg"
	for i := 0; i < 100000000; i++ {
		item = ma["dd"]
	}
	t.Log(item)
}
