package types

import (
	"testing"
)

func TestFinalMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final": true,
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(FinalMulti))
}
