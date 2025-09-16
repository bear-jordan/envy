package core

import (
	"log"
	"testing"

	"github.com/bear-jordan/envy/internal/testutils"
	"github.com/goccy/go-yaml"
)

func configYml() string {
	return `
%YAML 1.2
---
projects:
    example_project_1:
        description: Example Project
        backend: aws_ssm
        environments:
            - name: prod
              position: 1
              prefix: /prod/envy
            - name: stg
              position: 2
              prefix: /stg/envy
`
}

func projectYml() string {
	return `
%YAML 1.2
---
description: Example Project 2
backend: aws_ssm
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

func configFixture() *Config {
	var config Config
	yml := configYml()
	if err := yaml.Unmarshal([]byte(yml), &config); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	return &config
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

// func TestConfig(t *testing.T) {
//  pathTestConfigs := []struct {
//      message string
//      name    ConfigPath
//      want    string
//  }{
//      {"test path projects", PathProjects, ""},
//      {"test path backend", PathBackend, ""},
//      {"test path name", PathName, ""},
//      {"test environment prefix", PathEnvironmentPrefix, ""},
//      {"test environment position", PathEnvironmentPosition, ""},
//      {"test environment name", PathEnvironmentName, ""},
//  }
//  for _, tc := range pathTestConfigs {
//      t.Run(tc.message, func(t *testing.T) {
//          config.SetValue(tc.name, tc.want)
//          testutils.AssertDeepEqual(t, got, tc.want)
//      })
//  }
// }
