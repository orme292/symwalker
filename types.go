package swalker

import (
	"errors"
	"path/filepath"
)

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
}

type WalkerEntry struct {
	Path   string
	Marked bool
}

type WalkerResults []WalkerEntry

type History []string
type Pending []WalkerEntry

var history History
var pending Pending

func pathInHistory(path string) bool {
	count := 0
	for _, entry := range history {
		if entry == path {
			count++
		}
		if count >= 1 {
			return true
		}
	}
	return false
}

func symlinkTargetInHistory(path string) bool {
	linkPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}
	return pathInHistory(linkPath)
}

var (
	ErrStartPath = errors.New("StartPath should be accessible directory")
)
