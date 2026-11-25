package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProvisionResellerWorkspaceResponse struct {
	WorkspaceId       string  `json:"workspace_id" xml:"workspace_id" required:"true"`
	BusinessName      string  `json:"business_name" xml:"business_name" required:"true"`
	BillingTier       *string `json:"billing_tier" xml:"billing_tier" required:"true"`
	ReceptionistCount int     `json:"receptionist_count" xml:"receptionist_count" required:"true"`
	CallsLast30Days   int     `json:"calls_last_30_days" xml:"calls_last_30_days" required:"true"`
	CreatedAtUnix     int     `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// The reseller-managed subscription, when this workspace has one. None means the reseller sells nothing here: no tier is assigned yet, or the workspace pays for its own subscription (see charge_customer).
	Subscription *ResellerWorkspaceSubscription `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// The workspace pays for its own subscription, so its tier can't be assigned or changed here.
	ChargeCustomer *bool `json:"charge_customer,omitempty" xml:"charge_customer,omitempty"`
	// A charge for this workspace is unpaid. Assigning or changing a tier is refused until it clears, so surface it rather than letting it fail.
	HasPaymentIssue *bool `json:"has_payment_issue,omitempty" xml:"has_payment_issue,omitempty"`
	// When onboarding finished. Until then the onboarding wizard may replace or remove the workspace's business data.
	OnboardingCompletedAtUnix *int                `json:"onboarding_completed_at_unix,omitempty" xml:"onboarding_completed_at_unix,omitempty"`
	MemberResults             []*TeamInviteResult `json:"member_results,omitempty" xml:"member_results,omitempty"`
	extraProperties           map[string]interface{}
}

func (p *ProvisionResellerWorkspaceResponse) GetWorkspaceId() (value string) {
	if p == nil {
		return
	}
	return p.WorkspaceId
}

func (p *ProvisionResellerWorkspaceResponse) GetBusinessName() (value string) {
	if p == nil {
		return
	}
	return p.BusinessName
}

func (p *ProvisionResellerWorkspaceResponse) GetBillingTier() (value *string) {
	if p == nil {
		return
	}
	return p.BillingTier
}

func (p *ProvisionResellerWorkspaceResponse) GetReceptionistCount() (value int) {
	if p == nil {
		return
	}
	return p.ReceptionistCount
}

func (p *ProvisionResellerWorkspaceResponse) GetCallsLast30Days() (value int) {
	if p == nil {
		return
	}
	return p.CallsLast30Days
}

func (p *ProvisionResellerWorkspaceResponse) GetCreatedAtUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtUnix
}

func (p *ProvisionResellerWorkspaceResponse) GetSubscription() (value *ResellerWorkspaceSubscription) {
	if p == nil {
		return
	}
	return p.Subscription
}

func (p *ProvisionResellerWorkspaceResponse) GetChargeCustomer() (value *bool) {
	if p == nil {
		return
	}
	return p.ChargeCustomer
}

func (p *ProvisionResellerWorkspaceResponse) GetHasPaymentIssue() (value *bool) {
	if p == nil {
		return
	}
	return p.HasPaymentIssue
}

func (p *ProvisionResellerWorkspaceResponse) GetOnboardingCompletedAtUnix() (value *int) {
	if p == nil {
		return
	}
	return p.OnboardingCompletedAtUnix
}

func (p *ProvisionResellerWorkspaceResponse) GetMemberResults() (value []*TeamInviteResult) {
	if p == nil {
		return
	}
	return p.MemberResults
}

func (p *ProvisionResellerWorkspaceResponse) SetWorkspaceId(workspaceId string) {
	p.WorkspaceId = workspaceId
}

func (p *ProvisionResellerWorkspaceResponse) SetBusinessName(businessName string) {
	p.BusinessName = businessName
}

func (p *ProvisionResellerWorkspaceResponse) SetBillingTier(billingTier *string) {
	p.BillingTier = billingTier
}

func (p *ProvisionResellerWorkspaceResponse) SetReceptionistCount(receptionistCount int) {
	p.ReceptionistCount = receptionistCount
}

func (p *ProvisionResellerWorkspaceResponse) SetCallsLast30Days(callsLast30Days int) {
	p.CallsLast30Days = callsLast30Days
}

func (p *ProvisionResellerWorkspaceResponse) SetCreatedAtUnix(createdAtUnix int) {
	p.CreatedAtUnix = createdAtUnix
}

func (p *ProvisionResellerWorkspaceResponse) SetSubscription(subscription *ResellerWorkspaceSubscription) {
	p.Subscription = subscription
}

func (p *ProvisionResellerWorkspaceResponse) SetChargeCustomer(chargeCustomer *bool) {
	p.ChargeCustomer = chargeCustomer
}

func (p *ProvisionResellerWorkspaceResponse) SetHasPaymentIssue(hasPaymentIssue *bool) {
	p.HasPaymentIssue = hasPaymentIssue
}

func (p *ProvisionResellerWorkspaceResponse) SetOnboardingCompletedAtUnix(onboardingCompletedAtUnix *int) {
	p.OnboardingCompletedAtUnix = onboardingCompletedAtUnix
}

func (p *ProvisionResellerWorkspaceResponse) SetMemberResults(memberResults []*TeamInviteResult) {
	p.MemberResults = memberResults
}

func (p *ProvisionResellerWorkspaceResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProvisionResellerWorkspaceResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProvisionResellerWorkspaceResponse to string"
	}
	return string(jsonData)
}

func (p *ProvisionResellerWorkspaceResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProvisionResellerWorkspaceResponse) MarshalJSON() ([]byte, error) {
	type embed ProvisionResellerWorkspaceResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
