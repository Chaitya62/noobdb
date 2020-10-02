package type__test

import (
	"github.com/chaitya62/noobdb/tests/helpers"
	"github.com/chaitya62/noobdb/type"
	"testing"
)

func TestBoolean(t *testing.T) {
	boolean := &type_.Boolean{}

	t.Run("Implements Type interface", func(t *testing.T) {
		_, ok := interface{}(boolean).(type_.Type)
		if ok != true {
			t.Errorf("Integer does not implement type_.Type")
		}
	})

	t.Run("Serialize", func(t *testing.T) {

		t.Run("true", func(t *testing.T) {
			expected := []byte{1}

			boolean.SetValue(true)
			got := boolean.Serialize()

			helpers.EqualSlices(t, expected, got)

		})

		t.Run("false", func(t *testing.T) {
			expected := []byte{0}

			boolean.SetValue(false)
			got := boolean.Serialize()

			helpers.EqualSlices(t, expected, got)

		})
	})

	t.Run("Deserialize", func(t *testing.T) {

		t.Run("any input - true", func(t *testing.T) {
			input := []byte{25}
			expected := true

			ok := boolean.Deserialize(input)

			if ok != nil {
				t.Errorf("Got error: %v", ok)
			}

			got := boolean.GetValue()

			helpers.Equals(t, expected, got)

		})

		t.Run("correct 0 input - false", func(t *testing.T) {
			input := []byte{0}
			expected := false

			ok := boolean.Deserialize(input)

			if ok != nil {
				t.Errorf("Got error: %v", ok)
			}

			got := boolean.GetValue()

			helpers.Equals(t, expected, got)

		})

		t.Run("incorrect input", func(t *testing.T) {
			input := []byte{}
			boolean.SetValue(true)
			expected := false

			ok := boolean.Deserialize(input)

			if ok == nil {
				t.Errorf("Expect error to be raised")
			}

			got := boolean.GetValue()

			helpers.NotEquals(t, expected, got)

		})
	})

}
