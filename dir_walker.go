package symwalker

import (
	"os"
	"path/filepath"
)

func startLoop(conf *Conf) (res Results, err error) {
	hist := make(history, 0)
	err = dirWalk(conf, conf.StartPath, conf.Depth, res, &hist)
	return
}

func dirWalk(conf *Conf, path string, remDepth int, res Results, hist *history) (err error) {
	// if the remaining depth (remDepth) is 0, do not continue.
	if remDepth == 0 {
		return NewError(OpsMaxDepthReached, s("%q will not be walked", path))
	}
	remDepth = depth(remDepth)

	var skip bool
	// todo: check if dir has already been walked (in history), if it has, skip = true
	if hist.count(path) > 1 {
		skip = true
	} else {
		err = hist.add(path, emptyString)
		// todo: determine how to handle a potential symlink loop. currently, returning an error does nothing
		if err != nil {
			return err // the only possible error is currently ErrWalkPotentialSymlinkLoop
		}
	}

	if skip {
		return nil
	}
	readable, err := isReadable(path)
	if err != nil {
		return err // error type is already ErrWalkPathNotReadable
	}
	if readable {
		dirents, err := os.ReadDir(path)
		if err != nil {
			return NewError(ErrWalkCouldNotListFiles, err.Error())
		}
		for _, ent := range dirents {
			info, err := ent.Info()
			if err != nil {
				// todo: work out how this should report problems when reading file info
				res.add(path, os.ModeIrregular, err)
				continue
			}
			// Only directory entries that match the target type will be added to the results slice.
			// If the entry type matches conf.TargetType, then, if the entry is a Regular File (a regular
			// file is a directory entry with no ModeType bits set), then a final GlobPattern match
			// check is made. If the entry name matches the glob pattern, it is entered into the results
			// slice. If the entry is not a regular file, but is still a match, it is entered into the
			// results slice and the ModeType is attached.
			if conf.TargetType.Matches(info.Mode()) {
				if conf.TargetType.IsRegular() && info.Mode().IsRegular() {
					if globMatch(ent.Name(), conf.GlobPattern) {
						res.add(path, info.Mode().Type(), nil)
					}
				} else {
					res.add(path, info.Mode().Type(), nil)
				}
			}
			// todo: adjust history so that it also records the symlink that led to the path recorded in the history.
			// That way, we can detect symbolic link loops.
			if info.IsDir() {
				// todo: walk directory
			}
			// todo: if file is of symlink type pointing to directory, are we following? walk it
			if info.Mode()&os.ModeSymlink == 0 {
				if leadsToDir(filepath.Join(path, ent.Name())) {
					// todo: symlink walker ran here
				}
			}
			// todo: if file is of symlink type pointing to non-dir, are we following?
			// outta here
		}
	}
	return nil
}
