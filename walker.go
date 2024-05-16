// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// SymWalker accepts a configuration object. It calls the function that
// starts the main directory walking loop. It returns a Results object.
func SymWalker(conf *SymConf) (results Results, err error) {

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

	conf.StartPath = fullPathSafe(conf.StartPath)

	ent := isEntType(conf.StartPath)

	noise(conf.Noisy, "Path %q is %q", conf.StartPath, ent.string())

	switch ent {
	case entTypeDir:

		err = dirWalk(conf, conf.StartPath, lineHistory{})

	case entTypeLink:

		target, err := filepath.EvalSymlinks(conf.StartPath)
		if err != nil {
			noise(conf.Noisy, "StartPath could not be resolved: %s", conf.StartPath)
			return errors.New("StartPath could not be resolved")
		}

		if isEntType(target) == entTypeDir {
			err = dirWalk(conf, conf.StartPath, lineHistory{})
		} else {
			noise(conf.Noisy, "StartPath must lead to directory: %s", conf.StartPath)
			return errors.New("StartPath must lead to directory")
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

	noise(conf.Noisy, "Reading %s", basePath)

	switch basePathEnt {
	case entTypeDir:

		if history.exceedsDepth(conf.Depth) {
			break
		}

		if history.pathExists(basePath) {
			noise(conf.Noisy, "Skipping directory because it was already walked: %s", basePath)
			return fmt.Errorf("path exists in line history: %s", basePath)
		}
		history = history.add(basePath)

		conf.dirs.add(basePath, conf.FileData)

		entries, err := os.ReadDir(basePath)
		if err != nil {
			return err
		}

		for _, entry := range entries {

			entryPath := joinPaths(basePath, entry.Name())
			processDirEntry(conf, entryPath, history)

		}

	case entTypeLink:

		if !conf.FollowSymlinks {
			noise(conf.Noisy, "Skipping symlink: %s", basePath)
			break
		}

		target, err := filepath.EvalSymlinks(basePath)
		if err != nil {
			noise(conf.Noisy, "Could not evaluate symlink: %s", basePath)
			return err
		}

		switch isEntType(target) {
		case entTypeDir:

			if history.exceedsDepth(conf.Depth) {
				break
			}

			if history.pathExists(target) {
				noise(conf.Noisy, "Skipping symlinked path already processed: %q=>%q", basePath, target)
				return fmt.Errorf("symlinked path already processed: %q=>%q", basePath, target)
			}
			history = history.add(target)

			conf.dirs.add(basePath, conf.FileData)

			entries, err := os.ReadDir(target)
			if err != nil {
				return err
			}

			for _, entry := range entries {

				entryPath := joinPaths(basePath, entry.Name())
				processDirEntry(conf, entryPath, history)

			}

		case entTypeLink:

			processDirEntry(conf, basePath, history)

		case entTypeFile:

			if !conf.WithoutFiles {
				conf.files.add(basePath, conf.FileData)
			}

		case entTypeOther, entTypeErrored:

			if !conf.WithoutFiles {
				conf.others.add(basePath, conf.FileData)
			}

		}

	case entTypeFile:

		if !conf.WithoutFiles {
			conf.files.add(basePath, conf.FileData)
		}

	case entTypeOther, entTypeErrored:

		if !conf.WithoutFiles {
			conf.others.add(basePath, conf.FileData)
		}

	}

	return err

}

// processDirEntry handles each directory entry from the dirWalk function.
// The purpose of processDirEntry is to determine how to handle each entry.
// If the provided path's entType is entTypeDir, it is passed to dirWalk.
// If it is entTypeLink, the target's target is evaluated and passed to processDirEntry.
// If the provided path (or the path's target) entType is: entTypeFile, entTypeOther,
// or entTypeErrored then the path is added to the results.
func processDirEntry(conf *SymConf, path string, history lineHistory) {

	noise(conf.Noisy, "Processing: %s", path)

	history = history.branch()

	ent := isEntType(path)

	switch ent {
	case entTypeDir:

		err := dirWalk(conf, path, history)
		if err != nil {
			break
		}

	case entTypeLink:

		if !conf.FollowSymlinks {
			noise(conf.Noisy, "Skipping symlink: %s", path)
			break
		}

		target, err := filepath.EvalSymlinks(path)
		if err != nil {
			noise(conf.Noisy, "Could not evaluate symlink: %s", path)
			return
		}

		noise(conf.Noisy, "Link:  %q links to => %q", path, target)

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

			linkTargetBase := joinPaths(path, linkTarget)
			processDirEntry(conf, linkTargetBase, history)

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
