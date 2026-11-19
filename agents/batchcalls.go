package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost struct {
	CallName            string                          `json:"call_name" xml:"call_name" required:"true"`
	AgentId             string                          `json:"agent_id" xml:"agent_id" required:"true"`
	Recipients          []*types.OutboundCallRecipient  `json:"recipients" xml:"recipients" required:"true" maxItems:"10000"`
	ScheduledTimeUnix   *int                            `json:"scheduled_time_unix,omitempty" xml:"scheduled_time_unix,omitempty"`
	AgentPhoneNumberId  *string                         `json:"agent_phone_number_id,omitempty" xml:"agent_phone_number_id,omitempty"`
	WhatsappParams      *types.BatchCallWhatsAppParams  `json:"whatsapp_params,omitempty" xml:"whatsapp_params,omitempty"`
	Timezone            *string                         `json:"timezone,omitempty" xml:"timezone,omitempty"`
	BranchId            *string                         `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	Environment         *string                         `json:"environment,omitempty" xml:"environment,omitempty"`
	TelephonyCallConfig *types.TelephonyCallConfigInput `json:"telephony_call_config,omitempty" xml:"telephony_call_config,omitempty"`
	// Maximum number of simultaneous calls for this batch. When set, dispatch is governed by this limit rather than workspace/agent capacity percentages.
	TargetConcurrencyLimit *int `json:"target_concurrency_limit,omitempty" xml:"target_concurrency_limit,omitempty" min:"1"`
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetCallName() (value string) {
	if b == nil {
		return
	}
	return b.CallName
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetRecipients() (value []*types.OutboundCallRecipient) {
	if b == nil {
		return
	}
	return b.Recipients
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetScheduledTimeUnix() (value *int) {
	if b == nil {
		return
	}
	return b.ScheduledTimeUnix
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetAgentPhoneNumberId() (value *string) {
	if b == nil {
		return
	}
	return b.AgentPhoneNumberId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetWhatsappParams() (value *types.BatchCallWhatsAppParams) {
	if b == nil {
		return
	}
	return b.WhatsappParams
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetTimezone() (value *string) {
	if b == nil {
		return
	}
	return b.Timezone
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetBranchId() (value *string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetEnvironment() (value *string) {
	if b == nil {
		return
	}
	return b.Environment
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetTelephonyCallConfig() (value *types.TelephonyCallConfigInput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetTargetConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.TargetConcurrencyLimit
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetCallName(callName string) {
	b.CallName = callName
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetRecipients(recipients []*types.OutboundCallRecipient) {
	b.Recipients = recipients
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetScheduledTimeUnix(scheduledTimeUnix *int) {
	b.ScheduledTimeUnix = scheduledTimeUnix
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetAgentPhoneNumberId(agentPhoneNumberId *string) {
	b.AgentPhoneNumberId = agentPhoneNumberId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetWhatsappParams(whatsappParams *types.BatchCallWhatsAppParams) {
	b.WhatsappParams = whatsappParams
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetTimezone(timezone *string) {
	b.Timezone = timezone
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetBranchId(branchId *string) {
	b.BranchId = branchId
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetEnvironment(environment *string) {
	b.Environment = environment
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetTelephonyCallConfig(telephonyCallConfig *types.TelephonyCallConfigInput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) SetTargetConcurrencyLimit(targetConcurrencyLimit *int) {
	b.TargetConcurrencyLimit = targetConcurrencyLimit
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost to string"
	}
	return string(jsonData)
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost) MarshalJSON() ([]byte, error) {
	type embed BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listBatchCallsRequestFieldLimit   = big.NewInt(1 << 0)
	listBatchCallsRequestFieldLastDoc = big.NewInt(1 << 1)
	listBatchCallsRequestFieldAgentId = big.NewInt(1 << 2)
)

type ListBatchCallsRequest struct {
	Limit   *int    `json:"-" url:"limit,omitempty"`
	LastDoc *string `json:"-" url:"last_doc,omitempty"`
	AgentId *string `json:"-" url:"agent_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListBatchCallsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBatchCallsRequest) SetLimit(limit *int) {
	l.Limit = limit
	l.require(listBatchCallsRequestFieldLimit)
}

// SetLastDoc sets the LastDoc field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBatchCallsRequest) SetLastDoc(lastDoc *string) {
	l.LastDoc = lastDoc
	l.require(listBatchCallsRequestFieldLastDoc)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBatchCallsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listBatchCallsRequestFieldAgentId)
}

var (
	exportBatchCallsRequestFieldLimit = big.NewInt(1 << 0)
)

type ExportBatchCallsRequest struct {
	Limit *int `json:"-" url:"limit,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (e *ExportBatchCallsRequest) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ExportBatchCallsRequest) SetLimit(limit *int) {
	e.Limit = limit
	e.require(exportBatchCallsRequestFieldLimit)
}
