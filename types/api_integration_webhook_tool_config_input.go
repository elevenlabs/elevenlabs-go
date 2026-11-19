package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ApiIntegrationWebhookToolConfigInput struct {
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
	ToolErrorHandlingMode *ToolErrorHandlingMode  `json:"tool_error_handling_mode,omitempty" xml:"tool_error_handling_mode,omitempty"`
	DynamicVariables      *DynamicVariablesConfig `json:"dynamic_variables,omitempty" xml:"dynamic_variables,omitempty"`
	ExecutionMode         *ToolExecutionMode      `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// The version of the API integration tool
	ToolVersion                *string `json:"tool_version,omitempty" xml:"tool_version,omitempty"`
	ApiIntegrationId           string  `json:"api_integration_id" xml:"api_integration_id" required:"true"`
	ApiIntegrationConnectionId string  `json:"api_integration_connection_id" xml:"api_integration_connection_id" required:"true"`
	// User overrides applied on top of the base api_schema
	ApiSchemaOverrides *ApiIntegrationWebhookOverrides `json:"api_schema_overrides,omitempty" xml:"api_schema_overrides,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *ApiIntegrationWebhookToolConfigInput) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *ApiIntegrationWebhookToolConfigInput) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *ApiIntegrationWebhookToolConfigInput) GetResponseTimeoutSecs() (value *int) {
	if a == nil {
		return
	}
	return a.ResponseTimeoutSecs
}

func (a *ApiIntegrationWebhookToolConfigInput) GetDisableInterruptions() (value *bool) {
	if a == nil {
		return
	}
	return a.DisableInterruptions
}

func (a *ApiIntegrationWebhookToolConfigInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if a == nil {
		return
	}
	return a.InterruptionMode
}

func (a *ApiIntegrationWebhookToolConfigInput) GetForcePreToolSpeech() (value *bool) {
	if a == nil {
		return
	}
	return a.ForcePreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if a == nil {
		return
	}
	return a.PreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigInput) GetAssignments() (value []*DynamicVariableAssignment) {
	if a == nil {
		return
	}
	return a.Assignments
}

func (a *ApiIntegrationWebhookToolConfigInput) GetToolCallSound() (value *ToolCallSoundType) {
	if a == nil {
		return
	}
	return a.ToolCallSound
}

func (a *ApiIntegrationWebhookToolConfigInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if a == nil {
		return
	}
	return a.ToolCallSoundBehavior
}

func (a *ApiIntegrationWebhookToolConfigInput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if a == nil {
		return
	}
	return a.ToolErrorHandlingMode
}

func (a *ApiIntegrationWebhookToolConfigInput) GetDynamicVariables() (value *DynamicVariablesConfig) {
	if a == nil {
		return
	}
	return a.DynamicVariables
}

func (a *ApiIntegrationWebhookToolConfigInput) GetExecutionMode() (value *ToolExecutionMode) {
	if a == nil {
		return
	}
	return a.ExecutionMode
}

func (a *ApiIntegrationWebhookToolConfigInput) GetToolVersion() (value *string) {
	if a == nil {
		return
	}
	return a.ToolVersion
}

func (a *ApiIntegrationWebhookToolConfigInput) GetApiIntegrationId() (value string) {
	if a == nil {
		return
	}
	return a.ApiIntegrationId
}

func (a *ApiIntegrationWebhookToolConfigInput) GetApiIntegrationConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ApiIntegrationConnectionId
}

func (a *ApiIntegrationWebhookToolConfigInput) GetApiSchemaOverrides() (value *ApiIntegrationWebhookOverrides) {
	if a == nil {
		return
	}
	return a.ApiSchemaOverrides
}

func (a *ApiIntegrationWebhookToolConfigInput) SetName(name string) {
	a.Name = name
}

func (a *ApiIntegrationWebhookToolConfigInput) SetDescription(description string) {
	a.Description = description
}

func (a *ApiIntegrationWebhookToolConfigInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	a.ResponseTimeoutSecs = responseTimeoutSecs
}

func (a *ApiIntegrationWebhookToolConfigInput) SetDisableInterruptions(disableInterruptions *bool) {
	a.DisableInterruptions = disableInterruptions
}

func (a *ApiIntegrationWebhookToolConfigInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	a.InterruptionMode = interruptionMode
}

func (a *ApiIntegrationWebhookToolConfigInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	a.ForcePreToolSpeech = forcePreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	a.PreToolSpeech = preToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigInput) SetAssignments(assignments []*DynamicVariableAssignment) {
	a.Assignments = assignments
}

func (a *ApiIntegrationWebhookToolConfigInput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	a.ToolCallSound = toolCallSound
}

func (a *ApiIntegrationWebhookToolConfigInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	a.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (a *ApiIntegrationWebhookToolConfigInput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	a.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (a *ApiIntegrationWebhookToolConfigInput) SetDynamicVariables(dynamicVariables *DynamicVariablesConfig) {
	a.DynamicVariables = dynamicVariables
}

func (a *ApiIntegrationWebhookToolConfigInput) SetExecutionMode(executionMode *ToolExecutionMode) {
	a.ExecutionMode = executionMode
}

func (a *ApiIntegrationWebhookToolConfigInput) SetToolVersion(toolVersion *string) {
	a.ToolVersion = toolVersion
}

func (a *ApiIntegrationWebhookToolConfigInput) SetApiIntegrationId(apiIntegrationId string) {
	a.ApiIntegrationId = apiIntegrationId
}

func (a *ApiIntegrationWebhookToolConfigInput) SetApiIntegrationConnectionId(apiIntegrationConnectionId string) {
	a.ApiIntegrationConnectionId = apiIntegrationConnectionId
}

func (a *ApiIntegrationWebhookToolConfigInput) SetApiSchemaOverrides(apiSchemaOverrides *ApiIntegrationWebhookOverrides) {
	a.ApiSchemaOverrides = apiSchemaOverrides
}

func (a *ApiIntegrationWebhookToolConfigInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ApiIntegrationWebhookToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApiIntegrationWebhookToolConfigInput to string"
	}
	return string(jsonData)
}

func (a *ApiIntegrationWebhookToolConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ApiIntegrationWebhookToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed ApiIntegrationWebhookToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
