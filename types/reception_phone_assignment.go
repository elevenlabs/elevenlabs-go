package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A phone number held by the workspace.
type ReceptionPhoneAssignment struct {
	// Phone number in E.164 format.
	PhoneNumber *PhoneNumber `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// True when the phone is bring-your-own: registered by the workspace using their own telephony provider (Twilio credentials or SIP trunk). BYO phones are never part of the shared pool, do not count against the plan's phone number limit, and are deleted (rather than returned to the pool) on release.
	IsByo *bool `json:"is_byo,omitempty" xml:"is_byo,omitempty"`
	// How calls to the number reach ElevenLabs: through Twilio or a SIP trunk.
	Provider *ReceptionPhoneProvider `json:"provider,omitempty" xml:"provider,omitempty"`
	// US/Canada NANP area code (NPA) as an integer.
	NanpAreaCode *int `json:"nanp_area_code,omitempty" xml:"nanp_area_code,omitempty"`
	// The receptionist that answers this number. Null when the workspace holds the number but no receptionist answers it yet.
	CustomerFacingAgentId *string `json:"customer_facing_agent_id,omitempty" xml:"customer_facing_agent_id,omitempty"`
	// Unix timestamp when this number was deactivated; it is released 14 days later. Null when active.
	DeactivatedAtUnix *int `json:"deactivated_at_unix,omitempty" xml:"deactivated_at_unix,omitempty"`
	// Why this number was deactivated. Null when active.
	DeactivationReason *ReceptionPhoneAssignmentDeactivationReason `json:"deactivation_reason,omitempty" xml:"deactivation_reason,omitempty"`
	// Null for numbers that need no paperwork. `held` while the workspace's own application is under review, `approved` once it passed and the number is fully usable.
	RegulatoryStatus *ReceptionPhoneRegulatoryStatus `json:"regulatory_status,omitempty" xml:"regulatory_status,omitempty"`
	// The regulated number type this number was issued as. Null for numbers that need no paperwork.
	RegulatoryNumberType *ApplicationNumberType `json:"regulatory_number_type,omitempty" xml:"regulatory_number_type,omitempty"`
	// ISO country the regulated number was issued in. Null for numbers that need no paperwork.
	RegulatoryCountry *CountryCode `json:"regulatory_country,omitempty" xml:"regulatory_country,omitempty"`
	// The phone number assignment's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	// Whether the number is active; false during the 14 days before a deactivated number is released.
	IsActive *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	// Whether inbound calls from the public may reach the receptionist.
	AcceptsPublicCalls *bool `json:"accepts_public_calls,omitempty" xml:"accepts_public_calls,omitempty"`
	extraProperties    map[string]interface{}
}

func (r *ReceptionPhoneAssignment) GetPhoneNumber() (value *PhoneNumber) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

func (r *ReceptionPhoneAssignment) GetIsByo() (value *bool) {
	if r == nil {
		return
	}
	return r.IsByo
}

func (r *ReceptionPhoneAssignment) GetProvider() (value *ReceptionPhoneProvider) {
	if r == nil {
		return
	}
	return r.Provider
}

func (r *ReceptionPhoneAssignment) GetNanpAreaCode() (value *int) {
	if r == nil {
		return
	}
	return r.NanpAreaCode
}

func (r *ReceptionPhoneAssignment) GetCustomerFacingAgentId() (value *string) {
	if r == nil {
		return
	}
	return r.CustomerFacingAgentId
}

func (r *ReceptionPhoneAssignment) GetDeactivatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.DeactivatedAtUnix
}

func (r *ReceptionPhoneAssignment) GetDeactivationReason() (value *ReceptionPhoneAssignmentDeactivationReason) {
	if r == nil {
		return
	}
	return r.DeactivationReason
}

func (r *ReceptionPhoneAssignment) GetRegulatoryStatus() (value *ReceptionPhoneRegulatoryStatus) {
	if r == nil {
		return
	}
	return r.RegulatoryStatus
}

func (r *ReceptionPhoneAssignment) GetRegulatoryNumberType() (value *ApplicationNumberType) {
	if r == nil {
		return
	}
	return r.RegulatoryNumberType
}

func (r *ReceptionPhoneAssignment) GetRegulatoryCountry() (value *CountryCode) {
	if r == nil {
		return
	}
	return r.RegulatoryCountry
}

func (r *ReceptionPhoneAssignment) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionPhoneAssignment) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionPhoneAssignment) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionPhoneAssignment) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionPhoneAssignment) GetAcceptsPublicCalls() (value *bool) {
	if r == nil {
		return
	}
	return r.AcceptsPublicCalls
}

func (r *ReceptionPhoneAssignment) SetPhoneNumber(phoneNumber *PhoneNumber) {
	r.PhoneNumber = phoneNumber
}

func (r *ReceptionPhoneAssignment) SetIsByo(isByo *bool) {
	r.IsByo = isByo
}

func (r *ReceptionPhoneAssignment) SetProvider(provider *ReceptionPhoneProvider) {
	r.Provider = provider
}

func (r *ReceptionPhoneAssignment) SetNanpAreaCode(nanpAreaCode *int) {
	r.NanpAreaCode = nanpAreaCode
}

func (r *ReceptionPhoneAssignment) SetCustomerFacingAgentId(customerFacingAgentId *string) {
	r.CustomerFacingAgentId = customerFacingAgentId
}

func (r *ReceptionPhoneAssignment) SetDeactivatedAtUnix(deactivatedAtUnix *int) {
	r.DeactivatedAtUnix = deactivatedAtUnix
}

func (r *ReceptionPhoneAssignment) SetDeactivationReason(deactivationReason *ReceptionPhoneAssignmentDeactivationReason) {
	r.DeactivationReason = deactivationReason
}

func (r *ReceptionPhoneAssignment) SetRegulatoryStatus(regulatoryStatus *ReceptionPhoneRegulatoryStatus) {
	r.RegulatoryStatus = regulatoryStatus
}

func (r *ReceptionPhoneAssignment) SetRegulatoryNumberType(regulatoryNumberType *ApplicationNumberType) {
	r.RegulatoryNumberType = regulatoryNumberType
}

func (r *ReceptionPhoneAssignment) SetRegulatoryCountry(regulatoryCountry *CountryCode) {
	r.RegulatoryCountry = regulatoryCountry
}

func (r *ReceptionPhoneAssignment) SetId(id string) {
	r.Id = id
}

func (r *ReceptionPhoneAssignment) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionPhoneAssignment) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionPhoneAssignment) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionPhoneAssignment) SetAcceptsPublicCalls(acceptsPublicCalls *bool) {
	r.AcceptsPublicCalls = acceptsPublicCalls
}

func (r *ReceptionPhoneAssignment) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionPhoneAssignment) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionPhoneAssignment to string"
	}
	return string(jsonData)
}

func (r *ReceptionPhoneAssignment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionPhoneAssignment) MarshalJSON() ([]byte, error) {
	type embed ReceptionPhoneAssignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
