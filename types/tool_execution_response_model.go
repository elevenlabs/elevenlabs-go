package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolExecutionResponseModel struct {
	// The ID of the tool that was executed
	ToolId string `json:"tool_id" xml:"tool_id" required:"true"`
	// The request/call ID associated with this tool execution
	ToolRequestId string `json:"tool_request_id" xml:"tool_request_id" required:"true"`
	// The ID of the conversation where the tool was executed
	ConversationId string `json:"conversation_id" xml:"conversation_id" required:"true"`
	// The ID of the agent that ran the tool
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The branch ID if the agent has branches
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// Unix timestamp when the tool was executed
	Timestamp float64 `json:"timestamp" xml:"timestamp" required:"true"`
	// How long the tool execution took
	LatencySecs float64 `json:"latency_secs" xml:"latency_secs" required:"true"`
	// Whether the tool execution failed
	IsError *bool `json:"is_error,omitempty" xml:"is_error,omitempty"`
	// LLM-extracted parameters sent to the tool (JSON string)
	RequestPayload *string `json:"request_payload,omitempty" xml:"request_payload,omitempty"`
	// Response returned by the tool
	ResponsePayload *string `json:"response_payload,omitempty" xml:"response_payload,omitempty"`
	// Error message if the tool execution failed
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// Error category (internal, customer_config, customer_auth, external_server, external_client, client_timeout, unknown)
	ErrorType       *string                                    `json:"error_type,omitempty" xml:"error_type,omitempty"`
	Id              string                                     `json:"id" xml:"id" required:"true"`
	ToolCallDetails *ToolExecutionResponseModelToolCallDetails `json:"tool_call_details,omitempty" xml:"tool_call_details,omitempty"`
	extraProperties map[string]interface{}
}

func (t *ToolExecutionResponseModel) GetToolId() (value string) {
	if t == nil {
		return
	}
	return t.ToolId
}

func (t *ToolExecutionResponseModel) GetToolRequestId() (value string) {
	if t == nil {
		return
	}
	return t.ToolRequestId
}

func (t *ToolExecutionResponseModel) GetConversationId() (value string) {
	if t == nil {
		return
	}
	return t.ConversationId
}

func (t *ToolExecutionResponseModel) GetAgentId() (value string) {
	if t == nil {
		return
	}
	return t.AgentId
}

func (t *ToolExecutionResponseModel) GetBranchId() (value *string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *ToolExecutionResponseModel) GetTimestamp() (value float64) {
	if t == nil {
		return
	}
	return t.Timestamp
}

func (t *ToolExecutionResponseModel) GetLatencySecs() (value float64) {
	if t == nil {
		return
	}
	return t.LatencySecs
}

func (t *ToolExecutionResponseModel) GetIsError() (value *bool) {
	if t == nil {
		return
	}
	return t.IsError
}

func (t *ToolExecutionResponseModel) GetRequestPayload() (value *string) {
	if t == nil {
		return
	}
	return t.RequestPayload
}

func (t *ToolExecutionResponseModel) GetResponsePayload() (value *string) {
	if t == nil {
		return
	}
	return t.ResponsePayload
}

func (t *ToolExecutionResponseModel) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *ToolExecutionResponseModel) GetErrorType() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorType
}

func (t *ToolExecutionResponseModel) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *ToolExecutionResponseModel) GetToolCallDetails() (value *ToolExecutionResponseModelToolCallDetails) {
	if t == nil {
		return
	}
	return t.ToolCallDetails
}

func (t *ToolExecutionResponseModel) SetToolId(toolId string) {
	t.ToolId = toolId
}

func (t *ToolExecutionResponseModel) SetToolRequestId(toolRequestId string) {
	t.ToolRequestId = toolRequestId
}

func (t *ToolExecutionResponseModel) SetConversationId(conversationId string) {
	t.ConversationId = conversationId
}

func (t *ToolExecutionResponseModel) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *ToolExecutionResponseModel) SetBranchId(branchId *string) {
	t.BranchId = branchId
}

func (t *ToolExecutionResponseModel) SetTimestamp(timestamp float64) {
	t.Timestamp = timestamp
}

func (t *ToolExecutionResponseModel) SetLatencySecs(latencySecs float64) {
	t.LatencySecs = latencySecs
}

func (t *ToolExecutionResponseModel) SetIsError(isError *bool) {
	t.IsError = isError
}

func (t *ToolExecutionResponseModel) SetRequestPayload(requestPayload *string) {
	t.RequestPayload = requestPayload
}

func (t *ToolExecutionResponseModel) SetResponsePayload(responsePayload *string) {
	t.ResponsePayload = responsePayload
}

func (t *ToolExecutionResponseModel) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *ToolExecutionResponseModel) SetErrorType(errorType *string) {
	t.ErrorType = errorType
}

func (t *ToolExecutionResponseModel) SetId(id string) {
	t.Id = id
}

func (t *ToolExecutionResponseModel) SetToolCallDetails(toolCallDetails *ToolExecutionResponseModelToolCallDetails) {
	t.ToolCallDetails = toolCallDetails
}

func (t *ToolExecutionResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolExecutionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolExecutionResponseModel to string"
	}
	return string(jsonData)
}

func (t *ToolExecutionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolExecutionResponseModel) MarshalJSON() ([]byte, error) {
	type embed ToolExecutionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
