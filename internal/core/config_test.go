package core

import (
	"testing"

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

func TestConfig(t *testing.T) {
	var config Config
	yml := ymlFixture()

	if err := yaml.Unmarshal([]byte(yml), &config); err != nil {
		t.Errorf("unable to parse: %v", err)
	}
}
