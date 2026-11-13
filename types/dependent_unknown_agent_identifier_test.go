package types

import (
	"testing"
)

func TestDependentUnknownAgentIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}

func TestDependentUnknownAgentIdentifierRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}

func TestDependentUnknownAgentIdentifierRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}

func TestDependentUnknownAgentIdentifierRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}

func TestDependentUnknownAgentIdentifierRoundTrip5(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}

func TestDependentUnknownAgentIdentifierRoundTrip6(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(DependentUnknownAgentIdentifier))
}
