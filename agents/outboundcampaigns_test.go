package agents

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSubmitCampaignRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "call_name": "value",
  "agent_id": "value",
  "recipients": [
    {
      "phone_number": "value",
      "phone_numbers": [
        "value"
      ],
      "whatsapp_user_id": "value",
      "id": "value",
      "conversation_initiation_client_data": {
        "conversation_config_override": {
          "asr": {
            "keywords": [
              "value"
            ]
          },
          "turn": {
            "soft_timeout_config": {
              "message": "value",
              "additional_soft_timeout_messages": [
                "value"
              ]
            },
            "noisy_environment_message": "value"
          },
          "tts": {
            "model_id": "eleven_turbo_v2",
            "voice_id": "value",
            "supported_voices": [
              {
                "label": "value",
                "voice_id": "value",
                "description": "value",
                "language": "value",
                "model_family": "turbo",
                "optimize_streaming_latency": 123,
                "stability": 0.5,
                "speed": 1.2,
                "similarity_boost": 0.5,
                "clone_strength": 0.5
              }
            ],
            "stability": 9.99,
            "speed": 9.99,
            "similarity_boost": 9.99,
            "clone_strength": 9.99,
            "pronunciation_dictionary_locators": [
              {
                "pronunciation_dictionary_id": "value",
                "version_id": "value"
              }
            ]
          },
          "conversation": {
            "text_only": true,
            "max_duration_seconds": 123
          },
          "agent": {
            "first_message": "value",
            "language": "value",
            "max_conversation_duration_message": "value",
            "prompt": {
              "prompt": "value",
              "llm": "gpt-4o-mini",
              "reasoning_effort": "none",
              "tool_ids": [
                "value"
              ],
              "tool_ids_strict": true,
              "native_mcp_server_ids": [
                "value"
              ],
              "knowledge_base": [
                {
                  "type": "file",
                  "name": "value",
                  "id": "value",
                  "usage_mode": "prompt"
                }
              ]
            }
          }
        },
        "custom_llm_extra_body": {
          "key": {}
        },
        "user_id": "value",
        "source_info": {
          "source": "unknown",
          "version": "value"
        },
        "branch_id": "value",
        "environment": "value",
        "starting_workflow_node_id": "value",
        "procedure_ids": [
          "value"
        ],
        "start_procedure_ids": [
          "value"
        ],
        "dynamic_variables": {
          "key": {}
        }
      }
    }
  ],
  "scheduled_time_unix": 123,
  "timezone": "value",
  "branch_id": "value",
  "environment": "value",
  "target_concurrency_limit": 123,
  "idempotency_key": "value",
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
    "date_overrides": [
      {
        "start_date": "2024-01-01",
        "end_date": "2024-01-01",
        "name": "value"
      }
    ],
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
  "variant": "simple",
  "channel_config": {
    "channel": "phone",
    "agent_phone_number_id": "value",
    "mode_config": {
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
      },
      "mode": "voice"
    }
  }
}`
	assertRoundTrip(t, wire, new(SubmitCampaignRequest))
}

func TestSubmitCampaignRequestChannelConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "channel": "phone",
  "agent_phone_number_id": "value",
  "mode_config": {
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
    },
    "mode": "voice"
  }
}`
	assertRoundTrip(t, wire, new(SubmitCampaignRequestChannelConfig))
}

func TestListOutboundCampaignsRequestModeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListOutboundCampaignsRequestMode
		wire    string
	}
	cases := []enumCase{
		{variant: ListOutboundCampaignsRequestModeVoice, wire: "\"voice\""},
		{variant: ListOutboundCampaignsRequestModeMessage, wire: "\"message\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListOutboundCampaignsRequestMode
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
