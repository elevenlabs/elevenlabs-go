package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProvisionResellerWorkspaceSubscription struct {
	TierId          string                                               `json:"tier_id" xml:"tier_id" required:"true"`
	BillingPeriod   *ProvisionResellerWorkspaceSubscriptionBillingPeriod `json:"billing_period,omitempty" xml:"billing_period,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProvisionResellerWorkspaceSubscription) GetTierId() (value string) {
	if p == nil {
		return
	}
	return p.TierId
}

func (p *ProvisionResellerWorkspaceSubscription) GetBillingPeriod() (value *ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	if p == nil {
		return
	}
	return p.BillingPeriod
}

func (p *ProvisionResellerWorkspaceSubscription) SetTierId(tierId string) {
	p.TierId = tierId
}

func (p *ProvisionResellerWorkspaceSubscription) SetBillingPeriod(billingPeriod *ProvisionResellerWorkspaceSubscriptionBillingPeriod) {
	p.BillingPeriod = billingPeriod
}

func (p *ProvisionResellerWorkspaceSubscription) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProvisionResellerWorkspaceSubscription) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProvisionResellerWorkspaceSubscription to string"
	}
	return string(jsonData)
}

func (p *ProvisionResellerWorkspaceSubscription) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProvisionResellerWorkspaceSubscription
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProvisionResellerWorkspaceSubscription(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProvisionResellerWorkspaceSubscription) MarshalJSON() ([]byte, error) {
	type embed ProvisionResellerWorkspaceSubscription
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
