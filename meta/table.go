package meta

type Table struct {
	Name    string      `json:"name"`
	Text    string      `json:"text"`
	Comment string      `json:"comment,omitempty"`
	Charset string      `json:"charset"`
	State   SchemaState `json:"state"`
	Columns []*Column   `json:"cols"`
}

func (t *Table) AddColumn(col *Column) *Table {
	t.Columns = append(t.Columns, col)
	return t
}

func (t *Table) RemoveColumn(name string) *Table {
	for i, v := range t.Columns {
		if v.Name == name {
			t.Columns = append(t.Columns[:i], t.Columns[i+1:]...)
			return t
		}
	}
	return t
}

func (t *Table) ModifyColumn(col *Column) *Table {
	for i, v := range t.Columns {
		if v.Name == col.Name {
			t.Columns[i] = col
			return t
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

func NewTable(name string) *Table {
	cols := make([]*Column, 0, 10)
	return &Table{Name: name, Text: name, Columns: cols}
}
