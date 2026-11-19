package types

import (
	"testing"
)

func TestGetVoicesV2ResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voices": [
    {
      "voice_id": "value",
      "name": "value",
      "samples": [
        {
          "sample_id": "value",
          "file_name": "value",
          "mime_type": "value",
          "size_bytes": 123,
          "hash": "value",
          "duration_secs": 9.99,
          "remove_background_noise": true,
          "has_isolated_audio": true,
          "has_isolated_audio_preview": true,
          "speaker_separation": {
            "voice_id": "value",
            "sample_id": "value",
            "status": "not_started",
            "speakers": {
              "key": {
                "speaker_id": "value",
                "duration_secs": 9.99,
                "utterances": [
                  {
                    "start": 9.99,
                    "end": 9.99
                  }
                ]
              }
            },
            "selected_speaker_ids": [
              "value"
            ]
          },
          "trim_start": 123,
          "trim_end": 123
        }
      ],
      "category": "generated",
      "fine_tuning": {
        "is_allowed_to_fine_tune": true,
        "state": {
          "key": "not_started"
        },
        "verification_failures": [
          "value"
        ],
        "verification_attempts_count": 123,
        "manual_verification_requested": true,
        "language": "value",
        "progress": {
          "key": 9.99
        },
        "message": {
          "key": "value"
        },
        "dataset_duration_seconds": 9.99,
        "verification_attempts": [
          {
            "text": "value",
            "date_unix": 123,
            "accepted": true,
            "similarity": 9.99,
            "levenshtein_distance": 9.99,
            "recording": {
              "recording_id": "value",
              "mime_type": "value",
              "size_bytes": 123,
              "upload_date_unix": 123,
              "transcription": "value"
            }
          }
        ],
        "slice_ids": [
          "value"
        ],
        "manual_verification": {
          "extra_text": "value",
          "request_time_unix": 123,
          "files": [
            {
              "file_id": "value",
              "file_name": "value",
              "mime_type": "value",
              "size_bytes": 123,
              "upload_date_unix": 123
            }
          ]
        },
        "max_verification_attempts": 123,
        "next_max_verification_attempts_reset_unix_ms": 123,
        "finetuning_state": {}
      },
      "labels": {
        "key": "value"
      },
      "description": "value",
      "preview_url": "value",
      "available_for_tiers": [
        "value"
      ],
      "settings": {
        "stability": 0.5,
        "use_speaker_boost": true,
        "similarity_boost": 0.5,
        "style": 9.99,
        "speed": 9.99,
        "clone_strength": 0.5
      },
      "sharing": {
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
      },
      "high_quality_base_model_ids": [
        "value"
      ],
      "verified_languages": [
        {
          "language": "value",
          "model_id": "value",
          "accent": "value",
          "locale": "value",
          "preview_url": "value"
        }
      ],
      "collection_ids": [
        "value"
      ],
      "safety_control": "NONE",
      "voice_verification": {
        "requires_verification": true,
        "is_verified": true,
        "verification_failures": [
          "value"
        ],
        "verification_attempts_count": 123,
        "language": "value",
        "verification_attempts": [
          {
            "text": "value",
            "date_unix": 123,
            "accepted": true,
            "similarity": 9.99,
            "levenshtein_distance": 9.99,
            "recording": {
              "recording_id": "value",
              "mime_type": "value",
              "size_bytes": 123,
              "upload_date_unix": 123,
              "transcription": "value"
            }
          }
        ]
      },
      "permission_on_resource": "value",
      "is_owner": true,
      "is_legacy": true,
      "is_mixed": true,
      "favorited_at_unix": 123,
      "created_at_unix": 123,
      "is_bookmarked": true,
      "recording_quality": "studio",
      "labelling_status": "in_review",
      "recording_quality_reason": "value"
    }
  ],
  "has_more": true,
  "total_count": 123,
  "next_page_token": "value"
}`
	assertRoundTrip(t, wire, new(GetVoicesV2Response))
}
