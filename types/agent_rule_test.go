package types

import (
	"testing"
)

func TestAgentRuleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "priority": "strict",
  "id": "value",
  "created_at_unix": 123
}`
	assertRoundTrip(t, wire, new(AgentRule))
}

func TestAgentRuleRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "priority": "strict",
  "id": "value",
  "created_at_unix": 123
}`
	assertRoundTrip(t, wire, new(AgentRule))
}
