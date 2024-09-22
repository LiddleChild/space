package config

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func EnsureConfigDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func EnsureStartUpScript() error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	zshrc := path.Join(homePath, ".zshrc")
	bs, err := os.ReadFile(zshrc)
	if err != nil {
		return err
	}

	beginIndex := strings.Index(string(bs), "# space begin")

	if beginIndex >= 0 {
		return nil
	}

	shell := os.Getenv("SHELL")
	cmd := fmt.Sprintf("cat %s >> %s", path.Join(AppConfig.metadata.directory, "scripts", "rc.zsh"), zshrc)
	err = exec.Command(shell, "-c", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}
