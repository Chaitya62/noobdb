package buffer_test

import (
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/tests/helpers"
	"testing"
)

func TestPageTable(t *testing.T) {

	pt := &buffer.PageTable{}
	pt.Init()

	t.Run("Test Setter/Getter", func(t *testing.T) {
		pt.InsertOrUpdate(uint32(10), 14)

		if val, ok := pt.Get(uint32(10)); ok {
			helpers.Equals(t, 14, val)
		} else {
			t.Errorf("should return a value for key 10")
		}

	})

	t.Run("When key is not present", func(t *testing.T) {
		val, ok := pt.Get(uint32(1))
		if ok {
			t.Errorf("Key 1 should not exist")
		}

		helpers.Equals(t, -1, val)
	})

}
