package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BatchCallResponse struct {
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
	TargetConcurrencyLimit *int    `json:"target_concurrency_limit" xml:"target_concurrency_limit" required:"true" min:"1"`
	AgentName              string  `json:"agent_name" xml:"agent_name" required:"true"`
	BranchName             *string `json:"branch_name" xml:"branch_name" required:"true"`
	extraProperties        map[string]interface{}
}

func (b *BatchCallResponse) GetId() (value string) {
	if b == nil {
		return
	}
	return b.Id
}

func (b *BatchCallResponse) GetPhoneNumberId() (value *string) {
	if b == nil {
		return
	}
	return b.PhoneNumberId
}

func (b *BatchCallResponse) GetPhoneProvider() (value *TelephonyProvider) {
	if b == nil {
		return
	}
	return b.PhoneProvider
}

func (b *BatchCallResponse) GetWhatsappParams() (value *BatchCallWhatsAppParams) {
	if b == nil {
		return
	}
	return b.WhatsappParams
}

func (b *BatchCallResponse) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BatchCallResponse) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BatchCallResponse) GetBranchId() (value *string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BatchCallResponse) GetEnvironment() (value *string) {
	if b == nil {
		return
	}
	return b.Environment
}

func (b *BatchCallResponse) GetCreatedAtUnix() (value int) {
	if b == nil {
		return
	}
	return b.CreatedAtUnix
}

func (b *BatchCallResponse) GetScheduledTimeUnix() (value int) {
	if b == nil {
		return
	}
	return b.ScheduledTimeUnix
}

func (b *BatchCallResponse) GetTimezone() (value *string) {
	if b == nil {
		return
	}
	return b.Timezone
}

func (b *BatchCallResponse) GetTotalCallsDispatched() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsDispatched
}

func (b *BatchCallResponse) GetTotalCallsScheduled() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsScheduled
}

func (b *BatchCallResponse) GetTotalCallsFinished() (value int) {
	if b == nil {
		return
	}
	return b.TotalCallsFinished
}

func (b *BatchCallResponse) GetLastUpdatedAtUnix() (value int) {
	if b == nil {
		return
	}
	return b.LastUpdatedAtUnix
}

func (b *BatchCallResponse) GetStatus() (value BatchCallStatus) {
	if b == nil {
		return
	}
	return b.Status
}

func (b *BatchCallResponse) GetRetryCount() (value int) {
	if b == nil {
		return
	}
	return b.RetryCount
}

func (b *BatchCallResponse) GetTelephonyCallConfig() (value *TelephonyCallConfigOutput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BatchCallResponse) GetTargetConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.TargetConcurrencyLimit
}

func (b *BatchCallResponse) GetAgentName() (value string) {
	if b == nil {
		return
	}
	return b.AgentName
}

func (b *BatchCallResponse) GetBranchName() (value *string) {
	if b == nil {
		return
	}
	return b.BranchName
}

func (b *BatchCallResponse) SetId(id string) {
	b.Id = id
}

func (b *BatchCallResponse) SetPhoneNumberId(phoneNumberId *string) {
	b.PhoneNumberId = phoneNumberId
}

func (b *BatchCallResponse) SetPhoneProvider(phoneProvider *TelephonyProvider) {
	b.PhoneProvider = phoneProvider
}

func (b *BatchCallResponse) SetWhatsappParams(whatsappParams *BatchCallWhatsAppParams) {
	b.WhatsappParams = whatsappParams
}

func (b *BatchCallResponse) SetName(name string) {
	b.Name = name
}

func (b *BatchCallResponse) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BatchCallResponse) SetBranchId(branchId *string) {
	b.BranchId = branchId
}

func (b *BatchCallResponse) SetEnvironment(environment *string) {
	b.Environment = environment
}

func (b *BatchCallResponse) SetCreatedAtUnix(createdAtUnix int) {
	b.CreatedAtUnix = createdAtUnix
}

func (b *BatchCallResponse) SetScheduledTimeUnix(scheduledTimeUnix int) {
	b.ScheduledTimeUnix = scheduledTimeUnix
}

func (b *BatchCallResponse) SetTimezone(timezone *string) {
	b.Timezone = timezone
}

func (b *BatchCallResponse) SetTotalCallsDispatched(totalCallsDispatched int) {
	b.TotalCallsDispatched = totalCallsDispatched
}

func (b *BatchCallResponse) SetTotalCallsScheduled(totalCallsScheduled int) {
	b.TotalCallsScheduled = totalCallsScheduled
}

func (b *BatchCallResponse) SetTotalCallsFinished(totalCallsFinished int) {
	b.TotalCallsFinished = totalCallsFinished
}

func (b *BatchCallResponse) SetLastUpdatedAtUnix(lastUpdatedAtUnix int) {
	b.LastUpdatedAtUnix = lastUpdatedAtUnix
}

func (b *BatchCallResponse) SetStatus(status BatchCallStatus) {
	b.Status = status
}

func (b *BatchCallResponse) SetRetryCount(retryCount int) {
	b.RetryCount = retryCount
}

func (b *BatchCallResponse) SetTelephonyCallConfig(telephonyCallConfig *TelephonyCallConfigOutput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BatchCallResponse) SetTargetConcurrencyLimit(targetConcurrencyLimit *int) {
	b.TargetConcurrencyLimit = targetConcurrencyLimit
}

func (b *BatchCallResponse) SetAgentName(agentName string) {
	b.AgentName = agentName
}

func (b *BatchCallResponse) SetBranchName(branchName *string) {
	b.BranchName = branchName
}

func (b *BatchCallResponse) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchCallResponse) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchCallResponse to string"
	}
	return string(jsonData)
}

func (b *BatchCallResponse) UnmarshalJSON(data []byte) error {
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

func (b *BatchCallResponse) MarshalJSON() ([]byte, error) {
	type embed BatchCallResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
