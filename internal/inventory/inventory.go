package inventory

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Host struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Port int    `yaml:"port"`
}

type Group struct {
	Hosts []Host `yaml:"hosts"`
}

type Inventory struct {
	Groups map[string]Group `yaml:"groups"`
}

func LoadInventory(filePath string) (*Inventory, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var inv Inventory
	if err := yaml.Unmarshal(data, &inv); err != nil {
		return nil, err
	}

	return &inv, nil
}
