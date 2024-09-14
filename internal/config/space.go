package config

import (
	"os"

	"github.com/LiddleChild/space/internal/models"
)

func (cfg *Config) CreateSpace(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	space := models.Space{
		Name: name,
		Path: pwd,
	}

	cfg.Spaces = append(cfg.Spaces, space)

	err = cfg.Save()
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) GetSpaces() []models.Space {
	return cfg.Spaces
}
