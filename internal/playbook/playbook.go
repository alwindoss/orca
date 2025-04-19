package playbook

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Task struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Become  bool   `yaml:"become"`
}

type Play struct {
	Name  string `yaml:"name"`
	Hosts string `yaml:"hosts"`
	Tasks []Task `yaml:"tasks"`
}

func LoadPlaybook(filePath string) ([]Play, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var plays []Play
	if err := yaml.Unmarshal(data, &plays); err != nil {
		return nil, err
	}

	return plays, nil
}
