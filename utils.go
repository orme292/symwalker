// Package symwalker is a directory tree walker with symlink loop protection.
// It builds a separate history for each branching sub-directory below
// a given starting path. When evaluating a new symlink that targets a
// directory, the branch history is checked before walking the directory.
package symwalker

import (
	"log"
	"os"
	"path/filepath"
)

// joinPaths concatenates two path strings and returns the cleaned result.
// It is a wrapper for filepath.Join. filepath.Join cleans the path.
// Parameters:
// - start: the starting path string (usually filepath.Dir(string))
// - end: the ending path string (usually filepath.Base(string))
// Returns:
// - the concatenated and cleaned path string
func joinPaths(start string, end string) string {
	return filepath.Join(start, end)
}

// fullPathUnsafe returns the absolute and cleaned version of the given path.
// It is a wrapper for filepath.Abs. The filepath.Abs function cleans the
// path.
// Parameters:
// - path: the path string to be processed
// Returns:
// - the absolute and cleaned path string
// Considered unsafe because potential errors from filepath.Abs are ignored.
func fullPathUnsafe(path string) string {
	path, _ = filepath.Abs(path)
	return path
}

// noise logs the provided message if the noisy parameter is true.
// Parameters:
// - noisy: a boolean indicating if the message should be logged (conf.Noisy)
// - f: a format string for the log message (same as fmt.Sprintf)
// - v: optional values to be formatted in the log message
func noise(noisy bool, f string, v ...interface{}) {
	if noisy {
		log.Printf(f, v...)
	}
}

// isReadable checks if a file or directory at the given path is readable.
// It opens the path using os.Open. If there are any errors during this process,
// it returns false and the error. Otherwise, it returns true and a nil error.
// Parameters:
// - path: the path of the file or directory to be checked
// Returns:
// - a boolean indicating if the path is readable
// - an error if there was an error during the process
func isReadable(path string) (bool, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
