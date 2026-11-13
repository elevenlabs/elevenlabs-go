package types

import (
	"testing"
)

func TestUpdateSpeechEngineRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(UpdateSpeechEngineRequest))
}
