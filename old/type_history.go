package old

type referrerList []string

type historyEntry struct {
	path      string
	count     uint8
	referrers referrerList // a referrer is the path to a  symlink that targeted the path in the entry.
}

type history []*historyEntry

// add adds, or increments the count of, a historyEntry in the history slice. If a historyEntry's
// path value matches the provided path, then the historyEntry's count is incremented. After, if
// a referrer path is provided it is appended to the historyEntry's referrers slice.
// If no historyEntry's path matches the provided path, then a new historyEntry is appended
// to the history slice. The only potential return error is ErrWalkPotentialSymlinkLoop, which
// is returned if a historyEntry's path value matches the path, and the historyEntry's referrerList
// already contains an occurrence of the provided referrer.
func (h *history) add(path, referrer string) error {
	path = format(path)
	for _, entry := range *h {
		if entry.path == path {
			entry.count++
			if referrer != emptyString {
				if entry.referrers.add(referrer) > 1 {
					return NewError(ErrWalkPotentialSymlinkLoop, s("%q => %q", path, referrer))
				}
			}
			return nil
		}
	}
	*h = append(*h, &historyEntry{
		path:      path,
		count:     1,
		referrers: []string{referrer},
	})
	return nil
}

// count searches the history slice for an entry whose path matches
// the provided path. If it does, the entry's count value is returned.
func (h *history) count(path string) int {
	for _, entry := range *h {
		if entry.path == path {
			return int(entry.count)
		}
	}
	return 0
}

// count searches the referrerList slice for occurrences of
// the provided path string, and then returns a cumulative count.
func (rl *referrerList) count(path string) int {
	var count int
	if path == emptyString {
		return count
	}
	for _, referrer := range *rl {
		if referrer == path {
			count++
		}
	}
	return count
}

// add adds an entry to the referrerList and returns referrerList.count
func (rl *referrerList) add(referrer string) int {
	*rl = append(*rl, referrer)
	return rl.count(referrer)
}
