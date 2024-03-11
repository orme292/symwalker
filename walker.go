package swalker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func SymWalker(conf *SymConf) (Res WalkerResults, err error) {
	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return nil, err
	}
	sType := isType(conf.StartPath)
	switch sType {
	case symTypeDir:
		fmt.Println("Start WalkLoop")
		nRes, nErr := startWalkLoop(conf, conf.StartPath, "")
		if nErr != nil {
			return
		}
		Res = append(Res, nRes...)
	default:
		return Res, errors.New("StartPath should be accessible directory")
	}
	return
}

func startWalkLoop(conf *SymConf, path string, referrer string) (Res WalkerResults, err error) {
	readable, err := isReadable(f(path))
	if err != nil {
		return
	}
	if !readable {
		return Res, fmt.Errorf("path is not readable: %s", path)
	}

	sType := isType(path)
	switch sType {
	case symTypeDir:
		Res = append(Res, WalkerEntry{Path: path})
		nRes, err := dirWalk(conf, path)
		Res = append(Res, nRes...)
		if err != nil {
			break
		}
	}
	return
}

func dirWalk(conf *SymConf, base string) (Res WalkerResults, err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}
	dirents, err := os.ReadDir(base)
	if err != nil {
		return
	}
	for _, ent := range dirents {
		info, err := os.Lstat(j(base, ent.Name()))
		if err != nil {
			continue
		}
		workPath := j(base, ent.Name())
		switch isTypeFromInfo(info) {
		case symTypeDir:
			Res = append(Res, WalkerEntry{Path: workPath})
			nRes, err := dirWalk(conf, workPath)
			if err != nil {
				break
			}
			Res = append(Res, nRes...)
		case symTypeLink:
			if resolvesToDir(workPath) {
				Res = append(Res, WalkerEntry{Path: workPath})
				nRes, err := dirWalk(conf, workPath)
				if err != nil {
					break
				}
				Res = append(Res, nRes...)
			}
		}

	}
	return
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
		return resolvesToDir(j(path, filepath.Base(workPath)))
	}
	return false
}
