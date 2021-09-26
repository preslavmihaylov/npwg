package tftp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalAckDatagram(t *testing.T) {
	d := NewAck(5)
	expected := []byte{
		0x00, 0x04, // OpAck
		0x00, 0x05, // Block #5
	}
	actual, err := d.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalAckDatagram(t *testing.T) {
	input := []byte{
		0x00, 0x04, // OpAck
		0x00, 0x05, // Block #5
	}

	result, err := NewAckFromBinary(input)
	assert.NoError(t, err)
	assert.Equal(t, uint16(5), result.Block)
}
