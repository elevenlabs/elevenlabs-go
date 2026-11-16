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

func TestCreatePodcastWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioClient(t, 200, `{
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

	var request types.BodyCreatePodcastV1StudioPodcastsPost
	if err := json.Unmarshal([]byte(`{
  "safety-identifier": "value",
  "model_id": "value",
  "mode": {
    "type": "conversation",
    "conversation": {
      "host_voice_id": "value",
      "guest_voice_id": "value"
    }
  },
  "source": {
    "type": "text",
    "text": "value"
  },
  "quality_preset": "standard",
  "duration_scale": "short",
  "language": "va",
  "intro": "value",
  "outro": "value",
  "instructions_prompt": "value",
  "highlights": [
    "valuexxxxx"
  ],
  "callback_url": "value",
  "apply_text_normalization": "auto"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	safetyIdentifierValue := string("value")
	request.SafetyIdentifier = &safetyIdentifierValue
	result, err := client.CreatePodcast(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioClient(gotResponse, []byte(`{
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

	req := receiveWireStudioClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/podcasts" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/podcasts")
	}
	if got := req.Header.Get("safety-identifier"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "safety-identifier", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsStudioClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreatePodcastWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioClient(t, 422, `{
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

	var request types.BodyCreatePodcastV1StudioPodcastsPost
	if err := json.Unmarshal([]byte(`{
  "safety-identifier": "value",
  "model_id": "value",
  "mode": {
    "type": "conversation",
    "conversation": {
      "host_voice_id": "value",
      "guest_voice_id": "value"
    }
  },
  "source": {
    "type": "text",
    "text": "value"
  },
  "quality_preset": "standard",
  "duration_scale": "short",
  "language": "va",
  "intro": "value",
  "outro": "value",
  "instructions_prompt": "value",
  "highlights": [
    "valuexxxxx"
  ],
  "callback_url": "value",
  "apply_text_normalization": "auto"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	safetyIdentifierValue := string("value")
	request.SafetyIdentifier = &safetyIdentifierValue
	_, err := client.CreatePodcast(context.Background(), &request)
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

type wireRequestStudioClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireStudioClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireStudioClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestStudioClient) {
	t.Helper()
	recorded := make(chan wireRequestStudioClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestStudioClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireStudioClient(t *testing.T, recorded <-chan wireRequestStudioClient) wireRequestStudioClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestStudioClient{}
	}
}

// wireJSONContainsStudioClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsStudioClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualStudioClient(got, want)
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

// setWireValueStudioClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueStudioClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualStudioClient(got, want []byte) bool {
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
