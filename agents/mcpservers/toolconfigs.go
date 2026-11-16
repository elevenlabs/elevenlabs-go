package mcpservers

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type McpToolConfigOverrideCreateRequestModel struct {
	// Environment whose values are used when the MCP server URL, headers, or auth connection reference environment variables. Mirrors the environment a conversation would run in; defaults to production.
	Environment *string `json:"-" url:"environment,omitempty" explode:"true" serializationStyle:"form" queryParam:"environment"`
	// ID of the MCP Server.
	McpServerId string `json:"-" url:"-"`
	// DEPRECATED: use `pre_tool_speech` instead. If set, overrides the server's force_pre_tool_speech setting for this tool.
	ForcePreToolSpeech *bool `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	// If set, overrides the server's pre_tool_speech setting for this tool.
	PreToolSpeech *types.PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If set, overrides the server's disable_interruptions setting for this tool.
	DisableInterruptions *bool `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	// If set, overrides the server's interruption_mode setting for this tool.
	InterruptionMode *types.ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
	ToolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// If set, overrides the server's tool_call_sound_behavior setting for this tool
	ToolCallSoundBehavior *types.ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// If set, overrides the server's execution_mode setting for this tool
	ExecutionMode *types.ToolExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// If set, overrides the server's response timeout for this MCP tool.
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// Dynamic variable assignments for this MCP tool
	Assignments []*types.DynamicVariableAssignment `json:"assignments,omitempty" xml:"assignments,omitempty"`
	// Mapping of json path to input override configuration
	InputOverrides map[string]*McpToolConfigOverrideCreateRequestModelInputOverridesValue `json:"input_overrides,omitempty"`
	// Mock responses with optional parameter conditions. Evaluated top-to-bottom; first match wins.
	ResponseMocks []*types.ToolResponseMockConfigInput `json:"response_mocks,omitempty" xml:"response_mocks,omitempty"`
	// The name of the MCP tool
	ToolName string `json:"tool_name" xml:"tool_name" required:"true"`
}

func (m *McpToolConfigOverrideCreateRequestModel) GetEnvironment() (value *string) {
	if m == nil {
		return
	}
	return m.Environment
}

func (m *McpToolConfigOverrideCreateRequestModel) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpToolConfigOverrideCreateRequestModel) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpToolConfigOverrideCreateRequestModel) GetPreToolSpeech() (value *types.PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpToolConfigOverrideCreateRequestModel) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpToolConfigOverrideCreateRequestModel) GetInterruptionMode() (value *types.ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpToolConfigOverrideCreateRequestModel) GetToolCallSound() (value *McpToolConfigOverrideCreateRequestModelToolCallSound) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpToolConfigOverrideCreateRequestModel) GetToolCallSoundBehavior() (value *types.ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpToolConfigOverrideCreateRequestModel) GetExecutionMode() (value *types.ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpToolConfigOverrideCreateRequestModel) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpToolConfigOverrideCreateRequestModel) GetAssignments() (value []*types.DynamicVariableAssignment) {
	if m == nil {
		return
	}
	return m.Assignments
}

func (m *McpToolConfigOverrideCreateRequestModel) GetInputOverrides() (value map[string]*McpToolConfigOverrideCreateRequestModelInputOverridesValue) {
	if m == nil {
		return
	}
	return m.InputOverrides
}

func (m *McpToolConfigOverrideCreateRequestModel) GetResponseMocks() (value []*types.ToolResponseMockConfigInput) {
	if m == nil {
		return
	}
	return m.ResponseMocks
}

func (m *McpToolConfigOverrideCreateRequestModel) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolConfigOverrideCreateRequestModel) SetEnvironment(environment *string) {
	m.Environment = environment
}

func (m *McpToolConfigOverrideCreateRequestModel) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpToolConfigOverrideCreateRequestModel) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpToolConfigOverrideCreateRequestModel) SetPreToolSpeech(preToolSpeech *types.PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpToolConfigOverrideCreateRequestModel) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpToolConfigOverrideCreateRequestModel) SetInterruptionMode(interruptionMode *types.ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpToolConfigOverrideCreateRequestModel) SetToolCallSound(toolCallSound *McpToolConfigOverrideCreateRequestModelToolCallSound) {
	m.ToolCallSound = toolCallSound
}

func (m *McpToolConfigOverrideCreateRequestModel) SetToolCallSoundBehavior(toolCallSoundBehavior *types.ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpToolConfigOverrideCreateRequestModel) SetExecutionMode(executionMode *types.ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpToolConfigOverrideCreateRequestModel) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpToolConfigOverrideCreateRequestModel) SetAssignments(assignments []*types.DynamicVariableAssignment) {
	m.Assignments = assignments
}

func (m *McpToolConfigOverrideCreateRequestModel) SetInputOverrides(inputOverrides map[string]*McpToolConfigOverrideCreateRequestModelInputOverridesValue) {
	m.InputOverrides = inputOverrides
}

func (m *McpToolConfigOverrideCreateRequestModel) SetResponseMocks(responseMocks []*types.ToolResponseMockConfigInput) {
	m.ResponseMocks = responseMocks
}

