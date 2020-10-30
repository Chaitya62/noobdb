package main

import (
	"fmt"
	"github.com/chaitya62/noobdb/buffer"
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

func main() {

	var TupleChannel chan page.TupleImpl
	dmi := diskio.NewDiskManagerImpl("schema.txt")

	bpm := new(buffer.BufferPoolManager)
	bpm.Init(5, dmi)

	// move this logic to schema module
	schema_schema := [6]string{"INTEGER", "INTEGER", "VARCHAR", "INTEGER", "VARCHAR", "VARCHAR"}

	//for i := 0; i < 10; i++ {
	//	schemaPage := new(page.SchemaPage)
	//	schemaPage.Init()
	//	fmt.Println("TABLE DATA: ", schemaPage.GetData())
	//	fmt.Println("FSP: ", schemaPage.GetFreeSpacePointer())

	//	var schemaTuple page.TupleImpl
	//	schemaTuple.Init(schema_schema[:])
	//	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_NAME, "id")
	//	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_TYPE, "INTEGER")
	//	schemaTuple.SetValueFor(page.SCHEMA_TABLE_NAME, "schema_table")

	//	schemaTuple.PrintTuple()

	//	fmt.Println("TUPLE DATA: ", schemaTuple.GetData())

	//	schemaPage.InsertTuple(&schemaTuple)

	//	var schema_id int64
	//	var table_id int64
	//	var column_pos int64
	//	schema_id = 1
	//	table_id = int64(i)
	//	column_pos = 1

	//	schemaTuple.SetValueFor(page.SCHEMA_ID, schema_id)
	//	schemaTuple.SetValueFor(page.SCHEMA_TABLE_ID, table_id)
	//	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_NAME, "first_name")
	//	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_TYPE, "VARCHAR")
	//	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_POSITION, column_pos)

	//	schemaTuple.PrintTuple()

	//	fmt.Println("TUPLE SIZE: ", schemaTuple.GetSize())

	//	schemaPage.InsertTuple(&schemaTuple)

	//	fmt.Println("TABLE DATA: ", schemaPage.GetData())
	//	fmt.Println("FSP: ", schemaPage.GetFreeSpacePointer())

	//	//x := new(page.PageImpl)
	//	dmi.WritePage(uint32(i), schemaPage)

	//}

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

	//_data := x.GetData()
	//for i := 0; i < 1000; i++ {
	//}
	//x.SetData(_data[:])
	//dmi.WritePage(10, x)

	//var xn page.Page

	//xn = dmi.ReadPage(12)

	//fmt.Printf("%v\n", xn.GetData())
	//fmt.Println("%v", xn.GetPageId())
	//x.ResetMemory()
	//fmt.Printf("%v", x.GetData())
}
