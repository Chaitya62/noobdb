package type_

import (
	"fmt"
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

func (b *Boolean) Deserialize(data []byte) {
	if len(data) != 1 {
		fmt.Println("Invalid memory block")
	}
	b.val = (data[0]&1 == 1)
}

func (b *Boolean) GetValue() interface{} {
	return b.val
}
