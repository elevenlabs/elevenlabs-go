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

func TestCreateFromRulesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWirePronunciationdictionariesClient(t, 200, `{
  "id": "value",
  "name": "value",
  "created_by": "value",
  "creation_time_unix": 123,
  "version_id": "value",
  "version_rules_num": 123,
  "description": "value",
  "permission_on_resource": "admin"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost
	if err := json.Unmarshal([]byte(`{
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ],
  "name": "value",
  "description": "value",
  "workspace_access": "admin"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateFromRules(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesClient(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "created_by": "value",
  "creation_time_unix": 123,
  "version_id": "value",
  "version_rules_num": 123,
  "description": "value",
  "permission_on_resource": "admin"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "created_by": "value",
  "creation_time_unix": 123,
  "version_id": "value",
  "version_rules_num": 123,
  "description": "value",
  "permission_on_resource": "admin"
}`)
	}

	req := receiveWirePronunciationdictionariesClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/pronunciation-dictionaries/add-from-rules" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/add-from-rules")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsPronunciationdictionariesClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateFromRulesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesClient(t, 422, `{
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

	var request types.BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost
	if err := json.Unmarshal([]byte(`{
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ],
  "name": "value",
  "description": "value",
  "workspace_access": "admin"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateFromRules(context.Background(), &request)
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

	server, recorded := serveWirePronunciationdictionariesClient(t, 200, `{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesClient(gotResponse, []byte(`{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`)
	}

	req := receiveWirePronunciationdictionariesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/pronunciation-dictionaries/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesClient(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"))
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

	server, recorded := serveWirePronunciationdictionariesClient(t, 200, `{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "archived": true,
  "name": "value"
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
	if !wireJSONBytesEqualPronunciationdictionariesClient(gotResponse, []byte(`{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value"
}`)
	}

	req := receiveWirePronunciationdictionariesClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/pronunciation-dictionaries/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsPronunciationdictionariesClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesClient(t, 422, `{
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

	var request types.BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch
	if err := json.Unmarshal([]byte(`{
  "pronunciation_dictionary_id": "value",
  "archived": true,
  "name": "value"
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWirePronunciationdictionariesClient(t, 200, `{
  "pronunciation_dictionaries": [
    {
      "id": "value",
      "latest_version_id": "value",
      "latest_version_rules_num": 123,
      "name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123,
      "description": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	sortDirectionValue := string("value")
	includeArchivedValue := bool(true)
	request := &types.ListPronunciationDictionariesRequest{
		Cursor:          &cursorValue,
		PageSize:        &pageSizeValue,
		SortDirection:   &sortDirectionValue,
		IncludeArchived: &includeArchivedValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualPronunciationdictionariesClient(gotResponse, []byte(`{
  "pronunciation_dictionaries": [
    {
      "id": "value",
      "latest_version_id": "value",
      "latest_version_rules_num": 123,
      "name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123,
      "description": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pronunciation_dictionaries": [
    {
      "id": "value",
      "latest_version_id": "value",
      "latest_version_rules_num": 123,
      "name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123,
      "description": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWirePronunciationdictionariesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/pronunciation-dictionaries" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/pronunciation-dictionaries")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("sort_direction"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "sort_direction", got, "value")
	}
	if got := req.Query.Get("include_archived"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_archived", got, "true")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWirePronunciationdictionariesClient(t, 422, `{
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

	cursorValue := string("value")
	pageSizeValue := int(100)
	sortDirectionValue := string("value")
	includeArchivedValue := bool(true)
	request := &types.ListPronunciationDictionariesRequest{
		Cursor:          &cursorValue,
		PageSize:        &pageSizeValue,
		SortDirection:   &sortDirectionValue,
		IncludeArchived: &includeArchivedValue,
	}
	_, err := client.List(context.Background(), request)
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

type wireRequestPronunciationdictionariesClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWirePronunciationdictionariesClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWirePronunciationdictionariesClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestPronunciationdictionariesClient) {
	t.Helper()
	recorded := make(chan wireRequestPronunciationdictionariesClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestPronunciationdictionariesClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWirePronunciationdictionariesClient(t *testing.T, recorded <-chan wireRequestPronunciationdictionariesClient) wireRequestPronunciationdictionariesClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestPronunciationdictionariesClient{}
	}
}

// wireJSONContainsPronunciationdictionariesClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsPronunciationdictionariesClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualPronunciationdictionariesClient(got, want)
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

// setWireValuePronunciationdictionariesClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValuePronunciationdictionariesClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualPronunciationdictionariesClient(got, want []byte) bool {
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
