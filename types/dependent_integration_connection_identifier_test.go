package types

import (
	"testing"
)

func TestDependentIntegrationConnectionIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(DependentIntegrationConnectionIdentifier))
}
