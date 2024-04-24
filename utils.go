package symwalker

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
	objf "github.com/orme292/objectify"
)

// getCur returns the current directory path string.
// It uses osext.ExecutableFolder to get the path of the current executable.
// If there is an error, it returns "/" as the default path.
//
// Returns:
// - the current directory path string
func getCur() string {
	cur, err := osext.ExecutableFolder()
	if err != nil {
		cur = "/"
	}
	return cur
}

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

// fullPathSafe returns the absolute and cleaned version of the given path.
// It is a wrapper for filepath.Abs. The filepath.Abs function cleans the
// path.
// Parameters:
// - path: the path string to be processed
// Returns:
// - the absolute and cleaned path string
// Considered unsafe because potential errors from filepath.Abs are ignored.
func fullPathSafe(path string) string {
	path, err := filepath.Abs(curPathIf(path))
	if err != nil {
		path, err = os.Getwd()
		if err != nil {
			path = "/"
		}
	}
	return path
}

// curPathIf calls getCur() on the given path if it is an empty string.
// If the provided path is not empty, it does nothing and returns the path.
// Parameters:
// - path: the path string to be processed
// Returns:
// - the given path if non-empty, otherwise it returns the current directory path string.
func curPathIf(path string) string {
	path = strings.TrimSpace(path)
	if path == "" || path == "." {
		path = getCur()
	}
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

func getFileData(path string) (*objf.FileObj, error) {
	return objf.File(path, objf.SetsAll())
}

// isReadable checks if a file or directory at the given path is readable.
// It opens the path using os.Open. If there are any errors during this process,
// it returns false and the error. Otherwise, it returns true and a nil error.
// Parameters:
// - path: the path of the file or directory to be checked
// Returns:
// - a boolean indicating if the path is readable
// - an error if there was an error during the process
func isReadable(path string, ent entType) (bool, error) {
	if ent == entTypeOther {
		return false, nil
	}
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
