package agents

import (
	"testing"
)

func TestBodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parent_version_id": "value",
  "name": "value",
  "description": "value",
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
  "include_draft": true
}`
	assertRoundTrip(t, wire, new(BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost))
}

func TestBodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "description": "value",
  "is_archived": true,
  "protection_status": "writer_perms_required",
  "auto_rebase": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch))
}

func TestBodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "archive_source_branch": true,
  "force": true,
  "conflict_resolutions": [
    {
      "conflict_id": "value",
      "choice": "source",
      "custom_value": "value"
    }
  ],
  "source_tip_id": "value",
  "target_tip_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost))
}

func TestBodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conflict_resolutions": [
    {
      "conflict_id": "value",
      "choice": "source",
      "custom_value": "value"
    }
  ],
  "branch_tip_id": "value",
  "source_tip_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost))
}
