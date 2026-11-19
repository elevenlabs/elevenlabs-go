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
type BuiltInToolsOutput struct {
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

func (b *BuiltInToolsOutput) GetTransferToAgent() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.TransferToAgent
}

func (b *BuiltInToolsOutput) GetEndCall() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.EndCall
}

func (b *BuiltInToolsOutput) GetLanguageDetection() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.LanguageDetection
}

func (b *BuiltInToolsOutput) GetTransferToNumber() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.TransferToNumber
}

func (b *BuiltInToolsOutput) GetSkipTurn() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.SkipTurn
}

func (b *BuiltInToolsOutput) GetPlayKeypadTouchTone() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.PlayKeypadTouchTone
}

func (b *BuiltInToolsOutput) GetVoicemailDetection() (value *SystemToolConfigOutput) {
	if b == nil {
		return
	}
	return b.VoicemailDetection
}

func (b *BuiltInToolsOutput) SetTransferToAgent(transferToAgent *SystemToolConfigOutput) {
	b.TransferToAgent = transferToAgent
}

func (b *BuiltInToolsOutput) SetEndCall(endCall *SystemToolConfigOutput) {
	b.EndCall = endCall
}

func (b *BuiltInToolsOutput) SetLanguageDetection(languageDetection *SystemToolConfigOutput) {
	b.LanguageDetection = languageDetection
}

func (b *BuiltInToolsOutput) SetTransferToNumber(transferToNumber *SystemToolConfigOutput) {
	b.TransferToNumber = transferToNumber
}

func (b *BuiltInToolsOutput) SetSkipTurn(skipTurn *SystemToolConfigOutput) {
	b.SkipTurn = skipTurn
}

func (b *BuiltInToolsOutput) SetPlayKeypadTouchTone(playKeypadTouchTone *SystemToolConfigOutput) {
	b.PlayKeypadTouchTone = playKeypadTouchTone
}

func (b *BuiltInToolsOutput) SetVoicemailDetection(voicemailDetection *SystemToolConfigOutput) {
	b.VoicemailDetection = voicemailDetection
}

func (b *BuiltInToolsOutput) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BuiltInToolsOutput) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BuiltInToolsOutput to string"
	}
	return string(jsonData)
}

func (b *BuiltInToolsOutput) UnmarshalJSON(data []byte) error {
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

func (b *BuiltInToolsOutput) MarshalJSON() ([]byte, error) {
	type embed BuiltInToolsOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
