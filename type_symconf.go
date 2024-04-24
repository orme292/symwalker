package symwalker

// SymConf is the configuration object for symwalker. It also
// acts as storage for results. Each exported field can be set
// manually by instantiating a SymConf object, or functional
// options below can be used.
type SymConf struct {
	StartPath      string
	Depth          int
	FileData       bool
	FollowSymlinks bool

	Noisy bool

	WithoutFiles bool

	dirs   DirEntries
	files  DirEntries
	others DirEntries
}

const (
	FLAT     = 1
	INFINITE = 0
)

// NewSymConf implements the functional options pattern to handle
// the package configuration dependency. Example:
/*
	conf := NewSymConf(
		WithStartPath("/tests/users"),
		WithFollowedSymLinks(),
		WithLogging(),
	)
*/
func NewSymConf(path string, options ...OptFunc) *SymConf {

	path = fullPathSafe(path)

	c := &SymConf{
		StartPath: path,
	}
	for _, option := range options {
		option(c)
	}
	return c

}

type OptFunc func(*SymConf)

// WithDepth limits how many levels SymWalker will walk below
// the provided StartPath directory. Value n can be set to
// any number. INFINITE is equal to 0 or infinite depth. FLAT
// is equal to 1 -- SymWalker will not go below the top level.
func WithDepth(n int) OptFunc {
	return func(c *SymConf) {
		c.Depth = n
	}
}

func WithFileData() OptFunc {
	return func(c *SymConf) {
		c.FileData = true
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

// WithoutFiles is a flag that SymWalker checks before evaluating
// non-directory directory entries or symlinks
func WithoutFiles() OptFunc {
	return func(c *SymConf) {
		c.WithoutFiles = true
	}
}
