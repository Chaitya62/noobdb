package schema

// How to find details about a table
// Get the columns from a schema table
// Get the path of where the table is stored from table name
// Each table is a new file

import (
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
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
