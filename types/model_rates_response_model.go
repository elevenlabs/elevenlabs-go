package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ModelRatesResponseModel struct {
	// The cost multiplier for characters.
	CharacterCostMultiplier float64 `json:"character_cost_multiplier" xml:"character_cost_multiplier" required:"true"`
	// Discount multiplier applied to cost estimates. Defaults to 1.0 (no discount).
	CostDiscountMultiplier *float64 `json:"cost_discount_multiplier,omitempty" xml:"cost_discount_multiplier,omitempty"`
	extraProperties        map[string]interface{}
}

func (m *ModelRatesResponseModel) GetCharacterCostMultiplier() (value float64) {
	if m == nil {
		return
	}
	return m.CharacterCostMultiplier
}

func (m *ModelRatesResponseModel) GetCostDiscountMultiplier() (value *float64) {
	if m == nil {
		return
	}
	return m.CostDiscountMultiplier
}

func (m *ModelRatesResponseModel) SetCharacterCostMultiplier(characterCostMultiplier float64) {
	m.CharacterCostMultiplier = characterCostMultiplier
}

func (m *ModelRatesResponseModel) SetCostDiscountMultiplier(costDiscountMultiplier *float64) {
	m.CostDiscountMultiplier = costDiscountMultiplier
}

func (m *ModelRatesResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ModelRatesResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ModelRatesResponseModel to string"
	}
	return string(jsonData)
}

func (m *ModelRatesResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias ModelRatesResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = ModelRatesResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ModelRatesResponseModel) MarshalJSON() ([]byte, error) {
	type embed ModelRatesResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
