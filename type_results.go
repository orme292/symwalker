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

// pathAddedWithExistsCheck will check if the path already exists in the `re` object.
// If it exists, the function returns FALSE (since it was not added). If the path does not exist,
// it is added and the function returns TRUE (since the path was added).
func (re *ResultEntries) pathAddedWithExistsCheck(path string) bool {
	if re.pathExists(path) {
		return false
	}
	re.add(path)
	return true
}

func (re *ResultEntries) PopOut(index int) {
	*re = append((*re)[:index], (*re)[index+1:]...)
}
