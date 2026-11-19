package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleTemplateModel struct {
	Key             string `json:"key" xml:"key" required:"true"`
	Label           string `json:"label" xml:"label" required:"true"`
	RequiresHighFps *bool  `json:"requires_high_fps,omitempty" xml:"requires_high_fps,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleTemplateModel) GetKey() (value string) {
	if c == nil {
		return
	}
	return c.Key
}

func (c *CaptionStyleTemplateModel) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CaptionStyleTemplateModel) GetRequiresHighFps() (value *bool) {
	if c == nil {
		return
	}
	return c.RequiresHighFps
}

func (c *CaptionStyleTemplateModel) SetKey(key string) {
	c.Key = key
}

func (c *CaptionStyleTemplateModel) SetLabel(label string) {
	c.Label = label
}

func (c *CaptionStyleTemplateModel) SetRequiresHighFps(requiresHighFps *bool) {
	c.RequiresHighFps = requiresHighFps
}

func (c *CaptionStyleTemplateModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleTemplateModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleTemplateModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleTemplateModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleTemplateModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleTemplateModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleTemplateModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleTemplateModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
