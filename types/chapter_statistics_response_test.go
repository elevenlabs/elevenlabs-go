package types

import (
	"testing"
)

func TestChapterStatisticsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "characters_unconverted": 123,
  "characters_converted": 123,
  "paragraphs_converted": 123,
  "paragraphs_unconverted": 123,
  "credits_needed_to_convert": 123,
  "voice_statistics": [
    {
      "project_voice_ref_id": "value",
      "characters_unconverted": 123,
      "characters_converted": 123,
      "credits_needed_to_convert": 123,
      "voice_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ChapterStatisticsResponse))
}

func TestChapterStatisticsResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "characters_unconverted": 123,
  "characters_converted": 123,
  "paragraphs_converted": 123,
  "paragraphs_unconverted": 123,
  "credits_needed_to_convert": 123,
  "voice_statistics": [
    {
      "project_voice_ref_id": "value",
      "characters_unconverted": 123,
      "characters_converted": 123,
      "credits_needed_to_convert": 123,
      "voice_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ChapterStatisticsResponse))
}
