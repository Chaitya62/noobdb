package type_

// for now we will only have numeric types
//import "fmt"

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
	Deserialize(b []byte) error
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

func TypeFromTupleFactory(type_name string, data_ []byte, curr_pos uint64) (uint64, Type) {

	if len(data_) == 0 {
		return uint64(0), nil
	}

	type_obj := TypeFactory(type_name)
	type_obj.Deserialize(data_[curr_pos:])
	next_pos := curr_pos + type_obj.GetSize()

	return next_pos, type_obj

}
