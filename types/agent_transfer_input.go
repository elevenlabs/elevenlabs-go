package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentTransferInput struct {
	AgentId                            *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	BranchId                           *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	NodeId                             *string `json:"node_id,omitempty" xml:"node_id,omitempty"`
	Condition                          string  `json:"condition" xml:"condition" required:"true"`
	DelayMs                            *int    `json:"delay_ms,omitempty" xml:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty" xml:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty" xml:"enable_transferred_agent_first_message,omitempty"`
	IsWorkflowNodeTransfer             *bool   `json:"is_workflow_node_transfer,omitempty" xml:"is_workflow_node_transfer,omitempty"`
	// Defines whether TTS client overrides should be carried over to the transferred agent.
	PreserveClientTtsOverrides *bool `json:"preserve_client_tts_overrides,omitempty" xml:"preserve_client_tts_overrides,omitempty"`
	// Defines whether the base voice, TTS model and voice settings should be carried over to the transferred agent.
	PreserveVoiceSettings *bool `json:"preserve_voice_settings,omitempty" xml:"preserve_voice_settings,omitempty"`
	extraProperties       map[string]interface{}
}

func (a *AgentTransferInput) GetAgentId() (value *string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentTransferInput) GetBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentTransferInput) GetNodeId() (value *string) {
	if a == nil {
		return
	}
	return a.NodeId
}

func (a *AgentTransferInput) GetCondition() (value string) {
	if a == nil {
		return
	}
	return a.Condition
}

func (a *AgentTransferInput) GetDelayMs() (value *int) {
	if a == nil {
		return
	}
	return a.DelayMs
}

func (a *AgentTransferInput) GetTransferMessage() (value *string) {
	if a == nil {
		return
	}
	return a.TransferMessage
}

func (a *AgentTransferInput) GetEnableTransferredAgentFirstMessage() (value *bool) {
	if a == nil {
		return
	}
	return a.EnableTransferredAgentFirstMessage
}

func (a *AgentTransferInput) GetIsWorkflowNodeTransfer() (value *bool) {
	if a == nil {
		return
	}
	return a.IsWorkflowNodeTransfer
}

func (a *AgentTransferInput) GetPreserveClientTtsOverrides() (value *bool) {
	if a == nil {
		return
	}
	return a.PreserveClientTtsOverrides
}

func (a *AgentTransferInput) GetPreserveVoiceSettings() (value *bool) {
	if a == nil {
		return
	}
	return a.PreserveVoiceSettings
}

func (a *AgentTransferInput) SetAgentId(agentId *string) {
	a.AgentId = agentId
}

func (a *AgentTransferInput) SetBranchId(branchId *string) {
	a.BranchId = branchId
}

func (a *AgentTransferInput) SetNodeId(nodeId *string) {
	a.NodeId = nodeId
}

func (a *AgentTransferInput) SetCondition(condition string) {
	a.Condition = condition
}

func (a *AgentTransferInput) SetDelayMs(delayMs *int) {
	a.DelayMs = delayMs
}

func (a *AgentTransferInput) SetTransferMessage(transferMessage *string) {
	a.TransferMessage = transferMessage
}

func (a *AgentTransferInput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage *bool) {
	a.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (a *AgentTransferInput) SetIsWorkflowNodeTransfer(isWorkflowNodeTransfer *bool) {
	a.IsWorkflowNodeTransfer = isWorkflowNodeTransfer
}

func (a *AgentTransferInput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides *bool) {
	a.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (a *AgentTransferInput) SetPreserveVoiceSettings(preserveVoiceSettings *bool) {
	a.PreserveVoiceSettings = preserveVoiceSettings
}

func (a *AgentTransferInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTransferInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTransferInput to string"
	}
	return string(jsonData)
}

func (a *AgentTransferInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTransferInput) MarshalJSON() ([]byte, error) {
	type embed AgentTransferInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
