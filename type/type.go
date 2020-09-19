package type_

// for now we will only have numeric types
// introduce characters, text etc later

const (
	BOOLEAN = iota
	INTEGER
	DECIMAL
	REALNUMBER
)

type Type interface {
	GetTypeID() int8
	GetTypeName() string
	GetValue() interface{}
	SetValue(interface{})
	Serialize() []byte
	Deserialize(b []byte)
}
