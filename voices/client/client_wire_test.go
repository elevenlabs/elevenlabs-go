package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	withSettingsValue := bool(true)
	request := &types.GetVoicesRequest{
		WithSettings: &withSettingsValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/value")
	}
	if got := req.Query.Get("with_settings"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "with_settings", got, "true")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	withSettingsValue := bool(true)
	request := &types.GetVoicesRequest{
		WithSettings: &withSettingsValue,
	}
	_, err := client.Get(context.Background(), string("value"), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/voices/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Delete(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestSearchWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	nextPageTokenValue := string("value")
	pageSizeValue := int(123)
	searchValue := string("value")
	sortValue := string("value")
	sortDirectionValue := string("value")
	voiceTypeValue := string("value")
	categoryValue := string("value")
	fineTuningStateValue := string("value")
	collectionIdValue := string("value")
	genderValue := string("value")
	ageValue := string("value")
	accentValue := string("value")
	minNoticePeriodDaysValue := int(123)
	includeCustomRatesValue := bool(true)
	includeLiveModeratedValue := bool(true)
	highQualityValue := bool(true)
	includeTotalCountValue := bool(true)
	request := &types.SearchVoicesRequest{
		NextPageToken:        &nextPageTokenValue,
		PageSize:             &pageSizeValue,
		Search:               &searchValue,
		Sort:                 &sortValue,
		SortDirection:        &sortDirectionValue,
		VoiceType:            &voiceTypeValue,
		Category:             &categoryValue,
		FineTuningState:      &fineTuningStateValue,
		CollectionId:         &collectionIdValue,
		Gender:               &genderValue,
		Age:                  &ageValue,
		Accent:               &accentValue,
		MinNoticePeriodDays:  &minNoticePeriodDaysValue,
		IncludeCustomRates:   &includeCustomRatesValue,
		IncludeLiveModerated: &includeLiveModeratedValue,
		HighQuality:          &highQualityValue,
		IncludeTotalCount:    &includeTotalCountValue,
	}
	result, err := client.Search(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v2/voices" {
		t.Fatalf("path: got %q want %q", req.Path, "/v2/voices")
	}
	if got := req.Query.Get("next_page_token"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "next_page_token", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("sort"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "sort", got, "value")
	}
	if got := req.Query.Get("sort_direction"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "sort_direction", got, "value")
	}
	if got := req.Query.Get("voice_type"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "voice_type", got, "value")
	}
	if got := req.Query.Get("category"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "category", got, "value")
	}
	if got := req.Query.Get("fine_tuning_state"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "fine_tuning_state", got, "value")
	}
	if got := req.Query.Get("collection_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "collection_id", got, "value")
	}
	if got := req.Query.Get("gender"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "gender", got, "value")
	}
	if got := req.Query.Get("age"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "age", got, "value")
	}
	if got := req.Query.Get("accent"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "accent", got, "value")
	}
	if got := req.Query.Get("min_notice_period_days"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "min_notice_period_days", got, "123")
	}
	if got := req.Query.Get("include_custom_rates"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_custom_rates", got, "true")
	}
	if got := req.Query.Get("include_live_moderated"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_live_moderated", got, "true")
	}
	if got := req.Query.Get("high_quality"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "high_quality", got, "true")
	}
	if got := req.Query.Get("include_total_count"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_total_count", got, "true")
	}
}

func TestSearchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	nextPageTokenValue := string("value")
	pageSizeValue := int(123)
	searchValue := string("value")
	sortValue := string("value")
	sortDirectionValue := string("value")
	voiceTypeValue := string("value")
	categoryValue := string("value")
	fineTuningStateValue := string("value")
	collectionIdValue := string("value")
	genderValue := string("value")
	ageValue := string("value")
	accentValue := string("value")
	minNoticePeriodDaysValue := int(123)
	includeCustomRatesValue := bool(true)
	includeLiveModeratedValue := bool(true)
	highQualityValue := bool(true)
	includeTotalCountValue := bool(true)
	request := &types.SearchVoicesRequest{
		NextPageToken:        &nextPageTokenValue,
		PageSize:             &pageSizeValue,
		Search:               &searchValue,
		Sort:                 &sortValue,
		SortDirection:        &sortDirectionValue,
		VoiceType:            &voiceTypeValue,
		Category:             &categoryValue,
		FineTuningState:      &fineTuningStateValue,
		CollectionId:         &collectionIdValue,
		Gender:               &genderValue,
		Age:                  &ageValue,
		Accent:               &accentValue,
		MinNoticePeriodDays:  &minNoticePeriodDaysValue,
		IncludeCustomRates:   &includeCustomRatesValue,
		IncludeLiveModerated: &includeLiveModeratedValue,
		HighQuality:          &highQualityValue,
		IncludeTotalCount:    &includeTotalCountValue,
	}
	_, err := client.Search(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestReplicateToIsolatedEnvironmentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
  "voice_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.ReplicateVoiceToIsolatedEnvironmentRequestModel
	if err := json.Unmarshal([]byte(`{
  "voice_id": "value",
  "target_workspace_id": "value",
  "preserve_voice_id": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.ReplicateToIsolatedEnvironment(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
  "voice_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "voice_id": "value"
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/voices/value/replicate-to-isolated-environment" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/value/replicate-to-isolated-environment")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsVoicesClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestReplicateToIsolatedEnvironmentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.ReplicateVoiceToIsolatedEnvironmentRequestModel
	if err := json.Unmarshal([]byte(`{
  "voice_id": "value",
  "target_workspace_id": "value",
  "preserve_voice_id": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.ReplicateToIsolatedEnvironment(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestShareWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
  "voice_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost
	if err := json.Unmarshal([]byte(`{
  "voice_id": "value",
  "public_user_id": "value",
  "new_name": "value",
  "bookmarked": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Share(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
  "voice_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "voice_id": "value"
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/voices/add/value/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/add/value/value2")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsVoicesClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestShareWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost
	if err := json.Unmarshal([]byte(`{
  "voice_id": "value",
  "public_user_id": "value",
  "new_name": "value",
  "bookmarked": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Share(context.Background(), string("value"), string("value2"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestGetSharedWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesClient(t, 200, `{
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(123)
	genderValue := string("value")
	ageValue := string("value")
	accentValue := string("value")
	languageValue := string("value")
	localeValue := string("value")
	searchValue := string("value")
	featuredValue := bool(true)
	minNoticePeriodDaysValue := int(123)
	includeCustomRatesValue := bool(true)
	includeLiveModeratedValue := bool(true)
	readerAppEnabledValue := bool(true)
	ownerIdValue := string("value")
	pageValue := int(123)
	request := &types.GetSharedVoicesRequest{
		PageSize:             &pageSizeValue,
		Gender:               &genderValue,
		Age:                  &ageValue,
		Accent:               &accentValue,
		Language:             &languageValue,
		Locale:               &localeValue,
		Search:               &searchValue,
		Featured:             &featuredValue,
		MinNoticePeriodDays:  &minNoticePeriodDaysValue,
		IncludeCustomRates:   &includeCustomRatesValue,
		IncludeLiveModerated: &includeLiveModeratedValue,
		ReaderAppEnabled:     &readerAppEnabledValue,
		OwnerId:              &ownerIdValue,
		Page:                 &pageValue,
	}
	result, err := client.GetShared(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesClient(gotResponse, []byte(`{
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
}`)
	}

	req := receiveWireVoicesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/shared-voices" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/shared-voices")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("gender"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "gender", got, "value")
	}
	if got := req.Query.Get("age"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "age", got, "value")
	}
	if got := req.Query.Get("accent"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "accent", got, "value")
	}
	if got := req.Query.Get("language"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "language", got, "value")
	}
	if got := req.Query.Get("locale"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "locale", got, "value")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("featured"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "featured", got, "true")
	}
	if got := req.Query.Get("min_notice_period_days"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "min_notice_period_days", got, "123")
	}
	if got := req.Query.Get("include_custom_rates"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_custom_rates", got, "true")
	}
	if got := req.Query.Get("include_live_moderated"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_live_moderated", got, "true")
	}
	if got := req.Query.Get("reader_app_enabled"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "reader_app_enabled", got, "true")
	}
	if got := req.Query.Get("owner_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "owner_id", got, "value")
	}
	if got := req.Query.Get("page"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page", got, "123")
	}
}

func TestGetSharedWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	pageSizeValue := int(123)
	genderValue := string("value")
	ageValue := string("value")
	accentValue := string("value")
	languageValue := string("value")
	localeValue := string("value")
	searchValue := string("value")
	featuredValue := bool(true)
	minNoticePeriodDaysValue := int(123)
	includeCustomRatesValue := bool(true)
	includeLiveModeratedValue := bool(true)
	readerAppEnabledValue := bool(true)
	ownerIdValue := string("value")
	pageValue := int(123)
	request := &types.GetSharedVoicesRequest{
		PageSize:             &pageSizeValue,
		Gender:               &genderValue,
		Age:                  &ageValue,
		Accent:               &accentValue,
		Language:             &languageValue,
		Locale:               &localeValue,
		Search:               &searchValue,
		Featured:             &featuredValue,
		MinNoticePeriodDays:  &minNoticePeriodDaysValue,
		IncludeCustomRates:   &includeCustomRatesValue,
		IncludeLiveModerated: &includeLiveModeratedValue,
		ReaderAppEnabled:     &readerAppEnabledValue,
		OwnerId:              &ownerIdValue,
		Page:                 &pageValue,
	}
	_, err := client.GetShared(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

type wireRequestVoicesClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesClient) {
	t.Helper()
	recorded := make(chan wireRequestVoicesClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if responseBody != "" {
			_, _ = w.Write([]byte(responseBody))
		}
	}))
	t.Cleanup(server.Close)
	return server, recorded
}

func receiveWireVoicesClient(t *testing.T, recorded <-chan wireRequestVoicesClient) wireRequestVoicesClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesClient{}
	}
}

// wireJSONContainsVoicesClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesClient(got, want)
	}
	for key, wantValue := range wantObject {
		if key == "stream" {
			continue
		}
		if gotValue, ok := gotObject[key]; !ok || !reflect.DeepEqual(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// setWireValueVoicesClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesClient(t *testing.T, target any, value any) {
	t.Helper()
	field := reflect.ValueOf(target).Elem()
	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}
	switch typed := value.(type) {
	case string:
		field.SetString(typed)
	case int:
		field.SetInt(int64(typed))
	case float64:
		field.SetFloat(typed)
	default:
		t.Fatalf("unsupported value %T", value)
	}
}

func wireJSONBytesEqualVoicesClient(got, want []byte) bool {
	var gotValue any
	var wantValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		return false
	}
	if err := json.Unmarshal(want, &wantValue); err != nil {
		return false
	}
	return reflect.DeepEqual(gotValue, wantValue)
}
