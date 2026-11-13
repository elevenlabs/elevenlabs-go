package types

import (
	"testing"
)

func TestMediaGenerationCreateResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending"
}`
	assertRoundTrip(t, wire, new(MediaGenerationCreateResponse))
}

func TestMediaGenerationCreateResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending"
}`
	assertRoundTrip(t, wire, new(MediaGenerationCreateResponse))
}

func TestMediaGenerationCreateResponseRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending"
}`
	assertRoundTrip(t, wire, new(MediaGenerationCreateResponse))
}
