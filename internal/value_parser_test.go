package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBool(t *testing.T) {
	var result bool
	var err error

	result, err = parseBool("True")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)

	result, err = parseBool("False")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)

	result, err = parseBool("Yes")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)

	result, err = parseBool("NO")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, result)

	result, err = parseBool("unknown")
	assert.Error(t, err)
	assert.Equal(t, false, result)
}
