package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DubbingVoiceSettings struct {
	// How strongly the dubbed speakers clone the source voices, 0 to 10.
	CloningStrength *int `json:"cloning_strength,omitempty" xml:"cloning_strength,omitempty" min:"0" max:"10"`
	extraProperties map[string]interface{}
}

func (d *DubbingVoiceSettings) GetCloningStrength() (value *int) {
	if d == nil {
		return
	}
	return d.CloningStrength
}

func (d *DubbingVoiceSettings) SetCloningStrength(cloningStrength *int) {
	d.CloningStrength = cloningStrength
}

func (d *DubbingVoiceSettings) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingVoiceSettings to string"
	}
	return string(jsonData)
}

func (d *DubbingVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler DubbingVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DubbingVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed DubbingVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
