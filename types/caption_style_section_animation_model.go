package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleSectionAnimationModel struct {
	EnterType       CaptionStyleSectionAnimationModelEnterType `json:"enter_type" xml:"enter_type" required:"true"`
	ExitType        CaptionStyleSectionAnimationModelExitType  `json:"exit_type" xml:"exit_type" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleSectionAnimationModel) GetEnterType() (value CaptionStyleSectionAnimationModelEnterType) {
	if c == nil {
		return
	}
	return c.EnterType
}

func (c *CaptionStyleSectionAnimationModel) GetExitType() (value CaptionStyleSectionAnimationModelExitType) {
	if c == nil {
		return
	}
	return c.ExitType
}

func (c *CaptionStyleSectionAnimationModel) SetEnterType(enterType CaptionStyleSectionAnimationModelEnterType) {
	c.EnterType = enterType
}

func (c *CaptionStyleSectionAnimationModel) SetExitType(exitType CaptionStyleSectionAnimationModelExitType) {
	c.ExitType = exitType
}

func (c *CaptionStyleSectionAnimationModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleSectionAnimationModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleSectionAnimationModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleSectionAnimationModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleSectionAnimationModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleSectionAnimationModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleSectionAnimationModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleSectionAnimationModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
