package types

import (
	"testing"
)

func TestGenerationReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "generation",
  "generation_id": "value"
}`
	assertRoundTrip(t, wire, new(GenerationReference))
}

func TestGenerationReferenceRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "generation",
  "generation_id": "value"
}`
	assertRoundTrip(t, wire, new(GenerationReference))
}
