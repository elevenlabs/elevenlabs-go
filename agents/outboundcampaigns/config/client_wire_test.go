package config

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

	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsConfig(t, 200, `{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsConfig(gotResponse, []byte(`{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)
	}

	req := receiveWireAgentsOutboundcampaignsConfig(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/config" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/config")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsConfig(t, 422, `{
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

	_, err := client.Get(context.Background())
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

	server, recorded := serveWireAgentsOutboundcampaignsConfig(t, 200, `{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request outboundcampaigns.ReplaceWorkspaceConfigRequest
	if err := json.Unmarshal([]byte(`{
  "expected_version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsConfig(gotResponse, []byte(`{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`)
	}

	req := receiveWireAgentsOutboundcampaignsConfig(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/convai/outbound-campaigns/config" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/config")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsOutboundcampaignsConfig(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsConfig(t, 422, `{
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

	var request outboundcampaigns.ReplaceWorkspaceConfigRequest
	if err := json.Unmarshal([]byte(`{
  "expected_version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), &request)
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

type wireRequestAgentsOutboundcampaignsConfig struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsOutboundcampaignsConfig records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsOutboundcampaignsConfig(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsOutboundcampaignsConfig) {
	t.Helper()
	recorded := make(chan wireRequestAgentsOutboundcampaignsConfig, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsOutboundcampaignsConfig{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsOutboundcampaignsConfig(t *testing.T, recorded <-chan wireRequestAgentsOutboundcampaignsConfig) wireRequestAgentsOutboundcampaignsConfig {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsOutboundcampaignsConfig{}
	}
}

// wireJSONContainsAgentsOutboundcampaignsConfig reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsOutboundcampaignsConfig(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsOutboundcampaignsConfig(got, want)
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

// setWireValueAgentsOutboundcampaignsConfig sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsOutboundcampaignsConfig(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsOutboundcampaignsConfig(got, want []byte) bool {
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
