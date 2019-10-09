package meta

import (
	"encoding/json"
	"fmt"
)

type InstanceInfo struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	Comment string `json:"comment,omitempty"`
}

type Instance struct {
	InstanceInfo
	Databases map[string]*Database `json:"-"`
}

func (i *Instance) GetMName() string {
	return fmt.Sprintf("%s%s", i.Name, config.Meta.InstanceSuffix)
}

func (i *Instance) GetPath() string {
	return fmt.Sprintf("%s%s%s", config.Meta.Root, Separator, i.GetMName())
}

func (i *Instance) CreateDatabase(name string) *Database {
	db := &Database{
		DatabaseInfo: DatabaseInfo{
			Name: name,
			Text: name,
		},
		Tables:   make(map[string]*Table),
		Instance: i,
	}
	i.Databases[name] = db
	return db
}

func (i *Instance) CreateDatabaseWithInfo(info DatabaseInfo) *Database {
	db := &Database{
		DatabaseInfo: info,
		Tables:       make(map[string]*Table),
		Instance:     i,
	}
	i.Databases[info.Name] = db
	return db
}

func (i *Instance) GetDatabase(name string) *Database {
	return i.Databases[name]
}

func (i *Instance) Store() error {
	return storeInstance(i)
}

func (i *Instance) Load(data []byte) error {
	return json.Unmarshal(data, i)
}
