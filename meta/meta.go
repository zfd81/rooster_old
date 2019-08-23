package meta

import (
	"bytes"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

func WriteFile(filename string, ins *Instance) error {
	data, err := bson.Marshal(ins)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func ReadFile(filename string) (*Instance, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ins := NewInstance()
	err = bson.Unmarshal(data, ins)
	bytes.HasPrefix()
	if err != nil {
		return nil, err
	}
	return ins, err
}
