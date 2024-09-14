package config

import (
	"fmt"
	"os"

	"github.com/LiddleChild/space/internal/models"
)

func (cfg *Config) CreateSpace(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, space := range cfg.Spaces {
		if space.Name == name {
			return fmt.Errorf("%s already exists", space.Name)
		}
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
