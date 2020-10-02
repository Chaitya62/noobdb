package type__test

import (
	"github.com/chaitya62/noobdb/tests/helpers"
	"github.com/chaitya62/noobdb/type"
	"testing"
)

func TestVarchar(t *testing.T) {
	varchar := &type_.Varchar{}

	t.Run("Implements Type interface", func(t *testing.T) {
		_, ok := interface{}(varchar).(type_.Type)
		if ok != true {
			t.Errorf("Varchar does not implement type_.Type")
		}
	})

	t.Run("Serialize", func(t *testing.T) {
		expected := []byte{83, 0, 0, 0, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 90, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 40, 41, 45, 44, 34, 39, 33, 64, 35, 36, 37, 94, 38, 42, 91, 93, 123, 125, 92, 124, 96}

		varchar.SetValue("ABCDEFGHIJKLMNOPQRSTUVWZYZabcdefghijklmnopqrstuvwxyz1234567890()-,\"'!@#$%^&*[]{}\\|`")
		got := varchar.Serialize()

		helpers.EqualSlices(t, expected, got)

	})

	t.Run("Deserialize", func(t *testing.T) {
		input := []byte{4, 0, 0, 0, 96, 65, 66, 67}
		expected := "`ABC"

		varchar.Deserialize(input)

		got := varchar.GetValue()

		if got != expected {
			t.Errorf("Got %v, Expect: %v", got, expected)
		}

	})

}
