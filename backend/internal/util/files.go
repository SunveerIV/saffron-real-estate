package util

import (
	"errors"
	"os"
	"path/filepath"
)

func FindBackendDirectory() (string, error) {
	if dir := os.Getenv("BACKEND_DIR"); dir != "" {
		return dir, nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find backend directory")
		}
		dir = parent
	}
}
