package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleWordAnimationModel struct {
	EnterType       CaptionStyleWordAnimationModelEnterType `json:"enter_type" xml:"enter_type" required:"true"`
	ExitType        CaptionStyleWordAnimationModelExitType  `json:"exit_type" xml:"exit_type" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleWordAnimationModel) GetEnterType() (value CaptionStyleWordAnimationModelEnterType) {
	if c == nil {
		return
	}
	return c.EnterType
}

func (c *CaptionStyleWordAnimationModel) GetExitType() (value CaptionStyleWordAnimationModelExitType) {
	if c == nil {
		return
	}
	return c.ExitType
}

func (c *CaptionStyleWordAnimationModel) SetEnterType(enterType CaptionStyleWordAnimationModelEnterType) {
	c.EnterType = enterType
}

func (c *CaptionStyleWordAnimationModel) SetExitType(exitType CaptionStyleWordAnimationModelExitType) {
	c.ExitType = exitType
}

func (c *CaptionStyleWordAnimationModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleWordAnimationModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleWordAnimationModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleWordAnimationModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleWordAnimationModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleWordAnimationModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleWordAnimationModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleWordAnimationModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
