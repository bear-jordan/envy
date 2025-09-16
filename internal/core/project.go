package core

import "errors"

type Project struct {
	ProjectDescription  string `yaml:"description"`
	BackendProviderType string `yaml:"backend_type"`
	BackendProvider     BackendProvider
	Environments        []Environment `yaml:"environments"`
}

func (p *Project) UpdateBackendProvider(provider BackendProvider) error {
	p.BackendProvider = provider
	return nil
}

func (p *Project) AddEnvironment(environment Environment) error {
	return errors.New("Unable to add environment.")
}

func (p *Project) RemoveEnvironment(environment Environment) error {
	return errors.New("Unable to add environment.")
}
