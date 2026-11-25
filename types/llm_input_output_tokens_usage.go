package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmInputOutputTokensUsage struct {
	Input                       *LlmTokensCategoryUsage `json:"input,omitempty" xml:"input,omitempty"`
	InputCacheRead              *LlmTokensCategoryUsage `json:"input_cache_read,omitempty" xml:"input_cache_read,omitempty"`
	InputCacheWrite             *LlmTokensCategoryUsage `json:"input_cache_write,omitempty" xml:"input_cache_write,omitempty"`
	OutputTotal                 *LlmTokensCategoryUsage `json:"output_total,omitempty" xml:"output_total,omitempty"`
	OutputTotalReasoningPortion *LlmTokensCategoryUsage `json:"output_total_reasoning_portion,omitempty" xml:"output_total_reasoning_portion,omitempty"`
	extraProperties             map[string]interface{}
}

func (l *LlmInputOutputTokensUsage) GetInput() (value *LlmTokensCategoryUsage) {
	if l == nil {
		return
	}
	return l.Input
}

func (l *LlmInputOutputTokensUsage) GetInputCacheRead() (value *LlmTokensCategoryUsage) {
	if l == nil {
		return
	}
	return l.InputCacheRead
}

func (l *LlmInputOutputTokensUsage) GetInputCacheWrite() (value *LlmTokensCategoryUsage) {
	if l == nil {
		return
	}
	return l.InputCacheWrite
}

func (l *LlmInputOutputTokensUsage) GetOutputTotal() (value *LlmTokensCategoryUsage) {
	if l == nil {
		return
	}
	return l.OutputTotal
}

func (l *LlmInputOutputTokensUsage) GetOutputTotalReasoningPortion() (value *LlmTokensCategoryUsage) {
	if l == nil {
		return
	}
	return l.OutputTotalReasoningPortion
}

func (l *LlmInputOutputTokensUsage) SetInput(input *LlmTokensCategoryUsage) {
	l.Input = input
}

func (l *LlmInputOutputTokensUsage) SetInputCacheRead(inputCacheRead *LlmTokensCategoryUsage) {
	l.InputCacheRead = inputCacheRead
}

func (l *LlmInputOutputTokensUsage) SetInputCacheWrite(inputCacheWrite *LlmTokensCategoryUsage) {
	l.InputCacheWrite = inputCacheWrite
}

func (l *LlmInputOutputTokensUsage) SetOutputTotal(outputTotal *LlmTokensCategoryUsage) {
	l.OutputTotal = outputTotal
}

func (l *LlmInputOutputTokensUsage) SetOutputTotalReasoningPortion(outputTotalReasoningPortion *LlmTokensCategoryUsage) {
	l.OutputTotalReasoningPortion = outputTotalReasoningPortion
}

func (l *LlmInputOutputTokensUsage) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmInputOutputTokensUsage) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmInputOutputTokensUsage to string"
	}
	return string(jsonData)
}

func (l *LlmInputOutputTokensUsage) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmInputOutputTokensUsage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmInputOutputTokensUsage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmInputOutputTokensUsage) MarshalJSON() ([]byte, error) {
	type embed LlmInputOutputTokensUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
