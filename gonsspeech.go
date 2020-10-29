package gonsspeech

/*
#cgo CFLAGS: -x objective-c

#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

NSSpeechSynthesizer *speechSynth = NULL;

int NsSpeechInit(){
if(speechSynth) return 0;
speechSynth=[[NSSpeechSynthesizer alloc] init];
[speechSynth setRate:240];
[speechSynth setVolume:0.9];
return 1;
}

int NsSpeechFree(){
if(!speechSynth) return 0;
speechSynth=NULL;
return 1;
}

int NsSpeechSpeak(char *str){
if(!speechSynth) return 0;
NSString *nsstr = [NSString stringWithCString: str encoding:NSUTF8StringEncoding];
[speechSynth startSpeakingString:nsstr];
usleep(5000);
    return 1;
}

int NsSpeechStop(){
if(!speechSynth) return 0;
[speechSynth stopSpeaking];
    return 1;
}

int NsSpeechSetRate(float rate){
if(!speechSynth) return 0;
if(rate<=0) return -1;
[speechSynth setRate:rate];
return 1;
}

int NsSpeechIsSpeaking(){
if(!speechSynth) return 0;
return [speechSynth isSpeaking] == YES ? 1 : 2;
}


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

func NsSpeechStop() error {
	ret := C.NsSpeechStop()
	if ret == 0 {
		return errors.New("NsSpeechSynthesizer interface has not been initialized.")
	}
	return nil
}

func NsSpeechSetRate(rate float64) error {
	ret := C.NsSpeechSetRate(C.float(rate))
	if ret == 0 {
		return errors.New("NsSpeechSynthesizer interface has not been initialized.")
	}
	if ret == -1 {
		return errors.New("rate value is out of range")
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
