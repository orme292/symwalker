// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

// DirEntry holds directory entry information.
// marked is a flag used to indicate whether the DirEntry
// has been processed/walked.
type DirEntry struct {
	Path   string
	marked bool
}
