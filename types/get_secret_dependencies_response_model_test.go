package types

import (
	"testing"
)

func TestGetSecretDependenciesResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dependencies": [
    {
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "next_cursor": "value"
}`
	assertRoundTrip(t, wire, new(GetSecretDependenciesResponseModel))
}
