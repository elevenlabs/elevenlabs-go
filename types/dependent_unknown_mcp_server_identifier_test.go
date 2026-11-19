package types

import (
	"testing"
)

func TestDependentUnknownMcpServerIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownMcpServerIdentifier))
}

func TestDependentUnknownMcpServerIdentifierRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownMcpServerIdentifier))
}
