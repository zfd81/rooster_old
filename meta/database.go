package meta

type DatabaseInfo struct {
	Name    string   `json:"name"`
	Text    string   `json:"text"`
	Comment string   `json:"comment"`
	Charset string   `json:"charset"`
	Tables  []*Table `json:"tables"`
}

type Database struct {
	DatabaseInfo
	TableMap map[string]*Table `json:"-"`
}

func NewDatabase(name string) *Database {
	tables := make([]*Table, 0, 50)
	tableMap := make(map[string]*Table)
	db := &Database{TableMap: tableMap}
	db.Name = name
	db.Text = name
	db.Tables = tables
	return db
}

func (d *Database) AddTable(table *Table) *Database {
	d.Tables = append(d.Tables, table)
	d.TableMap[table.Name] = table
	return d
}

func (d *Database) RemoveTable(name string) *Database {
	for i, v := range d.Tables {
		if v.Name == name {
			d.Tables = append(d.Tables[:i], d.Tables[i+1:]...)
			delete(d.TableMap, name)
			return d
		}
	}
	return d
}

func (d *Database) GetTable(name string) *Table {
	return d.TableMap[name]
}
