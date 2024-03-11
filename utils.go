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

	_, err = os.Open(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func resolvesToDir(path string) bool {
	workPath, err := filepath.EvalSymlinks(f(path))
	if err != nil {
		return false
	}
	switch isType(workPath) {
	case symTypeDir:
		return true
	case symTypeLink:
		return resolvesToDir(workPath)
	}
	return false
}
