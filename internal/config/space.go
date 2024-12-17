package config

import (
	"fmt"
	"os"
	"slices"
	"time"

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

	cfg.Spaces[name] = &models.Space{
		Name:       name,
		Path:       pwd,
		LastOpened: time.Now(),
	}

	err = cfg.Save()
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) GetSpaces() []*models.Space {
	spaces := make([]*models.Space, 0, len(cfg.Spaces))
	for _, val := range cfg.Spaces {
		spaces = append(spaces, val)
	}

	slices.SortFunc(spaces, func(a, b *models.Space) int {
		return cfg.Spaces[a.Name].LastOpened.Compare(cfg.Spaces[b.Name].LastOpened) * -1
	})

	return spaces
}

func (cfg *Config) GetSpaceNames() []string {
	names := make([]string, 0, len(cfg.Spaces))
	for _, val := range cfg.GetSpaces() {
		names = append(names, val.Name)
	}

	return names
}

func (cfg *Config) GetSpace(name string) (*models.Space, error) {
	space, ok := cfg.Spaces[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	return space, nil
}

func (cfg *Config) RemoveSpace(name string) error {
	_, ok := cfg.Spaces[name]
	if !ok {
		return fmt.Errorf("%s does not exist", name)
	}

	delete(cfg.Spaces, name)

	err := cfg.Save()
	if err != nil {
		return err
	}

	return nil
}
