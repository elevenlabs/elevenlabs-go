package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Detailed response model for a batch call including all recipients.
type BatchCallDetailedResponse struct {
	Id                   string                     `json:"id" xml:"id" required:"true"`
	PhoneNumberId        *string                    `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	PhoneProvider        *TelephonyProvider         `json:"phone_provider" xml:"phone_provider" required:"true"`
	WhatsappParams       *BatchCallWhatsAppParams   `json:"whatsapp_params" xml:"whatsapp_params" required:"true"`
	Name                 string                     `json:"name" xml:"name" required:"true"`
	AgentId              string                     `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId             *string                    `json:"branch_id" xml:"branch_id" required:"true"`
	Environment          *string                    `json:"environment" xml:"environment" required:"true"`
	CreatedAtUnix        int                        `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	ScheduledTimeUnix    int                        `json:"scheduled_time_unix" xml:"scheduled_time_unix" required:"true"`
	Timezone             *string                    `json:"timezone" xml:"timezone" required:"true"`
	TotalCallsDispatched int                        `json:"total_calls_dispatched" xml:"total_calls_dispatched" required:"true"`
	TotalCallsScheduled  int                        `json:"total_calls_scheduled" xml:"total_calls_scheduled" required:"true"`
	TotalCallsFinished   int                        `json:"total_calls_finished" xml:"total_calls_finished" required:"true"`
	LastUpdatedAtUnix    int                        `json:"last_updated_at_unix" xml:"last_updated_at_unix" required:"true"`
	Status               BatchCallStatus            `json:"status" xml:"status" required:"true"`
	RetryCount           int                        `json:"retry_count" xml:"retry_count" required:"true"`
	TelephonyCallConfig  *TelephonyCallConfigOutput `json:"telephony_call_config" xml:"telephony_call_config" required:"true"`
	// Maximum number of simultaneous calls for this batch. When set, dispatch is governed by this limit rather than workspace/agent capacity percentages.
	TargetConcurrencyLimit *int                                  `json:"target_concurrency_limit" xml:"target_concurrency_limit" required:"true" min:"1"`
	AgentName              string                                `json:"agent_name" xml:"agent_name" required:"true"`
	BranchName             *string                               `json:"branch_name" xml:"branch_name" required:"true"`
	Recipients             []*OutboundCallRecipientResponseModel `json:"recipients" xml:"recipients" required:"true"`
	extraProperties        map[string]interface{}
}

func (b *BatchCallDetailedResponse) GetId() (value string) {
	if b == nil {
		return
	}
	return b.Id
}

func (b *BatchCallDetailedResponse) GetPhoneNumberId() (value *string) {
	if b == nil {
		return
	}
	return b.PhoneNumberId
}

func (b *BatchCallDetailedResponse) GetPhoneProvider() (value *TelephonyProvider) {
	if b == nil {
		return
	}
	return b.PhoneProvider
}

func (b *BatchCallDetailedResponse) GetWhatsappParams() (value *BatchCallWhatsAppParams) {
	if b == nil {
		return
	}
	return b.WhatsappParams
}

func (b *BatchCallDetailedResponse) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BatchCallDetailedResponse) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BatchCallDetailedResponse) GetBranchId() (value *string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BatchCallDetailedResponse) GetEnvironment() (value *string) {
	if b == nil {
		return
	}
	return b.Environment
}

func (b *BatchCallDetailedResponse) GetCreatedAtUnix() (value int) {
	if b == nil {
		return
	}
	return b.CreatedAtUnix
}

func (b *BatchCallDetailedResponse) GetScheduledTimeUnix() (value int) {
	if b == nil {
		return
	}
	return b.ScheduledTimeUnix
}

func (b *BatchCallDetailedResponse) GetTimezone() (value *string) {
	if b == nil {
		return
	}
	return b.Timezone
}

func (b *BatchCallDetailedResponse) GetTotalCallsDispatched() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsDispatched
}

func (b *BatchCallDetailedResponse) GetTotalCallsScheduled() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsScheduled
}

func (b *BatchCallDetailedResponse) GetTotalCallsFinished() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsFinished
}

func (b *BatchCallDetailedResponse) GetLastUpdatedAtUnix() (value int) {
	if b == nil {
		return
	}
	return b.LastUpdatedAtUnix
}

func (b *BatchCallDetailedResponse) GetStatus() (value BatchCallStatus) {
	if b == nil {
		return
	}
	return b.Status
}

func (b *BatchCallDetailedResponse) GetRetryCount() (value int) {
	if b == nil {
		return
	}
	return b.RetryCount
}

func (b *BatchCallDetailedResponse) GetTelephonyCallConfig() (value *TelephonyCallConfigOutput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BatchCallDetailedResponse) GetTargetConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.TargetConcurrencyLimit
}

func (b *BatchCallDetailedResponse) GetAgentName() (value string) {
	if b == nil {
		return
	}
	return b.AgentName
}

func (b *BatchCallDetailedResponse) GetBranchName() (value *string) {
	if b == nil {
		return
	}
	return b.BranchName
}

func (b *BatchCallDetailedResponse) GetRecipients() (value []*OutboundCallRecipientResponseModel) {
	if b == nil {
		return
	}
	return b.Recipients
}

func (b *BatchCallDetailedResponse) SetId(id string) {
	b.Id = id
}

func (b *BatchCallDetailedResponse) SetPhoneNumberId(phoneNumberId *string) {
	b.PhoneNumberId = phoneNumberId
}

func (b *BatchCallDetailedResponse) SetPhoneProvider(phoneProvider *TelephonyProvider) {
	b.PhoneProvider = phoneProvider
}

func (b *BatchCallDetailedResponse) SetWhatsappParams(whatsappParams *BatchCallWhatsAppParams) {
	b.WhatsappParams = whatsappParams
}

func (b *BatchCallDetailedResponse) SetName(name string) {
	b.Name = name
}

func (b *BatchCallDetailedResponse) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BatchCallDetailedResponse) SetBranchId(branchId *string) {
	b.BranchId = branchId
}

func (b *BatchCallDetailedResponse) SetEnvironment(environment *string) {
	b.Environment = environment
}

func (b *BatchCallDetailedResponse) SetCreatedAtUnix(createdAtUnix int) {
	b.CreatedAtUnix = createdAtUnix
}

func (b *BatchCallDetailedResponse) SetScheduledTimeUnix(scheduledTimeUnix int) {
	b.ScheduledTimeUnix = scheduledTimeUnix
}

func (b *BatchCallDetailedResponse) SetTimezone(timezone *string) {
	b.Timezone = timezone
}

func (b *BatchCallDetailedResponse) SetTotalCallsDispatched(totalCallsDispatched int) {
	b.TotalCallsDispatched = totalCallsDispatched
}

func (b *BatchCallDetailedResponse) SetTotalCallsScheduled(totalCallsScheduled int) {
	b.TotalCallsScheduled = totalCallsScheduled
}

func (b *BatchCallDetailedResponse) SetTotalCallsFinished(totalCallsFinished int) {
	b.TotalCallsFinished = totalCallsFinished
}

func (b *BatchCallDetailedResponse) SetLastUpdatedAtUnix(lastUpdatedAtUnix int) {
	b.LastUpdatedAtUnix = lastUpdatedAtUnix
}

func (b *BatchCallDetailedResponse) SetStatus(status BatchCallStatus) {
	b.Status = status
}

func (b *BatchCallDetailedResponse) SetRetryCount(retryCount int) {
	b.RetryCount = retryCount
}

func (b *BatchCallDetailedResponse) SetTelephonyCallConfig(telephonyCallConfig *TelephonyCallConfigOutput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BatchCallDetailedResponse) SetTargetConcurrencyLimit(targetConcurrencyLimit *int) {
	b.TargetConcurrencyLimit = targetConcurrencyLimit
}

func (b *BatchCallDetailedResponse) SetAgentName(agentName string) {
	b.AgentName = agentName
}

func (b *BatchCallDetailedResponse) SetBranchName(branchName *string) {
	b.BranchName = branchName
}

func (b *BatchCallDetailedResponse) SetRecipients(recipients []*OutboundCallRecipientResponseModel) {
	b.Recipients = recipients
}

func (b *BatchCallDetailedResponse) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchCallDetailedResponse) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchCallDetailedResponse to string"
	}
	return string(jsonData)
}

func (b *BatchCallDetailedResponse) UnmarshalJSON(data []byte) error {
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

func (b *BatchCallDetailedResponse) MarshalJSON() ([]byte, error) {
	type embed BatchCallDetailedResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
