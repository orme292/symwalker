package symwalker

// lineHistory holds a directory entry branch line history.
type lineHistory []string

// pathExists checks if the given path exists in
// the lineHistory slice. It returns true if the path
// exists, otherwise it returns false.
func (lh lineHistory) pathExists(path string) bool {
	for i := range lh {
		if lh[i] == path {
			return true
		}
	}
	return false
}

// add appends the given path to the lineHistory slice.
// Parameters:
// - path: the path string to be added
// Returns:
// - a new lineHistory slice with the path appended
func (lh lineHistory) add(path string) lineHistory {
	return append(lh, path)
}

// branch eliminates the lineHistory slice's pointer to the
// underlying array and creates a separate slice with
// the contents of the original lineHistory.
// This keeps each branching directory's lineHistory focused to
// a walk's particular line.
// Parameters:
// - lh: the lineHistory slice to be refreshed
// Returns:
// - a new lineHistory slice with the elements from the current lineHistory slice appended
func (lh lineHistory) branch() lineHistory {
	var newLineHistory lineHistory
	return append(newLineHistory, lh...)
}
