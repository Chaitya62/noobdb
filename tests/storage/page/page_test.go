package page_test

import (
	"github.com/chaitya62/noobdb/storage/page"
	"github.com/chaitya62/noobdb/tests/helpers"
	"testing"
)

func TestPageImpl(t *testing.T) {

	t.Run("Implements Page interface", func(t *testing.T) {
		pageImpl := &page.PageImpl{}
		_, ok := interface{}(pageImpl).(page.Page)
		if ok != true {
			t.Errorf("PageImpl does not implement page.Page")
		}
	})

	t.Run("SetData", func(t *testing.T) {

		t.Run("SetData - page_id and data", func(t *testing.T) {
			pageImpl := &page.PageImpl{}
			data_slice := make([]byte, page.PAGE_SIZE)

			expected_page_id := uint32(12)
			data_slice[0] = byte(expected_page_id)
			data_slice[1] = byte(expected_page_id << 8)
			data_slice[2] = byte(expected_page_id << 16)
			data_slice[3] = byte(expected_page_id << 24)

			ok := pageImpl.SetData(data_slice)
			if ok != nil {
				t.Errorf("pageImpl.SetData failed")
			}

			helpers.Equals(t, expected_page_id, pageImpl.GetPageId())
			helpers.EqualSlices(t, data_slice, pageImpl.GetData())
		})

		t.Run("SetData - invalid slice", func(t *testing.T) {
			pageImpl := &page.PageImpl{}

			invalid_page_size := 1235
			data_slice := make([]byte, invalid_page_size)

			ok := pageImpl.SetData(data_slice)

			if ok == nil {
				t.Errorf("Page data cannot be of size %v", invalid_page_size)
			}

		})

	})

	t.Run("ResetMemory", func(t *testing.T) {
		pageImpl := &page.PageImpl{}
		data_slice := make([]byte, page.PAGE_SIZE)

		expected_page_id := uint32(12)
		data_slice[0] = byte(expected_page_id)
		data_slice[1] = byte(expected_page_id << 8)
		data_slice[2] = byte(expected_page_id << 16)
		data_slice[3] = byte(expected_page_id << 24)

		ok := pageImpl.SetData(data_slice)
		if ok != nil {
			t.Errorf("pageImpl.SetData failed")
		}

		pageImpl.ResetMemory()

		helpers.NotEquals(t, expected_page_id, pageImpl.GetPageId())
		helpers.NotEqualSlices(t, data_slice, pageImpl.GetData())

	})

	t.Run("SetPageId", func(t *testing.T) {
		pageImpl := &page.PageImpl{}
		data_slice := make([]byte, page.PAGE_SIZE)

		expected_page_id := uint32(12)
		data_slice[0] = byte(expected_page_id)
		data_slice[1] = byte(expected_page_id << 8)
		data_slice[2] = byte(expected_page_id << 16)
		data_slice[3] = byte(expected_page_id << 24)

		pageImpl.SetPageId(expected_page_id)

		helpers.Equals(t, expected_page_id, pageImpl.GetPageId())
		helpers.EqualSlices(t, data_slice, pageImpl.GetData())

	})

	t.Run("InvalidPage", func(t *testing.T) {
		invalidPage := page.InvalidPage()

		helpers.Equals(t, uint32(page.INVALID_PAGE_ID), invalidPage.GetPageId())

	})
}
