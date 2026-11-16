package types

import (
	"testing"
)

func TestVoiceSharingResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "enabled",
  "history_item_sample_id": "value",
  "date_unix": 123,
  "whitelisted_emails": [
    "value"
  ],
  "public_owner_id": "value",
  "original_voice_id": "value",
  "financial_rewards_enabled": true,
  "free_users_allowed": true,
  "live_moderation_enabled": true,
  "rate": 9.99,
  "fiat_rate": 9.99,
  "notice_period": 123,
  "disable_at_unix": 123,
  "voice_mixing_allowed": true,
  "featured": true,
  "category": "generated",
  "reader_app_enabled": true,
  "image_url": "value",
  "ban_reason": "value",
  "liked_by_count": 123,
  "cloned_by_count": 123,
  "name": "value",
  "description": "value",
  "labels": {
    "key": "value"
  },
  "review_status": "not_requested",
  "review_message": "value",
  "enabled_in_library": true,
  "instagram_username": "value",
  "twitter_username": "value",
  "youtube_username": "value",
  "tiktok_username": "value",
  "moderation_check": {
    "date_checked_unix": 123,
    "name_value": "value",
    "name_check": true,
    "description_value": "value",
    "description_check": true,
    "sample_ids": [
      "value"
    ],
    "sample_checks": [
      9.99
    ],
    "captcha_ids": [
      "value"
    ],
    "captcha_checks": [
      9.99
    ]
  },
  "reader_restricted_on": [
    {
      "resource_type": "read",
      "resource_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(VoiceSharingResponse))
}
