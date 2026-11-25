package types

import (
	"testing"
)

func TestCampaignLeadsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "leads": [
    {
      "lead_id": "value",
      "recipient_id": "value",
      "phone_number": "value",
      "phone_numbers": [
        "value"
      ],
      "phone_numbers_index": 123,
      "timezone_override": "value",
      "user_id": "value",
      "attempt_count": 123,
      "failure_count": 123,
      "callback_count": 123,
      "redirect_count": 123,
      "do_not_call": true,
      "terminal_disposition": "campaign_ended",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "dynamic_variables": {
        "key": {}
      },
      "overrides": {
        "language": "value",
        "first_message": "value",
        "prompt": "value",
        "voice_id": "value"
      },
      "delivery": {
        "status": "pending",
        "stage": "queued",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_message": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(CampaignLeadsResponse))
}
