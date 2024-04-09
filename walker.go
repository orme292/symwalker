package swalker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SymWalker(conf *SymConf) (Results WalkerResults, err error) {
	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return nil, err
	}

	loopResults, loopErr := startWalkLoop(conf, conf.StartPath)
	if loopErr != nil {
		return
	}

	Results = append(Results, loopResults...)

	return
}

func startWalkLoop(conf *SymConf, path string) (Results WalkerResults, err error) {
	readable, err := isReadable(f(path))
	if err != nil {
		return
	}
	if !readable {
		return Results, fmt.Errorf("path is not readable: %s", path)
	}

	pathType := isPathType(path)
	switch pathType {
	case entTypeDir:

		nRes, err := dirWalk(conf, path)
		Results = append(Results, nRes...)
		if err != nil {
			break
		}

		// deal with pending entries
	}
	return
}

func dirWalk(conf *SymConf, base string) (Results WalkerResults, err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}

	history = append(history, base)

	if conf.Noisy {
		log.Println("Reading ", base)
	}

	dirEnts, err := os.ReadDir(base)
	if err != nil {
		return
	}

	for _, ent := range dirEnts {
		info, err := os.Lstat(j(base, ent.Name()))
		if err != nil {
			continue
		}

		workPath := j(base, ent.Name())

		switch pathTypeFromInfo(info) {
		case entTypeDir:

			Results = append(Results, WalkerEntry{Path: workPath})
			nRes, err := dirWalk(conf, workPath)
			if err != nil {
				break
			}
			Results = append(Results, nRes...)

		case entTypeLink:

			if conf.FollowSymlinks {
				pending = append(pending, WalkerEntry{Path: workPath})
			}

		}
	}

	return
}
