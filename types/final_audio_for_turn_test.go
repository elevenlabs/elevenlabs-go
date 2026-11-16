package types

import (
	"testing"
)

func TestFinalAudioForTurnRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final_audio_for_turn": true
}`
	assertRoundTrip(t, wire, new(FinalAudioForTurn))
}
