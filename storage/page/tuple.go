package page

import (
	"fmt"
	"github.com/chaitya62/noobdb/type"
)

type Tuple interface {
	ReadTuple(data_ []byte)
	SetValueFor(column_i uint64, val interface{})
	GetValueFor(column_i uint64) interface{}
	GetSize() uint64
	GetData() []byte
	PrintTuple()
}

type TupleImpl struct {
	//TODO: Add more columns for meta data / auto increament etc
	_data   []type_.Type
	_schema []string
}

func (st *TupleImpl) Init(s_ []string) {
	// handler if init called multiple times
	st._schema = append(st._schema, s_...)
	for _, s := range st._schema {
		st._data = append(st._data, type_.TypeFactory(s))
	}
}

func (st *TupleImpl) ReadTuple(data_ []byte) {
	if len(data_) == 0 {
		return
	}
	var curr_pos uint64
	st._data = []type_.Type{}
	for _, s := range st._schema {
		next_pos, type_obj := type_.TypeFromTupleFactory(s, data_, curr_pos)
		curr_pos = next_pos
		st._data = append(st._data, type_obj)
	}
}

func (st *TupleImpl) PrintTuple() {
	fmt.Printf("| ")
	for _, s := range st._data {
		fmt.Printf(" %v |", s.GetValue())
	}
	fmt.Printf("\n")
}

func (st *TupleImpl) SetValueFor(column_i uint64, val interface{}) {
	st._data[column_i].SetValue(val)
}

func (st *TupleImpl) GetValueFor(column_i uint64) interface{} {
	return st._data[column_i].GetValue()
}

func (st *TupleImpl) GetSize() uint64 {
	var size_ uint64
	for _, s := range st._data {
		size_ += s.GetSize()
	}

	return size_
}

func (st *TupleImpl) GetData() []byte {
	_data := make([]byte, st.GetSize())
	var index int
	for _, s := range st._data {
		index += copy(_data[index:], s.Serialize())
	}
	return _data
}
