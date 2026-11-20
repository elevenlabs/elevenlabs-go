package agenttools

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// ToolOptions are the options of a tool's Options method. Every kind of tool takes the same
// struct; a field that doesn't apply to the tool's kind is reported by Build as an invalid_option
// error. Nil means "not set": Options merges, and never clears a field.
//
// The json tags are the options' canonical names (the API field names), so options stored as JSON
// decode into this struct directly.
//
//	| Kind    | Fields                                                                                                                       |
//	|---------|------------------------------------------------------------------------------------------------------------------------------|
//	| all     | Name, Description, ResponseTimeoutSecs, InterruptionMode, PreToolSpeech, ToolCallSound, ToolCallSoundBehavior, ToolErrorHandlingMode |
//	| webhook | DynamicVariablePlaceholders, ExecutionMode, ContentType, ResponseFilter, FollowRedirects, FollowRedirectsAllowedDomains      |
//	| client  | DynamicVariablePlaceholders, ExecutionMode, ExpectsResponse                                                                  |
//	| system  | none beyond "all"                                                                                                            |
type ToolOptions struct {
	// Name is the name the LLM calls the tool by.
	Name *string `json:"name,omitempty"`
	// Description tells the LLM when to call the tool. A system tool leaves it out by default, so
	// the API uses its own.
	Description *string `json:"description,omitempty"`
	// ResponseTimeoutSecs is how long to wait for the tool: 5–300 for a webhook tool, 1–120 for a
	// client tool, or -1 (no limit) for a client tool with ExpectsResponse.
	ResponseTimeoutSecs   *int                         `json:"response_timeout_secs,omitempty"`
	InterruptionMode      *types.ToolInterruptionMode  `json:"interruption_mode,omitempty"`
	PreToolSpeech         *types.PreToolSpeechMode     `json:"pre_tool_speech,omitempty"`
	ToolCallSound         *types.ToolCallSoundType     `json:"tool_call_sound,omitempty"`
	ToolCallSoundBehavior *types.ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty"`
	ToolErrorHandlingMode *types.ToolErrorHandlingMode `json:"tool_error_handling_mode,omitempty"`

	// DynamicVariablePlaceholders are placeholder values for dynamic variables, written to
	// dynamic_variables.dynamic_variable_placeholders (webhook, client).
	DynamicVariablePlaceholders map[string]any           `json:"dynamic_variable_placeholders,omitempty"`
	ExecutionMode               *types.ToolExecutionMode `json:"execution_mode,omitempty"`

	// ContentType and ResponseFilter are written to api_schema (webhook).
	ContentType    *types.WebhookToolApiSchemaConfigInputContentType `json:"content_type,omitempty"`
	ResponseFilter *types.ResponseFilter                             `json:"response_filter,omitempty"`
	// FollowRedirects needs FollowRedirectsAllowedDomains (webhook).
	FollowRedirects               *bool    `json:"follow_redirects,omitempty"`
	FollowRedirectsAllowedDomains []string `json:"follow_redirects_allowed_domains,omitempty"`

	// ExpectsResponse makes the agent wait for the client's response and pass it to the LLM
	// (client).
	ExpectsResponse *bool `json:"expects_response,omitempty"`
}

var commonOptions = []string{
	"name", "description", "response_timeout_secs", "interruption_mode", "pre_tool_speech",
	"tool_call_sound", "tool_call_sound_behavior", "tool_error_handling_mode",
}

var webhookOptions = append(append([]string{}, commonOptions...),
	"dynamic_variable_placeholders", "execution_mode", "content_type", "response_filter",
	"follow_redirects", "follow_redirects_allowed_domains")

var clientOptions = append(append([]string{}, commonOptions...),
	"dynamic_variable_placeholders", "execution_mode", "expects_response")

type optionEntry struct {
	key   string
	value any
}

// entries lists the options that are set, in a fixed order, as Go values.
func (o *ToolOptions) entries() []optionEntry {
	if o == nil {
		return nil
	}
	var out []optionEntry
	add := func(key string, set bool, value any) {
		if set {
			out = append(out, optionEntry{key, value})
		}
	}
	add("name", o.Name != nil, o.Name)
	add("description", o.Description != nil, o.Description)
	add("response_timeout_secs", o.ResponseTimeoutSecs != nil, o.ResponseTimeoutSecs)
	add("interruption_mode", o.InterruptionMode != nil, o.InterruptionMode)
	add("pre_tool_speech", o.PreToolSpeech != nil, o.PreToolSpeech)
	add("tool_call_sound", o.ToolCallSound != nil, o.ToolCallSound)
	add("tool_call_sound_behavior", o.ToolCallSoundBehavior != nil, o.ToolCallSoundBehavior)
	add("tool_error_handling_mode", o.ToolErrorHandlingMode != nil, o.ToolErrorHandlingMode)
	add("dynamic_variable_placeholders", o.DynamicVariablePlaceholders != nil, o.DynamicVariablePlaceholders)
	add("execution_mode", o.ExecutionMode != nil, o.ExecutionMode)
	add("content_type", o.ContentType != nil, o.ContentType)
	add("response_filter", o.ResponseFilter != nil, o.ResponseFilter)
	add("follow_redirects", o.FollowRedirects != nil, o.FollowRedirects)
	add("follow_redirects_allowed_domains", o.FollowRedirectsAllowedDomains != nil, o.FollowRedirectsAllowedDomains)
	add("expects_response", o.ExpectsResponse != nil, o.ExpectsResponse)
	return out
}

