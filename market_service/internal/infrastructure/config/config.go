package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"app"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Logger struct {
		Level    string `yaml:"level"`
		Logstash struct {
			Address string `yaml:"address"`
			Enabled bool   `yaml:"enabled"`
		}
	} `yaml:"logger"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	log.Printf("Config loaded from %s", configPath)

	return &config, nil
}
