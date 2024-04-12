package swalker

type HistoryEntries []string

func (he *HistoryEntries) add(p string) {
	if !he.pathExists(p) {
		*he = append(*he, p)
	}
}

func (he *HistoryEntries) pathExists(p string) bool {
	for _, entry := range *he {
		if entry == p {
			return true
		}
	}
	return false
}
