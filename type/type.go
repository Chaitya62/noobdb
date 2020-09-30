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

func TypeFactory(type_name string) Type {

	var type_obj Type

	switch type_name {
	case "INTEGER":
		type_obj = new(Integer)
	case "VARCHAR":
		type_obj = new(Varchar)
	case "BOOLEAN":
		type_obj = new(Boolean)
	}

	return type_obj
}
