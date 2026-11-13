package types

import (
	"testing"
)

func TestProjectExternalAudioResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "external_audio_id": "value",
  "filename": "value",
  "signed_url": "value",
  "offset_ms": 123,
  "duration_ms": 123,
  "start_time_ms": 123,
  "end_time_ms": 123,
  "order": "value",
  "track_id": "value",
  "created_at_ms": 123,
  "updated_at_ms": 123,
  "volume_gain_db": 9.99,
  "muted": true,
  "fade_in_ms": 123,
  "fade_out_ms": 123,
  "source_external_audio_id": "value",
  "source_asset_id": "value",
  "source_platform_asset_id": "value",
  "pending_blocks_metadata": {
    "target_global_offset_ms": 123,
    "block_ids": [
      "value"
    ]
  },
  "pending_external_audios_metadata": {
    "target_global_offset_ms": 123,
    "external_audio_ids": [
      "value"
    ]
  },
  "speech_imported": true,
  "pending_task": {
    "type": "preprocessing",
    "progress": 9.99,
    "started_at_ms": 123,
    "updated_at_ms": 123,
    "metadata": {
      "key": {}
    }
  },
  "error": "value",
  "current_snapshot_id": "value",
  "source_context": {
    "source_type": "song",
    "song_id": "value",
    "chat_id": "value",
    "title": "value",
    "description": "value",
    "genres": [
      "value"
    ],
    "languages": [
      "value"
    ],
    "is_explicit": true,
    "bpm": 123,
    "generation_settings": {
      "key": {}
    }
  },
  "analysis": {
    "status": "processing",
    "data": {
      "title": "value",
      "description": "value",
      "content_type": "value",
      "overall_pacing": "value",
      "segments": [
        {
          "start_ms": 123,
          "end_ms": 123,
          "description": "value",
          "segment_type": "value",
          "has_speech": true,
          "has_music": true,
          "pacing": "value"
        }
      ],
      "key_moments": [
        {
          "timestamp_ms": 123,
          "type": "value",
          "description": "value"
        }
      ]
    },
    "updated_at_ms": 123
  },
  "transcription": {
    "status": "processing",
    "data": {
      "language_code": "value",
      "text": "value",
      "words": [
        "value"
      ],
      "word_start_times_ms": [
        123
      ],
      "word_end_times_ms": [
        123
      ],
      "word_speaker_ids": [
        "value"
      ]
    },
    "updated_at_ms": 123
  },
  "type": "audio",
  "import_speech_progress": 9.99
}`
	assertRoundTrip(t, wire, new(ProjectExternalAudioResponseModel))
}
