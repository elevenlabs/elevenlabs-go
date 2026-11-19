package types

import (
	"testing"
)

func TestProjectResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "name": "value",
  "create_date_unix": 123,
  "created_by_user_id": "value",
  "default_title_voice_ref_id": "value",
  "default_paragraph_voice_ref_id": "value",
  "default_model_id": "value",
  "last_conversion_date_unix": 123,
  "can_be_downloaded": true,
  "title": "value",
  "author": "value",
  "description": "value",
  "genres": [
    "value"
  ],
  "cover_image_url": "value",
  "target_audience": "children",
  "language": "value",
  "content_type": "value",
  "original_publication_date": "value",
  "mature_content": true,
  "isbn_number": "value",
  "volume_normalization": true,
  "state": "creating",
  "access_level": "admin",
  "fiction": "fiction",
  "quality_check_on": true,
  "quality_check_on_when_bulk_convert": true,
  "creation_meta": {
    "creation_progress": 9.99,
    "status": "draft",
    "type": "blank"
  },
  "source_type": "blank",
  "chapters_enabled": true,
  "captions_enabled": true,
  "caption_style": {
    "template": {
      "key": "value",
      "label": "value",
      "requires_high_fps": true
    },
    "text_font": "value",
    "text_scale": 9.99,
    "text_color": "value",
    "text_align": "start",
    "text_style": "normal",
    "text_weight": "normal",
    "text_transform": "none",
    "text_blend_mode": "normal",
    "text_shadow": {
      "enabled": true,
      "color": "value",
      "opacity": 9.99,
      "blur": 9.99,
      "offset_x": 9.99,
      "offset_y": 9.99
    },
    "text_outline": {
      "enabled": true,
      "color": "value",
      "opacity": 9.99,
      "width": 9.99
    },
    "background_enabled": true,
    "background_color": "value",
    "background_opacity": 9.99,
    "background_blur": 9.99,
    "background_border_radius": 9.99,
    "word_highlights_enabled": true,
    "word_highlights_color": "value",
    "word_highlights_background_color": "value",
    "word_highlights_opacity": 9.99,
    "word_highlights_border_radius": 9.99,
    "word_highlights_blur": 9.99,
    "section_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "word_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "character_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "cursor_enabled": true,
    "width_pct": 9.99,
    "horizontal_placement": {
      "align": "left",
      "translate_pct": 9.99
    },
    "vertical_placement": {
      "align": "top",
      "translate_pct": 9.99
    },
    "auto_break_enabled": true,
    "max_lines_per_section": 123,
    "max_words_per_line": 123
  },
  "caption_style_template_overrides": {
    "key": {
      "template": {
        "key": "value",
        "label": "value",
        "requires_high_fps": true
      },
      "text_font": "value",
      "text_scale": 9.99,
      "text_color": "value",
      "text_align": "start",
      "text_style": "normal",
      "text_weight": "normal",
      "text_transform": "none",
      "text_blend_mode": "normal",
      "text_shadow": {
        "enabled": true,
        "color": "value",
        "opacity": 9.99,
        "blur": 9.99,
        "offset_x": 9.99,
        "offset_y": 9.99
      },
      "text_outline": {
        "enabled": true,
        "color": "value",
        "opacity": 9.99,
        "width": 9.99
      },
      "background_enabled": true,
      "background_color": "value",
      "background_opacity": 9.99,
      "background_blur": 9.99,
      "background_border_radius": 9.99,
      "word_highlights_enabled": true,
      "word_highlights_color": "value",
      "word_highlights_background_color": "value",
      "word_highlights_opacity": 9.99,
      "word_highlights_border_radius": 9.99,
      "word_highlights_blur": 9.99,
      "section_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "word_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "character_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "cursor_enabled": true,
      "width_pct": 9.99,
      "horizontal_placement": {
        "align": "left",
        "translate_pct": 9.99
      },
      "vertical_placement": {
        "align": "top",
        "translate_pct": 9.99
      },
      "auto_break_enabled": true,
      "max_lines_per_section": 123,
      "max_words_per_line": 123
    }
  },
  "public_share_id": "value",
  "aspect_ratio": "16:9",
  "agent_settings": {
    "tool_settings": {
      "key": {
        "skip_confirmation": true
      }
    }
  },
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ProjectResponse))
}

func TestProjectResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "name": "value",
  "create_date_unix": 123,
  "created_by_user_id": "value",
  "default_title_voice_ref_id": "value",
  "default_paragraph_voice_ref_id": "value",
  "default_model_id": "value",
  "last_conversion_date_unix": 123,
  "can_be_downloaded": true,
  "title": "value",
  "author": "value",
  "description": "value",
  "genres": [
    "value"
  ],
  "cover_image_url": "value",
  "target_audience": "children",
  "language": "value",
  "content_type": "value",
  "original_publication_date": "value",
  "mature_content": true,
  "isbn_number": "value",
  "volume_normalization": true,
  "state": "creating",
  "access_level": "admin",
  "fiction": "fiction",
  "quality_check_on": true,
  "quality_check_on_when_bulk_convert": true,
  "creation_meta": {
    "creation_progress": 9.99,
    "status": "draft",
    "type": "blank"
  },
  "source_type": "blank",
  "chapters_enabled": true,
  "captions_enabled": true,
  "caption_style": {
    "template": {
      "key": "value",
      "label": "value",
      "requires_high_fps": true
    },
    "text_font": "value",
    "text_scale": 9.99,
    "text_color": "value",
    "text_align": "start",
    "text_style": "normal",
    "text_weight": "normal",
    "text_transform": "none",
    "text_blend_mode": "normal",
    "text_shadow": {
      "enabled": true,
      "color": "value",
      "opacity": 9.99,
      "blur": 9.99,
      "offset_x": 9.99,
      "offset_y": 9.99
    },
    "text_outline": {
      "enabled": true,
      "color": "value",
      "opacity": 9.99,
      "width": 9.99
    },
    "background_enabled": true,
    "background_color": "value",
    "background_opacity": 9.99,
    "background_blur": 9.99,
    "background_border_radius": 9.99,
    "word_highlights_enabled": true,
    "word_highlights_color": "value",
    "word_highlights_background_color": "value",
    "word_highlights_opacity": 9.99,
    "word_highlights_border_radius": 9.99,
    "word_highlights_blur": 9.99,
    "section_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "word_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "character_animation": {
      "enter_type": "none",
      "exit_type": "none"
    },
    "cursor_enabled": true,
    "width_pct": 9.99,
    "horizontal_placement": {
      "align": "left",
      "translate_pct": 9.99
    },
    "vertical_placement": {
      "align": "top",
      "translate_pct": 9.99
    },
    "auto_break_enabled": true,
    "max_lines_per_section": 123,
    "max_words_per_line": 123
  },
  "caption_style_template_overrides": {
    "key": {
      "template": {
        "key": "value",
        "label": "value",
        "requires_high_fps": true
      },
      "text_font": "value",
      "text_scale": 9.99,
      "text_color": "value",
      "text_align": "start",
      "text_style": "normal",
      "text_weight": "normal",
      "text_transform": "none",
      "text_blend_mode": "normal",
      "text_shadow": {
        "enabled": true,
        "color": "value",
        "opacity": 9.99,
        "blur": 9.99,
        "offset_x": 9.99,
        "offset_y": 9.99
      },
      "text_outline": {
        "enabled": true,
        "color": "value",
        "opacity": 9.99,
        "width": 9.99
      },
      "background_enabled": true,
      "background_color": "value",
      "background_opacity": 9.99,
      "background_blur": 9.99,
      "background_border_radius": 9.99,
      "word_highlights_enabled": true,
      "word_highlights_color": "value",
      "word_highlights_background_color": "value",
      "word_highlights_opacity": 9.99,
      "word_highlights_border_radius": 9.99,
      "word_highlights_blur": 9.99,
      "section_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "word_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "character_animation": {
        "enter_type": "none",
        "exit_type": "none"
      },
      "cursor_enabled": true,
      "width_pct": 9.99,
      "horizontal_placement": {
        "align": "left",
        "translate_pct": 9.99
      },
      "vertical_placement": {
        "align": "top",
        "translate_pct": 9.99
      },
      "auto_break_enabled": true,
      "max_lines_per_section": 123,
      "max_words_per_line": 123
    }
  },
  "public_share_id": "value",
  "aspect_ratio": "16:9",
  "agent_settings": {
    "tool_settings": {
      "key": {
        "skip_confirmation": true
      }
    }
  },
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ProjectResponse))
}
