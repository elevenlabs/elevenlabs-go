package types

import (
	"testing"
)

func TestTranslateInputAudioChunkRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "input_audio_chunk",
  "audio_base_64": "value"
}`
	assertRoundTrip(t, wire, new(TranslateInputAudioChunk))
}
