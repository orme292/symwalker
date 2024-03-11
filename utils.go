package swalker

import (
	"os"
	"path/filepath"
)

func isReadable(path string) (bool, error) {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return false, err
	}

	_, err = os.Open(path)
	if err != nil {
		return false, err
	}
	return true, nil
}
