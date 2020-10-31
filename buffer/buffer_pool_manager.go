package buffer

import (
	"container/list"
	"fmt"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
)

type BufferPoolManager struct {
	pages      []page.Page
	dmi        diskio.DiskManager
	free_list  *list.List
	page_table *PageTable
	replacer   Replacer
}

func GetNewBufferPoolManager(buffer_size int, dmi diskio.DiskManager) *BufferPoolManager {
	bpm := new(BufferPoolManager)
	bpm.Init(buffer_size, dmi)
	return bpm
}

func (bpm *BufferPoolManager) Init(buffer_size int, dmi diskio.DiskManager) {
	bpm.pages = make([]page.Page, buffer_size)
	bpm.dmi = dmi
	bpm.free_list = list.New()
	bpm.replacer = NewClockReplacer(buffer_size)
	for i := 0; i < buffer_size; i++ {
		bpm.free_list.PushBack(i)
	}
	bpm.page_table = &PageTable{}
	bpm.page_table.Init()
}

func (bpm *BufferPoolManager) PinPage(page_id uint32) bool {
	val, ok := bpm.page_table.Get(page_id)
	if ok {
		bpm.replacer.PinPage(val)
	}
	return ok
}

func (bpm *BufferPoolManager) UnPinPage(page_id uint32) bool {
	val, ok := bpm.page_table.Get(page_id)
	if ok {
		bpm.replacer.UnPinPage(val)
	}
	return ok
}

// not thread safe
func (bpm *BufferPoolManager) GetPage(page_id uint32) page.Page {
	// if page is not there
	// fetch it via dmi
	if val, ok := bpm.page_table.Get(page_id); ok {
		return bpm.pages[val]
	}

	fmt.Println("Didn't find page, Allocating new memory")

	var free_frame_i int
	free_frame_i = -1
	if bpm.free_list.Len() != 0 {
		front := bpm.free_list.Front()
		free_frame_i = front.Value.(int)
		bpm.free_list.Remove(front)
	} else {
		frame_i, ok := bpm.replacer.GetNextVictim()
		if ok {
			free_frame_i = frame_i
			// add logic to if page is dirty write it to disk
		} else {
			fmt.Println("Couldn't find a page")
		}
	}

	fmt.Println("Allocating frame", free_frame_i, page_id)

	if free_frame_i == -1 {
		return page.InvalidPage()
	}

	bpm.pages[free_frame_i] = bpm.dmi.ReadPage(page_id)
	bpm.page_table.InsertOrUpdate(page_id, free_frame_i)
	return bpm.pages[free_frame_i]

}
