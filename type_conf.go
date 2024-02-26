package symwalker

type Conf struct {
	StartPath      string
	Depth          int
	FollowSymLinks bool
	TargetType     WalkTarget
	Blob           string

	// used for testing
	debug bool
}
