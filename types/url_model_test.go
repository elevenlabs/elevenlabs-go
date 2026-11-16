package types

import (
	"testing"
)

func TestUrlModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url_string": "value",
  "scheme": "value",
  "hostname": "value",
  "port": 123,
  "path": "value",
  "query_string": "value"
}`
	assertRoundTrip(t, wire, new(UrlModel))
}
