package symwalker

import (
	"os"
)

// WalkerResult represents the result of walking a directory. It contains
// the path of the file or directory, its file mode, and any error encountered
// when trying to read file info from the path.
type WalkerResult struct {
	Path   string
	IsType os.FileMode
	Error  WalkOpErr
}

// Results is a type representing a collection of WalkerResult pointers.
// It is used to store the results of a walker operation.
type Results []*WalkerResult

// add appends a new WalkerResult to the Results slice
// Parameter p is the path of the result
// Parameter tp is the file mode of the result
// Parameter e is the WalkOpErr error of the result
func (r *Results) add(p string, tp os.FileMode, e WalkOpErr) {
	*r = append(*r, &WalkerResult{
		Path:   p,
		IsType: tp,
		Error:  e,
	})
}

// GetErrored returns a slice of WalkerResult with non-nil Error fields.
// It iterates over the Results slice and appends any WalkerResult that
// has a non-nil Error to the newResults slice.
// Returns []*WalkerResult
// Note: This method does not modify the original Results slice.
func (r *Results) GetErrored() []*WalkerResult {
	newResults := make([]*WalkerResult, 0)
	for _, result := range *r {
		if result.Error != nil {
			newResults = append(newResults, result)
		}
	}
	return newResults
}

// Get returns a slice of WalkerResult pointers with no errors.
// It iterates through the Results slice and adds the results with no errors to
// the newResults slice.
// Returns []*WalkerResult
// Note: This method does not modify the original Results slice.
func (r *Results) Get() []*WalkerResult {
	newResults := make([]*WalkerResult, 0)
	for _, result := range *r {
		if result.Error == nil {
			newResults = append(newResults, result)
		}
	}
	return newResults
}
