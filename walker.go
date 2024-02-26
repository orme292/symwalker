package symwalker

import (
	"os"
)

func SymWalker(conf *Conf) (res *Results, err error) {
	// Check to make sure that the StartPath is readable and also that it is
	// not a Symlinked file or path. Even though the StartPath, as a symlink,
	// could be pointing to a directory, it doesn't seem to make sense to
	// walk it, since all paths under it would be different from the StartPath.
	// Finally, we check to make sure the StartPath is a directory.
	info, e := os.Lstat(conf.StartPath)
	if e != nil {
		return nil, NewError(ErrStartPathNotReadable, e.Error())
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		return nil, NewError(ErrStartPathIsSymlink,
			s("%q => %q", conf.StartPath, evalLinkNonFatal(conf.StartPath)))
	}
	if info.Mode()&os.ModeDir == 0 {
		return nil, NewError(ErrStartPathIsNotDir, s("%q", conf.StartPath))
	}

	//
	if conf.Depth >= 0 {
		conf.Depth++
	}
	return nil, nil
}
