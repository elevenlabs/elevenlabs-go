package types

import (
	"testing"
)

func TestShareOptionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "id": "value",
  "type": "user"
}`
	assertRoundTrip(t, wire, new(ShareOptionResponseModel))
}
