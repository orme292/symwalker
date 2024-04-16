package swalker

import (
	"testing"
)

func TestSymWalker(t *testing.T) {

	conf := NewSymConf(
		WithStartPath("/tests/users"),
		FollowsSymLinks(),
		WithLogging(),
	)

	res, err := SymWalker(conf)
	if err != nil {
		t.Errorf("SymWalker returned an error: %s", err.Error())
		t.Fail()
	}

	pass := checkTestResults(res, t)
	if !pass {
		t.Fail()
	}

}

func checkTestResults(r Results, t *testing.T) bool {

	expected := make(map[string]bool)
	expected["/tests/users"] = false
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
	expected["/tests/users/david/documents/rogue"] = false
	expected["/tests/users/david/documents/rogue/documents"] = false
	expected["/tests/users/david/documents/rogue/downloads"] = false
	expected["/tests/users/david/documents/rogue/others"] = false
	expected["/tests/users/david/documents/rogue/others/more"] = false
	expected["/tests/users/david/documents/rogue/others/more/directories"] = false
	expected["/tests/users/david/documents/rogue/others/more/directories/to"] = false
	expected["/tests/users/david/documents/rogue/others/more/directories/to/find"] = false
	expected["/tests/users/david/documents/rogue/pictures"] = false
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
	expected["/tests/users/docs1"] = false
	expected["/tests/users/docs2"] = false
	expected["/tests/users/docs1/more"] = false
	expected["/tests/users/docs2/more"] = false

	var pass = true

	for _, dir := range r.Dirs {
		if _, ok := expected[dir.Path]; ok {
			if expected[dir.Path] == true {
				t.Errorf("Unexpected duplicate: %s", dir.Path)
				pass = false
			} else if expected[dir.Path] == false {
				t.Logf("OK: %s", dir.Path)
				expected[dir.Path] = true
			}
		} else {
			t.Errorf("Unexpected Path: %s", dir.Path)
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
