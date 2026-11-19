package types

import (
	"testing"
)

func TestGetLibraryVoicesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voices": [
    {
      "public_owner_id": "value",
      "voice_id": "value",
      "date_unix": 123,
      "name": "value",
      "accent": "value",
      "gender": "value",
      "age": "value",
      "descriptive": "value",
      "use_case": "value",
      "category": "generated",
      "language": "value",
      "locale": "value",
      "description": "value",
      "preview_url": "value",
      "usage_character_count_1y": 123,
      "usage_character_count_7d": 123,
      "play_api_usage_character_count_1y": 123,
      "cloned_by_count": 123,
      "rate": 9.99,
      "fiat_rate": 9.99,
      "free_users_allowed": true,
      "live_moderation_enabled": true,
      "featured": true,
      "verified_languages": [
        {
          "language": "value",
          "model_id": "value",
          "accent": "value",
          "locale": "value",
          "preview_url": "value"
        }
      ],
      "notice_period": 123,
      "instagram_username": "value",
      "twitter_username": "value",
      "youtube_username": "value",
      "tiktok_username": "value",
      "image_url": "value",
      "is_added_by_user": true,
      "is_bookmarked": true
    }
  ],
  "has_more": true,
  "total_count": 123,
  "last_sort_id": "value"
}`
	assertRoundTrip(t, wire, new(GetLibraryVoicesResponse))
}
