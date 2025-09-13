package core

import (
	"sync"
)

/*
Requirements:
- A user should be able to set a project
- A user should be able to set their backend provider
- A user should be able to set their environment names
- A user should be able to map environment names to prefixes
- Data should be stored in a yml file and available to the program
- I want to be able to extend this to the cli and within a tui

```yml
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
```
*/

type ConfigPath string

const (
	PathProjects            ConfigPath = "projects"
	PathBackend             ConfigPath = "backend"
	PathName                ConfigPath = "name"
	PathEnvironmentPrefix   ConfigPath = "environments.%s.prefix"
	PathEnvironmentPosition ConfigPath = "environments.%s.position"
	PathEnvironmentName     ConfigPath = "environments.%s.name"
)

type Setter interface {
	SetValue()
}

type Environment struct {
	EnvironmentName string `yaml:"name"`
	DisplayPosition int    `yaml:"position"`
	SearchPrefix    string `yaml:"prefix"`
}

type Project struct {
	ProjectName     string        `yaml:"name"`
	BackendProvider string        `yaml:"backend"`
	Environments    []Environment `yaml:"environments"`
}

type Config struct {
	Mu             sync.RWMutex
	CurrentProject string
	Projects       map[string]Project `yaml:"projects"`
}

func (c *Config) SetValue(value_key ConfigPath, value any) {
	// Lock config
	c.Mu.Lock()
	defer c.Mu.Unlock()

	// Match key
}

func (c *Config) SetProject(project_key string) {}

func (c *Config) ListProjects() []string {
	project_keys := make([]string, 0, len(c.Projects))
	for key := range c.Projects {
		project_keys = append(project_keys, key)
	}

	return project_keys
}
