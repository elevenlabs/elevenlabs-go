package reception

import (
	"testing"
)

func TestReceptionCreateAssetRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "is_active": true,
  "available_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "photo_url": "value",
  "location_id": "value",
  "customer_facing_name": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionCreateAssetRequest))
}

func TestReceptionUpdateAssetRequestPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "is_active": true,
  "available_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "photo_url": "value",
  "location_id": "value",
  "customer_facing_name": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateAssetRequestPatch))
}
