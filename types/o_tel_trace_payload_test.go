package types

import (
	"testing"
)

func TestOTelTracePayloadRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resourceSpans": [
    {
      "resource": {
        "attributes": [
          {
            "key": "value",
            "value": {
              "key": "value"
            }
          }
        ]
      },
      "scopeSpans": [
        {
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
        }
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(OTelTracePayload))
}
