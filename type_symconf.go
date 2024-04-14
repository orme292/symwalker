package swalker

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
	Noisy          bool
	pending        PendingEntries
	history        historyEntries
	results        ResultEntries
}
