package core

type Environment struct {
	EnvironmentName string `yaml:"name"`
	DisplayPosition int    `yaml:"position"`
	SearchPrefix    string `yaml:"prefix"`
}
