package gonsspeech
import (
"testing"
"github.com/stretchr/testify/assert"
)

func TestNsSpeech_Init(t *testing.T) {
err := NsSpeechInit()
assert.NoError(t, err, "NsSpeechInit")
err = NsSpeechInit()
assert.EqualError(t, err, "Cannot initialize NsSpeechSynthesizer interface.")
NsSpeechFree()
}
