package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetToolCallUnitTestResponseModel struct {
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
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not evaluated.
	ToolCallParameters *UnitTestToolCallEvaluationModelOutput `json:"tool_call_parameters,omitempty" xml:"tool_call_parameters,omitempty"`
	// If set to True this test will pass if any tool call returned by the LLM matches the criteria. Otherwise it will fail if more than one tool is returned by the agent.
	CheckAnyToolMatches *bool  `json:"check_any_tool_matches,omitempty" xml:"check_any_tool_matches,omitempty"`
	Id                  string `json:"id" xml:"id" required:"true"`
	Name                string `json:"name" xml:"name" required:"true"`
	extraProperties     map[string]interface{}
}

func (g *GetToolCallUnitTestResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetToolCallUnitTestResponseModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if g == nil {
		return
	}
	return g.FromConversationMetadata
}

func (g *GetToolCallUnitTestResponseModel) GetDynamicVariables() (value map[string]any) {
	if g == nil {
		return
	}
	return g.DynamicVariables
}

func (g *GetToolCallUnitTestResponseModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if g == nil {
		return
	}
	return g.ChatHistory
}

func (g *GetToolCallUnitTestResponseModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if g == nil {
		return
	}
	return g.ConversationInitiationSource
}

func (g *GetToolCallUnitTestResponseModel) GetEnvironment() (value *string) {
	if g == nil {
		return
	}
	return g.Environment
}

func (g *GetToolCallUnitTestResponseModel) GetToolCallParameters() (value *UnitTestToolCallEvaluationModelOutput) {
	if g == nil {
		return
	}
	return g.ToolCallParameters
}

func (g *GetToolCallUnitTestResponseModel) GetCheckAnyToolMatches() (value *bool) {
	if g == nil {
		return
	}
	return g.CheckAnyToolMatches
}

func (g *GetToolCallUnitTestResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetToolCallUnitTestResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetToolCallUnitTestResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetToolCallUnitTestResponseModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	g.FromConversationMetadata = fromConversationMetadata
}

func (g *GetToolCallUnitTestResponseModel) SetDynamicVariables(dynamicVariables map[string]any) {
	g.DynamicVariables = dynamicVariables
}

func (g *GetToolCallUnitTestResponseModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	g.ChatHistory = chatHistory
}

func (g *GetToolCallUnitTestResponseModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	g.ConversationInitiationSource = conversationInitiationSource
}

func (g *GetToolCallUnitTestResponseModel) SetEnvironment(environment *string) {
	g.Environment = environment
}

func (g *GetToolCallUnitTestResponseModel) SetToolCallParameters(toolCallParameters *UnitTestToolCallEvaluationModelOutput) {
	g.ToolCallParameters = toolCallParameters
}

func (g *GetToolCallUnitTestResponseModel) SetCheckAnyToolMatches(checkAnyToolMatches *bool) {
	g.CheckAnyToolMatches = checkAnyToolMatches
}

func (g *GetToolCallUnitTestResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetToolCallUnitTestResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetToolCallUnitTestResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetToolCallUnitTestResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetToolCallUnitTestResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetToolCallUnitTestResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetToolCallUnitTestResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetToolCallUnitTestResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
