package swalker

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrPathNotFound = errors.New("Path not found in Pending")
)

func (p Pending) Mark(path string) (err error) {
	for i := range p {
		if p[i].Path == path {
			p[i].Marked = true
			return
		}
	}
	return ErrPathNotFound
}

func walkPending(conf *SymConf, base string) (Results WalkerResults, err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}

	info, err := os.Lstat(base)
	if err != nil {
		return
	}

	history = append(history, base)

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
