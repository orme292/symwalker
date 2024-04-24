package symwalker

import (
	objf "github.com/orme292/objectify"
)

// DirEntry holds directory entry information.
// marked is a flag used to indicate whether the DirEntry
// has been processed/walked.
type DirEntry struct {
	Path    string
	FileObj *objf.FileObj
	marked  bool
}
