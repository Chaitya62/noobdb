package type__test

import (
	"github.com/chaitya62/noobdb/tests/helpers"
	"github.com/chaitya62/noobdb/type"
	"testing"
)

func TestInteger(t *testing.T) {
	integer := &type_.Integer{}

	t.Run("Implements Type interface", func(t *testing.T) {
		_, ok := interface{}(integer).(type_.Type)
		if ok != true {
			t.Errorf("Integer does not implement type_.Type")
		}
	})

	t.Run("Serialize", func(t *testing.T) {
		expected := []byte{255, 1, 0, 0, 0, 0, 0, 0}

		integer.SetValue(int64(511))
		got := integer.Serialize()

		helpers.EqualSlices(t, expected, got)

	})

	t.Run("Deserialize", func(t *testing.T) {

		t.Run("correct input", func(t *testing.T) {
			input := []byte{254, 1, 0, 0, 0, 0, 0, 0}
			expected := int64(510)

			ok := integer.Deserialize(input)

			if ok != nil {
				t.Errorf("Got error: %v", ok)
			}

			got := integer.GetValue()

			helpers.Equals(t, expected, got)
		})

		t.Run("correct input", func(t *testing.T) {
			input := []byte{254, 1, 0, 0, 0, 0, 0}
			integer.SetValue(int64(0))
			expected := int64(0)

			ok := integer.Deserialize(input)

			if ok == nil {
				t.Errorf("Should raise error")
			}

			got := integer.GetValue()
			helpers.Equals(t, expected, got)

		})

	})

}
