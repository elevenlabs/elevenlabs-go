package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResponseUnitTestModel struct {
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
	extraProperties map[string]interface{}
}

func (r *ResponseUnitTestModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if r == nil {
		return
	}
	return r.FromConversationMetadata
}

func (r *ResponseUnitTestModel) GetDynamicVariables() (value map[string]any) {
	if r == nil {
		return
	}
	return r.DynamicVariables
}

func (r *ResponseUnitTestModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if r == nil {
		return
	}
	return r.ChatHistory
}

func (r *ResponseUnitTestModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if r == nil {
		return
	}
	return r.ConversationInitiationSource
}

func (r *ResponseUnitTestModel) GetEnvironment() (value *string) {
	if r == nil {
		return
	}
	return r.Environment
}

func (r *ResponseUnitTestModel) GetSuccessCondition() (value *string) {
	if r == nil {
		return
	}
	return r.SuccessCondition
}

func (r *ResponseUnitTestModel) GetSuccessExamples() (value []*AgentSuccessfulResponseExample) {
	if r == nil {
		return
	}
	return r.SuccessExamples
}

func (r *ResponseUnitTestModel) GetFailureExamples() (value []*AgentFailureResponseExample) {
	if r == nil {
		return
	}
	return r.FailureExamples
}

func (r *ResponseUnitTestModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	r.FromConversationMetadata = fromConversationMetadata
}

func (r *ResponseUnitTestModel) SetDynamicVariables(dynamicVariables map[string]any) {
	r.DynamicVariables = dynamicVariables
}

func (r *ResponseUnitTestModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	r.ChatHistory = chatHistory
}

func (r *ResponseUnitTestModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	r.ConversationInitiationSource = conversationInitiationSource
}

func (r *ResponseUnitTestModel) SetEnvironment(environment *string) {
	r.Environment = environment
}

func (r *ResponseUnitTestModel) SetSuccessCondition(successCondition *string) {
	r.SuccessCondition = successCondition
}

func (r *ResponseUnitTestModel) SetSuccessExamples(successExamples []*AgentSuccessfulResponseExample) {
	r.SuccessExamples = successExamples
}

func (r *ResponseUnitTestModel) SetFailureExamples(failureExamples []*AgentFailureResponseExample) {
	r.FailureExamples = failureExamples
}

func (r *ResponseUnitTestModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResponseUnitTestModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResponseUnitTestModel to string"
	}
	return string(jsonData)
}

func (r *ResponseUnitTestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResponseUnitTestModel) MarshalJSON() ([]byte, error) {
	type embed ResponseUnitTestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
