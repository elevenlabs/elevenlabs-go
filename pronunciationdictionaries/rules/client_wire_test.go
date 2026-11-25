package rules

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
	pronunciationdictionaries "github.com/elevenlabs/elevenlabs-go/pronunciationdictionaries"
)

func TestSetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWirePronunciationdictionariesRules(t, 200, `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request pronunciationdictionaries.BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Set(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesRules(gotResponse, []byte(`{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)
	}

	req := receiveWirePronunciationdictionariesRules(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/pronunciation-dictionaries/value/set-rules" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/value/set-rules")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsPronunciationdictionariesRules(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesRules(t, 422, `{
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

	var request pronunciationdictionaries.BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Set(context.Background(), string("value"), &request)
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

func TestAddWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWirePronunciationdictionariesRules(t, 200, `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request pronunciationdictionaries.PronunciationDictionary
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Add(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesRules(gotResponse, []byte(`{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)
	}

	req := receiveWirePronunciationdictionariesRules(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/pronunciation-dictionaries/value/add-rules" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/value/add-rules")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsPronunciationdictionariesRules(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAddWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesRules(t, 422, `{
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

	var request pronunciationdictionaries.PronunciationDictionary
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Add(context.Background(), string("value"), &request)
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

func TestRemoveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWirePronunciationdictionariesRules(t, 200, `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request pronunciationdictionaries.RemovePronunciationDictionaryRulesRequest
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rule_strings": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Remove(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesRules(gotResponse, []byte(`{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`)
	}

	req := receiveWirePronunciationdictionariesRules(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/pronunciation-dictionaries/value/remove-rules" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/value/remove-rules")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsPronunciationdictionariesRules(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestRemoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesRules(t, 422, `{
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

	var request pronunciationdictionaries.RemovePronunciationDictionaryRulesRequest
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "rule_strings": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Remove(context.Background(), string("value"), &request)
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

type wireRequestPronunciationdictionariesRules struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWirePronunciationdictionariesRules records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWirePronunciationdictionariesRules(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestPronunciationdictionariesRules) {
	t.Helper()
	recorded := make(chan wireRequestPronunciationdictionariesRules, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestPronunciationdictionariesRules{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWirePronunciationdictionariesRules(t *testing.T, recorded <-chan wireRequestPronunciationdictionariesRules) wireRequestPronunciationdictionariesRules {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestPronunciationdictionariesRules{}
	}
}

// wireJSONContainsPronunciationdictionariesRules reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsPronunciationdictionariesRules(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualPronunciationdictionariesRules(got, want)
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

// setWireValuePronunciationdictionariesRules sets a named string, int or float field, or the value a pointer field points to.
func setWireValuePronunciationdictionariesRules(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualPronunciationdictionariesRules(got, want []byte) bool {
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
