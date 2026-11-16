package types

import (
	"testing"
)

func TestImageAnalysisResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ImageAnalysisResult))
}
