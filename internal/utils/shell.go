package utils

import (
	"os"
	"slices"
	"syscall"
)

func Shell(envs ...string) error {
	shell := os.Getenv("SHELL")
	environ := slices.Insert(syscall.Environ(), 0, envs...)
	err := syscall.Exec(shell, []string{shell}, environ)
	if err != nil {
		return err
	}

	return nil
}
