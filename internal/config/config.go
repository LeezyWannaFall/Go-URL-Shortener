package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Storage struct {
		Type	 string `yaml:"type"`
		Postgres struct {
			Host     string `yaml:"host"`
			Database string `yaml:"database"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Port	 int    `yaml:"port"`
		} `yaml:"postgres"`
	} `yaml:"storage"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if envType := os.Getenv("STORAGE_TYPE"); envType != "" {
        cfg.Storage.Type = envType
    }

    if envUser := os.Getenv("POSTGRES_USER"); envUser != "" {
        cfg.Storage.Postgres.User = envUser
    }

    if envPass := os.Getenv("POSTGRES_PASSWORD"); envPass != "" {
        cfg.Storage.Postgres.Password = envPass
    }

    if envDB := os.Getenv("POSTGRES_DB"); envDB != "" {
        cfg.Storage.Postgres.Database = envDB
    }

	return &cfg, nil
}
