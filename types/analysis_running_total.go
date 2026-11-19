package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Cumulative LLM cost of running post-call analysis on this conversation.
type AnalysisRunningTotal struct {
	Price            *float64           `json:"price,omitempty" xml:"price,omitempty"`
	Charge           *int               `json:"charge,omitempty" xml:"charge,omitempty"`
	Runs             *int               `json:"runs,omitempty" xml:"runs,omitempty"`
	PricePerFeature  map[string]float64 `json:"price_per_feature,omitempty"`
	ChargePerFeature map[string]int     `json:"charge_per_feature,omitempty"`
	extraProperties  map[string]interface{}
}

func (a *AnalysisRunningTotal) GetPrice() (value *float64) {
	if a == nil {
		return
	}
	return a.Price
}

func (a *AnalysisRunningTotal) GetCharge() (value *int) {
	if a == nil {
		return
	}
	return a.Charge
}

func (a *AnalysisRunningTotal) GetRuns() (value *int) {
	if a == nil {
		return
	}
	return a.Runs
}

func (a *AnalysisRunningTotal) GetPricePerFeature() (value map[string]float64) {
	if a == nil {
		return
	}
	return a.PricePerFeature
}

func (a *AnalysisRunningTotal) GetChargePerFeature() (value map[string]int) {
	if a == nil {
		return
	}
	return a.ChargePerFeature
}

func (a *AnalysisRunningTotal) SetPrice(price *float64) {
	a.Price = price
}

func (a *AnalysisRunningTotal) SetCharge(charge *int) {
	a.Charge = charge
}

func (a *AnalysisRunningTotal) SetRuns(runs *int) {
	a.Runs = runs
}

func (a *AnalysisRunningTotal) SetPricePerFeature(pricePerFeature map[string]float64) {
	a.PricePerFeature = pricePerFeature
}

func (a *AnalysisRunningTotal) SetChargePerFeature(chargePerFeature map[string]int) {
	a.ChargePerFeature = chargePerFeature
}

func (a *AnalysisRunningTotal) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisRunningTotal) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisRunningTotal to string"
	}
	return string(jsonData)
}

func (a *AnalysisRunningTotal) UnmarshalJSON(data []byte) error {
	type unmarshaler AnalysisRunningTotal
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AnalysisRunningTotal(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisRunningTotal) MarshalJSON() ([]byte, error) {
	type embed AnalysisRunningTotal
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
