// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

import (
	"fmt"
	"os"
	"path/filepath"
)

// SymWalker accepts a configuration object. It calls the function that
// starts the main directory walking loop. It returns a Results object.
func SymWalker(conf *SymConf) (results Results, err error) {

	conf.StartPath, err = filepath.Abs(conf.StartPath)
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

// startWalkLoop is the where the main walk loop begins. It uses
// the configuration object's StartPath value to start.
// The function calls isReadable, which attempts to open the StartPath.
// On success, isEntType returns the entType of the StartPath.
// If the entType is entTypeDir, then the StartPath is a directory, and
// is walked. Otherwise, the function returns. StartPath CANNOT be
// a file or a symlink.
func startWalkLoop(conf *SymConf) (err error) {

	ent := isEntType(conf.StartPath)

	readable, err := isReadable(fullPathSafe(conf.StartPath), ent)
	if err != nil || !readable {
		noise(conf.Noisy, "not readable: %s", conf.StartPath)
		return fmt.Errorf("StartPath is not readable: %s", conf.StartPath)
	}

	switch ent {
	case entTypeDir:

		err := dirWalk(conf, conf.StartPath, lineHistory{})
		if err != nil {
			break
		}

	default:

		noise(conf.Noisy, "StartPath must be a directory: %s", conf.StartPath)
		return fmt.Errorf("startPath must be a directory")

	}

	return nil

}

// dirWalk IS the primary loop. It accepts a pointer to the configuration object,
// a basePath, and history, which is a lineHistory object. When first calling dirWalk, lineHistory will
// be a new, empty, lineHistory object.
// First, the lineHistory slice is refreshed, and "dereferenced".
// The basePath is checked for readability. After which, basePath's entType is checked.
// If the entType is entTypeDir, then the directory is opened and each directory entry is passed
// to processDirEntry.
// -- if the entType is entTypeFile, entTypeOther, entTypeErrored, it is added to results.
// -- If the entType is entTypeLink, then the target path is evaluated for it's own entType.
// -- If the basePath's target is entTypeDir, it is opened and each directory entry is passed to
// processDirEntry.
// If the basePath's target is entTypeLink, it is passed to processDirEntry.
// If the basePath's target is entTypeFile, entTypeOther, entTypeErrored, it is added to results.
func dirWalk(conf *SymConf, basePath string, history lineHistory) (err error) {

	noise(conf.Noisy, "dirWalk: %s", basePath)

	history = history.branch()

	basePathEnt := isEntType(basePath)

	readable, err := isReadable(fullPathSafe(basePath), basePathEnt)
	if err != nil || !readable {
		err = fmt.Errorf("not readable: %s", basePath)
		return
	}

	noise(conf.Noisy, "Reading %s", basePath)

	switch basePathEnt {
	case entTypeDir:

		if history.exceedsDepth(conf.Depth) {
			break
		}

		if history.pathExists(basePath) {
			noise(conf.Noisy, "Path already processed: %s", basePath)
			return fmt.Errorf("path already processed: %s", basePath)
		}
		history = history.add(basePath)

		conf.dirs.add(basePath, conf.FileData)

		entries, err := os.ReadDir(basePath)
		if err != nil {
			return err
		}

		for _, entry := range entries {

			workPath := joinPaths(basePath, entry.Name())
			processDirEntry(conf, workPath, history)

		}

	case entTypeLink:

		if !conf.FollowSymlinks {
			noise(conf.Noisy, "Not evaluating symlink: %s", basePath)
			break
		}
		target, err := filepath.EvalSymlinks(basePath)
		if err != nil {
			return err
		}

		switch isEntType(target) {
		case entTypeDir:

			if history.exceedsDepth(conf.Depth) {
				break
			}

			if history.pathExists(target) {
				noise(conf.Noisy, "Path already processed: %s", target)
				return fmt.Errorf("path already processed: %s", target)
			}
			history = history.add(target)

			conf.dirs.add(basePath, conf.FileData)

			entries, err := os.ReadDir(target)
			if err != nil {
				return err
			}

			for _, entry := range entries {

				workPath := joinPaths(basePath, entry.Name())
				processDirEntry(conf, workPath, history)

			}

		case entTypeLink:

			processDirEntry(conf, basePath, history)

		case entTypeFile:

			if conf.WithoutFiles {
				break
			}

			conf.files.add(basePath, conf.FileData)

		case entTypeOther, entTypeErrored:

			if conf.WithoutFiles {
				break
			}

			conf.others.add(basePath, conf.FileData)

		}

	case entTypeFile:

		if conf.WithoutFiles {
			break
		}

		conf.files.add(basePath, conf.FileData)

	case entTypeOther, entTypeErrored:

		if conf.WithoutFiles {
			break
		}

		conf.others.add(basePath, conf.FileData)

	}

	return

}

// processDirEntry handles each directory entry from the dirWalk function.
// The purpose of processDirEntry is to determine how to handle each entry.
// If the provided path's entType is entTypeDir, it is passed to dirWalk.
// If it is entTypeLink, the target's target is evaluated and passed to processDirEntry.
// If the provided path (or the path's target) entType is: entTypeFile, entTypeOther,
// or entTypeErrored then the path is added to the results.
func processDirEntry(conf *SymConf, path string, history lineHistory) {

	noise(conf.Noisy, "processDirEntry: %s", path)

	history = history.branch()

	ent := isEntType(path)

	readable, err := isReadable(fullPathSafe(path), ent)
	if err != nil || !readable {
		noise(conf.Noisy, "Not readable: %s", path)
		return
	}

	switch ent {
	case entTypeDir:

		err := dirWalk(conf, path, history)
		if err != nil {
			break
		}

	case entTypeLink:

		if !conf.FollowSymlinks {
			noise(conf.Noisy, "Not evaluating symlink: %s", path)
			break
		}

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

			linkTarget, err := filepath.EvalSymlinks(target)
			if err != nil {
				return
			}

			workPath := joinPaths(path, linkTarget)
			processDirEntry(conf, workPath, history)

		case entTypeFile:

			if !conf.WithoutFiles {
				noise(conf.Noisy, "Process (%s): %s", isEntType(path).string(), path)
				conf.files.add(path, conf.FileData)
			}

		case entTypeOther, entTypeErrored:

			if !conf.WithoutFiles {
				noise(conf.Noisy, "Process (%s): %s", isEntType(path).string(), path)
				conf.others.add(path, conf.FileData)
			}

		}

	case entTypeFile:

		if !conf.WithoutFiles {
			noise(conf.Noisy, "Process (%s): %s", isEntType(path).string(), path)
			conf.files.add(path, conf.FileData)
		}

	case entTypeOther, entTypeErrored:

		if !conf.WithoutFiles {
			noise(conf.Noisy, "Process (%s): %s", isEntType(path).string(), path)
			conf.others.add(path, conf.FileData)
		}

	}

}
