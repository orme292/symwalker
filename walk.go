package symwalker

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walker(root string, target WalkTarget, follow bool) (targets []string, err error) {
	info, err := os.Stat(root)
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		panic(WalkErrRootIsSymlink)
	}
	if !info.Mode().IsDir() {
		panic(WalkErrRootIsFile)
	} else {
		dTargets, err := dirWalker(root, target.Is().(os.FileMode), follow, nil)
		if err != nil {
			panic(err)
		}
		targets = append(targets, dTargets...)
	}
	return nil, nil
}

func dirWalker(p string, t os.FileMode, follow bool, re error) (targets []string, err error) {
	if re != nil {
		panic(err)
	}
	files, err := os.ReadDir(p)
	if err != nil {
		return
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting info: ", err)
			continue
		}
		if info.Mode().Type() == t {
			fmt.Println("Target: ", filepath.Join(p, file.Name()))
			targets = append(targets, filepath.Join(p, file.Name()))
		}
		if info.Mode().Type() == os.ModeDir {
			dTargets, err := dirWalker(filepath.Join(p, file.Name()), t, follow, re)
			if err != nil {
				return nil, err
			}
			targets = append(targets, dTargets...)
		}
		if info.Mode().Type()&os.ModeSymlink == os.ModeSymlink {
			if follow {
				dTargets, err := linkWalker(filepath.Join(p, file.Name()), t, re)
				if err != nil {
					return nil, err
				}
				targets = append(targets, dTargets...)

			}
		}
	}
	return
}

func linkWalker(base string, t os.FileMode, re error) (targets []string, err error) {
	if re != nil {
		panic(re)
	}

	symlink, err := filepath.EvalSymlinks(base)
	if err != nil {
		panic(err)
	}

	symlinkInfo, err := os.Lstat(base)
	if err != nil {
		return nil, WalkErrSymLinkLoop
	}

	if symlinkInfo.Mode().Type() == t {
		// fmt.Println("Symlinked Target: ", base)
		targets = append(targets, base)
	}

	files, err := os.ReadDir(symlink)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			panic(err)
		}
		if info.Mode().Type() == t {
			fmt.Println("Symlinked Target: ", filepath.Join(base, file.Name()))
		}
		if info.Mode().Type() == os.ModeDir {
			dTargets, err := dirWalker(filepath.Join(base, file.Name()), t, true, re)
			if err != nil {
				panic(err)
			}
			targets = append(targets, dTargets...)
		}
		if info.Mode().Type()&os.ModeSymlink == os.ModeSymlink {
			dTargets, err := linkWalker(filepath.Join(base, file.Name()), t, re)
			if err != nil {
				panic(err)
			}
			targets = append(targets, dTargets...)
		}
	}
	return
}
