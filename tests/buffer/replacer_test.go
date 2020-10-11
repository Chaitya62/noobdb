package buffer_test

import (
	"github.com/chaitya62/noobdb/buffer"
	"github.com/chaitya62/noobdb/tests/helpers"
	"testing"
)

func TestClockReplacer(t *testing.T) {

	// not a unit test techincally
	t.Run("Check victim detection", func(t *testing.T) {

		t.Run("Detech victim properly", func(t *testing.T) {
			clockReplacer := buffer.NewClockReplacer(10)
			clockReplacer.UnPinPage(5)
			clockReplacer.UnPinPage(1)
			clockReplacer.UnPinPage(6)

			page1, ok := clockReplacer.GetNextVictim()

			if !ok {
				t.Errorf("Should return a frame_id")
			}

			helpers.Equals(t, 1, page1)

		})

		t.Run("When there are no victims", func(t *testing.T) {
			clockReplacer := buffer.NewClockReplacer(10)

			_, ok := clockReplacer.GetNextVictim()

			if ok {
				t.Errorf("Should return no frame_id")
			}

		})

	})

	t.Run("Check Size detection", func(t *testing.T) {
		clockReplacer := buffer.NewClockReplacer(10)
		clockReplacer.UnPinPage(5)
		clockReplacer.UnPinPage(1)
		clockReplacer.UnPinPage(6)

		helpers.Equals(t, uint32(3), clockReplacer.Size())

	})

	t.Run("Check PinPage", func(t *testing.T) {
		clockReplacer := buffer.NewClockReplacer(10)
		clockReplacer.UnPinPage(5)
		clockReplacer.UnPinPage(1)
		clockReplacer.UnPinPage(6)

		clockReplacer.PinPage(1)
		page1, ok := clockReplacer.GetNextVictim()

		if !ok {
			t.Errorf("Should return a frame_id")
		}

		helpers.Equals(t, 5, page1)

	})
}
