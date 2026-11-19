package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type QuoteInfo struct {
	// The quoted price for this item in USD. Use the order's total_amount_usd for the combined order total.
	AmountUsd       float64 `json:"amount_usd" xml:"amount_usd" required:"true"`
	extraProperties map[string]interface{}
}

func (q *QuoteInfo) GetAmountUsd() (value float64) {
	if q == nil {
		return
	}
	return q.AmountUsd
}

func (q *QuoteInfo) SetAmountUsd(amountUsd float64) {
	q.AmountUsd = amountUsd
}

func (q *QuoteInfo) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QuoteInfo) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QuoteInfo to string"
	}
	return string(jsonData)
}

func (q *QuoteInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, q); err != nil {
		return err
	}
	type alias QuoteInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*q = QuoteInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *q)
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QuoteInfo) MarshalJSON() ([]byte, error) {
	type embed QuoteInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*q),
	}
	return json.Marshal(marshaler)
}
