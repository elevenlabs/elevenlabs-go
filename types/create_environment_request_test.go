package types

import (
	"testing"
)

func TestCreateEnvironmentRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(CreateEnvironmentRequest))
}
