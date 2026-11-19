package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubscriptionUsageResponseModel struct {
	// The rollover credits quota.
	RolloverCreditsQuota int `json:"rollover_credits_quota" xml:"rollover_credits_quota" required:"true"`
	// The subscription cycle credits quota.
	SubscriptionCycleCreditsQuota int `json:"subscription_cycle_credits_quota" xml:"subscription_cycle_credits_quota" required:"true"`
	// The manually gifted credits quota.
	ManuallyGiftedCreditsQuota int `json:"manually_gifted_credits_quota" xml:"manually_gifted_credits_quota" required:"true"`
	// The payg credits quota.
	PaygCreditsQuota *int `json:"payg_credits_quota,omitempty" xml:"payg_credits_quota,omitempty"`
	// The rollover credits used.
	RolloverCreditsUsed int `json:"rollover_credits_used" xml:"rollover_credits_used" required:"true"`
	// The subscription cycle credits used.
	SubscriptionCycleCreditsUsed int `json:"subscription_cycle_credits_used" xml:"subscription_cycle_credits_used" required:"true"`
	// The manually gifted credits used.
	ManuallyGiftedCreditsUsed int `json:"manually_gifted_credits_used" xml:"manually_gifted_credits_used" required:"true"`
	// The payg credits used.
	PaygCreditsUsed *int `json:"payg_credits_used,omitempty" xml:"payg_credits_used,omitempty"`
	// The paid usage based credits used.
	PaidUsageBasedCreditsUsed int `json:"paid_usage_based_credits_used" xml:"paid_usage_based_credits_used" required:"true"`
	// The actual reported credits.
	ActualReportedCredits int `json:"actual_reported_credits" xml:"actual_reported_credits" required:"true"`
	extraProperties       map[string]interface{}
}

func (s *SubscriptionUsageResponseModel) GetRolloverCreditsQuota() (value int) {
	if s == nil {
		return
	}
	return s.RolloverCreditsQuota
}

func (s *SubscriptionUsageResponseModel) GetSubscriptionCycleCreditsQuota() (value int) {
	if s == nil {
		return
	}
	return s.SubscriptionCycleCreditsQuota
}

func (s *SubscriptionUsageResponseModel) GetManuallyGiftedCreditsQuota() (value int) {
	if s == nil {
		return
	}
	return s.ManuallyGiftedCreditsQuota
}

func (s *SubscriptionUsageResponseModel) GetPaygCreditsQuota() (value *int) {
	if s == nil {
		return
	}
	return s.PaygCreditsQuota
}

func (s *SubscriptionUsageResponseModel) GetRolloverCreditsUsed() (value int) {
	if s == nil {
		return
	}
	return s.RolloverCreditsUsed
}

func (s *SubscriptionUsageResponseModel) GetSubscriptionCycleCreditsUsed() (value int) {
	if s == nil {
		return
	}
	return s.SubscriptionCycleCreditsUsed
}

func (s *SubscriptionUsageResponseModel) GetManuallyGiftedCreditsUsed() (value int) {
	if s == nil {
		return
	}
	return s.ManuallyGiftedCreditsUsed
}

func (s *SubscriptionUsageResponseModel) GetPaygCreditsUsed() (value *int) {
	if s == nil {
		return
	}
	return s.PaygCreditsUsed
}

func (s *SubscriptionUsageResponseModel) GetPaidUsageBasedCreditsUsed() (value int) {
	if s == nil {
		return
	}
	return s.PaidUsageBasedCreditsUsed
}

func (s *SubscriptionUsageResponseModel) GetActualReportedCredits() (value int) {
	if s == nil {
		return
	}
	return s.ActualReportedCredits
}

func (s *SubscriptionUsageResponseModel) SetRolloverCreditsQuota(rolloverCreditsQuota int) {
	s.RolloverCreditsQuota = rolloverCreditsQuota
}

func (s *SubscriptionUsageResponseModel) SetSubscriptionCycleCreditsQuota(subscriptionCycleCreditsQuota int) {
	s.SubscriptionCycleCreditsQuota = subscriptionCycleCreditsQuota
}

func (s *SubscriptionUsageResponseModel) SetManuallyGiftedCreditsQuota(manuallyGiftedCreditsQuota int) {
	s.ManuallyGiftedCreditsQuota = manuallyGiftedCreditsQuota
}

func (s *SubscriptionUsageResponseModel) SetPaygCreditsQuota(paygCreditsQuota *int) {
	s.PaygCreditsQuota = paygCreditsQuota
}

func (s *SubscriptionUsageResponseModel) SetRolloverCreditsUsed(rolloverCreditsUsed int) {
	s.RolloverCreditsUsed = rolloverCreditsUsed
}

func (s *SubscriptionUsageResponseModel) SetSubscriptionCycleCreditsUsed(subscriptionCycleCreditsUsed int) {
	s.SubscriptionCycleCreditsUsed = subscriptionCycleCreditsUsed
}

func (s *SubscriptionUsageResponseModel) SetManuallyGiftedCreditsUsed(manuallyGiftedCreditsUsed int) {
	s.ManuallyGiftedCreditsUsed = manuallyGiftedCreditsUsed
}

func (s *SubscriptionUsageResponseModel) SetPaygCreditsUsed(paygCreditsUsed *int) {
	s.PaygCreditsUsed = paygCreditsUsed
}

func (s *SubscriptionUsageResponseModel) SetPaidUsageBasedCreditsUsed(paidUsageBasedCreditsUsed int) {
	s.PaidUsageBasedCreditsUsed = paidUsageBasedCreditsUsed
}

func (s *SubscriptionUsageResponseModel) SetActualReportedCredits(actualReportedCredits int) {
	s.ActualReportedCredits = actualReportedCredits
}

func (s *SubscriptionUsageResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubscriptionUsageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubscriptionUsageResponseModel to string"
	}
	return string(jsonData)
}

func (s *SubscriptionUsageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubscriptionUsageResponseModel) MarshalJSON() ([]byte, error) {
	type embed SubscriptionUsageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
