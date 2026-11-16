package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmInfoModel struct {
	Llm Llm `json:"llm" xml:"llm" required:"true"`
	// Whether this is a pinned checkpoint version of a model rather than a top-level alias.
	IsCheckpoint bool `json:"is_checkpoint" xml:"is_checkpoint" required:"true"`
	// Maximum number of output tokens the model can generate.
	MaxTokensLimit int `json:"max_tokens_limit" xml:"max_tokens_limit" required:"true"`
	// Maximum number of input context tokens the model supports.
	MaxContextLimit int `json:"max_context_limit" xml:"max_context_limit" required:"true"`
	// Whether the model supports image file inputs during conversations.
	SupportsImageInput bool `json:"supports_image_input" xml:"supports_image_input" required:"true"`
	// Whether the model supports document (PDF) file inputs during conversations.
	SupportsDocumentInput bool `json:"supports_document_input" xml:"supports_document_input" required:"true"`
	// Whether the model supports calling multiple tools in parallel.
	SupportsParallelToolCalls bool `json:"supports_parallel_tool_calls" xml:"supports_parallel_tool_calls" required:"true"`
	// Available reasoning effort levels for this model. Null if the model does not support configurable reasoning.
	AvailableReasoningEfforts []LlmReasoningEffort `json:"available_reasoning_efforts,omitempty" xml:"available_reasoning_efforts,omitempty"`
	// Available service tiers for this model. Always includes at least the standard tier.
	AvailableServiceTiers []LlmServiceTier `json:"available_service_tiers" xml:"available_service_tiers" required:"true"`
	// Deprecation information if this model is deprecated or scheduled for deprecation. Null if the model is not affected.
	DeprecationInfo *LlmDeprecationInfoModel `json:"deprecation_info,omitempty" xml:"deprecation_info,omitempty"`
	// Regional processing surcharge details if this model has additional costs in the current deployment region. Null if no surcharge applies.
	RegionalProcessingSurcharge *RegionalProcessingSurchargeInfo `json:"regional_processing_surcharge,omitempty" xml:"regional_processing_surcharge,omitempty"`
	extraProperties             map[string]interface{}
}

func (l *LlmInfoModel) GetLlm() (value Llm) {
	if l == nil {
		return
	}
	return l.Llm
}

func (l *LlmInfoModel) GetIsCheckpoint() (value bool) {
	if l == nil {
		return
	}
	return l.IsCheckpoint
}

func (l *LlmInfoModel) GetMaxTokensLimit() (value int) {
	if l == nil {
		return
	}
	return l.MaxTokensLimit
}

func (l *LlmInfoModel) GetMaxContextLimit() (value int) {
	if l == nil {
		return
	}
	return l.MaxContextLimit
}

func (l *LlmInfoModel) GetSupportsImageInput() (value bool) {
	if l == nil {
		return
	}
	return l.SupportsImageInput
}

func (l *LlmInfoModel) GetSupportsDocumentInput() (value bool) {
	if l == nil {
		return
	}
	return l.SupportsDocumentInput
}

func (l *LlmInfoModel) GetSupportsParallelToolCalls() (value bool) {
	if l == nil {
		return
	}
	return l.SupportsParallelToolCalls
}

func (l *LlmInfoModel) GetAvailableReasoningEfforts() (value []LlmReasoningEffort) {
	if l == nil {
		return
	}
	return l.AvailableReasoningEfforts
}

func (l *LlmInfoModel) GetAvailableServiceTiers() (value []LlmServiceTier) {
	if l == nil {
		return
	}
	return l.AvailableServiceTiers
}

func (l *LlmInfoModel) GetDeprecationInfo() (value *LlmDeprecationInfoModel) {
	if l == nil {
		return
	}
	return l.DeprecationInfo
}

func (l *LlmInfoModel) GetRegionalProcessingSurcharge() (value *RegionalProcessingSurchargeInfo) {
	if l == nil {
		return
	}
	return l.RegionalProcessingSurcharge
}

func (l *LlmInfoModel) SetLlm(llm Llm) {
	l.Llm = llm
}

func (l *LlmInfoModel) SetIsCheckpoint(isCheckpoint bool) {
	l.IsCheckpoint = isCheckpoint
}

func (l *LlmInfoModel) SetMaxTokensLimit(maxTokensLimit int) {
	l.MaxTokensLimit = maxTokensLimit
}

func (l *LlmInfoModel) SetMaxContextLimit(maxContextLimit int) {
	l.MaxContextLimit = maxContextLimit
}

func (l *LlmInfoModel) SetSupportsImageInput(supportsImageInput bool) {
	l.SupportsImageInput = supportsImageInput
}

func (l *LlmInfoModel) SetSupportsDocumentInput(supportsDocumentInput bool) {
	l.SupportsDocumentInput = supportsDocumentInput
}

func (l *LlmInfoModel) SetSupportsParallelToolCalls(supportsParallelToolCalls bool) {
	l.SupportsParallelToolCalls = supportsParallelToolCalls
}

func (l *LlmInfoModel) SetAvailableReasoningEfforts(availableReasoningEfforts []LlmReasoningEffort) {
	l.AvailableReasoningEfforts = availableReasoningEfforts
}

func (l *LlmInfoModel) SetAvailableServiceTiers(availableServiceTiers []LlmServiceTier) {
	l.AvailableServiceTiers = availableServiceTiers
}

func (l *LlmInfoModel) SetDeprecationInfo(deprecationInfo *LlmDeprecationInfoModel) {
	l.DeprecationInfo = deprecationInfo
}

func (l *LlmInfoModel) SetRegionalProcessingSurcharge(regionalProcessingSurcharge *RegionalProcessingSurchargeInfo) {
	l.RegionalProcessingSurcharge = regionalProcessingSurcharge
}

func (l *LlmInfoModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmInfoModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmInfoModel to string"
	}
	return string(jsonData)
}

func (l *LlmInfoModel) UnmarshalJSON(data []byte) error {
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

func (l *LlmInfoModel) MarshalJSON() ([]byte, error) {
	type embed LlmInfoModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
