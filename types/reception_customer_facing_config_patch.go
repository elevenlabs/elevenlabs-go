package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Rules, procedures and transfer rules are changed through their own endpoints, not here.
type ReceptionCustomerFacingConfigPatch struct {
	AdditionalInstructions *StrippedStr `json:"additional_instructions,omitempty" xml:"additional_instructions,omitempty"`
	// Personality traits for the receptionist, e.g. `['professional', 'concise']`.
	ToneTraits []StrippedStr `json:"tone_traits,omitempty" xml:"tone_traits,omitempty"`
	// Details the agent should always try to collect from callers (presets + custom labels, e.g. ['full name', 'phone number']).
	DetailsToCollect []StrippedStr `json:"details_to_collect,omitempty" xml:"details_to_collect,omitempty"`
	// Model the receptionist runs on. Null uses the default.
	Llm *Llm `json:"llm,omitempty" xml:"llm,omitempty"`
	// Reasoning effort for the receptionist's model, up to `medium` and only among the efforts the model supports. Null uses the model's default.
	ReasoningEffort *LlmReasoningEffort `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// When true, the whole knowledge base is given to the receptionist up front instead of searched on each turn. Faster, but only for small knowledge bases; larger ones are still searched.
	DisableRag *bool `json:"disable_rag,omitempty" xml:"disable_rag,omitempty"`
	// Filter out far-field human speech before transcription.
	BackgroundVoiceDetection *bool `json:"background_voice_detection,omitempty" xml:"background_voice_detection,omitempty"`
	// How quickly the receptionist decides the caller has finished speaking.
	TurnEagerness *ReceptionTurnEagerness `json:"turn_eagerness,omitempty" xml:"turn_eagerness,omitempty"`
	// ElevenLabs voice ID
	VoiceId  *string        `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	Language *SupportedLang `json:"language,omitempty" xml:"language,omitempty"`
	// Additional language codes the agent can communicate in
	AdditionalLanguages []SupportedLang `json:"additional_languages,omitempty" xml:"additional_languages,omitempty"`
	// The first message the agent says when starting a conversation
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// How the receptionist's greeting plays.
	FirstMessageMode *FirstMessageMode `json:"first_message_mode,omitempty" xml:"first_message_mode,omitempty"`
	// Allow the end_call system tool during conversations.
	EnableEndCallTool *bool `json:"enable_end_call_tool,omitempty" xml:"enable_end_call_tool,omitempty"`
	// Allow the skip_turn system tool during conversations.
	EnableSkipTurnTool *bool `json:"enable_skip_turn_tool,omitempty" xml:"enable_skip_turn_tool,omitempty"`
	// Allow the play_keypad_touch_tone (DTMF) system tool during conversations. When enabled, the receptionist can play touch-tone digits on a live call to navigate automated phone menus or answer carrier prompts such as a 'press 1 to accept' call-screening step.
	EnablePlayKeypadTouchToneTool *bool `json:"enable_play_keypad_touch_tone_tool,omitempty" xml:"enable_play_keypad_touch_tone_tool,omitempty"`
	// Show the business's own number, rather than the caller's, as the caller ID on transferred calls.
	HideCallerNumberOnTransfers *bool `json:"hide_caller_number_on_transfers,omitempty" xml:"hide_caller_number_on_transfers,omitempty"`
	// Keep recording after a call is transferred and add that audio to the call recording.
	RecordAfterTransfers *bool `json:"record_after_transfers,omitempty" xml:"record_after_transfers,omitempty"`
	// When enabled, the receptionist may only look up or modify appointments and client records tied to the caller's verified phone number.
	EnableSecureMode *bool `json:"enable_secure_mode,omitempty" xml:"enable_secure_mode,omitempty"`
	// When enabled, inbound calls to US numbers are screened before the receptionist answers: calls whose caller ID failed carrier (STIR/SHAKEN) verification are rejected, and unknown callers whose caller ID the carrier could not vouch for must press a key to be connected.
	EnableSpamProtection *bool `json:"enable_spam_protection,omitempty" xml:"enable_spam_protection,omitempty"`
	extraProperties      map[string]interface{}
}

func (r *ReceptionCustomerFacingConfigPatch) GetAdditionalInstructions() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.AdditionalInstructions
}

func (r *ReceptionCustomerFacingConfigPatch) GetToneTraits() (value []StrippedStr) {
	if r == nil {
		return
	}
	return r.ToneTraits
}

func (r *ReceptionCustomerFacingConfigPatch) GetDetailsToCollect() (value []StrippedStr) {
	if r == nil {
		return
	}
	return r.DetailsToCollect
}

func (r *ReceptionCustomerFacingConfigPatch) GetLlm() (value *Llm) {
	if r == nil {
		return
	}
	return r.Llm
}

func (r *ReceptionCustomerFacingConfigPatch) GetReasoningEffort() (value *LlmReasoningEffort) {
	if r == nil {
		return
	}
	return r.ReasoningEffort
}

func (r *ReceptionCustomerFacingConfigPatch) GetDisableRag() (value *bool) {
	if r == nil {
		return
	}
	return r.DisableRag
}

func (r *ReceptionCustomerFacingConfigPatch) GetBackgroundVoiceDetection() (value *bool) {
	if r == nil {
		return
	}
	return r.BackgroundVoiceDetection
}

func (r *ReceptionCustomerFacingConfigPatch) GetTurnEagerness() (value *ReceptionTurnEagerness) {
	if r == nil {
		return
	}
	return r.TurnEagerness
}

func (r *ReceptionCustomerFacingConfigPatch) GetVoiceId() (value *string) {
	if r == nil {
		return
	}
	return r.VoiceId
}

