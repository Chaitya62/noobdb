package schema

// How to find details about a table
// Get the columns from a schema table
// Get the path of where the table is stored from table name
// Each table is a new file

import (
	//	"fmt"
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
)

const (
	SCHEMA_ID = iota
	SCHEMA_TABLE_ID
	SCHEMA_TABLE_NAME
	SCHEMA_COLUMN_POSITION
	SCHEMA_COLUMN_NAME
	SCHEMA_COLUMN_TYPE
)

const schema_file = "schema.txt"

type SchemaTable struct {
	dmi     diskio.DiskManager
	bpm     *buffer.BufferPoolManager
	_schema [6]string
}

func (st *SchemaTable) Init() {

	st._schema = [6]string{"INTEGER", "INTEGER", "VARCHAR", "INTEGER", "VARCHAR", "VARCHAR"}
	st.dmi = diskio.NewDiskManagerImpl(schema_file)
	st.bpm = buffer.GetNewBufferPoolManager(3, st.dmi)
}

func (st *SchemaTable) GetDefaultRow() page.Tuple {
	tuple := new(page.TupleImpl)
	tuple.Init(st._schema[:])
	return tuple
}

func (st *SchemaTable) Insert(tuple page.Tuple) error {
	no_of_pages := st.dmi.GetNumberOfPages()
	var page_id uint32
	var _schemaPage *page.SchemaPage
	if no_of_pages == 0 {
		// create a new page
		// when schema table is empty
		page_id = uint32(0)
		_page := st.bpm.AllocatePage(page_id)
		_pageImpl := _page.(*page.PageImpl)
		_schemaPage = &page.SchemaPage{PageImpl: *_pageImpl}
		_schemaPage.Init()
	} else {
		page_id = uint32(no_of_pages - 1)
		_page := st.bpm.GetPage(page_id)
		_pageImpl := _page.(*page.PageImpl)
		_schemaPage = &page.SchemaPage{PageImpl: *_pageImpl}
	}

	st.bpm.PinPage(page_id)
	err := _schemaPage.InsertTuple(tuple)
	st.bpm.UnPinPage(page_id)

	// page is full
	// allocate new page
	if err != nil {
		page_id = uint32(page_id + 1)
		_page := st.bpm.AllocatePage(page_id)
		_pageImpl := _page.(*page.PageImpl)
		_schemaPage = &page.SchemaPage{PageImpl: *_pageImpl}
		_schemaPage.Init()

		// we assume each tuple can fit in a page of size 4096 bytes
		st.bpm.PinPage(page_id)
		err = _schemaPage.InsertTuple(tuple)
		st.bpm.UnPinPage(page_id)

	}

	return nil

}

//TODO: Convert this to an iterator design pattern
// https://ewencp.org/blog/golang-iterators/index.html
func (st *SchemaTable) Iterator() {
	no_of_pages := st.dmi.GetNumberOfPages()
	for i := 0; i < no_of_pages; i++ {

		// find a beter way to do this conversion
		_page := st.bpm.GetPage(uint32(i)).(*page.PageImpl)
		schemaPage := page.SchemaPage{PageImpl: *_page}
		no_of_tuples := schemaPage.GetNumberOfTuples()
		for j := uint16(0); j < no_of_tuples; j++ {
			tuple_data := schemaPage.ReadTuple(j)
			var schemaTuple page.TupleImpl
			schemaTuple.Init(st._schema[:])
			schemaTuple.ReadTuple(tuple_data)
			schemaTuple.PrintTuple()

		}
	}
}

// sanatize and persist everything to disk
func (st *SchemaTable) Close() {
	st.bpm.PersistAll()
}
