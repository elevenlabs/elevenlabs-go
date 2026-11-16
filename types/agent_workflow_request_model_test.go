package types

import (
	"testing"
)

func TestAgentWorkflowRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AgentWorkflowRequestModel))
}

func TestAgentWorkflowRequestModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AgentWorkflowRequestModel))
}

func TestAgentWorkflowRequestModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AgentWorkflowRequestModel))
}
