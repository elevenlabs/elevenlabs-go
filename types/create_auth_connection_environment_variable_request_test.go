package types

import (
	"testing"
)

func TestCreateAuthConnectionEnvironmentVariableRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "auth_connection",
  "label": "value",
  "values": {
    "key": {
      "auth_connection_id": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(CreateAuthConnectionEnvironmentVariableRequest))
}
