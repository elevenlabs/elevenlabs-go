package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResellerWorkspaceSummary struct {
	WorkspaceId       string  `json:"workspace_id" xml:"workspace_id" required:"true"`
	BusinessName      string  `json:"business_name" xml:"business_name" required:"true"`
	BillingTier       *string `json:"billing_tier" xml:"billing_tier" required:"true"`
	ReceptionistCount int     `json:"receptionist_count" xml:"receptionist_count" required:"true"`
	CallsLast30Days   int     `json:"calls_last_30_days" xml:"calls_last_30_days" required:"true"`
	CreatedAtUnix     int     `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// The reseller-managed subscription, when this workspace has one. None means nothing is sold here yet and a tier can be assigned.
	Subscription *ResellerWorkspaceSubscription `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// A charge for this workspace is unpaid. Assigning or changing a tier is refused until it clears, so surface it rather than letting it fail.
	HasPaymentIssue *bool `json:"has_payment_issue,omitempty" xml:"has_payment_issue,omitempty"`
	// When onboarding finished. Until then the onboarding wizard may replace or remove the workspace's business data.
	OnboardingCompletedAtUnix *int `json:"onboarding_completed_at_unix,omitempty" xml:"onboarding_completed_at_unix,omitempty"`
	extraProperties           map[string]interface{}
}

func (r *ResellerWorkspaceSummary) GetWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.WorkspaceId
}

func (r *ResellerWorkspaceSummary) GetBusinessName() (value string) {
	if r == nil {
		return
	}
	return r.BusinessName
}

func (r *ResellerWorkspaceSummary) GetBillingTier() (value *string) {
	if r == nil {
		return
	}
	return r.BillingTier
}

func (r *ResellerWorkspaceSummary) GetReceptionistCount() (value int) {
	if r == nil {
		return
	}
	return r.ReceptionistCount
}

func (r *ResellerWorkspaceSummary) GetCallsLast30Days() (value int) {
	if r == nil {
		return
	}
	return r.CallsLast30Days
}

func (r *ResellerWorkspaceSummary) GetCreatedAtUnix() (value int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ResellerWorkspaceSummary) GetSubscription() (value *ResellerWorkspaceSubscription) {
	if r == nil {
		return
	}
	return r.Subscription
}

func (r *ResellerWorkspaceSummary) GetHasPaymentIssue() (value *bool) {
	if r == nil {
		return
	}
	return r.HasPaymentIssue
}

func (r *ResellerWorkspaceSummary) GetOnboardingCompletedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.OnboardingCompletedAtUnix
}

func (r *ResellerWorkspaceSummary) SetWorkspaceId(workspaceId string) {
	r.WorkspaceId = workspaceId
}

func (r *ResellerWorkspaceSummary) SetBusinessName(businessName string) {
	r.BusinessName = businessName
}

func (r *ResellerWorkspaceSummary) SetBillingTier(billingTier *string) {
	r.BillingTier = billingTier
}

func (r *ResellerWorkspaceSummary) SetReceptionistCount(receptionistCount int) {
	r.ReceptionistCount = receptionistCount
}

func (r *ResellerWorkspaceSummary) SetCallsLast30Days(callsLast30Days int) {
	r.CallsLast30Days = callsLast30Days
}

func (r *ResellerWorkspaceSummary) SetCreatedAtUnix(createdAtUnix int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ResellerWorkspaceSummary) SetSubscription(subscription *ResellerWorkspaceSubscription) {
	r.Subscription = subscription
}

func (r *ResellerWorkspaceSummary) SetHasPaymentIssue(hasPaymentIssue *bool) {
	r.HasPaymentIssue = hasPaymentIssue
}

func (r *ResellerWorkspaceSummary) SetOnboardingCompletedAtUnix(onboardingCompletedAtUnix *int) {
	r.OnboardingCompletedAtUnix = onboardingCompletedAtUnix
}

func (r *ResellerWorkspaceSummary) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerWorkspaceSummary) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerWorkspaceSummary to string"
	}
	return string(jsonData)
}

func (r *ResellerWorkspaceSummary) UnmarshalJSON(data []byte) error {
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

func (r *ResellerWorkspaceSummary) MarshalJSON() ([]byte, error) {
	type embed ResellerWorkspaceSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
