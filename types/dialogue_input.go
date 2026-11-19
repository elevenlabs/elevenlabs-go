package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DialogueInput struct {
	// The text to be converted into speech.
	Text string `json:"text" xml:"text" required:"true"`
	// The ID of the voice to be used for the generation.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DialogueInput) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DialogueInput) GetVoiceId() (value string) {
	if d == nil {
		return
	}
	return d.VoiceId
}

func (d *DialogueInput) SetText(text string) {
	d.Text = text
}

func (d *DialogueInput) SetVoiceId(voiceId string) {
	d.VoiceId = voiceId
}

func (d *DialogueInput) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DialogueInput) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DialogueInput to string"
	}
	return string(jsonData)
}

func (d *DialogueInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DialogueInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DialogueInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DialogueInput) MarshalJSON() ([]byte, error) {
	type embed DialogueInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
