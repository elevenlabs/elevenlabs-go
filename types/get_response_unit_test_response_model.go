package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetResponseUnitTestResponseModel struct {
	// The access information for the requesting user on this test.
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata *TestFromConversationMetadataOutput `json:"from_conversation_metadata,omitempty" xml:"from_conversation_metadata,omitempty"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]any                                    `json:"dynamic_variables,omitempty"`
	ChatHistory      []*ConversationHistoryTranscriptCommonModelOutput `json:"chat_history,omitempty" xml:"chat_history,omitempty" maxItems:"200"`
	// Simulate the test as if the conversation originated from this channel.
	ConversationInitiationSource *ConversationInitiationSource `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	// The environment to resolve environment-specific variable values against when running this test (URL, headers, auth connections). If not provided, defaults to 'production'. For simulation tests, simulation_environment takes precedence when set.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// A prompt that evaluates whether the agent's response is successful. Should return True or False.
	SuccessCondition *string `json:"success_condition,omitempty" xml:"success_condition,omitempty"`
	// Non-empty list of example responses that should be considered successful
	SuccessExamples []*AgentSuccessfulResponseExample `json:"success_examples,omitempty" xml:"success_examples,omitempty" maxItems:"5"`
	// Non-empty list of example responses that should be considered failures
	FailureExamples []*AgentFailureResponseExample `json:"failure_examples,omitempty" xml:"failure_examples,omitempty" maxItems:"5"`
	Id              string                         `json:"id" xml:"id" required:"true"`
	Name            string                         `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetResponseUnitTestResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetResponseUnitTestResponseModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if g == nil {
		return
	}
	return g.FromConversationMetadata
}

func (g *GetResponseUnitTestResponseModel) GetDynamicVariables() (value map[string]any) {
	if g == nil {
		return
	}
	return g.DynamicVariables
}

func (g *GetResponseUnitTestResponseModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if g == nil {
		return
	}
	return g.ChatHistory
}

func (g *GetResponseUnitTestResponseModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if g == nil {
		return
	}
	return g.ConversationInitiationSource
}

func (g *GetResponseUnitTestResponseModel) GetEnvironment() (value *string) {
	if g == nil {
		return
	}
	return g.Environment
}

func (g *GetResponseUnitTestResponseModel) GetSuccessCondition() (value *string) {
	if g == nil {
		return
	}
	return g.SuccessCondition
}

func (g *GetResponseUnitTestResponseModel) GetSuccessExamples() (value []*AgentSuccessfulResponseExample) {
	if g == nil {
		return
	}
	return g.SuccessExamples
}

func (g *GetResponseUnitTestResponseModel) GetFailureExamples() (value []*AgentFailureResponseExample) {
	if g == nil {
		return
	}
	return g.FailureExamples
}

func (g *GetResponseUnitTestResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetResponseUnitTestResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetResponseUnitTestResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetResponseUnitTestResponseModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	g.FromConversationMetadata = fromConversationMetadata
}

func (g *GetResponseUnitTestResponseModel) SetDynamicVariables(dynamicVariables map[string]any) {
	g.DynamicVariables = dynamicVariables
}

func (g *GetResponseUnitTestResponseModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	g.ChatHistory = chatHistory
}

func (g *GetResponseUnitTestResponseModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	g.ConversationInitiationSource = conversationInitiationSource
}

func (g *GetResponseUnitTestResponseModel) SetEnvironment(environment *string) {
	g.Environment = environment
}

func (g *GetResponseUnitTestResponseModel) SetSuccessCondition(successCondition *string) {
	g.SuccessCondition = successCondition
}

func (g *GetResponseUnitTestResponseModel) SetSuccessExamples(successExamples []*AgentSuccessfulResponseExample) {
	g.SuccessExamples = successExamples
}

func (g *GetResponseUnitTestResponseModel) SetFailureExamples(failureExamples []*AgentFailureResponseExample) {
	g.FailureExamples = failureExamples
}

func (g *GetResponseUnitTestResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetResponseUnitTestResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetResponseUnitTestResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetResponseUnitTestResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetResponseUnitTestResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetResponseUnitTestResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetResponseUnitTestResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetResponseUnitTestResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
