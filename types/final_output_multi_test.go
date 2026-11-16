package types

import (
	"testing"
)

func TestFinalOutputMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final": true,
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(FinalOutputMulti))
}
