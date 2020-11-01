package buffer

import (
	"container/list"
	"fmt"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
)

type BufferPoolManager struct {
	pages       []page.Page
	dmi         diskio.DiskManager
	free_list   *list.List
	page_table  *PageTable
	replacer    Replacer
	buffer_size int
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
	bpm.buffer_size = buffer_size
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

func (bpm *BufferPoolManager) PersistAll() {
	for i := 0; i < bpm.buffer_size; i++ {
		if bpm.pages[i] != nil {
			if bpm.pages[i].IsDirty() {
				bpm.dmi.WritePage(bpm.pages[i].GetPageId(), bpm.pages[i])
			}
		}
	}
}

func (bpm *BufferPoolManager) AllocatePage(page_id uint32) page.Page {
	_page := bpm.GetPage(page_id) // this will written a page with invalid id
	_page.SetPageId(page_id)
	bpm.dmi.WritePage(page_id, _page)
	return _page
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

			// delete entry from PageTable
			bpm.page_table.Delete(bpm.pages[frame_i].GetPageId())

			// persist dirty page to disk
			if bpm.pages[frame_i].IsDirty() {
				bpm.dmi.WritePage(bpm.pages[frame_i].GetPageId(), bpm.pages[frame_i])
			}

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
