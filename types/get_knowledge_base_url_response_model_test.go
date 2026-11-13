package types

import (
	"testing"
)

func TestGetKnowledgeBaseUrlResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "folder_path": [
    {
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseUrlResponseModel))
}

func TestGetKnowledgeBaseUrlResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "folder_path": [
    {
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseUrlResponseModel))
}
