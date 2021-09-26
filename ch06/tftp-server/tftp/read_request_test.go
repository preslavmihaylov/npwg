package tftp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalReadRequest(t *testing.T) {
	rr := NewReadRequest("foo.txt", "octet")

	expected := []byte{
		0x00, 0x01, // RRQ Op Code
		'f', 'o', 'o', '.', 't', 'x', 't',
		0x00,
		'o', 'c', 't', 'e', 't',
		0x00,
	}
	actual, err := rr.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalReadRequest(t *testing.T) {
	payload := []byte{
		0x00, 0x01, // RRQ Op Code
		'f', 'o', 'o', '.', 't', 'x', 't',
		0x00,
		'o', 'c', 't', 'e', 't',
		0x00,
	}

	expected := NewReadRequest("foo.txt", "octet")
	actual, err := NewReadRequestFromBinary(payload)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalReadRequestErrorsOnUnsupportedMode(t *testing.T) {
	payload := []byte{
		0x00, 0x01, // RRQ Op Code
		'f', 'o', 'o', '.', 't', 'x', 't',
		0x00,
		'n', 'o', 'o',
		0x00,
	}

	_, err := NewReadRequestFromBinary(payload)
	assert.Error(t, err)
}
