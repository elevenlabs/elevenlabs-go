package types

import (
	"testing"
)

func TestClientErrorErrorEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": 123,
  "error_name": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(ClientErrorErrorEvent))
}
