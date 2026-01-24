package config

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	LogLevel string       `yaml:"log_level"`
	Google   GoogleConfig `yaml:"google"`
}

type GoogleConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

var Current Config

func Load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(homeDir, ".sympal", "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return createDefault(configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &Current)
}

func createDefault(path string) error {
	Current = Config{
		LogLevel: "info",
	}

	data, err := yaml.Marshal(&Current)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
