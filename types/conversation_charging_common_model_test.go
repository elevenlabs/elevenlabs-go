package types

import (
	"testing"
)

func TestConversationChargingCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ConversationChargingCommonModel))
}

func TestConversationChargingCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ConversationChargingCommonModel))
}
