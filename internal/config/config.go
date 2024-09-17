package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/LiddleChild/space/internal/models"
	"github.com/spf13/cobra"
)

var AppConfig *Config

func init() {
	var err error
	AppConfig, err = Load()
	cobra.CheckErr(err)
}

type ConfigMetadata struct {
	directory string
	filename  string
}

type Config struct {
	metadata ConfigMetadata           `json:"-"`
	Spaces   map[string]*models.Space `json:"spaces"`
}

func Load() (*Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		metadata: ConfigMetadata{
			directory: path.Join(homePath, ".config/space"),
			filename:  "settings.json",
		},
		Spaces: make(map[string]*models.Space),
	}

	err = cfg.readConfigFile(path.Join(cfg.metadata.directory, cfg.metadata.filename))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) Save() error {
	cfg.ensureConfigDirectory(cfg.metadata.directory)

	absoluteConfigPath := path.Join(cfg.metadata.directory, cfg.metadata.filename)
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

func (cfg *Config) ensureConfigDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cfg *Config) readConfigFile(filename string) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return err
	}

	return nil
}
