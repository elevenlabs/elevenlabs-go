package finetunes

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireMusicFinetunes(t, 200, `{
  "finetunes": [
    {
      "id": "value",
      "name": "value",
      "tags": [
        "value"
      ],
      "primary_genre": "value",
      "model_id": "value",
      "created_at": "2024-01-01T00:00:00Z",
      "visibility": "private",
      "created_by": "self",
      "status": "pending",
      "training_progress": 9.99,
      "failure_reason": "audio_processing_failed"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(123)
	request := &music.ListFinetunesRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualMusicFinetunes(gotResponse, []byte(`{
  "finetunes": [
    {
      "id": "value",
      "name": "value",
      "tags": [
        "value"
      ],
      "primary_genre": "value",
      "model_id": "value",
      "created_at": "2024-01-01T00:00:00Z",
      "visibility": "private",
      "created_by": "self",
      "status": "pending",
      "training_progress": 9.99,
      "failure_reason": "audio_processing_failed"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "finetunes": [
    {
      "id": "value",
      "name": "value",
      "tags": [
        "value"
      ],
      "primary_genre": "value",
      "model_id": "value",
      "created_at": "2024-01-01T00:00:00Z",
      "visibility": "private",
      "created_by": "self",
      "status": "pending",
      "training_progress": 9.99,
      "failure_reason": "audio_processing_failed"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireMusicFinetunes(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/music/finetunes" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/music/finetunes")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireMusicFinetunes(t, 422, `{
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
	pageSizeValue := int(123)
	request := &music.ListFinetunesRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireMusicFinetunes(t, 200, `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
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
	if !wireJSONBytesEqualMusicFinetunes(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)
	}

	req := receiveWireMusicFinetunes(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/music/finetunes/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/music/finetunes/value")
	}
}

func TestGetWireError403(t *testing.T) {
	t.Parallel()

	server, _ := serveWireMusicFinetunes(t, 403, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Get(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 403 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 403 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 403)
	}
}

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireMusicFinetunes(t, 200, `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request music.UpdateMusicFinetuneRequestModel
	if err := json.Unmarshal([]byte(`{
  "finetune_id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "visibility": "private"
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
	if !wireJSONBytesEqualMusicFinetunes(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)
	}

	req := receiveWireMusicFinetunes(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/music/finetunes/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/music/finetunes/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsMusicFinetunes(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError403(t *testing.T) {
	t.Parallel()

	server, _ := serveWireMusicFinetunes(t, 403, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request music.UpdateMusicFinetuneRequestModel
	if err := json.Unmarshal([]byte(`{
  "finetune_id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "visibility": "private"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 403 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 403 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 403)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireMusicFinetunes(t, 200, `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualMusicFinetunes(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "model_id": "value",
  "created_at": "2024-01-01T00:00:00Z",
  "visibility": "private",
  "created_by": "self",
  "status": "pending",
  "training_progress": 9.99,
  "failure_reason": "audio_processing_failed"
}`)
	}

	req := receiveWireMusicFinetunes(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/music/finetunes/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/music/finetunes/value")
	}
}

func TestDeleteWireError403(t *testing.T) {
	t.Parallel()

	server, _ := serveWireMusicFinetunes(t, 403, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Delete(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 403 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 403 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 403)
	}
}

type wireRequestMusicFinetunes struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireMusicFinetunes records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireMusicFinetunes(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestMusicFinetunes) {
	t.Helper()
	recorded := make(chan wireRequestMusicFinetunes, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestMusicFinetunes{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireMusicFinetunes(t *testing.T, recorded <-chan wireRequestMusicFinetunes) wireRequestMusicFinetunes {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestMusicFinetunes{}
	}
}

// wireJSONContainsMusicFinetunes reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsMusicFinetunes(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualMusicFinetunes(got, want)
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

// setWireValueMusicFinetunes sets a named string, int or float field, or the value a pointer field points to.
func setWireValueMusicFinetunes(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualMusicFinetunes(got, want []byte) bool {
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
