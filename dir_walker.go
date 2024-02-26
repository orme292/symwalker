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

	var skip bool
	// todo: check if dir has already been walked (in history), if it has, skip = true

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
			hist.add(filepath.Join(path, file.Name()))
			inf, err := file.Info()
			if err != nil {
				// todo: report a real error
				panic(err)
			}
			// todo: if file is of target type, add to results
			// todo: if file is of dir type, walk it
			// todo: if file is of symlink type pointing to directory, are we following? walk it
			// todo: if file is of symlink type pointing to non-dir, are we following? if symlink target is target type, add to results
			// outta here
		}
	}
	return nil
}
