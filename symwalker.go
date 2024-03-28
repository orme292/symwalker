package swalker

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrPathNotFound = errors.New("Path not found in Pending")
)

type Pending []WalkerEntry

func (p Pending) Mark(path string) (err error) {
	for i := range p {
		if p[i].Path == path {
			p[i].Marked = true
			return
		}
	}
	return ErrPathNotFound
}

func symWalk(conf *SymConf, base string) (Res WalkerResults, err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
	}

	if conf.Noisy {
		log.Println("Following ", base)
	}

	switch isPathType(base) {
	case symTypeLink:
		actual, err := filepath.EvalSymlinks(base)
		if err != nil {
			return
		}
		dirents, err := os.ReadDir(actual)
		if err != nil {
			return
		}
		for _, ent := range dirents {
			switch isPathType(j(base, ent.Name())) {
			case symTypeLink:
				p = append(p, j(base, ent.Name()))
			case symTypeDir:
				// how to deal with results?
			}
		}
	}
}
