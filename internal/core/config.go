package core

import (
	"errors"
	"log"

	"github.com/goccy/go-yaml"
)

type Config struct {
	CurrentProject string
	Projects       map[string]Project `yaml:"projects"`
}

func UnmarshalYml(specification string) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(specification), &config)
	if err != nil {
		log.Fatalf("Cannot unmarshal specification. %v", err)
	}
	for _, project := range config.Projects {
		project.BackendProvider, err = BackendProviderFactory(project.BackendProviderType)
		if err != nil {
			return nil, err
		}
	}
	return &config, nil
}

// Add a project to a config
func (c *Config) AddProject(id string, project Project) error {
	_, ok := c.Projects[id]
	if ok {
		return errors.New("Unable to add project.")
	}
	c.Projects[id] = project

	return nil
}

// Remove a project from the config
func (c *Config) RemoveProject(id string) error {
	_, ok := c.Projects[id]
	if ok {
		delete(c.Projects, id)
		return nil
	}

	return errors.New("ID does not exist.")
}

// Set the current project
func (c *Config) SetCurrentProject(id string) error {
	_, ok := c.Projects[id]
	if ok {
		c.CurrentProject = id
		return nil
	}

	return errors.New("ID does not exist.")
}
