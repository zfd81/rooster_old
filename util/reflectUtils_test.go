package util

import "testing"

func TestStructIterator(t *testing.T) {
	type User struct {
		Name string
		Age  int
		Sex  string
	}

	user := &User{"zfd", 25, "boy"}

	StructIterator(user, func(index int, key string, value interface{}) {
		t.Log(index)
		t.Log(key)
		t.Log(value)
	})
}
