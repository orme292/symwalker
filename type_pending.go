package swalker

import (
	"log"
)

type PendingEntries []*DirEntry

func (pe *PendingEntries) pathExists(path string) bool {
	for _, entry := range *pe {
		if entry.Path == path {
			return true
		}
	}
	return false
}

func (pe *PendingEntries) add(path string) {
	*pe = append(*pe, &DirEntry{Path: path})
}

func (pe *PendingEntries) markByPath(path string) {
	for index := range *pe {
		if (*pe)[index].Path == path {
			log.Println("MARK: ", path)
			(*pe)[index].Marked = true
		}
	}
}

func (pe *PendingEntries) markByIndex(index int) {
	if index < 0 || index < len(*pe) {
		log.Println("MARK: ", (*pe)[index].Path)
		(*pe)[index].Marked = true
	}
}

func (pe *PendingEntries) unmarkedEntriesExist() bool {
	for index := range *pe {
		if (*pe)[index].Marked == false {
			return true
		}
	}
	return false
}
