package id

import (
	"bytes"
	"strconv"
)

// ID is a type to represent incremental identifier
// in database
type ID uint64

func (i ID) Uint64() uint64 { return uint64(i) }
func (i ID) String() string { return strconv.FormatUint(i.Uint64(), 10) }

// SQLIn constructs part of SQL query with placeholders values
func SQLIn(ids ...ID) (string, []interface{}) {
	if l := len(ids); l == 0 {
		return "", nil
	} else if l == 1 {
		return "=?", []interface{}{ids[0].Uint64()}
	}

	var out []interface{}
	buf := bytes.NewBufferString("IN (")
	for i, j := range ids {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('?')
		out = append(out, j.Uint64())
	}
	buf.WriteRune(')')
	return buf.String(), out
}
