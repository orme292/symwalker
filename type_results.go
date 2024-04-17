package symwalker

// DirEntries holds each directory entry.
type DirEntries []DirEntry

// add adds a new directory entry to the DirEntries
// slice if the path does not already exist.
func (re *DirEntries) add(p string) {
	if !re.pathExists(p) {
		*re = append(*re, DirEntry{Path: p})
	}
}

// pathExists checks if a given path already exists
// in the DirEntries slice. It returns true if the
// path exists, otherwise it returns false.
func (re *DirEntries) pathExists(p string) bool {
	for _, entry := range *re {
		if entry.Path == p {
			return true
		}
	}
	return false
}

// Results contains each recordable directory entry
// returned from the directory walk loop. This is the
// exported SymWalk function return object.
type Results struct {
	Dirs   DirEntries
	Files  DirEntries
	Others DirEntries
}

// newResults returns a Results object that is
// populated with values from the SymConf object
// passed to it.
func newResults(conf SymConf) Results {
	return Results{
		Dirs:   conf.dirs,
		Files:  conf.files,
		Others: conf.others,
	}
}
