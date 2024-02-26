package symwalker

import (
	"os"
)

type WalkTarget int

const (
	TargetDir WalkTarget = iota
	TargetSymlink
	TargetRegular
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
