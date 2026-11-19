package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ProvisionResellerWorkspaceRequest struct {
	BusinessName string `json:"business_name" xml:"business_name" required:"true" maxLength:"100" minLength:"1"`
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

func (p *ProvisionResellerWorkspaceRequest) GetBusinessName() (value string) {
	if p == nil {
		return
	}
	return p.BusinessName
}

func (p *ProvisionResellerWorkspaceRequest) GetSubscription() (value *types.ProvisionResellerWorkspaceSubscription) {
	if p == nil {
		return
	}
	return p.Subscription
}

func (p *ProvisionResellerWorkspaceRequest) GetPhoneNumber() (value *string) {
	if p == nil {
		return
	}
	return p.PhoneNumber
}

func (p *ProvisionResellerWorkspaceRequest) GetMembers() (value []*types.TeamInviteRequest) {
	if p == nil {
		return
	}
	return p.Members
}

func (p *ProvisionResellerWorkspaceRequest) GetCompleteOnboarding() (value *bool) {
	if p == nil {
		return
	}
	return p.CompleteOnboarding
}

func (p *ProvisionResellerWorkspaceRequest) GetIdempotencyKey() (value *string) {
	if p == nil {
		return
	}
	return p.IdempotencyKey
}

func (p *ProvisionResellerWorkspaceRequest) SetBusinessName(businessName string) {
	p.BusinessName = businessName
}

func (p *ProvisionResellerWorkspaceRequest) SetSubscription(subscription *types.ProvisionResellerWorkspaceSubscription) {
	p.Subscription = subscription
}

func (p *ProvisionResellerWorkspaceRequest) SetPhoneNumber(phoneNumber *string) {
	p.PhoneNumber = phoneNumber
}

func (p *ProvisionResellerWorkspaceRequest) SetMembers(members []*types.TeamInviteRequest) {
	p.Members = members
}

func (p *ProvisionResellerWorkspaceRequest) SetCompleteOnboarding(completeOnboarding *bool) {
	p.CompleteOnboarding = completeOnboarding
}

func (p *ProvisionResellerWorkspaceRequest) SetIdempotencyKey(idempotencyKey *string) {
	p.IdempotencyKey = idempotencyKey
}

func (p *ProvisionResellerWorkspaceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p ProvisionResellerWorkspaceRequest) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProvisionResellerWorkspaceRequest to string"
	}
	return string(jsonData)
}

func (p *ProvisionResellerWorkspaceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}

func (p *ProvisionResellerWorkspaceRequest) MarshalJSON() ([]byte, error) {
	type embed ProvisionResellerWorkspaceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
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
