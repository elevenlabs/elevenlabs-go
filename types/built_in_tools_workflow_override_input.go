package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BuiltInToolsWorkflowOverrideInput struct {
	// The transfer to agent tool
	TransferToAgent *SystemToolConfigInput `json:"transfer_to_agent,omitempty" xml:"transfer_to_agent,omitempty"`
	// The end call tool
	EndCall *SystemToolConfigInput `json:"end_call,omitempty" xml:"end_call,omitempty"`
	// The language detection tool
	LanguageDetection *SystemToolConfigInput `json:"language_detection,omitempty" xml:"language_detection,omitempty"`
	// The transfer to number tool
	TransferToNumber *SystemToolConfigInput `json:"transfer_to_number,omitempty" xml:"transfer_to_number,omitempty"`
	// The skip turn tool
	SkipTurn *SystemToolConfigInput `json:"skip_turn,omitempty" xml:"skip_turn,omitempty"`
	// The play DTMF tool
	PlayKeypadTouchTone *SystemToolConfigInput `json:"play_keypad_touch_tone,omitempty" xml:"play_keypad_touch_tone,omitempty"`
	// The voicemail detection tool
	VoicemailDetection *SystemToolConfigInput `json:"voicemail_detection,omitempty" xml:"voicemail_detection,omitempty"`
	extraProperties    map[string]interface{}
}

func (b *BuiltInToolsWorkflowOverrideInput) GetTransferToAgent() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.TransferToAgent
}

func (b *BuiltInToolsWorkflowOverrideInput) GetEndCall() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.EndCall
}

func (b *BuiltInToolsWorkflowOverrideInput) GetLanguageDetection() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.LanguageDetection
}

func (b *BuiltInToolsWorkflowOverrideInput) GetTransferToNumber() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.TransferToNumber
}

func (b *BuiltInToolsWorkflowOverrideInput) GetSkipTurn() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.SkipTurn
}

func (b *BuiltInToolsWorkflowOverrideInput) GetPlayKeypadTouchTone() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.PlayKeypadTouchTone
}

func (b *BuiltInToolsWorkflowOverrideInput) GetVoicemailDetection() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.VoicemailDetection
}

func (b *BuiltInToolsWorkflowOverrideInput) SetTransferToAgent(transferToAgent *SystemToolConfigInput) {
	b.TransferToAgent = transferToAgent
}

func (b *BuiltInToolsWorkflowOverrideInput) SetEndCall(endCall *SystemToolConfigInput) {
	b.EndCall = endCall
}

func (b *BuiltInToolsWorkflowOverrideInput) SetLanguageDetection(languageDetection *SystemToolConfigInput) {
	b.LanguageDetection = languageDetection
}

func (b *BuiltInToolsWorkflowOverrideInput) SetTransferToNumber(transferToNumber *SystemToolConfigInput) {
	b.TransferToNumber = transferToNumber
}

func (b *BuiltInToolsWorkflowOverrideInput) SetSkipTurn(skipTurn *SystemToolConfigInput) {
	b.SkipTurn = skipTurn
}

func (b *BuiltInToolsWorkflowOverrideInput) SetPlayKeypadTouchTone(playKeypadTouchTone *SystemToolConfigInput) {
	b.PlayKeypadTouchTone = playKeypadTouchTone
}

func (b *BuiltInToolsWorkflowOverrideInput) SetVoicemailDetection(voicemailDetection *SystemToolConfigInput) {
	b.VoicemailDetection = voicemailDetection
}

func (b *BuiltInToolsWorkflowOverrideInput) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BuiltInToolsWorkflowOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BuiltInToolsWorkflowOverrideInput to string"
	}
	return string(jsonData)
}

func (b *BuiltInToolsWorkflowOverrideInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BuiltInToolsWorkflowOverrideInput) MarshalJSON() ([]byte, error) {
	type embed BuiltInToolsWorkflowOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
