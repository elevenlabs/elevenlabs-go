package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpServerConfigOutput struct {
	// Defines the MCP server-level approval policy for tool execution.
	ApprovalPolicy *McpApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	// List of tool approval hashes for per-tool approval when approval_policy is REQUIRE_APPROVAL_PER_TOOL
	ToolApprovalHashes []*McpToolApprovalHash `json:"tool_approval_hashes,omitempty" xml:"tool_approval_hashes,omitempty"`
	// Supported MCP server transport types.
	Transport *McpServerTransport `json:"transport,omitempty" xml:"transport,omitempty"`
	// The URL of the MCP server, if this contains a secret please store as a workspace secret, otherwise store as a plain string. Must use https
	Url *McpServerConfigOutputUrl `json:"url" xml:"url" required:"true"`
	// The secret token (Authorization header) stored as a workspace secret or in-place secret
	SecretToken *McpServerConfigOutputSecretToken `json:"secret_token,omitempty" xml:"secret_token,omitempty"`
	// The headers included in the request
	RequestHeaders map[string]*McpServerConfigOutputRequestHeadersValue `json:"request_headers,omitempty"`
	// Entries sent in the MCP `_meta` field of tools/call requests. Values may be JSON scalars, or references to a workspace secret, dynamic variable, or environment variable resolved per call.
	RequestMeta map[string]*McpServerConfigOutputRequestMetaValue `json:"request_meta,omitempty"`
	// Optional auth connection to use for authentication with this MCP server
	AuthConnection *McpServerConfigOutputAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
	Name           string                               `json:"name" xml:"name" required:"true"`
	Description    *string                              `json:"description,omitempty" xml:"description,omitempty"`
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
	ToolConfigOverrides []*McpToolConfigOverrideOutput `json:"tool_config_overrides,omitempty" xml:"tool_config_overrides,omitempty"`
	// Whether to disable HTTP compression for this MCP server. Enable this if the server does not support compressed responses.
	DisableCompression *bool `json:"disable_compression,omitempty" xml:"disable_compression,omitempty"`
	extraProperties    map[string]interface{}
}

func (m *McpServerConfigOutput) GetApprovalPolicy() (value *McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpServerConfigOutput) GetToolApprovalHashes() (value []*McpToolApprovalHash) {
	if m == nil {
		return
	}
	return m.ToolApprovalHashes
}

func (m *McpServerConfigOutput) GetTransport() (value *McpServerTransport) {
	if m == nil {
		return
	}
	return m.Transport
}

func (m *McpServerConfigOutput) GetUrl() (value *McpServerConfigOutputUrl) {
	if m == nil {
		return
	}
	return m.Url
}

func (m *McpServerConfigOutput) GetSecretToken() (value *McpServerConfigOutputSecretToken) {
	if m == nil {
		return
	}
	return m.SecretToken
}

func (m *McpServerConfigOutput) GetRequestHeaders() (value map[string]*McpServerConfigOutputRequestHeadersValue) {
	if m == nil {
		return
	}
	return m.RequestHeaders
}

func (m *McpServerConfigOutput) GetRequestMeta() (value map[string]*McpServerConfigOutputRequestMetaValue) {
	if m == nil {
		return
	}
	return m.RequestMeta
}

func (m *McpServerConfigOutput) GetAuthConnection() (value *McpServerConfigOutputAuthConnection) {
	if m == nil {
		return
	}
	return m.AuthConnection
}

func (m *McpServerConfigOutput) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *McpServerConfigOutput) GetDescription() (value *string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *McpServerConfigOutput) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpServerConfigOutput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpServerConfigOutput) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpServerConfigOutput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpServerConfigOutput) GetToolCallSound() (value *ToolCallSoundType) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpServerConfigOutput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpServerConfigOutput) GetExecutionMode() (value *ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpServerConfigOutput) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpServerConfigOutput) GetToolConfigOverrides() (value []*McpToolConfigOverrideOutput) {
	if m == nil {
		return
	}
	return m.ToolConfigOverrides
}

func (m *McpServerConfigOutput) GetDisableCompression() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableCompression
}

func (m *McpServerConfigOutput) SetApprovalPolicy(approvalPolicy *McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpServerConfigOutput) SetToolApprovalHashes(toolApprovalHashes []*McpToolApprovalHash) {
	m.ToolApprovalHashes = toolApprovalHashes
}

func (m *McpServerConfigOutput) SetTransport(transport *McpServerTransport) {
	m.Transport = transport
}

func (m *McpServerConfigOutput) SetUrl(url *McpServerConfigOutputUrl) {
	m.Url = url
}

func (m *McpServerConfigOutput) SetSecretToken(secretToken *McpServerConfigOutputSecretToken) {
	m.SecretToken = secretToken
}

func (m *McpServerConfigOutput) SetRequestHeaders(requestHeaders map[string]*McpServerConfigOutputRequestHeadersValue) {
	m.RequestHeaders = requestHeaders
}

func (m *McpServerConfigOutput) SetRequestMeta(requestMeta map[string]*McpServerConfigOutputRequestMetaValue) {
	m.RequestMeta = requestMeta
}

func (m *McpServerConfigOutput) SetAuthConnection(authConnection *McpServerConfigOutputAuthConnection) {
	m.AuthConnection = authConnection
}

func (m *McpServerConfigOutput) SetName(name string) {
	m.Name = name
}

func (m *McpServerConfigOutput) SetDescription(description *string) {
	m.Description = description
}

func (m *McpServerConfigOutput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpServerConfigOutput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpServerConfigOutput) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpServerConfigOutput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpServerConfigOutput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	m.ToolCallSound = toolCallSound
}

func (m *McpServerConfigOutput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpServerConfigOutput) SetExecutionMode(executionMode *ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpServerConfigOutput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpServerConfigOutput) SetToolConfigOverrides(toolConfigOverrides []*McpToolConfigOverrideOutput) {
	m.ToolConfigOverrides = toolConfigOverrides
}

func (m *McpServerConfigOutput) SetDisableCompression(disableCompression *bool) {
	m.DisableCompression = disableCompression
}

func (m *McpServerConfigOutput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpServerConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerConfigOutput to string"
	}
	return string(jsonData)
}

func (m *McpServerConfigOutput) UnmarshalJSON(data []byte) error {
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

func (m *McpServerConfigOutput) MarshalJSON() ([]byte, error) {
	type embed McpServerConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
