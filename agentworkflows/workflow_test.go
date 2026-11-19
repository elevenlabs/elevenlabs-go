package agentworkflows_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/types"
)

func codesOf(t *testing.T, err error) []string {
	t.Helper()
	var verr *agentworkflows.ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected a *ValidationError, got %v", err)
	}
	return verr.Codes()
}

func contains(codes []string, code agentworkflows.IssueCode) bool {
	for _, c := range codes {
		if c == string(code) {
			return true
		}
	}
	return false
}

func llm(text string) *agentworkflows.Condition {
	c := agentworkflows.ConditionLLM(text, nil)
	return &c
}

// The spec's one-chain example, using Up() to get back to triage.
func TestSpecChainExample(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
		Tool("lookup_order", []string{"tool_get_order"}, nil).
		OnResult(agentworkflows.ResultBranches{
			Success: func(ok *agentworkflows.Branch) {
				ok.Converse("resolve", "Explain the order status", nil).End(nil)
			},
			Failure: func(failed *agentworkflows.Branch) { failed.GoTo("triage", nil) },
		}).
		Up().
		When(agentworkflows.ConditionLLM("Caller wants a human", nil), nil).
		Transfer("human", agentworkflows.DestinationPhone("+15551234567"), nil)

	out, err := wf.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	nodes := out["nodes"].(map[string]any)
	triage := nodes["triage"].(map[string]any)
	want := []any{"triage__lookup_order", "triage__human"}
	if got := triage["edge_order"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("triage edge_order = %v, want %v", got, want)
	}
	// lookup_order.OnFailure().GoTo("triage") loops back along triage → lookup_order.
	edges := out["edges"].(map[string]any)
	if _, ok := edges["lookup_order__triage"]; ok {
		t.Fatal("the failure loop must be merged into triage__lookup_order")
	}
	back := edges["triage__lookup_order"].(map[string]any)["backward_condition"]
	if !reflect.DeepEqual(back, map[string]any{"type": "result", "successful": false}) {
		t.Fatalf("backward_condition = %v", back)
	}

	model, err := wf.Build(nil)
	if err != nil {
		t.Fatal(err)
	}
	if model.Nodes["human"].PhoneNumber == nil || model.Nodes["triage"].OverrideAgent == nil {
		t.Fatal("Build must return typed nodes")
	}
	if _, err := json.Marshal(model); err != nil {
		t.Fatalf("the request model must serialize: %v", err)
	}
}

func TestBackwardConditionIntoStartIsRejected(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.When(agentworkflows.ConditionLLM("Begin", nil), &agentworkflows.WhenOptions{Back: llm("Restart")}).
		Converse("greet", "Hello", nil)
	_, err := wf.Build(nil)
	if !contains(codesOf(t, err), agentworkflows.EdgeIntoStart) {
		t.Fatalf("expected edge_into_start, got %v", err)
	}
}

func TestBackwardConditionOutOfEndIsRejected(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Triage", nil).
		When(agentworkflows.ConditionLLM("Done", nil), &agentworkflows.WhenOptions{Back: llm("Resume")}).
		End(nil)
	_, err := wf.Build(nil)
	if !contains(codesOf(t, err), agentworkflows.EndHasOutgoingEdges) {
		t.Fatalf("expected end_has_outgoing_edges, got %v", err)
	}
}

func TestBackwardLoopBetweenSubagentsIsAllowed(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("a", "A", nil).
		When(agentworkflows.ConditionLLM("To B", nil), &agentworkflows.WhenOptions{Back: llm("Back to A")}).
		Converse("b", "B", nil).
		End(nil)
	if _, err := wf.Build(nil); err != nil {
		t.Fatal(err)
	}
}

func TestLoadedNullBackwardConditionIsNoTransition(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("greet", "Hi", nil).End(nil)
	out, err := wf.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range out["edges"].(map[string]any) {
		e.(map[string]any)["backward_condition"] = nil
	}
	loaded, err := agentworkflows.FromJSON(out)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := loaded.Build(nil); err != nil {
		t.Fatalf("an explicit null backward_condition must not invent a reverse transition: %v", err)
	}
}

func TestInsertBetweenBidirectionalEdgeIsRejected(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("a", "A", nil).
		When(agentworkflows.ConditionLLM("to b", nil), &agentworkflows.WhenOptions{Back: llm("back to a")}).
		Converse("b", "B", nil)
	got := wf.InsertBetween("a", "b", func(from *agentworkflows.Branch) *agentworkflows.Node {
		return from.Converse("mid", "M", nil)
	})
	if got != nil {
		t.Fatal("inserting into a bidirectional edge must return nil")
	}
	_, err := wf.Build(nil)
	if !contains(codesOf(t, err), agentworkflows.CannotInsertBidirectional) {
		t.Fatalf("expected cannot_insert_bidirectional, got %v", err)
	}
}

