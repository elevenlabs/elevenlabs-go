package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DubbingLanguageResponseVoiceSettings struct {
	// How strongly the dubbed speakers clone the source voices, 0 to 10.
	CloningStrength *int `json:"cloning_strength,omitempty" xml:"cloning_strength,omitempty" min:"0" max:"10"`
	extraProperties map[string]interface{}
}

func (d *DubbingLanguageResponseVoiceSettings) GetCloningStrength() (value *int) {
	if d == nil {
		return
	}
	return d.CloningStrength
}

func (d *DubbingLanguageResponseVoiceSettings) SetCloningStrength(cloningStrength *int) {
	d.CloningStrength = cloningStrength
}

func (d *DubbingLanguageResponseVoiceSettings) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingLanguageResponseVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingLanguageResponseVoiceSettings to string"
	}
	return string(jsonData)
}

func (d *DubbingLanguageResponseVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler DubbingLanguageResponseVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DubbingLanguageResponseVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingLanguageResponseVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed DubbingLanguageResponseVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
