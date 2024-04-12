package swalker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func SymWalker(conf *SymConf) (results ResultEntries, err error) {
	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return nil, err
	}

	loopResults, loopErr := startWalkLoop(conf, conf.StartPath)
	if loopErr != nil {
		return
	}

	results = append(results, loopResults...)

	return
}

func startWalkLoop(conf *SymConf, path string) (results ResultEntries, err error) {
	readable, err := isReadable(f(path))
	if err != nil {
		return
	}
	if !readable {
		return results, fmt.Errorf("path is not readable: %s", path)
	}

	pathType := isEntType(path)
	switch pathType {
	case entTypeDir:

		nRes, err := dirWalk(conf, path)
		results = append(results, nRes...)
		if err != nil {
			break
		}

		// deal with pending entries
	}
	return
}

func dirWalk(conf *SymConf, base string) (results ResultEntries, err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}

	conf.history = append(conf.history, base)

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

		switch entTypeFromInfo(info) {
		case entTypeDir:

			results.add(workPath)
			entries, err := dirWalk(conf, workPath)
			if err != nil {
				break
			}
			results.combine(entries)

		case entTypeLink:

			if conf.FollowSymlinks {
				log.Println("PENDING: ", workPath)
				conf.pending.add(workPath)
			}

		}
	}

	return
}
