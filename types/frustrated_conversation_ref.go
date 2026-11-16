package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FrustratedConversationRef struct {
	ConversationId          string                                    `json:"conversation_id" xml:"conversation_id" required:"true"`
	AgentId                 string                                    `json:"agent_id" xml:"agent_id" required:"true"`
	StartTimeUnixSecs       int                                       `json:"start_time_unix_secs" xml:"start_time_unix_secs" required:"true"`
	OverallLabel            ConversationSentimentAnalysisOverallLabel `json:"overall_label" xml:"overall_label" required:"true"`
	OverallSentimentScore   float64                                   `json:"overall_sentiment_score" xml:"overall_sentiment_score" required:"true"`
	OverallFrustrationScore float64                                   `json:"overall_frustration_score" xml:"overall_frustration_score" required:"true"`
	extraProperties         map[string]interface{}
}

func (f *FrustratedConversationRef) GetConversationId() (value string) {
	if f == nil {
		return
	}
	return f.ConversationId
}

func (f *FrustratedConversationRef) GetAgentId() (value string) {
	if f == nil {
		return
	}
	return f.AgentId
}

func (f *FrustratedConversationRef) GetStartTimeUnixSecs() (value int) {
	if f == nil {
		return
	}
	return f.StartTimeUnixSecs
}

func (f *FrustratedConversationRef) GetOverallLabel() (value ConversationSentimentAnalysisOverallLabel) {
	if f == nil {
		return
	}
	return f.OverallLabel
}

func (f *FrustratedConversationRef) GetOverallSentimentScore() (value float64) {
	if f == nil {
		return
	}
	return f.OverallSentimentScore
}

func (f *FrustratedConversationRef) GetOverallFrustrationScore() (value float64) {
	if f == nil {
		return
	}
	return f.OverallFrustrationScore
}

func (f *FrustratedConversationRef) SetConversationId(conversationId string) {
	f.ConversationId = conversationId
}

func (f *FrustratedConversationRef) SetAgentId(agentId string) {
	f.AgentId = agentId
}

func (f *FrustratedConversationRef) SetStartTimeUnixSecs(startTimeUnixSecs int) {
	f.StartTimeUnixSecs = startTimeUnixSecs
}

func (f *FrustratedConversationRef) SetOverallLabel(overallLabel ConversationSentimentAnalysisOverallLabel) {
	f.OverallLabel = overallLabel
}

func (f *FrustratedConversationRef) SetOverallSentimentScore(overallSentimentScore float64) {
	f.OverallSentimentScore = overallSentimentScore
}

func (f *FrustratedConversationRef) SetOverallFrustrationScore(overallFrustrationScore float64) {
	f.OverallFrustrationScore = overallFrustrationScore
}

func (f *FrustratedConversationRef) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FrustratedConversationRef) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FrustratedConversationRef to string"
	}
	return string(jsonData)
}

func (f *FrustratedConversationRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FrustratedConversationRef
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FrustratedConversationRef(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FrustratedConversationRef) MarshalJSON() ([]byte, error) {
	type embed FrustratedConversationRef
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
