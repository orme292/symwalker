package symwalker

import (
	"os"
	"path/filepath"
)

func SymWalker(conf *Conf) (res *Results, err error) {
	conf.StartPath = absPathNonFatal(filepath.Clean(conf.StartPath))

	// conf.GlobPattern must have a value. conf.GlobCheck() applies a default
	// and checks that it's valid.
	e := conf.GlobCheck()
	if e != nil {
		return nil, e // Error is already type ErrConfGlobalMalformed
	}

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

	// A depth setting of 0 or less in the configuration is equal to infinite depth.
	if conf.Depth <= 0 {
		conf.Depth = DepthInfinite
	}
	return nil, nil
}
