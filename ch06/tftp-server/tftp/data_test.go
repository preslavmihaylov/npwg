package tftp

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalDataRequest(t *testing.T) {
	d := NewData(5, strings.NewReader("some data"))

	expected := []byte{
		0x00, 0x03, // OpData
		0x00, 0x05, // Block
		's', 'o', 'm', 'e', ' ', 'd', 'a', 't', 'a',
	}
	actual, err := d.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUnmarshalDataRequest(t *testing.T) {
	input := []byte{
		0x00, 0x03, // OpData
		0x00, 0x05, // Block
		's', 'o', 'm', 'e', ' ', 'd', 'a', 't', 'a',
	}

	result, err := NewDataFromBinary(input)
	assert.NoError(t, err)

	bs, err := ioutil.ReadAll(result.Payload)
	assert.NoError(t, err)
	assert.Equal(t, "some data", string(bs))
}
