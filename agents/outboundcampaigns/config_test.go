package outboundcampaigns

import (
	"testing"
)

func TestReplaceWorkspaceConfigRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "expected_version": 123,
  "config": {
    "policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    },
    "agent_policies": {
      "key": {
        "timezone": "value",
        "retry": {
          "retry_delay_seconds": 123,
          "max_short_retries": 30,
          "max_total_calls": 30,
          "max_callbacks": 30,
          "max_redirects": 30,
          "interval_placement": "start",
          "rotate_step_seconds": 123,
          "max_consecutive_failures_per_number": 30
        },
        "data_collection": {
          "callback_time": "value",
          "redirect_phone_number": "value",
          "stop": "value",
          "mappings": {
            "key": "value"
          }
        },
        "followup_context": {
          "call_type": true,
          "previous_conversation_id": true,
          "previous_summary": true,
          "previous_transcript": true
        }
      }
    },
    "whatsapp_voice_policy": {
      "timezone": "value",
      "retry": {
        "retry_delay_seconds": 123,
        "max_short_retries": 30,
        "max_total_calls": 30,
        "max_callbacks": 30,
        "max_redirects": 30,
        "interval_placement": "start",
        "rotate_step_seconds": 123,
        "max_consecutive_failures_per_number": 30
      },
      "data_collection": {
        "callback_time": "value",
        "redirect_phone_number": "value",
        "stop": "value",
        "mappings": {
          "key": "value"
        }
      },
      "followup_context": {
        "call_type": true,
        "previous_conversation_id": true,
        "previous_summary": true,
        "previous_transcript": true
      }
    }
  }
}`
	assertRoundTrip(t, wire, new(ReplaceWorkspaceConfigRequest))
}
