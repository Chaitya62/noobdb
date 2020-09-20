package type_

// for now we will only have numeric types
// introduce characters, text etc later

const (
	BOOLEAN = iota
	INTEGER
	DECIMAL
	REALNUMBER
	VARCHAR
)

type Type interface {
	GetTypeID() int8
	GetTypeName() string
	GetSize() uint64 // size in number of bytes
	GetValue() interface{}
	SetValue(interface{})
	Serialize() []byte
	Deserialize(b []byte)
}
