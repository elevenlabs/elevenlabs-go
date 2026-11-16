package types

import (
	"testing"
)

func TestProjectVideoResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "video_id": "value",
  "filename": "value",
  "signed_url": "value",
  "signed_preview_url": "value",
  "offset_ms": 123,
  "duration_ms": 123,
  "volume_gain_db": 9.99,
  "muted": true,
  "fade_in_ms": 123,
  "fade_out_ms": 123,
  "width": 123,
  "height": 123,
  "codec": "value",
  "order": "value",
  "created_at_ms": 123,
  "updated_at_ms": 123,
  "error": "value",
  "thumbnail_interval_seconds": 9.99,
  "thumbnail_size": [
    123
  ],
  "thumbnail_sheets": [
    {
      "start_thumbnail_index": 123,
      "thumbnail_count": 123,
      "signed_cloud_url": "value"
    }
  ],
  "start_time_ms": 123,
  "end_time_ms": 123,
  "asset_preview_signed_url": "value",
  "source_video_id": "value",
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
  "audio_track_ready": true,
  "export_format_ready": true,
  "current_snapshot_id": "value",
  "source_context": {
    "source_type": "generation",
    "generation_id": "value",
    "prompt": "value",
    "model_id": "value",
    "model_provider": "value",
    "generation_session_id": "value",
    "session_iteration_id": "value",
    "model_parameters": {
      "key": {}
    },
    "extend_video": {
      "generation_id": "value",
      "content_asset_id": "value",
      "template_node_id": "value",
      "studio_clip": {
        "project_id": "value",
        "chapter_id": "value",
        "clip_type": "video",
        "clip_id": "value",
        "block_id": "value",
        "preview_url": "value"
      }
    }
  },
  "analysis": {
    "status": "processing",
    "data": {
      "title": "value",
      "description": "value",
      "content_type": "value",
      "overall_pacing": "value",
      "subjects": [
        {
          "name": "value",
          "description": "value"
        }
      ],
      "segments": [
        {
          "start_ms": 123,
          "end_ms": 123,
          "description": "value",
          "subjects": [
            "value"
          ],
          "shot_type": "value",
          "camera_movement": "value",
          "transition_in": "value",
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
  "type": "video",
  "canvas_placement": {
    "x_relative": 9.99,
    "y_relative": 9.99,
    "scale_x": 9.99,
    "scale_y": 9.99,
    "pivot_x": 9.99,
    "pivot_y": 9.99,
    "skew_x": 9.99,
    "skew_y": 9.99,
    "crop_top": 9.99,
    "crop_right": 9.99,
    "crop_bottom": 9.99,
    "crop_left": 9.99,
    "flip_x": true,
    "flip_y": true
  },
  "animation": {
    "enter_effect": "none",
    "enter_duration_ms": 123,
    "exit_effect": "none",
    "exit_duration_ms": 123
  },
  "playback_speed": 9.99,
  "opacity": 0.5,
  "track_id": "value",
  "preview_job_progress": 9.99,
  "import_speech_progress": 9.99
}`
	assertRoundTrip(t, wire, new(ProjectVideoResponseModel))
}
