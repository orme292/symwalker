package old

import (
	"errors"
	"path/filepath"
)

type Conf struct {
	StartPath      string
	Depth          int
	FollowSymLinks bool
	TargetType     WalkTarget
	GlobPattern    string // only applies when the WalkTarget is Regular, not symlink or dir

	// used for testing
	debug bool
}

func (c *Conf) GlobCheck() error {
	// GlobPattern must have a value, default is to match all files ("*")
	if c.GlobPattern == emptyString {
		c.GlobPattern = GlobMatchAll
		return nil
	}
	_, err := filepath.Match(c.GlobPattern, "/users/sample/grow/451614/!$!$!")
	// filepath.Match only throws one type of error, filepath.ErrBadPattern
	if errors.Is(err, filepath.ErrBadPattern) {
		return NewError(ErrConfGlobalMalformed, c.GlobPattern)
	}
	return nil
}
