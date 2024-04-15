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
	if err != nil || !readable {
		return fmt.Errorf("path is not readable: %s", conf.StartPath)
	}

	startPathEntType := isEntType(conf.StartPath)
	switch startPathEntType {
	case entTypeDir:

		err := dirWalk(conf, conf.StartPath)
		if err != nil {
			break
		}

		_ = startPendingWalkLoop(conf)

	}
	return
}

func dirWalk(conf *SymConf, base string) (err error) {
	readable, err := isReadable(f(base))
	if err != nil || !readable {
		err = fmt.Errorf("path is not readable: %s", base)
		return
	}

	// if the base path already exists in the history, then exit.
	// otherwise, add it to the history
	if conf.history.pathAddedWithExistsCheck(base) == false {
		noise(conf.Noisy, "Loop found: %s", base)
	}

	noise(conf.Noisy, "Reading %s", base)

	// read the directory
	dirEntries, err := os.ReadDir(base)
	if err != nil {
		return
	}

	for _, entry := range dirEntries {
		workPath := j(base, entry.Name())
		err = processDirEntry(conf, workPath)
		if err != nil {
			noise(conf.Noisy, "Error processing: %s (%s)", workPath, err.Error())
			continue
		}
	}

	return
}

func processDirEntry(conf *SymConf, path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}

	switch entTypeFromInfo(info) {
	case entTypeDir:

		conf.results.add(path)
		err := dirWalk(conf, path)
		if err != nil {
			break
		}

	case entTypeLink:

		if conf.FollowSymlinks {
			if resolvesToDir(path) {
				noise(conf.Noisy, "Link pending: %s (>>dir)", path)
				conf.pending.add(path)
			} else {
				noise(conf.Noisy, "Link skipped: %s (!>>dir)", path)
			}
		}
	}

	return nil
}

func startPendingWalkLoop(conf *SymConf) (err error) {
	endlessLoopProtection := conf.maxLoops
	for {
		for index := range conf.pending {
			if conf.pending[index].Marked == false {
				err = pendingWalk(conf, index)
			}
		}
		if !conf.pending.unmarkedEntriesExist() || endlessLoopProtection <= 0 {
			if endlessLoopProtection <= 0 {
				noise(conf.Noisy, "Loop protection breached: %d", conf.maxLoops)
			}
			break
		}
		endlessLoopProtection--
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
