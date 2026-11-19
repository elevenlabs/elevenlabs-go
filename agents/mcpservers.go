package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Request model for creating/updating an MCP Server configuration.
type McpServerRequestModel struct {
	Config *types.McpServerConfigInput `json:"config" xml:"config" required:"true"`
}

func (m *McpServerRequestModel) GetConfig() (value *types.McpServerConfigInput) {
	if m == nil {
		return
	}
	return m.Config
}

func (m *McpServerRequestModel) SetConfig(config *types.McpServerConfigInput) {
	m.Config = config
}

func (m *McpServerRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m McpServerRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpServerRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpServerRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpServerRequestModel(tmp)
	return nil
}

func (m *McpServerRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpServerRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

// Unified request model for updating MCP Server configuration.
type McpServerConfigUpdateRequestModel struct {
	// ID of the MCP Server.
	McpServerId string `json:"-" url:"-"`
	// The approval mode to set for the MCP server
	ApprovalPolicy *types.McpApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	// DEPRECATED: use `pre_tool_speech` instead. If set, overrides the server's force_pre_tool_speech setting for this tool.
	ForcePreToolSpeech *bool `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	// If set, overrides the server's pre_tool_speech setting for this tool.
	PreToolSpeech *types.PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If set, overrides the server's disable_interruptions setting for this tool.
	DisableInterruptions *bool `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	// If set, overrides the server's interruption_mode setting for this tool.
	InterruptionMode *types.ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// Predefined tool call sound type to play during tool execution for all tools from this MCP server
	ToolCallSound *types.ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// Determines when the tool call sound should play for all tools from this MCP server
	ToolCallSoundBehavior *types.ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// If set, overrides the server's execution_mode setting for this tool
	ExecutionMode *types.ToolExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// The maximum time in seconds to wait for each MCP tool call to complete.
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// The headers to include in requests to the MCP server
	RequestHeaders map[string]*McpServerConfigUpdateRequestModelRequestHeadersValue `json:"request_headers,omitempty"`
	// Entries sent in the MCP `_meta` field of tools/call requests. Values may be JSON scalars, or references to a workspace secret, dynamic variable, or environment variable resolved per call.
	RequestMeta map[string]*McpServerConfigUpdateRequestModelRequestMetaValue `json:"request_meta,omitempty"`
	// Whether to disable HTTP compression for this MCP server
	DisableCompression *bool `json:"disable_compression,omitempty" xml:"disable_compression,omitempty"`
	// Optional secret token for authentication with this MCP server
	SecretToken *types.ConvAiSecretLocator `json:"secret_token,omitempty" xml:"secret_token,omitempty"`
	// Optional auth connection to use for authentication with this MCP server
	AuthConnection *McpServerConfigUpdateRequestModelAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
}

func (m *McpServerConfigUpdateRequestModel) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpServerConfigUpdateRequestModel) GetApprovalPolicy() (value *types.McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpServerConfigUpdateRequestModel) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpServerConfigUpdateRequestModel) GetPreToolSpeech() (value *types.PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpServerConfigUpdateRequestModel) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpServerConfigUpdateRequestModel) GetInterruptionMode() (value *types.ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpServerConfigUpdateRequestModel) GetToolCallSound() (value *types.ToolCallSoundType) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpServerConfigUpdateRequestModel) GetToolCallSoundBehavior() (value *types.ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpServerConfigUpdateRequestModel) GetExecutionMode() (value *types.ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpServerConfigUpdateRequestModel) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpServerConfigUpdateRequestModel) GetRequestHeaders() (value map[string]*McpServerConfigUpdateRequestModelRequestHeadersValue) {
	if m == nil {
		return
	}
	return m.RequestHeaders
}

func (m *McpServerConfigUpdateRequestModel) GetRequestMeta() (value map[string]*McpServerConfigUpdateRequestModelRequestMetaValue) {
	if m == nil {
		return
	}
	return m.RequestMeta
}

func (m *McpServerConfigUpdateRequestModel) GetDisableCompression() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableCompression
}

func (m *McpServerConfigUpdateRequestModel) GetSecretToken() (value *types.ConvAiSecretLocator) {
	if m == nil {
		return
	}
	return m.SecretToken
}

func (m *McpServerConfigUpdateRequestModel) GetAuthConnection() (value *McpServerConfigUpdateRequestModelAuthConnection) {
	if m == nil {
		return
	}
	return m.AuthConnection
}

func (m *McpServerConfigUpdateRequestModel) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpServerConfigUpdateRequestModel) SetApprovalPolicy(approvalPolicy *types.McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpServerConfigUpdateRequestModel) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpServerConfigUpdateRequestModel) SetPreToolSpeech(preToolSpeech *types.PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpServerConfigUpdateRequestModel) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpServerConfigUpdateRequestModel) SetInterruptionMode(interruptionMode *types.ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpServerConfigUpdateRequestModel) SetToolCallSound(toolCallSound *types.ToolCallSoundType) {
	m.ToolCallSound = toolCallSound
}

func (m *McpServerConfigUpdateRequestModel) SetToolCallSoundBehavior(toolCallSoundBehavior *types.ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpServerConfigUpdateRequestModel) SetExecutionMode(executionMode *types.ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpServerConfigUpdateRequestModel) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpServerConfigUpdateRequestModel) SetRequestHeaders(requestHeaders map[string]*McpServerConfigUpdateRequestModelRequestHeadersValue) {
	m.RequestHeaders = requestHeaders
}

