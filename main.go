package main

import (
	"fmt"
	//"github.com/chaitya62/noobdb/storage/disk"
	//"github.com/chaitya62/noobdb/storage/page"
	"github.com/chaitya62/noobdb/type"
)

func main() {

	x := []byte{12, 3, 4, 5, 6}
	x1 := []byte{12, 3, 4, 5, 6}

	x2 := append(x1, x...)

	fmt.Println(len(x2))

	//return

	fmt.Printf("%v %v %v %v\n\n", type_.INTEGER, type_.BOOLEAN, type_.DECIMAL, type_.REALNUMBER)

	var tx type_.Type

	tx = new(type_.Integer)

	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("TypeID: ", tx.GetTypeID())
	fmt.Println("TypeName: ", tx.GetTypeName())
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())
	tx.Deserialize([]byte{255, 255, 255, 255, 255, 255, 255, 255})
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())
	tx.SetValue(int64(123234))
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())

	tx = new(type_.Boolean)

	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("TypeID: ", tx.GetTypeID())
	fmt.Println("TypeName: ", tx.GetTypeName())
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())
	tx.Deserialize([]byte{1})
	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())

	tx = new(type_.Varchar)

	fmt.Println("Size: ", tx.GetSize())

	fmt.Println("TypeID: ", tx.GetTypeID())
	tx.SetValue("WHAT")
	fmt.Println("TypeName: ", tx.GetTypeName())
	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("Value: ", tx.GetValue())
	fmt.Println("Serialized: ", tx.Serialize())
	tx.Deserialize([]byte{5, 0, 0, 0, 'W', 'N', 'G', 80, '1'})
	fmt.Println("TypeID: ", tx.GetTypeID())
	fmt.Println("TypeName: ", tx.GetTypeName())
	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("Value: ", tx.GetValue())
	tx.Deserialize([]byte{0, 0, 0, 0})
	fmt.Println("TypeID: ", tx.GetTypeID())
	fmt.Println("TypeName: ", tx.GetTypeName())
	fmt.Println("Size: ", tx.GetSize())
	fmt.Println("Value: ", tx.GetValue())

	//x := new(page.PageImpl)
	//dmi := diskio.NewDiskManagerImpl("db.txt")
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
