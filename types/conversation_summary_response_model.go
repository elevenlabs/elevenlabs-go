package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationSummaryResponseModel struct {
	AgentId           string  `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId          *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	VersionId         *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	AgentName         *string `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
	ConversationId    string  `json:"conversation_id" xml:"conversation_id" required:"true"`
	StartTimeUnixSecs int     `json:"start_time_unix_secs" xml:"start_time_unix_secs" required:"true"`
	CallDurationSecs  int     `json:"call_duration_secs" xml:"call_duration_secs" required:"true"`
	// Total cost of the conversation in credits.
	Cost *int `json:"cost,omitempty" xml:"cost,omitempty"`
	// Total fiat cost of the conversation in USD. `None` for conversations that predate fiat cost tracking.
	CostFiat *float64 `json:"cost_fiat,omitempty" xml:"cost_fiat,omitempty"`
	// Number of transcript turns with text; tool-only turns are not counted.
	MessageCount int `json:"message_count" xml:"message_count" required:"true"`
	// Number of user transcript turns with text; tool-only turns are not counted. `null` for older conversations before this field whose transcript was deleted.
	UserMessageCount *int `json:"user_message_count,omitempty" xml:"user_message_count,omitempty"`
	// Number of agent transcript turns with text; tool-only turns are not counted. `null` for older conversations before this field whose transcript was deleted.
	AgentMessageCount            *int                                        `json:"agent_message_count,omitempty" xml:"agent_message_count,omitempty"`
	Status                       ConversationSummaryResponseModelStatus      `json:"status" xml:"status" required:"true"`
	TerminationReason            *string                                     `json:"termination_reason,omitempty" xml:"termination_reason,omitempty"`
	CallSuccessful               EvaluationSuccessResult                     `json:"call_successful" xml:"call_successful" required:"true"`
	CallSuccessScore             *float64                                    `json:"call_success_score,omitempty" xml:"call_success_score,omitempty"`
	TranscriptSummary            *string                                     `json:"transcript_summary,omitempty" xml:"transcript_summary,omitempty"`
	CallSummaryTitle             *string                                     `json:"call_summary_title,omitempty" xml:"call_summary_title,omitempty"`
	MainLanguage                 *string                                     `json:"main_language,omitempty" xml:"main_language,omitempty"`
	ConversationInitiationSource *ConversationInitiationSource               `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	ToolNames                    []string                                    `json:"tool_names,omitempty" xml:"tool_names,omitempty"`
	Direction                    *TelephonyDirection                         `json:"direction,omitempty" xml:"direction,omitempty"`
	Rating                       *float64                                    `json:"rating,omitempty" xml:"rating,omitempty"`
	SentimentAnalysis            *ConversationSentimentAnalysis              `json:"sentiment_analysis,omitempty" xml:"sentiment_analysis,omitempty"`
	DataCollectionResults        map[string]any                              `json:"data_collection_results,omitempty"`
	EvaluationCriteriaResults    map[string]*EvaluationCriteriaSummaryResult `json:"evaluation_criteria_results,omitempty"`
	// Conversation tag ids assigned to this conversation.
	TagIds          []string `json:"tag_ids,omitempty" xml:"tag_ids,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationSummaryResponseModel) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *ConversationSummaryResponseModel) GetBranchId() (value *string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *ConversationSummaryResponseModel) GetVersionId() (value *string) {
	if c == nil {
		return
	}
	return c.VersionId
}

func (c *ConversationSummaryResponseModel) GetAgentName() (value *string) {
	if c == nil {
		return
	}
	return c.AgentName
}

func (c *ConversationSummaryResponseModel) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *ConversationSummaryResponseModel) GetStartTimeUnixSecs() (value int) {
	if c == nil {
		return
	}
	return c.StartTimeUnixSecs
}

func (c *ConversationSummaryResponseModel) GetCallDurationSecs() (value int) {
	if c == nil {
		return
	}
	return c.CallDurationSecs
}

func (c *ConversationSummaryResponseModel) GetCost() (value *int) {
	if c == nil {
		return
	}
	return c.Cost
}

func (c *ConversationSummaryResponseModel) GetCostFiat() (value *float64) {
	if c == nil {
		return
	}
	return c.CostFiat
}

func (c *ConversationSummaryResponseModel) GetMessageCount() (value int) {
	if c == nil {
		return
	}
	return c.MessageCount
}

func (c *ConversationSummaryResponseModel) GetUserMessageCount() (value *int) {
	if c == nil {
		return
	}
	return c.UserMessageCount
}

func (c *ConversationSummaryResponseModel) GetAgentMessageCount() (value *int) {
	if c == nil {
		return
	}
	return c.AgentMessageCount
}

func (c *ConversationSummaryResponseModel) GetStatus() (value ConversationSummaryResponseModelStatus) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *ConversationSummaryResponseModel) GetTerminationReason() (value *string) {
	if c == nil {
		return
	}
	return c.TerminationReason
}

func (c *ConversationSummaryResponseModel) GetCallSuccessful() (value EvaluationSuccessResult) {
	if c == nil {
		return
	}
	return c.CallSuccessful
}

func (c *ConversationSummaryResponseModel) GetCallSuccessScore() (value *float64) {
	if c == nil {
		return
	}
	return c.CallSuccessScore
}

