package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryMetadataCommonModel struct {
	StartTimeUnixSecs    int  `json:"start_time_unix_secs" xml:"start_time_unix_secs" required:"true"`
	AcceptedTimeUnixSecs *int `json:"accepted_time_unix_secs,omitempty" xml:"accepted_time_unix_secs,omitempty"`
	CallDurationSecs     int  `json:"call_duration_secs" xml:"call_duration_secs" required:"true"`
	// Seconds the caller was held in the concurrency wait queue. Excluded from call_duration_secs and from billed time. None when the conversation was never queued.
	QueueWaitSecs       *float64                                         `json:"queue_wait_secs,omitempty" xml:"queue_wait_secs,omitempty"`
	Cost                *int                                             `json:"cost,omitempty" xml:"cost,omitempty"`
	DeletionSettings    *ConversationDeletionSettings                    `json:"deletion_settings,omitempty" xml:"deletion_settings,omitempty"`
	Feedback            *ConversationHistoryFeedbackCommonModel          `json:"feedback,omitempty" xml:"feedback,omitempty"`
	AuthorizationMethod *AuthorizationMethod                             `json:"authorization_method,omitempty" xml:"authorization_method,omitempty"`
	Charging            *ConversationChargingCommonModel                 `json:"charging,omitempty" xml:"charging,omitempty"`
	PhoneCall           *ConversationHistoryMetadataCommonModelPhoneCall `json:"phone_call,omitempty" xml:"phone_call,omitempty"`
	BatchCall           *ConversationHistoryBatchCallModel               `json:"batch_call,omitempty" xml:"batch_call,omitempty"`
	TerminationReason   *string                                          `json:"termination_reason,omitempty" xml:"termination_reason,omitempty"`
	Error               *ConversationHistoryErrorCommonModel             `json:"error,omitempty" xml:"error,omitempty"`
	Warnings            []string                                         `json:"warnings,omitempty" xml:"warnings,omitempty"`
	MainLanguage        *string                                          `json:"main_language,omitempty" xml:"main_language,omitempty"`
	RagUsage            *ConversationHistoryRagUsageCommonModel          `json:"rag_usage,omitempty" xml:"rag_usage,omitempty"`
	TextOnly            *bool                                            `json:"text_only,omitempty" xml:"text_only,omitempty"`
	FeaturesUsage       *FeaturesUsageCommonModel                        `json:"features_usage,omitempty" xml:"features_usage,omitempty"`
	ElevenAssistant     *ConversationHistoryElevenAssistantCommonModel   `json:"eleven_assistant,omitempty" xml:"eleven_assistant,omitempty"`
	InitiatorId         *string                                          `json:"initiator_id,omitempty" xml:"initiator_id,omitempty"`
	// Enum representing the possible sources for conversation initiation.
	ConversationInitiationSource        *ConversationInitiationSource   `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	ConversationInitiationSourceVersion *string                         `json:"conversation_initiation_source_version,omitempty" xml:"conversation_initiation_source_version,omitempty" maxLength:"50"`
	Timezone                            *string                         `json:"timezone,omitempty" xml:"timezone,omitempty"`
	AsyncMetadata                       *AsyncConversationMetadata      `json:"async_metadata,omitempty" xml:"async_metadata,omitempty"`
	Whatsapp                            *WhatsAppConversationInfo       `json:"whatsapp,omitempty" xml:"whatsapp,omitempty"`
	Sms                                 *SmsConversationInfo            `json:"sms,omitempty" xml:"sms,omitempty"`
	AgentCreatedFrom                    *AgentDefinitionSource          `json:"agent_created_from,omitempty" xml:"agent_created_from,omitempty"`
	AgentLastUpdatedFrom                *AgentDefinitionSource          `json:"agent_last_updated_from,omitempty" xml:"agent_last_updated_from,omitempty"`
	VoiceRewards                        []*ConversationVoiceRewardModel `json:"voice_rewards,omitempty" xml:"voice_rewards,omitempty"`
	// Total fiat cost of the conversation in USD, i.e. the sum of the LLM price and the non-LLM platform price (the fiat analogue of ``cost``). ``None`` when neither is set (e.g. conversations that predate fiat cost tracking).
	CostFiat        *float64 `json:"cost_fiat,omitempty" xml:"cost_fiat,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryMetadataCommonModel) GetStartTimeUnixSecs() (value int) {
	if c == nil {
		return
	}
	return c.StartTimeUnixSecs
}

func (c *ConversationHistoryMetadataCommonModel) GetAcceptedTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.AcceptedTimeUnixSecs
}

func (c *ConversationHistoryMetadataCommonModel) GetCallDurationSecs() (value int) {
	if c == nil {
		return
	}
	return c.CallDurationSecs
}

func (c *ConversationHistoryMetadataCommonModel) GetQueueWaitSecs() (value *float64) {
	if c == nil {
		return
	}
	return c.QueueWaitSecs
}

