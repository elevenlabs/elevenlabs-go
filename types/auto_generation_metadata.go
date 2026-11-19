package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Metadata for an automatically generated simulation test.
//
// Tracks the origin and context of tests that were auto-generated from
// conversations, including the source conversation, topic, and evaluation
// criteria the test covers.
type AutoGenerationMetadata struct {
	// Agent whose simulation library this test belongs to.
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// Stable topic id (from topic discovery) this test covers.
	StableTopicId string `json:"stable_topic_id" xml:"stable_topic_id" required:"true"`
	// Evaluation criterion id this test targets.
	CriteriaId string `json:"criteria_id" xml:"criteria_id" required:"true"`
	// Human-readable topic label at generation time.
	TopicLabel string `json:"topic_label" xml:"topic_label" required:"true"`
	// Conversation the test was generated from.
	SourceConversationId string `json:"source_conversation_id" xml:"source_conversation_id" required:"true"`
	// Unix timestamp (seconds) when the test was generated.
	GeneratedAtUnix int `json:"generated_at_unix" xml:"generated_at_unix" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AutoGenerationMetadata) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AutoGenerationMetadata) GetStableTopicId() (value string) {
	if a == nil {
		return
	}
	return a.StableTopicId
}

func (a *AutoGenerationMetadata) GetCriteriaId() (value string) {
	if a == nil {
		return
	}
	return a.CriteriaId
}

func (a *AutoGenerationMetadata) GetTopicLabel() (value string) {
	if a == nil {
		return
	}
	return a.TopicLabel
}

func (a *AutoGenerationMetadata) GetSourceConversationId() (value string) {
	if a == nil {
		return
	}
	return a.SourceConversationId
}

func (a *AutoGenerationMetadata) GetGeneratedAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.GeneratedAtUnix
}

func (a *AutoGenerationMetadata) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AutoGenerationMetadata) SetStableTopicId(stableTopicId string) {
	a.StableTopicId = stableTopicId
}

func (a *AutoGenerationMetadata) SetCriteriaId(criteriaId string) {
	a.CriteriaId = criteriaId
}

func (a *AutoGenerationMetadata) SetTopicLabel(topicLabel string) {
	a.TopicLabel = topicLabel
}

func (a *AutoGenerationMetadata) SetSourceConversationId(sourceConversationId string) {
	a.SourceConversationId = sourceConversationId
}

func (a *AutoGenerationMetadata) SetGeneratedAtUnix(generatedAtUnix int) {
	a.GeneratedAtUnix = generatedAtUnix
}

func (a *AutoGenerationMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AutoGenerationMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AutoGenerationMetadata to string"
	}
	return string(jsonData)
}

func (a *AutoGenerationMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AutoGenerationMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AutoGenerationMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AutoGenerationMetadata) MarshalJSON() ([]byte, error) {
	type embed AutoGenerationMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
