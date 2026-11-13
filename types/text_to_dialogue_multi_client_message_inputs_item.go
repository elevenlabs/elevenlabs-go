package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TextToDialogueMultiClientMessageInputsItem struct {
	// Text to convert to speech.
	Text string `json:"text" xml:"text" required:"true"`
	// Voice ID for this segment.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Signal that this input starts a new speaker turn for prosody continuity.
	NewTurn         *bool `json:"new_turn,omitempty" xml:"new_turn,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TextToDialogueMultiClientMessageInputsItem) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TextToDialogueMultiClientMessageInputsItem) GetVoiceId() (value string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TextToDialogueMultiClientMessageInputsItem) GetNewTurn() (value *bool) {
	if t == nil {
		return
	}
	return t.NewTurn
}

func (t *TextToDialogueMultiClientMessageInputsItem) SetText(text string) {
	t.Text = text
}

func (t *TextToDialogueMultiClientMessageInputsItem) SetVoiceId(voiceId string) {
	t.VoiceId = voiceId
}

func (t *TextToDialogueMultiClientMessageInputsItem) SetNewTurn(newTurn *bool) {
	t.NewTurn = newTurn
}

func (t *TextToDialogueMultiClientMessageInputsItem) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueMultiClientMessageInputsItem) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueMultiClientMessageInputsItem to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueMultiClientMessageInputsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TextToDialogueMultiClientMessageInputsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TextToDialogueMultiClientMessageInputsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueMultiClientMessageInputsItem) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueMultiClientMessageInputsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
