package meta

import (
	"encoding/json"
	"io/ioutil"
)

type SchemaState byte

const (
	// StateNone means this schema element is absent and can't be used.
	StateNone SchemaState = iota
	// StateDeleteOnly means we can only delete items for this schema element.
	StateDeleteOnly
	// StateWriteOnly means we can use any write operation on this schema element,
	// but outer can't read the changed data.
	StateWriteOnly
	// StateWriteReorganization means we are re-organizing whole data after write only state.
	StateWriteReorganization
	// StateDeleteReorganization means we are re-organizing whole data after delete only state.
	StateDeleteReorganization
	// StatePublic means this schema element is ok for all write and read operations.
	StatePublic
)

var (
	instances = make([]Instance, 0, 10)
)

func (s SchemaState) String() string {
	switch s {
	case StateDeleteOnly:
		return "delete only"
	case StateWriteOnly:
		return "write only"
	case StateWriteReorganization:
		return "write reorganization"
	case StateDeleteReorganization:
		return "delete reorganization"
	case StatePublic:
		return "public"
	default:
		return "none"
	}
}

func writeFile(filename string, in interface{}) error {
	//data, err := bson.Marshal(in)
	data, err := json.Marshal(in)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func readFile(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//return bson.Unmarshal(data, out)
	return json.Unmarshal(data, out)
}

func ReadMeta(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			ReadInstance(file.Name())
		}
	}
	return nil
}

func CreateInstance(name string) *Instance {
	ins := NewInstance()
	ins.Name = name
	return ins
}
