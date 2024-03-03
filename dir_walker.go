package symwalker

import (
	"os"
	"path/filepath"
)

func startLoop(conf *Conf) (res Results, err error) {
	hist := make(history, 0)
	err = dirWalk(conf, conf.StartPath, conf.Depth, &res, &hist)
	return
}

func dirWalk(conf *Conf, path string, remDepth int, res *Results, hist *history) (err error) {
	// if the remaining depth (remDepth) is 0, do not continue.
	if remDepth == 0 {
		return NewError(OpsMaxDepthReached, s("depth exceeded before %q", path))
	}
	remDepth = depth(remDepth)

	// hist.count() checks to see if the current path has already been
	// visited. If it has, then we return instead of continuing.
	if hist.count(path) > 1 {
		return nil
	}

	// Otherwise, we add this to the history, with an empty referrer.
	_ = hist.add(path, emptyString)

	// Confirm that the path is readable before continuing. If not,
	// return the error, which will be ErrWalkPathNotReadable.
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
				// If not able to get the info for a particular directory entry, then
				// it is still added to the results struct as an irregular file. The
				// returned error is attached. We just continue on after this.
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

			// If the entry is a directory, we walk it.
			if info.IsDir() {
				err = dirWalk(conf, filepath.Join(path, ent.Name()), remDepth, res, hist)
				if err != nil {
					return err
				}
			}

			// If configured to follow symlinks, and the entry is a symlink, then we check if it
			// ultimately leads to a directory. If it does, we walk it down the line. Otherwise,
			// we see if the symlink leads to a file that matches the TargetType. If the TargetType
			// is a directory, then it will be picked up by the previous statement.
			if conf.FollowSymLinks && info.Mode()&os.ModeSymlink == 0 {
				if leadsToDir(filepath.Join(path, ent.Name())) {
					// todo: symlink walker ran here
				} else {
					if leadsToTarget(filepath.Join(path, ent.Name()), conf.TargetType) {
						// todo: what do we do here?
					}
				}
			}
		}
	}
	return nil
}
