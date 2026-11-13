package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmDeprecationConfigModel struct {
	// Number of days before the provider deprecation date when warnings start being shown.
	WarningStartDays int `json:"warning_start_days" xml:"warning_start_days" required:"true"`
	// Number of days before the provider deprecation date when traffic starts being routed to the replacement model.
	FallbackStartDays int `json:"fallback_start_days" xml:"fallback_start_days" required:"true"`
	// Number of days before the provider deprecation date when all traffic is routed to the replacement model.
	FallbackCompleteDays int `json:"fallback_complete_days" xml:"fallback_complete_days" required:"true"`
	// Percentage of traffic routed to the replacement model when fallback begins.
	FallbackStartPercentage int `json:"fallback_start_percentage" xml:"fallback_start_percentage" required:"true"`
	// Percentage of traffic routed to the replacement model when fallback is complete.
	FallbackCompletePercentage int `json:"fallback_complete_percentage" xml:"fallback_complete_percentage" required:"true"`
	extraProperties            map[string]interface{}
}

func (l *LlmDeprecationConfigModel) GetWarningStartDays() (value int) {
	if l == nil {
		return
	}
	return l.WarningStartDays
}

func (l *LlmDeprecationConfigModel) GetFallbackStartDays() (value int) {
	if l == nil {
		return
	}
	return l.FallbackStartDays
}

func (l *LlmDeprecationConfigModel) GetFallbackCompleteDays() (value int) {
	if l == nil {
		return
	}
	return l.FallbackCompleteDays
}

func (l *LlmDeprecationConfigModel) GetFallbackStartPercentage() (value int) {
	if l == nil {
		return
	}
	return l.FallbackStartPercentage
}

func (l *LlmDeprecationConfigModel) GetFallbackCompletePercentage() (value int) {
	if l == nil {
		return
	}
	return l.FallbackCompletePercentage
}

func (l *LlmDeprecationConfigModel) SetWarningStartDays(warningStartDays int) {
	l.WarningStartDays = warningStartDays
}

func (l *LlmDeprecationConfigModel) SetFallbackStartDays(fallbackStartDays int) {
	l.FallbackStartDays = fallbackStartDays
}

func (l *LlmDeprecationConfigModel) SetFallbackCompleteDays(fallbackCompleteDays int) {
	l.FallbackCompleteDays = fallbackCompleteDays
}

func (l *LlmDeprecationConfigModel) SetFallbackStartPercentage(fallbackStartPercentage int) {
	l.FallbackStartPercentage = fallbackStartPercentage
}

func (l *LlmDeprecationConfigModel) SetFallbackCompletePercentage(fallbackCompletePercentage int) {
	l.FallbackCompletePercentage = fallbackCompletePercentage
}

func (l *LlmDeprecationConfigModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmDeprecationConfigModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmDeprecationConfigModel to string"
	}
	return string(jsonData)
}

func (l *LlmDeprecationConfigModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmDeprecationConfigModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmDeprecationConfigModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmDeprecationConfigModel) MarshalJSON() ([]byte, error) {
	type embed LlmDeprecationConfigModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
