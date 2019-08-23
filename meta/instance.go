package meta

type Instance struct {
	Name      string
	Text      string
	Comment   string
	Charset   string
	State     SchemaState
	Databases []*Database
}

func NewInstance() *Instance {
	return &Instance{}
}
