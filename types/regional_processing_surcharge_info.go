package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RegionalProcessingSurchargeInfo struct {
	// The surcharge multiplier applied to this model's pricing (e.g. 1.1 for a 10% surcharge).
	Multiplier      float64 `json:"multiplier" xml:"multiplier" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RegionalProcessingSurchargeInfo) GetMultiplier() (value float64) {
	if r == nil {
		return
	}
	return r.Multiplier
}

func (r *RegionalProcessingSurchargeInfo) SetMultiplier(multiplier float64) {
	r.Multiplier = multiplier
}

func (r *RegionalProcessingSurchargeInfo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RegionalProcessingSurchargeInfo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegionalProcessingSurchargeInfo to string"
	}
	return string(jsonData)
}

func (r *RegionalProcessingSurchargeInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RegionalProcessingSurchargeInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RegionalProcessingSurchargeInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RegionalProcessingSurchargeInfo) MarshalJSON() ([]byte, error) {
	type embed RegionalProcessingSurchargeInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
