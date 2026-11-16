package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Accumulated charge for a single :class:`PlatformCategory`.
type PlatformCategoryUsage struct {
	Credits         *int     `json:"credits,omitempty" xml:"credits,omitempty"`
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty"`
	Quantity        *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PlatformCategoryUsage) GetCredits() (value *int) {
	if p == nil {
		return
	}
	return p.Credits
}

func (p *PlatformCategoryUsage) GetPrice() (value *float64) {
	if p == nil {
		return
	}
	return p.Price
}

func (p *PlatformCategoryUsage) GetQuantity() (value *float64) {
	if p == nil {
		return
	}
	return p.Quantity
}

func (p *PlatformCategoryUsage) SetCredits(credits *int) {
	p.Credits = credits
}

func (p *PlatformCategoryUsage) SetPrice(price *float64) {
	p.Price = price
}

func (p *PlatformCategoryUsage) SetQuantity(quantity *float64) {
	p.Quantity = quantity
}

func (p *PlatformCategoryUsage) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlatformCategoryUsage) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlatformCategoryUsage to string"
	}
	return string(jsonData)
}

func (p *PlatformCategoryUsage) UnmarshalJSON(data []byte) error {
	type unmarshaler PlatformCategoryUsage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PlatformCategoryUsage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlatformCategoryUsage) MarshalJSON() ([]byte, error) {
	type embed PlatformCategoryUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
