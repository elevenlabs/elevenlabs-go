package types

import (
	"testing"
)

func TestMediaGenerationInProgressResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending"
}`
	assertRoundTrip(t, wire, new(MediaGenerationInProgressResponse))
}
