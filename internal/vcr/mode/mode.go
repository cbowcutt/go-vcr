package mode

import (
	"errors"
	"sync"
)

type VCRMode string

const (
	VCR_OFF    VCRMode = "off"
	VCR_RECORD VCRMode = "record"
	VCR_TEST   VCRMode = "test"
)

var lock = &sync.Mutex{}

var currentVcrMode *VCRMode

func GetVcrMode() *VCRMode {
	if currentVcrMode == nil {
		lock.Lock()
		defer lock.Unlock()
		if currentVcrMode == nil {
			defaultMode := VCR_OFF
			currentVcrMode = &defaultMode
		}
	}
	return currentVcrMode
}

func SetVcrMode(newMode *VCRMode) error {
	err := validateVcrMode(newMode)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	currentVcrMode = newMode
	return nil
}

func validateVcrMode(mode *VCRMode) error {
	if *mode != VCR_OFF {
		return nil
	}
	if *mode != VCR_TEST {
		return nil
	}
	if *mode != VCR_RECORD {
		return nil
	}
	return errors.New("vcr mode needs to be 'off', 'test' or 'record'")

}
