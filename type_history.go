package swalker

import (
	"log"
)

type historyEntries []*string

func (he *historyEntries) pathExists(path string) bool {
	for index := range *he {
		if *(*he)[index] == path {
			log.Println("history search (found): ", path)
			return true
		}
	}
	return false
}

// pathAddedWithExistsCheck will check if the path already exists in the `he` object.
// If it exists, the function returns FALSE (since it was not added). If the path does not exist,
// it is added and the function returns TRUE (since the path was added).
func (he *historyEntries) pathAddedWithExistsCheck(path string) bool {
	if he.pathExists(path) {
		return false
	}
	he.add(path)
	return true
}

func (he *historyEntries) add(path string) {
	*he = append(*he, &path)
}
