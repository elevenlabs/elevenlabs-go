package types

import (
	"testing"
)

func TestKnowledgeBaseDocumentMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at_unix_secs": 123,
  "last_updated_at_unix_secs": 123,
  "size_bytes": 123
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDocumentMetadataResponseModel))
}

func TestKnowledgeBaseDocumentMetadataResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at_unix_secs": 123,
  "last_updated_at_unix_secs": 123,
  "size_bytes": 123
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDocumentMetadataResponseModel))
}
