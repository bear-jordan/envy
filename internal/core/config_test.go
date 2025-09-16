package core

import (
	"log"
	"testing"

	"github.com/bear-jordan/envy/internal/testutils"
)

func configFixtureYml() string {
	return `
%YAML 1.2
---
projects:
    project_1:
        description: Example Project
        backend_type: mock
        environments:
            - name: prod
              position: 1
              prefix: /prod/envy
            - name: stg
              position: 2
              prefix: /stg/envy
`
}

func configFixture() *Config {
	specification := configFixtureYml()
	config, err := UnmarshalYml(specification)
	if err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	return config
}

func TestAddProject(t *testing.T) {
	t.Run("add a new project", func(t *testing.T) {
		config := configFixture()
		project := projectFixture()

		err := config.AddProject("project_2", *project)
		testutils.AssertNoError(t, err)
	})
	t.Run("add a duplicate project fails", func(t *testing.T) {
		config := configFixture()
		project := projectFixture()

		_ = config.AddProject("project_2", *project)
		err := config.AddProject("project_2", *project)
		testutils.AssertError(t, err)
	})

}

func TestRemoveProject(t *testing.T) {
	t.Run("test removing a valid project", func(t *testing.T) {
		config := configFixture()

		err := config.RemoveProject("project_1")
		testutils.AssertNoError(t, err)
	})
	t.Run("test removing an invalid project", func(t *testing.T) {
		config := configFixture()

		err := config.RemoveProject("should_fail")
		testutils.AssertError(t, err)
	})
}

func TestSetCurrentProject(t *testing.T) {
	t.Run("test set valid project", func(t *testing.T) {
		// Setup config
		config := configFixture()
		project := projectFixture()
		err := config.AddProject("project_2", *project)
		testutils.AssertNoError(t, err)

		// Set project
		err = config.SetCurrentProject("project_2")
		testutils.AssertNoError(t, err)
	})
	t.Run("test set invalid project", func(t *testing.T) {
		// Setup config
		config := configFixture()
		project := projectFixture()
		err := config.AddProject("project_2", *project)
		testutils.AssertNoError(t, err)

		// Set project
		err = config.SetCurrentProject("should_fail")
		testutils.AssertError(t, err)

	})
}

func TestUnmarshalSpecification(t *testing.T) {
	t.Run("test unmarshalling a valid string", func(t *testing.T) {})
	t.Run("test unmarshalling an invalid string", func(t *testing.T) {})
	t.Run("test unmarshalling a valid path", func(t *testing.T) {})
	t.Run("test unmarshalling an invalid path", func(t *testing.T) {})
}
