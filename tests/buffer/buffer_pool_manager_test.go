package buffer_test

import (
	//	"fmt"
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/storage/disk"
	"github.com/chaitya62/noobdb/storage/page"
	"github.com/chaitya62/noobdb/tests/helpers"
	"io/ioutil"
	"os"
	"testing"
)

func TestBufferPoolManager(t *testing.T) {
	path, err := os.Getwd()
	tmpdir, err := ioutil.TempDir(path, "tmp_dir*")

	// should be panic ?
	if err != nil {
		t.Errorf("Error creating tmp dir %v", err)
	}

	tmpfile, err := ioutil.TempFile(tmpdir, "db_*")

	defer func() {
		// clean up
		os.Remove(tmpfile.Name())
		os.RemoveAll(tmpdir)
	}()

	dmi := diskio.NewDiskManagerImpl(tmpfile.Name())

	//TODO: Add a init function for BufferPoolManager which takes care of DiskManager
	bpm := new(buffer.BufferPoolManager)
	bpm.Init(5, dmi)

	t.Run("should return invalid page if it doesn't exists", func(t *testing.T) {
		_page := bpm.GetPage(uint32(1245))

		helpers.Equals(t, uint32(page.INVALID_PAGE_ID), _page.GetPageId())

	})

}
