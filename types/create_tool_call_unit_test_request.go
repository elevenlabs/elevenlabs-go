package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateToolCallUnitTestRequest struct {
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
	// The ID of the parent folder. If not provided, the test will be created at the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateToolCallUnitTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if c == nil {
		return
	}
	return c.FromConversationMetadata
}

func (c *CreateToolCallUnitTestRequest) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *CreateToolCallUnitTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if c == nil {
		return
	}
	return c.ChatHistory
}

func (c *CreateToolCallUnitTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSource
}

func (c *CreateToolCallUnitTestRequest) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *CreateToolCallUnitTestRequest) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CreateToolCallUnitTestRequest) GetToolCallParameters() (value *UnitTestToolCallEvaluationModelInput) {
	if c == nil {
		return
	}
	return c.ToolCallParameters
}

func (c *CreateToolCallUnitTestRequest) GetCheckAnyToolMatches() (value *bool) {
	if c == nil {
		return
	}
	return c.CheckAnyToolMatches
}

func (c *CreateToolCallUnitTestRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateToolCallUnitTestRequest) GetParentFolderId() (value *string) {
	if c == nil {
		return
	}
	return c.ParentFolderId
}

func (c *CreateToolCallUnitTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	c.FromConversationMetadata = fromConversationMetadata
}

func (c *CreateToolCallUnitTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *CreateToolCallUnitTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	c.ChatHistory = chatHistory
}

func (c *CreateToolCallUnitTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	c.ConversationInitiationSource = conversationInitiationSource
}

func (c *CreateToolCallUnitTestRequest) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *CreateToolCallUnitTestRequest) SetType(type_ *string) {
	c.Type = type_
}

func (c *CreateToolCallUnitTestRequest) SetToolCallParameters(toolCallParameters *UnitTestToolCallEvaluationModelInput) {
	c.ToolCallParameters = toolCallParameters
}

func (c *CreateToolCallUnitTestRequest) SetCheckAnyToolMatches(checkAnyToolMatches *bool) {
	c.CheckAnyToolMatches = checkAnyToolMatches
}

func (c *CreateToolCallUnitTestRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateToolCallUnitTestRequest) SetParentFolderId(parentFolderId *string) {
	c.ParentFolderId = parentFolderId
}

func (c *CreateToolCallUnitTestRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateToolCallUnitTestRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateToolCallUnitTestRequest to string"
	}
	return string(jsonData)
}

func (c *CreateToolCallUnitTestRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateToolCallUnitTestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateToolCallUnitTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