func (r *ReceptionCustomerFacingConfigPatch) GetLanguage() (value *SupportedLang) {
	if r == nil {
		return
	}
	return r.Language
}

func (r *ReceptionCustomerFacingConfigPatch) GetAdditionalLanguages() (value []SupportedLang) {
	if r == nil {
		return
	}
	return r.AdditionalLanguages
}

func (r *ReceptionCustomerFacingConfigPatch) GetFirstMessage() (value *string) {
	if r == nil {
		return
	}
	return r.FirstMessage
}

func (r *ReceptionCustomerFacingConfigPatch) GetFirstMessageMode() (value *FirstMessageMode) {
	if r == nil {
		return
	}
	return r.FirstMessageMode
}

func (r *ReceptionCustomerFacingConfigPatch) GetEnableEndCallTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableEndCallTool
}

func (r *ReceptionCustomerFacingConfigPatch) GetEnableSkipTurnTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSkipTurnTool
}

func (r *ReceptionCustomerFacingConfigPatch) GetEnablePlayKeypadTouchToneTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnablePlayKeypadTouchToneTool
}

func (r *ReceptionCustomerFacingConfigPatch) GetHideCallerNumberOnTransfers() (value *bool) {
	if r == nil {
		return
	}
	return r.HideCallerNumberOnTransfers
}

func (r *ReceptionCustomerFacingConfigPatch) GetRecordAfterTransfers() (value *bool) {
	if r == nil {
		return
	}
	return r.RecordAfterTransfers
}

func (r *ReceptionCustomerFacingConfigPatch) GetEnableSecureMode() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSecureMode
}

func (r *ReceptionCustomerFacingConfigPatch) GetEnableSpamProtection() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSpamProtection
}

func (r *ReceptionCustomerFacingConfigPatch) SetAdditionalInstructions(additionalInstructions *StrippedStr) {
	r.AdditionalInstructions = additionalInstructions
}

func (r *ReceptionCustomerFacingConfigPatch) SetToneTraits(toneTraits []StrippedStr) {
	r.ToneTraits = toneTraits
}

func (r *ReceptionCustomerFacingConfigPatch) SetDetailsToCollect(detailsToCollect []StrippedStr) {
	r.DetailsToCollect = detailsToCollect
}

func (r *ReceptionCustomerFacingConfigPatch) SetLlm(llm *Llm) {
	r.Llm = llm
}

func (r *ReceptionCustomerFacingConfigPatch) SetReasoningEffort(reasoningEffort *LlmReasoningEffort) {
	r.ReasoningEffort = reasoningEffort
}

func (r *ReceptionCustomerFacingConfigPatch) SetDisableRag(disableRag *bool) {
	r.DisableRag = disableRag
}

func (r *ReceptionCustomerFacingConfigPatch) SetBackgroundVoiceDetection(backgroundVoiceDetection *bool) {
	r.BackgroundVoiceDetection = backgroundVoiceDetection
}

func (r *ReceptionCustomerFacingConfigPatch) SetTurnEagerness(turnEagerness *ReceptionTurnEagerness) {
	r.TurnEagerness = turnEagerness
}

func (r *ReceptionCustomerFacingConfigPatch) SetVoiceId(voiceId *string) {
	r.VoiceId = voiceId
}

func (r *ReceptionCustomerFacingConfigPatch) SetLanguage(language *SupportedLang) {
	r.Language = language
}

func (r *ReceptionCustomerFacingConfigPatch) SetAdditionalLanguages(additionalLanguages []SupportedLang) {
	r.AdditionalLanguages = additionalLanguages
}

func (r *ReceptionCustomerFacingConfigPatch) SetFirstMessage(firstMessage *string) {
	r.FirstMessage = firstMessage
}

func (r *ReceptionCustomerFacingConfigPatch) SetFirstMessageMode(firstMessageMode *FirstMessageMode) {
	r.FirstMessageMode = firstMessageMode
}

func (r *ReceptionCustomerFacingConfigPatch) SetEnableEndCallTool(enableEndCallTool *bool) {
	r.EnableEndCallTool = enableEndCallTool
}

func (r *ReceptionCustomerFacingConfigPatch) SetEnableSkipTurnTool(enableSkipTurnTool *bool) {
	r.EnableSkipTurnTool = enableSkipTurnTool
}

func (r *ReceptionCustomerFacingConfigPatch) SetEnablePlayKeypadTouchToneTool(enablePlayKeypadTouchToneTool *bool) {
	r.EnablePlayKeypadTouchToneTool = enablePlayKeypadTouchToneTool
}

func (r *ReceptionCustomerFacingConfigPatch) SetHideCallerNumberOnTransfers(hideCallerNumberOnTransfers *bool) {
	r.HideCallerNumberOnTransfers = hideCallerNumberOnTransfers
}

func (r *ReceptionCustomerFacingConfigPatch) SetRecordAfterTransfers(recordAfterTransfers *bool) {
	r.RecordAfterTransfers = recordAfterTransfers
}

func (r *ReceptionCustomerFacingConfigPatch) SetEnableSecureMode(enableSecureMode *bool) {
	r.EnableSecureMode = enableSecureMode
}

func (r *ReceptionCustomerFacingConfigPatch) SetEnableSpamProtection(enableSpamProtection *bool) {
	r.EnableSpamProtection = enableSpamProtection
}

func (r *ReceptionCustomerFacingConfigPatch) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionCustomerFacingConfigPatch) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCustomerFacingConfigPatch to string"
	}
	return string(jsonData)
}

func (r *ReceptionCustomerFacingConfigPatch) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionCustomerFacingConfigPatch) MarshalJSON() ([]byte, error) {
	type embed ReceptionCustomerFacingConfigPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
