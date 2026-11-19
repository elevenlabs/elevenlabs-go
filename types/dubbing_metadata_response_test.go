package types

import (
	"testing"
)

func TestDubbingMetadataResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dubbing_id": "value",
  "name": "value",
  "status": "value",
  "source_language": "value",
  "target_languages": [
    "value"
  ],
  "editable": true,
  "created_at": "2024-01-01T00:00:00Z",
  "media_metadata": {
    "content_type": "value",
    "duration": 9.99
  },
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingMetadataResponse))
}
