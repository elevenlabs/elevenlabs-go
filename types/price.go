package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Currency/amount pair.
type Price struct {
	Amount          string   `json:"amount" xml:"amount" required:"true" pattern:"^(?!^[-+.]*$)[+-]?0*\\d*\\.?\\d*$"`
	Currency        Currency `json:"currency" xml:"currency" required:"true"`
	extraProperties map[string]interface{}
}

func (p *Price) GetAmount() (value string) {
	if p == nil {
		return
	}
	return p.Amount
}

func (p *Price) GetCurrency() (value Currency) {
	if p == nil {
		return
	}
	return p.Currency
}

func (p *Price) SetAmount(amount string) {
	p.Amount = amount
}

func (p *Price) SetCurrency(currency Currency) {
	p.Currency = currency
}

func (p *Price) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p Price) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Price to string"
	}
	return string(jsonData)
}

func (p *Price) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias Price
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = Price(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *Price) MarshalJSON() ([]byte, error) {
	type embed Price
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
