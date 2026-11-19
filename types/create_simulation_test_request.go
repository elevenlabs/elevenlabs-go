package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateSimulationTestRequest struct {
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
	ToolMockOverrides map[string][]*ToolResponseMockConfigInput `json:"tool_mock_overrides,omitempty"`
	// LLM model to use for evaluating simulation results.
	EvaluationModel *Llm `json:"evaluation_model,omitempty" xml:"evaluation_model,omitempty"`
	// LLM model for the simulated user.
	SimulatedUserModel *Llm   `json:"simulated_user_model,omitempty" xml:"simulated_user_model,omitempty"`
	Name               string `json:"name" xml:"name" required:"true"`
	// The ID of the parent folder. If not provided, the test will be created at the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateSimulationTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if c == nil {
		return
	}
	return c.FromConversationMetadata
}

func (c *CreateSimulationTestRequest) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *CreateSimulationTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if c == nil {
		return
	}
	return c.ChatHistory
}

func (c *CreateSimulationTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSource
}

func (c *CreateSimulationTestRequest) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *CreateSimulationTestRequest) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CreateSimulationTestRequest) GetSuccessCondition() (value *string) {
	if c == nil {
		return
	}
	return c.SuccessCondition
}

func (c *CreateSimulationTestRequest) GetSuccessConditions() (value []string) {
	if c == nil {
		return
	}
	return c.SuccessConditions
}

func (c *CreateSimulationTestRequest) GetSimulationScenario() (value *string) {
	if c == nil {
		return
	}
	return c.SimulationScenario
}

func (c *CreateSimulationTestRequest) GetSimulationMaxTurns() (value *int) {
	if c == nil {
		return
	}
	return c.SimulationMaxTurns
}

func (c *CreateSimulationTestRequest) GetSimulationEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.SimulationEnvironment
}

func (c *CreateSimulationTestRequest) GetToolMockConfig() (value *SimulationToolMockBehaviorConfig) {
	if c == nil {
		return
	}
	return c.ToolMockConfig
}

func (c *CreateSimulationTestRequest) GetToolMockOverrides() (value map[string][]*ToolResponseMockConfigInput) {
	if c == nil {
		return
	}
	return c.ToolMockOverrides
}

func (c *CreateSimulationTestRequest) GetEvaluationModel() (value *Llm) {
	if c == nil {
		return
	}
	return c.EvaluationModel
}

func (c *CreateSimulationTestRequest) GetSimulatedUserModel() (value *Llm) {
	if c == nil {
		return
	}
	return c.SimulatedUserModel
}

func (c *CreateSimulationTestRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateSimulationTestRequest) GetParentFolderId() (value *string) {
	if c == nil {
		return
	}
	return c.ParentFolderId
}

func (c *CreateSimulationTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	c.FromConversationMetadata = fromConversationMetadata
}

func (c *CreateSimulationTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *CreateSimulationTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	c.ChatHistory = chatHistory
}

func (c *CreateSimulationTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	c.ConversationInitiationSource = conversationInitiationSource
}

func (c *CreateSimulationTestRequest) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *CreateSimulationTestRequest) SetType(type_ *string) {
	c.Type = type_
}

func (c *CreateSimulationTestRequest) SetSuccessCondition(successCondition *string) {
	c.SuccessCondition = successCondition
}

func (c *CreateSimulationTestRequest) SetSuccessConditions(successConditions []string) {
	c.SuccessConditions = successConditions
}

func (c *CreateSimulationTestRequest) SetSimulationScenario(simulationScenario *string) {
	c.SimulationScenario = simulationScenario
}

func (c *CreateSimulationTestRequest) SetSimulationMaxTurns(simulationMaxTurns *int) {
	c.SimulationMaxTurns = simulationMaxTurns
}

func (c *CreateSimulationTestRequest) SetSimulationEnvironment(simulationEnvironment *string) {
	c.SimulationEnvironment = simulationEnvironment
}

func (c *CreateSimulationTestRequest) SetToolMockConfig(toolMockConfig *SimulationToolMockBehaviorConfig) {
	c.ToolMockConfig = toolMockConfig
}

func (c *CreateSimulationTestRequest) SetToolMockOverrides(toolMockOverrides map[string][]*ToolResponseMockConfigInput) {
	c.ToolMockOverrides = toolMockOverrides
}

func (c *CreateSimulationTestRequest) SetEvaluationModel(evaluationModel *Llm) {
	c.EvaluationModel = evaluationModel
}

func (c *CreateSimulationTestRequest) SetSimulatedUserModel(simulatedUserModel *Llm) {
	c.SimulatedUserModel = simulatedUserModel
}

func (c *CreateSimulationTestRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateSimulationTestRequest) SetParentFolderId(parentFolderId *string) {
	c.ParentFolderId = parentFolderId
}

func (c *CreateSimulationTestRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateSimulationTestRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateSimulationTestRequest to string"
	}
	return string(jsonData)
}

func (c *CreateSimulationTestRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateSimulationTestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSimulationTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
