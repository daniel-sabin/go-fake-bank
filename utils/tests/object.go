package test

import (
	"reflect"
	"testing"
)

func AssertEquals[object any](t *testing.T, got, want object) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("objects are not equal")
	}
}
