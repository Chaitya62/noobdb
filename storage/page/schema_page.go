package page

// 4096 bits
// 24 bits meta-data header
// slots and tuples

import (
	"fmt"
	"github.com/chaitya62/noobdb/type"
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

func (sp *SchemaPage) InsertTuple(tp SchemaTuple) error {
	tp_size := tp.GetSize()

	fp := sp.GetFreeSpacePointer()
	number_of_tps := sp.GetNumberOfTuples()
	slot_ends := (SLOT_OFFSET + SLOT_SIZE*number_of_tps)
	space_left := fp - slot_ends

	if uint64(space_left) < tp_size {
		//TODO: ADD CUSTOM STANDARD ERROR TYPES TO DATABASE
		return nil
	}

	// assuming all tuples fit in one page
	//TODO: Implement handling for Tuple OVERFLOW
	start_at := fp - uint16(tp_size)

	//TODO: Error handling
	fmt.Println("HERE")
	// intert tupple
	copy(sp._data[start_at:fp], tp.GetData())

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

const (
	SCHEMA_ID = iota
	SCHEMA_TABLE_ID
	SCHEMA_TABLE_NAME
	SCHEMA_COLUMN_POSITION
	SCHEMA_COLUMN_NAME
	SCHEMA_COLUMN_TYPE
)

// hard coding the tuples schema for schemaPage / Table
//TODO: CREATE A TUPLE INTERFACE
type SchemaTuple struct {
	//TODO: Add more columns for meta data / auto increament etc
	_data []type_.Type
}

func (st *SchemaTuple) Init() {
	schema_schema := []string{"INTEGER", "INTEGER", "VARCHAR", "INTEGER", "VARCHAR", "VARCHAR"}
	for _, s := range schema_schema {
		st._data = append(st._data, type_.TypeFactory(s))
	}

	st._data[SCHEMA_COLUMN_NAME].SetValue("id")
	st._data[SCHEMA_TABLE_NAME].SetValue("schema_table")
	st._data[SCHEMA_COLUMN_TYPE].SetValue("INTEGER")
}

func (st *SchemaTuple) GetSize() uint64 {
	var size_ uint64
	for _, s := range st._data {
		size_ += s.GetSize()
	}

	return size_
}

func (st *SchemaTuple) GetData() []byte {
	_data := make([]byte, st.GetSize())
	var index int
	for _, s := range st._data {
		index += copy(_data[index:], s.Serialize())
	}
	//return append(st._data[0].Serialize(), append(st._data[1].Serialize(), append(st._data[2].Serialize(), append(st._data[3].Serialize(), append(st._data[4].Serialize(), st._data[5].Serialize()...)...)...)...)...)
	return _data
}
