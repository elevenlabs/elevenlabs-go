package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateResponseUnitTestRequest struct {
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
	// The ID of the parent folder. If not provided, the test will be moved to the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpdateResponseUnitTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if u == nil {
		return
	}
	return u.FromConversationMetadata
}

func (u *UpdateResponseUnitTestRequest) GetDynamicVariables() (value map[string]any) {
	if u == nil {
		return
	}
	return u.DynamicVariables
}

func (u *UpdateResponseUnitTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if u == nil {
		return
	}
	return u.ChatHistory
}

func (u *UpdateResponseUnitTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if u == nil {
		return
	}
	return u.ConversationInitiationSource
}

func (u *UpdateResponseUnitTestRequest) GetEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.Environment
}

func (u *UpdateResponseUnitTestRequest) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UpdateResponseUnitTestRequest) GetSuccessCondition() (value *string) {
	if u == nil {
		return
	}
	return u.SuccessCondition
}

func (u *UpdateResponseUnitTestRequest) GetSuccessExamples() (value []*AgentSuccessfulResponseExample) {
	if u == nil {
		return
	}
	return u.SuccessExamples
}

func (u *UpdateResponseUnitTestRequest) GetFailureExamples() (value []*AgentFailureResponseExample) {
	if u == nil {
		return
	}
	return u.FailureExamples
}

func (u *UpdateResponseUnitTestRequest) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateResponseUnitTestRequest) GetParentFolderId() (value *string) {
	if u == nil {
		return
	}
	return u.ParentFolderId
}

func (u *UpdateResponseUnitTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	u.FromConversationMetadata = fromConversationMetadata
}

func (u *UpdateResponseUnitTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	u.DynamicVariables = dynamicVariables
}

func (u *UpdateResponseUnitTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	u.ChatHistory = chatHistory
}

func (u *UpdateResponseUnitTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	u.ConversationInitiationSource = conversationInitiationSource
}

func (u *UpdateResponseUnitTestRequest) SetEnvironment(environment *string) {
	u.Environment = environment
}

func (u *UpdateResponseUnitTestRequest) SetType(type_ *string) {
	u.Type = type_
}

func (u *UpdateResponseUnitTestRequest) SetSuccessCondition(successCondition *string) {
	u.SuccessCondition = successCondition
}

func (u *UpdateResponseUnitTestRequest) SetSuccessExamples(successExamples []*AgentSuccessfulResponseExample) {
	u.SuccessExamples = successExamples
}

func (u *UpdateResponseUnitTestRequest) SetFailureExamples(failureExamples []*AgentFailureResponseExample) {
	u.FailureExamples = failureExamples
}

func (u *UpdateResponseUnitTestRequest) SetName(name string) {
	u.Name = name
}

func (u *UpdateResponseUnitTestRequest) SetParentFolderId(parentFolderId *string) {
	u.ParentFolderId = parentFolderId
}

func (u *UpdateResponseUnitTestRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateResponseUnitTestRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateResponseUnitTestRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateResponseUnitTestRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpdateResponseUnitTestRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateResponseUnitTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
