package type__test

import (
	//	"fmt"
	"github.com/chaitya62/noobdb/type"
	"reflect"
	"testing"
)

func TestTypeFactory(t *testing.T) {

	t.Run("Integer", func(t *testing.T) {

		integer := type_.TypeFactory("INTEGER")
		got := reflect.TypeOf(integer)
		expected := reflect.TypeOf(&type_.Integer{})

		if got != expected {
			t.Errorf("Got %v, Expect: %v", got, expected)
		}
	})

	t.Run("Boolean", func(t *testing.T) {
		boolean := type_.TypeFactory("BOOLEAN")
		got := reflect.TypeOf(boolean)
		expected := reflect.TypeOf(&type_.Boolean{})

		if got != expected {
			t.Errorf("Got %v, Expect: %v", got, expected)
		}
	})

	t.Run("Varchar", func(t *testing.T) {
		varchar := type_.TypeFactory("VARCHAR")
		got := reflect.TypeOf(varchar)
		expected := reflect.TypeOf(&type_.Varchar{})

		if got != expected {
			t.Errorf("Got %v, Expect: %v", got, expected)
		}
	})
}

func TestTypeFromTupleFactory(t *testing.T) {

	t.Run("Tuple Integer", func(t *testing.T) {
		tuple := []byte{245, 0, 0, 0, 0, 0, 0, 0}

		start_pos := uint64(0)
		expected_pos := start_pos + type_.INTEGER_SIZE

		got_pos, got_obj := type_.TypeFromTupleFactory("INTEGER", tuple, start_pos)

		if got_pos != expected_pos {
			t.Errorf("Got next pos %v, Expect next pos: %v", got_pos, expected_pos)
		}

		got_type := reflect.TypeOf(got_obj)
		expected_type := reflect.TypeOf(&type_.Integer{})

		if got_type != expected_type {
			t.Errorf("Got type %v, Expect type: %v", got_type, expected_type)
		}

		got_val := got_obj.GetValue()
		expected_val := int64(245)

		if got_val != expected_val {
			t.Errorf("Got value %v, Expect value: %v", got_val, expected_val)
		}
	})
}
