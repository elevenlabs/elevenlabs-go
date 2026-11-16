package types

import (
	"testing"
)

func TestGetSecretDependenciesResponseModelDependenciesZeroItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(GetSecretDependenciesResponseModelDependenciesZeroItem))
}
