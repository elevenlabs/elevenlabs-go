package types

import (
	"testing"
)

func TestDependentAvailableAgentIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}

func TestDependentAvailableAgentIdentifierRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}

func TestDependentAvailableAgentIdentifierRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}

func TestDependentAvailableAgentIdentifierRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}

func TestDependentAvailableAgentIdentifierRoundTrip5(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}

func TestDependentAvailableAgentIdentifierRoundTrip6(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(DependentAvailableAgentIdentifier))
}
