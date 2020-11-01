package main

import (
	"fmt"
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/schema"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
	"time"
	//"github.com/chaitya62/noobdb/type"
)

func printTuple(tuple page.TupleImpl) {
	tuple.PrintTuple()
}

func accessMemory(i int, bpm *buffer.BufferPoolManager, schema_schema []string, tupleChannel chan<- page.TupleImpl) {

	time.Sleep(100 * time.Millisecond)

	Page := bpm.GetPage(uint32(i)).(*page.PageImpl)
	//Page := dmi.ReadPage(uint32(i)).(*page.PageImpl)

	schemaPageR := page.SchemaPage{PageImpl: *Page}

	if schemaPageR.GetPageId() == page.INVALID_PAGE_ID {
		return
	}

	//fmt.Println(schemaPageR.GetData())

	tuple := schemaPageR.ReadTuple(0)
	tuple2 := schemaPageR.ReadTuple(1)

	//fmt.Println(tuple)
	//fmt.Println(tuple2)

	if i >= 7 {
		bpm.PinPage(uint32(i))
	}

	var schemaTupleR page.TupleImpl

	schemaTupleR.Init(schema_schema[:])

	schemaTupleR.ReadTuple(tuple)
	//schemaTupleR.PrintTuple()
	//tupleChannel <- schemaTupleR

	schemaTupleR.ReadTuple(tuple2)
	//schemaTupleR.PrintTuple()
	tupleChannel <- schemaTupleR

}

func testBufferPoolManagerAsync() {

	var TupleChannel chan page.TupleImpl
	dmi := diskio.NewDiskManagerImpl("schema.txt")

	bpm := new(buffer.BufferPoolManager)
	bpm.Init(5, dmi)

	// move this logic to schema module
	schema_schema := [6]string{"INTEGER", "INTEGER", "VARCHAR", "INTEGER", "VARCHAR", "VARCHAR"}

	//  Read Schema table from a page

	TupleChannel = make(chan page.TupleImpl, 1)

	go func() {

		count := 0
		for {
			select {
			case tuple := <-TupleChannel:
				fmt.Println("cnt", count)
				count = count + 1
				printTuple(tuple)
			default:
				println("Waiting for data")
				time.Sleep(100 * time.Millisecond)
			}
		}

	}()

	for i := 0; i < 10; i++ {
		accessMemory(i, bpm, schema_schema[:], TupleChannel)
	}

}

func main() {
	schema_table := new(schema.SchemaTable)
	schema_table.Init()

	//for i := 0; i < 1000; i++ {
	//	tuple := schema_table.GetDefaultRow()
	//	tuple.SetValueFor(schema.SCHEMA_ID, int64(i))
	//	tuple.SetValueFor(schema.SCHEMA_TABLE_ID, int64(0i))
	//	tuple.SetValueFor(schema.SCHEMA_TABLE_NAME, "first_table")
	//	tuple.SetValueFor(schema.SCHEMA_COLUMN_NAME, "id")
	//	tuple.SetValueFor(schema.SCHEMA_COLUMN_POSITION, int64(1))
	//	tuple.SetValueFor(schema.SCHEMA_COLUMN_TYPE, "INTEGER")
	//	schema_table.Insert(tuple)
	//}

	schema_table.Close()
	schema_table.Iterator()
}
