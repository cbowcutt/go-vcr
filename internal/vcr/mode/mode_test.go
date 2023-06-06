package mode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetVcrMode_default(t *testing.T) {
	mode := GetVcrMode()
	assert.Equal(t, VCR_OFF, *mode)
}

func Test_SetVcrMode_VCR_RECORD(t *testing.T) {
	newMode := VCR_RECORD
	err := SetVcrMode(&newMode)
	assert.Nil(t, err)
	actualMode := GetVcrMode()
	assert.Equal(t, newMode, *actualMode )
}

func Test_SetVcrMode_VCR_OFF(t *testing.T) {
	newMode := VCR_OFF
	err := SetVcrMode(&newMode)
	assert.Nil(t, err)
	actualMode := GetVcrMode()
	assert.Equal(t, newMode, *actualMode )
}

func Test_SetVcrMode_VCR_TEST(t *testing.T) {
	newMode := VCR_TEST
	err := SetVcrMode(&newMode)
	assert.Nil(t, err)
	actualMode := GetVcrMode()
	assert.Equal(t, newMode, *actualMode )
}

func Test_SetVcrMode_Error(t *testing.T) {
	newMode := VCRMode("hello")
	err := SetVcrMode(&newMode)
	assert.Nil(t, err)
	actualMode := GetVcrMode()
	assert.Equal(t, newMode, *actualMode )
}
