package types

import (
	"testing"
)

func TestConversationHistoryMetadataCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time_unix_secs": 123,
  "accepted_time_unix_secs": 123,
  "call_duration_secs": 123,
  "queue_wait_secs": 9.99,
  "cost": 123,
  "deletion_settings": {
    "deletion_time_unix_secs": 123,
    "deleted_logs_at_time_unix_secs": 123,
    "deleted_audio_at_time_unix_secs": 123,
    "deleted_transcript_at_time_unix_secs": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true
  },
  "feedback": {
    "type": "thumbs",
    "overall_score": "like",
    "likes": 123,
    "dislikes": 123,
    "rating": 123,
    "comment": "value"
  },
  "authorization_method": "invalid",
  "charging": {
    "dev_discount": true,
    "is_burst": true,
    "tier": "value",
    "llm_usage": {
      "irreversible_generation": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total_reasoning_portion": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "initiated_generation": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total_reasoning_portion": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      }
    },
    "llm_price": 9.99,
    "llm_charge": 123,
    "call_charge": 123,
    "platform_charge": 123,
    "platform_usage": {
      "category_usage": {
        "key": {
          "credits": 123,
          "price": 9.99,
          "quantity": 9.99
        }
      }
    },
    "platform_price": 9.99,
    "free_minutes_consumed": 9.99,
    "free_llm_dollars_consumed": 9.99,
    "tts_usage": {
      "primary_tts_model": "value",
      "total_audio_output_seconds": 9.99,
      "total_characters": 123,
      "per_voice_usage": [
        {
          "voice_id": "value",
          "audio_output_seconds": 9.99
        }
      ]
    },
    "asr_usage": {
      "asr_model": "value",
      "total_transcription_calls": 123,
      "total_audio_input_seconds": 9.99
    },
    "analysis": {
      "total": {
        "price": 9.99,
        "charge": 123,
        "runs": 123,
        "price_per_feature": {
          "key": 9.99
        },
        "charge_per_feature": {
          "key": 123
        }
      },
      "last_run": {
        "price": 9.99,
        "charge": 123,
        "price_per_feature": {
          "key": 9.99
        },
        "charge_per_feature": {
          "key": 123
        }
      }
    },
    "voice_surcharge_charge": 123
  },
  "phone_call": {
    "direction": "inbound",
    "phone_number_id": "value",
    "agent_number": "value",
    "external_number": "value",
    "call_result": "answered",
    "type": "twilio",
    "stream_sid": "value",
    "call_sid": "value"
  },
  "batch_call": {
    "batch_call_id": "value",
    "batch_call_recipient_id": "value",
    "campaign": {
      "campaign_id": "value",
      "campaign_lead_id": "value",
      "cold_attempt": 123,
      "conversation_type": "initial"
    }
  },
  "termination_reason": "value",
  "error": {
    "code": 123,
    "reason": "value"
  },
  "warnings": [
    "value"
  ],
  "main_language": "value",
  "rag_usage": {
    "usage_count": 123,
    "embedding_model": "value"
  },
  "text_only": true,
  "features_usage": {
    "language_detection": {
      "enabled": true,
      "used": true
    },
    "transfer_to_agent": {
      "enabled": true,
      "used": true
    },
    "transfer_to_number": {
      "enabled": true,
      "used": true
    },
    "multivoice": {
      "enabled": true,
      "used": true
    },
    "dtmf_tones": {
      "enabled": true,
      "used": true
    },
    "external_mcp_servers": {
      "enabled": true,
      "used": true
    },
    "pii_zrm_workspace": true,
    "pii_zrm_agent": true,
    "tool_dynamic_variable_updates": {
      "enabled": true,
      "used": true
    },
    "is_livekit": true,
    "voicemail_detection": {
      "enabled": true,
      "used": true
    },
    "dtmf_input": {
      "enabled": true,
      "used": true
    },
    "workflow": {
      "enabled": true,
      "tool_node": {
        "enabled": true,
        "used": true
      },
      "standalone_agent_node": {
        "enabled": true,
        "used": true
      },
      "phone_number_node": {
        "enabled": true,
        "used": true
      },
      "end_node": {
        "enabled": true,
        "used": true
      },
      "override_agent_node": {
        "enabled": true,
        "used": true
      }
    },
    "agent_testing": {
      "enabled": true,
      "tests_ran_after_last_modification": true,
      "tests_ran_in_last_7_days": true
    },
    "versioning": {
      "enabled": true,
      "used": true
    },
    "file_input": {
      "enabled": true,
      "used": true
    },
    "freeform_procedure": {
      "enabled": true,
      "used": true
    },
    "structured_procedure": {
      "enabled": true,
      "used": true
    }
  },
  "eleven_assistant": {
    "is_eleven_assistant": true
  },
  "initiator_id": "value",
  "conversation_initiation_source": "unknown",
  "conversation_initiation_source_version": "value",
  "timezone": "value",
  "async_metadata": {
    "delivery_status": "pending",
    "delivery_timestamp": 123,
    "delivery_error": "value",
    "external_system": "value",
    "external_id": "value",
    "external_link": "value",
    "retry_count": 123,
    "last_retry_timestamp": 123,
    "last_processed_external_message_id": "value",
    "integration_trigger_id": "value",
    "platform_metadata": {
      "conversation_initiation_source": "plain_integration",
      "thread_id": "value",
      "workspace_id": "value"
    }
  },
  "whatsapp": {
    "direction": "inbound",
    "whatsapp_phone_number_id": "value",
    "whatsapp_user_id": "value",
    "awaiting_first_user_message": true,
    "recipient": {
      "phone": "value",
      "bsuid": "value"
    },
    "call_permission_request": {
      "state": "blocked",
      "requested_at_unix": 123,
      "responded_at_unix": 123,
      "permanent": true,
      "granted_by_user_call": true
    },
    "call_result": "answered",
    "opening_message": {
      "template_name": "value",
      "wamid": "value",
      "sent_at_unix": 123,
      "delivered_at_unix": 123,
      "read_at_unix": 123,
      "replied_at_unix": 123,
      "failed_at_unix": 123,
      "error_code": "value",
      "pricing_category": "value",
      "billable": true
    }
  },
  "sms": {
    "direction": "inbound",
    "phone_number_id": "value",
    "sms_user_phone_number": "value",
    "agent_phone_number": "value"
  },
  "agent_created_from": "cli",
  "agent_last_updated_from": "cli",
  "voice_rewards": [
    {
      "voice_id": "value",
      "reward_usd_cents": 9.99
    }
  ],
  "cost_fiat": 9.99
}`
	assertRoundTrip(t, wire, new(ConversationHistoryMetadataCommonModel))
}
