package swalker

import (
	"testing"
)

func TestSymWalker(t *testing.T) {

	conf := NewSymConf(
		WithStartPath("/tests/users"),
		FollowsSymLinks(),
		//WithLogging(),
	)

	res, err := SymWalker(conf)
	if err != nil {
		t.Errorf("SymWalker returned an error: %s", err.Error())
		t.Fail()
	}

	pass := checkExpectedResults(getExpectedDirs(), res.Dirs, t)
	if !pass {
		t.Fail()
	}
	pass = checkExpectedResults(getExpectedFiles(), res.Files, t)
	if !pass {
		t.Fail()
	}

}

func checkExpectedResults(e map[string]bool, d DirEntries, t *testing.T) bool {

	var pass = true

	for _, entry := range d {
		if _, ok := e[entry.Path]; ok {
			if e[entry.Path] == true {
				t.Errorf("Unexpected duplicate: %s", entry.Path)
				pass = false
			} else if e[entry.Path] == false {
				t.Logf("OK: %s", entry.Path)
				e[entry.Path] = true
			}
		} else {
			t.Errorf("Unexpected Path: %s", entry.Path)
		}
	}

	for index := range e {
		if e[index] == false {
			t.Errorf("Missing Path: %s", index)
			pass = false
		}
	}

	if pass == true {
		t.Logf("TEST PASSED")
	} else {
		t.Logf("TEST FAILED")
	}

	return pass
}

func getExpectedDirs() map[string]bool {
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
	return expected
}

