package type_

import (
	"errors"
)

type Boolean struct {
	val bool
}

func (b *Boolean) GetTypeID() int8 {
	return BOOLEAN
}

func (b *Boolean) GetTypeName() string {
	return "BOOLEAN"
}

func (b *Boolean) SetValue(v interface{}) {
	b.val = v.(bool)
}

func (b *Boolean) Serialize() []byte {
	data := make([]byte, 1)
	if b.val {
		data[0] = 1
	}
	return data
}

func (b *Boolean) Deserialize(data []byte) error {
	if len(data) < 1 {
		return errors.New("Invalid size of byte slice")
	}
	b.val = (data[0]&1 == 1)

	return nil
}

func (b *Boolean) GetSize() uint64 {
	return uint64(1)
}

func (b *Boolean) GetValue() interface{} {
	return b.val
}
