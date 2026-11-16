package agents

import (
	"testing"
)

func TestBodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_config": {
    "key": {}
  },
  "platform_settings": {
    "key": {}
  },
  "workflow": {
    "edges": {
      "key": {
        "source": "value",
        "target": "value",
        "forward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "backward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "reroutes": [
          {
            "x": 9.99,
            "y": 9.99
          }
        ]
      }
    },
    "nodes": {
      "key": {
        "type": "start",
        "position": {
          "x": 9.99,
          "y": 9.99
        },
        "edge_order": [
          "value"
        ]
      }
    },
    "prevent_subagent_loops": true
  },
  "name": "value",
  "tags": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost))
}