// TransferOptions are the options of one SystemTool.Transfer entry, with the workflow builder's
// transfer option names. The first group applies to transfer_to_number, the second to
// transfer_to_agent; a field of the other group is reported by Build as invalid_option.
type TransferOptions struct {
	// TransferType is "conference" (the API default) or "sip_refer", which a SIP URI destination
	// needs.
	TransferType         *types.TransferTypeEnum                          `json:"transfer_type,omitempty"`
	CustomSipHeaders     []*types.PhoneNumberTransferCustomSipHeadersItem `json:"custom_sip_headers,omitempty"`
	SipReferPlayDialtone *bool                                            `json:"sip_refer_play_dialtone,omitempty"`
	Uui                  *types.UuiTransferConfig                         `json:"uui,omitempty"`
	PostDialDigits       *types.PhoneNumberTransferPostDialDigits         `json:"post_dial_digits,omitempty"`

	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`
	PreserveClientTtsOverrides         *bool   `json:"preserve_client_tts_overrides,omitempty"`
	PreserveVoiceSettings              *bool   `json:"preserve_voice_settings,omitempty"`
}

var phoneTransferOptions = []string{
	"transfer_type", "custom_sip_headers", "sip_refer_play_dialtone", "uui", "post_dial_digits",
}

var agentTransferOptions = []string{
	"delay_ms", "transfer_message", "enable_transferred_agent_first_message",
	"preserve_client_tts_overrides", "preserve_voice_settings",
}

func (o *TransferOptions) entries() []optionEntry {
	if o == nil {
		return nil
	}
	var out []optionEntry
	add := func(key string, set bool, value any) {
		if set {
			out = append(out, optionEntry{key, value})
		}
	}
	add("transfer_type", o.TransferType != nil, o.TransferType)
	add("custom_sip_headers", o.CustomSipHeaders != nil, o.CustomSipHeaders)
	add("sip_refer_play_dialtone", o.SipReferPlayDialtone != nil, o.SipReferPlayDialtone)
	add("uui", o.Uui != nil, o.Uui)
	add("post_dial_digits", o.PostDialDigits != nil, o.PostDialDigits)
	add("delay_ms", o.DelayMs != nil, o.DelayMs)
	add("transfer_message", o.TransferMessage != nil, o.TransferMessage)
	add("enable_transferred_agent_first_message", o.EnableTransferredAgentFirstMessage != nil, o.EnableTransferredAgentFirstMessage)
	add("preserve_client_tts_overrides", o.PreserveClientTtsOverrides != nil, o.PreserveClientTtsOverrides)
	add("preserve_voice_settings", o.PreserveVoiceSettings != nil, o.PreserveVoiceSettings)
	return out
}

// acceptedOptions converts the entries in allowed to wire values, in order, and records the rest
// (and values that can't be converted to JSON) as invalid_option issues.
func acceptedOptions(entries []optionEntry, allowed []string, kind string, issues *[]Issue) []optionEntry {
	var out []optionEntry
	for _, entry := range entries {
		if !contains(allowed, entry.key) {
			*issues = append(*issues, errorIssue(InvalidOption, entry.key, fmt.Sprintf("Option %q doesn't apply to %s tools", entry.key, kind)))
			continue
		}
		wire, err := wirejson.Normalize(entry.value)
		if err != nil {
			*issues = append(*issues, errorIssue(InvalidOption, entry.key, fmt.Sprintf("Option %q can't be converted to JSON: %v", entry.key, err)))
			continue
		}
		out = append(out, optionEntry{entry.key, wire})
	}
	return out
}

// AssignOptions are the optional settings of WebhookTool.Assign.
type AssignOptions struct {
	// Sanitize removes the value from the tool response before the LLM and the transcript see it.
	Sanitize *bool `json:"sanitize,omitempty"`
	// PreserveNativeType stores lists and objects as they are instead of as JSON strings.
	PreserveNativeType *bool `json:"preserve_native_type,omitempty"`
}

// BuildOptions are the optional settings of Build and BuildConfig.
type BuildOptions struct {
	// OnWarning is called once per warning (unknown_tool_type).
	OnWarning func(Issue)
}

// SkipTurnOptions are the params of SystemSkipTurn.
type SkipTurnOptions struct {
	// WaitTimeoutSecs is how long to wait before the agent checks in; -1 waits for the user.
	WaitTimeoutSecs *float64 `json:"wait_timeout_secs,omitempty"`
}

// LanguageDetectionOptions are the params of SystemLanguageDetection.
type LanguageDetectionOptions struct {
	// OnlyAtConversationStart only switches language during the first two user turns.
	OnlyAtConversationStart *bool `json:"only_at_conversation_start,omitempty"`
}

// PlayKeypadTouchToneOptions are the params of SystemPlayKeypadTouchTone.
type PlayKeypadTouchToneOptions struct {
	// UseOutOfBandDtmf sends tones as RTP events instead of in-band audio (SIP trunk numbers only).
	UseOutOfBandDtmf *bool `json:"use_out_of_band_dtmf,omitempty"`
	// SuppressTurnAfterDtmf keeps the agent quiet after playing tones, so speech doesn't confuse
	// an IVR.
	SuppressTurnAfterDtmf *bool `json:"suppress_turn_after_dtmf,omitempty"`
}

// VoicemailDetectionOptions are the params of SystemVoicemailDetection.
type VoicemailDetectionOptions struct {
	// VoicemailMessage is the message to leave on voicemail. Without one, the call ends.
	VoicemailMessage *string `json:"voicemail_message,omitempty"`
}

// KnowledgeBaseOptions are the params of SystemKnowledgeBase.
type KnowledgeBaseOptions struct {
	EnabledStrategies []types.SearchStrategy `json:"enabled_strategies,omitempty"`
}

// TransferToNumberOptions are the params of SystemTransferToNumber.
type TransferToNumberOptions struct {
	// EnableClientMessage plays a message to the caller while they wait for the transfer.
	EnableClientMessage *bool `json:"enable_client_message,omitempty"`
}
