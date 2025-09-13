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
