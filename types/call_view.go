package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CallView struct {
	IntentId   string                   `json:"intent_id" xml:"intent_id" required:"true"`
	LeadId     string                   `json:"lead_id" xml:"lead_id" required:"true"`
	IntentType CampaignConversationType `json:"intent_type" xml:"intent_type" required:"true"`
	// `1` for the first call and `+1` per retry of it; unset for callbacks, redirects and their retries.
	ColdAttempt *int       `json:"cold_attempt" xml:"cold_attempt" required:"true"`
	Status      CallStatus `json:"status" xml:"status" required:"true"`
	PhoneNumber string     `json:"phone_number" xml:"phone_number" required:"true"`
	// When the call is due to be placed.
	ScheduledTimeUnix int     `json:"scheduled_time_unix" xml:"scheduled_time_unix" required:"true"`
	ConversationId    *string `json:"conversation_id" xml:"conversation_id" required:"true"`
	// The call's final dialing status, set once it ends.
	Outcome *BatchCallRecipientStatus `json:"outcome" xml:"outcome" required:"true"`
	// How the call ended.
	CallResult *CallResult `json:"call_result" xml:"call_result" required:"true"`
	// When the call ended.
	OutcomeAtUnix *int `json:"outcome_at_unix" xml:"outcome_at_unix" required:"true"`
	// What the campaign did after the call, such as queue a retry or a callback, or finish the lead.
	DecisionDisposition *DecisionDisposition `json:"decision_disposition" xml:"decision_disposition" required:"true"`
	// Data collection results this call copied into the lead's dynamic variables for later calls, with secrets redacted.
	CollectedDynamicVariables map[string]any `json:"collected_dynamic_variables,omitempty"`
	// The message's delivery; set on message campaigns.
	Delivery        *LeadDeliveryView `json:"delivery,omitempty" xml:"delivery,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CallView) GetIntentId() (value string) {
	if c == nil {
		return
	}
	return c.IntentId
}

func (c *CallView) GetLeadId() (value string) {
	if c == nil {
		return
	}
	return c.LeadId
}

func (c *CallView) GetIntentType() (value CampaignConversationType) {
	if c == nil {
		return
	}
	return c.IntentType
}

func (c *CallView) GetColdAttempt() (value *int) {
	if c == nil {
		return
	}
	return c.ColdAttempt
}

func (c *CallView) GetStatus() (value CallStatus) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *CallView) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

func (c *CallView) GetScheduledTimeUnix() (value int) {
	if c == nil {
		return
	}
	return c.ScheduledTimeUnix
}

func (c *CallView) GetConversationId() (value *string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *CallView) GetOutcome() (value *BatchCallRecipientStatus) {
	if c == nil {
		return
	}
	return c.Outcome
}

func (c *CallView) GetCallResult() (value *CallResult) {
	if c == nil {
		return
	}
	return c.CallResult
}

func (c *CallView) GetOutcomeAtUnix() (value *int) {
	if c == nil {
		return
	}
	return c.OutcomeAtUnix
}

func (c *CallView) GetDecisionDisposition() (value *DecisionDisposition) {
	if c == nil {
		return
	}
	return c.DecisionDisposition
}

func (c *CallView) GetCollectedDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.CollectedDynamicVariables
}

func (c *CallView) GetDelivery() (value *LeadDeliveryView) {
	if c == nil {
		return
	}
	return c.Delivery
}

func (c *CallView) SetIntentId(intentId string) {
	c.IntentId = intentId
}

func (c *CallView) SetLeadId(leadId string) {
	c.LeadId = leadId
}

func (c *CallView) SetIntentType(intentType CampaignConversationType) {
	c.IntentType = intentType
}

func (c *CallView) SetColdAttempt(coldAttempt *int) {
	c.ColdAttempt = coldAttempt
}

func (c *CallView) SetStatus(status CallStatus) {
	c.Status = status
}

func (c *CallView) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *CallView) SetScheduledTimeUnix(scheduledTimeUnix int) {
	c.ScheduledTimeUnix = scheduledTimeUnix
}

func (c *CallView) SetConversationId(conversationId *string) {
	c.ConversationId = conversationId
}

func (c *CallView) SetOutcome(outcome *BatchCallRecipientStatus) {
	c.Outcome = outcome
}

func (c *CallView) SetCallResult(callResult *CallResult) {
	c.CallResult = callResult
}

func (c *CallView) SetOutcomeAtUnix(outcomeAtUnix *int) {
	c.OutcomeAtUnix = outcomeAtUnix
}

func (c *CallView) SetDecisionDisposition(decisionDisposition *DecisionDisposition) {
	c.DecisionDisposition = decisionDisposition
}

func (c *CallView) SetCollectedDynamicVariables(collectedDynamicVariables map[string]any) {
	c.CollectedDynamicVariables = collectedDynamicVariables
}

func (c *CallView) SetDelivery(delivery *LeadDeliveryView) {
	c.Delivery = delivery
}

func (c *CallView) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CallView) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CallView to string"
	}
	return string(jsonData)
}

func (c *CallView) UnmarshalJSON(data []byte) error {
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

func (c *CallView) MarshalJSON() ([]byte, error) {
	type embed CallView
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
