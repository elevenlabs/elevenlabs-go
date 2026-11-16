package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmUsageCalculatorResponseModel struct {
	LlmPrices       []*LlmUsageCalculatorLlmResponseModel `json:"llm_prices" xml:"llm_prices" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LlmUsageCalculatorResponseModel) GetLlmPrices() (value []*LlmUsageCalculatorLlmResponseModel) {
	if l == nil {
		return
	}
	return l.LlmPrices
}

func (l *LlmUsageCalculatorResponseModel) SetLlmPrices(llmPrices []*LlmUsageCalculatorLlmResponseModel) {
	l.LlmPrices = llmPrices
}

func (l *LlmUsageCalculatorResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmUsageCalculatorResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageCalculatorResponseModel to string"
	}
	return string(jsonData)
}

func (l *LlmUsageCalculatorResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmUsageCalculatorResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmUsageCalculatorResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmUsageCalculatorResponseModel) MarshalJSON() ([]byte, error) {
	type embed LlmUsageCalculatorResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
