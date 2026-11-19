package types

import (
	"testing"
)

func TestCreateEnvironmentVariablesRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "string",
  "label": "value",
  "values": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(CreateEnvironmentVariablesRequestBody))
}
