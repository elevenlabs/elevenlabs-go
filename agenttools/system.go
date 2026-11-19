package agenttools

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// SystemTool is a built-in tool the agent runs itself, such as ending or transferring the call.
// Create one with a System* constructor. Its name defaults to its system_tool_type, and its
// description is left out so the API uses its own; both can be set with Options.
type SystemTool struct {
	core
	systemToolType string
	// params holds the params fields other than system_tool_type.
	params map[string]any
	// malformedParams is whether the loaded params were missing, not an object, or without a
	// system_tool_type. They are then kept in config as loaded, and Build reports them.
	malformedParams bool
}

func newSystemTool(systemToolType string, opts any) *SystemTool {
	params := map[string]any{}
	if opts != nil {
		if wire, err := wirejson.Normalize(opts); err == nil {
			if m, ok := wire.(map[string]any); ok {
				params = m
			}
		}
	}
	return &SystemTool{
		core:           core{config: map[string]any{"name": systemToolType}},
		systemToolType: systemToolType,
		params:         params,
	}
}

func loadSystem(o *object) *SystemTool {
	value, _ := o.get("params")
	params, isObject := value.(*object)
	var systemToolType string
	if isObject {
		systemToolType, _ = params.vals["system_tool_type"].(string)
	}
	if systemToolType == "" {
		// Kept as loaded; the generated params union can't take it.
		return &SystemTool{core: core{config: configOf(o)}, params: map[string]any{}, malformedParams: true}
	}
	rest := plain(params).(map[string]any)
	delete(rest, "system_tool_type")
	return &SystemTool{core: core{config: configOf(o, "params")}, systemToolType: systemToolType, params: rest}
}

// SystemEndCall lets the agent end the call.
func SystemEndCall() *SystemTool { return newSystemTool("end_call", nil) }

// SystemSkipTurn lets the agent stay quiet when the user asks for a moment. opts may be nil.
func SystemSkipTurn(opts *SkipTurnOptions) *SystemTool { return newSystemTool("skip_turn", opts) }

// SystemLanguageDetection lets the agent switch language. opts may be nil.
func SystemLanguageDetection(opts *LanguageDetectionOptions) *SystemTool {
	return newSystemTool("language_detection", opts)
}

// SystemPlayKeypadTouchTone lets the agent play keypad tones, for phone menus. opts may be nil.
func SystemPlayKeypadTouchTone(opts *PlayKeypadTouchToneOptions) *SystemTool {
	return newSystemTool("play_keypad_touch_tone", opts)
}

// SystemVoicemailDetection lets the agent detect voicemail and leave a message or hang up. opts
// may be nil.
func SystemVoicemailDetection(opts *VoicemailDetectionOptions) *SystemTool {
	return newSystemTool("voicemail_detection", opts)
}

// SystemKnowledgeBase lets the agent search its knowledge base. opts may be nil.
func SystemKnowledgeBase(opts *KnowledgeBaseOptions) *SystemTool {
	return newSystemTool("knowledge_base", opts)
}

// SystemTransferToNumber lets the agent transfer the call to a phone number or SIP URI. Add
// destinations with Transfer. opts may be nil.
func SystemTransferToNumber(opts *TransferToNumberOptions) *SystemTool {
	t := newSystemTool("transfer_to_number", opts)
	t.params["transfers"] = []any{}
	return t
}

// SystemTransferToAgent lets the agent hand the conversation to another agent. Add destinations
// with Transfer.
func SystemTransferToAgent() *SystemTool {
	t := newSystemTool("transfer_to_agent", nil)
	t.params["transfers"] = []any{}
	return t
}

func (t *SystemTool) impl() toolImpl { return t }

// Name is the tool's name, as the LLM sees it.
func (t *SystemTool) Name() string { return t.name() }

// Kind is ToolKindSystem.
func (t *SystemTool) Kind() ToolKind { return ToolKindSystem }

// SystemToolType is the API's system_tool_type, such as "end_call".
func (t *SystemTool) SystemToolType() string { return t.systemToolType }

// AsWebhook is nil: the tool isn't a webhook tool.
func (t *SystemTool) AsWebhook() *WebhookTool { return nil }

// AsClient is nil: the tool isn't a client tool.
func (t *SystemTool) AsClient() *ClientTool { return nil }

// AsSystem is the tool itself.
func (t *SystemTool) AsSystem() *SystemTool { return t }

// Build validates the tool and returns the request model tools.create and tools.update accept.
// opts may be nil.
func (t *SystemTool) Build(opts *BuildOptions) (*types.ToolRequestModel, error) {
	return build(t, opts)
}

