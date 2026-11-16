package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Put a workspace with no subscription onto one of the reseller's tiers.
type AssignResellerTierRequest struct {
	// ID of the customer workspace. Returned when you provision one; list them with `GET /v1/reception/reseller/workspaces`.
	WorkspaceId   string                                                     `json:"-" url:"-"`
	TierId        string                                                     `json:"tier_id" xml:"tier_id" required:"true"`
	BillingPeriod *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod `json:"billing_period,omitempty" xml:"billing_period,omitempty"`
	// Client-generated key that makes a retry safe. The same key and body within 24 hours returns the first result without running again. While that first request is still running, a retry gets 409 `idempotency_key_in_progress`; a different body with the same key gets 400 `idempotency_key_reused`. A request that fails is not saved, so it can be retried with the same key.
	IdempotencyKey *string `json:"idempotency_key,omitempty" xml:"idempotency_key,omitempty" maxLength:"255"`
}

func (a *AssignResellerTierRequest) GetWorkspaceId() (value string) {
	if a == nil {
		return
	}
	return a.WorkspaceId
}

func (a *AssignResellerTierRequest) GetTierId() (value string) {
	if a == nil {
		return
	}
	return a.TierId
}

func (a *AssignResellerTierRequest) GetBillingPeriod() (value *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	if a == nil {
		return
	}
	return a.BillingPeriod
}

func (a *AssignResellerTierRequest) GetIdempotencyKey() (value *string) {
	if a == nil {
		return
	}
	return a.IdempotencyKey
}

func (a *AssignResellerTierRequest) SetWorkspaceId(workspaceId string) {
	a.WorkspaceId = workspaceId
}

func (a *AssignResellerTierRequest) SetTierId(tierId string) {
	a.TierId = tierId
}

func (a *AssignResellerTierRequest) SetBillingPeriod(billingPeriod *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	a.BillingPeriod = billingPeriod
}

func (a *AssignResellerTierRequest) SetIdempotencyKey(idempotencyKey *string) {
	a.IdempotencyKey = idempotencyKey
}

func (a *AssignResellerTierRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AssignResellerTierRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssignResellerTierRequest to string"
	}
	return string(jsonData)
}

func (a *AssignResellerTierRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

func (a *AssignResellerTierRequest) MarshalJSON() ([]byte, error) {
	type embed AssignResellerTierRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

// Move a workspace to another of the reseller's tiers.
type ChangeResellerTierRequest struct {
	// ID of the customer workspace. Returned when you provision one; list them with `GET /v1/reception/reseller/workspaces`.
	WorkspaceId   string                                                     `json:"-" url:"-"`
	TierId        string                                                     `json:"tier_id" xml:"tier_id" required:"true"`
	BillingPeriod *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod `json:"billing_period,omitempty" xml:"billing_period,omitempty"`
	// Client-generated key that makes a retry safe. The same key and body within 24 hours returns the first result without running again. While that first request is still running, a retry gets 409 `idempotency_key_in_progress`; a different body with the same key gets 400 `idempotency_key_reused`. A request that fails is not saved, so it can be retried with the same key.
	IdempotencyKey *string `json:"idempotency_key,omitempty" xml:"idempotency_key,omitempty" maxLength:"255"`
}

func (c *ChangeResellerTierRequest) GetWorkspaceId() (value string) {
	if c == nil {
		return
	}
	return c.WorkspaceId
}

func (c *ChangeResellerTierRequest) GetTierId() (value string) {
	if c == nil {
		return
	}
	return c.TierId
}

func (c *ChangeResellerTierRequest) GetBillingPeriod() (value *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	if c == nil {
		return
	}
	return c.BillingPeriod
}

func (c *ChangeResellerTierRequest) GetIdempotencyKey() (value *string) {
	if c == nil {
		return
	}
	return c.IdempotencyKey
}

func (c *ChangeResellerTierRequest) SetWorkspaceId(workspaceId string) {
	c.WorkspaceId = workspaceId
}

func (c *ChangeResellerTierRequest) SetTierId(tierId string) {
	c.TierId = tierId
}

func (c *ChangeResellerTierRequest) SetBillingPeriod(billingPeriod *types.ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	c.BillingPeriod = billingPeriod
}

func (c *ChangeResellerTierRequest) SetIdempotencyKey(idempotencyKey *string) {
	c.IdempotencyKey = idempotencyKey
}

func (c *ChangeResellerTierRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c ChangeResellerTierRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChangeResellerTierRequest to string"
	}
	return string(jsonData)
}

func (c *ChangeResellerTierRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *ChangeResellerTierRequest) MarshalJSON() ([]byte, error) {
	type embed ChangeResellerTierRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
