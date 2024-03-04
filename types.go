package symwalker

const (
	emptyString = ""

	DepthInfinite = -1

	GlobMatchAll = "*"
)

type link struct {
	next     string
	base     string
	referrer string
}
