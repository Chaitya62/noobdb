package helpers

import (
	"reflect"
	"testing"
)

func EqualTypes(t *testing.T, expected, got interface{}) {
	expected_type := reflect.TypeOf(expected)
	got_type := reflect.TypeOf(got)
	if expected_type != got_type {
		t.Errorf("Excepted type: %v, Got type %v", expected_type, got_type)
	}
}

func Equals(t *testing.T, expected, got interface{}) {
	if got != expected {
		t.Errorf("Expected %v, %T  Got %v %T", expected, expected, got, got)
	}
}

func NotEquals(t *testing.T, not_expected, got interface{}) {
	if got == not_expected {
		t.Errorf("Expected %v and %v to be different", not_expected, got)
	}
}

func EqualSlices(t *testing.T, expected, got interface{}) {
	ok := reflect.DeepEqual(expected, got)

	if !ok {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func NotEqualSlices(t *testing.T, expected, got interface{}) {
	ok := reflect.DeepEqual(expected, got)

	if ok {
		t.Errorf("Expected %v not to be equal to %v", got, expected)
	}
}
