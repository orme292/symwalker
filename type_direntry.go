package symwalker

// DirEntry holds directory entry information.
// Marked is a flag used to indicate whether the DirEntry
// has been processed/walked.
type DirEntry struct {
	Path   string
	Marked bool
}
