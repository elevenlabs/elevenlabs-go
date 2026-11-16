package types

import (
	"testing"
)

func TestSubtitleOrderItemRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "subtitles",
  "media_ids": [
    "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  ],
  "source_language": "value",
  "destination_languages": [
    "value"
  ],
  "cue_options": {
    "min_duration_ms": 123,
    "max_duration_ms": 4000,
    "max_lines_per_cue": 3,
    "max_chars_per_line": 50,
    "max_chars_per_s": 30,
    "min_gap_between_cues_frames": 6
  },
  "sdh": true,
  "instructions": "value"
}`
	assertRoundTrip(t, wire, new(SubtitleOrderItemRequest))
}

func TestSubtitleOrderItemRequestRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "subtitles",
  "media_ids": [
    "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  ],
  "source_language": "value",
  "destination_languages": [
    "value"
  ],
  "cue_options": {
    "min_duration_ms": 123,
    "max_duration_ms": 4000,
    "max_lines_per_cue": 3,
    "max_chars_per_line": 50,
    "max_chars_per_s": 30,
    "min_gap_between_cues_frames": 6
  },
  "sdh": true,
  "instructions": "value"
}`
	assertRoundTrip(t, wire, new(SubtitleOrderItemRequest))
}
