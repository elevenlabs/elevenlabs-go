package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoSubject struct {
	Name            string `json:"name" xml:"name" required:"true"`
	Description     string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VideoSubject) GetName() (value string) {
	if v == nil {
		return
	}
	return v.Name
}

func (v *VideoSubject) GetDescription() (value string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VideoSubject) SetName(name string) {
	v.Name = name
}

func (v *VideoSubject) SetDescription(description string) {
	v.Description = description
}

func (v *VideoSubject) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoSubject) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoSubject to string"
	}
	return string(jsonData)
}

func (v *VideoSubject) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VideoSubject
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VideoSubject(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoSubject) MarshalJSON() ([]byte, error) {
	type embed VideoSubject
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
