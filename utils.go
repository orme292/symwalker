package swalker

import (
	"log"
	"os"
	"path/filepath"
)

func joinUnsafe(start string, end string) string {
	return filepath.Clean(filepath.Join(start, end))
}

func fullPathUnsafe(path string) string {
	path, _ = filepath.Abs(filepath.Clean(path))
	return path
}

func noise(noisy bool, f string, v ...interface{}) {
	if noisy {
		log.Printf(f, v...)
	}
}

func isReadable(path string) (bool, error) {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return false, err
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
