package swalker

type SymConf struct {
	StartPath      string
	FollowSymlinks bool
}

type WalkerEntry struct {
	Path string
}

type WalkerResults []WalkerEntry

type History []WalkerEntry

var history History
