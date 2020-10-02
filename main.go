package main

import (
	"fmt"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
	//"github.com/chaitya62/noobdb/type"
)

func main() {

	x := []byte{12, 3, 4, 5, 6}
	x1 := []byte{12, 3, 4, 5, 6}

	x2 := append(x1, x...)

	fmt.Println(len(x2))

	//return

	//	var tx type_.Type
	//
	//	tx = new(type_.Integer)
	//
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("TypeID: ", tx.GetTypeID())
	//	fmt.Println("TypeName: ", tx.GetTypeName())
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//	tx.Deserialize([]byte{255, 255, 255, 255, 255, 255, 255, 255})
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//	tx.SetValue(int64(123234))
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//
	//	tx = new(type_.Boolean)
	//
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("TypeID: ", tx.GetTypeID())
	//	fmt.Println("TypeName: ", tx.GetTypeName())
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//	tx.Deserialize([]byte{1})
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//
	//	tx = new(type_.Varchar)
	//
	//	fmt.Println("Size: ", tx.GetSize())
	//
	//	fmt.Println("TypeID: ", tx.GetTypeID())
	//	tx.SetValue("WHAT")
	//	fmt.Println("TypeName: ", tx.GetTypeName())
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("Value: ", tx.GetValue())
	//	fmt.Println("Serialized: ", tx.Serialize())
	//	tx.Deserialize([]byte{5, 0, 0, 0, 'W', 'N', 'G', 80, '1'})
	//	fmt.Println("TypeID: ", tx.GetTypeID())
	//	fmt.Println("TypeName: ", tx.GetTypeName())
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("Value: ", tx.GetValue())
	//	tx.Deserialize([]byte{0, 0, 0, 0})
	//	fmt.Println("TypeID: ", tx.GetTypeID())
	//	fmt.Println("TypeName: ", tx.GetTypeName())
	//	fmt.Println("Size: ", tx.GetSize())
	//	fmt.Println("Value: ", tx.GetValue())

	// Creaet a schema table

	dmi := diskio.NewDiskManagerImpl("schema.txt")

	schemaPage := new(page.SchemaPage)
	schemaPage.Init()
	fmt.Println("TABLE DATA: ", schemaPage.GetData())
	fmt.Println("FSP: ", schemaPage.GetFreeSpacePointer())

	var schemaTuple page.SchemaTuple
	schemaTuple.Init()
	schemaTuple.InitDefault()

	schemaTuple.PrintTuple()

	fmt.Println("TUPLE DATA: ", schemaTuple.GetData())

	schemaPage.InsertTuple(schemaTuple)

	var schema_id int64
	var table_id int64
	var column_pos int64
	schema_id = 1
	table_id = 0
	column_pos = 1

	schemaTuple.SetValueFor(page.SCHEMA_ID, schema_id)
	schemaTuple.SetValueFor(page.SCHEMA_TABLE_ID, table_id)
	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_NAME, "first_name")
	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_TYPE, "VARCHAR")
	schemaTuple.SetValueFor(page.SCHEMA_COLUMN_POSITION, column_pos)

	schemaTuple.PrintTuple()

	fmt.Println("TUPLE SIZE: ", schemaTuple.GetSize())

	schemaPage.InsertTuple(schemaTuple)

	fmt.Println("TABLE DATA: ", schemaPage.GetData())
	fmt.Println("FSP: ", schemaPage.GetFreeSpacePointer())

	//x := new(page.PageImpl)
	dmi.WritePage(0, schemaPage)

	//  Read Schema table from a page

	Page := dmi.ReadPage(0).(*page.PageImpl)

	schemaPageR := page.SchemaPage{PageImpl: *Page}

	fmt.Println(schemaPageR.GetData())

	tuple := schemaPageR.ReadTuple(0)
	tuple2 := schemaPageR.ReadTuple(1)

	fmt.Println(tuple)
	fmt.Println(tuple2)

	var schemaTupleR page.SchemaTuple

	schemaTupleR.Init()

	schemaTupleR.ReadTuple(tuple)
	schemaTupleR.PrintTuple()

	schemaTupleR.ReadTuple(tuple2)
	schemaTupleR.PrintTuple()

	//_data := x.GetData()
	//for i := 0; i < 1000; i++ {
	//	_data[i] = byte(i)
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
