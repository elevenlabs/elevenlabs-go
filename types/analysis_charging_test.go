package types

import (
	"testing"
)

func TestAnalysisChargingRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": {
    "price": 9.99,
    "charge": 123,
    "runs": 123,
    "price_per_feature": {
      "key": 9.99
    },
    "charge_per_feature": {
      "key": 123
    }
  },
  "last_run": {
    "price": 9.99,
    "charge": 123,
    "price_per_feature": {
      "key": 9.99
    },
    "charge_per_feature": {
      "key": 123
    }
  }
}`
	assertRoundTrip(t, wire, new(AnalysisCharging))
}
