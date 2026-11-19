package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReceptionCustomerFacingConfigOutput struct {
	AdditionalInstructions *StrippedStr `json:"additional_instructions,omitempty" xml:"additional_instructions,omitempty"`
	// Personality traits for the receptionist, e.g. `['professional', 'concise']`.
	ToneTraits []StrippedStr `json:"tone_traits,omitempty" xml:"tone_traits,omitempty" maxItems:"20"`
	// Details the agent should always try to collect from callers (presets + custom labels, e.g. ['full name', 'phone number']).
	DetailsToCollect []StrippedStr `json:"details_to_collect,omitempty" xml:"details_to_collect,omitempty" maxItems:"30"`
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
	// When true, callers cannot interrupt the receptionist while the first message is being delivered. Use when the greeting carries a recording or legal notice callers must hear in full.
	DisableFirstMessageInterruptions *bool `json:"disable_first_message_interruptions,omitempty" xml:"disable_first_message_interruptions,omitempty"`
	// Allow the end_call system tool during conversations.
	EnableEndCallTool *bool `json:"enable_end_call_tool,omitempty" xml:"enable_end_call_tool,omitempty"`
	// Allow the skip_turn system tool during conversations.
	EnableSkipTurnTool *bool `json:"enable_skip_turn_tool,omitempty" xml:"enable_skip_turn_tool,omitempty"`
	// Allow the play_keypad_touch_tone (DTMF) system tool during conversations. When enabled, the receptionist can play touch-tone digits on a live call to navigate automated phone menus or answer carrier prompts such as a 'press 1 to accept' call-screening step.
	EnablePlayKeypadTouchToneTool *bool `json:"enable_play_keypad_touch_tone_tool,omitempty" xml:"enable_play_keypad_touch_tone_tool,omitempty"`
	// Behavioral rules this agent must follow.
	AgentRules []*AgentRule `json:"agent_rules,omitempty" xml:"agent_rules,omitempty" maxItems:"50"`
	// Phone numbers this agent can transfer calls to.
	TransferRules []*ReceptionTransferRule `json:"transfer_rules,omitempty" xml:"transfer_rules,omitempty" maxItems:"100"`
	// Show the business's own number, rather than the caller's, as the caller ID on transferred calls.
	HideCallerNumberOnTransfers *bool `json:"hide_caller_number_on_transfers,omitempty" xml:"hide_caller_number_on_transfers,omitempty"`
	// Keep recording after a call is transferred and add that audio to the call recording.
	RecordAfterTransfers *bool `json:"record_after_transfers,omitempty" xml:"record_after_transfers,omitempty"`
	// Standard operating procedures (SOPs) this agent can run on calls.
	Procedures []*ReceptionProcedure `json:"procedures,omitempty" xml:"procedures,omitempty" maxItems:"20"`
	// When enabled, the receptionist may only look up or modify appointments and client records tied to the caller's verified phone number.
	EnableSecureMode *bool `json:"enable_secure_mode,omitempty" xml:"enable_secure_mode,omitempty"`
	// When enabled, inbound calls to US numbers are screened before the receptionist answers: calls whose caller ID failed carrier (STIR/SHAKEN) verification are rejected, and unknown callers whose caller ID the carrier could not vouch for must press a key to be connected.
	EnableSpamProtection *bool `json:"enable_spam_protection,omitempty" xml:"enable_spam_protection,omitempty"`
	extraProperties      map[string]interface{}
}

func (r *ReceptionCustomerFacingConfigOutput) GetAdditionalInstructions() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.AdditionalInstructions
}

func (r *ReceptionCustomerFacingConfigOutput) GetToneTraits() (value []StrippedStr) {
	if r == nil {
		return
	}
	return r.ToneTraits
}

func (r *ReceptionCustomerFacingConfigOutput) GetDetailsToCollect() (value []StrippedStr) {
	if r == nil {
		return
	}
	return r.DetailsToCollect
}

