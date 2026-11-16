package reception

import (
	"testing"
)

func TestReceptionCreateServiceRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ReceptionCreateServiceRequest))
}

func TestReceptionUpdateServiceRequestPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateServiceRequestPatch))
}

func TestReceptionConfirmServiceUpdateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ReceptionConfirmServiceUpdateRequest))
}
