package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traggo/server/test"
)

func TestReadLines(t *testing.T) {
	reader := &test.MockReader{
		Bytes:   []byte("foo\nbar\nbaz"),
		Pointer: 0,
		Err:     nil,
	}

	lines, err := ReadLines(reader)
	assert.NoError(t, err)

	assert.Equal(t, []string{"foo", "bar", "baz"}, lines)
}

func TestReadLines_Error(t *testing.T) {
	reader := &test.MockReader{
		Bytes:   []byte("foo\nbar\nbaz"),
		Pointer: 0,
		Err:     os.ErrNotExist,
	}

	lines, err := ReadLines(reader)
	assert.Error(t, err)

	assert.Equal(t, os.ErrNotExist, err)
	assert.Nil(t, lines)
}
