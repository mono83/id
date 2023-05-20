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

// ParseSlice parses multiple values
func ParseSlice(slice []string) (out []ID, err error) {
	if len(slice) == 0 {
		return
	}

	var u ID
	for _, s := range slice {
		u, err = Parse(s)
		if err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return
}
