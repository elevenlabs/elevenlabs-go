package types

import (
	"testing"
)

func TestConversationHistoryAnalysisCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation_criteria_results": {
    "key": {
      "criteria_id": "value",
      "result": "success",
      "rationale": "value",
      "scoring_mode": "binary",
      "score": 123,
      "max_score": 123
    }
  },
  "data_collection_results": {
    "key": {
      "data_collection_id": "value",
      "name": "value",
      "value": {},
      "json_schema": {
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
        "is_omitted": true
      },
      "rationale": "value"
    }
  },
  "evaluation_criteria_results_list": [
    {
      "criteria_id": "value",
      "result": "success",
      "rationale": "value",
      "scoring_mode": "binary",
      "score": 123,
      "max_score": 123
    }
  ],
  "data_collection_results_list": [
    {
      "data_collection_id": "value",
      "name": "value",
      "value": {},
      "json_schema": {
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
        "is_omitted": true
      },
      "rationale": "value"
    }
  ],
  "call_successful": "success",
  "call_success_score": 9.99,
  "transcript_summary": "value",
  "call_summary_title": "value",
  "scoped": [
    {
      "scope": "conversation",
      "source_agent_id": "value",
      "source_branch_id": "value",
      "evaluation_criteria_results": {
        "key": {
          "criteria_id": "value",
          "result": "success",
          "rationale": "value",
          "scoring_mode": "binary",
          "score": 123,
          "max_score": 123
        }
      },
      "data_collection_results": {
        "key": {
          "data_collection_id": "value",
          "name": "value",
          "value": {},
          "json_schema": {
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
            "is_omitted": true
          },
          "rationale": "value"
        }
      },
      "successful": "success",
      "success_score": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryAnalysisCommonModel))
}

func TestConversationHistoryAnalysisCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation_criteria_results": {
    "key": {
      "criteria_id": "value",
      "result": "success",
      "rationale": "value",
      "scoring_mode": "binary",
      "score": 123,
      "max_score": 123
    }
  },
  "data_collection_results": {
    "key": {
      "data_collection_id": "value",
      "name": "value",
      "value": {},
      "json_schema": {
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
        "is_omitted": true
      },
      "rationale": "value"
    }
  },
  "evaluation_criteria_results_list": [
    {
      "criteria_id": "value",
      "result": "success",
      "rationale": "value",
      "scoring_mode": "binary",
      "score": 123,
      "max_score": 123
    }
  ],
  "data_collection_results_list": [
    {
      "data_collection_id": "value",
      "name": "value",
      "value": {},
      "json_schema": {
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
        "is_omitted": true
      },
      "rationale": "value"
    }
  ],
  "call_successful": "success",
  "call_success_score": 9.99,
  "transcript_summary": "value",
  "call_summary_title": "value",
  "scoped": [
    {
      "scope": "conversation",
      "source_agent_id": "value",
      "source_branch_id": "value",
      "evaluation_criteria_results": {
        "key": {
          "criteria_id": "value",
          "result": "success",
          "rationale": "value",
          "scoring_mode": "binary",
          "score": 123,
          "max_score": 123
        }
      },
      "data_collection_results": {
        "key": {
          "data_collection_id": "value",
          "name": "value",
          "value": {},
          "json_schema": {
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
            "is_omitted": true
          },
          "rationale": "value"
        }
      },
      "successful": "success",
      "success_score": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryAnalysisCommonModel))
}
