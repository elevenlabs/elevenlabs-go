package types

import (
	"testing"
)

func TestAddVoiceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(AddVoiceResponseModel))
}

func TestAddVoiceResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(AddVoiceResponseModel))
}

func TestAddVoiceResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(AddVoiceResponseModel))
}
