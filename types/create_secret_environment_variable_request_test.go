package types

import (
	"testing"
)

func TestCreateSecretEnvironmentVariableRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "secret",
  "label": "value",
  "values": {
    "key": {
      "secret_id": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(CreateSecretEnvironmentVariableRequest))
}