func (r *ReceptionCustomerFacingConfigOutput) GetLlm() (value *Llm) {
	if r == nil {
		return
	}
	return r.Llm
}

func (r *ReceptionCustomerFacingConfigOutput) GetReasoningEffort() (value *LlmReasoningEffort) {
	if r == nil {
		return
	}
	return r.ReasoningEffort
}

func (r *ReceptionCustomerFacingConfigOutput) GetDisableRag() (value *bool) {
	if r == nil {
		return
	}
	return r.DisableRag
}

func (r *ReceptionCustomerFacingConfigOutput) GetBackgroundVoiceDetection() (value *bool) {
	if r == nil {
		return
	}
	return r.BackgroundVoiceDetection
}

func (r *ReceptionCustomerFacingConfigOutput) GetTurnEagerness() (value *ReceptionTurnEagerness) {
	if r == nil {
		return
	}
	return r.TurnEagerness
}

func (r *ReceptionCustomerFacingConfigOutput) GetVoiceId() (value *string) {
	if r == nil {
		return
	}
	return r.VoiceId
}

func (r *ReceptionCustomerFacingConfigOutput) GetLanguage() (value *SupportedLang) {
	if r == nil {
		return
	}
	return r.Language
}

func (r *ReceptionCustomerFacingConfigOutput) GetAdditionalLanguages() (value []SupportedLang) {
	if r == nil {
		return
	}
	return r.AdditionalLanguages
}

func (r *ReceptionCustomerFacingConfigOutput) GetFirstMessage() (value *string) {
	if r == nil {
		return
	}
	return r.FirstMessage
}

func (r *ReceptionCustomerFacingConfigOutput) GetDisableFirstMessageInterruptions() (value *bool) {
	if r == nil {
		return
	}
	return r.DisableFirstMessageInterruptions
}

func (r *ReceptionCustomerFacingConfigOutput) GetEnableEndCallTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableEndCallTool
}

func (r *ReceptionCustomerFacingConfigOutput) GetEnableSkipTurnTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSkipTurnTool
}

func (r *ReceptionCustomerFacingConfigOutput) GetEnablePlayKeypadTouchToneTool() (value *bool) {
	if r == nil {
		return
	}
	return r.EnablePlayKeypadTouchToneTool
}

func (r *ReceptionCustomerFacingConfigOutput) GetAgentRules() (value []*AgentRule) {
	if r == nil {
		return
	}
	return r.AgentRules
}

func (r *ReceptionCustomerFacingConfigOutput) GetTransferRules() (value []*ReceptionTransferRule) {
	if r == nil {
		return
	}
	return r.TransferRules
}

func (r *ReceptionCustomerFacingConfigOutput) GetHideCallerNumberOnTransfers() (value *bool) {
	if r == nil {
		return
	}
	return r.HideCallerNumberOnTransfers
}

func (r *ReceptionCustomerFacingConfigOutput) GetRecordAfterTransfers() (value *bool) {
	if r == nil {
		return
	}
	return r.RecordAfterTransfers
}

func (r *ReceptionCustomerFacingConfigOutput) GetProcedures() (value []*ReceptionProcedure) {
	if r == nil {
		return
	}
	return r.Procedures
}

func (r *ReceptionCustomerFacingConfigOutput) GetEnableSecureMode() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSecureMode
}

func (r *ReceptionCustomerFacingConfigOutput) GetEnableSpamProtection() (value *bool) {
	if r == nil {
		return
	}
	return r.EnableSpamProtection
}

func (r *ReceptionCustomerFacingConfigOutput) SetAdditionalInstructions(additionalInstructions *StrippedStr) {
	r.AdditionalInstructions = additionalInstructions
}

func (r *ReceptionCustomerFacingConfigOutput) SetToneTraits(toneTraits []StrippedStr) {
	r.ToneTraits = toneTraits
}

func (r *ReceptionCustomerFacingConfigOutput) SetDetailsToCollect(detailsToCollect []StrippedStr) {
	r.DetailsToCollect = detailsToCollect
}

