package old

import (
	"errors"
	"fmt"
)

type WalkOpErr error

var (
	ErrConfGlobalMalformed = errors.New("GlobPattern is malformed")

	ErrStartPathNotReadable = errors.New("StartPath could not be read")
	ErrStartPathIsSymlink   = errors.New("StartPath cannot be a symlink")
	ErrStartPathIsNotDir    = errors.New("StartPath is not a directory")

	ErrWalkFailedGeneral        = errors.New("error while walking directory")
	ErrWalkPathNotReadable      = errors.New("error reading from path")
	ErrWalkCouldNotListFiles    = errors.New("error listing files under path")
	ErrWalkPotentialSymlinkLoop = errors.New("potential symlink loop detected")

	OpsMaxDepthReached = errors.New("max depth has been reached")
)

// NewError creates a new error by wrapping an existing error and appending a custom message.
// Errors generated with NewError using Error Constants defined above will be comparable with
// errors.Is.
//
// Parameters:
//   - errString: the error string to be wrapped
//   - msg: the custom message to be appended
//
// Returns:
//   - error: the newly created error with the wrapped error and custom message
func NewError(e WalkOpErr, msg string) WalkOpErr {
	return fmt.Errorf("%w: %s", e, msg)
}
