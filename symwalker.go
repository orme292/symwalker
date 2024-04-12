package swalker

import (
	"errors"
)

var (
	ErrPathNotFound = errors.New("Path not found in PendingEntries")
)

func walkPending(conf *SymConf, base string) (Results ResultEntries, err error) {
	return
}