func (c *ConversationSummaryResponseModel) GetTranscriptSummary() (value *string) {
	if c == nil {
		return
	}
	return c.TranscriptSummary
}

func (c *ConversationSummaryResponseModel) GetCallSummaryTitle() (value *string) {
	if c == nil {
		return
	}
	return c.CallSummaryTitle
}

func (c *ConversationSummaryResponseModel) GetMainLanguage() (value *string) {
	if c == nil {
		return
	}
	return c.MainLanguage
}

func (c *ConversationSummaryResponseModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSource
}

func (c *ConversationSummaryResponseModel) GetToolNames() (value []string) {
	if c == nil {
		return
	}
	return c.ToolNames
}

func (c *ConversationSummaryResponseModel) GetDirection() (value *TelephonyDirection) {
	if c == nil {
		return
	}
	return c.Direction
}

func (c *ConversationSummaryResponseModel) GetRating() (value *float64) {
	if c == nil {
		return
	}
	return c.Rating
}

func (c *ConversationSummaryResponseModel) GetSentimentAnalysis() (value *ConversationSentimentAnalysis) {
	if c == nil {
		return
	}
	return c.SentimentAnalysis
}

func (c *ConversationSummaryResponseModel) GetDataCollectionResults() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DataCollectionResults
}

func (c *ConversationSummaryResponseModel) GetEvaluationCriteriaResults() (value map[string]*EvaluationCriteriaSummaryResult) {
	if c == nil {
		return
	}
	return c.EvaluationCriteriaResults
}

func (c *ConversationSummaryResponseModel) GetTagIds() (value []string) {
	if c == nil {
		return
	}
	return c.TagIds
}

func (c *ConversationSummaryResponseModel) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *ConversationSummaryResponseModel) SetBranchId(branchId *string) {
	c.BranchId = branchId
}

func (c *ConversationSummaryResponseModel) SetVersionId(versionId *string) {
	c.VersionId = versionId
}

func (c *ConversationSummaryResponseModel) SetAgentName(agentName *string) {
	c.AgentName = agentName
}

func (c *ConversationSummaryResponseModel) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *ConversationSummaryResponseModel) SetStartTimeUnixSecs(startTimeUnixSecs int) {
	c.StartTimeUnixSecs = startTimeUnixSecs
}

func (c *ConversationSummaryResponseModel) SetCallDurationSecs(callDurationSecs int) {
	c.CallDurationSecs = callDurationSecs
}

func (c *ConversationSummaryResponseModel) SetCost(cost *int) {
	c.Cost = cost
}

func (c *ConversationSummaryResponseModel) SetCostFiat(costFiat *float64) {
	c.CostFiat = costFiat
}

func (c *ConversationSummaryResponseModel) SetMessageCount(messageCount int) {
	c.MessageCount = messageCount
}

func (c *ConversationSummaryResponseModel) SetUserMessageCount(userMessageCount *int) {
	c.UserMessageCount = userMessageCount
}

func (c *ConversationSummaryResponseModel) SetAgentMessageCount(agentMessageCount *int) {
	c.AgentMessageCount = agentMessageCount
}

func (c *ConversationSummaryResponseModel) SetStatus(status ConversationSummaryResponseModelStatus) {
	c.Status = status
}

func (c *ConversationSummaryResponseModel) SetTerminationReason(terminationReason *string) {
	c.TerminationReason = terminationReason
}

func (c *ConversationSummaryResponseModel) SetCallSuccessful(callSuccessful EvaluationSuccessResult) {
	c.CallSuccessful = callSuccessful
}

func (c *ConversationSummaryResponseModel) SetCallSuccessScore(callSuccessScore *float64) {
	c.CallSuccessScore = callSuccessScore
}

func (c *ConversationSummaryResponseModel) SetTranscriptSummary(transcriptSummary *string) {
	c.TranscriptSummary = transcriptSummary
}

func (c *ConversationSummaryResponseModel) SetCallSummaryTitle(callSummaryTitle *string) {
	c.CallSummaryTitle = callSummaryTitle
}

func (c *ConversationSummaryResponseModel) SetMainLanguage(mainLanguage *string) {
	c.MainLanguage = mainLanguage
}

func (c *ConversationSummaryResponseModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	c.ConversationInitiationSource = conversationInitiationSource
}

func (c *ConversationSummaryResponseModel) SetToolNames(toolNames []string) {
	c.ToolNames = toolNames
}

func (c *ConversationSummaryResponseModel) SetDirection(direction *TelephonyDirection) {
	c.Direction = direction
}

func (c *ConversationSummaryResponseModel) SetRating(rating *float64) {
	c.Rating = rating
}

func (c *ConversationSummaryResponseModel) SetSentimentAnalysis(sentimentAnalysis *ConversationSentimentAnalysis) {
	c.SentimentAnalysis = sentimentAnalysis
}

func (c *ConversationSummaryResponseModel) SetDataCollectionResults(dataCollectionResults map[string]any) {
	c.DataCollectionResults = dataCollectionResults
}

func (c *ConversationSummaryResponseModel) SetEvaluationCriteriaResults(evaluationCriteriaResults map[string]*EvaluationCriteriaSummaryResult) {
	c.EvaluationCriteriaResults = evaluationCriteriaResults
}

func (c *ConversationSummaryResponseModel) SetTagIds(tagIds []string) {
	c.TagIds = tagIds
}

func (c *ConversationSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationSummaryResponseModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
