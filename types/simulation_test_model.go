package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SimulationTestModel struct {
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
	SimulatedUserModel *Llm `json:"simulated_user_model,omitempty" xml:"simulated_user_model,omitempty"`
	extraProperties    map[string]interface{}
}

func (s *SimulationTestModel) GetFromConversationMetadata() (value *TestFromConversationMetadataOutput) {
	if s == nil {
		return
	}
	return s.FromConversationMetadata
}

func (s *SimulationTestModel) GetDynamicVariables() (value map[string]any) {
	if s == nil {
		return
	}
	return s.DynamicVariables
}

func (s *SimulationTestModel) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if s == nil {
		return
	}
	return s.ChatHistory
}

func (s *SimulationTestModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if s == nil {
		return
	}
	return s.ConversationInitiationSource
}

func (s *SimulationTestModel) GetEnvironment() (value *string) {
	if s == nil {
		return
	}
	return s.Environment
}

func (s *SimulationTestModel) GetSuccessCondition() (value *string) {
	if s == nil {
		return
	}
	return s.SuccessCondition
}

func (s *SimulationTestModel) GetSuccessConditions() (value []string) {
	if s == nil {
		return
	}
	return s.SuccessConditions
}

func (s *SimulationTestModel) GetSimulationScenario() (value *string) {
	if s == nil {
		return
	}
	return s.SimulationScenario
}

func (s *SimulationTestModel) GetSimulationMaxTurns() (value *int) {
	if s == nil {
		return
	}
	return s.SimulationMaxTurns
}

func (s *SimulationTestModel) GetSimulationEnvironment() (value *string) {
	if s == nil {
		return
	}
	return s.SimulationEnvironment
}

func (s *SimulationTestModel) GetToolMockConfig() (value *SimulationToolMockBehaviorConfig) {
	if s == nil {
		return
	}
	return s.ToolMockConfig
}

func (s *SimulationTestModel) GetToolMockOverrides() (value map[string][]*ToolResponseMockConfigOutput) {
	if s == nil {
		return
	}
	return s.ToolMockOverrides
}

func (s *SimulationTestModel) GetEvaluationModel() (value *Llm) {
	if s == nil {
		return
	}
	return s.EvaluationModel
}

func (s *SimulationTestModel) GetSimulatedUserModel() (value *Llm) {
	if s == nil {
		return
	}
	return s.SimulatedUserModel
}

func (s *SimulationTestModel) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataOutput) {
	s.FromConversationMetadata = fromConversationMetadata
}

func (s *SimulationTestModel) SetDynamicVariables(dynamicVariables map[string]any) {
	s.DynamicVariables = dynamicVariables
}

func (s *SimulationTestModel) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelOutput) {
	s.ChatHistory = chatHistory
}

func (s *SimulationTestModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	s.ConversationInitiationSource = conversationInitiationSource
}

func (s *SimulationTestModel) SetEnvironment(environment *string) {
	s.Environment = environment
}

func (s *SimulationTestModel) SetSuccessCondition(successCondition *string) {
	s.SuccessCondition = successCondition
}

func (s *SimulationTestModel) SetSuccessConditions(successConditions []string) {
	s.SuccessConditions = successConditions
}

func (s *SimulationTestModel) SetSimulationScenario(simulationScenario *string) {
	s.SimulationScenario = simulationScenario
}

func (s *SimulationTestModel) SetSimulationMaxTurns(simulationMaxTurns *int) {
	s.SimulationMaxTurns = simulationMaxTurns
}

func (s *SimulationTestModel) SetSimulationEnvironment(simulationEnvironment *string) {
	s.SimulationEnvironment = simulationEnvironment
}

func (s *SimulationTestModel) SetToolMockConfig(toolMockConfig *SimulationToolMockBehaviorConfig) {
	s.ToolMockConfig = toolMockConfig
}

func (s *SimulationTestModel) SetToolMockOverrides(toolMockOverrides map[string][]*ToolResponseMockConfigOutput) {
	s.ToolMockOverrides = toolMockOverrides
}

func (s *SimulationTestModel) SetEvaluationModel(evaluationModel *Llm) {
	s.EvaluationModel = evaluationModel
}

func (s *SimulationTestModel) SetSimulatedUserModel(simulatedUserModel *Llm) {
	s.SimulatedUserModel = simulatedUserModel
}

func (s *SimulationTestModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SimulationTestModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SimulationTestModel to string"
	}
	return string(jsonData)
}

func (s *SimulationTestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SimulationTestModel) MarshalJSON() ([]byte, error) {
	type embed SimulationTestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
