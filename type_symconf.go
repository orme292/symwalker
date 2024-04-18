// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

// SymConf is the configuration object for symwalker. It also
// acts as storage for results. Each exported field can be set
// manually by instantiating a SymConf object, or functional
// options below can be used.
type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
	dirs           DirEntries
	files          DirEntries
	others         DirEntries
}

// NewSymConf implements the functional options pattern to handle
// the package configuration dependency. Example:
/*
	conf := NewSymConf(
		WithStartPath("/tests/users"),
		WithFollowedSymLinks(),
		WithLogging(),
	)
*/
func NewSymConf(options ...OptFunc) *SymConf {

	c := &SymConf{}
	for _, option := range options {
		option(c)
	}
	return c

}

type OptFunc func(*SymConf)

// WithStartPath accepts the path in which the directory walk
// will start. It MUST be a directory, not a file, or a symlink to
// a directory.
func WithStartPath(startPath string) OptFunc {
	return func(c *SymConf) {
		c.StartPath = startPath
	}
}

// WithFollowedSymLinks is a flag that the directory walk functions
// will check before evaluating symlinks.
func WithFollowedSymLinks() OptFunc {
	return func(c *SymConf) {
		c.FollowSymlinks = true
	}
}

// WithLogging is a flag used to determine whether log messages are
// printed to the screen.
func WithLogging() OptFunc {
	return func(c *SymConf) {
		c.Noisy = true
	}
}
