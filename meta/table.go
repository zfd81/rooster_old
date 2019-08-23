package meta

type Table struct {
	Name    string
	Text    string
	Comment string
	Charset string
	State   SchemaState
	Columns []*Column
}
