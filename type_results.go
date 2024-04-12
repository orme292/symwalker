package swalker

type DirEntry struct {
	Path   string
	Marked bool
}

func (de *DirEntry) Mark() {
	de.Marked = true
}

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

func (re *ResultEntries) PopOut(index int) ResultEntries {
	return append((*re)[:index], (*re)[index+1:]...)
}
