package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Body for giving a receptionist a phone number.
type CreatePhoneAssignmentRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// The receptionist that answers calls to the number.
	CustomerFacingAgentId string `json:"customer_facing_agent_id" xml:"customer_facing_agent_id" required:"true"`
	// Preferred US or Canadian area codes, e.g. `["415", "628"]`. Codes outside the business's country are ignored, and any number is given when none of these is available.
	AreaCodes []string `json:"area_codes,omitempty" xml:"area_codes,omitempty"`
	// Local or mobile, in a country that requires paperwork before a number takes calls. Leave unset to try mobile first. Ignored in the US and Canada.
	NumberType *types.ApplicationNumberType `json:"number_type,omitempty" xml:"number_type,omitempty"`
	// Client-generated key that makes a retry safe. The same key and body within 24 hours returns the first result without running again. While that first request is still running, a retry gets 409 `idempotency_key_in_progress`; a different body with the same key gets 400 `idempotency_key_reused`. A request that fails is not saved, so it can be retried with the same key.
	IdempotencyKey *string `json:"idempotency_key,omitempty" xml:"idempotency_key,omitempty" maxLength:"255"`
}

func (c *CreatePhoneAssignmentRequest) GetXiWorkspaceId() (value string) {
	if c == nil {
		return
	}
	return c.XiWorkspaceId
}

func (c *CreatePhoneAssignmentRequest) GetCustomerFacingAgentId() (value string) {
	if c == nil {
		return
	}
	return c.CustomerFacingAgentId
}

func (c *CreatePhoneAssignmentRequest) GetAreaCodes() (value []string) {
	if c == nil {
		return
	}
	return c.AreaCodes
}

func (c *CreatePhoneAssignmentRequest) GetNumberType() (value *types.ApplicationNumberType) {
	if c == nil {
		return
	}
	return c.NumberType
}

func (c *CreatePhoneAssignmentRequest) GetIdempotencyKey() (value *string) {
	if c == nil {
		return
	}
	return c.IdempotencyKey
}

func (c *CreatePhoneAssignmentRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	c.XiWorkspaceId = xiWorkspaceId
}

func (c *CreatePhoneAssignmentRequest) SetCustomerFacingAgentId(customerFacingAgentId string) {
	c.CustomerFacingAgentId = customerFacingAgentId
}

func (c *CreatePhoneAssignmentRequest) SetAreaCodes(areaCodes []string) {
	c.AreaCodes = areaCodes
}

func (c *CreatePhoneAssignmentRequest) SetNumberType(numberType *types.ApplicationNumberType) {
	c.NumberType = numberType
}

func (c *CreatePhoneAssignmentRequest) SetIdempotencyKey(idempotencyKey *string) {
	c.IdempotencyKey = idempotencyKey
}

func (c *CreatePhoneAssignmentRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreatePhoneAssignmentRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePhoneAssignmentRequest to string"
	}
	return string(jsonData)
}

func (c *CreatePhoneAssignmentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreatePhoneAssignmentRequest) MarshalJSON() ([]byte, error) {
	type embed CreatePhoneAssignmentRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

// Body for attaching a phone number to a different receptionist.
type ReassignPhoneRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the phone number assignment. Returned when you assign a number; list them with `GET /v1/reception/phone-numbers/assignments`.
	PhoneAssignmentId string `json:"-" url:"-"`
	// The receptionist that answers calls to the number from now on.
	CustomerFacingAgentId string `json:"customer_facing_agent_id" xml:"customer_facing_agent_id" required:"true"`
}

func (r *ReassignPhoneRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReassignPhoneRequest) GetPhoneAssignmentId() (value string) {
	if r == nil {
		return
	}
	return r.PhoneAssignmentId
}

func (r *ReassignPhoneRequest) GetCustomerFacingAgentId() (value string) {
	if r == nil {
		return
	}
	return r.CustomerFacingAgentId
}

func (r *ReassignPhoneRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReassignPhoneRequest) SetPhoneAssignmentId(phoneAssignmentId string) {
	r.PhoneAssignmentId = phoneAssignmentId
}

func (r *ReassignPhoneRequest) SetCustomerFacingAgentId(customerFacingAgentId string) {
	r.CustomerFacingAgentId = customerFacingAgentId
}

func (r *ReassignPhoneRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReassignPhoneRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReassignPhoneRequest to string"
	}
	return string(jsonData)
}

func (r *ReassignPhoneRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReassignPhoneRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReassignPhoneRequest(tmp)
	return nil
}

func (r *ReassignPhoneRequest) MarshalJSON() ([]byte, error) {
	type embed ReassignPhoneRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listPhoneNumbersRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListPhoneNumbersRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListPhoneNumbersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListPhoneNumbersRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listPhoneNumbersRequestFieldXiWorkspaceId)
}

var (
	deletePhoneNumbersRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeletePhoneNumbersRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeletePhoneNumbersRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeletePhoneNumbersRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deletePhoneNumbersRequestFieldXiWorkspaceId)
}
