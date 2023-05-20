package id

import "strconv"

// Parse parses string into ID.
func Parse(s string) (ID, error) {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(x), nil
}
