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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireTexttovoiceClient(t, 200, `{
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

	var request types.BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost
	if err := json.Unmarshal([]byte(`{
  "voice_name": "value",
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "generated_voice_id": "value",
  "labels": {
    "key": "value"
  },
  "played_not_selected_voice_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualTexttovoiceClient(gotResponse, []byte(`{
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

	req := receiveWireTexttovoiceClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/text-to-voice" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/text-to-voice")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsTexttovoiceClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError409(t *testing.T) {
	t.Parallel()

	server, _ := serveWireTexttovoiceClient(t, 409, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost
	if err := json.Unmarshal([]byte(`{
  "voice_name": "value",
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "generated_voice_id": "value",
  "labels": {
    "key": "value"
  },
  "played_not_selected_voice_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 409 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 409 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 409)
	}
}

func TestDesignWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireTexttovoiceClient(t, 200, `{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.VoiceDesignRequestModel
	if err := json.Unmarshal([]byte(`{
  "output_format": "mp3_22050_32",
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "model_id": "eleven_multilingual_ttv_v2",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "should_enhance": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "quality": 0.5,
  "reference_audio_base64": "value",
  "prompt_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Design(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualTexttovoiceClient(gotResponse, []byte(`{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)
	}

	req := receiveWireTexttovoiceClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/text-to-voice/design" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/text-to-voice/design")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsTexttovoiceClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestDesignWireError409(t *testing.T) {
	t.Parallel()

	server, _ := serveWireTexttovoiceClient(t, 409, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.VoiceDesignRequestModel
	if err := json.Unmarshal([]byte(`{
  "output_format": "mp3_22050_32",
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "model_id": "eleven_multilingual_ttv_v2",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "should_enhance": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "quality": 0.5,
  "reference_audio_base64": "value",
  "prompt_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Design(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 409 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 409 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 409)
	}
}

func TestRemixWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireTexttovoiceClient(t, 200, `{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.VoiceRemixRequestModel
	if err := json.Unmarshal([]byte(`{
  "output_format": "mp3_22050_32",
  "voice_id": "value",
  "voice_description": "value",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "prompt_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Remix(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualTexttovoiceClient(gotResponse, []byte(`{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`)
	}

	req := receiveWireTexttovoiceClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/text-to-voice/value/remix" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/text-to-voice/value/remix")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsTexttovoiceClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestRemixWireError409(t *testing.T) {
	t.Parallel()

	server, _ := serveWireTexttovoiceClient(t, 409, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.VoiceRemixRequestModel
	if err := json.Unmarshal([]byte(`{
  "output_format": "mp3_22050_32",
  "voice_id": "value",
  "voice_description": "value",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "prompt_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Remix(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 409 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 409 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 409)
	}
}

type wireRequestTexttovoiceClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireTexttovoiceClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireTexttovoiceClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestTexttovoiceClient) {
	t.Helper()
	recorded := make(chan wireRequestTexttovoiceClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestTexttovoiceClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireTexttovoiceClient(t *testing.T, recorded <-chan wireRequestTexttovoiceClient) wireRequestTexttovoiceClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestTexttovoiceClient{}
	}
}

// wireJSONContainsTexttovoiceClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsTexttovoiceClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualTexttovoiceClient(got, want)
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

// setWireValueTexttovoiceClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueTexttovoiceClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualTexttovoiceClient(got, want []byte) bool {
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
