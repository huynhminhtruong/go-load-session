package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Cookie struct {
	Name   string `yaml:"name"`
	Value  string `yaml:"value"`
	Domain string `yaml:"domain"`
	Path   string `yaml:"path"`
}

type CookieConfig struct {
	Cookies []Cookie `yaml:"cookies"`
}

func LoadCookieConfig(path string) (*CookieConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg CookieConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
