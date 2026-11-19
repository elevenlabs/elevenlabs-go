package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost struct {
	// A specification that will be used to simulate a conversation between an agent and an AI user.
	SimulationSpecification *ConversationSimulationSpecification `json:"simulation_specification" xml:"simulation_specification" required:"true"`
	// A list of evaluation criteria to test
	ExtraEvaluationCriteria []*PromptEvaluationCriteria `json:"extra_evaluation_criteria,omitempty" xml:"extra_evaluation_criteria,omitempty"`
	// Maximum number of new turns to generate in the conversation simulation
	NewTurnsLimit   *int `json:"new_turns_limit,omitempty" xml:"new_turns_limit,omitempty" max:"100"`
	extraProperties map[string]interface{}
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) GetSimulationSpecification() (value *ConversationSimulationSpecification) {
	if b == nil {
		return
	}
	return b.SimulationSpecification
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) GetExtraEvaluationCriteria() (value []*PromptEvaluationCriteria) {
	if b == nil {
		return
	}
	return b.ExtraEvaluationCriteria
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) GetNewTurnsLimit() (value *int) {
	if b == nil {
		return
	}
	return b.NewTurnsLimit
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) SetSimulationSpecification(simulationSpecification *ConversationSimulationSpecification) {
	b.SimulationSpecification = simulationSpecification
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) SetExtraEvaluationCriteria(extraEvaluationCriteria []*PromptEvaluationCriteria) {
	b.ExtraEvaluationCriteria = extraEvaluationCriteria
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) SetNewTurnsLimit(newTurnsLimit *int) {
	b.NewTurnsLimit = newTurnsLimit
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost to string"
	}
	return string(jsonData)
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost) MarshalJSON() ([]byte, error) {
	type embed BodySimulatesAConversationStreamV1ConvaiAgentsAgentIdSimulateConversationStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
