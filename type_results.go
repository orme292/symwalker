package symwalker

import (
	objf "github.com/orme292/objectify"
)

// DirEntries holds each directory entry.
type DirEntries []DirEntry

// add appends a new DirEntry to the DirEntries slice.
// path is checked to see if it exists before adding. If it does, no
// new element is appended.
// If withData flag is true, then the FileObj field is populated
// with file data from objectify, and the path field, then appended.
// If withData is false, then the Path field is populated and appended.
func (re *DirEntries) add(path string, withData bool) {
	var de DirEntry
	if !re.pathExists(path) {
		if withData {
			fo, err := objf.File(path, objf.SetsAll())
			if err != nil {
				fo = &objf.FileObj{}
			}
			de = DirEntry{Path: path, FileObj: fo}
		} else {
			de = DirEntry{Path: path}
		}
		*re = append(*re, de)
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
