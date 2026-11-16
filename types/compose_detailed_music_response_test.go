package types

import (
	"testing"
)

func TestComposeDetailedMusicResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "composition_plan": {
    "key": {}
  },
  "song_metadata": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(ComposeDetailedMusicResponse))
}
