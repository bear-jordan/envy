package testutils

import (
	"reflect"
	"testing"
)

func AssertDeepEqual(t testing.TB, got any, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("An error occurred: %s", err)
	}
}

func AssertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("No error occurred")
	}
}
