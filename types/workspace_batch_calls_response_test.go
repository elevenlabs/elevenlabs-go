package types

import (
	"testing"
)

func TestWorkspaceBatchCallsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "batch_calls": [
    {
      "id": "value",
      "phone_number_id": "value",
      "phone_provider": "twilio",
      "whatsapp_params": {
        "whatsapp_phone_number_id": "value",
        "whatsapp_call_permission_request_template_name": "value",
        "whatsapp_call_permission_request_template_language_code": "value"
      },
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "environment": "value",
      "created_at_unix": 123,
      "scheduled_time_unix": 123,
      "timezone": "value",
      "total_calls_dispatched": 123,
      "total_calls_scheduled": 123,
      "total_calls_finished": 123,
      "last_updated_at_unix": 123,
      "status": "pending",
      "retry_count": 123,
      "telephony_call_config": {
        "ringing_timeout_secs": 123,
        "twilio_call_recording_enabled": true,
        "twilio_machine_detection": {
          "mode": "enable",
          "actions": [
            {
              "action": "hang_up",
              "answered_by": [
                "human"
              ]
            }
          ]
        }
      },
      "target_concurrency_limit": 123,
      "agent_name": "value",
      "branch_name": "value"
    }
  ],
  "next_doc": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(WorkspaceBatchCallsResponse))
}
