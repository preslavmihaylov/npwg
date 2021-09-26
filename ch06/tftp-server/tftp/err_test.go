package tftp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalErrDatagram(t *testing.T) {
	d := NewErr(ErrNotFound, "file not found")

	expected := []byte{
		0x00, 0x05, // OpErr
		0x00, 0x01, // ErrNotFound
		'f', 'i', 'l', 'e', ' ', 'n', 'o', 't', ' ', 'f', 'o', 'u', 'n', 'd',
		0x00,
	}
	actual, err := d.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalErrDatagram(t *testing.T) {
	input := []byte{
		0x00, 0x05, // OpErr
		0x00, 0x01, // ErrNotFound
		'f', 'i', 'l', 'e', ' ', 'n', 'o', 't', ' ', 'f', 'o', 'u', 'n', 'd',
		0x00,
	}

	result, err := NewErrFromBinary(input)
	assert.NoError(t, err)
	assert.Equal(t, ErrNotFound, result.Code)
	assert.Equal(t, "file not found", result.Message)
}
