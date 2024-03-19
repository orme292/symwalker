package swalker

import (
	"fmt"
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

	for _, entry := range res {
		fmt.Printf("Path: %s\n", entry.Path)
	}

}

func checkTestResults(wr WalkerResults, t *testing.T) bool {
	lenExpected := 24
	if len(wr) != lenExpected {
		t.Logf("checkTestResults: wr length wrong. expected %d, got %d", lenExpected, len(wr))
		return false
	}

	expected := make(map[string]int)
	expected["/tests/users/andrew"] = 0
	expected["/tests/users/andrew/downloads"] = 0
	expected["/tests/users/andrew/documents"] = 0
	expected["/tests/users/andrew/pictures"] = 0
	expected["/tests/users/brian"] = 0
	expected["/tests/users/brian/downloads"] = 0
	expected["/tests/users/brian/documents"] = 0
	expected["/tests/users/brian/pictures"] = 0
	expected["/tests/users/carolyn"] = 0
	expected["/tests/users/carolyn/downloads"] = 0
	expected["/tests/users/carolyn/documents"] = 0
	expected["/tests/users/carolyn/pictures"] = 0
	expected["/tests/users/david"] = 0
	expected["/tests/users/david/downloads"] = 0
	expected["/tests/users/david/documents"] = 0
	expected["/tests/users/david/pictures"] = 0
	expected["/tests/users/erin"] = 0
	expected["/tests/users/erin/downloads"] = 0
	expected["/tests/users/erin/documents"] = 0
	expected["/tests/users/erin/pictures"] = 0
	expected["/tests/users/frank"] = 0
	expected["/tests/users/frank/downloads"] = 0
	expected["/tests/users/frank/documents"] = 0
	expected["/tests/users/frank/pictures"] = 0

	return true
}
