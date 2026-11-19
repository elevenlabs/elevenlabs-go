package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ImageSubject struct {
	Name            string `json:"name" xml:"name" required:"true"`
	Description     string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (i *ImageSubject) GetName() (value string) {
	if i == nil {
		return
	}
	return i.Name
}

func (i *ImageSubject) GetDescription() (value string) {
	if i == nil {
		return
	}
	return i.Description
}

func (i *ImageSubject) SetName(name string) {
	i.Name = name
}

func (i *ImageSubject) SetDescription(description string) {
	i.Description = description
}

func (i *ImageSubject) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i ImageSubject) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: ImageSubject to string"
	}
	return string(jsonData)
}

func (i *ImageSubject) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias ImageSubject
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = ImageSubject(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *ImageSubject) MarshalJSON() ([]byte, error) {
	type embed ImageSubject
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
