package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PendingSubscriptionSwitchResponseModel struct {
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The tier to change to.
	NextTier          PendingSubscriptionSwitchResponseModelNextTier `json:"next_tier" xml:"next_tier" required:"true"`
	NextBillingPeriod BillingPeriod                                  `json:"next_billing_period" xml:"next_billing_period" required:"true"`
	// The timestamp of the change.
	TimestampSeconds int `json:"timestamp_seconds" xml:"timestamp_seconds" required:"true"`
	extraProperties  map[string]interface{}
}

func (p *PendingSubscriptionSwitchResponseModel) GetKind() (value *string) {
	if p == nil {
		return
	}
	return p.Kind
}

func (p *PendingSubscriptionSwitchResponseModel) GetNextTier() (value PendingSubscriptionSwitchResponseModelNextTier) {
	if p == nil {
		return
	}
	return p.NextTier
}

func (p *PendingSubscriptionSwitchResponseModel) GetNextBillingPeriod() (value BillingPeriod) {
	if p == nil {
		return
	}
	return p.NextBillingPeriod
}

func (p *PendingSubscriptionSwitchResponseModel) GetTimestampSeconds() (value int) {
	if p == nil {
		return
	}
	return p.TimestampSeconds
}

func (p *PendingSubscriptionSwitchResponseModel) SetKind(kind *string) {
	p.Kind = kind
}

func (p *PendingSubscriptionSwitchResponseModel) SetNextTier(nextTier PendingSubscriptionSwitchResponseModelNextTier) {
	p.NextTier = nextTier
}

func (p *PendingSubscriptionSwitchResponseModel) SetNextBillingPeriod(nextBillingPeriod BillingPeriod) {
	p.NextBillingPeriod = nextBillingPeriod
}

func (p *PendingSubscriptionSwitchResponseModel) SetTimestampSeconds(timestampSeconds int) {
	p.TimestampSeconds = timestampSeconds
}

func (p *PendingSubscriptionSwitchResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PendingSubscriptionSwitchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PendingSubscriptionSwitchResponseModel to string"
	}
	return string(jsonData)
}

func (p *PendingSubscriptionSwitchResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PendingSubscriptionSwitchResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PendingSubscriptionSwitchResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PendingSubscriptionSwitchResponseModel) MarshalJSON() ([]byte, error) {
	type embed PendingSubscriptionSwitchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
