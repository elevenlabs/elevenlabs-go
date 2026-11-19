package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowStandaloneAgentNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// The ID of the agent to transfer the conversation to. None means transfer within the current agent.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// Optional branch ID for the target agent. None follows the target agent's traffic deployment, or uses its main branch when node_id is set.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// Optional target node ID in the destination agent's workflow. When set, the transfer starts at this node instead of the default entry node.
	NodeId *string `json:"node_id,omitempty" xml:"node_id,omitempty" minLength:"1"`
	// Artificial delay in milliseconds applied before transferring the conversation.
	DelayMs *int `json:"delay_ms,omitempty" xml:"delay_ms,omitempty"`
	// Optional message sent to the user before the transfer is initiated.
	TransferMessage *string `json:"transfer_message,omitempty" xml:"transfer_message,omitempty"`
	// Whether to enable the transferred agent to send its configured first message after the transfer.
	EnableTransferredAgentFirstMessage *bool `json:"enable_transferred_agent_first_message,omitempty" xml:"enable_transferred_agent_first_message,omitempty"`
	// Defines whether TTS client overrides should be carried over to the transferred agent.
	PreserveClientTtsOverrides *bool `json:"preserve_client_tts_overrides,omitempty" xml:"preserve_client_tts_overrides,omitempty"`
	// Defines whether the base voice, TTS model and voice settings should be carried over to the transferred agent.
	PreserveVoiceSettings *bool `json:"preserve_voice_settings,omitempty" xml:"preserve_voice_settings,omitempty"`
	extraProperties       map[string]interface{}
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetAgentId() (value *string) {
	if w == nil {
		return
	}
	return w.AgentId
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetBranchId() (value *string) {
	if w == nil {
		return
	}
	return w.BranchId
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetNodeId() (value *string) {
	if w == nil {
		return
	}
	return w.NodeId
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetDelayMs() (value *int) {
	if w == nil {
		return
	}
	return w.DelayMs
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetTransferMessage() (value *string) {
	if w == nil {
		return
	}
	return w.TransferMessage
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetEnableTransferredAgentFirstMessage() (value *bool) {
	if w == nil {
		return
	}
	return w.EnableTransferredAgentFirstMessage
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetPreserveClientTtsOverrides() (value *bool) {
	if w == nil {
		return
	}
	return w.PreserveClientTtsOverrides
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetPreserveVoiceSettings() (value *bool) {
	if w == nil {
		return
	}
	return w.PreserveVoiceSettings
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetAgentId(agentId *string) {
	w.AgentId = agentId
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetBranchId(branchId *string) {
	w.BranchId = branchId
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetNodeId(nodeId *string) {
	w.NodeId = nodeId
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetDelayMs(delayMs *int) {
	w.DelayMs = delayMs
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetTransferMessage(transferMessage *string) {
	w.TransferMessage = transferMessage
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage *bool) {
	w.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides *bool) {
	w.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (w *WorkflowStandaloneAgentNodeModelInput) SetPreserveVoiceSettings(preserveVoiceSettings *bool) {
	w.PreserveVoiceSettings = preserveVoiceSettings
}

func (w *WorkflowStandaloneAgentNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStandaloneAgentNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStandaloneAgentNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStandaloneAgentNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStandaloneAgentNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStandaloneAgentNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
