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
	studio "github.com/elevenlabs/elevenlabs-go/studio"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
  "projects": [
    {
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
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
  "projects": [
    {
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
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "projects": [
    {
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
    }
  ]
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

	_, err := client.List(context.Background())
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

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
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
  "quality_preset": "standard",
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ],
  "pronunciation_dictionary_versions": [
    {
      "version_id": "value",
      "version_rules_num": 123,
      "pronunciation_dictionary_id": "value",
      "dictionary_name": "value",
      "version_name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123
    }
  ],
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "apply_text_normalization": "auto",
  "experimental": {
    "key": {}
  },
  "assets": [
    {
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
    }
  ],
  "voices": [
    {
      "project_voice_ref_id": "value",
      "voice_id": "value",
      "alias": "value",
      "stability": 9.99,
      "similarity_boost": 9.99,
      "style": 9.99,
      "is_pinned": true,
      "use_speaker_boost": true,
      "volume_gain": 9.99,
      "speed": 9.99
    }
  ],
  "base_voices": [
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
  "publishing_read": {
    "read_id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "word_count": 123,
    "char_count": 123,
    "chapters": [
      {
        "chapter_name": "value",
        "word_count": 123,
        "char_count": 123,
        "starting_char_offset": 123,
        "has_parsed_html": true,
        "has_summary": true,
        "duration_seconds": 9.99,
        "file_number": "value",
        "is_fallback_name": true,
        "chapter_id": "value"
      }
    ],
    "title": "value",
    "author": "value",
    "description": "value",
    "article_image_url": "value",
    "language": "value",
    "locale": "value",
    "display_mode": "text",
    "genre": [
      "Fantasy"
    ],
    "fiction": "value",
    "content_type": "value",
    "original_file_type": "value",
    "target_audience": "children",
    "mature_content": true,
    "safesearch_adult": true,
    "origin": "value",
    "publication_date": "value",
    "isbn": "value",
    "ean": "value",
    "legal_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "content_guidelines_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "last_updated_from_project_unix": 123,
    "publishing_project_id": "value",
    "publishing_state": "value",
    "publisher_profile_id": "value",
    "quality_score": 123,
    "publisher": "value",
    "copyright": "value",
    "subtitle": "value",
    "distribution_territories": [
      "value"
    ],
    "edition": "value",
    "contributors": [
      {
        "name": "value",
        "role": "value",
        "bio": "value",
        "profile_id": "value"
      }
    ],
    "payout_type": "none",
    "list_price": 9.99,
    "currency": "usd",
    "original_audio_project_export_id": "value",
    "original_audio_document_id": "value",
    "series_id": "value",
    "volume": 123,
    "published_at_unix": 123,
    "read_slug": "value",
    "preview_audio_object": {
      "voice_id": "value",
      "text": "value",
      "audio_url": "value",
      "hls_manifest_url": "value",
      "dash_manifest_url": "value",
      "is_auto_generated": true,
      "generated_at_unix": 123
    },
    "sample_config": {
      "is_sample": true,
      "parent_id": "value",
      "parent_type": "read",
      "chapter_ids": [
        "value"
      ]
    },
    "review": {
      "review_status": "approved",
      "reviewed_at_unix": 123,
      "reviewed_by": "value",
      "reject_reasons": [
        "lacks_structure"
      ],
      "scores_breakdown": {
        "key": 123
      },
      "rejected_details": "value",
      "explanation": "value"
    },
    "voice_id": "value",
    "can_use_assistant": true,
    "is_voice_changer_on": true,
    "restricted_to_user_email_domains": [
      "value"
    ]
  },
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	shareIdValue := string("value")
	request := &studio.GetProjectsRequest{
		ShareId: &shareIdValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
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
  "quality_preset": "standard",
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ],
  "pronunciation_dictionary_versions": [
    {
      "version_id": "value",
      "version_rules_num": 123,
      "pronunciation_dictionary_id": "value",
      "dictionary_name": "value",
      "version_name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123
    }
  ],
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "apply_text_normalization": "auto",
  "experimental": {
    "key": {}
  },
  "assets": [
    {
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
    }
  ],
  "voices": [
    {
      "project_voice_ref_id": "value",
      "voice_id": "value",
      "alias": "value",
      "stability": 9.99,
      "similarity_boost": 9.99,
      "style": 9.99,
      "is_pinned": true,
      "use_speaker_boost": true,
      "volume_gain": 9.99,
      "speed": 9.99
    }
  ],
  "base_voices": [
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
  "publishing_read": {
    "read_id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "word_count": 123,
    "char_count": 123,
    "chapters": [
      {
        "chapter_name": "value",
        "word_count": 123,
        "char_count": 123,
        "starting_char_offset": 123,
        "has_parsed_html": true,
        "has_summary": true,
        "duration_seconds": 9.99,
        "file_number": "value",
        "is_fallback_name": true,
        "chapter_id": "value"
      }
    ],
    "title": "value",
    "author": "value",
    "description": "value",
    "article_image_url": "value",
    "language": "value",
    "locale": "value",
    "display_mode": "text",
    "genre": [
      "Fantasy"
    ],
    "fiction": "value",
    "content_type": "value",
    "original_file_type": "value",
    "target_audience": "children",
    "mature_content": true,
    "safesearch_adult": true,
    "origin": "value",
    "publication_date": "value",
    "isbn": "value",
    "ean": "value",
    "legal_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "content_guidelines_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "last_updated_from_project_unix": 123,
    "publishing_project_id": "value",
    "publishing_state": "value",
    "publisher_profile_id": "value",
    "quality_score": 123,
    "publisher": "value",
    "copyright": "value",
    "subtitle": "value",
    "distribution_territories": [
      "value"
    ],
    "edition": "value",
    "contributors": [
      {
        "name": "value",
        "role": "value",
        "bio": "value",
        "profile_id": "value"
      }
    ],
    "payout_type": "none",
    "list_price": 9.99,
    "currency": "usd",
    "original_audio_project_export_id": "value",
    "original_audio_document_id": "value",
    "series_id": "value",
    "volume": 123,
    "published_at_unix": 123,
    "read_slug": "value",
    "preview_audio_object": {
      "voice_id": "value",
      "text": "value",
      "audio_url": "value",
      "hls_manifest_url": "value",
      "dash_manifest_url": "value",
      "is_auto_generated": true,
      "generated_at_unix": 123
    },
    "sample_config": {
      "is_sample": true,
      "parent_id": "value",
      "parent_type": "read",
      "chapter_ids": [
        "value"
      ]
    },
    "review": {
      "review_status": "approved",
      "reviewed_at_unix": 123,
      "reviewed_by": "value",
      "reject_reasons": [
        "lacks_structure"
      ],
      "scores_breakdown": {
        "key": 123
      },
      "rejected_details": "value",
      "explanation": "value"
    },
    "voice_id": "value",
    "can_use_assistant": true,
    "is_voice_changer_on": true,
    "restricted_to_user_email_domains": [
      "value"
    ]
  },
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
  "quality_preset": "standard",
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ],
  "pronunciation_dictionary_versions": [
    {
      "version_id": "value",
      "version_rules_num": 123,
      "pronunciation_dictionary_id": "value",
      "dictionary_name": "value",
      "version_name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123
    }
  ],
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "apply_text_normalization": "auto",
  "experimental": {
    "key": {}
  },
  "assets": [
    {
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
    }
  ],
  "voices": [
    {
      "project_voice_ref_id": "value",
      "voice_id": "value",
      "alias": "value",
      "stability": 9.99,
      "similarity_boost": 9.99,
      "style": 9.99,
      "is_pinned": true,
      "use_speaker_boost": true,
      "volume_gain": 9.99,
      "speed": 9.99
    }
  ],
  "base_voices": [
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
  "publishing_read": {
    "read_id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "word_count": 123,
    "char_count": 123,
    "chapters": [
      {
        "chapter_name": "value",
        "word_count": 123,
        "char_count": 123,
        "starting_char_offset": 123,
        "has_parsed_html": true,
        "has_summary": true,
        "duration_seconds": 9.99,
        "file_number": "value",
        "is_fallback_name": true,
        "chapter_id": "value"
      }
    ],
    "title": "value",
    "author": "value",
    "description": "value",
    "article_image_url": "value",
    "language": "value",
    "locale": "value",
    "display_mode": "text",
    "genre": [
      "Fantasy"
    ],
    "fiction": "value",
    "content_type": "value",
    "original_file_type": "value",
    "target_audience": "children",
    "mature_content": true,
    "safesearch_adult": true,
    "origin": "value",
    "publication_date": "value",
    "isbn": "value",
    "ean": "value",
    "legal_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "content_guidelines_terms": {
      "terms": "value",
      "start_date": "value",
      "end_date": "value"
    },
    "last_updated_from_project_unix": 123,
    "publishing_project_id": "value",
    "publishing_state": "value",
    "publisher_profile_id": "value",
    "quality_score": 123,
    "publisher": "value",
    "copyright": "value",
    "subtitle": "value",
    "distribution_territories": [
      "value"
    ],
    "edition": "value",
    "contributors": [
      {
        "name": "value",
        "role": "value",
        "bio": "value",
        "profile_id": "value"
      }
    ],
    "payout_type": "none",
    "list_price": 9.99,
    "currency": "usd",
    "original_audio_project_export_id": "value",
    "original_audio_document_id": "value",
    "series_id": "value",
    "volume": 123,
    "published_at_unix": 123,
    "read_slug": "value",
    "preview_audio_object": {
      "voice_id": "value",
      "text": "value",
      "audio_url": "value",
      "hls_manifest_url": "value",
      "dash_manifest_url": "value",
      "is_auto_generated": true,
      "generated_at_unix": 123
    },
    "sample_config": {
      "is_sample": true,
      "parent_id": "value",
      "parent_type": "read",
      "chapter_ids": [
        "value"
      ]
    },
    "review": {
      "review_status": "approved",
      "reviewed_at_unix": 123,
      "reviewed_by": "value",
      "reject_reasons": [
        "lacks_structure"
      ],
      "scores_breakdown": {
        "key": 123
      },
      "rejected_details": "value",
      "explanation": "value"
    },
    "voice_id": "value",
    "can_use_assistant": true,
    "is_voice_changer_on": true,
    "restricted_to_user_email_domains": [
      "value"
    ]
  },
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value"
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value")
	}
	if got := req.Query.Get("share_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "share_id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

	shareIdValue := string("value")
	request := &studio.GetProjectsRequest{
		ShareId: &shareIdValue,
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
  "project": {
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
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request studio.BodyUpdateStudioProjectV1StudioProjectsProjectIdPost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "name": "value",
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value",
  "title": "value",
  "author": "value",
  "isbn_number": "value",
  "volume_normalization": true
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
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
  "project": {
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
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "project": {
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
  }
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/projects/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsStudioProjectsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

	var request studio.BodyUpdateStudioProjectV1StudioProjectsProjectIdPost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "name": "value",
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value",
  "title": "value",
  "author": "value",
  "isbn_number": "value",
  "volume_normalization": true
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

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
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
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/studio/projects/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

func TestConvertWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Convert(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/projects/value/convert" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/convert")
	}
}

func TestConvertWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

	_, err := client.Convert(context.Background(), string("value"))
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

func TestGetMutedTracksWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsClient(t, 200, `{
  "chapter_ids": [
    "value"
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetMutedTracks(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsClient(gotResponse, []byte(`{
  "chapter_ids": [
    "value"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapter_ids": [
    "value"
  ]
}`)
	}

	req := receiveWireStudioProjectsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value/muted-tracks" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/muted-tracks")
	}
}

func TestGetMutedTracksWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsClient(t, 422, `{
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

	_, err := client.GetMutedTracks(context.Background(), string("value"))
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

type wireRequestStudioProjectsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireStudioProjectsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireStudioProjectsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestStudioProjectsClient) {
	t.Helper()
	recorded := make(chan wireRequestStudioProjectsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestStudioProjectsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireStudioProjectsClient(t *testing.T, recorded <-chan wireRequestStudioProjectsClient) wireRequestStudioProjectsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestStudioProjectsClient{}
	}
}

// wireJSONContainsStudioProjectsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsStudioProjectsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualStudioProjectsClient(got, want)
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

// setWireValueStudioProjectsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueStudioProjectsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualStudioProjectsClient(got, want []byte) bool {
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
