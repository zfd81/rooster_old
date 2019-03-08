package util

import "testing"

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
