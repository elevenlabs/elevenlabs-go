package types

import (
	"testing"
)

func TestOTelSpanRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "traceId": "value",
  "spanId": "value",
  "parentSpanId": "value",
  "name": "value",
  "kind": 123,
  "startTimeUnixNano": "value",
  "endTimeUnixNano": "value",
  "attributes": [
    {
      "key": "value",
      "value": {
        "key": "value"
      }
    }
  ],
  "status": {
    "code": 123,
    "message": "value"
  }
}`
	assertRoundTrip(t, wire, new(OTelSpan))
}
