package symwalker

import (
	"os"
)

type WalkErr string

func (e WalkErr) Error() string {
	return string(e)
}
func (e WalkErr) String() string {
	return string(e)
}
func (e WalkErr) Is(err error) bool {
	return string(e) == err.Error()
}

func (e WalkErr) As(target interface{}) bool {
	switch v := target.(type) {
	case *WalkErr:
		*v = e
		return true
	default:
		return false
	}
}

const (
	empty = ""

	WalkErrRootIsSymlink          WalkErr = "Root directory cannot be a symlink."
	WalkErrRootIsFile             WalkErr = "Root directory cannot be a file."
	WalkErrSymLinkLoop            WalkErr = "Possible symlink loop detected."
	WalkErrFilePasssedToDirWalker WalkErr = "A file was passed to the DirWalker"
)

type WalkTarget int

const (
	Dir WalkTarget = iota
	Symlink
	Regular
)

func (wt WalkTarget) String() string {
	return [...]string{"Dir", "Symlink", "Regular"}[wt]
}

func (wt WalkTarget) IsDir() bool {
	return wt == Dir
}

func (wt WalkTarget) IsSymlink() bool {
	return wt == Symlink
}

func (wt WalkTarget) IsRegular() bool {
	return wt == Regular
}

func (wt WalkTarget) Is() any {
	switch wt {
	case Dir:
		return os.ModeDir
	case Symlink:
		return os.ModeSymlink
	case Regular:
		return 0
	default:
		return 0
	}
}
