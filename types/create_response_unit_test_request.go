package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateResponseUnitTestRequest struct {
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
	// A prompt that evaluates whether the agent's response is successful. Should return True or False.
	SuccessCondition *string `json:"success_condition,omitempty" xml:"success_condition,omitempty"`
	// Non-empty list of example responses that should be considered successful
	SuccessExamples []*AgentSuccessfulResponseExample `json:"success_examples,omitempty" xml:"success_examples,omitempty" maxItems:"5"`
	// Non-empty list of example responses that should be considered failures
	FailureExamples []*AgentFailureResponseExample `json:"failure_examples,omitempty" xml:"failure_examples,omitempty" maxItems:"5"`
	Name            string                         `json:"name" xml:"name" required:"true"`
	// The ID of the parent folder. If not provided, the test will be created at the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateResponseUnitTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if c == nil {
		return
	}
	return c.FromConversationMetadata
}

func (c *CreateResponseUnitTestRequest) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *CreateResponseUnitTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if c == nil {
		return
	}
	return c.ChatHistory
}

func (c *CreateResponseUnitTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSource
}

func (c *CreateResponseUnitTestRequest) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *CreateResponseUnitTestRequest) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CreateResponseUnitTestRequest) GetSuccessCondition() (value *string) {
	if c == nil {
		return
	}
	return c.SuccessCondition
}

func (c *CreateResponseUnitTestRequest) GetSuccessExamples() (value []*AgentSuccessfulResponseExample) {
	if c == nil {
		return
	}
	return c.SuccessExamples
}

func (c *CreateResponseUnitTestRequest) GetFailureExamples() (value []*AgentFailureResponseExample) {
	if c == nil {
		return
	}
	return c.FailureExamples
}

func (c *CreateResponseUnitTestRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateResponseUnitTestRequest) GetParentFolderId() (value *string) {
	if c == nil {
		return
	}
	return c.ParentFolderId
}

func (c *CreateResponseUnitTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	c.FromConversationMetadata = fromConversationMetadata
}

func (c *CreateResponseUnitTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *CreateResponseUnitTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	c.ChatHistory = chatHistory
}

func (c *CreateResponseUnitTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	c.ConversationInitiationSource = conversationInitiationSource
}

func (c *CreateResponseUnitTestRequest) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *CreateResponseUnitTestRequest) SetType(type_ *string) {
	c.Type = type_
}

func (c *CreateResponseUnitTestRequest) SetSuccessCondition(successCondition *string) {
	c.SuccessCondition = successCondition
}

func (c *CreateResponseUnitTestRequest) SetSuccessExamples(successExamples []*AgentSuccessfulResponseExample) {
	c.SuccessExamples = successExamples
}

func (c *CreateResponseUnitTestRequest) SetFailureExamples(failureExamples []*AgentFailureResponseExample) {
	c.FailureExamples = failureExamples
}

func (c *CreateResponseUnitTestRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateResponseUnitTestRequest) SetParentFolderId(parentFolderId *string) {
	c.ParentFolderId = parentFolderId
}

func (c *CreateResponseUnitTestRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateResponseUnitTestRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateResponseUnitTestRequest to string"
	}
	return string(jsonData)
}

func (c *CreateResponseUnitTestRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateResponseUnitTestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateResponseUnitTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
