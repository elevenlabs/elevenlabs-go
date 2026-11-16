package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleHorizontalPlacementModel struct {
	Align           CaptionStyleHorizontalPlacementModelAlign `json:"align" xml:"align" required:"true"`
	TranslatePct    float64                                   `json:"translate_pct" xml:"translate_pct" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleHorizontalPlacementModel) GetAlign() (value CaptionStyleHorizontalPlacementModelAlign) {
	if c == nil {
		return
	}
	return c.Align
}

func (c *CaptionStyleHorizontalPlacementModel) GetTranslatePct() (value float64) {
	if c == nil {
		return
	}
	return c.TranslatePct
}

func (c *CaptionStyleHorizontalPlacementModel) SetAlign(align CaptionStyleHorizontalPlacementModelAlign) {
	c.Align = align
}

func (c *CaptionStyleHorizontalPlacementModel) SetTranslatePct(translatePct float64) {
	c.TranslatePct = translatePct
}

func (c *CaptionStyleHorizontalPlacementModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleHorizontalPlacementModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleHorizontalPlacementModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleHorizontalPlacementModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleHorizontalPlacementModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleHorizontalPlacementModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleHorizontalPlacementModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleHorizontalPlacementModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
