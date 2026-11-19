package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Model struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id" xml:"model_id" required:"true"`
	// The name of the model.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Whether the model can be finetuned.
	CanBeFinetuned *bool `json:"can_be_finetuned,omitempty" xml:"can_be_finetuned,omitempty"`
	// Whether the model can do text-to-speech.
	CanDoTextToSpeech *bool `json:"can_do_text_to_speech,omitempty" xml:"can_do_text_to_speech,omitempty"`
	// Whether the model can do voice conversion.
	CanDoVoiceConversion *bool `json:"can_do_voice_conversion,omitempty" xml:"can_do_voice_conversion,omitempty"`
	// Whether the model can use style.
	CanUseStyle *bool `json:"can_use_style,omitempty" xml:"can_use_style,omitempty"`
	// Whether the model can use speaker boost.
	CanUseSpeakerBoost *bool `json:"can_use_speaker_boost,omitempty" xml:"can_use_speaker_boost,omitempty"`
	// Whether the model serves pro voices.
	ServesProVoices *bool `json:"serves_pro_voices,omitempty" xml:"serves_pro_voices,omitempty"`
	// The cost factor for the model.
	TokenCostFactor *float64 `json:"token_cost_factor,omitempty" xml:"token_cost_factor,omitempty"`
	// The description of the model.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether the model requires alpha access.
	RequiresAlphaAccess *bool `json:"requires_alpha_access,omitempty" xml:"requires_alpha_access,omitempty"`
	// Deprecated. Not enforced; use `maximum_text_length_per_request` instead.
	MaxCharactersRequestFreeUser *int `json:"max_characters_request_free_user,omitempty" xml:"max_characters_request_free_user,omitempty"`
	// Deprecated. Not enforced; use `maximum_text_length_per_request` instead.
	MaxCharactersRequestSubscribedUser *int `json:"max_characters_request_subscribed_user,omitempty" xml:"max_characters_request_subscribed_user,omitempty"`
	// The maximum number of input text characters accepted per request for this model. Longer requests are rejected.
	MaximumTextLengthPerRequest *int `json:"maximum_text_length_per_request,omitempty" xml:"maximum_text_length_per_request,omitempty"`
	// The languages supported by the model.
	Languages  []*LanguageResponse      `json:"languages,omitempty" xml:"languages,omitempty"`
	ModelRates *ModelRatesResponseModel `json:"model_rates,omitempty" xml:"model_rates,omitempty"`
	// The concurrency group for the model.
	ConcurrencyGroup *string `json:"concurrency_group,omitempty" xml:"concurrency_group,omitempty"`
	extraProperties  map[string]interface{}
}

func (m *Model) GetModelId() (value string) {
	if m == nil {
		return
	}
	return m.ModelId
}

func (m *Model) GetName() (value *string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *Model) GetCanBeFinetuned() (value *bool) {
	if m == nil {
		return
	}
	return m.CanBeFinetuned
}

func (m *Model) GetCanDoTextToSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.CanDoTextToSpeech
}

func (m *Model) GetCanDoVoiceConversion() (value *bool) {
	if m == nil {
		return
	}
	return m.CanDoVoiceConversion
}

func (m *Model) GetCanUseStyle() (value *bool) {
	if m == nil {
		return
	}
	return m.CanUseStyle
}

func (m *Model) GetCanUseSpeakerBoost() (value *bool) {
	if m == nil {
		return
	}
	return m.CanUseSpeakerBoost
}

func (m *Model) GetServesProVoices() (value *bool) {
	if m == nil {
		return
	}
	return m.ServesProVoices
}

func (m *Model) GetTokenCostFactor() (value *float64) {
	if m == nil {
		return
	}
	return m.TokenCostFactor
}

func (m *Model) GetDescription() (value *string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *Model) GetRequiresAlphaAccess() (value *bool) {
	if m == nil {
		return
	}
	return m.RequiresAlphaAccess
}

func (m *Model) GetMaxCharactersRequestFreeUser() (value *int) {
	if m == nil {
		return
	}
	return m.MaxCharactersRequestFreeUser
}

func (m *Model) GetMaxCharactersRequestSubscribedUser() (value *int) {
	if m == nil {
		return
	}
	return m.MaxCharactersRequestSubscribedUser
}

func (m *Model) GetMaximumTextLengthPerRequest() (value *int) {
	if m == nil {
		return
	}
	return m.MaximumTextLengthPerRequest
}

func (m *Model) GetLanguages() (value []*LanguageResponse) {
	if m == nil {
		return
	}
	return m.Languages
}

func (m *Model) GetModelRates() (value *ModelRatesResponseModel) {
	if m == nil {
		return
	}
	return m.ModelRates
}

func (m *Model) GetConcurrencyGroup() (value *string) {
	if m == nil {
		return
	}
	return m.ConcurrencyGroup
}

func (m *Model) SetModelId(modelId string) {
	m.ModelId = modelId
}

func (m *Model) SetName(name *string) {
	m.Name = name
}

func (m *Model) SetCanBeFinetuned(canBeFinetuned *bool) {
	m.CanBeFinetuned = canBeFinetuned
}

func (m *Model) SetCanDoTextToSpeech(canDoTextToSpeech *bool) {
	m.CanDoTextToSpeech = canDoTextToSpeech
}

func (m *Model) SetCanDoVoiceConversion(canDoVoiceConversion *bool) {
	m.CanDoVoiceConversion = canDoVoiceConversion
}

func (m *Model) SetCanUseStyle(canUseStyle *bool) {
	m.CanUseStyle = canUseStyle
}

func (m *Model) SetCanUseSpeakerBoost(canUseSpeakerBoost *bool) {
	m.CanUseSpeakerBoost = canUseSpeakerBoost
}

func (m *Model) SetServesProVoices(servesProVoices *bool) {
	m.ServesProVoices = servesProVoices
}

func (m *Model) SetTokenCostFactor(tokenCostFactor *float64) {
	m.TokenCostFactor = tokenCostFactor
}

func (m *Model) SetDescription(description *string) {
	m.Description = description
}

func (m *Model) SetRequiresAlphaAccess(requiresAlphaAccess *bool) {
	m.RequiresAlphaAccess = requiresAlphaAccess
}

func (m *Model) SetMaxCharactersRequestFreeUser(maxCharactersRequestFreeUser *int) {
	m.MaxCharactersRequestFreeUser = maxCharactersRequestFreeUser
}

func (m *Model) SetMaxCharactersRequestSubscribedUser(maxCharactersRequestSubscribedUser *int) {
	m.MaxCharactersRequestSubscribedUser = maxCharactersRequestSubscribedUser
}

func (m *Model) SetMaximumTextLengthPerRequest(maximumTextLengthPerRequest *int) {
	m.MaximumTextLengthPerRequest = maximumTextLengthPerRequest
}

func (m *Model) SetLanguages(languages []*LanguageResponse) {
	m.Languages = languages
}

func (m *Model) SetModelRates(modelRates *ModelRatesResponseModel) {
	m.ModelRates = modelRates
}

func (m *Model) SetConcurrencyGroup(concurrencyGroup *string) {
	m.ConcurrencyGroup = concurrencyGroup
}

func (m *Model) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m Model) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: Model to string"
	}
	return string(jsonData)
}

func (m *Model) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias Model
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = Model(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *Model) MarshalJSON() ([]byte, error) {
	type embed Model
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
