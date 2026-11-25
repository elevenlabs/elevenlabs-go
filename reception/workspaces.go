package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionProvisionResellerWorkspaceRequest struct {
	BusinessName string `json:"business_name" xml:"business_name" required:"true" maxLength:"100" minLength:"1"`
	// The business's vertical, used with `complete_onboarding`. Without it, onboarding picks the vertical as usual.
	Vertical *types.ReceptionVertical `json:"vertical,omitempty" xml:"vertical,omitempty"`
	// Bill the workspace's own customer instead of the reseller: it checks out and sees its billing itself, and the reseller earns a commission. Only for a reseller on standard billing.
	ChargeCustomer *bool `json:"charge_customer,omitempty" xml:"charge_customer,omitempty"`
	// Paid tier to subscribe to. Omitted leaves the workspace with no subscription.
	Subscription *types.ProvisionResellerWorkspaceSubscription `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// A free number from the reseller's phone pool, reserved for the workspace until onboarding binds it to the receptionist.
	PhoneNumber *string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// People to invite once the workspace exists, capped by the tier's team size. Best-effort: each outcome is in `member_results`.
	Members []*types.TeamInviteRequest `json:"members,omitempty" xml:"members,omitempty" maxItems:"100"`
	// Finish setup right away, as `POST /v1/reception/reseller/workspaces/{workspace_id}/complete-onboarding` does. Best-effort: `onboarding_completed_at_unix` stays null on failure, and that call can be retried.
	CompleteOnboarding *bool `json:"complete_onboarding,omitempty" xml:"complete_onboarding,omitempty"`
	// Client-generated key that makes a retry safe. The same key and body within 24 hours returns the first result without running again. While that first request is still running, a retry gets 409 `idempotency_key_in_progress`; a different body with the same key gets 400 `idempotency_key_reused`. A request that fails is not saved, so it can be retried with the same key.
	IdempotencyKey *string `json:"idempotency_key,omitempty" xml:"idempotency_key,omitempty" maxLength:"255"`
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetBusinessName() (value string) {
	if r == nil {
		return
	}
	return r.BusinessName
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetVertical() (value *types.ReceptionVertical) {
	if r == nil {
		return
	}
	return r.Vertical
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetChargeCustomer() (value *bool) {
	if r == nil {
		return
	}
	return r.ChargeCustomer
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetSubscription() (value *types.ProvisionResellerWorkspaceSubscription) {
	if r == nil {
		return
	}
	return r.Subscription
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetPhoneNumber() (value *string) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetMembers() (value []*types.TeamInviteRequest) {
	if r == nil {
		return
	}
	return r.Members
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetCompleteOnboarding() (value *bool) {
	if r == nil {
		return
	}
	return r.CompleteOnboarding
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetIdempotencyKey() (value *string) {
	if r == nil {
		return
	}
	return r.IdempotencyKey
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetBusinessName(businessName string) {
	r.BusinessName = businessName
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetVertical(vertical *types.ReceptionVertical) {
	r.Vertical = vertical
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetChargeCustomer(chargeCustomer *bool) {
	r.ChargeCustomer = chargeCustomer
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetSubscription(subscription *types.ProvisionResellerWorkspaceSubscription) {
	r.Subscription = subscription
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetPhoneNumber(phoneNumber *string) {
	r.PhoneNumber = phoneNumber
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetMembers(members []*types.TeamInviteRequest) {
	r.Members = members
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetCompleteOnboarding(completeOnboarding *bool) {
	r.CompleteOnboarding = completeOnboarding
}

func (r *ReceptionProvisionResellerWorkspaceRequest) SetIdempotencyKey(idempotencyKey *string) {
	r.IdempotencyKey = idempotencyKey
}

func (r *ReceptionProvisionResellerWorkspaceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionProvisionResellerWorkspaceRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionProvisionResellerWorkspaceRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionProvisionResellerWorkspaceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionProvisionResellerWorkspaceRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionProvisionResellerWorkspaceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listWorkspacesRequestFieldCursor   = big.NewInt(1 << 0)
	listWorkspacesRequestFieldPageSize = big.NewInt(1 << 1)
)

type ListWorkspacesRequest struct {
	Cursor   *string `json:"-" url:"cursor,omitempty"`
	PageSize *int    `json:"-" url:"page_size,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListWorkspacesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListWorkspacesRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listWorkspacesRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListWorkspacesRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listWorkspacesRequestFieldPageSize)
}
