package meta

type Column struct {
	Name         string      `json:"name"`
	Text         string      `json:"text,omitempty"`
	Comment      string      `json:"comment,omitempty"`
	DataType     string      `json:"dataType"`
	Length       int         `json:"length,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	State        SchemaState `json:"state"`
	Index        int         `json:"index"`
}
