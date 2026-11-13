package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmTokensCategoryUsage struct {
	// Number of tokens in this category.
	Tokens *int `json:"tokens,omitempty" xml:"tokens,omitempty"`
	// Price of these tokens in USD.
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmTokensCategoryUsage) GetTokens() (value *int) {
	if l == nil {
		return
	}
	return l.Tokens
}

func (l *LlmTokensCategoryUsage) GetPrice() (value *float64) {
	if l == nil {
		return
	}
	return l.Price
}

func (l *LlmTokensCategoryUsage) SetTokens(tokens *int) {
	l.Tokens = tokens
}

func (l *LlmTokensCategoryUsage) SetPrice(price *float64) {
	l.Price = price
}

func (l *LlmTokensCategoryUsage) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmTokensCategoryUsage) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmTokensCategoryUsage to string"
	}
	return string(jsonData)
}

func (l *LlmTokensCategoryUsage) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmTokensCategoryUsage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmTokensCategoryUsage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmTokensCategoryUsage) MarshalJSON() ([]byte, error) {
	type embed LlmTokensCategoryUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
