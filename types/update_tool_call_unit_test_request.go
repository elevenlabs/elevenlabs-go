package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateToolCallUnitTestRequest struct {
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata *TestFromConversationMetadataInput `json:"from_conversation_metadata,omitempty" xml:"from_conversation_metadata,omitempty"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]any                                   `json:"dynamic_variables,omitempty"`
	ChatHistory      []*ConversationHistoryTranscriptCommonModelInput `json:"chat_history,omitempty" xml:"chat_history,omitempty" maxItems:"200"`
	// Simulate the test as if the conversation originated from this channel.
	ConversationInitiationSource *ConversationInitiationSource `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	// The environment to resolve environment-specific variable values against when running this test (URL, headers, auth connections). If not provided, defaults to 'production'. For simulation tests, simulation_environment takes precedence when set.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not evaluated.
	ToolCallParameters *UnitTestToolCallEvaluationModelInput `json:"tool_call_parameters,omitempty" xml:"tool_call_parameters,omitempty"`
	// If set to True this test will pass if any tool call returned by the LLM matches the criteria. Otherwise it will fail if more than one tool is returned by the agent.
	CheckAnyToolMatches *bool  `json:"check_any_tool_matches,omitempty" xml:"check_any_tool_matches,omitempty"`
	Name                string `json:"name" xml:"name" required:"true"`
	// The ID of the parent folder. If not provided, the test will be moved to the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpdateToolCallUnitTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if u == nil {
		return
	}
	return u.FromConversationMetadata
}

func (u *UpdateToolCallUnitTestRequest) GetDynamicVariables() (value map[string]any) {
	if u == nil {
		return
	}
	return u.DynamicVariables
}

func (u *UpdateToolCallUnitTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if u == nil {
		return
	}
	return u.ChatHistory
}

func (u *UpdateToolCallUnitTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if u == nil {
		return
	}
	return u.ConversationInitiationSource
}

func (u *UpdateToolCallUnitTestRequest) GetEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.Environment
}

func (u *UpdateToolCallUnitTestRequest) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UpdateToolCallUnitTestRequest) GetToolCallParameters() (value *UnitTestToolCallEvaluationModelInput) {
	if u == nil {
		return
	}
	return u.ToolCallParameters
}

func (u *UpdateToolCallUnitTestRequest) GetCheckAnyToolMatches() (value *bool) {
	if u == nil {
		return
	}
	return u.CheckAnyToolMatches
}

func (u *UpdateToolCallUnitTestRequest) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateToolCallUnitTestRequest) GetParentFolderId() (value *string) {
	if u == nil {
		return
	}
	return u.ParentFolderId
}

func (u *UpdateToolCallUnitTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	u.FromConversationMetadata = fromConversationMetadata
}

func (u *UpdateToolCallUnitTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	u.DynamicVariables = dynamicVariables
}

func (u *UpdateToolCallUnitTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	u.ChatHistory = chatHistory
}

func (u *UpdateToolCallUnitTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	u.ConversationInitiationSource = conversationInitiationSource
}

func (u *UpdateToolCallUnitTestRequest) SetEnvironment(environment *string) {
	u.Environment = environment
}

func (u *UpdateToolCallUnitTestRequest) SetType(type_ *string) {
	u.Type = type_
}

func (u *UpdateToolCallUnitTestRequest) SetToolCallParameters(toolCallParameters *UnitTestToolCallEvaluationModelInput) {
	u.ToolCallParameters = toolCallParameters
}

func (u *UpdateToolCallUnitTestRequest) SetCheckAnyToolMatches(checkAnyToolMatches *bool) {
	u.CheckAnyToolMatches = checkAnyToolMatches
}

func (u *UpdateToolCallUnitTestRequest) SetName(name string) {
	u.Name = name
}

func (u *UpdateToolCallUnitTestRequest) SetParentFolderId(parentFolderId *string) {
	u.ParentFolderId = parentFolderId
}

func (u *UpdateToolCallUnitTestRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateToolCallUnitTestRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateToolCallUnitTestRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateToolCallUnitTestRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateToolCallUnitTestRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateToolCallUnitTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
