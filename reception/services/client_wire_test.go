package services

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

	server, recorded := serveWireReceptionServices(t, 200, `[
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	staffIdValue := string("value")
	assetIdValue := string("value")
	isActiveValue := bool(true)
	request := &reception.ListServicesRequest{
		StaffId:       &staffIdValue,
		AssetId:       &assetIdValue,
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
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`[
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  }
]`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/services" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services")
	}
	if got := req.Query.Get("staff_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "staff_id", got, "value")
	}
	if got := req.Query.Get("asset_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "asset_id", got, "value")
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

	server, _ := serveWireReceptionServices(t, 422, `{
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

	staffIdValue := string("value")
	assetIdValue := string("value")
	isActiveValue := bool(true)
	request := &reception.ListServicesRequest{
		StaffId:       &staffIdValue,
		AssetId:       &assetIdValue,
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

	server, recorded := serveWireReceptionServices(t, 200, `{
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionCreateServiceRequest
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ]
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
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`{
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
}`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/services" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionServices(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionServices(t, 422, `{
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

	var request reception.ReceptionCreateServiceRequest
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ]
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

	server, recorded := serveWireReceptionServices(t, 200, `{
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.GetServicesRequest{
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
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`{
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "duration_minutes": 123,
  "price": 9.99
}`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/services/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionServices(t, 422, `{
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

	request := &reception.GetServicesRequest{
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

	server, recorded := serveWireReceptionServices(t, 200, `{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionUpdateServiceRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "service_id": "value",
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "duration_minutes": 123,
  "price": 9.99
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
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/reception/services/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionServices(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionServices(t, 422, `{
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

	var request reception.ReceptionUpdateServiceRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "service_id": "value",
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "duration_minutes": 123,
  "price": 9.99
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

	server, recorded := serveWireReceptionServices(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.DeleteServicesRequest{
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
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/services/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionServices(t, 422, `{
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

	request := &reception.DeleteServicesRequest{
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

func TestConfirmUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionServices(t, 200, `{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionConfirmServiceUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "reprice": true,
  "xi-workspace-id": "value",
  "service_id": "value",
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "duration_minutes": 123,
  "price": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	repriceValue := bool(true)
	request.Reprice = &repriceValue
	request.XiWorkspaceId = "value"
	result, err := client.ConfirmUpdate(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionServices(gotResponse, []byte(`{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "service": {
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
    "variants": [
      {
        "id": "value",
        "duration_minutes": 123,
        "price": 9.99,
        "is_active": true
      }
    ],
    "kind": "classic",
    "buffer_minutes": 123,
    "service_location": "business",
    "add_ons": [
      {
        "id": "value",
        "name": "value",
        "description": "value",
        "add_duration_minutes": 123,
        "add_price": 9.99
      }
    ],
    "max_participants": 123,
    "capacity": 123,
    "is_active": true,
    "any_staff_member": true,
    "staff_ids": [
      "value"
    ],
    "asset_ids": [
      "value"
    ],
    "advance_booking_minutes": 123,
    "location_ids": [
      "value"
    ],
    "display_on_booking_page": true,
    "travel_minutes": 123,
    "travel_before": true,
    "travel_after": true,
    "promotions": [
      {
        "id": "value",
        "name": "value",
        "target": "service",
        "variant_id": "value",
        "add_on_id": "value",
        "promotional_price": 9.99,
        "days_of_week": [
          6
        ],
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "is_active": true
      }
    ],
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "duration_minutes": 123,
    "price": 9.99
  },
  "has_affected_events": true
}`)
	}

	req := receiveWireReceptionServices(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/services/value/confirm-update" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/services/value/confirm-update")
	}
	if got := req.Query.Get("reprice"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "reprice", got, "true")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionServices(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestConfirmUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionServices(t, 422, `{
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

	var request reception.ReceptionConfirmServiceUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "reprice": true,
  "xi-workspace-id": "value",
  "service_id": "value",
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
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123,
  "capacity": 123,
  "is_active": true,
  "any_staff_member": true,
  "staff_ids": [
    "value"
  ],
  "asset_ids": [
    "value"
  ],
  "advance_booking_minutes": 123,
  "location_ids": [
    "value"
  ],
  "display_on_booking_page": true,
  "travel_minutes": 123,
  "travel_before": true,
  "travel_after": true,
  "promotions": [
    {
      "id": "value",
      "name": "value",
      "target": "service",
      "variant_id": "value",
      "add_on_id": "value",
      "promotional_price": 9.99,
      "days_of_week": [
        6
      ],
      "start_date": "2024-01-01",
      "end_date": "2024-01-01",
      "is_active": true
    }
  ],
  "duration_minutes": 123,
  "price": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	repriceValue := bool(true)
	request.Reprice = &repriceValue
	request.XiWorkspaceId = "value"
	_, err := client.ConfirmUpdate(context.Background(), string("value"), &request)
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

type wireRequestReceptionServices struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionServices records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionServices(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionServices) {
	t.Helper()
	recorded := make(chan wireRequestReceptionServices, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionServices{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionServices(t *testing.T, recorded <-chan wireRequestReceptionServices) wireRequestReceptionServices {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionServices{}
	}
}

// wireJSONContainsReceptionServices reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionServices(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionServices(got, want)
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

// setWireValueReceptionServices sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionServices(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionServices(got, want []byte) bool {
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
