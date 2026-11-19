package types

import (
	"testing"
)

func TestKnowledgeBaseToolInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled_strategies": [
    "cat"
  ]
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseToolInfo))
}

func TestKnowledgeBaseToolInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled_strategies": [
    "cat"
  ]
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseToolInfo))
}
