package symwalker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func s(f string, v ...any) string {
	return fmt.Sprintf(f, v...)
}

func evalLinkNonFatal(path string) string {
	path, _ = filepath.EvalSymlinks(path)
	return path
}

func absPathNonFatal(path string) string {
	path, _ = filepath.Abs(filepath.Clean(path))
	return path
}

func isReadable(path string) (bool, error) {
	_, err := os.Open(absPathNonFatal(path))
	if err != nil {
		return false, NewError(ErrWalkPathNotReadable, err.Error())
	}
	return true, nil
}

func format(path string) string {
	clean := filepath.Clean(path)
	if clean[0] == '.' {
		clean = absPathNonFatal(path)
	} else if len(clean) > 0 && clean[0] != '/' {
		clean = s("%s%s", "/", clean)
	}
	return strings.TrimSuffix(clean, "/")
}
