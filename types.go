package swalker

import (
	"errors"
)

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
}

type WalkerEntry struct {
	Path string
}

type WalkerResults []WalkerEntry

type History []string

var history History

func pathInHistory(path string) bool {
	count := 0
	for _, entry := range history {
		if entry == path {
			count++
		}
		if count > 1 {
			return true
		}
	}
	return false
}

var (
	ErrStartPath = errors.New("StartPath should be accessible directory")
)
