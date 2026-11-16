package types

import (
	"testing"
)

func TestMediaGenerationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "generating"
}`
	assertRoundTrip(t, wire, new(MediaGenerationResponse))
}

func TestMediaGenerationResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "generating"
}`
	assertRoundTrip(t, wire, new(MediaGenerationResponse))
}

func TestMediaGenerationResponseRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "generating"
}`
	assertRoundTrip(t, wire, new(MediaGenerationResponse))
}
