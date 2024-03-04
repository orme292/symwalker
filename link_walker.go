package symwalker

import (
	"os"
	"path/filepath"
)

// /Users/aorme/symlinkpath => /test/path1
func linkWalk(conf *Conf, link link, remDepth int, res *Results, hist *history) (err error) {
	// todo: check depth
	if remDepth == 0 {
		return NewError(OpsMaxDepthReached, s("depth exceeded before %q", path))
	}
	remDepth = depth(remDepth)

	// todo: record in history.
	if hist.count(link.referrer) > 1 {
		// todo: how do we handle this
		return nil
	}
	_ = hist.add(path, referrer)

	slink, err := filepath.EvalSymlinks(path)
	if err != nil {
		return err
	}

	readable, err := isReadable(slink)
	if err != nil {
		return err
	}
	if readable {
		sinfo, err := os.Lstat(slink)
		if err != nil {
			return err
		}
		if sinfo.Mode()&
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
	_, err = os.Lstat(targetPath)
	// todo: WHAT IF THE TARGET IS A SYMLINK?
	if err != nil {
		// todo: which errors could happen here? how to handle?
		return false
	}
	return false
}
