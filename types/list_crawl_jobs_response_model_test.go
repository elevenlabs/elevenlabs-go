package types

import (
	"testing"
)

func TestListCrawlJobsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "crawl_jobs": [
    {
      "type": "discovery",
      "seed_url": "value",
      "pattern": "value",
      "exclude_patterns": [
        "value"
      ],
      "max_pages": 123,
      "auto_discover": true,
      "status": "queued",
      "pages_identified": 123,
      "pages_scraped": 123,
      "pages_skipped": 123,
      "pages_failed": 123,
      "root_folder_id": "value",
      "updated_at": 123,
      "id": "value",
      "created_at": 123
    }
  ],
  "next_cursor": "value"
}`
	assertRoundTrip(t, wire, new(ListCrawlJobsResponseModel))
}
