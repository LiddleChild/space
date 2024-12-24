package utils

import (
	"fmt"
	"os"
	"syscall"
)

func IsInSession() bool {
	return os.Getenv("TERM_PROGRAM") == "tmux"
}

func NewSession(path string) error {
	shell := os.Getenv("SHELL")
	err := syscall.Exec(shell, []string{shell, "-c", fmt.Sprintf("(tmux new -c %s)", path)}, syscall.Environ())
	if err != nil {
		return err
	}

	return nil
}
