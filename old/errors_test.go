package old

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	errMap := map[error]error{
		ErrStartPathNotReadable: ErrStartPathNotReadable,
		ErrStartPathIsSymlink:   ErrStartPathIsSymlink,
		ErrStartPathIsNotDir:    ErrStartPathIsNotDir,
		ErrWalkFailedGeneral:    ErrWalkFailedGeneral,
		ErrWalkPathNotReadable:  ErrWalkPathNotReadable,
	}

	for arg, want := range errMap {

		got := NewError(arg, "test message")
		if !errors.Is(got, want) {
			t.Errorf("FAIL: GOT: %v, WANTED ERROR: %v", got, want)
		} else {
			t.Logf("PASS: %v", arg)
		}
	}
}
