package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TextToDialogueClientMessageInputsItem struct {
	// Text to convert to speech.
	Text string `json:"text" xml:"text" required:"true"`
	// Voice ID for this segment.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Signal that this input starts a new speaker turn for prosody continuity.
	NewTurn         *bool `json:"new_turn,omitempty" xml:"new_turn,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TextToDialogueClientMessageInputsItem) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TextToDialogueClientMessageInputsItem) GetVoiceId() (value string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TextToDialogueClientMessageInputsItem) GetNewTurn() (value *bool) {
	if t == nil {
		return
	}
	return t.NewTurn
}

func (t *TextToDialogueClientMessageInputsItem) SetText(text string) {
	t.Text = text
}

func (t *TextToDialogueClientMessageInputsItem) SetVoiceId(voiceId string) {
	t.VoiceId = voiceId
}

func (t *TextToDialogueClientMessageInputsItem) SetNewTurn(newTurn *bool) {
	t.NewTurn = newTurn
}

func (t *TextToDialogueClientMessageInputsItem) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueClientMessageInputsItem) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueClientMessageInputsItem to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueClientMessageInputsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TextToDialogueClientMessageInputsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TextToDialogueClientMessageInputsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueClientMessageInputsItem) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueClientMessageInputsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
