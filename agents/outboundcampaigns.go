package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type SubmitCampaignRequest struct {
	// The campaign's name.
	CallName string `json:"call_name" xml:"call_name" required:"true" maxLength:"256" minLength:"1"`
	// The agent that places the calls.
	AgentId string `json:"agent_id" xml:"agent_id" required:"true" maxLength:"256" minLength:"1"`
	// Who to call. Each recipient becomes a lead of the campaign.
	Recipients []*types.SubmitCampaignRecipient `json:"recipients" xml:"recipients" required:"true" minItems:"1" maxItems:"10000"`
	// When to start calling; omit to start now. Calls still wait for the next calling window.
	ScheduledTimeUnix *int `json:"scheduled_time_unix,omitempty" xml:"scheduled_time_unix,omitempty" min:"0"`
	// IANA timezone to read the calling windows in, overriding the policy's.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The agent branch to call with.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty" maxLength:"256"`
	// The environment to resolve the agent's environment variables in.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty" maxLength:"256"`
	// The most calls to place at once.
	TargetConcurrencyLimit *int `json:"target_concurrency_limit,omitempty" xml:"target_concurrency_limit,omitempty" min:"1"`
	// Makes the request safe to retry: a repeat with the same key returns the campaign it created.
	IdempotencyKey *string `json:"idempotency_key,omitempty" xml:"idempotency_key,omitempty" maxLength:"256" minLength:"1"`
	// Policy to snapshot onto the campaign. When set, it wins over the workspace default and any per-agent override; when omitted, the workspace configuration must exist and supplies the policy.
	Policy *types.PolicyConfigInput `json:"policy,omitempty" xml:"policy,omitempty"`
	// `simple` runs one attempt per recipient with always-on calling windows and nothing collected, and cannot be combined with `policy`. `advanced` uses `policy` or the workspace configuration, and is the default for phone campaigns.
	Variant       *types.CampaignVariant              `json:"variant,omitempty" xml:"variant,omitempty"`
	ChannelConfig *SubmitCampaignRequestChannelConfig `json:"channel_config" xml:"channel_config" required:"true"`
}

func (s *SubmitCampaignRequest) GetCallName() (value string) {
	if s == nil {
		return
	}
	return s.CallName
}

func (s *SubmitCampaignRequest) GetAgentId() (value string) {
	if s == nil {
		return
	}
	return s.AgentId
}

func (s *SubmitCampaignRequest) GetRecipients() (value []*types.SubmitCampaignRecipient) {
	if s == nil {
		return
	}
	return s.Recipients
}

func (s *SubmitCampaignRequest) GetScheduledTimeUnix() (value *int) {
	if s == nil {
		return
	}
	return s.ScheduledTimeUnix
}

func (s *SubmitCampaignRequest) GetTimezone() (value *string) {
	if s == nil {
		return
	}
	return s.Timezone
}

func (s *SubmitCampaignRequest) GetBranchId() (value *string) {
	if s == nil {
		return
	}
	return s.BranchId
}

func (s *SubmitCampaignRequest) GetEnvironment() (value *string) {
	if s == nil {
		return
	}
	return s.Environment
}

func (s *SubmitCampaignRequest) GetTargetConcurrencyLimit() (value *int) {
	if s == nil {
		return
	}
	return s.TargetConcurrencyLimit
}

func (s *SubmitCampaignRequest) GetIdempotencyKey() (value *string) {
	if s == nil {
		return
	}
	return s.IdempotencyKey
}

func (s *SubmitCampaignRequest) GetPolicy() (value *types.PolicyConfigInput) {
	if s == nil {
		return
	}
	return s.Policy
}

func (s *SubmitCampaignRequest) GetVariant() (value *types.CampaignVariant) {
	if s == nil {
		return
	}
	return s.Variant
}

func (s *SubmitCampaignRequest) GetChannelConfig() (value *SubmitCampaignRequestChannelConfig) {
	if s == nil {
		return
	}
	return s.ChannelConfig
}

func (s *SubmitCampaignRequest) SetCallName(callName string) {
	s.CallName = callName
}

func (s *SubmitCampaignRequest) SetAgentId(agentId string) {
	s.AgentId = agentId
}

func (s *SubmitCampaignRequest) SetRecipients(recipients []*types.SubmitCampaignRecipient) {
	s.Recipients = recipients
}

func (s *SubmitCampaignRequest) SetScheduledTimeUnix(scheduledTimeUnix *int) {
	s.ScheduledTimeUnix = scheduledTimeUnix
}

func (s *SubmitCampaignRequest) SetTimezone(timezone *string) {
	s.Timezone = timezone
}

func (s *SubmitCampaignRequest) SetBranchId(branchId *string) {
	s.BranchId = branchId
}

func (s *SubmitCampaignRequest) SetEnvironment(environment *string) {
	s.Environment = environment
}

func (s *SubmitCampaignRequest) SetTargetConcurrencyLimit(targetConcurrencyLimit *int) {
	s.TargetConcurrencyLimit = targetConcurrencyLimit
}

func (s *SubmitCampaignRequest) SetIdempotencyKey(idempotencyKey *string) {
	s.IdempotencyKey = idempotencyKey
}

func (s *SubmitCampaignRequest) SetPolicy(policy *types.PolicyConfigInput) {
	s.Policy = policy
}

func (s *SubmitCampaignRequest) SetVariant(variant *types.CampaignVariant) {
	s.Variant = variant
}

func (s *SubmitCampaignRequest) SetChannelConfig(channelConfig *SubmitCampaignRequestChannelConfig) {
	s.ChannelConfig = channelConfig
}