func (c *ConversationHistoryMetadataCommonModel) GetCost() (value *int) {
	if c == nil {
		return
	}
	return c.Cost
}

func (c *ConversationHistoryMetadataCommonModel) GetDeletionSettings() (value *ConversationDeletionSettings) {
	if c == nil {
		return
	}
	return c.DeletionSettings
}

func (c *ConversationHistoryMetadataCommonModel) GetFeedback() (value *ConversationHistoryFeedbackCommonModel) {
	if c == nil {
		return
	}
	return c.Feedback
}

func (c *ConversationHistoryMetadataCommonModel) GetAuthorizationMethod() (value *AuthorizationMethod) {
	if c == nil {
		return
	}
	return c.AuthorizationMethod
}

func (c *ConversationHistoryMetadataCommonModel) GetCharging() (value *ConversationChargingCommonModel) {
	if c == nil {
		return
	}
	return c.Charging
}

func (c *ConversationHistoryMetadataCommonModel) GetPhoneCall() (value *ConversationHistoryMetadataCommonModelPhoneCall) {
	if c == nil {
		return
	}
	return c.PhoneCall
}

func (c *ConversationHistoryMetadataCommonModel) GetBatchCall() (value *ConversationHistoryBatchCallModel) {
	if c == nil {
		return
	}
	return c.BatchCall
}

func (c *ConversationHistoryMetadataCommonModel) GetTerminationReason() (value *string) {
	if c == nil {
		return
	}
	return c.TerminationReason
}

func (c *ConversationHistoryMetadataCommonModel) GetError() (value *ConversationHistoryErrorCommonModel) {
	if c == nil {
		return
	}
	return c.Error
}

func (c *ConversationHistoryMetadataCommonModel) GetWarnings() (value []string) {
	if c == nil {
		return
	}
	return c.Warnings
}

func (c *ConversationHistoryMetadataCommonModel) GetMainLanguage() (value *string) {
	if c == nil {
		return
	}
	return c.MainLanguage
}

func (c *ConversationHistoryMetadataCommonModel) GetRagUsage() (value *ConversationHistoryRagUsageCommonModel) {
	if c == nil {
		return
	}
	return c.RagUsage
}

func (c *ConversationHistoryMetadataCommonModel) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationHistoryMetadataCommonModel) GetFeaturesUsage() (value *FeaturesUsageCommonModel) {
	if c == nil {
		return
	}
	return c.FeaturesUsage
}

func (c *ConversationHistoryMetadataCommonModel) GetElevenAssistant() (value *ConversationHistoryElevenAssistantCommonModel) {
	if c == nil {
		return
	}
	return c.ElevenAssistant
}

func (c *ConversationHistoryMetadataCommonModel) GetInitiatorId() (value *string) {
	if c == nil {
		return
	}
	return c.InitiatorId
}

func (c *ConversationHistoryMetadataCommonModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSource
}

func (c *ConversationHistoryMetadataCommonModel) GetConversationInitiationSourceVersion() (value *string) {
	if c == nil {
		return
	}
	return c.ConversationInitiationSourceVersion
}

func (c *ConversationHistoryMetadataCommonModel) GetTimezone() (value *string) {
	if c == nil {
		return
	}
	return c.Timezone
}

func (c *ConversationHistoryMetadataCommonModel) GetAsyncMetadata() (value *AsyncConversationMetadata) {
	if c == nil {
		return
	}
	return c.AsyncMetadata
}

func (c *ConversationHistoryMetadataCommonModel) GetWhatsapp() (value *WhatsAppConversationInfo) {
	if c == nil {
		return
	}
	return c.Whatsapp
}

func (c *ConversationHistoryMetadataCommonModel) GetSms() (value *SmsConversationInfo) {
	if c == nil {
		return
	}
	return c.Sms
}

func (c *ConversationHistoryMetadataCommonModel) GetAgentCreatedFrom() (value *AgentDefinitionSource) {
	if c == nil {
		return
	}
	return c.AgentCreatedFrom
}

func (c *ConversationHistoryMetadataCommonModel) GetAgentLastUpdatedFrom() (value *AgentDefinitionSource) {
	if c == nil {
		return
	}
	return c.AgentLastUpdatedFrom
}

func (c *ConversationHistoryMetadataCommonModel) GetVoiceRewards() (value []*ConversationVoiceRewardModel) {
	if c == nil {
		return
	}
	return c.VoiceRewards
}

func (c *ConversationHistoryMetadataCommonModel) GetCostFiat() (value *float64) {
	if c == nil {
		return
	}
	return c.CostFiat
}

func (c *ConversationHistoryMetadataCommonModel) SetStartTimeUnixSecs(startTimeUnixSecs int) {
	c.StartTimeUnixSecs = startTimeUnixSecs
}

func (c *ConversationHistoryMetadataCommonModel) SetAcceptedTimeUnixSecs(acceptedTimeUnixSecs *int) {
	c.AcceptedTimeUnixSecs = acceptedTimeUnixSecs
}

