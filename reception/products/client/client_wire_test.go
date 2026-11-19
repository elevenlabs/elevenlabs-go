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
	reception "github.com/elevenlabs/elevenlabs-go/reception"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsClient(t, 200, `[
  {
    "pricing_mode": "fixed",
    "intake_questions": [
      {
        "id": "value",
        "label": "value",
        "kind": "text",
        "options": [
          "value"
        ],
        "required": true,
        "is_active": true
      }
    ],
    "groups": [
      {
        "label": "value",
        "rank": 123
      }
    ],
    "ungrouped_rank": 123,
    "name": "value",
    "description": "value",
    "price": 9.99,
    "is_active": true,
    "orderable": true,
    "allowed_fulfillment_types": [
      "pickup"
    ],
    "location_ids": [
      "value"
    ],
    "stock": {
      "key": 123
    },
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	groupValue := string("value")
	isActiveValue := bool(true)
	request := &reception.ListProductsRequest{
		Group:         &groupValue,
		IsActive:      &isActiveValue,
		XiWorkspaceId: "value",
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsClient(gotResponse, []byte(`[
  {
    "pricing_mode": "fixed",
    "intake_questions": [
      {
        "id": "value",
        "label": "value",
        "kind": "text",
        "options": [
          "value"
        ],
        "required": true,
        "is_active": true
      }
    ],
    "groups": [
      {
        "label": "value",
        "rank": 123
      }
    ],
    "ungrouped_rank": 123,
    "name": "value",
    "description": "value",
    "price": 9.99,
    "is_active": true,
    "orderable": true,
    "allowed_fulfillment_types": [
      "pickup"
    ],
    "location_ids": [
      "value"
    ],
    "stock": {
      "key": 123
    },
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "pricing_mode": "fixed",
    "intake_questions": [
      {
        "id": "value",
        "label": "value",
        "kind": "text",
        "options": [
          "value"
        ],
        "required": true,
        "is_active": true
      }
    ],
    "groups": [
      {
        "label": "value",
        "rank": 123
      }
    ],
    "ungrouped_rank": 123,
    "name": "value",
    "description": "value",
    "price": 9.99,
    "is_active": true,
    "orderable": true,
    "allowed_fulfillment_types": [
      "pickup"
    ],
    "location_ids": [
      "value"
    ],
    "stock": {
      "key": 123
    },
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)
	}

	req := receiveWireReceptionProductsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/products" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products")
	}
	if got := req.Query.Get("group"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "group", got, "value")
	}
	if got := req.Query.Get("is_active"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "is_active", got, "true")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsClient(t, 422, `{
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

	groupValue := string("value")
	isActiveValue := bool(true)
	request := &reception.ListProductsRequest{
		Group:         &groupValue,
		IsActive:      &isActiveValue,
		XiWorkspaceId: "value",
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsClient(t, 200, `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionCreateProductRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "track_quantity": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsClient(gotResponse, []byte(`{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionProductsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/products" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionProductsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsClient(t, 422, `{
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

	var request reception.ReceptionCreateProductRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "track_quantity": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsClient(t, 200, `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.GetProductsRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsClient(gotResponse, []byte(`{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionProductsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/products/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsClient(t, 422, `{
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

	request := &reception.GetProductsRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Get(context.Background(), string("value"), request)
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

	server, recorded := serveWireReceptionProductsClient(t, 200, `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionUpdateProductRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "track_quantity": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsClient(gotResponse, []byte(`{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionProductsClient(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/reception/products/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionProductsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsClient(t, 422, `{
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

	var request reception.ReceptionUpdateProductRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "track_quantity": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsClient(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.DeleteProductsRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsClient(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionProductsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/products/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsClient(t, 422, `{
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

	request := &reception.DeleteProductsRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Delete(context.Background(), string("value"), request)
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

type wireRequestReceptionProductsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionProductsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionProductsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionProductsClient) {
	t.Helper()
	recorded := make(chan wireRequestReceptionProductsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionProductsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionProductsClient(t *testing.T, recorded <-chan wireRequestReceptionProductsClient) wireRequestReceptionProductsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionProductsClient{}
	}
}

// wireJSONContainsReceptionProductsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionProductsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionProductsClient(got, want)
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

// setWireValueReceptionProductsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionProductsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionProductsClient(got, want []byte) bool {
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
