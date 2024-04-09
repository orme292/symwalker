package swalker

import (
	"fmt"
	"os"
	"path/filepath"
)

type entType string

var (
	entTypeDir     entType = "dir"
	entTypeLink    entType = "link"
	entTypeFile    entType = "file"
	entTypeOther   entType = "other"
	entTypeErrored entType = "errored"
)

func (st entType) string() string {
	return fmt.Sprintf("%s", st)
}

func isPathType(path string) entType {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return entTypeErrored
	}

	info, err := os.Lstat(path)
	if err != nil {
		return entTypeErrored
	}

	return pathTypeFromInfo(info)
}

func pathTypeFromInfo(info os.FileInfo) (st entType) {
	if info.IsDir() {
		return entTypeDir
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		return entTypeLink
	}
	if info.Mode().IsRegular() {
		return entTypeFile
	}
	return entTypeOther
}