func (m *McpToolConfigOverrideCreateRequestModel) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolConfigOverrideCreateRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m McpToolConfigOverrideCreateRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolConfigOverrideCreateRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpToolConfigOverrideCreateRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}

func (m *McpToolConfigOverrideCreateRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpToolConfigOverrideCreateRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

type McpToolConfigOverrideCreateRequestModelInputOverridesValue struct {
	Source          string
	Constant        *types.ConstantSchemaOverride
	DynamicVariable *types.DynamicVariableSchemaOverride
	Llm             *types.LlmSchemaOverride
	Omit            *types.OmitSchemaOverride
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) GetConstant() *types.ConstantSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Constant == nil {
		return nil
	}
	return m.Constant
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) GetDynamicVariable() *types.DynamicVariableSchemaOverride {
	if m == nil {
		return nil
	}
	if m.DynamicVariable == nil {
		return nil
	}
	return m.DynamicVariable
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) GetLlm() *types.LlmSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Llm == nil {
		return nil
	}
	return m.Llm
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) GetOmit() *types.OmitSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Omit == nil {
		return nil
	}
	return m.Omit
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) GetSource() string {
	if m == nil {
		return ""
	}
	return m.Source
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", m)
	}
	switch unmarshaler.Source {
	case "constant":
		value := new(types.ConstantSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Constant = value
	case "dynamic_variable":
		value := new(types.DynamicVariableSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.DynamicVariable = value
	case "llm":
		value := new(types.LlmSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Llm = value
	case "omit":
		value := new(types.OmitSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Omit = value
	}
	return nil
}

func (m McpToolConfigOverrideCreateRequestModelInputOverridesValue) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Constant != nil {
		return internal.MarshalJSONWithDiscriminant(m.Constant, "source", "constant")
	}
	if m.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(m.DynamicVariable, "source", "dynamic_variable")
	}
	if m.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(m.Llm, "source", "llm")
	}
	if m.Omit != nil {
		return internal.MarshalJSONWithDiscriminant(m.Omit, "source", "omit")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type McpToolConfigOverrideCreateRequestModelInputOverridesValueVisitor interface {
	VisitConstant(*types.ConstantSchemaOverride) error
	VisitDynamicVariable(*types.DynamicVariableSchemaOverride) error
	VisitLlm(*types.LlmSchemaOverride) error
	VisitOmit(*types.OmitSchemaOverride) error
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) Accept(visitor McpToolConfigOverrideCreateRequestModelInputOverridesValueVisitor) error {
	if m.Constant != nil {
		return visitor.VisitConstant(m.Constant)
	}
	if m.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(m.DynamicVariable)
	}
	if m.Llm != nil {
		return visitor.VisitLlm(m.Llm)
	}
	if m.Omit != nil {
		return visitor.VisitOmit(m.Omit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *McpToolConfigOverrideCreateRequestModelInputOverridesValue) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Constant != nil {
		fields = append(fields, "constant")
	}
	if m.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if m.Llm != nil {
		fields = append(fields, "llm")
	}
	if m.Omit != nil {
		fields = append(fields, "omit")
	}
	if len(fields) == 0 {
		if m.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Source)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Source != "" {
		field := fields[0]
		if m.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Source,
				m,
			)
		}
	}
	return nil
}

