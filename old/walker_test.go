package old

import (
	"errors"
	"fmt"
	"testing"
)

// TestSymWalker_StartPathErrors tests the behavior of the SymWalker function when encountering start path errors.
// Arguments:
// - `/Users/aorme/test`: StartPath with no errors.
// - `/Users/aorme`: StartPath with no errors.
// - `/Users/aorme/test/working/level4`: StartPath is a symlink.
// - `/Users/aorme/test/working/file1.md`: StartPath is a symlink.
// - `Users/aorme/test/working/data_file_1.dat`: StartPath is not a directory.
// For each argument, it initializes a `Conf` struct and invokes the SymWalker function.
// It compares the resulting error with the expected error and logs the results.
// If the resulting error is nil and the expected error is nil, the test has passed.
// If the resulting error and the expected error match, it prints the error details.
func TestSymWalker_StartPathErrors(t *testing.T) {
	// TODO: Create a file structure that is testable.
	args := make(map[string]error)
	args["/tests/start/dir_that_does_not_exist"] = ErrStartPathNotReadable
	args["/tests/symlink/dirs/link_to_readable_dir_1"] = ErrStartPathIsSymlink
	args["/tests/start/readable_1.file"] = ErrStartPathIsNotDir

	for path, want := range args {
		fmt.Println("Testing: ", path)
		conf := &Conf{
			StartPath:      path,
			Depth:          DepthInfinite,
			FollowSymLinks: true,
			TargetType:     TargetDir,
			GlobPattern:    GlobMatchAll,
		}
		_, got := SymWalker(conf)
		if got != nil {
			if errors.Is(got, want) == false {
				t.Errorf("!FAIL:\n\tWANT: %s\n\tGOT: %v", want.Error(), got.Error())
			} else {
				t.Logf("PASS: WANT: %v (GOT: %v)", want.Error(), got.Error())
			}
		} else {
			t.Errorf("!FAIL:\n\tWANT %s\n\tGOT: nil", want.Error())
		}
	}
}

func TestSymWalker_GlobCheck(t *testing.T) {
	path := "/tests/start"
	conf := &Conf{
		StartPath:      path,
		Depth:          DepthInfinite,
		FollowSymLinks: true,
		TargetType:     TargetRegular,
		GlobPattern:    GlobMatchAll,
	}
	_, got := SymWalker(conf)
	if got != nil {
		t.Errorf("FAIL!:\n\tWANT: nil\n\tGOT: %v", got.Error())
	} else {
		t.Log("PASS: WANT: nil (GOT: nil)")
	}

	conf.GlobPattern = "[-]"
	want := ErrConfGlobalMalformed
	_, got = SymWalker(conf)
	if !errors.Is(got, want) && got != nil {
		t.Errorf("FAIL!:\n\tWANT: %T\n\tGOT: %T", want, got)
	} else if errors.Is(got, want) {
		t.Logf("PASS: WANT: %v (GOT: %v)", want.Error(), got.Error())
	} else if got == nil {
		t.Errorf("FAIL: WANT: %v (GOT: nil)", want.Error())
	}
}
