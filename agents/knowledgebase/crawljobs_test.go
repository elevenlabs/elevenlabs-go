package knowledgebase

import (
	"testing"
)

func TestBodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "max_depth": 123,
  "max_pages": 123,
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "sitemap_urls": [
    "value"
  ],
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "auto_discover": true,
  "minimum_frequency_days": 123
}`
	assertRoundTrip(t, wire, new(BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost))
}
