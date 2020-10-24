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
	err = NsSpeechSpeak("test")
	assert.NoError(t, err, "NsSpeechSpeak")
	time.Sleep(2 * time.Second)
	NsSpeechFree()
}

func TestNsSpeech_Speak_Error(t *testing.T) {
	err := NsSpeechSpeak("test")
	assert.EqualError(t, err, "NsSpeechSynthesizer interface has not been initialized.")
}

func TestNsSpeech_IsSpeaking(t *testing.T) {
	err := NsSpeechInit()
	assert.NoError(t, err, "NsSpeechInit")

	err = NsSpeechSpeak("This is a long long long long long long long long long long long loooooooooooooooooooong text")
	assert.NoError(t, err, "NsSpeechSpeak")
	time.Sleep(2 * time.Second)

	speaking, err := NsSpeechIsSpeaking()
	assert.NoError(t, err, "NsSpeechIsSpeaking")
	assert.True(t, speaking, "NsSpeechIsSpeaking")

	err = NsSpeechStop()
	time.Sleep(time.Second)
	assert.NoError(t, err, "NsSpeechStop")
	speaking, err = NsSpeechIsSpeaking()
	assert.NoError(t, err, "NsSpeechIsSpeaking")
	assert.False(t,speaking,"NsSpeechIsSpeaking")

	NsSpeechFree()
}

func TestNsSpeech_IsSpeaking_Error(t *testing.T) {
	_, err := NsSpeechIsSpeaking()
	assert.EqualError(t, err, "NsSpeechSynthesizer interface has not been initialized.")
}
