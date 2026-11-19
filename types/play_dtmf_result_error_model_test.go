package types

import (
	"testing"
)

func TestPlayDtmfResultErrorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "play_dtmf_error",
  "status": "error",
  "error": "value",
  "details": "value"
}`
	assertRoundTrip(t, wire, new(PlayDtmfResultErrorModel))
}

func TestPlayDtmfResultErrorModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "play_dtmf_error",
  "status": "error",
  "error": "value",
  "details": "value"
}`
	assertRoundTrip(t, wire, new(PlayDtmfResultErrorModel))
}
