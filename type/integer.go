package type_

import (
	"errors"
)

const INTEGER_SIZE = 8

type Integer struct {
	val int64
}

func (i *Integer) GetTypeID() int8 {
	return INTEGER
}

func (i *Integer) GetTypeName() string {
	return "INTEGER"
}

func (i *Integer) SetValue(v interface{}) {
	i.val = v.(int64)
}

// https://golang.org/src/encoding/binary/binary.go
func (i *Integer) Serialize() []byte {
	data := make([]byte, INTEGER_SIZE)

	data[0] = byte(i.val) // first 8-bit LSB
	data[1] = byte(i.val >> 8)
	data[2] = byte(i.val >> 16)
	data[3] = byte(i.val >> 24)
	data[4] = byte(i.val >> 32)
	data[5] = byte(i.val >> 40)
	data[6] = byte(i.val >> 48)
	data[7] = byte(i.val >> 56) // last 8-bit MSB

	return data
}

func (i *Integer) GetSize() uint64 {
	return uint64(INTEGER_SIZE)
}

func (i *Integer) Deserialize(data []byte) error {
	if len(data) < 8 {
		return errors.New("Invalid size of byte slice")
	}
	i.val = (int64(data[7])<<56 | int64(data[6])<<48 | int64(data[5])<<40 | int64(data[4])<<32 | int64(data[3])<<24 | int64(data[2])<<16 | int64(data[1])<<8 | int64(data[0]))
	return nil
}

func (i *Integer) GetValue() interface{} {
	return i.val
}
