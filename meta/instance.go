package meta

import "os"

const (
	instanceFileName = "instance.m"
)

type InstanceInfo struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	Comment string `json:"comment,omitempty"`
}

type Instance struct {
	InstanceInfo
	Databases []*Database `json:"-"`
}

//func CreateDatabase() *Database {
//	return
//}

func NewInstance() *Instance {
	return &Instance{}
}

func ReadInstance(path string) *Instance {
	filePath := path + "/" + instanceFileName
	_, err := os.Stat(filePath)
	if err != nil {
		return nil
	}
	ins := NewInstance()
	err = readFile(filePath, ins)
	if err != nil {
		return nil
	}
	return ins
}

func WriteInstance(path string, ins *Instance) error {
	filePath := path + "/" + instanceFileName
	return writeFile(filePath, ins)
}
