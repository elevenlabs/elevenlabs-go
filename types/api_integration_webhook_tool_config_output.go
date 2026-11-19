package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ApiIntegrationWebhookToolConfigOutput struct {
	type_ string
	Name  string `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description" xml:"description" required:"true"`
	// The maximum time in seconds to wait for the tool call to complete. Must be between 5 and 300 seconds (inclusive).
	ResponseTimeoutSecs int `json:"response_timeout_secs" xml:"response_timeout_secs" required:"true" min:"5" max:"300"`
	// DEPRECATED: use `interruption_mode` instead. If true, the user will not be able to interrupt the agent while this tool is running.
	DisableInterruptions bool                 `json:"disable_interruptions" xml:"disable_interruptions" required:"true"`
	InterruptionMode     ToolInterruptionMode `json:"interruption_mode" xml:"interruption_mode" required:"true"`
	// DEPRECATED: use `pre_tool_speech` instead. If true, the agent will speak before the tool call.
	ForcePreToolSpeech bool              `json:"force_pre_tool_speech" xml:"force_pre_tool_speech" required:"true"`
	PreToolSpeech      PreToolSpeechMode `json:"pre_tool_speech" xml:"pre_tool_speech" required:"true"`
	// Configuration for extracting values from tool responses and assigning them to dynamic variables
	Assignments []*DynamicVariableAssignment `json:"assignments" xml:"assignments" required:"true"`
	// Predefined tool call sound type to play during tool execution. If not specified, no tool call sound will be played.
	ToolCallSound *ToolCallSoundType `json:"tool_call_sound" xml:"tool_call_sound" required:"true"`
	// Determines how the tool call sound should be played.
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior" xml:"tool_call_sound_behavior" required:"true"`
	// Controls how tool errors are processed before being shared with the agent.
	ToolErrorHandlingMode ToolErrorHandlingMode   `json:"tool_error_handling_mode" xml:"tool_error_handling_mode" required:"true"`
	DynamicVariables      *DynamicVariablesConfig `json:"dynamic_variables" xml:"dynamic_variables" required:"true"`
	ExecutionMode         ToolExecutionMode       `json:"execution_mode" xml:"execution_mode" required:"true"`
	// The version of the API integration tool
	ToolVersion                string `json:"tool_version" xml:"tool_version" required:"true"`
	ApiIntegrationId           string `json:"api_integration_id" xml:"api_integration_id" required:"true"`
	ApiIntegrationConnectionId string `json:"api_integration_connection_id" xml:"api_integration_connection_id" required:"true"`
	// User overrides applied on top of the base api_schema
	ApiSchemaOverrides *ApiIntegrationWebhookOverrides `json:"api_schema_overrides" xml:"api_schema_overrides" required:"true"`
	extraProperties    map[string]interface{}
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetResponseTimeoutSecs() (value int) {
	if a == nil {
		return
	}
	return a.ResponseTimeoutSecs
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetDisableInterruptions() (value bool) {
	if a == nil {
		return
	}
	return a.DisableInterruptions
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetInterruptionMode() (value ToolInterruptionMode) {
	if a == nil {
		return
	}
	return a.InterruptionMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetForcePreToolSpeech() (value bool) {
	if a == nil {
		return
	}
	return a.ForcePreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetPreToolSpeech() (value PreToolSpeechMode) {
	if a == nil {
		return
	}
	return a.PreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetAssignments() (value []*DynamicVariableAssignment) {
	if a == nil {
		return
	}
	return a.Assignments
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetToolCallSound() (value *ToolCallSoundType) {
	if a == nil {
		return
	}
	return a.ToolCallSound
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetToolCallSoundBehavior() (value ToolCallSoundBehavior) {
	if a == nil {
		return
	}
	return a.ToolCallSoundBehavior
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetToolErrorHandlingMode() (value ToolErrorHandlingMode) {
	if a == nil {
		return
	}
	return a.ToolErrorHandlingMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetDynamicVariables() (value *DynamicVariablesConfig) {
	if a == nil {
		return
	}
	return a.DynamicVariables
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetExecutionMode() (value ToolExecutionMode) {
	if a == nil {
		return
	}
	return a.ExecutionMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetToolVersion() (value string) {
	if a == nil {
		return
	}
	return a.ToolVersion
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetApiIntegrationId() (value string) {
	if a == nil {
		return
	}
	return a.ApiIntegrationId
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetApiIntegrationConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ApiIntegrationConnectionId
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetApiSchemaOverrides() (value *ApiIntegrationWebhookOverrides) {
	if a == nil {
		return
	}
	return a.ApiSchemaOverrides
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetName(name string) {
	a.Name = name
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetDescription(description string) {
	a.Description = description
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetResponseTimeoutSecs(responseTimeoutSecs int) {
	a.ResponseTimeoutSecs = responseTimeoutSecs
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetDisableInterruptions(disableInterruptions bool) {
	a.DisableInterruptions = disableInterruptions
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetInterruptionMode(interruptionMode ToolInterruptionMode) {
	a.InterruptionMode = interruptionMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetForcePreToolSpeech(forcePreToolSpeech bool) {
	a.ForcePreToolSpeech = forcePreToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetPreToolSpeech(preToolSpeech PreToolSpeechMode) {
	a.PreToolSpeech = preToolSpeech
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetAssignments(assignments []*DynamicVariableAssignment) {
	a.Assignments = assignments
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	a.ToolCallSound = toolCallSound
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetToolCallSoundBehavior(toolCallSoundBehavior ToolCallSoundBehavior) {
	a.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetToolErrorHandlingMode(toolErrorHandlingMode ToolErrorHandlingMode) {
	a.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetDynamicVariables(dynamicVariables *DynamicVariablesConfig) {
	a.DynamicVariables = dynamicVariables
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetExecutionMode(executionMode ToolExecutionMode) {
	a.ExecutionMode = executionMode
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetToolVersion(toolVersion string) {
	a.ToolVersion = toolVersion
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetApiIntegrationId(apiIntegrationId string) {
	a.ApiIntegrationId = apiIntegrationId
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetApiIntegrationConnectionId(apiIntegrationConnectionId string) {
	a.ApiIntegrationConnectionId = apiIntegrationConnectionId
}

func (a *ApiIntegrationWebhookToolConfigOutput) SetApiSchemaOverrides(apiSchemaOverrides *ApiIntegrationWebhookOverrides) {
	a.ApiSchemaOverrides = apiSchemaOverrides
}

func (a *ApiIntegrationWebhookToolConfigOutput) Type() string {
	return a.type_
}

func (a *ApiIntegrationWebhookToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ApiIntegrationWebhookToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApiIntegrationWebhookToolConfigOutput to string"
	}
	return string(jsonData)
}

func (a *ApiIntegrationWebhookToolConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "api_integration_webhook" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "api_integration_webhook", *literals.Type)
	}
	a.type_ = "api_integration_webhook"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ApiIntegrationWebhookToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ApiIntegrationWebhookToolConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "api_integration_webhook",
	}
	return json.Marshal(marshaler)
}
