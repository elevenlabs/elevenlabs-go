package receptionists

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

// Request model for creating a new transfer rule.
type ReceptionCreateTransferRuleRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId     string      `json:"-" url:"-"`
	Label       StrippedStr `json:"label" xml:"label" required:"true"`
	PhoneNumber PhoneNumber `json:"phone_number" xml:"phone_number" required:"true"`
	Condition   StrippedStr `json:"condition" xml:"condition" required:"true"`
	// DTMF digits automatically played after the transfer destination picks up (e.g. '1' to answer a 'press 1 to accept' screen, or 'www1234' to key an extension after a pause). Valid characters are 0-9, *, #, w (0.5s pause) and W (1s pause). Twilio transfers only.
	PostDialDigits *string `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty" maxLength:"50" minLength:"1" pattern:"^[0-9*#wW]*$"`
}

func (r *ReceptionCreateTransferRuleRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateTransferRuleRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionCreateTransferRuleRequest) GetLabel() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Label
}

func (r *ReceptionCreateTransferRuleRequest) GetPhoneNumber() (value PhoneNumber) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

func (r *ReceptionCreateTransferRuleRequest) GetCondition() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Condition
}

func (r *ReceptionCreateTransferRuleRequest) GetPostDialDigits() (value *string) {
	if r == nil {
		return
	}
	return r.PostDialDigits
}

func (r *ReceptionCreateTransferRuleRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateTransferRuleRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionCreateTransferRuleRequest) SetLabel(label StrippedStr) {
	r.Label = label
}

func (r *ReceptionCreateTransferRuleRequest) SetPhoneNumber(phoneNumber PhoneNumber) {
	r.PhoneNumber = phoneNumber
}

func (r *ReceptionCreateTransferRuleRequest) SetCondition(condition StrippedStr) {
	r.Condition = condition
}

func (r *ReceptionCreateTransferRuleRequest) SetPostDialDigits(postDialDigits *string) {
	r.PostDialDigits = postDialDigits
}

func (r *ReceptionCreateTransferRuleRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateTransferRuleRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateTransferRuleRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateTransferRuleRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionCreateTransferRuleRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateTransferRuleRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

// Request model for updating an existing transfer rule.
type ReceptionUpdateTransferRuleRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the transfer rule. Returned when you add one, and listed in the receptionist's `config.transfer_rules`.
	RuleId string `json:"-" url:"-"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId     string      `json:"-" url:"-"`
	Label       StrippedStr `json:"label" xml:"label" required:"true"`
	PhoneNumber PhoneNumber `json:"phone_number" xml:"phone_number" required:"true"`
	Condition   StrippedStr `json:"condition" xml:"condition" required:"true"`
	// DTMF digits automatically played after the transfer destination picks up (e.g. '1' to answer a 'press 1 to accept' screen, or 'www1234' to key an extension after a pause). Valid characters are 0-9, *, #, w (0.5s pause) and W (1s pause). Twilio transfers only.
	PostDialDigits *string `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty" maxLength:"50" minLength:"1" pattern:"^[0-9*#wW]*$"`
}

func (r *ReceptionUpdateTransferRuleRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateTransferRuleRequest) GetRuleId() (value string) {
	if r == nil {
		return
	}
	return r.RuleId
}

func (r *ReceptionUpdateTransferRuleRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionUpdateTransferRuleRequest) GetLabel() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Label
}

func (r *ReceptionUpdateTransferRuleRequest) GetPhoneNumber() (value PhoneNumber) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

func (r *ReceptionUpdateTransferRuleRequest) GetCondition() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Condition
}

func (r *ReceptionUpdateTransferRuleRequest) GetPostDialDigits() (value *string) {
	if r == nil {
		return
	}
	return r.PostDialDigits
}

func (r *ReceptionUpdateTransferRuleRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateTransferRuleRequest) SetRuleId(ruleId string) {
	r.RuleId = ruleId
}

func (r *ReceptionUpdateTransferRuleRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionUpdateTransferRuleRequest) SetLabel(label StrippedStr) {
	r.Label = label
}

func (r *ReceptionUpdateTransferRuleRequest) SetPhoneNumber(phoneNumber PhoneNumber) {
	r.PhoneNumber = phoneNumber
}

func (r *ReceptionUpdateTransferRuleRequest) SetCondition(condition StrippedStr) {
	r.Condition = condition
}

func (r *ReceptionUpdateTransferRuleRequest) SetPostDialDigits(postDialDigits *string) {
	r.PostDialDigits = postDialDigits
}

func (r *ReceptionUpdateTransferRuleRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateTransferRuleRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateTransferRuleRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateTransferRuleRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateTransferRuleRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateTransferRuleRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	deleteTransferRulesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteTransferRulesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteTransferRulesRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteTransferRulesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteTransferRulesRequestFieldXiWorkspaceId)
}
