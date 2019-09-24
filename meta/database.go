package meta

import (
	"encoding/json"
	"fmt"
)

type DatabaseInfo struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	Comment string `json:"comment,omitempty"`
	Charset string `json:"charset"`
}

type Database struct {
	DatabaseInfo
	Tables   map[string]*Table `json:"-"`
	Instance *Instance         `json:"-"`
}

func (d *Database) GetMName() string {
	return fmt.Sprintf("%s%s", config.Meta.DatabasePrefix, d.Name)
}

func (d *Database) GetPath() string {
	return fmt.Sprintf("%s%s%s", d.Instance.GetPath(), Separator, d.GetMName())
}

func (d *Database) CreateTable(name string) *Table {
	tbl := &Table{
		Name:     name,
		Text:     name,
		Columns:  make([]*Column, 0, 10),
		Database: d,
	}
	d.Tables[name] = tbl
	return tbl
}

func (d *Database) RemoveTable(name string) *Database {
	delete(d.Tables, name)
	return d
}

func (d *Database) GetTable(name string) *Table {
	return d.Tables[name]
}

func (d *Database) Store() error {
	return storeDatabase(d)
}

func (d *Database) Load(data []byte) error {
	err := json.Unmarshal(data, d)
	if err == nil {
		d.Instance.Databases[d.Name] = d
	}
	return err
}
