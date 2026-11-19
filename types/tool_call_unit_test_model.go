package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolCallUnitTestModel struct {
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata *TestFromConversationMetadataOutput `json:"from_conversation_metadata,omitempty" xml:"from_conversation_metadata,omitempty"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]any                                    `json:"dynamic_variables,omitempty"`
	ChatHistory      []*ConversationHistoryTranscriptCommonModelOutput `json:"chat_history,omitempty" xml:"chat_history,omitempty" maxItems:"200"`
	// Simulate the test as if the conversation originated from this channel.
	ConversationInitiationSource *ConversationInitiationSource `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	// The environment to resolve environment-specific variable values against when running this test (URL, headers, auth connections). If not provided, defaults to 'production'. For simulation tests, simulation_environment takes precedence when set.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not evaluated.
	ToolCallParameters *UnitTestToolCallEvaluationModelOutput `json:"tool_call_parameters,omitempty" xml:"tool_call_parameters,omitempty"`
	// If set to True this test will pass if any tool call returned by the LLM matches the criteria. Otherwise it will fail if more than one tool is returned by the agent.
	CheckAnyToolMatches *bool `json:"check_any_tool_matches,omitempty" xml:"check_any_tool_matches,omitempty"`
	extraProperties     map[string]interface{}
}

func (t *ToolCallUnitTestModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if t == nil {
		return
	}
	return t.FromConversationMetadata
}

func (t *ToolCallUnitTestModel) GetDynamicVariables() (value map[string]any) {
	if t == nil {
		return
	}
	return t.DynamicVariables
}

func (t *ToolCallUnitTestModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if t == nil {
		return
	}
	return t.ChatHistory
}

func (t *ToolCallUnitTestModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if t == nil {
		return
	}
	return t.ConversationInitiationSource
}

func (t *ToolCallUnitTestModel) GetEnvironment() (value *string) {
	if t == nil {
		return
	}
	return t.Environment
}

func (t *ToolCallUnitTestModel) GetToolCallParameters() (value *UnitTestToolCallEvaluationModelOutput) {
	if t == nil {
		return
	}
	return t.ToolCallParameters
}

func (t *ToolCallUnitTestModel) GetCheckAnyToolMatches() (value *bool) {
	if t == nil {
		return
	}
	return t.CheckAnyToolMatches
}

func (t *ToolCallUnitTestModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	t.FromConversationMetadata = fromConversationMetadata
}

func (t *ToolCallUnitTestModel) SetDynamicVariables(dynamicVariables map[string]any) {
	t.DynamicVariables = dynamicVariables
}

func (t *ToolCallUnitTestModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	t.ChatHistory = chatHistory
}

func (t *ToolCallUnitTestModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	t.ConversationInitiationSource = conversationInitiationSource
}

func (t *ToolCallUnitTestModel) SetEnvironment(environment *string) {
	t.Environment = environment
}

func (t *ToolCallUnitTestModel) SetToolCallParameters(toolCallParameters *UnitTestToolCallEvaluationModelOutput) {
	t.ToolCallParameters = toolCallParameters
}

func (t *ToolCallUnitTestModel) SetCheckAnyToolMatches(checkAnyToolMatches *bool) {
	t.CheckAnyToolMatches = checkAnyToolMatches
}

func (t *ToolCallUnitTestModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolCallUnitTestModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolCallUnitTestModel to string"
	}
	return string(jsonData)
}

func (t *ToolCallUnitTestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolCallUnitTestModel) MarshalJSON() ([]byte, error) {
	type embed ToolCallUnitTestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
