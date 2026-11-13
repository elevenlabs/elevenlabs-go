package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToAgentToolResultSuccessModelInput struct {
	Status                             *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	FromAgent                          string                                                `json:"from_agent" xml:"from_agent" required:"true"`
	ToAgent                            string                                                `json:"to_agent" xml:"to_agent" required:"true"`
	ToNode                             *string                                               `json:"to_node,omitempty" xml:"to_node,omitempty"`
	Condition                          string                                                `json:"condition" xml:"condition" required:"true"`
	DelayMs                            *int                                                  `json:"delay_ms,omitempty" xml:"delay_ms,omitempty"`
	TransferMessage                    *string                                               `json:"transfer_message,omitempty" xml:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool                                                 `json:"enable_transferred_agent_first_message,omitempty" xml:"enable_transferred_agent_first_message,omitempty"`
	BranchInfo                         *TransferToAgentToolResultSuccessModelInputBranchInfo `json:"branch_info,omitempty" xml:"branch_info,omitempty"`
	PreserveClientTtsOverrides         *bool                                                 `json:"preserve_client_tts_overrides,omitempty" xml:"preserve_client_tts_overrides,omitempty"`
	PreserveVoiceSettings              *bool                                                 `json:"preserve_voice_settings,omitempty" xml:"preserve_voice_settings,omitempty"`
	extraProperties                    map[string]interface{}
}

func (t *TransferToAgentToolResultSuccessModelInput) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToAgentToolResultSuccessModelInput) GetFromAgent() (value string) {
	if t == nil {
		return
	}
	return t.FromAgent
}

func (t *TransferToAgentToolResultSuccessModelInput) GetToAgent() (value string) {
	if t == nil {
		return
	}
	return t.ToAgent
}

func (t *TransferToAgentToolResultSuccessModelInput) GetToNode() (value *string) {
	if t == nil {
		return
	}
	return t.ToNode
}

func (t *TransferToAgentToolResultSuccessModelInput) GetCondition() (value string) {
	if t == nil {
		return
	}
	return t.Condition
}

func (t *TransferToAgentToolResultSuccessModelInput) GetDelayMs() (value *int) {
	if t == nil {
		return
	}
	return t.DelayMs
}

func (t *TransferToAgentToolResultSuccessModelInput) GetTransferMessage() (value *string) {
	if t == nil {
		return
	}
	return t.TransferMessage
}

func (t *TransferToAgentToolResultSuccessModelInput) GetEnableTransferredAgentFirstMessage() (value *bool) {
	if t == nil {
		return
	}
	return t.EnableTransferredAgentFirstMessage
}

func (t *TransferToAgentToolResultSuccessModelInput) GetBranchInfo() (value *TransferToAgentToolResultSuccessModelInputBranchInfo) {
	if t == nil {
		return
	}
	return t.BranchInfo
}

func (t *TransferToAgentToolResultSuccessModelInput) GetPreserveClientTtsOverrides() (value *bool) {
	if t == nil {
		return
	}
	return t.PreserveClientTtsOverrides
}

func (t *TransferToAgentToolResultSuccessModelInput) GetPreserveVoiceSettings() (value *bool) {
	if t == nil {
		return
	}
	return t.PreserveVoiceSettings
}

func (t *TransferToAgentToolResultSuccessModelInput) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToAgentToolResultSuccessModelInput) SetFromAgent(fromAgent string) {
	t.FromAgent = fromAgent
}

func (t *TransferToAgentToolResultSuccessModelInput) SetToAgent(toAgent string) {
	t.ToAgent = toAgent
}

func (t *TransferToAgentToolResultSuccessModelInput) SetToNode(toNode *string) {
	t.ToNode = toNode
}

func (t *TransferToAgentToolResultSuccessModelInput) SetCondition(condition string) {
	t.Condition = condition
}

func (t *TransferToAgentToolResultSuccessModelInput) SetDelayMs(delayMs *int) {
	t.DelayMs = delayMs
}

func (t *TransferToAgentToolResultSuccessModelInput) SetTransferMessage(transferMessage *string) {
	t.TransferMessage = transferMessage
}

func (t *TransferToAgentToolResultSuccessModelInput) SetEnableTransferredAgentFirstMessage(enableTransferredAgentFirstMessage *bool) {
	t.EnableTransferredAgentFirstMessage = enableTransferredAgentFirstMessage
}

func (t *TransferToAgentToolResultSuccessModelInput) SetBranchInfo(branchInfo *TransferToAgentToolResultSuccessModelInputBranchInfo) {
	t.BranchInfo = branchInfo
}

func (t *TransferToAgentToolResultSuccessModelInput) SetPreserveClientTtsOverrides(preserveClientTtsOverrides *bool) {
	t.PreserveClientTtsOverrides = preserveClientTtsOverrides
}

func (t *TransferToAgentToolResultSuccessModelInput) SetPreserveVoiceSettings(preserveVoiceSettings *bool) {
	t.PreserveVoiceSettings = preserveVoiceSettings
}

func (t *TransferToAgentToolResultSuccessModelInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToAgentToolResultSuccessModelInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToAgentToolResultSuccessModelInput to string"
	}
	return string(jsonData)
}

func (t *TransferToAgentToolResultSuccessModelInput) UnmarshalJSON(data []byte) error {
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

func (t *TransferToAgentToolResultSuccessModelInput) MarshalJSON() ([]byte, error) {
	type embed TransferToAgentToolResultSuccessModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
