package tiers

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
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionTiers(t, 200, `[
  {
    "tier_id": "value",
    "display_name": "value",
    "tier_economics": {
      "monthly_price_usd": 9.99,
      "annual_price_usd": 9.99,
      "included_minutes_per_month": 9.99,
      "overage_price_per_minute_usd": 9.99,
      "max_phone_numbers": 123,
      "concurrent_calls": 123,
      "max_active_locations": 123,
      "max_customer_facing_agents": 123,
      "max_team_members": 123,
      "assistant_daily_messages": 123
    }
  }
]`)

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
	if !wireJSONBytesEqualReceptionTiers(gotResponse, []byte(`[
  {
    "tier_id": "value",
    "display_name": "value",
    "tier_economics": {
      "monthly_price_usd": 9.99,
      "annual_price_usd": 9.99,
      "included_minutes_per_month": 9.99,
      "overage_price_per_minute_usd": 9.99,
      "max_phone_numbers": 123,
      "concurrent_calls": 123,
      "max_active_locations": 123,
      "max_customer_facing_agents": 123,
      "max_team_members": 123,
      "assistant_daily_messages": 123
    }
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "tier_id": "value",
    "display_name": "value",
    "tier_economics": {
      "monthly_price_usd": 9.99,
      "annual_price_usd": 9.99,
      "included_minutes_per_month": 9.99,
      "overage_price_per_minute_usd": 9.99,
      "max_phone_numbers": 123,
      "concurrent_calls": 123,
      "max_active_locations": 123,
      "max_customer_facing_agents": 123,
      "max_team_members": 123,
      "assistant_daily_messages": 123
    }
  }
]`)
	}

	req := receiveWireReceptionTiers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/reseller/tier-ladder" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/tier-ladder")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionTiers(t, 422, `{
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

type wireRequestReceptionTiers struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionTiers records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionTiers(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionTiers) {
	t.Helper()
	recorded := make(chan wireRequestReceptionTiers, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionTiers{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionTiers(t *testing.T, recorded <-chan wireRequestReceptionTiers) wireRequestReceptionTiers {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionTiers{}
	}
}

// wireJSONContainsReceptionTiers reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionTiers(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionTiers(got, want)
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

// setWireValueReceptionTiers sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionTiers(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionTiers(got, want []byte) bool {
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
