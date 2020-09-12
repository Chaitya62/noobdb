package main

import (
	"fmt"
	"github.com/chaitya62/mygodb/storage/page"
)

func main() {

	x := new(page.PageImpl)

	//fmt.Println(x.test())
	fmt.Printf("%v\n", x.GetData())
	fmt.Println("%v", x.GetPageId())
	x.ResetMemory()
	fmt.Printf("%v", x.GetData())
}
