package symwalker

import (
	"fmt"
	"log"
	"testing"
)

// todo: rewrite with new history type in mind
func TestHistoryMethods(t *testing.T) {

	paths := []string{
		"/users",
		"/users",
		"/users",
		"/users/andrew",
		"/users/andrew",
		"/users/andrew/documents",
		"/users/andrew/downloads",
		"/users/andrew/",
		"users/andrew",
		// There's no test for a relative path, like "./ssh", since the result would be a
		// random temporary path as the current working directory
	}

	wants := map[string]int{
		"/users":                  3,
		"/users/andrew":           4,
		"/users/andrew/documents": 1,
		"/users/andrew/downloads": 1,
	}

	hist := make(history)
	for _, path := range paths {
		hist.add(path)
	}

	fmt.Printf("%v+\n\t", hist)

	for arg, want := range wants {
		got := hist.count(arg)
		if got != want {
			t.Errorf("!FAIL: %q GOT %d, WANT %d", arg, want, got)
		} else {
			log.Printf("PASS: %q GOT %d", arg, got)
		}
	}
}
