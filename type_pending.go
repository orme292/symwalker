package swalker

type PendingEntries []DirEntry

func (pe *PendingEntries) add(p string) {
	if !pe.pathExists(p) {
		*pe = append(*pe, DirEntry{Path: p})
	}
}

func (pe *PendingEntries) mark(p string) {
	for i, entry := range *pe {
		if entry.Path == p {
			(*pe)[i].Mark()
			break
		}
	}
}

func (pe *PendingEntries) pathExists(p string) bool {
	for _, entry := range *pe {
		if entry.Path == p {
			return true
		}
	}
	return false
}

func (pe *PendingEntries) popOut(index int) PendingEntries {
	return append((*pe)[:index], (*pe)[index+1:]...)
}
