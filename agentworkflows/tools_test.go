package agentworkflows_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// namedTool stands in for an agenttools tool: anything with a name.
type namedTool struct{ name string }

func (t *namedTool) Name() string { return t.name }

func tool(name string) *namedTool { return &namedTool{name: name} }

func constantOverride(t *testing.T, value string) *types.WorkflowToolLocatorSchemaOverridesValue {
	t.Helper()
	var override types.WorkflowToolLocatorSchemaOverridesValue
	if err := json.Unmarshal([]byte(`{"source": "constant", "constant_value": "`+value+`"}`), &override); err != nil {
		t.Fatal(err)
	}
	return &override
}

func nodeOf(t *testing.T, out map[string]any, id string) map[string]any {
	t.Helper()
	node, ok := out["nodes"].(map[string]any)[id].(map[string]any)
	if !ok {
		t.Fatalf("no node %q in %v", id, out)
	}
	return node
}

func TestToolObjectsResolveToTheirIDs(t *testing.T) {
	lookup, show := tool("get_order"), tool("show_order")
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
		Tool("lookup", nil, &agentworkflows.Options{
			Tools: []agentworkflows.Tool{lookup},
			SchemaOverrides: map[string]map[string]*types.WorkflowToolLocatorSchemaOverridesValue{
				"get_order": {"path_params.order_id": constantOverride(t, "1")},
			},
		}).
		OnSuccess(nil).
		Converse("resolve", "Explain the order", &agentworkflows.Options{ToolIDs: []string{"tool_raw"}, Tools: []agentworkflows.Tool{show}})

	if got := wf.Tools(); !reflect.DeepEqual(got, []agentworkflows.Tool{lookup, show}) {
		t.Fatalf("Tools() = %v", got)
	}
	out, err := wf.ToJSON(&agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_g", "show_order": "tool_s"}})
	if err != nil {
		t.Fatal(err)
	}
	tools, _ := json.Marshal(nodeOf(t, out, "lookup")["tools"])
	if string(tools) != `[{"schema_overrides":{"path_params.order_id":{"constant_value":"1","source":"constant"}},"tool_id":"tool_g"}]` {
		t.Fatalf("lookup tools = %s", tools)
	}
	if got := nodeOf(t, out, "resolve")["additional_tool_ids"]; !reflect.DeepEqual(got, []any{"tool_raw", "tool_s"}) {
		t.Fatalf("resolve additional_tool_ids = %v", got)
	}

	// A Tool without an id is unresolved, in Build, ToJSON and ToJSONString alike.
	partial := &agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_g"}}
	if _, err := wf.Build(partial); !reflect.DeepEqual(codesOf(t, err), []string{"unresolved_tool"}) {
		t.Fatalf("Build: %v", err)
	}
	if _, err := wf.ToJSONString(nil); !reflect.DeepEqual(codesOf(t, err), []string{"unresolved_tool", "unresolved_tool"}) {
		t.Fatalf("ToJSONString: %v", err)
	}
	if _, err := wf.Build(&agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_g", "show_order": "tool_s"}}); err != nil {
		t.Fatalf("Build with every id: %v", err)
	}
}

func TestToolObjectsInToJSONString(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order")}})
	opts := &agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}}
	text, err := wf.ToJSONString(opts)
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := agentworkflows.FromJSON(text)
	if err != nil {
		t.Fatal(err)
	}
	out, err := loaded.ToJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := nodeOf(t, out, "lookup")["tools"]; !reflect.DeepEqual(got, []any{map[string]any{"tool_id": "tool_1"}}) {
		t.Fatalf("tools = %v", got)
	}
}

func TestDuplicateToolNames(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order"), tool("get_order")}})
	ids := &agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}}
	if _, err := wf.Build(ids); !reflect.DeepEqual(codesOf(t, err), []string{"duplicate_tool_name"}) {
		t.Fatalf("Build: %v", err)
	}

	// The same object twice is fine.
	same := tool("get_order")
	ok := agentworkflows.New(nil)
	ok.Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{same, same}})
	if _, err := ok.Build(ids); err != nil {
		t.Fatal(err)
	}

	// Another node still using a different tool of that name conflicts.
	two := agentworkflows.New(nil)
	two.Tool("a", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order")}}).
		OnSuccess(nil).Tool("b", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order")}})
	if _, err := two.Build(ids); !reflect.DeepEqual(codesOf(t, err), []string{"duplicate_tool_name"}) {
		t.Fatalf("Build: %v", err)
	}

	// A removed node's tool name can be reused.
	reused := agentworkflows.New(nil)
	reused.Tool("a", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order")}})
	reused.Remove("a")
	reused.Tool("b", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tool("get_order")}})
	if _, err := reused.Build(ids); err != nil {
		t.Fatal(err)
	}
	if len(reused.Tools()) != 1 {
		t.Fatalf("Tools() = %v", reused.Tools())
	}
}

func TestSchemaOverridesForAToolNotRun(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Tool("lookup", nil, &agentworkflows.Options{
		Tools: []agentworkflows.Tool{tool("get_order")},
		SchemaOverrides: map[string]map[string]*types.WorkflowToolLocatorSchemaOverridesValue{
			"other": {"x": constantOverride(t, "1")},
		},
	})
	_, err := wf.Build(&agentworkflows.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}})
	if !reflect.DeepEqual(codesOf(t, err), []string{"invalid_option"}) {
		t.Fatalf("Build: %v", err)
	}
}

func TestFromJSONWithTools(t *testing.T) {
	refund := tool("issue_refund")
	wf, err := agentworkflows.FromJSONWithTools(`{
		"nodes": {
			"start_node": {"type": "start"},
			"refunds": {"type": "override_agent", "label": "Refunds", "additional_tool_ids": ["tool_refund", "tool_other"]},
			"run": {"type": "tool", "tools": [{"tool_id": "tool_refund"}]}
		},
		"edges": {
			"e1": {"source": "start_node", "target": "refunds", "forward_condition": {"type": "unconditional"}},
			"e2": {"source": "refunds", "target": "run", "forward_condition": {"type": "llm", "condition": "Refund"}}
		}
	}`, map[string]agentworkflows.Tool{"tool_refund": refund})
	if err != nil {
		t.Fatal(err)
	}
	if got := wf.Tools(); !reflect.DeepEqual(got, []agentworkflows.Tool{refund}) {
		t.Fatalf("Tools() = %v", got)
	}
	out, err := wf.ToJSON(&agentworkflows.BuildOptions{ToolIDs: map[string]string{"issue_refund": "tool_new"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := nodeOf(t, out, "refunds")["additional_tool_ids"]; !reflect.DeepEqual(got, []any{"tool_new", "tool_other"}) {
		t.Fatalf("additional_tool_ids = %v", got)
	}
	if got := nodeOf(t, out, "run")["tools"]; !reflect.DeepEqual(got, []any{map[string]any{"tool_id": "tool_new"}}) {
		t.Fatalf("tools = %v", got)
	}
}
