package compositionplan

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
	music "github.com/elevenlabs/elevenlabs-go/music"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireMusicCompositionplan(t, 200, `{
  "positive_global_styles": [
    "value"
  ],
  "negative_global_styles": [
    "value"
  ],
  "sections": [
    {
      "section_name": "value",
      "positive_local_styles": [
        "value"
      ],
      "negative_local_styles": [
        "value"
      ],
      "duration_ms": 3000,
      "lines": [
        "value"
      ],
      "source_from": {
        "song_id": "value",
        "range": {
          "start_ms": 123,
          "end_ms": 123
        },
        "negative_ranges": [
          {
            "start_ms": 123,
            "end_ms": 123
          }
        ]
      }
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request music.BodyGenerateCompositionPlanV1MusicPlanPost
	if err := json.Unmarshal([]byte(`{
  "enable_logging": true,
  "prompt": "value",
  "music_length_ms": 3000,
  "source_composition_plan": {
    "positive_global_styles": [
      "value"
    ],
    "negative_global_styles": [
      "value"
    ],
    "sections": [
      {
        "section_name": "value",
        "positive_local_styles": [
          "value"
        ],
        "negative_local_styles": [
          "value"
        ],
        "duration_ms": 3000,
        "lines": [
          "value"
        ],
        "source_from": {
          "song_id": "value",
          "range": {
            "start_ms": 123,
            "end_ms": 123
          },
          "negative_ranges": [
            {
              "start_ms": 123,
              "end_ms": 123
            }
          ]
        }
      }
    ]
  },
  "model_id": "music_v1"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualMusicCompositionplan(gotResponse, []byte(`{
  "positive_global_styles": [
    "value"
  ],
  "negative_global_styles": [
    "value"
  ],
  "sections": [
    {
      "section_name": "value",
      "positive_local_styles": [
        "value"
      ],
      "negative_local_styles": [
        "value"
      ],
      "duration_ms": 3000,
      "lines": [
        "value"
      ],
      "source_from": {
        "song_id": "value",
        "range": {
          "start_ms": 123,
          "end_ms": 123
        },
        "negative_ranges": [
          {
            "start_ms": 123,
            "end_ms": 123
          }
        ]
      }
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "positive_global_styles": [
    "value"
  ],
  "negative_global_styles": [
    "value"
  ],
  "sections": [
    {
      "section_name": "value",
      "positive_local_styles": [
        "value"
      ],
      "negative_local_styles": [
        "value"
      ],
      "duration_ms": 3000,
      "lines": [
        "value"
      ],
      "source_from": {
        "song_id": "value",
        "range": {
          "start_ms": 123,
          "end_ms": 123
        },
        "negative_ranges": [
          {
            "start_ms": 123,
            "end_ms": 123
          }
        ]
      }
    }
  ]
}`)
	}

	req := receiveWireMusicCompositionplan(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/music/plan" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/music/plan")
	}
	if got := req.Query.Get("enable_logging"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "enable_logging", got, "true")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsMusicCompositionplan(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireMusicCompositionplan(t, 422, `{
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

	var request music.BodyGenerateCompositionPlanV1MusicPlanPost
	if err := json.Unmarshal([]byte(`{
  "enable_logging": true,
  "prompt": "value",
  "music_length_ms": 3000,
  "source_composition_plan": {
    "positive_global_styles": [
      "value"
    ],
    "negative_global_styles": [
      "value"
    ],
    "sections": [
      {
        "section_name": "value",
        "positive_local_styles": [
          "value"
        ],
        "negative_local_styles": [
          "value"
        ],
        "duration_ms": 3000,
        "lines": [
          "value"
        ],
        "source_from": {
          "song_id": "value",
          "range": {
            "start_ms": 123,
            "end_ms": 123
          },
          "negative_ranges": [
            {
              "start_ms": 123,
              "end_ms": 123
            }
          ]
        }
      }
    ]
  },
  "model_id": "music_v1"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	_, err := client.Create(context.Background(), &request)
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

type wireRequestMusicCompositionplan struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireMusicCompositionplan records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireMusicCompositionplan(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestMusicCompositionplan) {
	t.Helper()
	recorded := make(chan wireRequestMusicCompositionplan, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestMusicCompositionplan{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireMusicCompositionplan(t *testing.T, recorded <-chan wireRequestMusicCompositionplan) wireRequestMusicCompositionplan {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestMusicCompositionplan{}
	}
}

// wireJSONContainsMusicCompositionplan reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsMusicCompositionplan(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualMusicCompositionplan(got, want)
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

// setWireValueMusicCompositionplan sets a named string, int or float field, or the value a pointer field points to.
func setWireValueMusicCompositionplan(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualMusicCompositionplan(got, want []byte) bool {
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
