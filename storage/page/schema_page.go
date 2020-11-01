package page

// 4096 bits
// 24 bits meta-data header
// slots and tuples

import (
	"errors"
	"fmt"
	//"github.com/chaitya62/noobdb/type"
)

const TABLE_NAME_LIMT = 2048
const COLUMN_NAME_LIMT = 2048
const SLOT_OFFSET = 24
const SLOT_ID_SIZE = 4
const TUPLE_LOCATION_SIZE = 2
const SLOT_SIZE = 6

type SchemaPage struct {
	PageImpl
}

func (sp *SchemaPage) Init() {
	sp.SetFreeSpacePointer(4096)
}

func (sp *SchemaPage) shallowCopy(_page Page) {
	sp._data = _page.GetData()
}

func (sp *SchemaPage) GetHeader() []byte {
	// returning a slice returns the address space and not the copy
	// will have to re-think later if this is causes any unexpect behaviour
	return sp._data[:24]
}

// right now giving it only 2 bytes
// we have space for more
// but with 2 byte it means we can only have a page with maxsize 2^16 -> 65535
func (sp *SchemaPage) GetFreeSpacePointer() uint16 {
	return (uint16(sp._data[4]) | uint16(sp._data[5])<<8)
}

func (sp *SchemaPage) GetNumberOfTuples() uint16 {
	return (uint16(sp._data[6]) | uint16(sp._data[7])<<8)
}

func (sp *SchemaPage) UpdateNumberOfTuples(fsp uint16) error {
	sp._data[6] = byte(fsp)
	sp._data[7] = byte(fsp >> 8)
	return nil
}

// probably should be private
func (sp *SchemaPage) SetFreeSpacePointer(fsp uint16) error {
	sp._data[4] = byte(fsp)
	sp._data[5] = byte(fsp >> 8)
	return nil
}

func (sp *SchemaPage) GetSlotStart(i uint16) uint16 {
	x := SLOT_OFFSET + (SLOT_ID_SIZE+TUPLE_LOCATION_SIZE)*i
	return (uint16(sp._data[x+4]) | uint16(sp._data[x+5])<<8)
}

func (sp *SchemaPage) ReadTuple(i uint16) []byte {
	//TODO: error handling

	var end_at uint16
	end_at = 4096
	start_at := sp.GetSlotStart(i)
	if i > 0 {
		end_at = sp.GetSlotStart((i - 1))
	}

	return sp._data[start_at:end_at]

}

func (sp *SchemaPage) InsertTuple(tp Tuple) error {
	tp_size := tp.GetSize()

	fp := sp.GetFreeSpacePointer()
	number_of_tps := sp.GetNumberOfTuples()
	slot_ends := (SLOT_OFFSET + SLOT_SIZE*number_of_tps)
	space_left := fp - slot_ends

	if uint64(space_left) < tp_size {
		//TODO: ADD CUSTOM STANDARD ERROR TYPES TO DATABASE
		fmt.Println("PAGE is FULL")
		return errors.New("Page is full")
	}

	// assuming all tuples fit in one page
	//TODO: Implement handling for Tuple OVERFLOW
	start_at := fp - uint16(tp_size)

	//TODO: Error handling
	// intert tupple
	copy(sp._data[start_at:fp], tp.GetData())
	sp._dirty = true

	// insert slot
	//TODO: SLOT STRUCT ?
	var slot [(SLOT_ID_SIZE + TUPLE_LOCATION_SIZE)]byte

	slot_id := number_of_tps + 1
	slot[0] = byte(slot_id)
	slot[1] = byte(slot_id >> 8)
	slot[2] = byte(slot_id >> 16)
	slot[3] = byte(slot_id >> 24)
	slot[4] = byte(start_at)
	slot[5] = byte(start_at >> 8)

	// set slot
	copy(sp._data[slot_ends:slot_ends+(SLOT_SIZE)], slot[:])

	sp.SetFreeSpacePointer(start_at)
	sp.UpdateNumberOfTuples(number_of_tps + 1)

	return nil
	// check if space is there
}
