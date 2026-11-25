package types

import (
	"testing"
)

func TestReceptionGroupMembershipRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "rank": 123
}`
	assertRoundTrip(t, wire, new(ReceptionGroupMembership))
}

func TestReceptionGroupMembershipRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "rank": 123
}`
	assertRoundTrip(t, wire, new(ReceptionGroupMembership))
}

func TestReceptionGroupMembershipRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "rank": 123
}`
	assertRoundTrip(t, wire, new(ReceptionGroupMembership))
}

func TestReceptionGroupMembershipRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "rank": 123
}`
	assertRoundTrip(t, wire, new(ReceptionGroupMembership))
}
