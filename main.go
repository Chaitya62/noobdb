package main

import (
	"fmt"
	"github.com/chaitya62/mygodb/storage/disk"
	"github.com/chaitya62/mygodb/storage/page"
)

func main() {

	x := new(page.PageImpl)
	dmi := diskio.NewDiskManagerImpl("db.txt")
	_data := x.GetData()
	for i := 0; i < 1000; i++ {
		_data[i] = byte(i)
	}
	x.SetData(_data[:])
	dmi.WritePage(10, x)

	var xn page.Page

	xn = dmi.ReadPage(12)

	fmt.Printf("%v\n", xn.GetData())
	fmt.Println("%v", xn.GetPageId())
	//x.ResetMemory()
	//fmt.Printf("%v", x.GetData())
}
