package swalker

import (
	"errors"
)

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
}

type WalkerEntry struct {
	Path string
}

type WalkerResults []WalkerEntry

type History []string

var history History

func pathInHistory(path string) bool {
	count := 0
	for i := 0; i < len(history); i++ {
		if path == history[i] {
			count += 1
		}
		if count == 2 {
			return true
		}
	}
	return false
}

var (
	ErrStartPath = errors.New("StartPath should be accessible directory")
)
