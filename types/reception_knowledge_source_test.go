package types

import (
	"testing"
)

func TestReceptionKnowledgeSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionKnowledgeSource))
}
