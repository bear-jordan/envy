package core

import (
	"errors"
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
		description: Example Project
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
	PathDescription         ConfigPath = "description"
	PathEnvironmentPrefix   ConfigPath = "environments.%s.prefix"
	PathEnvironmentPosition ConfigPath = "environments.%s.position"
	PathEnvironmentName     ConfigPath = "environments.%s.name"
)

type Environment struct {
	EnvironmentName string `yaml:"name"`
	DisplayPosition int    `yaml:"position"`
	SearchPrefix    string `yaml:"prefix"`
}

type Project struct {
	ProjectDescription string        `yaml:"description"`
	BackendProvider    string        `yaml:"backend"`
	Environments       []Environment `yaml:"environments"`
}

type Config struct {
	CurrentProject string
	Projects       map[string]Project `yaml:"projects"`
}

func (c *Config) AddProject(id string, project Project) error {
	_, ok := c.Projects[id]
	if ok {
		return errors.New("Unable to add project.")
	}
	c.Projects[id] = project

	return nil
}
