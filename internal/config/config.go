package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/LiddleChild/space/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppConfig *Config

func init() {
	var err error
	AppConfig, err = Load()
	cobra.CheckErr(err)
}

type ConfigMetadata struct {
	directory string
	name      string
	ext       string
}

type Config struct {
	metadata ConfigMetadata          `json:"-"`
	Spaces   map[string]models.Space `json:"spaces"`
}

func Load() (*Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		metadata: ConfigMetadata{
			directory: path.Join(homePath, ".config/space"),
			name:      "settings",
			ext:       "json",
		},
		Spaces: make(map[string]models.Space),
	}

	viper.AddConfigPath(cfg.metadata.directory)
	viper.SetConfigName(cfg.metadata.name)
	viper.SetConfigType(cfg.metadata.ext)

	err = viper.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, err
	} else if err == nil {
		err = viper.Unmarshal(&cfg)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func (cfg *Config) ensureConfigDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cfg *Config) Save() error {
	cfg.ensureConfigDirectory(cfg.metadata.directory)

	absoluteConfigPath := path.Join(cfg.metadata.directory, fmt.Sprintf("%v.%v", cfg.metadata.name, cfg.metadata.ext))
	f, err := os.OpenFile(absoluteConfigPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
