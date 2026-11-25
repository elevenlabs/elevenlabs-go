package types

import (
	"testing"
)

func TestResellerTierResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "display_name": "value",
  "currency": "USD",
  "tier_economics": {
    "monthly_price_usd": 9.99,
    "annual_price_usd": 9.99,
    "included_minutes_per_month": 9.99,
    "overage_price_per_minute_usd": 9.99,
    "max_phone_numbers": 123,
    "concurrent_calls": 123,
    "max_active_locations": 123,
    "max_customer_facing_agents": 123,
    "max_team_members": 123,
    "assistant_daily_messages": 123
  }
}`
	assertRoundTrip(t, wire, new(ResellerTierResponse))
}
