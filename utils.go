package swalker

import (
	"fmt"
	"os"
	"path/filepath"
)

func s(f string, v ...interface{}) string {
	return fmt.Sprintf(f, v...)
}

func j(start string, end string) string {
	return filepath.Join(start, end)
}

func f(path string) string {
	path, _ = filepath.Abs(filepath.Clean(path))
	return path
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

// resolvesToDir checks if the given path resolves to a directory.
// It resolves symbolic links and recursively checks if the resolved path is a directory.
// If the resolved path is a directory, it returns true.
// If the resolved path is a symbolic link to a directory, it recursively calls itself with the resolved path.
// Otherwise, it returns false.
func resolvesToDir(path string) bool {
	workPath, err := filepath.EvalSymlinks(f(path))
	if err != nil {
		return false
	}

	switch isPathType(workPath) {
	case symTypeDir:
		return true
	case symTypeLink:
		return resolvesToDir(workPath)
	}
	return false
}
