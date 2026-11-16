package types

import (
	"testing"
)

func TestUpdateEnvironmentVariableRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "values": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(UpdateEnvironmentVariableRequest))
}
