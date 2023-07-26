package mode

import (
	"errors"
	"sync"
)

type VCRMode string

const (
	OFF      VCRMode = "off"
	RECORD   VCRMode = "record"
	PLAYBACK VCRMode = "playback"
)

var lock = &sync.Mutex{}

var currentVcrMode *VCRMode

func GetVcrMode() *VCRMode {
	if currentVcrMode == nil {
		lock.Lock()
		defer lock.Unlock()
		if currentVcrMode == nil {
			defaultMode := OFF
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
	if *mode != OFF {
		return nil
	}
	if *mode != PLAYBACK {
		return nil
	}
	if *mode != RECORD {
		return nil
	}
	return errors.New("vcr mode needs to be 'off', 'test' or 'record'")

}
