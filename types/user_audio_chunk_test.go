package types

import (
	"testing"
)

func TestUserAudioChunkRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_audio_chunk": "value"
}`
	assertRoundTrip(t, wire, new(UserAudioChunk))
}
