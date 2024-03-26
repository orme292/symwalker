package swalker

import (
	"testing"
)

func TestSymWalker(t *testing.T) {
	conf := SymConf{
		StartPath:      "/tests/users",
		FollowSymlinks: true,
		Noisy:          false,
	}

	res, err := SymWalker(&conf)
	if err != nil {
		t.Errorf("SymWalker returned an error: %s", err.Error())
		t.Fail()
	}

	pass := checkTestResults(res, t)
	if !pass {
		t.Fail()
	}

}

func checkTestResults(wr WalkerResults, t *testing.T) bool {
	expected := make(map[string]bool)
	expected["/tests/users/andrew"] = false
	expected["/tests/users/andrew/downloads"] = false
	expected["/tests/users/andrew/documents"] = false
	expected["/tests/users/andrew/pictures"] = false
	expected["/tests/users/brian"] = false
	expected["/tests/users/brian/downloads"] = false
	expected["/tests/users/brian/documents"] = false
	expected["/tests/users/brian/pictures"] = false
	expected["/tests/users/carolyn"] = false
	expected["/tests/users/carolyn/downloads"] = false
	expected["/tests/users/carolyn/documents"] = false
	expected["/tests/users/carolyn/pictures"] = false
	expected["/tests/users/david"] = false
	expected["/tests/users/david/downloads"] = false
	expected["/tests/users/david/documents"] = false
	expected["/tests/users/david/pictures"] = false
	expected["/tests/users/erin"] = false
	expected["/tests/users/erin/downloads"] = false
	expected["/tests/users/erin/documents"] = false
	expected["/tests/users/erin/pictures"] = false
	expected["/tests/users/frank"] = false
	expected["/tests/users/frank/downloads"] = false
	expected["/tests/users/frank/documents"] = false
	expected["/tests/users/frank/pictures"] = false
	expected["/tests/users/andrew/others"] = false
	expected["/tests/users/andrew/others/more"] = false
	expected["/tests/users/andrew/others/more/directories"] = false
	expected["/tests/users/andrew/others/more/directories/to"] = false
	expected["/tests/users/andrew/others/more/directories/to/find"] = false

	var pass = true

	for _, entry := range wr {
		if _, ok := expected[entry.Path]; ok {
			if expected[entry.Path] == true {
				t.Errorf("Unexpected duplicate: %s", entry.Path)
				pass = false
			} else if expected[entry.Path] == false {
				expected[entry.Path] = true
			}
		} else {
			t.Errorf("Unexpected Path: %s", entry.Path)
		}
	}

	for index := range expected {
		if expected[index] == false {
			t.Errorf("Missing Path: %s", index)
			pass = false
		}
	}
	return pass
}