func (s *SubmitCampaignRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (s SubmitCampaignRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitCampaignRequest to string"
	}
	return string(jsonData)
}

func (s *SubmitCampaignRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

func (s *SubmitCampaignRequest) MarshalJSON() ([]byte, error) {
	type embed SubmitCampaignRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

type SubmitCampaignRequestChannelConfig struct {
	Channel  string
	Phone    *types.PhoneChannelConfigInput
	Whatsapp *types.SubmitWhatsAppChannelConfig
}

func (s *SubmitCampaignRequestChannelConfig) GetPhone() *types.PhoneChannelConfigInput {
	if s == nil {
		return nil
	}
	if s.Phone == nil {
		return nil
	}
	return s.Phone
}

func (s *SubmitCampaignRequestChannelConfig) GetWhatsapp() *types.SubmitWhatsAppChannelConfig {
	if s == nil {
		return nil
	}
	if s.Whatsapp == nil {
		return nil
	}
	return s.Whatsapp
}

func (s *SubmitCampaignRequestChannelConfig) GetChannel() string {
	if s == nil {
		return ""
	}
	return s.Channel
}

func (s *SubmitCampaignRequestChannelConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Channel string `json:"channel"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Channel = unmarshaler.Channel
	if unmarshaler.Channel == "" {
		return fmt.Errorf("%T did not include discriminant channel", s)
	}
	switch unmarshaler.Channel {
	case "phone":
		value := new(types.PhoneChannelConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Phone = value
	case "whatsapp":
		value := new(types.SubmitWhatsAppChannelConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Whatsapp = value
	}
	return nil
}

func (s SubmitCampaignRequestChannelConfig) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(s.Phone, "channel", "phone")
	}
	if s.Whatsapp != nil {
		return internal.MarshalJSONWithDiscriminant(s.Whatsapp, "channel", "whatsapp")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SubmitCampaignRequestChannelConfigVisitor interface {
	VisitPhone(*types.PhoneChannelConfigInput) error
	VisitWhatsapp(*types.SubmitWhatsAppChannelConfig) error
}

func (s *SubmitCampaignRequestChannelConfig) Accept(visitor SubmitCampaignRequestChannelConfigVisitor) error {
	if s.Phone != nil {
		return visitor.VisitPhone(s.Phone)
	}
	if s.Whatsapp != nil {
		return visitor.VisitWhatsapp(s.Whatsapp)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SubmitCampaignRequestChannelConfig) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Phone != nil {
		fields = append(fields, "phone")
	}
	if s.Whatsapp != nil {
		fields = append(fields, "whatsapp")
	}
	if len(fields) == 0 {
		if s.Channel != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Channel)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Channel != "" {
		field := fields[0]
		if s.Channel != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Channel,
				s,
			)
		}
	}
	return nil
}

type ListOutboundCampaignsRequestMode string

const (
	ListOutboundCampaignsRequestModeVoice   ListOutboundCampaignsRequestMode = "voice"
	ListOutboundCampaignsRequestModeMessage ListOutboundCampaignsRequestMode = "message"
)

func NewListOutboundCampaignsRequestModeFromString(s string) (ListOutboundCampaignsRequestMode, error) {
	switch s {
	case "voice":
		return ListOutboundCampaignsRequestModeVoice, nil
	case "message":
		return ListOutboundCampaignsRequestModeMessage, nil
	}
	var t ListOutboundCampaignsRequestMode
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListOutboundCampaignsRequestMode) Ptr() *ListOutboundCampaignsRequestMode {
	return &e
}

var (
	listOutboundCampaignsRequestFieldPageSize = big.NewInt(1 << 0)
	listOutboundCampaignsRequestFieldStatus   = big.NewInt(1 << 1)
	listOutboundCampaignsRequestFieldAgentId  = big.NewInt(1 << 2)
	listOutboundCampaignsRequestFieldChannel  = big.NewInt(1 << 3)
	listOutboundCampaignsRequestFieldMode     = big.NewInt(1 << 4)
	listOutboundCampaignsRequestFieldSearch   = big.NewInt(1 << 5)
	listOutboundCampaignsRequestFieldCursor   = big.NewInt(1 << 6)
)

type ListOutboundCampaignsRequest struct {
	PageSize *int                              `json:"-" url:"page_size,omitempty"`
	Status   []*types.CampaignStatus           `json:"-" url:"status,omitempty"`
	AgentId  *string                           `json:"-" url:"agent_id,omitempty"`
	Channel  *types.OutboundChannel            `json:"-" url:"channel,omitempty"`
	Mode     *ListOutboundCampaignsRequestMode `json:"-" url:"mode,omitempty"`
	Search   *string                           `json:"-" url:"search,omitempty"`
	Cursor   *string                           `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListOutboundCampaignsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listOutboundCampaignsRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetStatus(status []*types.CampaignStatus) {
	l.Status = status
	l.require(listOutboundCampaignsRequestFieldStatus)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listOutboundCampaignsRequestFieldAgentId)
}

// SetChannel sets the Channel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetChannel(channel *types.OutboundChannel) {
	l.Channel = channel
	l.require(listOutboundCampaignsRequestFieldChannel)
}

// SetMode sets the Mode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetMode(mode *ListOutboundCampaignsRequestMode) {
	l.Mode = mode
	l.require(listOutboundCampaignsRequestFieldMode)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listOutboundCampaignsRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOutboundCampaignsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listOutboundCampaignsRequestFieldCursor)
}
