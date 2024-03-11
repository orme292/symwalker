package old

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

func fileInfoNonFatal(path string) os.FileInfo {
	info, _ := os.Lstat(path)
	return info
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

func allFileInfoNonFatal(path string) (info os.FileInfo, mode os.FileMode, link string, err error) {
	if ok, err := isReadable(path); !ok {
		return nil, os.ModeIrregular, emptyString, err
	}
	info, err = os.Lstat(path)
	if err != nil {
		return
	}
	mode = info.Mode()
	link, _ = filepath.EvalSymlinks(path)
	return
}

func globMatch(path string, pattern string) bool {
	found, err := filepath.Match(pattern, path)
	if err != nil {
		return false
	}
	return found
}

func depth(val int) int {
	if val != DepthInfinite {
		val--
	}
	return val
}
