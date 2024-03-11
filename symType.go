package swalker

import (
	"fmt"
	"os"
	"path/filepath"
)

type symType string

var (
	symTypeDir     symType = "dir"
	symTypeLink    symType = "link"
	symTypeFile    symType = "file"
	symTypeOther   symType = "other"
	symTypeErrored symType = "errored"
)

func (st symType) string() string {
	return fmt.Sprintf("%s", st)
}

func isType(path string) symType {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return symTypeErrored
	}

	info, err := os.Lstat(path)
	if err != nil {
		return symTypeErrored
	}

	return isTypeFromInfo(info)
}

func isTypeFromInfo(info os.FileInfo) (st symType) {
	if info.IsDir() {
		return symTypeDir
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		return symTypeLink
	}
	if info.Mode().IsRegular() {
		return symTypeFile
	}
	return symTypeOther
}