func (m *McpServerConfigUpdateRequestModel) SetRequestMeta(requestMeta map[string]*McpServerConfigUpdateRequestModelRequestMetaValue) {
	m.RequestMeta = requestMeta
}

func (m *McpServerConfigUpdateRequestModel) SetDisableCompression(disableCompression *bool) {
	m.DisableCompression = disableCompression
}

func (m *McpServerConfigUpdateRequestModel) SetSecretToken(secretToken *types.ConvAiSecretLocator) {
	m.SecretToken = secretToken
}

func (m *McpServerConfigUpdateRequestModel) SetAuthConnection(authConnection *McpServerConfigUpdateRequestModelAuthConnection) {
	m.AuthConnection = authConnection
}

func (m *McpServerConfigUpdateRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m McpServerConfigUpdateRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerConfigUpdateRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpServerConfigUpdateRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}

func (m *McpServerConfigUpdateRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpServerConfigUpdateRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

type McpServerConfigUpdateRequestModelRequestHeadersValue struct {
	String                string
	ConvAiSecretLocator   *types.ConvAiSecretLocator
	ConvAiDynamicVariable *types.ConvAiDynamicVariable
	ConvAiEnvVarLocator   *types.ConvAiEnvVarLocator
	typ                   string
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) GetConvAiSecretLocator() *types.ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) GetConvAiDynamicVariable() *types.ConvAiDynamicVariable {
	if m == nil {
		return nil
	}
	if m.ConvAiDynamicVariable == nil {
		return nil
	}
	return m.ConvAiDynamicVariable
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) GetConvAiEnvVarLocator() *types.ConvAiEnvVarLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiEnvVarLocator == nil {
		return nil
	}
	return m.ConvAiEnvVarLocator
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *types.ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		m.typ = "ConvAiSecretLocator"
		m.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *types.ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		m.typ = "ConvAiDynamicVariable"
		m.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	var valueConvAiEnvVarLocator *types.ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		m.typ = "ConvAiEnvVarLocator"
		m.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpServerConfigUpdateRequestModelRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return json.Marshal(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return json.Marshal(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return json.Marshal(m.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigUpdateRequestModelRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*types.ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*types.ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*types.ConvAiEnvVarLocator) error
}

func (m *McpServerConfigUpdateRequestModelRequestHeadersValue) Accept(visitor McpServerConfigUpdateRequestModelRequestHeadersValueVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(m.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigUpdateRequestModelRequestMetaValue struct {
	String                string
	Integer               int
	Double                float64
	Boolean               bool
	ConvAiSecretLocator   *types.ConvAiSecretLocator
	ConvAiDynamicVariable *types.ConvAiDynamicVariable
	ConvAiEnvVarLocator   *types.ConvAiEnvVarLocator
	typ                   string
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetInteger() int {
	if m == nil {
		return 0
	}
	return m.Integer
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetDouble() float64 {
	if m == nil {
		return 0
	}
	return m.Double
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetBoolean() bool {
	if m == nil {
		return false
	}
	return m.Boolean
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetConvAiSecretLocator() *types.ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetConvAiDynamicVariable() *types.ConvAiDynamicVariable {
	if m == nil {
		return nil
	}
	if m.ConvAiDynamicVariable == nil {
		return nil
	}
	return m.ConvAiDynamicVariable
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) GetConvAiEnvVarLocator() *types.ConvAiEnvVarLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiEnvVarLocator == nil {
		return nil
	}
	return m.ConvAiEnvVarLocator
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		m.typ = "Integer"
		m.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		m.typ = "Double"
		m.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		m.typ = "Boolean"
		m.Boolean = valueBoolean
		return nil
	}
	var valueConvAiSecretLocator *types.ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		m.typ = "ConvAiSecretLocator"
		m.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *types.ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		m.typ = "ConvAiDynamicVariable"
		m.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	var valueConvAiEnvVarLocator *types.ConvAiEnvVarLocator
	if err := json.Unmarshal(data, &valueConvAiEnvVarLocator); err == nil {
		m.typ = "ConvAiEnvVarLocator"
		m.ConvAiEnvVarLocator = valueConvAiEnvVarLocator
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpServerConfigUpdateRequestModelRequestMetaValue) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return json.Marshal(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return json.Marshal(m.Double)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return json.Marshal(m.Boolean)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return json.Marshal(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return json.Marshal(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return json.Marshal(m.ConvAiEnvVarLocator)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigUpdateRequestModelRequestMetaValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
	VisitConvAiSecretLocator(*types.ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*types.ConvAiDynamicVariable) error
	VisitConvAiEnvVarLocator(*types.ConvAiEnvVarLocator) error
}

func (m *McpServerConfigUpdateRequestModelRequestMetaValue) Accept(visitor McpServerConfigUpdateRequestModelRequestMetaValueVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return visitor.VisitInteger(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return visitor.VisitDouble(m.Double)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return visitor.VisitBoolean(m.Boolean)
	}
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiDynamicVariable" || m.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(m.ConvAiDynamicVariable)
	}
	if m.typ == "ConvAiEnvVarLocator" || m.ConvAiEnvVarLocator != nil {
		return visitor.VisitConvAiEnvVarLocator(m.ConvAiEnvVarLocator)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
