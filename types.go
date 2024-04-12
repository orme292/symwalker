package swalker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
	pending        PendingEntries
	history        HistoryEntries
}

var (
	ErrStartPath = errors.New("StartPath should be accessible directory")
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

func isEntType(path string) entType {
	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return entTypeErrored
	}

	info, err := os.Lstat(path)
	if err != nil {
		return entTypeErrored
	}

	return entTypeFromInfo(info)
}

func entTypeFromInfo(info os.FileInfo) (st entType) {
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
