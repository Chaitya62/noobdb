package type_

import (
	"strconv"
)

type Varchar struct {
	val  string
	_len uint32
}

func (vc *Varchar) GetTypeID() int8 {
	return INTEGER
}

func (vc *Varchar) GetTypeName() string {
	type_name := "VARCHAR(" + strconv.Itoa(int(vc._len)) + ")"
	return type_name
}

func (vc *Varchar) SetValue(v interface{}) {
	vc.val = v.(string)
	vc._len = uint32(len(vc.val))
}

func (vc *Varchar) Serialize() []byte {

	// 32-bit number t | char array for length t |

	string_data := []byte(vc.val)
	data := make([]byte, 4)

	data[0] = byte(vc._len) // first 8-bit LSB
	data[1] = byte(vc._len >> 8)
	data[2] = byte(vc._len >> 16)
	data[3] = byte(vc._len >> 24) // last 8-bit MSB

	result := append(data, string_data...)
	return result
}

func (vc *Varchar) GetSize() uint64 {
	return uint64(4 + uint64(vc._len))
}

// this seems tricky
// for now assuming that we will get the entire memory location with length and string
// might have to re-think the implementation later if the api becomes tedious to call
func (vc *Varchar) Deserialize(data []byte) {
	vc._len = (uint32(data[3])<<24 | uint32(data[2])<<16 | uint32(data[1])<<8 | uint32(data[0]))
	vc.val = string(data[4 : 4+vc._len])
}

func (vc *Varchar) GetValue() interface{} {
	return vc.val
}
