package swalker

import (
	"fmt"
	"os"
	"path/filepath"
)

func SymWalker(conf *SymConf) (results Results, err error) {

	conf.StartPath, err = filepath.Abs(filepath.Clean(conf.StartPath))
	if err != nil {
		return
	}

	loopErr := startWalkLoop(conf)
	if loopErr != nil {
		noise(conf.Noisy, "SymWalker encountered an error: %s", loopErr.Error())
		return
	}

	return newResults(*conf), nil

}

func startWalkLoop(conf *SymConf) (err error) {

	readable, err := isReadable(fullPathUnsafe(conf.StartPath))
	if err != nil || !readable {
		noise(conf.Noisy, "unable to read: %s", conf.StartPath)
		return fmt.Errorf("path is not readable: %s", conf.StartPath)
	}

	switch isEntType(conf.StartPath) {
	case entTypeDir:

		err := dirWalk(conf, conf.StartPath, lineHistory{})
		if err != nil {
			break
		}

	default:

		return fmt.Errorf("startPath must be a directory")

	}

	return nil
}

func dirWalk(conf *SymConf, basePath string, history lineHistory) (err error) {

	history = history.refresh()

	readable, err := isReadable(fullPathUnsafe(basePath))
	if err != nil || !readable {
		err = fmt.Errorf("path is not readable: %s", basePath)
		return
	}

	noise(conf.Noisy, "Reading %s", basePath)

	switch isEntType(basePath) {
	case entTypeDir:

		if history.pathExists(basePath) {
			noise(conf.Noisy, "Path already processed: %s", basePath)
			return fmt.Errorf("path already processed: %s", basePath)
		}
		history = history.add(basePath)

		conf.dirs.add(basePath)

		entries, err := os.ReadDir(basePath)
		if err != nil {
			return err
		}

		for _, entry := range entries {

			workPath := joinUnsafe(basePath, entry.Name())
			processDirEntry(conf, workPath, history)

		}

	case entTypeLink:

		target, err := filepath.EvalSymlinks(basePath)
		if err != nil {
			return err
		}

		switch isEntType(target) {
		case entTypeDir:

			if history.pathExists(target) {
				noise(conf.Noisy, "Path already processed: %s", target)
				return fmt.Errorf("path already processed: %s", target)
			}
			history = history.add(target)

			conf.dirs.add(basePath)

			entries, err := os.ReadDir(target)
			if err != nil {
				return err
			}

			for _, entry := range entries {

				workPath := joinUnsafe(basePath, entry.Name())
				processDirEntry(conf, workPath, history)

			}

		case entTypeFile:

			conf.files.add(basePath)

		case entTypeOther, entTypeErrored:

			conf.others.add(basePath)

		}

	case entTypeFile:

		conf.files.add(basePath)

	case entTypeOther, entTypeErrored:

		conf.others.add(basePath)

	}

	return

}

func processDirEntry(conf *SymConf, path string, history lineHistory) {

	//noise(conf.Noisy, "Processing %s", path)

	history = history.refresh()

	readable, err := isReadable(fullPathUnsafe(path))
	if err != nil || !readable {
		noise(conf.Noisy, "Unable to read: %s", path)
		return
	}

	noise(conf.Noisy, "Processing (%s): %s", isEntType(path).string(), path)

	switch isEntType(path) {
	case entTypeDir:

		err := dirWalk(conf, path, history)
		if err != nil {
			break
		}

	case entTypeLink:

		target, err := filepath.EvalSymlinks(path)
		if err != nil {
			return
		}

		noise(conf.Noisy, "Processing link: %s; target: %s; (leads to: %s)", path, target, isEntType(target).string())
		switch isEntType(target) {
		case entTypeDir:

			err = dirWalk(conf, path, history)
			if err != nil {
				break
			}

		case entTypeLink:

			processDirEntry(conf, path, history)

		case entTypeFile:

			conf.files.add(path)

		}

	case entTypeFile:

		conf.files.add(path)

	case entTypeOther:

		conf.others.add(path)
		noise(conf.Noisy, "Skipping %s", path)

	case entTypeErrored:

		conf.others.add(path)
		noise(conf.Noisy, "unable to determine entry type: %s", path)

	}

	return
}
