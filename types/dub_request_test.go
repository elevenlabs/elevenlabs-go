package types

import (
	"testing"
)

func TestDubRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "source_url": "value",
  "source_lang": "value",
  "target_lang": "value",
  "target_accent": "value",
  "num_speakers": 123,
  "watermark": true,
  "start_time": 123,
  "end_time": 123,
  "highest_resolution": true,
  "drop_background_audio": true,
  "use_profanity_filter": true,
  "dubbing_studio": true,
  "disable_voice_cloning": true,
  "mode": "automatic",
  "csv_fps": 9.99
}`
	assertRoundTrip(t, wire, new(DubRequest))
}
