package types

import (
	"testing"
)

func TestOTelScopeSpansRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "scope": {
    "name": "value",
    "version": "value"
  },
  "spans": [
    {
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
    }
  ]
}`
	assertRoundTrip(t, wire, new(OTelScopeSpans))
}
