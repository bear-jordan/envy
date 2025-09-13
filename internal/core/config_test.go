package core

import (
	"log"
	"testing"

	"github.com/bear-jordan/envy/internal/testutils"
	"github.com/goccy/go-yaml"
)

func ymlFixture() string {
	return `
%YAML 1.2
---
projects:
    example_project:
        name: Example Project
        backend: aws_ssm
        environments:
            - name: prod
              position: 1
              prefix: /prod/envy
            - name: stg
              position: 2
              prefix: /stg/envy
            - name: dev
              position: 3
              prefix: /dev/envy
`

}

func configFixture() *Config {
	var config Config
	yml := ymlFixture()
	if err := yaml.Unmarshal([]byte(yml), &config); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	return &config
}

func TestConfig(t *testing.T) {
	t.Run("test get projects", func(t *testing.T) {
		config := configFixture()
		got := config.ListProjects()
		want := []string{"example_project"}

		testutils.AssertDeepEqual(t, got, want)
	})

	pathTestConfigs := []struct {
		message string
		name    ConfigPath
		want    string
	}{
		{"", PathProjects, "bar"},
	}

	for _, tc := range pathTestConfigs {
		t.Run(tc.message, func(t *testing.T) {
			// config.SetValue(tc.name, tc.want)
			// testutils.AssertDeepEqual(t, got, tc.want)
		})
	}
}
