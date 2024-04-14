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

func (he *historyEntries) add(path string) {
	*he = append(*he, &path)
}
