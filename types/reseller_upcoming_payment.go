package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A workspace's next renewal, estimated before its invoice is created.
type ResellerUpcomingPayment struct {
	WorkspaceId   string `json:"workspace_id" xml:"workspace_id" required:"true"`
	WorkspaceName string `json:"workspace_name" xml:"workspace_name" required:"true"`
	// A reseller-managed paid subscription, as the reseller can act on it.
	//
	// Present only for a workspace the reseller actually sells to: untiered and
	// trial workspaces have no subscription to change or end, and are assigned to
	// instead.
	Subscription *ResellerWorkspaceSubscription `json:"subscription" xml:"subscription" required:"true"`
	// The plan the renewal charges upfront: the pending downgrade's tier when there is one.
	PlanUsd float64 `json:"plan_usd" xml:"plan_usd" required:"true"`
	// This period's usage beyond its included minutes so far, charged at the renewal and still growing.
	OverageSoFarUsd float64 `json:"overage_so_far_usd" xml:"overage_so_far_usd" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResellerUpcomingPayment) GetWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.WorkspaceId
}

func (r *ResellerUpcomingPayment) GetWorkspaceName() (value string) {
	if r == nil {
		return
	}
	return r.WorkspaceName
}

func (r *ResellerUpcomingPayment) GetSubscription() (value *ResellerWorkspaceSubscription) {
	if r == nil {
		return
	}
	return r.Subscription
}

func (r *ResellerUpcomingPayment) GetPlanUsd() (value float64) {
	if r == nil {
		return
	}
	return r.PlanUsd
}

func (r *ResellerUpcomingPayment) GetOverageSoFarUsd() (value float64) {
	if r == nil {
		return
	}
	return r.OverageSoFarUsd
}

func (r *ResellerUpcomingPayment) SetWorkspaceId(workspaceId string) {
	r.WorkspaceId = workspaceId
}

func (r *ResellerUpcomingPayment) SetWorkspaceName(workspaceName string) {
	r.WorkspaceName = workspaceName
}

func (r *ResellerUpcomingPayment) SetSubscription(subscription *ResellerWorkspaceSubscription) {
	r.Subscription = subscription
}

func (r *ResellerUpcomingPayment) SetPlanUsd(planUsd float64) {
	r.PlanUsd = planUsd
}

func (r *ResellerUpcomingPayment) SetOverageSoFarUsd(overageSoFarUsd float64) {
	r.OverageSoFarUsd = overageSoFarUsd
}

func (r *ResellerUpcomingPayment) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerUpcomingPayment) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerUpcomingPayment to string"
	}
	return string(jsonData)
}

func (r *ResellerUpcomingPayment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ResellerUpcomingPayment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ResellerUpcomingPayment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerUpcomingPayment) MarshalJSON() ([]byte, error) {
	type embed ResellerUpcomingPayment
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
