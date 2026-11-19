package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DialogueInputResponseModel struct {
	// The text of the dialogue input line.
	Text string `json:"text" xml:"text" required:"true"`
	// The ID of the voice used for this dialogue input line.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// The name of the voice used for this dialogue input line.
	VoiceName       string `json:"voice_name" xml:"voice_name" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DialogueInputResponseModel) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DialogueInputResponseModel) GetVoiceId() (value string) {
	if d == nil {
		return
	}
	return d.VoiceId
}

func (d *DialogueInputResponseModel) GetVoiceName() (value string) {
	if d == nil {
		return
	}
	return d.VoiceName
}

func (d *DialogueInputResponseModel) SetText(text string) {
	d.Text = text
}

func (d *DialogueInputResponseModel) SetVoiceId(voiceId string) {
	d.VoiceId = voiceId
}

func (d *DialogueInputResponseModel) SetVoiceName(voiceName string) {
	d.VoiceName = voiceName
}

func (d *DialogueInputResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DialogueInputResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DialogueInputResponseModel to string"
	}
	return string(jsonData)
}

func (d *DialogueInputResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DialogueInputResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DialogueInputResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DialogueInputResponseModel) MarshalJSON() ([]byte, error) {
	type embed DialogueInputResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
