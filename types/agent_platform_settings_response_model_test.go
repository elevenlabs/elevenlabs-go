package types

import (
	"testing"
)

func TestAgentPlatformSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation": {
    "criteria": [
      {
        "id": "value",
        "name": "value",
        "type": "prompt",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scope": "conversation",
        "llm": "gpt-4o-mini",
        "scoring_mode": "binary",
        "max_score": 123,
        "score_instructions": "value"
      }
    ]
  },
  "widget": {
    "variant": "tiny",
    "placement": "top-left",
    "expandable": "never",
    "avatar": {
      "type": "orb",
      "color_1": "value",
      "color_2": "value"
    },
    "feedback_mode": "none",
    "end_feedback": {
      "type": "rating"
    },
    "bg_color": "value",
    "text_color": "value",
    "btn_color": "value",
    "btn_text_color": "value",
    "border_color": "value",
    "focus_color": "value",
    "border_radius": 123,
    "btn_radius": 123,
    "action_text": "value",
    "start_call_text": "value",
    "end_call_text": "value",
    "expand_text": "value",
    "listening_text": "value",
    "speaking_text": "value",
    "shareable_page_text": "value",
    "shareable_page_show_terms": true,
    "terms_text": "value",
    "terms_html": "value",
    "terms_key": "value",
    "show_avatar_when_collapsed": true,
    "disable_banner": true,
    "override_link": "value",
    "markdown_link_allowed_hosts": [
      {
        "hostname": "value"
      }
    ],
    "markdown_link_include_www": true,
    "markdown_link_allow_http": true,
    "mic_muting_enabled": true,
    "transcript_enabled": true,
    "text_input_enabled": true,
    "conversation_mode_toggle_enabled": true,
    "default_expanded": true,
    "always_expanded": true,
    "dismissible": true,
    "show_agent_status": true,
    "show_conversation_id": true,
    "strip_audio_tags": true,
    "syntax_highlight_theme": "light",
    "text_contents": {
      "main_label": "value",
      "start_call": "value",
      "start_chat": "value",
      "new_call": "value",
      "end_call": "value",
      "mute_microphone": "value",
      "change_language": "value",
      "collapse": "value",
      "expand": "value",
      "copied": "value",
      "accept_terms": "value",
      "dismiss_terms": "value",
      "listening_status": "value",
      "speaking_status": "value",
      "connecting_status": "value",
      "chatting_status": "value",
      "queue_waiting_status": "value",
      "queue_waiting_status_short": "value",
      "input_label": "value",
      "input_placeholder": "value",
      "input_placeholder_text_only": "value",
      "input_placeholder_new_conversation": "value",
      "user_ended_conversation": "value",
      "agent_ended_conversation": "value",
      "conversation_id": "value",
      "error_occurred": "value",
      "queue_timed_out": "value",
      "copy_id": "value",
      "initiate_feedback": "value",
      "request_follow_up_feedback": "value",
      "thanks_for_feedback": "value",
      "thanks_for_feedback_details": "value",
      "follow_up_feedback_placeholder": "value",
      "submit": "value",
      "go_back": "value",
      "send_message": "value",
      "text_mode": "value",
      "voice_mode": "value",
      "switched_to_text_mode": "value",
      "switched_to_voice_mode": "value",
      "copy": "value",
      "download": "value",
      "wrap": "value",
      "agent_working": "value",
      "agent_done": "value",
      "agent_error": "value",
      "attach_file": "value",
      "remove_file": "value",
      "file_upload_error": "value",
      "file_type_unsupported": "value",
      "file_too_large": "value",
      "file_limit_reached": "value",
      "typing_indicator": "value",
      "rich_content_unavailable": "value"
    },
    "styles": {
      "base": "value",
      "base_hover": "value",
      "base_active": "value",
      "base_border": "value",
      "base_subtle": "value",
      "base_primary": "value",
      "base_error": "value",
      "accent": "value",
      "accent_hover": "value",
      "accent_active": "value",
      "accent_border": "value",
      "accent_subtle": "value",
      "accent_primary": "value",
      "overlay_padding": 9.99,
      "button_radius": 9.99,
      "input_radius": 9.99,
      "bubble_radius": 9.99,
      "sheet_radius": 9.99,
      "compact_sheet_radius": 9.99,
      "dropdown_sheet_radius": 9.99
    },
    "show_resize_button": true,
    "end_confirmation_enabled": true,
    "mobile_fullscreen": true,
    "transcript_speech_sync_enabled": true,
    "language_selector": true,
    "supports_text_only": true,
    "custom_avatar_path": "value",
    "language_presets": {
      "key": {
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "text_contents_translation": {
          "source": {
            "key": "value"
          },
          "text": {
            "key": "value"
          }
        },
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value",
        "terms_translation": {
          "source_hash": "value",
          "text": "value"
        }
      }
    }
  },
  "data_collection": {
    "key": {
      "type": "boolean",
      "description": "value",
      "enum": [
        "value"
      ],
      "is_system_provided": true,
      "dynamic_variable": "value",
      "allowed_values": {
        "dynamic_variable": "value"
      },
      "allowed_values_dynamic_variable": "value",
      "constant_value": "value",
      "is_omitted": true,
      "name": "value",
      "llm": "gpt-4o-mini"
    }
  },
  "data_collection_scopes": {
    "key": "conversation"
  },
  "analysis_items": {
    "evaluation_criteria": [
      {
        "source": "user",
        "analysis_item_id": "value",
        "version_id": "value",
        "additional_version_ids": [
          "value"
        ],
        "scope": "conversation",
        "weight": 9.99
      }
    ],
    "data_collection": [
      {
        "source": "user",
        "analysis_item_id": "value",
        "version_id": "value",
        "scope": "conversation"
      }
    ]
  },
  "overrides": {
    "conversation_config_override": {
      "asr": {
        "keywords": true
      },
      "turn": {
        "soft_timeout_config": {
          "message": true,
          "additional_soft_timeout_messages": true
        },
        "noisy_environment_message": true
      },
      "tts": {
        "model_id": true,
        "voice_id": true,
        "supported_voices": true,
        "stability": true,
        "speed": true,
        "similarity_boost": true,
        "clone_strength": true,
        "pronunciation_dictionary_locators": true
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": true
      },
      "agent": {
        "first_message": true,
        "language": true,
        "max_conversation_duration_message": true,
        "prompt": {
          "prompt": true,
          "llm": true,
          "reasoning_effort": true,
          "tool_ids": true,
          "tool_ids_strict": true,
          "native_mcp_server_ids": true,
          "knowledge_base": true
        }
      }
    },
    "custom_llm_extra_body": true,
    "enable_conversation_initiation_client_data_from_webhook": true,
    "enable_starting_workflow_node_id_from_client": true,
    "enable_procedure_ids_from_client": true
  },
  "workspace_overrides": {
    "conversation_initiation_client_data_webhook": {
      "url": "value",
      "request_headers": {
        "key": "value"
      }
    },
    "webhooks": {
      "post_call_webhook_id": "value",
      "events": [
        "transcript"
      ],
      "transcript_format": "json",
      "exclude_transcript": true,
      "send_audio": true
    }
  },
  "testing": {
    "attached_tests": [
      {
        "test_id": "value",
        "workflow_node_id": "value"
      }
    ]
  },
  "archived": true,
  "guardrails": {
    "version": "1",
    "focus": {
      "is_enabled": true
    },
    "prompt_injection": {
      "is_enabled": true
    },
    "content": {
      "execution_mode": "streaming",
      "config": {
        "sexual": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "violence": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "harassment": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "self_harm": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "profanity": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "religion_or_politics": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "medical_and_legal_information": {
          "is_enabled": true,
          "threshold": 9.99
        }
      },
      "trigger_action": {
        "type": "end_call"
      }
    },
    "custom": {
      "config": {
        "configs": [
          {
            "is_enabled": true,
            "name": "value",
            "prompt": "value",
            "execution_mode": "streaming",
            "model": "gemini-2.5-flash-lite",
            "service_tier": "standard",
            "history_message_count": 100,
            "trigger_action": {
              "type": "end_call"
            },
            "evaluate_full_response_only": true
          }
        ]
      }
    }
  },
  "summary_language": "value",
  "auto_translate_transcript_to_app_language": true,
  "auth": {
    "enable_auth": true,
    "allowlist": [
      {
        "hostname": "value"
      }
    ],
    "require_origin_header": true,
    "shareable_token": "value"
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "queueing_config": {
    "enabled": true,
    "wait_timeout_seconds": 123,
    "hold_audio": {
      "audio_path": "value",
      "audio_url": "value",
      "original_filename": "value",
      "duration_secs": 9.99,
      "size_bytes": 123
    }
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "trust_context": "unknown",
  "analysis_llm": "gpt-4o-mini",
  "topic_discovery": {},
  "sentiment_analysis": {},
  "alerting": {
    "monitor_configs": {
      "key": {
        "enabled": true,
        "threshold": 0.5,
        "relative_increase_threshold": 9.99,
        "min_failure_count": 123,
        "min_history_bucket_count": 123,
        "min_sample_count": 123,
        "suspect_trigger_threshold": 10,
        "auto_resolve_after_inactive_minutes": 123
      }
    },
    "auto_resolve_after_inactive_minutes": 123,
    "environments": [
      "value"
    ],
    "notifiers": [
      {
        "type": "webhook",
        "webhook_id": "value"
      }
    ]
  },
  "safety": {
    "is_blocked_ivc": true,
    "is_blocked_non_ivc": true,
    "ignore_safety_evaluation": true
  }
}`
	assertRoundTrip(t, wire, new(AgentPlatformSettingsResponseModel))
}

func TestAgentPlatformSettingsResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation": {
    "criteria": [
      {
        "id": "value",
        "name": "value",
        "type": "prompt",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scope": "conversation",
        "llm": "gpt-4o-mini",
        "scoring_mode": "binary",
        "max_score": 123,
        "score_instructions": "value"
      }
    ]
  },
  "widget": {
    "variant": "tiny",
    "placement": "top-left",
    "expandable": "never",
    "avatar": {
      "type": "orb",
      "color_1": "value",
      "color_2": "value"
    },
    "feedback_mode": "none",
    "end_feedback": {
      "type": "rating"
    },
    "bg_color": "value",
    "text_color": "value",
    "btn_color": "value",
    "btn_text_color": "value",
    "border_color": "value",
    "focus_color": "value",
    "border_radius": 123,
    "btn_radius": 123,
    "action_text": "value",
    "start_call_text": "value",
    "end_call_text": "value",
    "expand_text": "value",
    "listening_text": "value",
    "speaking_text": "value",
    "shareable_page_text": "value",
    "shareable_page_show_terms": true,
    "terms_text": "value",
    "terms_html": "value",
    "terms_key": "value",
    "show_avatar_when_collapsed": true,
    "disable_banner": true,
    "override_link": "value",
    "markdown_link_allowed_hosts": [
      {
        "hostname": "value"
      }
    ],
    "markdown_link_include_www": true,
    "markdown_link_allow_http": true,
    "mic_muting_enabled": true,
    "transcript_enabled": true,
    "text_input_enabled": true,
    "conversation_mode_toggle_enabled": true,
    "default_expanded": true,
    "always_expanded": true,
    "dismissible": true,
    "show_agent_status": true,
    "show_conversation_id": true,
    "strip_audio_tags": true,
    "syntax_highlight_theme": "light",
    "text_contents": {
      "main_label": "value",
      "start_call": "value",
      "start_chat": "value",
      "new_call": "value",
      "end_call": "value",
      "mute_microphone": "value",
      "change_language": "value",
      "collapse": "value",
      "expand": "value",
      "copied": "value",
      "accept_terms": "value",
      "dismiss_terms": "value",
      "listening_status": "value",
      "speaking_status": "value",
      "connecting_status": "value",
      "chatting_status": "value",
      "queue_waiting_status": "value",
      "queue_waiting_status_short": "value",
      "input_label": "value",
      "input_placeholder": "value",
      "input_placeholder_text_only": "value",
      "input_placeholder_new_conversation": "value",
      "user_ended_conversation": "value",
      "agent_ended_conversation": "value",
      "conversation_id": "value",
      "error_occurred": "value",
      "queue_timed_out": "value",
      "copy_id": "value",
      "initiate_feedback": "value",
      "request_follow_up_feedback": "value",
      "thanks_for_feedback": "value",
      "thanks_for_feedback_details": "value",
      "follow_up_feedback_placeholder": "value",
      "submit": "value",
      "go_back": "value",
      "send_message": "value",
      "text_mode": "value",
      "voice_mode": "value",
      "switched_to_text_mode": "value",
      "switched_to_voice_mode": "value",
      "copy": "value",
      "download": "value",
      "wrap": "value",
      "agent_working": "value",
      "agent_done": "value",
      "agent_error": "value",
      "attach_file": "value",
      "remove_file": "value",
      "file_upload_error": "value",
      "file_type_unsupported": "value",
      "file_too_large": "value",
      "file_limit_reached": "value",
      "typing_indicator": "value",
      "rich_content_unavailable": "value"
    },
    "styles": {
      "base": "value",
      "base_hover": "value",
      "base_active": "value",
      "base_border": "value",
      "base_subtle": "value",
      "base_primary": "value",
      "base_error": "value",
      "accent": "value",
      "accent_hover": "value",
      "accent_active": "value",
      "accent_border": "value",
      "accent_subtle": "value",
      "accent_primary": "value",
      "overlay_padding": 9.99,
      "button_radius": 9.99,
      "input_radius": 9.99,
      "bubble_radius": 9.99,
      "sheet_radius": 9.99,
      "compact_sheet_radius": 9.99,
      "dropdown_sheet_radius": 9.99
    },
    "show_resize_button": true,
    "end_confirmation_enabled": true,
    "mobile_fullscreen": true,
    "transcript_speech_sync_enabled": true,
    "language_selector": true,
    "supports_text_only": true,
    "custom_avatar_path": "value",
    "language_presets": {
      "key": {
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "text_contents_translation": {
          "source": {
            "key": "value"
          },
          "text": {
            "key": "value"
          }
        },
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value",
        "terms_translation": {
          "source_hash": "value",
          "text": "value"
        }
      }
    }
  },
  "data_collection": {
    "key": {
      "type": "boolean",
      "description": "value",
      "enum": [
        "value"
      ],
      "is_system_provided": true,
      "dynamic_variable": "value",
      "allowed_values": {
        "dynamic_variable": "value"
      },
      "allowed_values_dynamic_variable": "value",
      "constant_value": "value",
      "is_omitted": true,
      "name": "value",
      "llm": "gpt-4o-mini"
    }
  },
  "data_collection_scopes": {
    "key": "conversation"
  },
  "analysis_items": {
    "evaluation_criteria": [
      {
        "source": "user",
        "analysis_item_id": "value",
        "version_id": "value",
        "additional_version_ids": [
          "value"
        ],
        "scope": "conversation",
        "weight": 9.99
      }
    ],
    "data_collection": [
      {
        "source": "user",
        "analysis_item_id": "value",
        "version_id": "value",
        "scope": "conversation"
      }
    ]
  },
  "overrides": {
    "conversation_config_override": {
      "asr": {
        "keywords": true
      },
      "turn": {
        "soft_timeout_config": {
          "message": true,
          "additional_soft_timeout_messages": true
        },
        "noisy_environment_message": true
      },
      "tts": {
        "model_id": true,
        "voice_id": true,
        "supported_voices": true,
        "stability": true,
        "speed": true,
        "similarity_boost": true,
        "clone_strength": true,
        "pronunciation_dictionary_locators": true
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": true
      },
      "agent": {
        "first_message": true,
        "language": true,
        "max_conversation_duration_message": true,
        "prompt": {
          "prompt": true,
          "llm": true,
          "reasoning_effort": true,
          "tool_ids": true,
          "tool_ids_strict": true,
          "native_mcp_server_ids": true,
          "knowledge_base": true
        }
      }
    },
    "custom_llm_extra_body": true,
    "enable_conversation_initiation_client_data_from_webhook": true,
    "enable_starting_workflow_node_id_from_client": true,
    "enable_procedure_ids_from_client": true
  },
  "workspace_overrides": {
    "conversation_initiation_client_data_webhook": {
      "url": "value",
      "request_headers": {
        "key": "value"
      }
    },
    "webhooks": {
      "post_call_webhook_id": "value",
      "events": [
        "transcript"
      ],
      "transcript_format": "json",
      "exclude_transcript": true,
      "send_audio": true
    }
  },
  "testing": {
    "attached_tests": [
      {
        "test_id": "value",
        "workflow_node_id": "value"
      }
    ]
  },
  "archived": true,
  "guardrails": {
    "version": "1",
    "focus": {
      "is_enabled": true
    },
    "prompt_injection": {
      "is_enabled": true
    },
    "content": {
      "execution_mode": "streaming",
      "config": {
        "sexual": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "violence": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "harassment": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "self_harm": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "profanity": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "religion_or_politics": {
          "is_enabled": true,
          "threshold": 9.99
        },
        "medical_and_legal_information": {
          "is_enabled": true,
          "threshold": 9.99
        }
      },
      "trigger_action": {
        "type": "end_call"
      }
    },
    "custom": {
      "config": {
        "configs": [
          {
            "is_enabled": true,
            "name": "value",
            "prompt": "value",
            "execution_mode": "streaming",
            "model": "gemini-2.5-flash-lite",
            "service_tier": "standard",
            "history_message_count": 100,
            "trigger_action": {
              "type": "end_call"
            },
            "evaluate_full_response_only": true
          }
        ]
      }
    }
  },
  "summary_language": "value",
  "auto_translate_transcript_to_app_language": true,
  "auth": {
    "enable_auth": true,
    "allowlist": [
      {
        "hostname": "value"
      }
    ],
    "require_origin_header": true,
    "shareable_token": "value"
  },
  "call_limits": {
    "agent_concurrency_limit": 123,
    "daily_limit": 123,
    "bursting_enabled": true
  },
  "queueing_config": {
    "enabled": true,
    "wait_timeout_seconds": 123,
    "hold_audio": {
      "audio_path": "value",
      "audio_url": "value",
      "original_filename": "value",
      "duration_secs": 9.99,
      "size_bytes": 123
    }
  },
  "privacy": {
    "record_voice": true,
    "retention_days": 123,
    "delete_transcript_and_pii": true,
    "delete_audio": true,
    "apply_to_existing_conversations": true,
    "zero_retention_mode": true,
    "conversation_history_redaction": {
      "enabled": true,
      "entities": [
        "name"
      ],
      "excluded_data_collection_ids": [
        "value"
      ]
    }
  },
  "trust_context": "unknown",
  "analysis_llm": "gpt-4o-mini",
  "topic_discovery": {},
  "sentiment_analysis": {},
  "alerting": {
    "monitor_configs": {
      "key": {
        "enabled": true,
        "threshold": 0.5,
        "relative_increase_threshold": 9.99,
        "min_failure_count": 123,
        "min_history_bucket_count": 123,
        "min_sample_count": 123,
        "suspect_trigger_threshold": 10,
        "auto_resolve_after_inactive_minutes": 123
      }
    },
    "auto_resolve_after_inactive_minutes": 123,
    "environments": [
      "value"
    ],
    "notifiers": [
      {
        "type": "webhook",
        "webhook_id": "value"
      }
    ]
  },
  "safety": {
    "is_blocked_ivc": true,
    "is_blocked_non_ivc": true,
    "ignore_safety_evaluation": true
  }
}`
	assertRoundTrip(t, wire, new(AgentPlatformSettingsResponseModel))
}
