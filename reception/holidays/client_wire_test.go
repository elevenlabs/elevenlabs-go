package holidays

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

	server, recorded := serveWireReceptionHolidays(t, 200, `[
  {
    "name": "value",
    "start_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "end_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "holiday_type": "closed",
    "override_hours": [
      {
        "start_hour": 123,
        "start_minute": 123,
        "end_hour": 24,
        "end_minute": 123
      }
    ],
    "time_off_type": "staff",
    "resource_id": "value",
    "location_id": "value",
    "reason": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	startUnixValue := int(123)
	endUnixValue := int(123)
	resourceIdValue := string("value")
	locationIdValue := string("value")
	request := &reception.ListHolidaysRequest{
		StartUnix:     &startUnixValue,
		EndUnix:       &endUnixValue,
		ResourceId:    &resourceIdValue,
		LocationId:    &locationIdValue,
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
	if !wireJSONBytesEqualReceptionHolidays(gotResponse, []byte(`[
  {
    "name": "value",
    "start_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "end_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "holiday_type": "closed",
    "override_hours": [
      {
        "start_hour": 123,
        "start_minute": 123,
        "end_hour": 24,
        "end_minute": 123
      }
    ],
    "time_off_type": "staff",
    "resource_id": "value",
    "location_id": "value",
    "reason": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "name": "value",
    "start_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "end_date": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "holiday_type": "closed",
    "override_hours": [
      {
        "start_hour": 123,
        "start_minute": 123,
        "end_hour": 24,
        "end_minute": 123
      }
    ],
    "time_off_type": "staff",
    "resource_id": "value",
    "location_id": "value",
    "reason": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)
	}

	req := receiveWireReceptionHolidays(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/holidays" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/holidays")
	}
	if got := req.Query.Get("start_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "start_unix", got, "123")
	}
	if got := req.Query.Get("end_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "end_unix", got, "123")
	}
	if got := req.Query.Get("resource_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "resource_id", got, "value")
	}
	if got := req.Query.Get("location_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "location_id", got, "value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionHolidays(t, 422, `{
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

	startUnixValue := int(123)
	endUnixValue := int(123)
	resourceIdValue := string("value")
	locationIdValue := string("value")
	request := &reception.ListHolidaysRequest{
		StartUnix:     &startUnixValue,
		EndUnix:       &endUnixValue,
		ResourceId:    &resourceIdValue,
		LocationId:    &locationIdValue,
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

	server, recorded := serveWireReceptionHolidays(t, 200, `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionCreateHolidayRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
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
	if !wireJSONBytesEqualReceptionHolidays(gotResponse, []byte(`{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionHolidays(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/holidays" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/holidays")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionHolidays(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionHolidays(t, 422, `{
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

	var request reception.ReceptionCreateHolidayRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
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

	server, recorded := serveWireReceptionHolidays(t, 200, `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.GetHolidaysRequest{
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
	if !wireJSONBytesEqualReceptionHolidays(gotResponse, []byte(`{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionHolidays(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/holidays/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/holidays/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionHolidays(t, 422, `{
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

	request := &reception.GetHolidaysRequest{
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

	server, recorded := serveWireReceptionHolidays(t, 200, `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionUpdateHolidayRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "holiday_id": "value",
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
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
	if !wireJSONBytesEqualReceptionHolidays(gotResponse, []byte(`{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionHolidays(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/reception/holidays/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/holidays/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionHolidays(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionHolidays(t, 422, `{
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

	var request reception.ReceptionUpdateHolidayRequestPatch
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "holiday_id": "value",
  "name": "value",
  "start_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "end_date": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "holiday_type": "closed",
  "override_hours": [
    {
      "start_hour": 123,
      "start_minute": 123,
      "end_hour": 24,
      "end_minute": 123
    }
  ],
  "time_off_type": "staff",
  "resource_id": "value",
  "location_id": "value",
  "reason": "value"
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

	server, recorded := serveWireReceptionHolidays(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.DeleteHolidaysRequest{
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
	if !wireJSONBytesEqualReceptionHolidays(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionHolidays(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/holidays/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/holidays/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionHolidays(t, 422, `{
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

	request := &reception.DeleteHolidaysRequest{
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

type wireRequestReceptionHolidays struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionHolidays records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionHolidays(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionHolidays) {
	t.Helper()
	recorded := make(chan wireRequestReceptionHolidays, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionHolidays{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionHolidays(t *testing.T, recorded <-chan wireRequestReceptionHolidays) wireRequestReceptionHolidays {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionHolidays{}
	}
}

// wireJSONContainsReceptionHolidays reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionHolidays(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionHolidays(got, want)
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

// setWireValueReceptionHolidays sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionHolidays(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionHolidays(got, want []byte) bool {
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
