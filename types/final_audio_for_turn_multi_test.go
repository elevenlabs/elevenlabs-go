package types

import (
	"testing"
)

func TestFinalAudioForTurnMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final_audio_for_turn": true,
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(FinalAudioForTurnMulti))
}
