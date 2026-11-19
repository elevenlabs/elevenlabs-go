package items

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
	orders "github.com/elevenlabs/elevenlabs-go/productions/orders"
)

func TestUpsertWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireProductionsOrdersItems(t, 200, `{
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "quote": {
    "amount_usd": 9.99
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request orders.BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost
	if err := json.Unmarshal([]byte(`{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "request": {
    "item": {
      "kind": "dub",
      "media_id": "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa",
      "source_language": "value",
      "destination_languages": [
        "value"
      ],
      "include_captions": true,
      "include_source_captions": true,
      "instructions": "value",
      "captions_sdh": true
    },
    "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Upsert(context.Background(), OrderId("prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualProductionsOrdersItems(gotResponse, []byte(`{
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "quote": {
    "amount_usd": 9.99
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "quote": {
    "amount_usd": 9.99
  }
}`)
	}

	req := receiveWireProductionsOrdersItems(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/productions/orders/prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa/items" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/productions/orders/prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa/items")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsProductionsOrdersItems(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpsertWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireProductionsOrdersItems(t, 422, `{
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

	var request orders.BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost
	if err := json.Unmarshal([]byte(`{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "request": {
    "item": {
      "kind": "dub",
      "media_id": "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa",
      "source_language": "value",
      "destination_languages": [
        "value"
      ],
      "include_captions": true,
      "include_source_captions": true,
      "instructions": "value",
      "captions_sdh": true
    },
    "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Upsert(context.Background(), OrderId("prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa"), &request)
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

	server, recorded := serveWireProductionsOrdersItems(t, 200, `{
  "success": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Remove(context.Background(), OrderId("prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa"), ItemId("proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualProductionsOrdersItems(gotResponse, []byte(`{
  "success": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "success": true
}`)
	}

	req := receiveWireProductionsOrdersItems(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/productions/orders/prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa/items/proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/productions/orders/prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa/items/proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

func TestRemoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireProductionsOrdersItems(t, 422, `{
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

	_, err := client.Remove(context.Background(), OrderId("prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa"), ItemId("proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa"))
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

type wireRequestProductionsOrdersItems struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireProductionsOrdersItems records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireProductionsOrdersItems(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestProductionsOrdersItems) {
	t.Helper()
	recorded := make(chan wireRequestProductionsOrdersItems, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestProductionsOrdersItems{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireProductionsOrdersItems(t *testing.T, recorded <-chan wireRequestProductionsOrdersItems) wireRequestProductionsOrdersItems {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestProductionsOrdersItems{}
	}
}

// wireJSONContainsProductionsOrdersItems reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsProductionsOrdersItems(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualProductionsOrdersItems(got, want)
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

// setWireValueProductionsOrdersItems sets a named string, int or float field, or the value a pointer field points to.
func setWireValueProductionsOrdersItems(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualProductionsOrdersItems(got, want []byte) bool {
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
