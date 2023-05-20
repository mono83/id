package id

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	if x, err := Parse("13"); assert.NoError(t, err) {
		assert.Equal(t, ID(13), x)
	}
}
