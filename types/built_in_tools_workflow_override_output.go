package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BuiltInToolsWorkflowOverrideOutput struct {
	// The transfer to agent tool
	TransferToAgent *SystemToolConfigOutput `json:"transfer_to_agent,omitempty" xml:"transfer_to_agent,omitempty"`
	// The end call tool
	EndCall *SystemToolConfigOutput `json:"end_call,omitempty" xml:"end_call,omitempty"`
	// The language detection tool
	LanguageDetection *SystemToolConfigOutput `json:"language_detection,omitempty" xml:"language_detection,omitempty"`
	// The transfer to number tool
	TransferToNumber *SystemToolConfigOutput `json:"transfer_to_number,omitempty" xml:"transfer_to_number,omitempty"`
	// The skip turn tool
	SkipTurn *SystemToolConfigOutput `json:"skip_turn,omitempty" xml:"skip_turn,omitempty"`
	// The play DTMF tool
	PlayKeypadTouchTone *SystemToolConfigOutput `json:"play_keypad_touch_tone,omitempty" xml:"play_keypad_touch_tone,omitempty"`
	// The voicemail detection tool
	VoicemailDetection *SystemToolConfigOutput `json:"voicemail_detection,omitempty" xml:"voicemail_detection,omitempty"`
	extraProperties    map[string]interface{}
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetTransferToAgent() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.TransferToAgent
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetEndCall() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.EndCall
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetLanguageDetection() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.LanguageDetection
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetTransferToNumber() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.TransferToNumber
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetSkipTurn() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.SkipTurn
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetPlayKeypadTouchTone() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.PlayKeypadTouchTone
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetVoicemailDetection() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.VoicemailDetection
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetTransferToAgent(transferToAgent *SystemToolConfigOutput) {
	b.TransferToAgent = transferToAgent
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetEndCall(endCall *SystemToolConfigOutput) {
	b.EndCall = endCall
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetLanguageDetection(languageDetection *SystemToolConfigOutput) {
	b.LanguageDetection = languageDetection
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetTransferToNumber(transferToNumber *SystemToolConfigOutput) {
	b.TransferToNumber = transferToNumber
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetSkipTurn(skipTurn *SystemToolConfigOutput) {
	b.SkipTurn = skipTurn
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetPlayKeypadTouchTone(playKeypadTouchTone *SystemToolConfigOutput) {
	b.PlayKeypadTouchTone = playKeypadTouchTone
}

func (b *BuiltInToolsWorkflowOverrideOutput) SetVoicemailDetection(voicemailDetection *SystemToolConfigOutput) {
	b.VoicemailDetection = voicemailDetection
}

func (b *BuiltInToolsWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BuiltInToolsWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BuiltInToolsWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (b *BuiltInToolsWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (b *BuiltInToolsWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed BuiltInToolsWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
