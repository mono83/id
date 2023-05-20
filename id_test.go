package id

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQLIn(t *testing.T) {
	query, args := SQLIn()
	assert.Equal(t, "", query)
	assert.Len(t, args, 0)

	query, args = SQLIn(44)
	assert.Equal(t, "=?", query)
	assert.Equal(t, []interface{}{uint64(44)}, args)

	query, args = SQLIn(2, 65)
	assert.Equal(t, "IN (?,?)", query)
	assert.Equal(t, []interface{}{uint64(2), uint64(65)}, args)

	query, args = SQLIn(9, 8, 7, 6, 5)
	assert.Equal(t, "IN (?,?,?,?,?)", query)
	assert.Equal(t, []interface{}{uint64(9), uint64(8), uint64(7), uint64(6), uint64(5)}, args)
}
