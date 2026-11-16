package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A reseller-managed paid subscription, as the reseller can act on it.
//
// Present only for a workspace the reseller actually sells to: untiered and
// trial workspaces have no subscription to change or end, and are assigned to
// instead.
type ResellerWorkspaceSubscription struct {
	TierId string `json:"tier_id" xml:"tier_id" required:"true"`
	// Cadence being billed. None only for a subscription mid-provisioning.
	BillingPeriod *ResellerWorkspaceSubscriptionBillingPeriod `json:"billing_period" xml:"billing_period" required:"true"`
	// When the current period ends, and so when a deferred change lands.
	BillingPeriodEndUnix int `json:"billing_period_end_unix" xml:"billing_period_end_unix" required:"true"`
	// When the subscription ends. Set means a pending unassignment the reseller can still undo.
	CancelAtUnix *int `json:"cancel_at_unix,omitempty" xml:"cancel_at_unix,omitempty"`
	// Tier this workspace drops to at the period boundary, if a downgrade is already scheduled.
	PendingDowngradeTier *string `json:"pending_downgrade_tier,omitempty" xml:"pending_downgrade_tier,omitempty"`
	// When that scheduled downgrade takes effect.
	PendingDowngradeAtUnix *int `json:"pending_downgrade_at_unix,omitempty" xml:"pending_downgrade_at_unix,omitempty"`
	extraProperties        map[string]interface{}
}

func (r *ResellerWorkspaceSubscription) GetTierId() (value string) {
	if r == nil {
		return
	}
	return r.TierId
}

func (r *ResellerWorkspaceSubscription) GetBillingPeriod() (value *ResellerWorkspaceSubscriptionBillingPeriod) {
	if r == nil {
		return
	}
	return r.BillingPeriod
}

func (r *ResellerWorkspaceSubscription) GetBillingPeriodEndUnix() (value int) {
	if r == nil {
		return
	}
	return r.BillingPeriodEndUnix
}

func (r *ResellerWorkspaceSubscription) GetCancelAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CancelAtUnix
}

func (r *ResellerWorkspaceSubscription) GetPendingDowngradeTier() (value *string) {
	if r == nil {
		return
	}
	return r.PendingDowngradeTier
}

func (r *ResellerWorkspaceSubscription) GetPendingDowngradeAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.PendingDowngradeAtUnix
}

func (r *ResellerWorkspaceSubscription) SetTierId(tierId string) {
	r.TierId = tierId
}

func (r *ResellerWorkspaceSubscription) SetBillingPeriod(billingPeriod *ResellerWorkspaceSubscriptionBillingPeriod) {
	r.BillingPeriod = billingPeriod
}

func (r *ResellerWorkspaceSubscription) SetBillingPeriodEndUnix(billingPeriodEndUnix int) {
	r.BillingPeriodEndUnix = billingPeriodEndUnix
}

func (r *ResellerWorkspaceSubscription) SetCancelAtUnix(cancelAtUnix *int) {
	r.CancelAtUnix = cancelAtUnix
}

func (r *ResellerWorkspaceSubscription) SetPendingDowngradeTier(pendingDowngradeTier *string) {
	r.PendingDowngradeTier = pendingDowngradeTier
}

func (r *ResellerWorkspaceSubscription) SetPendingDowngradeAtUnix(pendingDowngradeAtUnix *int) {
	r.PendingDowngradeAtUnix = pendingDowngradeAtUnix
}

func (r *ResellerWorkspaceSubscription) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerWorkspaceSubscription) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerWorkspaceSubscription to string"
	}
	return string(jsonData)
}

func (r *ResellerWorkspaceSubscription) UnmarshalJSON(data []byte) error {
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

func (r *ResellerWorkspaceSubscription) MarshalJSON() ([]byte, error) {
	type embed ResellerWorkspaceSubscription
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
