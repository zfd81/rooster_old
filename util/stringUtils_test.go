package util

import (
	"reflect"
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
	type args struct {
		str      string
		open     string
		close    string
		replacer ReplacerFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReplaceBetween(tt.args.str, tt.args.open, tt.args.close, tt.args.replacer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceBetween() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReplaceBetween() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReplaceBetween() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
