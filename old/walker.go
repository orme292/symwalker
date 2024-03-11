package old

import (
	"os"
	"path/filepath"
)

func SymWalker(conf *Conf) (res Results, err error) {
	conf.StartPath = absPathNonFatal(filepath.Clean(conf.StartPath))

	// conf.GlobPattern must have a value. conf.GlobCheck() applies a default
	// if empty or checks that the provided pattern is valid. GlobPattern
	// must follow the rules set by the filepath library, spec. filepath.Match.
	// See: https://golang.org/pkg/path/filepath/#Match
	e := conf.GlobCheck()
	if e != nil {
		return nil, e // Error is already type ErrConfGlobalMalformed
	}

	// Check to make sure that the StartPath is readable and also that it is
	// not a Symlinked file or path. Even though the StartPath, as a symlink,
	// could be pointing to a directory, it doesn't seem to make sense to
	// walk it, since all paths under it would be different from the StartPath.
	// Finally, we check to make sure the StartPath is a directory and not a file.
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

	// A depth setting of 0 or less in the configuration is considered to be infinite.
	// When the walker is running, a value of 0 will cause it to stop walking. So,
	// if it is set to 0 when started, nothing would happen (which doesn't make a ton
	// of sense to me). It is set to -1 to avoid the problem. When not set to
	// -1 (DepthInfinite), the depth value is passed to each subsequent directory and
	// 1 is subtracted from it. Once it reaches a value of 0, it stops walking further.
	if conf.Depth <= 0 {
		conf.Depth = DepthInfinite
	}

	hist := make(history, 0)
	err = dirWalk(conf, conf.StartPath, conf.Depth, res, hist)
	return nil, nil
}
