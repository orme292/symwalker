package swalker

type DirEntries []DirEntry

func (re *DirEntries) add(p string) {
	if !re.pathExists(p) {
		*re = append(*re, DirEntry{Path: p})
	}
}

func (re *DirEntries) pathExists(p string) bool {
	for _, entry := range *re {
		if entry.Path == p {
			return true
		}
	}
	return false
}

type Results struct {
	Dirs   DirEntries
	Files  DirEntries
	Others DirEntries
}

func newResults(conf SymConf) Results {
	return Results{
		Dirs:   conf.dirs,
		Files:  conf.files,
		Others: conf.others,
	}
}
