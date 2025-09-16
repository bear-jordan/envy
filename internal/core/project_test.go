package core

import (
	"log"
	"testing"

	"github.com/bear-jordan/envy/internal/testutils"
	"github.com/goccy/go-yaml"
)

func projectYml() string {
	return `
%YAML 1.2
---
description: Example Project 2
backend_type: aws_ssm
environments:
    - name: prod
      position: 1
      prefix: /prod/envy
`

}

func projectFixture() *Project {
	var project Project
	yml := projectYml()
	if err := yaml.Unmarshal([]byte(yml), &project); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	return &project
}

// func TestUpdateDescription(t *testing.T) {
// 	t.Run("update with a valid description", func(t *testing.T) {
// 		project := projectFixture()
// 		want := "foo"
// 		err := project.UpdateDescription(want)
// 		testutils.AssertNoError(t, err)
// 		got := project.ProjectDescription
// 		testutils.AssertEqual(t, want, got)
// 	})
// }

func TestUpdateBackendProvider(t *testing.T) {
	t.Run("update with a valid provider", func(t *testing.T) {
		want := &MockBackendProvider{}
		project := projectFixture()
		provider, err := BackendProviderFactory("mock")
		testutils.AssertNoError(t, err)

		err = project.UpdateBackendProvider(provider)
		testutils.AssertNoError(t, err)

		got := project.BackendProvider
		testutils.AssertType(t, got, want)
	})
}

func TestAddEnvironment(t *testing.T) {}

func TestRemoveEnvironment(t *testing.T) {}
