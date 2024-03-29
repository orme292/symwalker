package swalker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SymWalker(conf *SymConf) (Res WalkerResults, err error) {
	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return nil, err
	}
	sType := isPathType(conf.StartPath)
	switch sType {
	case symTypeDir:
		fmt.Println("Start WalkLoop")
		nRes, nErr := startWalkLoop(conf, conf.StartPath)
		if nErr != nil {
			return
		}
		Res = append(Res, nRes...)
	default:
		return Res, ErrStartPath
	}
	return
}

func startWalkLoop(conf *SymConf, path string) (Res WalkerResults, err error) {
	readable, err := isReadable(f(path))
	if err != nil {
		return
	}
	if !readable {
		return Res, fmt.Errorf("path is not readable: %s", path)
	}

	sType := isPathType(path)
	switch sType {
	case symTypeDir:
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

	baseInfo, err := os.Lstat(base)
	if err != nil {
		return
	}
	if pathTypeFromInfo(baseInfo) == symTypeLink {
		if symlinkTargetInHistory(base) {
			return
		}
	} else {
		if pathInHistory(base) {
			return
		}
	}
	history = append(history, base)

	if conf.Noisy {
		log.Println("Reading ", base)
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
		switch pathTypeFromInfo(info) {
		case symTypeDir:
			Res = append(Res, WalkerEntry{Path: workPath})
			nRes, err := dirWalk(conf, workPath)
			if err != nil {
				break
			}
			Res = append(Res, nRes...)
		case symTypeLink:
			if resolvesToDir(workPath) && conf.FollowSymlinks {
				if symlinkTargetInHistory(workPath) {
					continue
				}
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
