package swalker

type lineHistory []string

func (lh lineHistory) pathExists(path string) bool {
	for i := range lh {
		if lh[i] == path {
			return true
		}
	}
	return false
}

func (lh lineHistory) add(path string) lineHistory {
	return append(lh, path)
}

func (lh lineHistory) refresh() lineHistory {
	var newLineHistory lineHistory
	return append(newLineHistory, lh...)
}
