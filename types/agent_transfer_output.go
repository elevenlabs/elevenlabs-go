package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentTransferOutput struct {
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

func (a *AgentTransferOutput) GetAgentId() (value *string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentTransferOutput) GetBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentTransferOutput) GetNodeId() (value *string) {
	if a == nil {
		return
	}
	return a.NodeId
}

func (a *AgentTransferOutput) GetCondition() (value string) {
	if a == nil {
		return
	}
	return a.Condition
}

func (a *AgentTransferOutput) GetDelayMs() (value *int) {
	if a == nil {
		return
	}
	return a.DelayMs
}

func (a *AgentTransferOutput) GetTransferMessage() (value *string) {
	if a == nil {
		return
	}
	return a.TransferMessage
}

func (a *AgentTransferOutput) GetEnableTransferredAgentFirstMessage() (value *bool) {
	if a == nil {
		return
	}
	return a.EnableTransferredAgentFirstMessage
}

func (a *AgentTransferOutput) GetIsWorkflowNodeTransfer() (value *bool) {
	if a == nil {
		return
	}
	return a.IsWorkflowNodeTransfer
}

func (a *AgentTransferOutput) GetPreserveClientTtsOverrides() (value *bool) {
	if a == nil {
		return
	}
	return a.PreserveClientTtsOverrides
}

func (a *AgentTransferOutput) GetPreserveVoiceSettings() (value *bool) {
	if a == nil {
		return
	}
	return a.PreserveVoiceSettings
}

func (a *AgentTransferOutput) SetAgentId(agentId *string) {
	a.AgentId = agentId
}

func (a *AgentTransferOutput) SetBranchId(branchId *string) {
	a.BranchId = branchId
}

func (a *AgentTransferOutput) SetNodeId(nodeId *string) {
	a.NodeId = nodeId
}

func (a *AgentTransferOutput) SetCondition(condition string) {
	a.Condition = condition
}

func (a *AgentTransferOutput) SetDelayMs(delayMs *int) {
	a.DelayMs = delayMs
}

func (a *AgentTransferOutput) SetTransferMessage(transferMessage *string) {
	a.TransferMessage = transferMessage
}

func (a *AgentTransferOutput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage *bool) {
	a.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (a *AgentTransferOutput) SetIsWorkflowNodeTransfer(isWorkflowNodeTransfer *bool) {
	a.IsWorkflowNodeTransfer = isWorkflowNodeTransfer
}

func (a *AgentTransferOutput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides *bool) {
	a.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (a *AgentTransferOutput) SetPreserveVoiceSettings(preserveVoiceSettings *bool) {
	a.PreserveVoiceSettings = preserveVoiceSettings
}

func (a *AgentTransferOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTransferOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTransferOutput to string"
	}
	return string(jsonData)
}

func (a *AgentTransferOutput) UnmarshalJSON(data []byte) error {
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

func (a *AgentTransferOutput) MarshalJSON() ([]byte, error) {
	type embed AgentTransferOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
