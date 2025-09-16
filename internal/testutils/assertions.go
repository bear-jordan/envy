package testutils

import (
	"reflect"
	"testing"
)

func AssertType(t testing.TB, got any, want any) {
	t.Helper()

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertEqual(t testing.TB, got any, want any) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

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
