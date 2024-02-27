package symwalker

// history is a type that represents a mapping of string keys to unsigned 8-bit integer values.
// It is used to track the count of occurrences of each string in history.
// Example usage:
//
//	hist := make(history)
//	hist.add("path/to/file")
//	count := hist.count("path/to/file")
type history map[string]uint8

// add increments the count for the given path in the history.
func (h *history) add(path string) {
	path = format(path)
	(*h)[path]++
}

// count returns the count value for the given path in the history.
func (h *history) count(path string) int {
	return int((*h)[path])
}

type histEnt struct {
	path     string
	count    uint8
	referrer []string //referrer is a symlink that targeted the path in the entry.
}

type historyA []*histEnt

func (hA *historyA) add(path, referrer string) error {
	path = format(path)
	for _, entry := range *hA {
		if entry.path == path {
			entry.count++
			if referrer != emptyString {
				// todo: must iterate through referrers and check for repeats.
				entry.referrer = append(entry.referrer, format(referrer))
			}
		}
		return nil
	}
	*hA = append(*hA, &histEnt{
		path:     path,
		count:    1,
		referrer: []string{referrer},
	})
	return nil
}
