package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpServerConfigInput struct {
	// Defines the MCP server-level approval policy for tool execution.
	ApprovalPolicy *McpApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	// List of tool approval hashes for per-tool approval when approval_policy is REQUIRE_APPROVAL_PER_TOOL
	ToolApprovalHashes []*McpToolApprovalHash `json:"tool_approval_hashes,omitempty" xml:"tool_approval_hashes,omitempty"`
	// Supported MCP server transport types.
	Transport *McpServerTransport `json:"transport,omitempty" xml:"transport,omitempty"`
	// The URL of the MCP server, if this contains a secret please store as a workspace secret, otherwise store as a plain string. Must use https
	Url *McpServerConfigOutputUrl `json:"url" xml:"url" required:"true"`
	// The secret token (Authorization header) stored as a workspace secret or in-place secret
	SecretToken *McpServerConfigInputSecretToken `json:"secret_token,omitempty" xml:"secret_token,omitempty"`
	// The headers included in the request
	RequestHeaders map[string]*McpServerConfigInputRequestHeadersValue `json:"request_headers,omitempty"`
	// Entries sent in the MCP `_meta` field of tools/call requests. Values may be JSON scalars, or references to a workspace secret, dynamic variable, or environment variable resolved per call.
	RequestMeta map[string]*McpServerConfigInputRequestMetaValue `json:"request_meta,omitempty"`
	// Optional auth connection to use for authentication with this MCP server
	AuthConnection *McpServerConfigInputAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
	Name           string                              `json:"name" xml:"name" required:"true"`
	Description    *string                             `json:"description,omitempty" xml:"description,omitempty"`
	// DEPRECATED: use `pre_tool_speech` instead. If true, all tools from this MCP server will require pre-tool execution speech.
	ForcePreToolSpeech *bool              `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	PreToolSpeech      *PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If true, the user will not be able to interrupt the agent while any tool from this MCP server is running.
	DisableInterruptions *bool                 `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	InterruptionMode     *ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// Predefined tool call sound type to play during tool execution for all tools from this MCP server
	ToolCallSound *ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// Determines how the tool call sound should be played.
	ToolCallSoundBehavior *ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	ExecutionMode         *ToolExecutionMode     `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// The maximum time in seconds to wait for each MCP tool call to complete. Must be between 5 and 300 seconds (inclusive).
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// List of per-tool configuration overrides that override the server-level defaults for specific tools
	ToolConfigOverrides []*McpToolConfigOverrideInput `json:"tool_config_overrides,omitempty" xml:"tool_config_overrides,omitempty"`
	// Whether to disable HTTP compression for this MCP server. Enable this if the server does not support compressed responses.
	DisableCompression *bool `json:"disable_compression,omitempty" xml:"disable_compression,omitempty"`
	extraProperties    map[string]interface{}
}

func (m *McpServerConfigInput) GetApprovalPolicy() (value *McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpServerConfigInput) GetToolApprovalHashes() (value []*McpToolApprovalHash) {
	if m == nil {
		return
	}
	return m.ToolApprovalHashes
}

func (m *McpServerConfigInput) GetTransport() (value *McpServerTransport) {
	if m == nil {
		return
	}
	return m.Transport
}

func (m *McpServerConfigInput) GetUrl() (value *McpServerConfigOutputUrl) {
	if m == nil {
		return
	}
	return m.Url
}

func (m *McpServerConfigInput) GetSecretToken() (value *McpServerConfigInputSecretToken) {
	if m == nil {
		return
	}
	return m.SecretToken
}

func (m *McpServerConfigInput) GetRequestHeaders() (value map[string]*McpServerConfigInputRequestHeadersValue) {
	if m == nil {
		return
	}
	return m.RequestHeaders
}

func (m *McpServerConfigInput) GetRequestMeta() (value map[string]*McpServerConfigInputRequestMetaValue) {
	if m == nil {
		return
	}
	return m.RequestMeta
}

func (m *McpServerConfigInput) GetAuthConnection() (value *McpServerConfigInputAuthConnection) {
	if m == nil {
		return
	}
	return m.AuthConnection
}

func (m *McpServerConfigInput) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *McpServerConfigInput) GetDescription() (value *string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *McpServerConfigInput) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpServerConfigInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpServerConfigInput) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpServerConfigInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpServerConfigInput) GetToolCallSound() (value *ToolCallSoundType) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpServerConfigInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpServerConfigInput) GetExecutionMode() (value *ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpServerConfigInput) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpServerConfigInput) GetToolConfigOverrides() (value []*McpToolConfigOverrideInput) {
	if m == nil {
		return
	}
	return m.ToolConfigOverrides
}

func (m *McpServerConfigInput) GetDisableCompression() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableCompression
}

func (m *McpServerConfigInput) SetApprovalPolicy(approvalPolicy *McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpServerConfigInput) SetToolApprovalHashes(toolApprovalHashes []*McpToolApprovalHash) {
	m.ToolApprovalHashes = toolApprovalHashes
}

func (m *McpServerConfigInput) SetTransport(transport *McpServerTransport) {
	m.Transport = transport
}

func (m *McpServerConfigInput) SetUrl(url *McpServerConfigOutputUrl) {
	m.Url = url
}

func (m *McpServerConfigInput) SetSecretToken(secretToken *McpServerConfigInputSecretToken) {
	m.SecretToken = secretToken
}

func (m *McpServerConfigInput) SetRequestHeaders(requestHeaders map[string]*McpServerConfigInputRequestHeadersValue) {
	m.RequestHeaders = requestHeaders
}

func (m *McpServerConfigInput) SetRequestMeta(requestMeta map[string]*McpServerConfigInputRequestMetaValue) {
	m.RequestMeta = requestMeta
}

func (m *McpServerConfigInput) SetAuthConnection(authConnection *McpServerConfigInputAuthConnection) {
	m.AuthConnection = authConnection
}

func (m *McpServerConfigInput) SetName(name string) {
	m.Name = name
}

func (m *McpServerConfigInput) SetDescription(description *string) {
	m.Description = description
}

func (m *McpServerConfigInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpServerConfigInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpServerConfigInput) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpServerConfigInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpServerConfigInput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	m.ToolCallSound = toolCallSound
}

func (m *McpServerConfigInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpServerConfigInput) SetExecutionMode(executionMode *ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpServerConfigInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpServerConfigInput) SetToolConfigOverrides(toolConfigOverrides []*McpToolConfigOverrideInput) {
	m.ToolConfigOverrides = toolConfigOverrides
}

func (m *McpServerConfigInput) SetDisableCompression(disableCompression *bool) {
	m.DisableCompression = disableCompression
}

func (m *McpServerConfigInput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpServerConfigInput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerConfigInput to string"
	}
	return string(jsonData)
}

func (m *McpServerConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpServerConfigInput) MarshalJSON() ([]byte, error) {
	type embed McpServerConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
