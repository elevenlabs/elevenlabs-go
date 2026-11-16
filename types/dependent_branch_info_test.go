package types

import (
	"testing"
)

func TestDependentBranchInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "agent_name": "value",
  "branch_id": "value",
  "branch_name": "value",
  "is_main": true
}`
	assertRoundTrip(t, wire, new(DependentBranchInfo))
}

func TestDependentBranchInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "agent_name": "value",
  "branch_id": "value",
  "branch_name": "value",
  "is_main": true
}`
	assertRoundTrip(t, wire, new(DependentBranchInfo))
}

func TestDependentBranchInfoRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "agent_name": "value",
  "branch_id": "value",
  "branch_name": "value",
  "is_main": true
}`
	assertRoundTrip(t, wire, new(DependentBranchInfo))
}

func TestDependentBranchInfoRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "agent_name": "value",
  "branch_id": "value",
  "branch_name": "value",
  "is_main": true
}`
	assertRoundTrip(t, wire, new(DependentBranchInfo))
}
