package types

import (
	"testing"
)

func TestInternalKbCrawlPageTimingRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "doc_id": "value",
  "read_ms": 123,
  "index_ms": 123,
  "html_len": 123,
  "skipped": true,
  "error": true
}`
	assertRoundTrip(t, wire, new(InternalKbCrawlPageTiming))
}
