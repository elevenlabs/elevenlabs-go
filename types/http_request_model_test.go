package types

import (
	"testing"
)

func TestHttpRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "http_method": "value",
  "url": {
    "url_string": "value",
    "scheme": "value",
    "hostname": "value",
    "port": 123,
    "path": "value",
    "query_string": "value"
  },
  "user_agent": "value",
  "x_forwarded_for": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(HttpRequestModel))
}
