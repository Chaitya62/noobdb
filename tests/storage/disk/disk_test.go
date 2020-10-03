package diskio_test

import (
	"fmt"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
	"github.com/chaitya62/noobdb/tests/helpers"
	"io/ioutil"
	"os"
	"testing"
)

func TestDiskManagerImpl(t *testing.T) {
	// setup
	path, err := os.Getwd()
	tmpdir, err := ioutil.TempDir(path, "tmp_dir*")

	if err != nil {
		t.Errorf("Error creating tmp dir %v", err)
	}

	tmpfile, err := ioutil.TempFile(tmpdir, "db_*")

	defer func() {
		// clean up
		os.Remove(tmpfile.Name())
		os.RemoveAll(tmpdir)
	}()

	fmt.Printf("Creating temp directory: %v and temp file: %v\n", tmpdir, tmpfile.Name())

	t.Run("Implements DiskManager interface", func(t *testing.T) {
		diskManagerImpl := &diskio.DiskManagerImpl{}
		_, ok := interface{}(diskManagerImpl).(diskio.DiskManager)
		if ok != true {
			t.Errorf("DiskManagerImpl does not implement diskio.DiskManager")
		}
	})

	t.Run("NewDiskManagerImpl", func(t *testing.T) {
		// couldn't mock initFile
		dmi := diskio.NewDiskManagerImpl(tmpfile.Name())
		helpers.EqualTypes(t, &diskio.DiskManagerImpl{}, dmi)
	})

	t.Run("WritePage", func(t *testing.T) {
		t.Run("Write a valid Page", func(t *testing.T) {

			dmi := diskio.NewDiskManagerImpl(tmpfile.Name())

			file_info_before, err1 := tmpfile.Stat()
			if err1 != nil {
				t.Errorf("Failed getting file stats: %v", err1)
			}

			dummy_page := &page.PageImpl{}

			ok := dmi.WritePage(dummy_page.GetPageId(), dummy_page)
			if ok != nil {
				t.Errorf("Failed writing Page")
			}

			file_info_after, err2 := tmpfile.Stat()
			if err2 != nil {
				t.Errorf("Failed getting file stats: %v", err2)
			}

			helpers.NotEquals(t, file_info_before.Size(), file_info_after.Size())
		})

		t.Run("Write a Invalid Page", func(t *testing.T) {
			dmi := diskio.NewDiskManagerImpl(tmpfile.Name())

			file_info_before, err1 := tmpfile.Stat()
			if err1 != nil {
				t.Errorf("Failed getting file stats: %v", err1)
			}

			dummy_page := page.InvalidPage()

			ok := dmi.WritePage(dummy_page.GetPageId(), dummy_page)
			if ok == nil {
				t.Errorf("Should have returned an error")
			}

			file_info_after, err2 := tmpfile.Stat()
			if err2 != nil {
				t.Errorf("Failed getting file stats: %v", err2)
			}

			helpers.Equals(t, file_info_before.Size(), file_info_after.Size())
		})
	})

	t.Run("ReadPage", func(t *testing.T) {
		// setup write a Page first before reading

		t.Run("Read a valid page", func(t *testing.T) {
			dmi := diskio.NewDiskManagerImpl(tmpfile.Name())
			dummpy_page := &page.PageImpl{}
			pd := make([]byte, page.PAGE_SIZE)
			pd[page.PAGE_SIZE-123] = 12
			pd[page.PAGE_SIZE-12] = 1
			dummpy_page.SetData(pd)
			dmi.WritePage(dummpy_page.GetPageId(), dummpy_page)
			read_page := dmi.ReadPage(0)

			helpers.EqualSlices(t, dummpy_page.GetData(), read_page.GetData())
		})

		t.Run("Read a valid page", func(t *testing.T) {
			dmi := diskio.NewDiskManagerImpl(tmpfile.Name())
			// this is a flaky test if the same file is used in any of the above test
			// and they somehow manage to write this many pages
			// for now this should work
			// can possibly get file stats and calculate a invalid page id
			invalid_page_id := uint32(3464)
			read_page := dmi.ReadPage(invalid_page_id)

			helpers.Equals(t, uint32(page.INVALID_PAGE_ID), read_page.GetPageId())
		})

	})
}
