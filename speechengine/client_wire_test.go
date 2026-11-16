package speechengine

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechengine(t, 200, `{
  "speech_engines": [
    {
      "speech_engine_id": "value",
      "name": "value",
      "voice_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "tags": [
        "value"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &types.ListSpeechEngineRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualSpeechengine(gotResponse, []byte(`{
  "speech_engines": [
    {
      "speech_engine_id": "value",
      "name": "value",
      "voice_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "tags": [
        "value"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "speech_engines": [
    {
      "speech_engine_id": "value",
      "name": "value",
      "voice_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "tags": [
        "value"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/speech-engine" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	pageSizeValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &types.ListSpeechEngineRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	_, err := client.List(context.Background(), request)
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechengine(t, 201, `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.CreateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  }
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
	if !wireJSONBytesEqualSpeechengine(gotResponse, []byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/speech-engine" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsSpeechengine(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	var request types.CreateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechengine(t, 200, `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualSpeechengine(gotResponse, []byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/speech-engine/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"))
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechengine(t, 200, `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.UpdateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualSpeechengine(gotResponse, []byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/speech-engine/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsSpeechengine(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	var request types.UpdateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
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

	server, recorded := serveWireSpeechengine(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/speech-engine/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	err := client.Delete(context.Background(), string("value"))
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

func TestDuplicateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechengine(t, 201, `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.DuplicateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "speech_engine_id": "value",
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Duplicate(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualSpeechengine(gotResponse, []byte(`{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "speech_engine_id": "value",
  "name": "value",
  "speech_engine": {
    "ws_url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "asr": {
    "quality": "high",
    "provider": "elevenlabs",
    "user_input_audio_format": "pcm_8000",
    "keywords": [
      "value"
    ]
  },
  "tts": {
    "model_id": "eleven_turbo_v2",
    "voice_id": "value",
    "supported_voices": [
      {
        "label": "value",
        "voice_id": "value",
        "description": "value",
        "language": "value",
        "model_family": "turbo",
        "optimize_streaming_latency": 123,
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5
      }
    ],
    "expressive_mode": true,
    "suggested_audio_tags": [
      {
        "tag": "value",
        "description": "value"
      }
    ],
    "agent_output_audio_format": "pcm_8000",
    "optimize_streaming_latency": 123,
    "stability": 0.5,
    "speed": 1.2,
    "similarity_boost": 0.5,
    "clone_strength": 0.5,
    "text_normalisation_type": "system_prompt",
    "pronunciation_dictionary_locators": [
      {
        "pronunciation_dictionary_id": "value",
        "version_id": "value"
      }
    ],
    "enable_phoneme_tags": true,
    "audio_effects": {
      "filter_preset_id": "value",
      "distance": 0.5,
      "environment_id": "value",
      "background_noise_id": "value",
      "send_level": 0.5,
      "pan": 0.5
    }
  },
  "turn": {
    "turn_timeout": 9.99,
    "initial_wait_time": 9.99,
    "silence_end_call_timeout": 9.99,
    "turn_eagerness": "patient",
    "spelling_patience": "auto",
    "speculative_turn": true,
    "retranscribe_on_turn_timeout": true,
    "turn_model": "turn_v2",
    "interruption_ignore_terms": [
      "value"
    ],
    "interruption_ignore_term_languages": [
      "value"
    ],
    "merge_with_default_ignore_terms": true,
    "transcribe_on_disabled_interruptions": true
  },
  "vad": {},
  "conversation": {
    "text_only": true,
    "max_duration_seconds": 123,
    "client_events": [
      "conversation_initiation_metadata"
    ],
    "file_input": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    },
    "monitoring_enabled": true,
    "monitoring_events": [
      "conversation_initiation_metadata"
    ],
    "dtmf_input_settings": {
      "dtmf_input_timeout": 9.99,
      "hash_terminator": true,
      "redact_input": true
    },
    "background_sound": {
      "source_type": "preset",
      "source_id": "office2",
      "volume": 0.5,
      "crossfade_loop": true
    },
    "source_attribution": true
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "auth": {
    "enable_auth": true
  },
  "language": "value",
  "cascade_timeout_seconds": 9.99,
  "tags": [
    "value"
  ],
  "overrides": {
    "first_message": true
  },
  "metadata": {
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "created_from": "cli",
    "last_updated_from": "cli"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireSpeechengine(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/speech-engine/value/duplicate" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-engine/value/duplicate")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsSpeechengine(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestDuplicateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechengine(t, 422, `{
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

	var request types.DuplicateSpeechEngineRequest
	if err := json.Unmarshal([]byte(`{
  "speech_engine_id": "value",
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Duplicate(context.Background(), string("value"), &request)
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

type wireRequestSpeechengine struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireSpeechengine records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireSpeechengine(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestSpeechengine) {
	t.Helper()
	recorded := make(chan wireRequestSpeechengine, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestSpeechengine{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireSpeechengine(t *testing.T, recorded <-chan wireRequestSpeechengine) wireRequestSpeechengine {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestSpeechengine{}
	}
}

// wireJSONContainsSpeechengine reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsSpeechengine(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualSpeechengine(got, want)
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

// setWireValueSpeechengine sets a named string, int or float field, or the value a pointer field points to.
func setWireValueSpeechengine(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualSpeechengine(got, want []byte) bool {
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
