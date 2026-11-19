package types

import (
	"testing"
)

func TestUserEvaluationAnalysisItemScoringRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "binary"
}`
	assertRoundTrip(t, wire, new(UserEvaluationAnalysisItemScoring))
}