func (r *ReceptionCustomerFacingConfigOutput) SetLlm(llm *Llm) {
	r.Llm = llm
}

func (r *ReceptionCustomerFacingConfigOutput) SetReasoningEffort(reasoningEffort *LlmReasoningEffort) {
	r.ReasoningEffort = reasoningEffort
}

func (r *ReceptionCustomerFacingConfigOutput) SetDisableRag(disableRag *bool) {
	r.DisableRag = disableRag
}

func (r *ReceptionCustomerFacingConfigOutput) SetBackgroundVoiceDetection(backgroundVoiceDetection *bool) {
	r.BackgroundVoiceDetection = backgroundVoiceDetection
}

func (r *ReceptionCustomerFacingConfigOutput) SetTurnEagerness(turnEagerness *ReceptionTurnEagerness) {
	r.TurnEagerness = turnEagerness
}

func (r *ReceptionCustomerFacingConfigOutput) SetVoiceId(voiceId *string) {
	r.VoiceId = voiceId
}

func (r *ReceptionCustomerFacingConfigOutput) SetLanguage(language *SupportedLang) {
	r.Language = language
}

func (r *ReceptionCustomerFacingConfigOutput) SetAdditionalLanguages(additionalLanguages []SupportedLang) {
	r.AdditionalLanguages = additionalLanguages
}

func (r *ReceptionCustomerFacingConfigOutput) SetFirstMessage(firstMessage *string) {
	r.FirstMessage = firstMessage
}

func (r *ReceptionCustomerFacingConfigOutput) SetDisableFirstMessageInterruptions(disableFirstMessageInterruptions *bool) {
	r.DisableFirstMessageInterruptions = disableFirstMessageInterruptions
}

func (r *ReceptionCustomerFacingConfigOutput) SetEnableEndCallTool(enableEndCallTool *bool) {
	r.EnableEndCallTool = enableEndCallTool
}

func (r *ReceptionCustomerFacingConfigOutput) SetEnableSkipTurnTool(enableSkipTurnTool *bool) {
	r.EnableSkipTurnTool = enableSkipTurnTool
}

func (r *ReceptionCustomerFacingConfigOutput) SetEnablePlayKeypadTouchToneTool(enablePlayKeypadTouchToneTool *bool) {
	r.EnablePlayKeypadTouchToneTool = enablePlayKeypadTouchToneTool
}

func (r *ReceptionCustomerFacingConfigOutput) SetAgentRules(agentRules []*AgentRule) {
	r.AgentRules = agentRules
}

func (r *ReceptionCustomerFacingConfigOutput) SetTransferRules(transferRules []*ReceptionTransferRule) {
	r.TransferRules = transferRules
}

func (r *ReceptionCustomerFacingConfigOutput) SetHideCallerNumberOnTransfers(hideCallerNumberOnTransfers *bool) {
	r.HideCallerNumberOnTransfers = hideCallerNumberOnTransfers
}

func (r *ReceptionCustomerFacingConfigOutput) SetRecordAfterTransfers(recordAfterTransfers *bool) {
	r.RecordAfterTransfers = recordAfterTransfers
}

func (r *ReceptionCustomerFacingConfigOutput) SetProcedures(procedures []*ReceptionProcedure) {
	r.Procedures = procedures
}

func (r *ReceptionCustomerFacingConfigOutput) SetEnableSecureMode(enableSecureMode *bool) {
	r.EnableSecureMode = enableSecureMode
}

func (r *ReceptionCustomerFacingConfigOutput) SetEnableSpamProtection(enableSpamProtection *bool) {
	r.EnableSpamProtection = enableSpamProtection
}

func (r *ReceptionCustomerFacingConfigOutput) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionCustomerFacingConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCustomerFacingConfigOutput to string"
	}
	return string(jsonData)
}

func (r *ReceptionCustomerFacingConfigOutput) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionCustomerFacingConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ReceptionCustomerFacingConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
