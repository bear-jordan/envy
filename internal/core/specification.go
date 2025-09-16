package core

type Specification interface {
	GetSpecificationString(path string) string
}

type MockSpecification struct {
	specification string
}

func (m *MockSpecification) GetSpecificationString(path string) string {
	return path
}

type FileSpecification struct{}

func (f *FileSpecification) GetSpecificationString() {}
