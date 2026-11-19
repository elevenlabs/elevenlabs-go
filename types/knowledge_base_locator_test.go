package types

import (
	"testing"
)

func TestKnowledgeBaseLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "name": "value",
  "id": "value",
  "usage_mode": "prompt"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseLocator))
}

func TestKnowledgeBaseLocatorRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "name": "value",
  "id": "value",
  "usage_mode": "prompt"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseLocator))
}

func TestKnowledgeBaseLocatorRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "name": "value",
  "id": "value",
  "usage_mode": "prompt"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseLocator))
}

func TestKnowledgeBaseLocatorRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "name": "value",
  "id": "value",
  "usage_mode": "prompt"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseLocator))
}
