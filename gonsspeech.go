package gonsspeech

/*
#cgo LDFLAGS: -L./objective-c -lnsspeech
#include "objective-c/splib.h"
*/
import "C"
import (
	"errors"
)

func NsSpeechInit() error {
	ret := C.NsSpeechInit()
	if ret == 0 {
		return errors.New("Cannot initialize NsSpeechSynthesizer interface.")
	}
	return nil
}

func NsSpeechSpeak(text string) error {
	ret := C.NsSpeechSpeak(C.CString(text))
	if ret == 0 {
		return errors.New("NsSpeechSynthesizer interface has not been initialized.")
	}
	return nil
}

func NsSpeechIsSpeaking() (bool, error) {
	switch C.NsSpeechIsSpeaking() {
	case 0:
		return false, errors.New("NsSpeechSynthesizer interface has not been initialized.")
	case 1:
		return true, nil
	case 2:
		return false, nil
	}
	return false, errors.New("Unrecognized error.")
}

func NsSpeechFree() {
	C.NsSpeechFree()
}
