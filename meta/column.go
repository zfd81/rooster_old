package meta

type Column struct {
	Name         string
	Text         string
	Comment      string
	dataType     string
	length       int
	DefaultValue string
	State        SchemaState
}
