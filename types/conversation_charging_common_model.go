package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationChargingCommonModel struct {
	DevDiscount *bool   `json:"dev_discount,omitempty" xml:"dev_discount,omitempty"`
	IsBurst     *bool   `json:"is_burst,omitempty" xml:"is_burst,omitempty"`
	Tier        *string `json:"tier,omitempty" xml:"tier,omitempty"`
	// LLM usage of a conversation, split by whether the generated response reached the user.
	LlmUsage *LlmCategoryUsage `json:"llm_usage,omitempty" xml:"llm_usage,omitempty"`
	// Total LLM price in USD. Equals the sum of the prices in `llm_usage.initiated_generation`.
	LlmPrice *float64 `json:"llm_price,omitempty" xml:"llm_price,omitempty"`
	// Credits charged for LLM usage.
	LlmCharge      *int `json:"llm_charge,omitempty" xml:"llm_charge,omitempty"`
	CallCharge     *int `json:"call_charge,omitempty" xml:"call_charge,omitempty"`
	PlatformCharge *int `json:"platform_charge,omitempty" xml:"platform_charge,omitempty"`
	// Per-category breakdown of ``platform_charge`` (the analogue of ``llm_usage``).
	PlatformUsage          *PlatformUsage             `json:"platform_usage,omitempty" xml:"platform_usage,omitempty"`
	PlatformPrice          *float64                   `json:"platform_price,omitempty" xml:"platform_price,omitempty"`
	FreeMinutesConsumed    *float64                   `json:"free_minutes_consumed,omitempty" xml:"free_minutes_consumed,omitempty"`
	FreeLlmDollarsConsumed *float64                   `json:"free_llm_dollars_consumed,omitempty" xml:"free_llm_dollars_consumed,omitempty"`
	TtsUsage               *ConversationTtsUsageModel `json:"tts_usage,omitempty" xml:"tts_usage,omitempty"`
	AsrUsage               *ConversationAsrUsageModel `json:"asr_usage,omitempty" xml:"asr_usage,omitempty"`
	Analysis               *AnalysisCharging          `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// Credits charged as the voice owner fee for the custom rate voices this conversation used. It is not included in `cost`.
	VoiceSurchargeCharge *int `json:"voice_surcharge_charge,omitempty" xml:"voice_surcharge_charge,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *ConversationChargingCommonModel) GetDevDiscount() (value *bool) {
	if c == nil {
		return
	}
	return c.DevDiscount
}

func (c *ConversationChargingCommonModel) GetIsBurst() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBurst
}

func (c *ConversationChargingCommonModel) GetTier() (value *string) {
	if c == nil {
		return
	}
	return c.Tier
}

func (c *ConversationChargingCommonModel) GetLlmUsage() (value *LlmCategoryUsage) {
	if c == nil {
		return
	}
	return c.LlmUsage
}

func (c *ConversationChargingCommonModel) GetLlmPrice() (value *float64) {
	if c == nil {
		return
	}
	return c.LlmPrice
}

func (c *ConversationChargingCommonModel) GetLlmCharge() (value *int) {
	if c == nil {
		return
	}
	return c.LlmCharge
}

func (c *ConversationChargingCommonModel) GetCallCharge() (value *int) {
	if c == nil {
		return
	}
	return c.CallCharge
}

func (c *ConversationChargingCommonModel) GetPlatformCharge() (value *int) {
	if c == nil {
		return
	}
	return c.PlatformCharge
}

func (c *ConversationChargingCommonModel) GetPlatformUsage() (value *PlatformUsage) {
	if c == nil {
		return
	}
	return c.PlatformUsage
}

func (c *ConversationChargingCommonModel) GetPlatformPrice() (value *float64) {
	if c == nil {
		return
	}
	return c.PlatformPrice
}

func (c *ConversationChargingCommonModel) GetFreeMinutesConsumed() (value *float64) {
	if c == nil {
		return
	}
	return c.FreeMinutesConsumed
}

func (c *ConversationChargingCommonModel) GetFreeLlmDollarsConsumed() (value *float64) {
	if c == nil {
		return
	}
	return c.FreeLlmDollarsConsumed
}

func (c *ConversationChargingCommonModel) GetTtsUsage() (value *ConversationTtsUsageModel) {
	if c == nil {
		return
	}
	return c.TtsUsage
}

func (c *ConversationChargingCommonModel) GetAsrUsage() (value *ConversationAsrUsageModel) {
	if c == nil {
		return
	}
	return c.AsrUsage
}

func (c *ConversationChargingCommonModel) GetAnalysis() (value *AnalysisCharging) {
	if c == nil {
		return
	}
	return c.Analysis
}

func (c *ConversationChargingCommonModel) GetVoiceSurchargeCharge() (value *int) {
	if c == nil {
		return
	}
	return c.VoiceSurchargeCharge
}

func (c *ConversationChargingCommonModel) SetDevDiscount(devDiscount *bool) {
	c.DevDiscount = devDiscount
}

func (c *ConversationChargingCommonModel) SetIsBurst(isBurst *bool) {
	c.IsBurst = isBurst
}

func (c *ConversationChargingCommonModel) SetTier(tier *string) {
	c.Tier = tier
}

func (c *ConversationChargingCommonModel) SetLlmUsage(llmUsage *LlmCategoryUsage) {
	c.LlmUsage = llmUsage
}

func (c *ConversationChargingCommonModel) SetLlmPrice(llmPrice *float64) {
	c.LlmPrice = llmPrice
}

func (c *ConversationChargingCommonModel) SetLlmCharge(llmCharge *int) {
	c.LlmCharge = llmCharge
}

func (c *ConversationChargingCommonModel) SetCallCharge(callCharge *int) {
	c.CallCharge = callCharge
}

func (c *ConversationChargingCommonModel) SetPlatformCharge(platformCharge *int) {
	c.PlatformCharge = platformCharge
}

func (c *ConversationChargingCommonModel) SetPlatformUsage(platformUsage *PlatformUsage) {
	c.PlatformUsage = platformUsage
}

func (c *ConversationChargingCommonModel) SetPlatformPrice(platformPrice *float64) {
	c.PlatformPrice = platformPrice
}

func (c *ConversationChargingCommonModel) SetFreeMinutesConsumed(freeMinutesConsumed *float64) {
	c.FreeMinutesConsumed = freeMinutesConsumed
}

func (c *ConversationChargingCommonModel) SetFreeLlmDollarsConsumed(freeLlmDollarsConsumed *float64) {
	c.FreeLlmDollarsConsumed = freeLlmDollarsConsumed
}

func (c *ConversationChargingCommonModel) SetTtsUsage(ttsUsage *ConversationTtsUsageModel) {
	c.TtsUsage = ttsUsage
}

func (c *ConversationChargingCommonModel) SetAsrUsage(asrUsage *ConversationAsrUsageModel) {
	c.AsrUsage = asrUsage
}

func (c *ConversationChargingCommonModel) SetAnalysis(analysis *AnalysisCharging) {
	c.Analysis = analysis
}

func (c *ConversationChargingCommonModel) SetVoiceSurchargeCharge(voiceSurchargeCharge *int) {
	c.VoiceSurchargeCharge = voiceSurchargeCharge
}

func (c *ConversationChargingCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationChargingCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationChargingCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationChargingCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationChargingCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationChargingCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
