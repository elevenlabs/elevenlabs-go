package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// System tools a conversational agent can be given.
//
// Deliberately not named ConversationalBuiltInTools: BuiltInToolsInput and
// BuiltInToolsOutput are part of the public API spec and the generated SDKs.
type BuiltInToolsInput struct {
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

func (b *BuiltInToolsInput) GetTransferToAgent() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.TransferToAgent
}

func (b *BuiltInToolsInput) GetEndCall() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.EndCall
}

func (b *BuiltInToolsInput) GetLanguageDetection() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.LanguageDetection
}

func (b *BuiltInToolsInput) GetTransferToNumber() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.TransferToNumber
}

func (b *BuiltInToolsInput) GetSkipTurn() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.SkipTurn
}

func (b *BuiltInToolsInput) GetPlayKeypadTouchTone() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.PlayKeypadTouchTone
}

func (b *BuiltInToolsInput) GetVoicemailDetection() (value *SystemToolConfigInput) {
	if b == nil {
		return
	}
	return b.VoicemailDetection
}

func (b *BuiltInToolsInput) SetTransferToAgent(transferToAgent *SystemToolConfigInput) {
	b.TransferToAgent = transferToAgent
}

func (b *BuiltInToolsInput) SetEndCall(endCall *SystemToolConfigInput) {
	b.EndCall = endCall
}

func (b *BuiltInToolsInput) SetLanguageDetection(languageDetection *SystemToolConfigInput) {
	b.LanguageDetection = languageDetection
}

func (b *BuiltInToolsInput) SetTransferToNumber(transferToNumber *SystemToolConfigInput) {
	b.TransferToNumber = transferToNumber
}

func (b *BuiltInToolsInput) SetSkipTurn(skipTurn *SystemToolConfigInput) {
	b.SkipTurn = skipTurn
}

func (b *BuiltInToolsInput) SetPlayKeypadTouchTone(playKeypadTouchTone *SystemToolConfigInput) {
	b.PlayKeypadTouchTone = playKeypadTouchTone
}

func (b *BuiltInToolsInput) SetVoicemailDetection(voicemailDetection *SystemToolConfigInput) {
	b.VoicemailDetection = voicemailDetection
}

func (b *BuiltInToolsInput) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BuiltInToolsInput) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BuiltInToolsInput to string"
	}
	return string(jsonData)
}

func (b *BuiltInToolsInput) UnmarshalJSON(data []byte) error {
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

func (b *BuiltInToolsInput) MarshalJSON() ([]byte, error) {
	type embed BuiltInToolsInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
