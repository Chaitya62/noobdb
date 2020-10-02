package helpers

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, expected, got interface{}) {
	if got != expected {
		t.Errorf("Expected %v, Got %v", expected, got)
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
