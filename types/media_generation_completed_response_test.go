package types

import (
	"testing"
)

func TestMediaGenerationCompletedResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "completed",
  "content_url": "value",
  "content_mime_type": "value"
}`
	assertRoundTrip(t, wire, new(MediaGenerationCompletedResponse))
}
