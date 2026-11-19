package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DiscountResponseModel struct {
	// The discount applied to the invoice. E.g. [20.0f] for 20% off.
	DiscountPercentOff *float64 `json:"discount_percent_off,omitempty" xml:"discount_percent_off,omitempty"`
	// The discount applied to the invoice. E.g. [20.0f] for 20 cents off.
	DiscountAmountOff *float64 `json:"discount_amount_off,omitempty" xml:"discount_amount_off,omitempty"`
	extraProperties   map[string]interface{}
}

func (d *DiscountResponseModel) GetDiscountPercentOff() (value *float64) {
	if d == nil {
		return
	}
	return d.DiscountPercentOff
}

func (d *DiscountResponseModel) GetDiscountAmountOff() (value *float64) {
	if d == nil {
		return
	}
	return d.DiscountAmountOff
}

func (d *DiscountResponseModel) SetDiscountPercentOff(discountPercentOff *float64) {
	d.DiscountPercentOff = discountPercentOff
}

func (d *DiscountResponseModel) SetDiscountAmountOff(discountAmountOff *float64) {
	d.DiscountAmountOff = discountAmountOff
}

func (d *DiscountResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DiscountResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiscountResponseModel to string"
	}
	return string(jsonData)
}

func (d *DiscountResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DiscountResponseModel) MarshalJSON() ([]byte, error) {
	type embed DiscountResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
