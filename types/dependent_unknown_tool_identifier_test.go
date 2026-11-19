package types

import (
	"testing"
)

func TestDependentUnknownToolIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "unknown",
  "id": "value"
}`
	assertRoundTrip(t, wire, new(DependentUnknownToolIdentifier))
}

func TestDependentUnknownToolIdentifierRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "unknown",
  "id": "value"
}`
	assertRoundTrip(t, wire, new(DependentUnknownToolIdentifier))
}
