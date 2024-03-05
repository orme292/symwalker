package symwalker

import (
	"os"
	"path/filepath"
)

// /Users/aorme/symlinkpath => /test/path1
func linkWalk(conf *Conf, basePath string, referrer string, remDepth int, res *Results, hist *history) (err error) {
	// todo: check depth
	if remDepth == 0 {
		return NewError(OpsMaxDepthReached, s("depth exceeded before %q", basePath))
	}
	remDepth = depth(remDepth)

	// todo: record in history.

	info, mode, link, err := allFileInfoNonFatal(basePath)
	if err != nil {
		return
	}

	if hist.count(basePath) > 0 {
		return NewError(ErrWalkPotentialSymlinkLoop, s("potential symlink loop detected at %q", basePath))
	}

	err = hist.add(basePath, referrer)
	if err != nil {
		return // The only possible error is a SymlinkLoop warning.
	}

	if conf.TargetType.Matches(mode) {
		res.add(basePath, mode, nil)
	}

	if info.IsDir() {
		dirents, err := os.ReadDir(basePath)
		if err != nil {
			return
		}
		for _, ent := range dirents {
			einfo, err := ent.Info()
			if err != nil {
				res.add(filepath.Join(basePath, ent.Name()), os.ModeIrregular, err)
				continue
			}

			if conf.TargetType.Matches(einfo.Mode()) {
				if conf.TargetType.IsRegular() && einfo.Mode().IsRegular() {
					if globMatch(ent.Name(), conf.GlobPattern) {
						res.add(filepath.Join(basePath, ent.Name()), einfo.Mode(), nil)
					}
				} else {
					res.add(filepath.Join(basePath, ent.Name()), einfo.Mode(), nil)
				}
			}

			if einfo.IsDir() || einfo.Mode()&os.ModeSymlink == 0 {
				if leadsToDir(filepath.Join(basePath, ent.Name())) ||
					leadsToTarget(filepath.Join(basePath, ent.Name()), conf.TargetType) {
					err = linkWalk(conf, filepath.Join(basePath, ent.Name()), basePath, remDepth, res, hist)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	if info.Mode()&os.ModeSymlink == 0 {
		if leadsToDir(filepath.Join(basePath, filepath.Base(link))) || leadsToTarget(filepath.Join(basePath, filepath.Base(link)), conf.TargetType) {
			err = linkWalk(conf, filepath.Join(basePath, filepath.Base(link)), basePath, remDepth, res, hist)
			if err != nil {
				return
			}
		}
	}

	// todo: evaluate link
	// todo: => if link evaluates to target type, record in results, continue
	// todo: => if link evaluates to dir, walk it
	// todo: => if link evaluates to another link...
	// todo: => => does the link ultimately evaluate to a directory? symDirWalk it
	// todo: => => does the link ultimately evaluate to a target? symwalk it
	// todo: => => if neither, move on
	return nil
}

// leadsToDir determines if the given path to a symlink ultimately leads to a directory.
func leadsToDir(path string) bool {
	targetPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		// todo: which errors could happen here? how to handle?
		return false
	}
	info, err := os.Lstat(targetPath)
	if err != nil {
		// todo: which errors could happen here? how to handle?
		return false
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		return leadsToDir(targetPath)
	}
	if info.IsDir() {
		return true
	}
	return false
}

// leadsToTarget determines if the given path to a symlink ultimately leads to the targeted type.
func leadsToTarget(path string, target WalkTarget) bool {
	targetPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		// todo: which errors could happen here? how to handle?
		return false
	}
	info, err := os.Lstat(targetPath)
	if err != nil {
		// todo: which errors could happen here? how to handle?
		return false
	}

	if target.Matches(info.Mode()) {
		return true
	}

	if info.Mode()&os.ModeSymlink == 0 {
		return leadsToTarget(targetPath, target)
	}

	return false
}
