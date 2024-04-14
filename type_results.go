package swalker

type ResultEntries []DirEntry

func (re *ResultEntries) add(p string) {
	if !re.pathExists(p) {
		*re = append(*re, DirEntry{Path: p})
	}
}

func (re *ResultEntries) combine(entries ResultEntries) {
	*re = append(*re, entries...)
}

func (re *ResultEntries) pathExists(p string) bool {
	for _, entry := range *re {
		if entry.Path == p {
			return true
		}
	}
	return false
}

func (re *ResultEntries) PopOut(index int) {
	*re = append((*re)[:index], (*re)[index+1:]...)
}
