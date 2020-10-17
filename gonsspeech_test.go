package gonsspeech

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNsSpeech_Init(t *testing.T) {
	err := NsSpeechInit()
	assert.NoError(t, err, "NsSpeechInit")
	err = NsSpeechInit()
	assert.EqualError(t, err, "Cannot initialize NsSpeechSynthesizer interface.")
	NsSpeechFree()
}

func TestNsSpeech_Speak(t *testing.T) {
	err := NsSpeechInit()
	assert.NoError(t, err, "NsSpeechInit")
	err := NsSpeechSpeak("test")
	assert.NoError(t, err, "NsSpeechSpeak")
	time.Sleep(2)
	NsSpeechFree()
}
