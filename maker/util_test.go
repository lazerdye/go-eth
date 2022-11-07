package maker

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsciiToByte(t *testing.T) {
	bat_a_test := AsciiToByte("BAT-A")

	bat_a_str := hex.EncodeToString(bat_a_test[:])
	assert.Equal(t, "4241542d41000000000000000000000000000000000000000000000000000000", bat_a_str)

	long_test := AsciiToByte("This is really long, much too long for the hex string it should be just cut off.")

	long_str := hex.EncodeToString(long_test[:])
	assert.Equal(t, "54686973206973207265616c6c79206c6f6e672c206d75636820746f6f206c6f", long_str)
}