func getExpectedFiles() map[string]bool {
	expected := make(map[string]bool)
	expected["/tests/users/andrew/documents/1-report.doc"] = false
	expected["/tests/users/andrew/documents/2-report.doc"] = false
	expected["/tests/users/andrew/documents/3-report.doc"] = false
	expected["/tests/users/andrew/documents/4-report.doc"] = false
	expected["/tests/users/andrew/documents/5-report.doc"] = false
	expected["/tests/users/andrew/documents/6-report.doc"] = false
	expected["/tests/users/andrew/downloads/a.part"] = false
	expected["/tests/users/andrew/downloads/b.part"] = false
	expected["/tests/users/andrew/downloads/c.part"] = false
	expected["/tests/users/andrew/downloads/d.part"] = false
	expected["/tests/users/andrew/downloads/e.part"] = false
	expected["/tests/users/andrew/downloads/f.part"] = false
	expected["/tests/users/andrew/downloads/g.part"] = false
	expected["/tests/users/andrew/linkedfile"] = false
	expected["/tests/users/andrew/pictures/1.jpg"] = false
	expected["/tests/users/andrew/pictures/2.jpg"] = false
	expected["/tests/users/andrew/pictures/3.jpg"] = false
	expected["/tests/users/andrew/pictures/4.jpg"] = false
	expected["/tests/users/andrew/pictures/5.jpg"] = false
	expected["/tests/users/brian/documents/1-report.doc"] = false
	expected["/tests/users/brian/documents/2-report.doc"] = false
	expected["/tests/users/brian/documents/3-report.doc"] = false
	expected["/tests/users/brian/documents/4-report.doc"] = false
	expected["/tests/users/brian/documents/5-report.doc"] = false
	expected["/tests/users/brian/documents/6-report.doc"] = false
	expected["/tests/users/brian/downloads/a.part"] = false
	expected["/tests/users/brian/downloads/b.part"] = false
	expected["/tests/users/brian/downloads/c.part"] = false
	expected["/tests/users/brian/downloads/d.part"] = false
	expected["/tests/users/brian/downloads/e.part"] = false
	expected["/tests/users/brian/downloads/f.part"] = false
	expected["/tests/users/brian/downloads/g.part"] = false
	expected["/tests/users/brian/pictures/1.jpg"] = false
	expected["/tests/users/brian/pictures/2.jpg"] = false
	expected["/tests/users/brian/pictures/3.jpg"] = false
	expected["/tests/users/brian/pictures/4.jpg"] = false
	expected["/tests/users/brian/pictures/5.jpg"] = false
	expected["/tests/users/carolyn/documents/1-report.doc"] = false
	expected["/tests/users/carolyn/documents/2-report.doc"] = false
	expected["/tests/users/carolyn/documents/3-report.doc"] = false
	expected["/tests/users/carolyn/documents/4-report.doc"] = false
	expected["/tests/users/carolyn/documents/5-report.doc"] = false
	expected["/tests/users/carolyn/documents/6-report.doc"] = false
	expected["/tests/users/carolyn/downloads/a.part"] = false
	expected["/tests/users/carolyn/downloads/b.part"] = false
	expected["/tests/users/carolyn/downloads/c.part"] = false
	expected["/tests/users/carolyn/downloads/d.part"] = false
	expected["/tests/users/carolyn/downloads/e.part"] = false
	expected["/tests/users/carolyn/downloads/f.part"] = false
	expected["/tests/users/carolyn/downloads/g.part"] = false
	expected["/tests/users/carolyn/pictures/1.jpg"] = false
	expected["/tests/users/carolyn/pictures/2.jpg"] = false
	expected["/tests/users/carolyn/pictures/3.jpg"] = false
	expected["/tests/users/carolyn/pictures/4.jpg"] = false
	expected["/tests/users/carolyn/pictures/5.jpg"] = false
	expected["/tests/users/david/documents/1-report.doc"] = false
	expected["/tests/users/david/documents/2-report.doc"] = false
	expected["/tests/users/david/documents/3-report.doc"] = false
	expected["/tests/users/david/documents/4-report.doc"] = false
	expected["/tests/users/david/documents/5-report.doc"] = false
	expected["/tests/users/david/documents/6-report.doc"] = false
	expected["/tests/users/david/downloads/a.part"] = false
	expected["/tests/users/david/downloads/b.part"] = false
	expected["/tests/users/david/downloads/c.part"] = false
	expected["/tests/users/david/downloads/d.part"] = false
	expected["/tests/users/david/downloads/e.part"] = false
	expected["/tests/users/david/downloads/f.part"] = false
	expected["/tests/users/david/downloads/g.part"] = false
	expected["/tests/users/david/pictures/1.jpg"] = false
	expected["/tests/users/david/pictures/2.jpg"] = false
	expected["/tests/users/david/pictures/3.jpg"] = false
	expected["/tests/users/david/pictures/4.jpg"] = false
	expected["/tests/users/david/pictures/5.jpg"] = false
	expected["/tests/users/david/documents/rogue/linkedfile"] = false
	expected["/tests/users/david/documents/rogue/pictures/1.jpg"] = false
	expected["/tests/users/david/documents/rogue/pictures/2.jpg"] = false
	expected["/tests/users/david/documents/rogue/pictures/3.jpg"] = false
	expected["/tests/users/david/documents/rogue/pictures/4.jpg"] = false
	expected["/tests/users/david/documents/rogue/pictures/5.jpg"] = false
	expected["/tests/users/docs1/more/img-1.jpg"] = false
	expected["/tests/users/docs1/more/img-2.jpg"] = false
	expected["/tests/users/docs1/more/img-3.jpg"] = false
	expected["/tests/users/docs1/more/img-4.jpg"] = false
	expected["/tests/users/docs1/more/img-5.jpg"] = false
	expected["/tests/users/docs1/more/img-6.jpg"] = false
	expected["/tests/users/docs2/more/img-1.jpg"] = false
	expected["/tests/users/docs2/more/img-2.jpg"] = false
	expected["/tests/users/docs2/more/img-3.jpg"] = false
	expected["/tests/users/docs2/more/img-4.jpg"] = false
	expected["/tests/users/docs2/more/img-5.jpg"] = false
	expected["/tests/users/docs2/more/img-6.jpg"] = false
	expected["/tests/users/docs1/doc-1.pdf"] = false
	expected["/tests/users/docs1/doc-2.pdf"] = false
	expected["/tests/users/docs1/doc-3.pdf"] = false
	expected["/tests/users/docs1/doc-4.pdf"] = false
	expected["/tests/users/docs1/doc-5.pdf"] = false
	expected["/tests/users/docs1/doc-6.pdf"] = false
	expected["/tests/users/docs2/doc-1.pdf"] = false
	expected["/tests/users/docs2/doc-2.pdf"] = false
	expected["/tests/users/docs2/doc-3.pdf"] = false
	expected["/tests/users/docs2/doc-4.pdf"] = false
	expected["/tests/users/docs2/doc-5.pdf"] = false
	expected["/tests/users/docs2/doc-6.pdf"] = false
	expected["/tests/users/erin/documents/1-report.doc"] = false
	expected["/tests/users/erin/documents/2-report.doc"] = false
	expected["/tests/users/erin/documents/3-report.doc"] = false
	expected["/tests/users/erin/documents/4-report.doc"] = false
	expected["/tests/users/erin/documents/5-report.doc"] = false
	expected["/tests/users/erin/documents/6-report.doc"] = false
	expected["/tests/users/erin/downloads/a.part"] = false
	expected["/tests/users/erin/downloads/b.part"] = false
	expected["/tests/users/erin/downloads/c.part"] = false
	expected["/tests/users/erin/downloads/d.part"] = false
	expected["/tests/users/erin/downloads/e.part"] = false
	expected["/tests/users/erin/downloads/f.part"] = false
	expected["/tests/users/erin/downloads/g.part"] = false
	expected["/tests/users/erin/pictures/1.jpg"] = false
	expected["/tests/users/erin/pictures/2.jpg"] = false
	expected["/tests/users/erin/pictures/3.jpg"] = false
	expected["/tests/users/erin/pictures/4.jpg"] = false
	expected["/tests/users/erin/pictures/5.jpg"] = false
	expected["/tests/users/frank/documents/1-report.doc"] = false
	expected["/tests/users/frank/documents/2-report.doc"] = false
	expected["/tests/users/frank/documents/3-report.doc"] = false
	expected["/tests/users/frank/documents/4-report.doc"] = false
	expected["/tests/users/frank/documents/5-report.doc"] = false
	expected["/tests/users/frank/documents/6-report.doc"] = false
	expected["/tests/users/frank/downloads/a.part"] = false
	expected["/tests/users/frank/downloads/b.part"] = false
	expected["/tests/users/frank/downloads/c.part"] = false
	expected["/tests/users/frank/downloads/d.part"] = false
	expected["/tests/users/frank/downloads/e.part"] = false
	expected["/tests/users/frank/downloads/f.part"] = false
	expected["/tests/users/frank/downloads/g.part"] = false
	expected["/tests/users/frank/pictures/1.jpg"] = false
	expected["/tests/users/frank/pictures/2.jpg"] = false
	expected["/tests/users/frank/pictures/3.jpg"] = false
	expected["/tests/users/frank/pictures/4.jpg"] = false
	expected["/tests/users/frank/pictures/5.jpg"] = false
	expected["/tests/users/david/documents/rogue/documents/1-report.doc"] = false
	expected["/tests/users/david/documents/rogue/documents/2-report.doc"] = false
	expected["/tests/users/david/documents/rogue/documents/3-report.doc"] = false
	expected["/tests/users/david/documents/rogue/documents/4-report.doc"] = false
	expected["/tests/users/david/documents/rogue/documents/5-report.doc"] = false
	expected["/tests/users/david/documents/rogue/documents/6-report.doc"] = false
	expected["/tests/users/david/documents/rogue/downloads/a.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/b.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/c.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/d.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/e.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/f.part"] = false
	expected["/tests/users/david/documents/rogue/downloads/g.part"] = false
	return expected
}
