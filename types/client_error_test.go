package types

import (
	"testing"
)

func TestClientErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "error_event": {
    "code": 123,
    "error_name": "value",
    "message": "value"
  },
  "type": "client_error"
}`
	assertRoundTrip(t, wire, new(ClientError))
}
