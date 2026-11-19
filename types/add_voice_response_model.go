package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddVoiceResponseModel struct {
	// The ID of the voice.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AddVoiceResponseModel) GetVoiceId() (value string) {
	if a == nil {
		return
	}
	return a.VoiceId
}

func (a *AddVoiceResponseModel) SetVoiceId(voiceId string) {
	a.VoiceId = voiceId
}

func (a *AddVoiceResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddVoiceResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddVoiceResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddVoiceResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddVoiceResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddVoiceResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddVoiceResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddVoiceResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
