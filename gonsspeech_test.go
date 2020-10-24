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

func TestNsSpeech_SetRate(t *testing.T) {
	err := NsSpeechInit()
	assert.NoError(t, err, "NsSpeechInit")

	err = NsSpeechSpeak("rate 240 default")
	assert.NoError(t, err, "NsSpeechSpeak")
	time.Sleep(2 * time.Second)

err := NsSpeechSetRate(480)
	assert.NoError(t, err, "NsSpeechSetRate")
	err = NsSpeechSpeak("rate 480")
	assert.NoError(t, err, "NsSpeechSpeak")
	time.Sleep(2 * time.Second)

	NsSpeechFree()
}

func TestNsSpeech_SetRate_Error(t *testing.T) {
	err := NsSpeechSetRate(400)
	assert.EqualError(t, err, "NsSpeechSynthesizer interface has not been initialized.")
	err = NsSpeechInit()
	assert.NoError(t, err, "NsSpeechInit")

	// There's no highest value limit since apple developers guide doesn't mention it. But I'm pretty sure there is.
	err = NsSpeechSetRate(-1)
	assert.EqualError(t, err, "rate value is out of range")
	NsSpeechFree()
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
	assert.False(t, speaking, "NsSpeechIsSpeaking")

	NsSpeechFree()
}

func TestNsSpeech_IsSpeaking_Error(t *testing.T) {
	_, err := NsSpeechIsSpeaking()
	assert.EqualError(t, err, "NsSpeechSynthesizer interface has not been initialized.")
}
