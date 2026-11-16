package types

import (
	"testing"
)

func TestSpeechEngineSummaryResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "speech_engine_id": "value",
  "name": "value",
  "voice_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "tags": [
    "value"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`
	assertRoundTrip(t, wire, new(SpeechEngineSummaryResponse))
}
