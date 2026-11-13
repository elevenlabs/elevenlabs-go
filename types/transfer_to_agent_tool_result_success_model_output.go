package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToAgentToolResultSuccessModelOutput struct {
	Status                             *string                                                `json:"status,omitempty" xml:"status,omitempty"`
	FromAgent                          string                                                 `json:"from_agent" xml:"from_agent" required:"true"`
	ToAgent                            string                                                 `json:"to_agent" xml:"to_agent" required:"true"`
	ToNode                             *string                                                `json:"to_node,omitempty" xml:"to_node,omitempty"`
	Condition                          string                                                 `json:"condition" xml:"condition" required:"true"`
	DelayMs                            *int                                                   `json:"delay_ms,omitempty" xml:"delay_ms,omitempty"`
	TransferMessage                    *string                                                `json:"transfer_message,omitempty" xml:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool                                                  `json:"enable_transferred_agent_first_message,omitempty" xml:"enable_transferred_agent_first_message,omitempty"`
	BranchInfo                         *TransferToAgentToolResultSuccessModelOutputBranchInfo `json:"branch_info,omitempty" xml:"branch_info,omitempty"`
	PreserveClientTtsOverrides         *bool                                                  `json:"preserve_client_tts_overrides,omitempty" xml:"preserve_client_tts_overrides,omitempty"`
	PreserveVoiceSettings              *bool                                                  `json:"preserve_voice_settings,omitempty" xml:"preserve_voice_settings,omitempty"`
	extraProperties                    map[string]interface{}
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetFromAgent() (value string) {
	if t == nil {
		return
	}
	return t.FromAgent
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetToAgent() (value string) {
	if t == nil {
		return
	}
	return t.ToAgent
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetToNode() (value *string) {
	if t == nil {
		return
	}
	return t.ToNode
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetCondition() (value string) {
	if t == nil {
		return
	}
	return t.Condition
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetDelayMs() (value *int) {
	if t == nil {
		return
	}
	return t.DelayMs
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetTransferMessage() (value *string) {
	if t == nil {
		return
	}
	return t.TransferMessage
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetEnableTransferredAgentFirstMessage() (value *bool) {
	if t == nil {
		return
	}
	return t.EnableTransferredAgentFirstMessage
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetBranchInfo() (value *TransferToAgentToolResultSuccessModelOutputBranchInfo) {
	if t == nil {
		return
	}
	return t.BranchInfo
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetPreserveClientTtsOverrides() (value *bool) {
	if t == nil {
		return
	}
	return t.PreserveClientTtsOverrides
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetPreserveVoiceSettings() (value *bool) {
	if t == nil {
		return
	}
	return t.PreserveVoiceSettings
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetFromAgent(fromAgent string) {
	t.FromAgent = fromAgent
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetToAgent(toAgent string) {
	t.ToAgent = toAgent
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetToNode(toNode *string) {
	t.ToNode = toNode
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetCondition(condition string) {
	t.Condition = condition
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetDelayMs(delayMs *int) {
	t.DelayMs = delayMs
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetTransferMessage(transferMessage *string) {
	t.TransferMessage = transferMessage
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage *bool) {
	t.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetBranchInfo(branchInfo *TransferToAgentToolResultSuccessModelOutputBranchInfo) {
	t.BranchInfo = branchInfo
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides *bool) {
	t.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (t *TransferToAgentToolResultSuccessModelOutput) SetPreserveVoiceSettings(preserveVoiceSettings *bool) {
	t.PreserveVoiceSettings = preserveVoiceSettings
}

func (t *TransferToAgentToolResultSuccessModelOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToAgentToolResultSuccessModelOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToAgentToolResultSuccessModelOutput to string"
	}
	return string(jsonData)
}

func (t *TransferToAgentToolResultSuccessModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "result_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToAgentToolResultSuccessModelOutput) MarshalJSON() ([]byte, error) {
	type embed TransferToAgentToolResultSuccessModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
