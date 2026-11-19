package llmusage

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

func TestCalculateForAgentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsLlmusage(t, 200, `{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.LlmUsageCalculatorRequestModel
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CalculateForAgent(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsLlmusage(gotResponse, []byte(`{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)
	}

	req := receiveWireAgentsLlmusage(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agent/value/llm-usage/calculate" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent/value/llm-usage/calculate")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsLlmusage(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCalculateForAgentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsLlmusage(t, 422, `{
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

	var request agents.LlmUsageCalculatorRequestModel
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CalculateForAgent(context.Background(), string("value"), &request)
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

func TestCalculateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsLlmusage(t, 200, `{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.LlmUsageCalculatorPublicRequestModel
	if err := json.Unmarshal([]byte(`{
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Calculate(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsLlmusage(gotResponse, []byte(`{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`)
	}

	req := receiveWireAgentsLlmusage(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/llm-usage/calculate" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/llm-usage/calculate")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsLlmusage(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCalculateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsLlmusage(t, 422, `{
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

	var request agents.LlmUsageCalculatorPublicRequestModel
	if err := json.Unmarshal([]byte(`{
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Calculate(context.Background(), &request)
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

type wireRequestAgentsLlmusage struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsLlmusage records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsLlmusage(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsLlmusage) {
	t.Helper()
	recorded := make(chan wireRequestAgentsLlmusage, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsLlmusage{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsLlmusage(t *testing.T, recorded <-chan wireRequestAgentsLlmusage) wireRequestAgentsLlmusage {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsLlmusage{}
	}
}

// wireJSONContainsAgentsLlmusage reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsLlmusage(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsLlmusage(got, want)
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

// setWireValueAgentsLlmusage sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsLlmusage(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsLlmusage(got, want []byte) bool {
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
