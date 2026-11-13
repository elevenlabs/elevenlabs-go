package types

import (
	"testing"
)

func TestImageAnalysisRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "processing",
  "data": {
    "title": "value",
    "description": "value",
    "content_type": "value",
    "mood_and_style": "value",
    "composition_notes": "value",
    "visible_text": "value",
    "subjects": [
      {
        "name": "value",
        "description": "value"
      }
    ]
  },
  "updated_at_ms": 123
}`
	assertRoundTrip(t, wire, new(ImageAnalysis))
}
