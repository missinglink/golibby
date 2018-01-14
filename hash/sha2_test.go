package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	{
		msg := ""
		digest := [32]byte{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a,
			0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64,
			0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}

		assert.Equal(t, digest, SHA256([]byte(msg)))
	}
	{
		msg := "foo"
		digest := [32]uint8{0x2c, 0x26, 0xb4, 0x6b, 0x68, 0xff, 0xc6, 0x8f, 0xf9, 0x9b, 0x45, 0x3c, 0x1d, 0x30, 0x41, 0x34, 0x13, 0x42, 0x2d, 0x70, 0x64, 0x83, 0xbf, 0xa0, 0xf9, 0x8a, 0x5e, 0x88, 0x62, 0x66, 0xe7, 0xae}

		assert.Equal(t, digest, SHA256([]byte(msg)))
	}
	{
		digest := [32]uint8{0x6c, 0x71, 0x46, 0x96, 0xd0, 0xdd, 0x37, 0x4b, 0x39, 0xe, 0xe, 0xe0, 0x61, 0xa1, 0xa, 0xe3, 0x5e, 0x9d, 0x2a, 0x2e, 0x26, 0xa7, 0x93, 0x71, 0xba, 0x15, 0xb2, 0xc, 0x90, 0x4e, 0xcd, 0x81}
		msg := `I like to think
(it has to be!)
of a cybernetic ecology
where we are free of our labors
and joined back to nature,
returned to our mammal
brothers and sisters,
and all watched over
by machines of loving grace.`

		assert.Equal(t, digest, SHA256([]byte(msg)))
	}
}
