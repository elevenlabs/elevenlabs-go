package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetSimulationTestResponseModel struct {
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
	// Deprecated legacy single success criterion. Use success_conditions instead. At least one of success_condition or success_conditions is required.
	SuccessCondition *string `json:"success_condition,omitempty" xml:"success_condition,omitempty"`
	// List of prompts that evaluate whether the simulation was successful. If provided, all criteria are evaluated and merged into a final result. Capped at the maximum number of evaluation criteria.
	SuccessConditions []string `json:"success_conditions,omitempty" xml:"success_conditions,omitempty" maxItems:"30"`
	// Description of the simulation scenario and user persona for simulation tests.
	SimulationScenario *string `json:"simulation_scenario,omitempty" xml:"simulation_scenario,omitempty"`
	// Maximum number of conversation turns for simulation tests.
	SimulationMaxTurns *int `json:"simulation_max_turns,omitempty" xml:"simulation_max_turns,omitempty" min:"1" max:"50"`
	// The environment to use when running this simulation test. If not provided, defaults to 'production'.
	SimulationEnvironment *string `json:"simulation_environment,omitempty" xml:"simulation_environment,omitempty"`
	// Simulation/preview-side config: tools are identified by IDs, resolved to names at runtime.
	ToolMockConfig *SimulationToolMockBehaviorConfig `json:"tool_mock_config,omitempty" xml:"tool_mock_config,omitempty"`
	// Test-specific response mocks, keyed by tool ID. Applied ahead of the tool's shared mocks and only within this test. Only take effect for tools that are mocked (see tool_mock_config).
	ToolMockOverrides map[string][]*ToolResponseMockConfigOutput `json:"tool_mock_overrides,omitempty"`
	// LLM model to use for evaluating simulation results.
	EvaluationModel *Llm `json:"evaluation_model,omitempty" xml:"evaluation_model,omitempty"`
	// LLM model for the simulated user.
	SimulatedUserModel *Llm   `json:"simulated_user_model,omitempty" xml:"simulated_user_model,omitempty"`
	Id                 string `json:"id" xml:"id" required:"true"`
	Name               string `json:"name" xml:"name" required:"true"`
	extraProperties    map[string]interface{}
}

func (g *GetSimulationTestResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetSimulationTestResponseModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if g == nil {
		return
	}
	return g.FromConversationMetadata
}

func (g *GetSimulationTestResponseModel) GetDynamicVariables() (value map[string]any) {
	if g == nil {
		return
	}
	return g.DynamicVariables
}

func (g *GetSimulationTestResponseModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if g == nil {
		return
	}
	return g.ChatHistory
}

func (g *GetSimulationTestResponseModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if g == nil {
		return
	}
	return g.ConversationInitiationSource
}

func (g *GetSimulationTestResponseModel) GetEnvironment() (value *string) {
	if g == nil {
		return
	}
	return g.Environment
}

func (g *GetSimulationTestResponseModel) GetSuccessCondition() (value *string) {
	if g == nil {
		return
	}
	return g.SuccessCondition
}

func (g *GetSimulationTestResponseModel) GetSuccessConditions() (value []string) {
	if g == nil {
		return
	}
	return g.SuccessConditions
}

func (g *GetSimulationTestResponseModel) GetSimulationScenario() (value *string) {
	if g == nil {
		return
	}
	return g.SimulationScenario
}

func (g *GetSimulationTestResponseModel) GetSimulationMaxTurns() (value *int) {
	if g == nil {
		return
	}
	return g.SimulationMaxTurns
}

func (g *GetSimulationTestResponseModel) GetSimulationEnvironment() (value *string) {
	if g == nil {
		return
	}
	return g.SimulationEnvironment
}

func (g *GetSimulationTestResponseModel) GetToolMockConfig() (value *SimulationToolMockBehaviorConfig) {
	if g == nil {
		return
	}
	return g.ToolMockConfig
}

func (g *GetSimulationTestResponseModel) GetToolMockOverrides() (value map[string][]*ToolResponseMockConfigOutput) {
	if g == nil {
		return
	}
	return g.ToolMockOverrides
}

func (g *GetSimulationTestResponseModel) GetEvaluationModel() (value *Llm) {
	if g == nil {
		return
	}
	return g.EvaluationModel
}

func (g *GetSimulationTestResponseModel) GetSimulatedUserModel() (value *Llm) {
	if g == nil {
		return
	}
	return g.SimulatedUserModel
}

func (g *GetSimulationTestResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetSimulationTestResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetSimulationTestResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetSimulationTestResponseModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	g.FromConversationMetadata = fromConversationMetadata
}

func (g *GetSimulationTestResponseModel) SetDynamicVariables(dynamicVariables map[string]any) {
	g.DynamicVariables = dynamicVariables
}

func (g *GetSimulationTestResponseModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	g.ChatHistory = chatHistory
}

func (g *GetSimulationTestResponseModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	g.ConversationInitiationSource = conversationInitiationSource
}

func (g *GetSimulationTestResponseModel) SetEnvironment(environment *string) {
	g.Environment = environment
}

func (g *GetSimulationTestResponseModel) SetSuccessCondition(successCondition *string) {
	g.SuccessCondition = successCondition
}

func (g *GetSimulationTestResponseModel) SetSuccessConditions(successConditions []string) {
	g.SuccessConditions = successConditions
}

func (g *GetSimulationTestResponseModel) SetSimulationScenario(simulationScenario *string) {
	g.SimulationScenario = simulationScenario
}

func (g *GetSimulationTestResponseModel) SetSimulationMaxTurns(simulationMaxTurns *int) {
	g.SimulationMaxTurns = simulationMaxTurns
}

func (g *GetSimulationTestResponseModel) SetSimulationEnvironment(simulationEnvironment *string) {
	g.SimulationEnvironment = simulationEnvironment
}

func (g *GetSimulationTestResponseModel) SetToolMockConfig(toolMockConfig *SimulationToolMockBehaviorConfig) {
	g.ToolMockConfig = toolMockConfig
}

func (g *GetSimulationTestResponseModel) SetToolMockOverrides(toolMockOverrides map[string][]*ToolResponseMockConfigOutput) {
	g.ToolMockOverrides = toolMockOverrides
}

func (g *GetSimulationTestResponseModel) SetEvaluationModel(evaluationModel *Llm) {
	g.EvaluationModel = evaluationModel
}

func (g *GetSimulationTestResponseModel) SetSimulatedUserModel(simulatedUserModel *Llm) {
	g.SimulatedUserModel = simulatedUserModel
}

func (g *GetSimulationTestResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetSimulationTestResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetSimulationTestResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetSimulationTestResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetSimulationTestResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetSimulationTestResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetSimulationTestResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetSimulationTestResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
