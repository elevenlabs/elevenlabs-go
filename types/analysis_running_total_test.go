package types

import (
	"testing"
)

func TestAnalysisRunningTotalRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "price": 9.99,
  "charge": 123,
  "runs": 123,
  "price_per_feature": {
    "key": 9.99
  },
  "charge_per_feature": {
    "key": 123
  }
}`
	assertRoundTrip(t, wire, new(AnalysisRunningTotal))
}
