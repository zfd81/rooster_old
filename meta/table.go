package meta

import (
	"encoding/json"
	"fmt"
)

type TableInfo struct {
	Name    string      `json:"name"`
	Text    string      `json:"text,omitempty"`
	Comment string      `json:"comment,omitempty"`
	Charset string      `json:"charset"`
	State   SchemaState `json:"state"`
	Columns []*Column   `json:"cols"`
}

type Table struct {
	TableInfo
	Database *Database `json:"-"`
}

func (t *Table) GetMName() string {
	return fmt.Sprintf("%s%s", t.Name, config.Meta.TableSuffix)
}

func (t *Table) GetPath() string {
	return fmt.Sprintf("%s%s%s", t.Database.GetPath(), Separator, t.GetMName())
}

func (t *Table) CreateColumn(name string, dataType string) *Column {
	col := &Column{
		Name:     name,
		Text:     name,
		DataType: dataType,
	}
	col.Index = len(t.Columns)
	t.Columns = append(t.Columns, col)
	return col
}

func (t *Table) RemoveColumn(name string) *Table {
	for i, v := range t.Columns {
		if v.Name == name {
			t.Columns = append(t.Columns[:i], t.Columns[i+1:]...)
			break
		}
	}
	for i, v := range t.Columns {
		v.Index = i
	}
	return t
}

func (t *Table) ModifyColumn(col *Column) *Table {
	for i, v := range t.Columns {
		if v.Name == col.Name {
			col.Index = v.Index
			t.Columns[i] = col
			break
		}
	}
	return t
}

func (t *Table) GetColumnIndex(name string) int {
	for i, v := range t.Columns {
		if v.Name == name {
			return i
		}
	}
	return -1
}

func (t *Table) GetColumn(name string) *Column {
	for _, v := range t.Columns {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (t *Table) Store() error {
	return storeTable(t)
}

func (t *Table) Load(data []byte) error {
	err := json.Unmarshal(data, t)
	if err == nil {
		t.Database.Tables[t.Name] = t
	}
	return err
}