// BuildConfig validates the tool and returns only its config. Its System field is the agent's
// prompt.built_in_tools entry for this tool. opts may be nil.
func (t *SystemTool) BuildConfig(opts *BuildOptions) (*types.ToolRequestModelToolConfig, error) {
	return buildConfig(t, opts)
}

// ToJSON validates like Build and returns the tool config in API wire format.
func (t *SystemTool) ToJSON() (map[string]any, error) { return toJSON(t) }

// ToJSONString validates like Build and returns the canonical file format.
func (t *SystemTool) ToJSONString() (string, error) { return toJSONString(t) }

// Options merges options for system tools: only the options every tool kind has (see
// ToolOptions). An option that doesn't apply is reported by Build as invalid_option.
func (t *SystemTool) Options(opts *ToolOptions) *SystemTool {
	for _, entry := range acceptedOptions(opts.entries(), commonOptions, "system", &t.issues) {
		t.setOption(entry.key, entry.value)
	}
	return t
}

// Transfer adds a destination to a transfer tool. condition is plain text telling the LLM when to
// transfer there. SystemTransferToNumber takes phone and SIP destinations (options TransferType,
// CustomSipHeaders, SipReferPlayDialtone, Uui, PostDialDigits); a SIP URI needs TransferType
// "sip_refer", since the API defaults to "conference". SystemTransferToAgent takes agent
// destinations (options DelayMs, TransferMessage, EnableTransferredAgentFirstMessage,
// PreserveClientTtsOverrides, PreserveVoiceSettings). opts may be nil.
func (t *SystemTool) Transfer(condition string, destination Destination, opts *TransferOptions) *SystemTool {
	toNumber := t.systemToolType == "transfer_to_number"
	if !toNumber && t.systemToolType != "transfer_to_agent" {
		t.issues = append(t.issues, errorIssue(InvalidOption, "transfer", fmt.Sprintf("Transfer only applies to transfer tools, not %s", t.systemToolType)))
		return t
	}
	if destination.Kind() == "" || destination.IsAgent() == toNumber {
		message := fmt.Sprintf("SystemTransferToAgent needs an agent destination (%q)", condition)
		if toNumber {
			message = fmt.Sprintf("SystemTransferToNumber needs a phone or SIP destination, not an agent (%q)", condition)
		}
		t.issues = append(t.issues, errorIssue(InvalidDestination, "transfer", message))
		return t
	}
	allowed, kind := agentTransferOptions, "transfer_to_agent"
	if toNumber {
		allowed, kind = phoneTransferOptions, "transfer_to_number"
	}
	settings := acceptedOptions(opts.entries(), allowed, kind, &t.issues)
	// A SIP URI held in a dynamic variable resolves to a SIP URI too, so it also needs SIP REFER.
	sipURI := destination.Kind() == agentworkflows.DestinationKindSIP || destination.Kind() == agentworkflows.DestinationKindSIPVariable
	sipRefer := opts != nil && opts.TransferType != nil && string(*opts.TransferType) == "sip_refer"
	if sipURI && !sipRefer {
		t.issues = append(t.issues, errorIssue(InvalidDestination, "transfer", fmt.Sprintf("A SIP URI destination needs TransferType \"sip_refer\" (%q)", condition)))
		return t
	}
	entry := destination.Wire()
	delete(entry, "type")
	for _, setting := range settings {
		entry[setting.key] = setting.value
	}
	entry["condition"] = condition
	transfers, _ := t.params["transfers"].([]any)
	t.params["transfers"] = append(append([]any(nil), transfers...), entry)
	return t
}

func (t *SystemTool) validate() []Issue {
	if t.systemToolType != "transfer_to_number" && t.systemToolType != "transfer_to_agent" {
		return nil
	}
	if transfers, _ := t.params["transfers"].([]any); len(transfers) == 0 {
		return []Issue{errorIssue(NoTransfers, "transfer", fmt.Sprintf("%s needs at least one Transfer", t.systemToolType))}
	}
	return nil
}

func (t *SystemTool) modelIssues() []Issue {
	if !t.malformedParams {
		return nil
	}
	return []Issue{errorIssue(InvalidModel, "params", "The loaded params aren't an object with a system_tool_type, so the tool doesn't match the generated tool config model. Fix it in the source JSON")}
}

func (t *SystemTool) toWire() map[string]any {
	wire := cloneMap(t.config)
	wire["type"] = "system"
	if !t.malformedParams {
		params := cloneMap(t.params)
		params["system_tool_type"] = t.systemToolType
		wire["params"] = params
	}
	return wire
}
