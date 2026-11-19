package types

import (
	"testing"
)

func TestMediaGenerationListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "generations": [
    {
      "id": "value",
      "status": "generating"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(MediaGenerationListResponse))
}

func TestMediaGenerationListResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "generations": [
    {
      "id": "value",
      "status": "generating"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(MediaGenerationListResponse))
}

func TestMediaGenerationListResponseRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "generations": [
    {
      "id": "value",
      "status": "generating"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(MediaGenerationListResponse))
}
