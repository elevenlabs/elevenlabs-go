package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentTopicResponseModel struct {
	TopicId           string                 `json:"topic_id" xml:"topic_id" required:"true"`
	Label             string                 `json:"label" xml:"label" required:"true"`
	Description       string                 `json:"description" xml:"description" required:"true"`
	ConversationCount int                    `json:"conversation_count" xml:"conversation_count" required:"true"`
	ParentTopicId     *string                `json:"parent_topic_id,omitempty" xml:"parent_topic_id,omitempty"`
	X2D               *float64               `json:"x_2d,omitempty" xml:"x_2d,omitempty"`
	Y2D               *float64               `json:"y_2d,omitempty" xml:"y_2d,omitempty"`
	Metrics           *TopicMetricsAggregate `json:"metrics,omitempty" xml:"metrics,omitempty"`
	// Success rate across the topic's evaluation criteria, weighted by scored conversations. Returned regardless of include_evaluation_criteria.
	SuccessRate     *float64 `json:"success_rate,omitempty" xml:"success_rate,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentTopicResponseModel) GetTopicId() (value string) {
	if a == nil {
		return
	}
	return a.TopicId
}

func (a *AgentTopicResponseModel) GetLabel() (value string) {
	if a == nil {
		return
	}
	return a.Label
}

func (a *AgentTopicResponseModel) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AgentTopicResponseModel) GetConversationCount() (value int) {
	if a == nil {
		return
	}
	return a.ConversationCount
}

func (a *AgentTopicResponseModel) GetParentTopicId() (value *string) {
	if a == nil {
		return
	}
	return a.ParentTopicId
}

func (a *AgentTopicResponseModel) GetX2D() (value *float64) {
	if a == nil {
		return
	}
	return a.X2D
}

func (a *AgentTopicResponseModel) GetY2D() (value *float64) {
	if a == nil {
		return
	}
	return a.Y2D
}

func (a *AgentTopicResponseModel) GetMetrics() (value *TopicMetricsAggregate) {
	if a == nil {
		return
	}
	return a.Metrics
}

func (a *AgentTopicResponseModel) GetSuccessRate() (value *float64) {
	if a == nil {
		return
	}
	return a.SuccessRate
}

func (a *AgentTopicResponseModel) SetTopicId(topicId string) {
	a.TopicId = topicId
}

func (a *AgentTopicResponseModel) SetLabel(label string) {
	a.Label = label
}

func (a *AgentTopicResponseModel) SetDescription(description string) {
	a.Description = description
}

func (a *AgentTopicResponseModel) SetConversationCount(conversationCount int) {
	a.ConversationCount = conversationCount
}

func (a *AgentTopicResponseModel) SetParentTopicId(parentTopicId *string) {
	a.ParentTopicId = parentTopicId
}

func (a *AgentTopicResponseModel) SetX2D(x2D *float64) {
	a.X2D = x2D
}

func (a *AgentTopicResponseModel) SetY2D(y2D *float64) {
	a.Y2D = y2D
}

func (a *AgentTopicResponseModel) SetMetrics(metrics *TopicMetricsAggregate) {
	a.Metrics = metrics
}

func (a *AgentTopicResponseModel) SetSuccessRate(successRate *float64) {
	a.SuccessRate = successRate
}

func (a *AgentTopicResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTopicResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTopicResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentTopicResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTopicResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentTopicResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
