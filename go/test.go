package main

/*
#cgo LDFLAGS: -L. -lnsspeech
#include "splib.h"
*/
import "C"
import "fmt"
import "time"

func main() {
	ret := C.NsSpeechInit()
	C.NsSpeechSpeak(C.CString("this is test test test"))
	time.Sleep(time.Second*5)
	ret = C.NsSpeechFree()
}
