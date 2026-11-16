package types

import (
	"testing"
)

func TestPlayDtmfResultSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "play_dtmf_success",
  "status": "success",
  "dtmf_tones": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(PlayDtmfResultSuccessModel))
}

func TestPlayDtmfResultSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "play_dtmf_success",
  "status": "success",
  "dtmf_tones": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(PlayDtmfResultSuccessModel))
}
