package whatsapp

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

func TestOutboundCallWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsWhatsapp(t, 200, `{
  "success": true,
  "message": "value",
  "conversation_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost
	if err := json.Unmarshal([]byte(`{
  "whatsapp_phone_number_id": "value",
  "whatsapp_user_id": "value",
  "whatsapp_call_permission_request_template_name": "value",
  "whatsapp_call_permission_request_template_language_code": "value",
  "agent_id": "value",
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
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
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.OutboundCall(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsapp(gotResponse, []byte(`{
  "success": true,
  "message": "value",
  "conversation_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "success": true,
  "message": "value",
  "conversation_id": "value"
}`)
	}

	req := receiveWireAgentsWhatsapp(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/whatsapp/outbound-call" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp/outbound-call")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsWhatsapp(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestOutboundCallWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsapp(t, 422, `{
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

	var request agents.BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost
	if err := json.Unmarshal([]byte(`{
  "whatsapp_phone_number_id": "value",
  "whatsapp_user_id": "value",
  "whatsapp_call_permission_request_template_name": "value",
  "whatsapp_call_permission_request_template_language_code": "value",
  "agent_id": "value",
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
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
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.OutboundCall(context.Background(), &request)
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

func TestOutboundMessageWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsWhatsapp(t, 200, `{
  "conversation_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost
	if err := json.Unmarshal([]byte(`{
  "whatsapp_phone_number_id": "value",
  "whatsapp_user_id": "value",
  "template_name": "value",
  "template_language_code": "value",
  "template_params": [
    {
      "type": "header",
      "parameters": [
        {
          "parameter_name": "value",
          "type": "text",
          "text": "value"
        }
      ]
    }
  ],
  "agent_id": "value",
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
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
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.OutboundMessage(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsapp(gotResponse, []byte(`{
  "conversation_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_id": "value"
}`)
	}

	req := receiveWireAgentsWhatsapp(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/whatsapp/outbound-message" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp/outbound-message")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsWhatsapp(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestOutboundMessageWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsapp(t, 422, `{
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

	var request agents.BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost
	if err := json.Unmarshal([]byte(`{
  "whatsapp_phone_number_id": "value",
  "whatsapp_user_id": "value",
  "template_name": "value",
  "template_language_code": "value",
  "template_params": [
    {
      "type": "header",
      "parameters": [
        {
          "parameter_name": "value",
          "type": "text",
          "text": "value"
        }
      ]
    }
  ],
  "agent_id": "value",
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
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
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.OutboundMessage(context.Background(), &request)
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

type wireRequestAgentsWhatsapp struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsWhatsapp records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsWhatsapp(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsWhatsapp) {
	t.Helper()
	recorded := make(chan wireRequestAgentsWhatsapp, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsWhatsapp{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsWhatsapp(t *testing.T, recorded <-chan wireRequestAgentsWhatsapp) wireRequestAgentsWhatsapp {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsWhatsapp{}
	}
}

// wireJSONContainsAgentsWhatsapp reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsWhatsapp(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsWhatsapp(got, want)
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

// setWireValueAgentsWhatsapp sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsWhatsapp(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsWhatsapp(got, want []byte) bool {
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
