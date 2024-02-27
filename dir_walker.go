package symwalker

import (
	"os"
	"path/filepath"
)

func startLoop(conf *Conf) (res Results, err error) {
	hist := make(history)
	hist.add(conf.StartPath)
	err = dirWalk(conf, conf.StartPath, conf.Depth, res, &hist)
	return
}

func dirWalk(conf *Conf, path string, remDepth int, res Results, hist *history) (err error) {
	if remDepth == 0 {
		return nil
	}
	remDepth = depth(remDepth)

	var skip bool
	// todo: check if dir has already been walked (in history), if it has, skip = true
	if hist.count(path) > 1 {
		skip = true
	}

	if skip {
		return nil
	}
	readable, err := isReadable(path)
	if err != nil {
		//todo: report a real error
		panic(err)
	}
	if readable {
		files, err := os.ReadDir(path)
		if err != nil {
			// todo: report a real error
			panic(err)
		}
		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				// todo: work out how this should report problems when reading file info
				res.add(path, os.ModeIrregular, err)
				continue
			}
			if conf.TargetType.Matches(info.Mode()) {
				if conf.TargetType.IsRegular() && info.Mode().IsRegular() {
					if globMatch(filepath.Join(path, file.Name()), conf.GlobPattern) {
						res.add(path, info.Mode().Type(), nil)
					}
				} else {
					res.add(path, info.Mode().Type(), nil)
				}
			}
			if info.IsDir() {
				// todo: walk directory
			}
			// todo: if file is of symlink type pointing to directory, are we following? walk it
			if info.Mode()&os.ModeSymlink == 0 {
				if leadsToDir(filepath.Join(path, file.Name())) {
					// todo: symlink walker ran here
				}
			}
			// todo: if file is of symlink type pointing to non-dir, are we following? if symlink target is target type, add to results
			// outta here
		}
	}
	return nil
}
