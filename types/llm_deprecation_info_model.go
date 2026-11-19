package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type LlmDeprecationInfoModel struct {
	Llm Llm `json:"llm" xml:"llm" required:"true"`
	// Whether this model is currently deprecated. True if the model is immediately deprecated or within the warning period.
	IsDeprecated bool `json:"is_deprecated" xml:"is_deprecated" required:"true"`
	// Whether this model is currently in the warning period before deprecation.
	IsInWarningPeriod *bool `json:"is_in_warning_period,omitempty" xml:"is_in_warning_period,omitempty"`
	// Whether traffic is currently being progressively routed to the replacement model.
	IsInFallbackPeriod *bool `json:"is_in_fallback_period,omitempty" xml:"is_in_fallback_period,omitempty"`
	// Current percentage of traffic being routed to the replacement model (0-100).
	FallbackPercentage *int `json:"fallback_percentage,omitempty" xml:"fallback_percentage,omitempty"`
	// The date when the model provider will deprecate this model. Null for immediately deprecated models.
	ProviderDeprecationDate *time.Time `json:"provider_deprecation_date,omitempty" xml:"provider_deprecation_date,omitempty"`
	// The model that replaces this deprecated model. Traffic will be automatically routed to this model.
	ReplacementModel *Llm `json:"replacement_model,omitempty" xml:"replacement_model,omitempty"`
	// Custom deprecation timing configuration for this model. Null if using the default configuration.
	DeprecationConfig *LlmDeprecationConfigModel `json:"deprecation_config,omitempty" xml:"deprecation_config,omitempty"`
	extraProperties   map[string]interface{}
}

func (l *LlmDeprecationInfoModel) GetLlm() (value Llm) {
	if l == nil {
		return
	}
	return l.Llm
}

func (l *LlmDeprecationInfoModel) GetIsDeprecated() (value bool) {
	if l == nil {
		return
	}
	return l.IsDeprecated
}

func (l *LlmDeprecationInfoModel) GetIsInWarningPeriod() (value *bool) {
	if l == nil {
		return
	}
	return l.IsInWarningPeriod
}

func (l *LlmDeprecationInfoModel) GetIsInFallbackPeriod() (value *bool) {
	if l == nil {
		return
	}
	return l.IsInFallbackPeriod
}

func (l *LlmDeprecationInfoModel) GetFallbackPercentage() (value *int) {
	if l == nil {
		return
	}
	return l.FallbackPercentage
}

func (l *LlmDeprecationInfoModel) GetProviderDeprecationDate() (value *time.Time) {
	if l == nil {
		return
	}
	return l.ProviderDeprecationDate
}

func (l *LlmDeprecationInfoModel) GetReplacementModel() (value *Llm) {
	if l == nil {
		return
	}
	return l.ReplacementModel
}

func (l *LlmDeprecationInfoModel) GetDeprecationConfig() (value *LlmDeprecationConfigModel) {
	if l == nil {
		return
	}
	return l.DeprecationConfig
}

func (l *LlmDeprecationInfoModel) SetLlm(llm Llm) {
	l.Llm = llm
}

func (l *LlmDeprecationInfoModel) SetIsDeprecated(isDeprecated bool) {
	l.IsDeprecated = isDeprecated
}

func (l *LlmDeprecationInfoModel) SetIsInWarningPeriod(isInWarningPeriod *bool) {
	l.IsInWarningPeriod = isInWarningPeriod
}

func (l *LlmDeprecationInfoModel) SetIsInFallbackPeriod(isInFallbackPeriod *bool) {
	l.IsInFallbackPeriod = isInFallbackPeriod
}

func (l *LlmDeprecationInfoModel) SetFallbackPercentage(fallbackPercentage *int) {
	l.FallbackPercentage = fallbackPercentage
}

func (l *LlmDeprecationInfoModel) SetProviderDeprecationDate(providerDeprecationDate *time.Time) {
	l.ProviderDeprecationDate = providerDeprecationDate
}

func (l *LlmDeprecationInfoModel) SetReplacementModel(replacementModel *Llm) {
	l.ReplacementModel = replacementModel
}

func (l *LlmDeprecationInfoModel) SetDeprecationConfig(deprecationConfig *LlmDeprecationConfigModel) {
	l.DeprecationConfig = deprecationConfig
}

func (l *LlmDeprecationInfoModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmDeprecationInfoModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmDeprecationInfoModel to string"
	}
	return string(jsonData)
}

func (l *LlmDeprecationInfoModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmDeprecationInfoModel) MarshalJSON() ([]byte, error) {
	type embed LlmDeprecationInfoModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
