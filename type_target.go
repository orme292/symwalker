package symwalker

import (
	"os"
)

type WalkTarget int

const (
	TargetRegular WalkTarget = iota
	TargetSymlink
	TargetDir
)

func (wt WalkTarget) String() string {
	return [...]string{"Dir", "Symlink", "Regular"}[wt]
}

func (wt WalkTarget) IsDir() bool {
	return wt == TargetDir
}

func (wt WalkTarget) IsSymlink() bool {
	return wt == TargetSymlink
}

func (wt WalkTarget) IsRegular() bool {
	return wt == TargetRegular
}

func (wt WalkTarget) Is() any {
	switch wt {
	case TargetDir:
		return os.ModeDir
	case TargetSymlink:
		return os.ModeSymlink
	case TargetRegular:
		return 0
	default:
		return 0
	}
}

func (wt WalkTarget) Matches(mode os.FileMode) bool {
	if wt.IsDir() {
		return wt.IsDir() == mode.IsDir()
	}
	if wt.IsRegular() {
		return wt.IsRegular() == mode.IsRegular()
	}
	if wt.IsSymlink() {
		return wt.IsSymlink() == (mode&os.ModeSymlink != 0)
	}
	return false
}
