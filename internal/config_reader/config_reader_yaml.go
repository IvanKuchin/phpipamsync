package config_reader

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type YAMLConfigReader struct {
}

func (config_reader *YAMLConfigReader) GetConfig() (*Config, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	filename := filepath.Join(home_dir, ".phpipamsync", "config")

	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}

	cfg := Config{}
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	// cfg, err := parseConfigFile(config_file)
	// if err != nil {
	// 	return nil, err
	// }

	return &cfg, nil
}

func NewYAMLConfigReader() (ConfigReader, error) {
	return &YAMLConfigReader{}, nil
}
