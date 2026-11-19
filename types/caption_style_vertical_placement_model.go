package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleVerticalPlacementModel struct {
	Align           CaptionStyleVerticalPlacementModelAlign `json:"align" xml:"align" required:"true"`
	TranslatePct    float64                                 `json:"translate_pct" xml:"translate_pct" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CaptionStyleVerticalPlacementModel) GetAlign() (value CaptionStyleVerticalPlacementModelAlign) {
	if c == nil {
		return
	}
	return c.Align
}

func (c *CaptionStyleVerticalPlacementModel) GetTranslatePct() (value float64) {
	if c == nil {
		return
	}
	return c.TranslatePct
}

func (c *CaptionStyleVerticalPlacementModel) SetAlign(align CaptionStyleVerticalPlacementModelAlign) {
	c.Align = align
}

func (c *CaptionStyleVerticalPlacementModel) SetTranslatePct(translatePct float64) {
	c.TranslatePct = translatePct
}

func (c *CaptionStyleVerticalPlacementModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleVerticalPlacementModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleVerticalPlacementModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleVerticalPlacementModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CaptionStyleVerticalPlacementModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CaptionStyleVerticalPlacementModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleVerticalPlacementModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleVerticalPlacementModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
