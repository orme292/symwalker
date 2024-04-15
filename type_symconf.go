package swalker

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
	maxLoops       int
	pending        PendingEntries
	history        historyEntries
	results        ResultEntries
}

const (
	MILLION = 1000000000
)

func NewSymConf(options ...func(*SymConf)) *SymConf {
	c := &SymConf{
		maxLoops: MILLION,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func WithStartPath(startPath string) func(*SymConf) {
	return func(c *SymConf) {
		c.StartPath = startPath
	}
}

func WithLogging() func(*SymConf) {
	return func(c *SymConf) {
		c.Noisy = true
	}
}

func FollowSymLinks() func(*SymConf) {
	return func(c *SymConf) {
		c.FollowSymlinks = true
	}
}

func OverridesMaxLoopsValue(maxLoops int) func(*SymConf) {
	return func(c *SymConf) {
		c.maxLoops = maxLoops
	}
}
