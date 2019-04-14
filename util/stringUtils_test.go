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
		// TODO: Add test cases.
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

	str := "select name, age, sex from table where id=#id# and age>#aaa#"
	var newStr string
	for i := 0; i < 1000000; i++ {
		newStr, _ = ReplaceBetween(str, "#", "#", func(index int, start int, end int, content string) (string, error) {
			return "?", nil
		})
	}
	t.Log(newStr)
}
