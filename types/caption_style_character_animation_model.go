package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleCharacterAnimationModel struct {
	EnterType       CaptionStyleCharacterAnimationModelEnterType `json:"enter_type" xml:"enter_type" required:"true"`
	ExitType        CaptionStyleCharacterAnimationModelExitType  `json:"exit_type" xml:"exit_type" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleCharacterAnimationModel) GetEnterType() (value CaptionStyleCharacterAnimationModelEnterType) {
	if c == nil {
		return
	}
	return c.EnterType
}

func (c *CaptionStyleCharacterAnimationModel) GetExitType() (value CaptionStyleCharacterAnimationModelExitType) {
	if c == nil {
		return
	}
	return c.ExitType
}

func (c *CaptionStyleCharacterAnimationModel) SetEnterType(enterType CaptionStyleCharacterAnimationModelEnterType) {
	c.EnterType = enterType
}

func (c *CaptionStyleCharacterAnimationModel) SetExitType(exitType CaptionStyleCharacterAnimationModelExitType) {
	c.ExitType = exitType
}

func (c *CaptionStyleCharacterAnimationModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleCharacterAnimationModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleCharacterAnimationModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleCharacterAnimationModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleCharacterAnimationModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleCharacterAnimationModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleCharacterAnimationModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleCharacterAnimationModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
