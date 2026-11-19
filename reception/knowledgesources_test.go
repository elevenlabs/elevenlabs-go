package reception

import (
	"testing"
)

func TestAddUrlRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "crawl_entire_site": true
}`
	assertRoundTrip(t, wire, new(AddUrlRequest))
}