func (c *ConversationHistoryMetadataCommonModel) SetCallDurationSecs(callDurationSecs int) {
	c.CallDurationSecs = callDurationSecs
}

func (c *ConversationHistoryMetadataCommonModel) SetQueueWaitSecs(queueWaitSecs *float64) {
	c.QueueWaitSecs = queueWaitSecs
}

func (c *ConversationHistoryMetadataCommonModel) SetCost(cost *int) {
	c.Cost = cost
}

func (c *ConversationHistoryMetadataCommonModel) SetDeletionSettings(deletionSettings *ConversationDeletionSettings) {
	c.DeletionSettings = deletionSettings
}

func (c *ConversationHistoryMetadataCommonModel) SetFeedback(feedback *ConversationHistoryFeedbackCommonModel) {
	c.Feedback = feedback
}

func (c *ConversationHistoryMetadataCommonModel) SetAuthorizationMethod(authorizationMethod *AuthorizationMethod) {
	c.AuthorizationMethod = authorizationMethod
}

func (c *ConversationHistoryMetadataCommonModel) SetCharging(charging *ConversationChargingCommonModel) {
	c.Charging = charging
}

func (c *ConversationHistoryMetadataCommonModel) SetPhoneCall(phoneCall *ConversationHistoryMetadataCommonModelPhoneCall) {
	c.PhoneCall = phoneCall
}

func (c *ConversationHistoryMetadataCommonModel) SetBatchCall(batchCall *ConversationHistoryBatchCallModel) {
	c.BatchCall = batchCall
}

func (c *ConversationHistoryMetadataCommonModel) SetTerminationReason(terminationReason *string) {
	c.TerminationReason = terminationReason
}

func (c *ConversationHistoryMetadataCommonModel) SetError(error *ConversationHistoryErrorCommonModel) {
	c.Error = error
}

func (c *ConversationHistoryMetadataCommonModel) SetWarnings(warnings []string) {
	c.Warnings = warnings
}

func (c *ConversationHistoryMetadataCommonModel) SetMainLanguage(mainLanguage *string) {
	c.MainLanguage = mainLanguage
}

func (c *ConversationHistoryMetadataCommonModel) SetRagUsage(ragUsage *ConversationHistoryRagUsageCommonModel) {
	c.RagUsage = ragUsage
}

func (c *ConversationHistoryMetadataCommonModel) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationHistoryMetadataCommonModel) SetFeaturesUsage(featuresUsage *FeaturesUsageCommonModel) {
	c.FeaturesUsage = featuresUsage
}

func (c *ConversationHistoryMetadataCommonModel) SetElevenAssistant(elevenAssistant *ConversationHistoryElevenAssistantCommonModel) {
	c.ElevenAssistant = elevenAssistant
}

func (c *ConversationHistoryMetadataCommonModel) SetInitiatorId(initiatorId *string) {
	c.InitiatorId = initiatorId
}

func (c *ConversationHistoryMetadataCommonModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	c.ConversationInitiationSource = conversationInitiationSource
}

func (c *ConversationHistoryMetadataCommonModel) SetConversationInitiationSourceVersion(conversationInitiationSourceVersion *string) {
	c.ConversationInitiationSourceVersion = conversationInitiationSourceVersion
}

func (c *ConversationHistoryMetadataCommonModel) SetTimezone(timezone *string) {
	c.Timezone = timezone
}

func (c *ConversationHistoryMetadataCommonModel) SetAsyncMetadata(asyncMetadata *AsyncConversationMetadata) {
	c.AsyncMetadata = asyncMetadata
}

func (c *ConversationHistoryMetadataCommonModel) SetWhatsapp(whatsapp *WhatsAppConversationInfo) {
	c.Whatsapp = whatsapp
}

func (c *ConversationHistoryMetadataCommonModel) SetSms(sms *SmsConversationInfo) {
	c.Sms = sms
}

func (c *ConversationHistoryMetadataCommonModel) SetAgentCreatedFrom(agentCreatedFrom *AgentDefinitionSource) {
	c.AgentCreatedFrom = agentCreatedFrom
}

func (c *ConversationHistoryMetadataCommonModel) SetAgentLastUpdatedFrom(agentLastUpdatedFrom *AgentDefinitionSource) {
	c.AgentLastUpdatedFrom = agentLastUpdatedFrom
}

func (c *ConversationHistoryMetadataCommonModel) SetVoiceRewards(voiceRewards []*ConversationVoiceRewardModel) {
	c.VoiceRewards = voiceRewards
}

func (c *ConversationHistoryMetadataCommonModel) SetCostFiat(costFiat *float64) {
	c.CostFiat = costFiat
}

func (c *ConversationHistoryMetadataCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryMetadataCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryMetadataCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryMetadataCommonModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryMetadataCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryMetadataCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
