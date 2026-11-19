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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireServiceaccountsClient(t, 200, `{
  "service-accounts": [
    {
      "service_account_user_id": "value",
      "name": "value",
      "created_at_unix": 123,
      "api-keys": [
        {
          "name": "value",
          "hint": "value",
          "key_id": "value",
          "service_account_user_id": "value",
          "created_at_unix": 123,
          "is_disabled": true,
          "permissions": [
            "text_to_speech"
          ],
          "disable_reason": "trial_ended",
          "character_limit": 123,
          "character_count": 123,
          "hashed_xi_api_key": "value",
          "allowed_ips": [
            "value"
          ],
          "third_party_disable_allowed": true,
          "platform_limits": {
            "credits": {
              "limit": 123,
              "usage": 123
            },
            "pvc": {
              "limit": 123
            },
            "concurrency": {
              "limit": 123
            },
            "dubbing_concurrency": {
              "limit": 123
            },
            "music_concurrency": {
              "limit": 123
            }
          },
          "expires_at_unix": 123
        }
      ],
      "default_sharing_groups": [
        {
          "group": {
            "name": "value",
            "id": "value",
            "members": [
              "value"
            ],
            "permissions": [
              "text_to_speech"
            ],
            "group_usage_limit": 123,
            "group_pvc_limit": 123,
            "group_concurrency_limit": 123,
            "group_dubbing_concurrency_limit": 123,
            "group_music_concurrency_limit": 123,
            "character_count": 123,
            "is_scim_synced": true,
            "scim_group": {
              "scim_external_id": "value",
              "display_name": "value",
              "created_at_unix": 123,
              "updated_at_unix": 123,
              "seat_type": "workspace_admin"
            },
            "scim_frozen": true
          },
          "permission_level": "admin"
        }
      ]
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
	if !wireJSONBytesEqualServiceaccountsClient(gotResponse, []byte(`{
  "service-accounts": [
    {
      "service_account_user_id": "value",
      "name": "value",
      "created_at_unix": 123,
      "api-keys": [
        {
          "name": "value",
          "hint": "value",
          "key_id": "value",
          "service_account_user_id": "value",
          "created_at_unix": 123,
          "is_disabled": true,
          "permissions": [
            "text_to_speech"
          ],
          "disable_reason": "trial_ended",
          "character_limit": 123,
          "character_count": 123,
          "hashed_xi_api_key": "value",
          "allowed_ips": [
            "value"
          ],
          "third_party_disable_allowed": true,
          "platform_limits": {
            "credits": {
              "limit": 123,
              "usage": 123
            },
            "pvc": {
              "limit": 123
            },
            "concurrency": {
              "limit": 123
            },
            "dubbing_concurrency": {
              "limit": 123
            },
            "music_concurrency": {
              "limit": 123
            }
          },
          "expires_at_unix": 123
        }
      ],
      "default_sharing_groups": [
        {
          "group": {
            "name": "value",
            "id": "value",
            "members": [
              "value"
            ],
            "permissions": [
              "text_to_speech"
            ],
            "group_usage_limit": 123,
            "group_pvc_limit": 123,
            "group_concurrency_limit": 123,
            "group_dubbing_concurrency_limit": 123,
            "group_music_concurrency_limit": 123,
            "character_count": 123,
            "is_scim_synced": true,
            "scim_group": {
              "scim_external_id": "value",
              "display_name": "value",
              "created_at_unix": 123,
              "updated_at_unix": 123,
              "seat_type": "workspace_admin"
            },
            "scim_frozen": true
          },
          "permission_level": "admin"
        }
      ]
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "service-accounts": [
    {
      "service_account_user_id": "value",
      "name": "value",
      "created_at_unix": 123,
      "api-keys": [
        {
          "name": "value",
          "hint": "value",
          "key_id": "value",
          "service_account_user_id": "value",
          "created_at_unix": 123,
          "is_disabled": true,
          "permissions": [
            "text_to_speech"
          ],
          "disable_reason": "trial_ended",
          "character_limit": 123,
          "character_count": 123,
          "hashed_xi_api_key": "value",
          "allowed_ips": [
            "value"
          ],
          "third_party_disable_allowed": true,
          "platform_limits": {
            "credits": {
              "limit": 123,
              "usage": 123
            },
            "pvc": {
              "limit": 123
            },
            "concurrency": {
              "limit": 123
            },
            "dubbing_concurrency": {
              "limit": 123
            },
            "music_concurrency": {
              "limit": 123
            }
          },
          "expires_at_unix": 123
        }
      ],
      "default_sharing_groups": [
        {
          "group": {
            "name": "value",
            "id": "value",
            "members": [
              "value"
            ],
            "permissions": [
              "text_to_speech"
            ],
            "group_usage_limit": 123,
            "group_pvc_limit": 123,
            "group_concurrency_limit": 123,
            "group_dubbing_concurrency_limit": 123,
            "group_music_concurrency_limit": 123,
            "character_count": 123,
            "is_scim_synced": true,
            "scim_group": {
              "scim_external_id": "value",
              "display_name": "value",
              "created_at_unix": 123,
              "updated_at_unix": 123,
              "seat_type": "workspace_admin"
            },
            "scim_frozen": true
          },
          "permission_level": "admin"
        }
      ]
    }
  ]
}`)
	}

	req := receiveWireServiceaccountsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/service-accounts" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/service-accounts")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireServiceaccountsClient(t, 422, `{
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireServiceaccountsClient(t, 200, `{
  "service-account-user-id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyCreateServiceAccountV1ServiceAccountsPost
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "default_sharing_groups": [
    {
      "group_id": "value",
      "permission_level": "admin"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualServiceaccountsClient(gotResponse, []byte(`{
  "service-account-user-id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "service-account-user-id": "value"
}`)
	}

	req := receiveWireServiceaccountsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/service-accounts" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/service-accounts")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsServiceaccountsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireServiceaccountsClient(t, 422, `{
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

	var request types.BodyCreateServiceAccountV1ServiceAccountsPost
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "default_sharing_groups": [
    {
      "group_id": "value",
      "permission_level": "admin"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
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

type wireRequestServiceaccountsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireServiceaccountsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireServiceaccountsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestServiceaccountsClient) {
	t.Helper()
	recorded := make(chan wireRequestServiceaccountsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestServiceaccountsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireServiceaccountsClient(t *testing.T, recorded <-chan wireRequestServiceaccountsClient) wireRequestServiceaccountsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestServiceaccountsClient{}
	}
}

// wireJSONContainsServiceaccountsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsServiceaccountsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualServiceaccountsClient(got, want)
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

// setWireValueServiceaccountsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueServiceaccountsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualServiceaccountsClient(got, want []byte) bool {
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
