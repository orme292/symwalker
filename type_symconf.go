package swalker

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
	dirs           DirEntries
	files          DirEntries
	others         DirEntries
}

func NewSymConf(options ...func(*SymConf)) *SymConf {

	c := &SymConf{}
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

func FollowsSymLinks() func(*SymConf) {
	return func(c *SymConf) {
		c.FollowSymlinks = true
	}
}
