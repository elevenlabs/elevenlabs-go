package types

import (
	"testing"
)

func TestProjectImageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "image_id": "value",
  "filename": "value",
  "signed_url": "value",
  "thumbnail_signed_url": "value",
  "type": "image",
  "source": "upload",
  "file_size_bytes": 123,
  "width": 123,
  "height": 123,
  "track_id": "value",
  "offset_ms": 123,
  "duration_ms": 123,
  "order": "value",
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
  "opacity": 0.5,
  "created_at_ms": 123,
  "updated_at_ms": 123,
  "current_snapshot_id": "value",
  "source_asset_id": "value",
  "source_platform_asset_id": "value",
  "error": "value",
  "pending_task": {
    "type": "preprocessing",
    "progress": 9.99,
    "started_at_ms": 123,
    "updated_at_ms": 123,
    "metadata": {
      "key": {}
    }
  },
  "analysis": {
    "status": "processing",
    "data": {
      "title": "value",
      "description": "value",
      "content_type": "value",
      "mood_and_style": "value",
      "composition_notes": "value",
      "visible_text": "value",
      "subjects": [
        {
          "name": "value",
          "description": "value"
        }
      ]
    },
    "updated_at_ms": 123
  }
}`
	assertRoundTrip(t, wire, new(ProjectImageResponseModel))
}
