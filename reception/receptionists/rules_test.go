package receptionists

import (
	"testing"
)

func TestReceptionCreateAgentRuleRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "priority": "strict"
}`
	assertRoundTrip(t, wire, new(ReceptionCreateAgentRuleRequest))
}

func TestReceptionUpdateAgentRuleRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "priority": "strict"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateAgentRuleRequest))
}
