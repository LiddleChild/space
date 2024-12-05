package utils

import (
	"os"
	"syscall"
)

func Goto(path string) error {
	err := syscall.Chdir(path)
	if err != nil {
		return err
	}

	shell := os.Getenv("SHELL")
	err = syscall.Exec(shell, []string{shell}, syscall.Environ())
	if err != nil {
		return err
	}

	return nil
}
