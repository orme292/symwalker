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

	loopErr := startWalkLoop(conf)
	if loopErr != nil {
		return
	}

	return conf.results, nil
}

func startWalkLoop(conf *SymConf) (err error) {
	readable, err := isReadable(f(conf.StartPath))
	if err != nil {
		return
	}
	if !readable {
		return fmt.Errorf("path is not readable: %s", conf.StartPath)
	}

	pathType := isEntType(conf.StartPath)
	switch pathType {
	case entTypeDir:

		err := dirWalk(conf, conf.StartPath)
		if err != nil {
			break
		}

		for index := range conf.pending {
			if conf.pending[index].Marked == false {
				fmt.Println("pending: ", conf.pending[index].Path)
			} else {
				fmt.Println("marked pending entry skipped")
			}
		}

		for index := range conf.history {
			fmt.Println("History Entry: ", *conf.history[index])
		}

		_ = startPendingWalk(conf)
		// deal with pending entries

		for index := range conf.history {
			fmt.Println("Post Pending History Entry: ", *conf.history[index])
		}
	}
	return
}

func dirWalk(conf *SymConf, base string) (err error) {
	readable, err := isReadable(f(base))
	if err != nil {
		return
	}
	if !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}

	// if the base path already exists in the history, then exit.
	// otherwise, add it to the history
	if conf.history.pathExists(base) {
		log.Println("LOOPED: ", base)
		return
	}
	log.Println("History Add (dirWalk): ", base)
	conf.history = append(conf.history, &base)

	// if noisy, print log message
	if conf.Noisy {
		log.Println("Reading ", base)
	}

	// read the directory
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

			conf.results.add(workPath)
			err := dirWalk(conf, workPath)
			if err != nil {
				break
			}

		case entTypeLink:

			if conf.FollowSymlinks {
				if resolvesToDir(workPath) {
					log.Println("PENDING: ", workPath)
					conf.pending.add(workPath)
				} else {
					log.Println("pending skip: ", workPath)
				}
			}

		}

	}

	return
}

func startPendingWalk(conf *SymConf) (err error) {
	for {
		for index := range conf.pending {
			if conf.pending[index].Marked == false {
				err = pendingWalk(conf, index)
			}
		}
		if !conf.pending.unmarkedEntriesExist() {
			break
		}
	}

	return nil
}

func pendingWalk(conf *SymConf, index int) (err error) {
	basePath := conf.pending[index].Path
	conf.pending.markByIndex(index)
	log.Println("evaluate pending: ", basePath)

	readable, err := isReadable(basePath)
	if err != nil || !readable {
		return
	}

	if conf.history.pathExists(basePath) {
		log.Println("LOOPED: ", basePath)
		return
	}
	conf.history.add(basePath)

	info, err := os.Lstat(basePath)
	if err != nil {
		return
	}

	var workPath string
	switch entTypeFromInfo(info) {
	case entTypeDir:

		log.Println("basepath is dir: ", basePath)
		workPath = basePath

	case entTypeLink:

		log.Println("basepath is link: ", basePath)
		targetPath, err := filepath.EvalSymlinks(basePath)
		if err != nil {
			return err
		}
		if conf.history.pathExists(targetPath) {
			log.Println("LOOP: ", basePath, " >> ", targetPath)
			return err
		}
		workPath = targetPath

	default:
		return
	}

	if !conf.results.pathExists(basePath) {
		conf.results.add(basePath)
	}

	dirEnts, err := os.ReadDir(workPath)
	if err != nil {
		return
	}

	for _, ent := range dirEnts {
		entInfo, err := os.Lstat(j(basePath, ent.Name()))
		if err != nil {
			continue
		}

		entWorkPath := j(basePath, ent.Name())

		entWorkPathEntType := entTypeFromInfo(entInfo)
		if (entWorkPathEntType == entTypeDir) || (entWorkPathEntType == entTypeLink) {

			log.Println("Add Entry & Pending: ", entWorkPath)
			conf.results.add(entWorkPath)
			conf.pending.add(entWorkPath)

		}
	}

	return
}
