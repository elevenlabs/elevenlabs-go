package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowStandaloneAgentNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order" xml:"edge_order" required:"true"`
	// The ID of the agent to transfer the conversation to. None means transfer within the current agent.
	AgentId *string `json:"agent_id" xml:"agent_id" required:"true"`
	// Optional branch ID for the target agent. None follows the target agent's traffic deployment, or uses its main branch when node_id is set.
	BranchId *string `json:"branch_id" xml:"branch_id" required:"true"`
	// Optional target node ID in the destination agent's workflow. When set, the transfer starts at this node instead of the default entry node.
	NodeId *string `json:"node_id" xml:"node_id" required:"true" minLength:"1"`
	// Artificial delay in milliseconds applied before transferring the conversation.
	DelayMs int `json:"delay_ms" xml:"delay_ms" required:"true"`
	// Optional message sent to the user before the transfer is initiated.
	TransferMessage *string `json:"transfer_message" xml:"transfer_message" required:"true"`
	// Whether to enable the transferred agent to send its configured first message after the transfer.
	EnableTransferredAgentFirstMessage bool `json:"enable_transferred_agent_first_message" xml:"enable_transferred_agent_first_message" required:"true"`
	// Defines whether TTS client overrides should be carried over to the transferred agent.
	PreserveClientTtsOverrides bool `json:"preserve_client_tts_overrides" xml:"preserve_client_tts_overrides" required:"true"`
	// Defines whether the base voice, TTS model and voice settings should be carried over to the transferred agent.
	PreserveVoiceSettings bool `json:"preserve_voice_settings" xml:"preserve_voice_settings" required:"true"`
	extraProperties       map[string]interface{}
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetAgentId() (value *string) {
	if w == nil {
		return
	}
	return w.AgentId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetBranchId() (value *string) {
	if w == nil {
		return
	}
	return w.BranchId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetNodeId() (value *string) {
	if w == nil {
		return
	}
	return w.NodeId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetDelayMs() (value int) {
	if w == nil {
		return
	}
	return w.DelayMs
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetTransferMessage() (value *string) {
	if w == nil {
		return
	}
	return w.TransferMessage
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetEnableTransferredAgentFirstMessage() (value bool) {
	if w == nil {
		return
	}
	return w.EnableTransferredAgentFirstMessage
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetPreserveClientTtsOverrides() (value bool) {
	if w == nil {
		return
	}
	return w.PreserveClientTtsOverrides
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetPreserveVoiceSettings() (value bool) {
	if w == nil {
		return
	}
	return w.PreserveVoiceSettings
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetAgentId(agentId *string) {
	w.AgentId = agentId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetBranchId(branchId *string) {
	w.BranchId = branchId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetNodeId(nodeId *string) {
	w.NodeId = nodeId
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetDelayMs(delayMs int) {
	w.DelayMs = delayMs
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetTransferMessage(transferMessage *string) {
	w.TransferMessage = transferMessage
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage bool) {
	w.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides bool) {
	w.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (w *WorkflowStandaloneAgentNodeModelOutput) SetPreserveVoiceSettings(preserveVoiceSettings bool) {
	w.PreserveVoiceSettings = preserveVoiceSettings
}

func (w *WorkflowStandaloneAgentNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowStandaloneAgentNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStandaloneAgentNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStandaloneAgentNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStandaloneAgentNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "standalone_agent" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "standalone_agent", *literals.Type)
	}
	w.type_ = "standalone_agent"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStandaloneAgentNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStandaloneAgentNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "standalone_agent",
	}
	return json.Marshal(marshaler)
}
