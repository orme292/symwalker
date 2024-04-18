// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

import (
	"os"
	"path/filepath"
)

// entType is a simplified file type system.
type entType string

var (
	entTypeDir     entType = "dir"
	entTypeLink    entType = "link"
	entTypeFile    entType = "file"
	entTypeOther   entType = "other"
	entTypeErrored entType = "errored"
)

// string returns the string representation of the entType.
func (st entType) string() string {
	return string(st)
}

// isEntType determines the entType by retrieving the path's info (with os.Lstat).
// The function returns the result received by calling the entTypeFromInfo function,
// which is passed the info from os.Lstat.
func isEntType(path string) entType {
	path, err := filepath.Abs(path)
	if err != nil {
		return entTypeErrored
	}

	info, err := os.Lstat(path)
	if err != nil {
		return entTypeErrored
	}

	return entTypeFromInfo(info)
}

// entTypeFromInfo determines the entType based on the os.FileInfo provided.
// If the FileInfo represents a directory, the function will return entTypeDir.
// If the FileInfo represents a symbolic link, the function will return entTypeLink.
// If the FileInfo represents a regular file, the function will return entTypeFile.
// Otherwise, it will return entTypeOther.
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
