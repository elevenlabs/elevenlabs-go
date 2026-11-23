package types

import (
	"testing"
)

func TestMultimodalMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "multimodal_message",
  "text": {
    "id": "value",
    "type": "user_message",
    "text": "value",
    "user_identifier": "value",
    "source_medium": "audio"
  },
  "file": {
    "type": "file_input",
    "file_id": "value"
  },
  "files": [
    {
      "type": "file_input",
      "file_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(MultimodalMessage))
}
