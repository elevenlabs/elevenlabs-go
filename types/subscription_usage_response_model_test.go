package types

import (
	"testing"
)

func TestSubscriptionUsageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "rollover_credits_quota": 123,
  "subscription_cycle_credits_quota": 123,
  "manually_gifted_credits_quota": 123,
  "payg_credits_quota": 123,
  "rollover_credits_used": 123,
  "subscription_cycle_credits_used": 123,
  "manually_gifted_credits_used": 123,
  "payg_credits_used": 123,
  "paid_usage_based_credits_used": 123,
  "actual_reported_credits": 123
}`
	assertRoundTrip(t, wire, new(SubscriptionUsageResponseModel))
}
