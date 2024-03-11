package swalker

import (
	"errors"
	"fmt"
	"path/filepath"
)

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
}

type WalkerEntry struct {
	Path string
}

type WalkerResults []*WalkerEntry

func SymWalker(conf *SymConf) (wRes WalkerResults, err error) {
	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return nil, err
	}

	sType := isType(conf.StartPath)
	switch sType {
	case symTypeDir:
		wRes = append(wRes, &WalkerEntry{Path: conf.StartPath})
	default:
		return nil, errors.New("StartPath should be accessible directory")
	}
	return
}

func walk(conf *SymConf, path string, referrer string, wRes WalkerResults) (err error) {
	readable, err := isReadable(path)
	if err != nil {
		return err
	}
	if !readable {
		return fmt.Errorf("path is not readable: %s", path)
	}

	sType := isType(path)
	switch sType {
	case symTypeDir:

	}
}
