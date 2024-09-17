package utils

import (
	"os"
	"syscall"
)

func Shell(envs ...string) error {
	shell := os.Getenv("SHELL")
	environ := append(syscall.Environ(), envs...)
	err := syscall.Exec(shell, []string{shell}, environ)
	if err != nil {
		return err
	}

	return nil
}
