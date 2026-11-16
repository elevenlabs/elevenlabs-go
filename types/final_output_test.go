package types

import (
	"testing"
)

func TestFinalOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final": true
}`
	assertRoundTrip(t, wire, new(FinalOutput))
}
