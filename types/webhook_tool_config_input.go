package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A webhook tool is a tool that calls an external webhook from our server
type WebhookToolConfigInput struct {
	Name string `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description" xml:"description" required:"true"`
	// The maximum time in seconds to wait for the tool call to complete. Must be between 5 and 300 seconds (inclusive).
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// DEPRECATED: use `interruption_mode` instead. If true, the user will not be able to interrupt the agent while this tool is running.
	DisableInterruptions *bool                 `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	InterruptionMode     *ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// DEPRECATED: use `pre_tool_speech` instead. If true, the agent will speak before the tool call.
	ForcePreToolSpeech *bool              `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	PreToolSpeech      *PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// Configuration for extracting values from tool responses and assigning them to dynamic variables
	Assignments []*DynamicVariableAssignment `json:"assignments,omitempty" xml:"assignments,omitempty"`
	// Predefined tool call sound type to play during tool execution. If not specified, no tool call sound will be played.
	ToolCallSound *ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// Determines how the tool call sound should be played.
	ToolCallSoundBehavior *ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// Controls how tool errors are processed before being shared with the agent.
	ToolErrorHandlingMode *ToolErrorHandlingMode           `json:"tool_error_handling_mode,omitempty" xml:"tool_error_handling_mode,omitempty"`
	DynamicVariables      *DynamicVariablesConfig          `json:"dynamic_variables,omitempty" xml:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode               `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	ApiSchema             *WebhookToolApiSchemaConfigInput `json:"api_schema" xml:"api_schema" required:"true"`
	// Whether to resolve a redirect from the endpoint and return the final response. One redirect is followed, as a GET without the request body; nothing configured on this tool (headers, authentication, client certificate) is sent to the redirect target. Both the endpoint and the redirect target must use HTTPS. Not supported for API integration tools.
	FollowRedirects *bool `json:"follow_redirects,omitempty" xml:"follow_redirects,omitempty"`
	// Domains a redirect may point at, e.g. 'test.example.com'. Required when following redirects, and a target outside the list is refused.
	FollowRedirectsAllowedDomains []string `json:"follow_redirects_allowed_domains,omitempty" xml:"follow_redirects_allowed_domains,omitempty" maxItems:"20"`
	extraProperties               map[string]interface{}
}

func (w *WebhookToolConfigInput) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WebhookToolConfigInput) GetDescription() (value string) {
	if w == nil {
		return
	}
	return w.Description
}

func (w *WebhookToolConfigInput) GetResponseTimeoutSecs() (value *int) {
	if w == nil {
		return
	}
	return w.ResponseTimeoutSecs
}

func (w *WebhookToolConfigInput) GetDisableInterruptions() (value *bool) {
	if w == nil {
		return
	}
	return w.DisableInterruptions
}

func (w *WebhookToolConfigInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if w == nil {
		return
	}
	return w.InterruptionMode
}

func (w *WebhookToolConfigInput) GetForcePreToolSpeech() (value *bool) {
	if w == nil {
		return
	}
	return w.ForcePreToolSpeech
}

func (w *WebhookToolConfigInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if w == nil {
		return
	}
	return w.PreToolSpeech
}

func (w *WebhookToolConfigInput) GetAssignments() (value []*DynamicVariableAssignment) {
	if w == nil {
		return
	}
	return w.Assignments
}

func (w *WebhookToolConfigInput) GetToolCallSound() (value *ToolCallSoundType) {
	if w == nil {
		return
	}
	return w.ToolCallSound
}

func (w *WebhookToolConfigInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if w == nil {
		return
	}
	return w.ToolCallSoundBehavior
}

func (w *WebhookToolConfigInput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if w == nil {
		return
	}
	return w.ToolErrorHandlingMode
}

func (w *WebhookToolConfigInput) GetDynamicVariables() (value *DynamicVariablesConfig) {
	if w == nil {
		return
	}
	return w.DynamicVariables
}

func (w *WebhookToolConfigInput) GetExecutionMode() (value *ToolExecutionMode) {
	if w == nil {
		return
	}
	return w.ExecutionMode
}

func (w *WebhookToolConfigInput) GetApiSchema() (value *WebhookToolApiSchemaConfigInput) {
	if w == nil {
		return
	}
	return w.ApiSchema
}

func (w *WebhookToolConfigInput) GetFollowRedirects() (value *bool) {
	if w == nil {
		return
	}
	return w.FollowRedirects
}

func (w *WebhookToolConfigInput) GetFollowRedirectsAllowedDomains() (value []string) {
	if w == nil {
		return
	}
	return w.FollowRedirectsAllowedDomains
}

func (w *WebhookToolConfigInput) SetName(name string) {
	w.Name = name
}

func (w *WebhookToolConfigInput) SetDescription(description string) {
	w.Description = description
}

func (w *WebhookToolConfigInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	w.ResponseTimeoutSecs = responseTimeoutSecs
}

func (w *WebhookToolConfigInput) SetDisableInterruptions(disableInterruptions *bool) {
	w.DisableInterruptions = disableInterruptions
}

func (w *WebhookToolConfigInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	w.InterruptionMode = interruptionMode
}

func (w *WebhookToolConfigInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	w.ForcePreToolSpeech = forcePreToolSpeech
}

func (w *WebhookToolConfigInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	w.PreToolSpeech = preToolSpeech
}

func (w *WebhookToolConfigInput) SetAssignments(assignments []*DynamicVariableAssignment) {
	w.Assignments = assignments
}

func (w *WebhookToolConfigInput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	w.ToolCallSound = toolCallSound
}

func (w *WebhookToolConfigInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	w.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (w *WebhookToolConfigInput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	w.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (w *WebhookToolConfigInput) SetDynamicVariables(dynamicVariables *DynamicVariablesConfig) {
	w.DynamicVariables = dynamicVariables
}

func (w *WebhookToolConfigInput) SetExecutionMode(executionMode *ToolExecutionMode) {
	w.ExecutionMode = executionMode
}

func (w *WebhookToolConfigInput) SetApiSchema(apiSchema *WebhookToolApiSchemaConfigInput) {
	w.ApiSchema = apiSchema
}

func (w *WebhookToolConfigInput) SetFollowRedirects(followRedirects *bool) {
	w.FollowRedirects = followRedirects
}

func (w *WebhookToolConfigInput) SetFollowRedirectsAllowedDomains(followRedirectsAllowedDomains []string) {
	w.FollowRedirectsAllowedDomains = followRedirectsAllowedDomains
}

func (w *WebhookToolConfigInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookToolConfigInput to string"
	}
	return string(jsonData)
}

func (w *WebhookToolConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed WebhookToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
