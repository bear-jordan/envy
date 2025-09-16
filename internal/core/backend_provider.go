package core

import "errors"

type BackendProvider interface {
	CreateValue() error
	UpdateValue() error
	DeleteValue() error
	TestConnection() error
}

// Factory
func BackendProviderFactory(backendType string) (BackendProvider, error) {
	switch backendType {
	case "mock":
		return &MockBackendProvider{}, nil
	default:
		return nil, errors.New("BackendProvider not found.")
	}
}

// Mock BackendProvider
type MockBackendProvider struct{}

func (t *MockBackendProvider) CreateValue() error {
	return nil
}

func (t *MockBackendProvider) UpdateValue() error {
	return nil
}

func (t *MockBackendProvider) DeleteValue() error {
	return nil
}

func (t *MockBackendProvider) TestConnection() error {
	return nil
}
