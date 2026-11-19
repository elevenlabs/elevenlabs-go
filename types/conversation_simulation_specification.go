package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A specification that will be used to simulate a conversation between an agent and an AI user.
type ConversationSimulationSpecification struct {
	SimulatedUserConfig *AgentConfig               `json:"simulated_user_config" xml:"simulated_user_config" required:"true"`
	ToolMockConfig      map[string]*ToolMockConfig `json:"tool_mock_config,omitempty"`
	// A partial conversation history to start the simulation from. If empty, simulation starts fresh.
	PartialConversationHistory []*ConversationHistoryTranscriptCommonModelInput `json:"partial_conversation_history,omitempty" xml:"partial_conversation_history,omitempty"`
	DynamicVariables           map[string]any                                   `json:"dynamic_variables,omitempty"`
	extraProperties            map[string]interface{}
}

func (c *ConversationSimulationSpecification) GetSimulatedUserConfig() (value *AgentConfig) {
	if c == nil {
		return
	}
	return c.SimulatedUserConfig
}

func (c *ConversationSimulationSpecification) GetToolMockConfig() (value map[string]*ToolMockConfig) {
	if c == nil {
		return
	}
	return c.ToolMockConfig
}

func (c *ConversationSimulationSpecification) GetPartialConversationHistory() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if c == nil {
		return
	}
	return c.PartialConversationHistory
}

func (c *ConversationSimulationSpecification) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *ConversationSimulationSpecification) SetSimulatedUserConfig(simulatedUserConfig *AgentConfig) {
	c.SimulatedUserConfig = simulatedUserConfig
}

func (c *ConversationSimulationSpecification) SetToolMockConfig(toolMockConfig map[string]*ToolMockConfig) {
	c.ToolMockConfig = toolMockConfig
}

func (c *ConversationSimulationSpecification) SetPartialConversationHistory(partialConversationHistory []*ConversationHistoryTranscriptCommonModelInput) {
	c.PartialConversationHistory = partialConversationHistory
}

func (c *ConversationSimulationSpecification) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *ConversationSimulationSpecification) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSimulationSpecification) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSimulationSpecification to string"
	}
	return string(jsonData)
}

func (c *ConversationSimulationSpecification) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationSimulationSpecification
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationSimulationSpecification(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationSimulationSpecification) MarshalJSON() ([]byte, error) {
	type embed ConversationSimulationSpecification
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
