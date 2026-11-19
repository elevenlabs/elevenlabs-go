package types

import (
	"testing"
)

func TestAudioRefChunkRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "song_id": "value",
  "range": {
    "start_ms": 123,
    "end_ms": 123
  }
}`
	assertRoundTrip(t, wire, new(AudioRefChunk))
}

func TestAudioRefChunkRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "song_id": "value",
  "range": {
    "start_ms": 123,
    "end_ms": 123
  }
}`
	assertRoundTrip(t, wire, new(AudioRefChunk))
}
