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

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsWidgetClient(t, 200, `{
  "agent_id": "value",
  "widget_config": {
    "variant": "tiny",
    "placement": "top-left",
    "expandable": "never",
    "avatar": {
      "type": "orb",
      "color_1": "value",
      "color_2": "value"
    },
    "feedback_mode": "none",
    "end_feedback": {
      "type": "rating"
    },
    "bg_color": "value",
    "text_color": "value",
    "btn_color": "value",
    "btn_text_color": "value",
    "border_color": "value",
    "focus_color": "value",
    "border_radius": 123,
    "btn_radius": 123,
    "action_text": "value",
    "start_call_text": "value",
    "end_call_text": "value",
    "expand_text": "value",
    "listening_text": "value",
    "speaking_text": "value",
    "shareable_page_text": "value",
    "shareable_page_show_terms": true,
    "terms_text": "value",
    "terms_html": "value",
    "terms_key": "value",
    "show_avatar_when_collapsed": true,
    "disable_banner": true,
    "override_link": "value",
    "markdown_link_allowed_hosts": [
      {
        "hostname": "value"
      }
    ],
    "markdown_link_include_www": true,
    "markdown_link_allow_http": true,
    "mic_muting_enabled": true,
    "transcript_enabled": true,
    "text_input_enabled": true,
    "conversation_mode_toggle_enabled": true,
    "default_expanded": true,
    "always_expanded": true,
    "dismissible": true,
    "show_agent_status": true,
    "show_conversation_id": true,
    "strip_audio_tags": true,
    "syntax_highlight_theme": "light",
    "text_contents": {
      "main_label": "value",
      "start_call": "value",
      "start_chat": "value",
      "new_call": "value",
      "end_call": "value",
      "mute_microphone": "value",
      "change_language": "value",
      "collapse": "value",
      "expand": "value",
      "copied": "value",
      "accept_terms": "value",
      "dismiss_terms": "value",
      "listening_status": "value",
      "speaking_status": "value",
      "connecting_status": "value",
      "chatting_status": "value",
      "queue_waiting_status": "value",
      "queue_waiting_status_short": "value",
      "input_label": "value",
      "input_placeholder": "value",
      "input_placeholder_text_only": "value",
      "input_placeholder_new_conversation": "value",
      "user_ended_conversation": "value",
      "agent_ended_conversation": "value",
      "conversation_id": "value",
      "error_occurred": "value",
      "queue_timed_out": "value",
      "copy_id": "value",
      "initiate_feedback": "value",
      "request_follow_up_feedback": "value",
      "thanks_for_feedback": "value",
      "thanks_for_feedback_details": "value",
      "follow_up_feedback_placeholder": "value",
      "submit": "value",
      "go_back": "value",
      "send_message": "value",
      "text_mode": "value",
      "voice_mode": "value",
      "switched_to_text_mode": "value",
      "switched_to_voice_mode": "value",
      "copy": "value",
      "download": "value",
      "wrap": "value",
      "agent_working": "value",
      "agent_done": "value",
      "agent_error": "value",
      "attach_file": "value",
      "remove_file": "value",
      "file_upload_error": "value",
      "file_type_unsupported": "value",
      "file_too_large": "value",
      "file_limit_reached": "value",
      "typing_indicator": "value",
      "rich_content_unavailable": "value"
    },
    "styles": {
      "base": "value",
      "base_hover": "value",
      "base_active": "value",
      "base_border": "value",
      "base_subtle": "value",
      "base_primary": "value",
      "base_error": "value",
      "accent": "value",
      "accent_hover": "value",
      "accent_active": "value",
      "accent_border": "value",
      "accent_subtle": "value",
      "accent_primary": "value",
      "overlay_padding": 9.99,
      "button_radius": 9.99,
      "input_radius": 9.99,
      "bubble_radius": 9.99,
      "sheet_radius": 9.99,
      "compact_sheet_radius": 9.99,
      "dropdown_sheet_radius": 9.99
    },
    "show_resize_button": true,
    "end_confirmation_enabled": true,
    "mobile_fullscreen": true,
    "transcript_speech_sync_enabled": true,
    "language": "value",
    "supported_language_overrides": [
      "value"
    ],
    "language_presets": {
      "key": {
        "first_message": "value",
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value"
      }
    },
    "text_only": true,
    "supports_text_only": true,
    "first_message": "value",
    "use_rtc": true,
    "file_input_config": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	conversationSignatureValue := string("value")
	request := &agents.GetWidgetRequest{
		ConversationSignature: &conversationSignatureValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWidgetClient(gotResponse, []byte(`{
  "agent_id": "value",
  "widget_config": {
    "variant": "tiny",
    "placement": "top-left",
    "expandable": "never",
    "avatar": {
      "type": "orb",
      "color_1": "value",
      "color_2": "value"
    },
    "feedback_mode": "none",
    "end_feedback": {
      "type": "rating"
    },
    "bg_color": "value",
    "text_color": "value",
    "btn_color": "value",
    "btn_text_color": "value",
    "border_color": "value",
    "focus_color": "value",
    "border_radius": 123,
    "btn_radius": 123,
    "action_text": "value",
    "start_call_text": "value",
    "end_call_text": "value",
    "expand_text": "value",
    "listening_text": "value",
    "speaking_text": "value",
    "shareable_page_text": "value",
    "shareable_page_show_terms": true,
    "terms_text": "value",
    "terms_html": "value",
    "terms_key": "value",
    "show_avatar_when_collapsed": true,
    "disable_banner": true,
    "override_link": "value",
    "markdown_link_allowed_hosts": [
      {
        "hostname": "value"
      }
    ],
    "markdown_link_include_www": true,
    "markdown_link_allow_http": true,
    "mic_muting_enabled": true,
    "transcript_enabled": true,
    "text_input_enabled": true,
    "conversation_mode_toggle_enabled": true,
    "default_expanded": true,
    "always_expanded": true,
    "dismissible": true,
    "show_agent_status": true,
    "show_conversation_id": true,
    "strip_audio_tags": true,
    "syntax_highlight_theme": "light",
    "text_contents": {
      "main_label": "value",
      "start_call": "value",
      "start_chat": "value",
      "new_call": "value",
      "end_call": "value",
      "mute_microphone": "value",
      "change_language": "value",
      "collapse": "value",
      "expand": "value",
      "copied": "value",
      "accept_terms": "value",
      "dismiss_terms": "value",
      "listening_status": "value",
      "speaking_status": "value",
      "connecting_status": "value",
      "chatting_status": "value",
      "queue_waiting_status": "value",
      "queue_waiting_status_short": "value",
      "input_label": "value",
      "input_placeholder": "value",
      "input_placeholder_text_only": "value",
      "input_placeholder_new_conversation": "value",
      "user_ended_conversation": "value",
      "agent_ended_conversation": "value",
      "conversation_id": "value",
      "error_occurred": "value",
      "queue_timed_out": "value",
      "copy_id": "value",
      "initiate_feedback": "value",
      "request_follow_up_feedback": "value",
      "thanks_for_feedback": "value",
      "thanks_for_feedback_details": "value",
      "follow_up_feedback_placeholder": "value",
      "submit": "value",
      "go_back": "value",
      "send_message": "value",
      "text_mode": "value",
      "voice_mode": "value",
      "switched_to_text_mode": "value",
      "switched_to_voice_mode": "value",
      "copy": "value",
      "download": "value",
      "wrap": "value",
      "agent_working": "value",
      "agent_done": "value",
      "agent_error": "value",
      "attach_file": "value",
      "remove_file": "value",
      "file_upload_error": "value",
      "file_type_unsupported": "value",
      "file_too_large": "value",
      "file_limit_reached": "value",
      "typing_indicator": "value",
      "rich_content_unavailable": "value"
    },
    "styles": {
      "base": "value",
      "base_hover": "value",
      "base_active": "value",
      "base_border": "value",
      "base_subtle": "value",
      "base_primary": "value",
      "base_error": "value",
      "accent": "value",
      "accent_hover": "value",
      "accent_active": "value",
      "accent_border": "value",
      "accent_subtle": "value",
      "accent_primary": "value",
      "overlay_padding": 9.99,
      "button_radius": 9.99,
      "input_radius": 9.99,
      "bubble_radius": 9.99,
      "sheet_radius": 9.99,
      "compact_sheet_radius": 9.99,
      "dropdown_sheet_radius": 9.99
    },
    "show_resize_button": true,
    "end_confirmation_enabled": true,
    "mobile_fullscreen": true,
    "transcript_speech_sync_enabled": true,
    "language": "value",
    "supported_language_overrides": [
      "value"
    ],
    "language_presets": {
      "key": {
        "first_message": "value",
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value"
      }
    },
    "text_only": true,
    "supports_text_only": true,
    "first_message": "value",
    "use_rtc": true,
    "file_input_config": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_id": "value",
  "widget_config": {
    "variant": "tiny",
    "placement": "top-left",
    "expandable": "never",
    "avatar": {
      "type": "orb",
      "color_1": "value",
      "color_2": "value"
    },
    "feedback_mode": "none",
    "end_feedback": {
      "type": "rating"
    },
    "bg_color": "value",
    "text_color": "value",
    "btn_color": "value",
    "btn_text_color": "value",
    "border_color": "value",
    "focus_color": "value",
    "border_radius": 123,
    "btn_radius": 123,
    "action_text": "value",
    "start_call_text": "value",
    "end_call_text": "value",
    "expand_text": "value",
    "listening_text": "value",
    "speaking_text": "value",
    "shareable_page_text": "value",
    "shareable_page_show_terms": true,
    "terms_text": "value",
    "terms_html": "value",
    "terms_key": "value",
    "show_avatar_when_collapsed": true,
    "disable_banner": true,
    "override_link": "value",
    "markdown_link_allowed_hosts": [
      {
        "hostname": "value"
      }
    ],
    "markdown_link_include_www": true,
    "markdown_link_allow_http": true,
    "mic_muting_enabled": true,
    "transcript_enabled": true,
    "text_input_enabled": true,
    "conversation_mode_toggle_enabled": true,
    "default_expanded": true,
    "always_expanded": true,
    "dismissible": true,
    "show_agent_status": true,
    "show_conversation_id": true,
    "strip_audio_tags": true,
    "syntax_highlight_theme": "light",
    "text_contents": {
      "main_label": "value",
      "start_call": "value",
      "start_chat": "value",
      "new_call": "value",
      "end_call": "value",
      "mute_microphone": "value",
      "change_language": "value",
      "collapse": "value",
      "expand": "value",
      "copied": "value",
      "accept_terms": "value",
      "dismiss_terms": "value",
      "listening_status": "value",
      "speaking_status": "value",
      "connecting_status": "value",
      "chatting_status": "value",
      "queue_waiting_status": "value",
      "queue_waiting_status_short": "value",
      "input_label": "value",
      "input_placeholder": "value",
      "input_placeholder_text_only": "value",
      "input_placeholder_new_conversation": "value",
      "user_ended_conversation": "value",
      "agent_ended_conversation": "value",
      "conversation_id": "value",
      "error_occurred": "value",
      "queue_timed_out": "value",
      "copy_id": "value",
      "initiate_feedback": "value",
      "request_follow_up_feedback": "value",
      "thanks_for_feedback": "value",
      "thanks_for_feedback_details": "value",
      "follow_up_feedback_placeholder": "value",
      "submit": "value",
      "go_back": "value",
      "send_message": "value",
      "text_mode": "value",
      "voice_mode": "value",
      "switched_to_text_mode": "value",
      "switched_to_voice_mode": "value",
      "copy": "value",
      "download": "value",
      "wrap": "value",
      "agent_working": "value",
      "agent_done": "value",
      "agent_error": "value",
      "attach_file": "value",
      "remove_file": "value",
      "file_upload_error": "value",
      "file_type_unsupported": "value",
      "file_too_large": "value",
      "file_limit_reached": "value",
      "typing_indicator": "value",
      "rich_content_unavailable": "value"
    },
    "styles": {
      "base": "value",
      "base_hover": "value",
      "base_active": "value",
      "base_border": "value",
      "base_subtle": "value",
      "base_primary": "value",
      "base_error": "value",
      "accent": "value",
      "accent_hover": "value",
      "accent_active": "value",
      "accent_border": "value",
      "accent_subtle": "value",
      "accent_primary": "value",
      "overlay_padding": 9.99,
      "button_radius": 9.99,
      "input_radius": 9.99,
      "bubble_radius": 9.99,
      "sheet_radius": 9.99,
      "compact_sheet_radius": 9.99,
      "dropdown_sheet_radius": 9.99
    },
    "show_resize_button": true,
    "end_confirmation_enabled": true,
    "mobile_fullscreen": true,
    "transcript_speech_sync_enabled": true,
    "language": "value",
    "supported_language_overrides": [
      "value"
    ],
    "language_presets": {
      "key": {
        "first_message": "value",
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value"
      }
    },
    "text_only": true,
    "supports_text_only": true,
    "first_message": "value",
    "use_rtc": true,
    "file_input_config": {
      "enabled": true,
      "max_files_in_memory": 30,
      "max_files_per_conversation": 123
    }
  }
}`)
	}

	req := receiveWireAgentsWidgetClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/widget" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/widget")
	}
	if got := req.Query.Get("conversation_signature"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "conversation_signature", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWidgetClient(t, 422, `{
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

	conversationSignatureValue := string("value")
	request := &agents.GetWidgetRequest{
		ConversationSignature: &conversationSignatureValue,
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

type wireRequestAgentsWidgetClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsWidgetClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsWidgetClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsWidgetClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsWidgetClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsWidgetClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsWidgetClient(t *testing.T, recorded <-chan wireRequestAgentsWidgetClient) wireRequestAgentsWidgetClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsWidgetClient{}
	}
}

// wireJSONContainsAgentsWidgetClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsWidgetClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsWidgetClient(got, want)
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

// setWireValueAgentsWidgetClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsWidgetClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsWidgetClient(got, want []byte) bool {
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
