package types

import (
	"testing"
)

func TestPlayDtmfToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "play_keypad_touch_tone",
  "use_out_of_band_dtmf": true,
  "suppress_turn_after_dtmf": true
}`
	assertRoundTrip(t, wire, new(PlayDtmfToolConfig))
}
