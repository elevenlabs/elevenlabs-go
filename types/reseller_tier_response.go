package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A tier you can sell.
type ResellerTierResponse struct {
	TierId      string `json:"tier_id" xml:"tier_id" required:"true"`
	DisplayName string `json:"display_name" xml:"display_name" required:"true"`
	// What the prices are in. The same for every tier on a ladder.
	Currency ReceptionAgentInfoResponseCurrency `json:"currency" xml:"currency" required:"true"`
	// A tier's prices and the limits shown on its tier card.
	TierEconomics   *ResellerTierEconomics `json:"tier_economics" xml:"tier_economics" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResellerTierResponse) GetTierId() (value string) {
	if r == nil {
		return
	}
	return r.TierId
}

func (r *ResellerTierResponse) GetDisplayName() (value string) {
	if r == nil {
		return
	}
	return r.DisplayName
}

func (r *ResellerTierResponse) GetCurrency() (value ReceptionAgentInfoResponseCurrency) {
	if r == nil {
		return
	}
	return r.Currency
}

func (r *ResellerTierResponse) GetTierEconomics() (value *ResellerTierEconomics) {
	if r == nil {
		return
	}
	return r.TierEconomics
}

func (r *ResellerTierResponse) SetTierId(tierId string) {
	r.TierId = tierId
}

func (r *ResellerTierResponse) SetDisplayName(displayName string) {
	r.DisplayName = displayName
}

func (r *ResellerTierResponse) SetCurrency(currency ReceptionAgentInfoResponseCurrency) {
	r.Currency = currency
}

func (r *ResellerTierResponse) SetTierEconomics(tierEconomics *ResellerTierEconomics) {
	r.TierEconomics = tierEconomics
}

func (r *ResellerTierResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerTierResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerTierResponse to string"
	}
	return string(jsonData)
}

func (r *ResellerTierResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ResellerTierResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ResellerTierResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerTierResponse) MarshalJSON() ([]byte, error) {
	type embed ResellerTierResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
