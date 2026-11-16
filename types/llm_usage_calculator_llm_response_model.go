package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmUsageCalculatorLlmResponseModel struct {
	Llm             Llm     `json:"llm" xml:"llm" required:"true"`
	PricePerMinute  float64 `json:"price_per_minute" xml:"price_per_minute" required:"true"`
	PricePerMessage float64 `json:"price_per_message" xml:"price_per_message" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LlmUsageCalculatorLlmResponseModel) GetLlm() (value Llm) {
	if l == nil {
		return
	}
	return l.Llm
}

func (l *LlmUsageCalculatorLlmResponseModel) GetPricePerMinute() (value float64) {
	if l == nil {
		return
	}
	return l.PricePerMinute
}

func (l *LlmUsageCalculatorLlmResponseModel) GetPricePerMessage() (value float64) {
	if l == nil {
		return
	}
	return l.PricePerMessage
}

func (l *LlmUsageCalculatorLlmResponseModel) SetLlm(llm Llm) {
	l.Llm = llm
}

func (l *LlmUsageCalculatorLlmResponseModel) SetPricePerMinute(pricePerMinute float64) {
	l.PricePerMinute = pricePerMinute
}

func (l *LlmUsageCalculatorLlmResponseModel) SetPricePerMessage(pricePerMessage float64) {
	l.PricePerMessage = pricePerMessage
}

func (l *LlmUsageCalculatorLlmResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmUsageCalculatorLlmResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageCalculatorLlmResponseModel to string"
	}
	return string(jsonData)
}

func (l *LlmUsageCalculatorLlmResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmUsageCalculatorLlmResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmUsageCalculatorLlmResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmUsageCalculatorLlmResponseModel) MarshalJSON() ([]byte, error) {
	type embed LlmUsageCalculatorLlmResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
