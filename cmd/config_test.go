package cmd

import "testing"

func testConfig(t *testing.T) {
	got := "foo"
	want := "foo"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
