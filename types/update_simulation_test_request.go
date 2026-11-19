package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateSimulationTestRequest struct {
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
	// The ID of the parent folder. If not provided, the test will be moved to the root level.
	ParentFolderId  *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpdateSimulationTestRequest) GetFromConversationMetadata() (value *TestFromConversationMetadataInput) {
	if u == nil {
		return
	}
	return u.FromConversationMetadata
}

func (u *UpdateSimulationTestRequest) GetDynamicVariables() (value map[string]any) {
	if u == nil {
		return
	}
	return u.DynamicVariables
}

func (u *UpdateSimulationTestRequest) GetChatHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if u == nil {
		return
	}
	return u.ChatHistory
}

func (u *UpdateSimulationTestRequest) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if u == nil {
		return
	}
	return u.ConversationInitiationSource
}

func (u *UpdateSimulationTestRequest) GetEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.Environment
}

func (u *UpdateSimulationTestRequest) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UpdateSimulationTestRequest) GetSuccessCondition() (value *string) {
	if u == nil {
		return
	}
	return u.SuccessCondition
}

func (u *UpdateSimulationTestRequest) GetSuccessConditions() (value []string) {
	if u == nil {
		return
	}
	return u.SuccessConditions
}

func (u *UpdateSimulationTestRequest) GetSimulationScenario() (value *string) {
	if u == nil {
		return
	}
	return u.SimulationScenario
}

func (u *UpdateSimulationTestRequest) GetSimulationMaxTurns() (value *int) {
	if u == nil {
		return
	}
	return u.SimulationMaxTurns
}

func (u *UpdateSimulationTestRequest) GetSimulationEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.SimulationEnvironment
}

func (u *UpdateSimulationTestRequest) GetToolMockConfig() (value *SimulationToolMockBehaviorConfig) {
	if u == nil {
		return
	}
	return u.ToolMockConfig
}

func (u *UpdateSimulationTestRequest) GetToolMockOverrides() (value map[string][]*ToolResponseMockConfigInput) {
	if u == nil {
		return
	}
	return u.ToolMockOverrides
}

func (u *UpdateSimulationTestRequest) GetEvaluationModel() (value *Llm) {
	if u == nil {
		return
	}
	return u.EvaluationModel
}

func (u *UpdateSimulationTestRequest) GetSimulatedUserModel() (value *Llm) {
	if u == nil {
		return
	}
	return u.SimulatedUserModel
}

func (u *UpdateSimulationTestRequest) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateSimulationTestRequest) GetParentFolderId() (value *string) {
	if u == nil {
		return
	}
	return u.ParentFolderId
}

func (u *UpdateSimulationTestRequest) SetFromConversationMetadata(fromConversationMetadata *TestFromConversationMetadataInput) {
	u.FromConversationMetadata = fromConversationMetadata
}

func (u *UpdateSimulationTestRequest) SetDynamicVariables(dynamicVariables map[string]any) {
	u.DynamicVariables = dynamicVariables
}

func (u *UpdateSimulationTestRequest) SetChatHistory(chatHistory []*ConversationHistoryTranscriptCommonModelInput) {
	u.ChatHistory = chatHistory
}

func (u *UpdateSimulationTestRequest) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	u.ConversationInitiationSource = conversationInitiationSource
}

func (u *UpdateSimulationTestRequest) SetEnvironment(environment *string) {
	u.Environment = environment
}

func (u *UpdateSimulationTestRequest) SetType(type_ *string) {
	u.Type = type_
}

func (u *UpdateSimulationTestRequest) SetSuccessCondition(successCondition *string) {
	u.SuccessCondition = successCondition
}

func (u *UpdateSimulationTestRequest) SetSuccessConditions(successConditions []string) {
	u.SuccessConditions = successConditions
}

func (u *UpdateSimulationTestRequest) SetSimulationScenario(simulationScenario *string) {
	u.SimulationScenario = simulationScenario
}

func (u *UpdateSimulationTestRequest) SetSimulationMaxTurns(simulationMaxTurns *int) {
	u.SimulationMaxTurns = simulationMaxTurns
}

func (u *UpdateSimulationTestRequest) SetSimulationEnvironment(simulationEnvironment *string) {
	u.SimulationEnvironment = simulationEnvironment
}

func (u *UpdateSimulationTestRequest) SetToolMockConfig(toolMockConfig *SimulationToolMockBehaviorConfig) {
	u.ToolMockConfig = toolMockConfig
}

func (u *UpdateSimulationTestRequest) SetToolMockOverrides(toolMockOverrides map[string][]*ToolResponseMockConfigInput) {
	u.ToolMockOverrides = toolMockOverrides
}

func (u *UpdateSimulationTestRequest) SetEvaluationModel(evaluationModel *Llm) {
	u.EvaluationModel = evaluationModel
}

func (u *UpdateSimulationTestRequest) SetSimulatedUserModel(simulatedUserModel *Llm) {
	u.SimulatedUserModel = simulatedUserModel
}

func (u *UpdateSimulationTestRequest) SetName(name string) {
	u.Name = name
}

func (u *UpdateSimulationTestRequest) SetParentFolderId(parentFolderId *string) {
	u.ParentFolderId = parentFolderId
}

func (u *UpdateSimulationTestRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateSimulationTestRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateSimulationTestRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateSimulationTestRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpdateSimulationTestRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateSimulationTestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
