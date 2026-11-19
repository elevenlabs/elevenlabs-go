package types

import (
	"testing"
)

func TestDirectPublishingReadResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "read_id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "word_count": 123,
  "char_count": 123,
  "chapters": [
    {
      "chapter_name": "value",
      "word_count": 123,
      "char_count": 123,
      "starting_char_offset": 123,
      "has_parsed_html": true,
      "has_summary": true,
      "duration_seconds": 9.99,
      "file_number": "value",
      "is_fallback_name": true,
      "chapter_id": "value"
    }
  ],
  "title": "value",
  "author": "value",
  "description": "value",
  "article_image_url": "value",
  "language": "value",
  "locale": "value",
  "display_mode": "text",
  "genre": [
    "Fantasy"
  ],
  "fiction": "value",
  "content_type": "value",
  "original_file_type": "value",
  "target_audience": "children",
  "mature_content": true,
  "safesearch_adult": true,
  "origin": "value",
  "publication_date": "value",
  "isbn": "value",
  "ean": "value",
  "legal_terms": {
    "terms": "value",
    "start_date": "value",
    "end_date": "value"
  },
  "content_guidelines_terms": {
    "terms": "value",
    "start_date": "value",
    "end_date": "value"
  },
  "last_updated_from_project_unix": 123,
  "publishing_project_id": "value",
  "publishing_state": "value",
  "publisher_profile_id": "value",
  "quality_score": 123,
  "publisher": "value",
  "copyright": "value",
  "subtitle": "value",
  "distribution_territories": [
    "value"
  ],
  "edition": "value",
  "contributors": [
    {
      "name": "value",
      "role": "value",
      "bio": "value",
      "profile_id": "value"
    }
  ],
  "payout_type": "none",
  "list_price": 9.99,
  "currency": "usd",
  "original_audio_project_export_id": "value",
  "original_audio_document_id": "value",
  "series_id": "value",
  "volume": 123,
  "published_at_unix": 123,
  "read_slug": "value",
  "preview_audio_object": {
    "voice_id": "value",
    "text": "value",
    "audio_url": "value",
    "hls_manifest_url": "value",
    "dash_manifest_url": "value",
    "is_auto_generated": true,
    "generated_at_unix": 123
  },
  "sample_config": {
    "is_sample": true,
    "parent_id": "value",
    "parent_type": "read",
    "chapter_ids": [
      "value"
    ]
  },
  "review": {
    "review_status": "approved",
    "reviewed_at_unix": 123,
    "reviewed_by": "value",
    "reject_reasons": [
      "lacks_structure"
    ],
    "scores_breakdown": {
      "key": 123
    },
    "rejected_details": "value",
    "explanation": "value"
  },
  "voice_id": "value",
  "can_use_assistant": true,
  "is_voice_changer_on": true,
  "restricted_to_user_email_domains": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(DirectPublishingReadResponseModel))
}
