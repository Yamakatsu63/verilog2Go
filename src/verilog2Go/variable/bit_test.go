package variable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	var bit Bit
	bit.InitializeBit(1)
	assert.Equal(t, bit.value, false)
	assert.Equal(t, bit.index, 1)
}

func TestSetValue(t *testing.T) {
	var bit Bit
	bit.InitializeBit(0)
	assert.Equal(t, bit.value, false)
	assert.Equal(t, bit.index, 0)
	bit.SetValue(true)
	assert.Equal(t, bit.value, true)
	assert.Equal(t, bit.index, 0)
}

func TestInitializeBitArray(t *testing.T) {
	var ba BitArray
	ba.InitializeBitArray(3)
	fmt.Println(ba.bits)
}
