package agentworkflows

import (
	"reflect"
	"testing"
)

// The node types read from the generated union. If a regeneration adds one, this list grows on
// its own; update the expectation then.
func TestGeneratedNodeTypes(t *testing.T) {
	want := map[string]bool{
		"end": true, "override_agent": true, "phone_number": true, "say": true, "standalone_agent": true,
		"start": true, "start_procedure": true, "tool": true, "update_state": true,
	}
	if got := generatedNodeTypes(); !reflect.DeepEqual(got, want) {
		t.Fatalf("generated node types = %v", got)
	}
	if modelKnowsNodeType("future_survey") || modelKnowsNodeType("") {
		t.Fatal("unknown types must not be known")
	}
}
