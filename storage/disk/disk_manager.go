package diskio

import (
	"errors"
	"fmt"
	"github.com/chaitya62/noobdb/storage/page"
	"os"
)

//TODO: Implement page delete logic to freeup disk space

type DiskManager interface {
	WritePage(page_id uint32, pg page.Page) error
	ReadPage(page_id uint32) page.Page
}

type DiskManagerImpl struct {
	db_file_name string
	db_file      *os.File
}

func NewDiskManagerImpl(db_file_name string) DiskManager {
	dmi := DiskManagerImpl{db_file_name, nil}
	dmi.initFile()
	return &dmi
}

func (dmi *DiskManagerImpl) initFile() error {
	if dmi.db_file == nil {
		file, err := os.OpenFile(dmi.db_file_name, os.O_RDWR|os.O_CREATE, 0644)
		dmi.db_file = file
		return err
	}
	return nil
}

func (dmi *DiskManagerImpl) WritePage(page_id uint32, pg page.Page) error {

	if page_id == page.INVALID_PAGE_ID {
		return errors.New("Cannot write Invalid Page")
	}

	var offset int64
	offset = page.PAGE_SIZE * int64(page_id)

	_page_data := pg.GetData()
	_, err := dmi.db_file.WriteAt(_page_data[:], offset)

	if err != nil {
		return err
	}

	err = dmi.db_file.Sync()

	return err
}

func (dmi *DiskManagerImpl) ReadPage(page_id uint32) page.Page {

	var offset int64
	offset = page.PAGE_SIZE * int64(page_id)

	file_info, err1 := dmi.db_file.Stat()

	if err1 != nil {
		fmt.Println(err1)
		return page.InvalidPage()
	}

	if file_info.Size() < offset {
		fmt.Println("No such page exists")
		return page.InvalidPage()
	}

	_page_data := [page.PAGE_SIZE]byte{}

	n, err := dmi.db_file.ReadAt(_page_data[:], offset)

	fmt.Printf("Read %v bytes\n", n)

	if err != nil {
		// handle panic later
	}

	newPage := new(page.PageImpl)
	newPage.SetData(_page_data[:])
	newPage.SetPageId(page_id)
	return newPage
}
