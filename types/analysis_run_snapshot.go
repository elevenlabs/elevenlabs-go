package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// LLM cost of the most recent post-call analysis pass on this conversation.
type AnalysisRunSnapshot struct {
	Price            *float64           `json:"price,omitempty" xml:"price,omitempty"`
	Charge           *int               `json:"charge,omitempty" xml:"charge,omitempty"`
	PricePerFeature  map[string]float64 `json:"price_per_feature,omitempty"`
	ChargePerFeature map[string]int     `json:"charge_per_feature,omitempty"`
	extraProperties  map[string]interface{}
}

func (a *AnalysisRunSnapshot) GetPrice() (value *float64) {
	if a == nil {
		return
	}
	return a.Price
}

func (a *AnalysisRunSnapshot) GetCharge() (value *int) {
	if a == nil {
		return
	}
	return a.Charge
}

func (a *AnalysisRunSnapshot) GetPricePerFeature() (value map[string]float64) {
	if a == nil {
		return
	}
	return a.PricePerFeature
}

func (a *AnalysisRunSnapshot) GetChargePerFeature() (value map[string]int) {
	if a == nil {
		return
	}
	return a.ChargePerFeature
}

func (a *AnalysisRunSnapshot) SetPrice(price *float64) {
	a.Price = price
}

func (a *AnalysisRunSnapshot) SetCharge(charge *int) {
	a.Charge = charge
}

func (a *AnalysisRunSnapshot) SetPricePerFeature(pricePerFeature map[string]float64) {
	a.PricePerFeature = pricePerFeature
}

func (a *AnalysisRunSnapshot) SetChargePerFeature(chargePerFeature map[string]int) {
	a.ChargePerFeature = chargePerFeature
}

func (a *AnalysisRunSnapshot) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisRunSnapshot) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisRunSnapshot to string"
	}
	return string(jsonData)
}

func (a *AnalysisRunSnapshot) UnmarshalJSON(data []byte) error {
	type unmarshaler AnalysisRunSnapshot
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AnalysisRunSnapshot(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisRunSnapshot) MarshalJSON() ([]byte, error) {
	type embed AnalysisRunSnapshot
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