type McpToolConfigOverrideUpdateRequestModel struct {
	// Environment whose values are used when the MCP server URL, headers, or auth connection reference environment variables. Mirrors the environment a conversation would run in; defaults to production.
	Environment *string `json:"-" url:"environment,omitempty" explode:"true" serializationStyle:"form" queryParam:"environment"`
	// Name of the MCP tool to update config overrides for.
	ToolName string `json:"-" url:"-"`
	// ID of the MCP Server.
	McpServerId string `json:"-" url:"-"`
	// DEPRECATED: use `pre_tool_speech` instead. If set, overrides the server's force_pre_tool_speech setting for this tool.
	ForcePreToolSpeech *bool `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	// If set, overrides the server's pre_tool_speech setting for this tool.
	PreToolSpeech *types.PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If set, overrides the server's disable_interruptions setting for this tool.
	DisableInterruptions *bool `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	// If set, overrides the server's interruption_mode setting for this tool.
	InterruptionMode *types.ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
	ToolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// If set, overrides the server's tool_call_sound_behavior setting for this tool
	ToolCallSoundBehavior *types.ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// If set, overrides the server's execution_mode setting for this tool
	ExecutionMode *types.ToolExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// If set, overrides the server's response timeout for this MCP tool.
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// Dynamic variable assignments for this MCP tool
	Assignments []*types.DynamicVariableAssignment `json:"assignments,omitempty" xml:"assignments,omitempty"`
	// Mapping of json path to input override configuration
	InputOverrides map[string]*McpToolConfigOverrideUpdateRequestModelInputOverridesValue `json:"input_overrides,omitempty"`
	// Mock responses with optional parameter conditions. Evaluated top-to-bottom; first match wins.
	ResponseMocks []*types.ToolResponseMockConfigInput `json:"response_mocks,omitempty" xml:"response_mocks,omitempty"`
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetEnvironment() (value *string) {
	if m == nil {
		return
	}
	return m.Environment
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetPreToolSpeech() (value *types.PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetInterruptionMode() (value *types.ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetToolCallSound() (value *McpToolConfigOverrideUpdateRequestModelToolCallSound) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetToolCallSoundBehavior() (value *types.ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetExecutionMode() (value *types.ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetAssignments() (value []*types.DynamicVariableAssignment) {
	if m == nil {
		return
	}
	return m.Assignments
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetInputOverrides() (value map[string]*McpToolConfigOverrideUpdateRequestModelInputOverridesValue) {
	if m == nil {
		return
	}
	return m.InputOverrides
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetResponseMocks() (value []*types.ToolResponseMockConfigInput) {
	if m == nil {
		return
	}
	return m.ResponseMocks
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetEnvironment(environment *string) {
	m.Environment = environment
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetPreToolSpeech(preToolSpeech *types.PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetInterruptionMode(interruptionMode *types.ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetToolCallSound(toolCallSound *McpToolConfigOverrideUpdateRequestModelToolCallSound) {
	m.ToolCallSound = toolCallSound
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetToolCallSoundBehavior(toolCallSoundBehavior *types.ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetExecutionMode(executionMode *types.ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetAssignments(assignments []*types.DynamicVariableAssignment) {
	m.Assignments = assignments
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetInputOverrides(inputOverrides map[string]*McpToolConfigOverrideUpdateRequestModelInputOverridesValue) {
	m.InputOverrides = inputOverrides
}

func (m *McpToolConfigOverrideUpdateRequestModel) SetResponseMocks(responseMocks []*types.ToolResponseMockConfigInput) {
	m.ResponseMocks = responseMocks
}

func (m *McpToolConfigOverrideUpdateRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m McpToolConfigOverrideUpdateRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolConfigOverrideUpdateRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpToolConfigOverrideUpdateRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}

func (m *McpToolConfigOverrideUpdateRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpToolConfigOverrideUpdateRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

type McpToolConfigOverrideUpdateRequestModelInputOverridesValue struct {
	Source          string
	Constant        *types.ConstantSchemaOverride
	DynamicVariable *types.DynamicVariableSchemaOverride
	Llm             *types.LlmSchemaOverride
	Omit            *types.OmitSchemaOverride
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) GetConstant() *types.ConstantSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Constant == nil {
		return nil
	}
	return m.Constant
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) GetDynamicVariable() *types.DynamicVariableSchemaOverride {
	if m == nil {
		return nil
	}
	if m.DynamicVariable == nil {
		return nil
	}
	return m.DynamicVariable
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) GetLlm() *types.LlmSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Llm == nil {
		return nil
	}
	return m.Llm
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) GetOmit() *types.OmitSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Omit == nil {
		return nil
	}
	return m.Omit
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) GetSource() string {
	if m == nil {
		return ""
	}
	return m.Source
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", m)
	}
	switch unmarshaler.Source {
	case "constant":
		value := new(types.ConstantSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Constant = value
	case "dynamic_variable":
		value := new(types.DynamicVariableSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.DynamicVariable = value
	case "llm":
		value := new(types.LlmSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Llm = value
	case "omit":
		value := new(types.OmitSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Omit = value
	}
	return nil
}

func (m McpToolConfigOverrideUpdateRequestModelInputOverridesValue) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Constant != nil {
		return internal.MarshalJSONWithDiscriminant(m.Constant, "source", "constant")
	}
	if m.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(m.DynamicVariable, "source", "dynamic_variable")
	}
	if m.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(m.Llm, "source", "llm")
	}
	if m.Omit != nil {
		return internal.MarshalJSONWithDiscriminant(m.Omit, "source", "omit")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type McpToolConfigOverrideUpdateRequestModelInputOverridesValueVisitor interface {
	VisitConstant(*types.ConstantSchemaOverride) error
	VisitDynamicVariable(*types.DynamicVariableSchemaOverride) error
	VisitLlm(*types.LlmSchemaOverride) error
	VisitOmit(*types.OmitSchemaOverride) error
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) Accept(visitor McpToolConfigOverrideUpdateRequestModelInputOverridesValueVisitor) error {
	if m.Constant != nil {
		return visitor.VisitConstant(m.Constant)
	}
	if m.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(m.DynamicVariable)
	}
	if m.Llm != nil {
		return visitor.VisitLlm(m.Llm)
	}
	if m.Omit != nil {
		return visitor.VisitOmit(m.Omit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *McpToolConfigOverrideUpdateRequestModelInputOverridesValue) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Constant != nil {
		fields = append(fields, "constant")
	}
	if m.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if m.Llm != nil {
		fields = append(fields, "llm")
	}
	if m.Omit != nil {
		fields = append(fields, "omit")
	}
	if len(fields) == 0 {
		if m.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Source)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Source != "" {
		field := fields[0]
		if m.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Source,
				m,
			)
		}
	}
	return nil
}