func TestPositionIsAcceptedForNodeKindsWithoutAFieldMap(t *testing.T) {
	wf, err := agentworkflows.FromJSON(`{
		"nodes": {
			"start_node": {"type": "start"},
			"announce": {"type": "say", "message": {"type": "literal", "text": "hello"}}
		},
		"edges": {
			"start_node__announce": {"source": "start_node", "target": "announce", "forward_condition": {"type": "unconditional"}}
		}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	x, y := 10.0, 20.0
	wf.Node("announce").Update(&agentworkflows.Options{Position: &types.PositionInput{X: &x, Y: &y}})
	out, err := wf.ToJSON()
	if err != nil {
		t.Fatalf("position must be accepted on a say node: %v", err)
	}
	position := out["nodes"].(map[string]any)["announce"].(map[string]any)["position"]
	if !reflect.DeepEqual(position, map[string]any{"x": json.Number("10"), "y": json.Number("20")}) {
		t.Fatalf("position = %v", position)
	}
}

func TestLoadedBackwardOnlyEdgeHasNoForwardTransition(t *testing.T) {
	wf, err := agentworkflows.FromJSON(`{
		"nodes": {
			"start_node": {"type": "start"},
			"hub": {"type": "override_agent", "label": "hub", "additional_prompt": "H"},
			"x": {"type": "override_agent", "label": "x", "additional_prompt": "X"},
			"end": {"type": "end"}
		},
		"edges": {
			"e1": {"source": "start_node", "target": "hub", "forward_condition": {"type": "llm", "condition": "go"}},
			"e2": {"source": "hub", "target": "x", "forward_condition": {"type": "llm", "condition": "go"}},
			"e3": {"source": "hub", "target": "end", "forward_condition": {"type": "llm", "condition": "go"}},
			"e4": {"source": "x", "target": "start_node", "backward_condition": {"type": "llm", "condition": "restart"}}
		}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := wf.ToJSON()
	if err != nil {
		t.Fatalf("a backward-only edge must not trip the forward routing checks: %v", err)
	}
	nodes := out["nodes"].(map[string]any)
	if order := nodes["x"].(map[string]any)["edge_order"]; !reflect.DeepEqual(order, []any{}) {
		t.Fatalf("x has no forward transition via e4: edge_order = %v", order)
	}
	if order := nodes["start_node"].(map[string]any)["edge_order"]; !reflect.DeepEqual(order, []any{"e1", "e4"}) {
		t.Fatalf("start_node leaves by e1 and by e4's backward condition: edge_order = %v", order)
	}
}

func TestMalformedLoadedFieldIsInvalidModel(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("greet", "Hi", nil).End(nil)
	out, err := wf.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range out["edges"].(map[string]any) {
		e.(map[string]any)["reroutes"] = map[string]any{"x": 1, "y": 2}
	}
	loaded, err := agentworkflows.FromJSON(out)
	if err != nil {
		t.Fatal(err)
	}
	_, err = loaded.Build(nil)
	if codes := codesOf(t, err); !reflect.DeepEqual(codes, []string{"invalid_model"}) {
		t.Fatalf("codes = %v", codes)
	}
	// ToJSONString doesn't type the output, so it still round-trips the field.
	if _, err := loaded.ToJSONString(); err != nil {
		t.Fatal(err)
	}
}

// A node type this SDK doesn't model is kept exactly as loaded, nulls included.
func TestUnknownNodeTypeIsAWarningAndRoundTrips(t *testing.T) {
	text := `{
  "edges": {
    "start_node__survey": {
      "forward_condition": {
        "type": "unconditional"
      },
      "source": "start_node",
      "target": "survey"
    }
  },
  "nodes": {
    "start_node": {
      "edge_order": [
        "start_node__survey"
      ],
      "position": {
        "x": 0,
        "y": 0
      },
      "type": "start"
    },
    "survey": {
      "edge_order": [],
      "position": {
        "x": 300,
        "y": 0
      },
      "retry": null,
      "survey_id": "srv_1",
      "type": "future_survey"
    }
  }
}
`
	wf, err := agentworkflows.FromJSON(text)
	if err != nil {
		t.Fatal(err)
	}
	var warnings []agentworkflows.Issue
	model, err := wf.Build(&agentworkflows.BuildOptions{OnWarning: func(issue agentworkflows.Issue) {
		warnings = append(warnings, issue)
	}})
	if err != nil {
		t.Fatal(err)
	}
	if len(warnings) != 1 || warnings[0].Code != agentworkflows.UnknownNodeType || warnings[0].NodeID != "survey" {
		t.Fatalf("warnings = %+v", warnings)
	}
	if model.Nodes["survey"].GetType() != "future_survey" {
		t.Fatalf("survey type = %q", model.Nodes["survey"].GetType())
	}
	got, err := wf.ToJSONString()
	if err != nil {
		t.Fatal(err)
	}
	if got != text {
		t.Fatalf("round trip changed the document:\n%s", got)
	}
}

func TestUnreachableNodeWarning(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Converse("a", "A", nil)
	wf.Node("a").When(agentworkflows.ConditionLLM("x", nil), nil).Converse("b", "B", nil)
	wf.Remove("a")
	var codes []agentworkflows.IssueCode
	if _, err := wf.Build(&agentworkflows.BuildOptions{OnWarning: func(i agentworkflows.Issue) { codes = append(codes, i.Code) }}); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(codes, []agentworkflows.IssueCode{agentworkflows.UnreachableNode}) {
		t.Fatalf("warnings = %v", codes)
	}
}

// Every fixture's expected output is a fixed point: FromJSON(s).ToJSONString() == s.
func TestFixtureOutputsRoundTrip(t *testing.T) {
	paths, _ := filepath.Glob(filepath.Join("testdata", "fixtures", "*.json"))
	checked := 0
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var fx fixture
		if err := json.Unmarshal(data, &fx); err != nil {
			t.Fatal(err)
		}
		if len(fx.Expected) == 0 {
			continue
		}
		first, err := agentworkflows.FromJSON(fx.Expected)
		if err != nil {
			t.Fatal(err)
		}
		s, err := first.ToJSONString()
		if err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		second, err := agentworkflows.FromJSON(s)
		if err != nil {
			t.Fatal(err)
		}
		again, err := second.ToJSONString()
		if err != nil {
			t.Fatal(err)
		}
		if again != s {
			t.Fatalf("%s: round trip changed the document", path)
		}
		checked++
	}
	if checked == 0 {
		t.Fatal("no fixture outputs checked")
	}
}

// Node order, the start node and unlisted edges follow document order, as JavaScript reads it.
func TestLoadKeepsDocumentOrder(t *testing.T) {
	wf, err := agentworkflows.FromJSON(`{
		"nodes": {
			"zeta": {"type": "override_agent", "label": "zeta", "additional_prompt": "Z"},
			"begin": {"type": "start"},
			"alpha": {"type": "override_agent", "label": "alpha", "additional_prompt": "A"}
		},
		"edges": {
			"to_zeta": {"source": "begin", "target": "zeta", "forward_condition": {"type": "llm", "condition": "z"}},
			"to_alpha": {"source": "begin", "target": "alpha", "forward_condition": {"type": "llm", "condition": "a"}}
		}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	if got := wf.NodeIDs(); !reflect.DeepEqual(got, []string{"zeta", "begin", "alpha"}) {
		t.Fatalf("NodeIDs = %v", got)
	}
	if wf.StartID() != "begin" {
		t.Fatalf("StartID = %q", wf.StartID())
	}
	out, err := wf.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	nodes := out["nodes"].(map[string]any)
	if order := nodes["begin"].(map[string]any)["edge_order"]; !reflect.DeepEqual(order, []any{"to_zeta", "to_alpha"}) {
		t.Fatalf("edge_order = %v", order)
	}
	// BFS visits zeta first, so it takes row 0 of column 1.
	if pos := nodes["zeta"].(map[string]any)["position"]; !reflect.DeepEqual(pos, map[string]any{"x": json.Number("300"), "y": json.Number("0")}) {
		t.Fatalf("zeta position = %v", pos)
	}
}

func TestNumbersAreWrittenAsJavaScriptWritesThem(t *testing.T) {
	wf, err := agentworkflows.FromJSON(`{
		"nodes": {
			"start_node": {"type": "start", "position": {"x": 300.0, "y": 9007199254740993}},
			"a": {"type": "override_agent", "label": "a", "additional_prompt": "A", "position": {"x": 1E-7, "y": 1e21}}
		},
		"edges": {}
	}`)
	if err != nil {
		t.Fatal(err)
	}
	s, err := wf.ToJSONString()
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"x": 300,`, `"y": 9007199254740993`, `"x": 1e-7,`, `"y": 1e+21`} {
		if !strings.Contains(s, want) {
			t.Errorf("output lacks %s:\n%s", want, s)
		}
	}
}

// A get-agent response holds every unset field as null (and the API's defaults for the rest).
// FromAgent reads the nulls as absent and keeps everything else.
func TestFromAgentLoadsAGetAgentResponse(t *testing.T) {
	response := `{
		"prevent_subagent_loops": false,
		"nodes": {
			"start_node": {"type": "start", "position": {"x": 0, "y": 0}, "edge_order": ["start_node__greet"]},
			"greet": {"type": "override_agent", "position": {"x": 300, "y": 0}, "edge_order": ["greet__end"],
				"conversation_config": null, "additional_prompt": "Hi", "additional_knowledge_base": [],
				"additional_tool_ids": [], "label": "greet", "entry_behavior": "auto", "allow_interruptions_on_entry": false},
			"end": {"type": "end", "position": {"x": 600, "y": 0}, "edge_order": [], "outcome": "success"}
		},
		"edges": {
			"start_node__greet": {"source": "start_node", "target": "greet", "forward_condition": {"type": "unconditional", "label": null}, "backward_condition": null, "reroutes": []},
			"greet__end": {"source": "greet", "target": "end", "forward_condition": {"type": "llm", "condition": "Done", "label": null}, "backward_condition": null, "reroutes": []}
		}
	}`
	var workflow types.AgentWorkflowResponseModel
	if err := json.Unmarshal([]byte(response), &workflow); err != nil {
		t.Fatal(err)
	}
	var want string
	for _, agent := range []any{&types.GetAgentResponseModel{AgentId: "agent_1", Workflow: &workflow}, &workflow} {
		loaded, err := agentworkflows.FromAgent(agent)
		if err != nil {
			t.Fatal(err)
		}
		got, err := loaded.ToJSONString()
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(got, "null") {
			t.Fatalf("FromAgent(%T) kept a null:\n%s", agent, got)
		}
		if !strings.Contains(got, `"allow_interruptions_on_entry": false`) || !strings.Contains(got, `"reroutes": []`) {
			t.Fatalf("FromAgent(%T) dropped loaded values:\n%s", agent, got)
		}
		if want == "" {
			want = got
		} else if got != want {
			t.Fatalf("FromAgent(%T) differs:\n%s\nwant:\n%s", agent, got, want)
		}
		if _, err := loaded.Build(nil); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := agentworkflows.FromAgent("not an agent"); err == nil {
		t.Fatal("FromAgent must reject other types")
	}
}

func TestExprBuilders(t *testing.T) {
	expr := agentworkflows.ExprAllOf(
		agentworkflows.ExprVariable("tier").Eq("gold"),
		agentworkflows.ExprLLM("How satisfied is the caller?", &agentworkflows.LLMValueOptions{Type: agentworkflows.LLMNumber}).Lt(5),
		agentworkflows.ExprIfElse(true, 1.5, nil),
	)
	want := map[string]any{
		"type": "and_operator",
		"children": []any{
			map[string]any{"type": "eq_operator", "left": map[string]any{"type": "dynamic_variable", "name": "tier"}, "right": map[string]any{"type": "string_literal", "value": "gold"}},
			map[string]any{"type": "lt_operator", "left": map[string]any{"type": "llm", "value_schema": map[string]any{"type": "number", "description": "How satisfied is the caller?"}}, "right": map[string]any{"type": "number_literal", "value": json.Number("5")}},
			map[string]any{"type": "conditional_operator", "condition": map[string]any{"type": "boolean_literal", "value": true}, "trueExpression": map[string]any{"type": "number_literal", "value": json.Number("1.5")}, "falseExpression": map[string]any{"type": "null_literal"}},
		},
	}
	if got := expr.Wire(); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v", got)
	}
}

// A response-shaped workflow's nulls are read as absent, at every level of a modelled node.
func TestFromJSONReadsNullsAsAbsent(t *testing.T) {
	wf, err := agentworkflows.FromJSON(`{"workflow": {
		"prevent_subagent_loops": null,
		"nodes": {
			"start_node": {"type": "start", "position": null, "edge_order": ["e1"]},
			"lookup": {"type": "tool", "tools": [{"tool_id": "t1", "schema_overrides": {"a": {"source": "constant", "constant_value": null, "description": null}}}]}
		},
		"edges": {
			"e1": {"source": "start_node", "target": "lookup", "forward_condition": {"type": "unconditional", "label": null}, "backward_condition": null, "reroutes": null}
		}
	}}`)
	if err != nil {
		t.Fatal(err)
	}
	s, err := wf.ToJSONString()
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Count(s, "null"); got != 1 || !strings.Contains(s, `"constant_value": null`) {
		t.Fatalf("only the constant's null is data:\n%s", s)
	}
}
