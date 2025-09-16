package core

import (
	"testing"

	"github.com/bear-jordan/envy/internal/testutils"
)

func TestBackendProviderFactory(t *testing.T) {
	t.Run("test generation of a valid provider", func(t *testing.T) {
		want := &MockBackendProvider{}
		got, err := BackendProviderFactory("mock")
		testutils.AssertNoError(t, err)
		testutils.AssertType(t, got, want)
	})
}
