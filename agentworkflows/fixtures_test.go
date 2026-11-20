package agentworkflows_test

// Cross-SDK conformance: runs every shared fixture (vendored from elevenlabs-js
// docs/agents-as-code/fixtures into testdata/fixtures) through the Go builder and compares
// ToJSONString(nil) byte for byte, or the sorted error codes from Build(). Mirrors the JS runner in
// tests/custom/agents/workflow/fixtureRunner.ts.

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

const fixtureCount = 10

type fixture struct {
	Name     string `json:"name"`
	Workflow *struct {
		PreventSubagentLoops *bool `json:"preventSubagentLoops"`
	} `json:"workflow"`
	From     json.RawMessage `json:"from"`
	Ops      []fixtureOp     `json:"ops"`
	Expected json.RawMessage `json:"expected"`
	Errors   []string        `json:"errors"`
}

type fixtureOp struct {
	Op          string                     `json:"op"`
	At          *string                    `json:"at"`
	When        json.RawMessage            `json:"when"`
	Back        json.RawMessage            `json:"back"`
	OnSuccess   bool                       `json:"onSuccess"`
	OnFailure   bool                       `json:"onFailure"`
	ID          string                     `json:"id"`
	Prompt      string                     `json:"prompt"`
	ToolIDs     []string                   `json:"toolIds"`
	Destination json.RawMessage            `json:"destination"`
	Target      string                     `json:"target"`
	Node        string                     `json:"node"`
	Source      string                     `json:"source"`
	Options     map[string]json.RawMessage `json:"options"`
	Ops         []fixtureOp                `json:"ops"`
	Success     *[]fixtureOp               `json:"success"`
	Failure     *[]fixtureOp               `json:"failure"`
}

func TestFixtures(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("testdata", "fixtures", "*.json"))
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(paths)
	if len(paths) != fixtureCount {
		t.Fatalf("expected %d fixtures, found %d", fixtureCount, len(paths))
	}
	for _, path := range paths {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var fx fixture
			if err := json.Unmarshal(data, &fx); err != nil {
				t.Fatal(err)
			}
			got, codes := runFixture(t, &fx)
			switch {
			case len(fx.Expected) > 0:
				expected, err := wirejson.Decode(fx.Expected)
				if err != nil {
					t.Fatal(err)
				}
				want, err := wirejson.String(expected)
				if err != nil {
					t.Fatal(err)
				}
				if codes != nil {
					t.Fatalf("%s: expected JSON but build failed: %v", fx.Name, codes)
				}
				if got != want {
					t.Fatalf("%s: JSON mismatch\n--- got ---\n%s\n--- want ---\n%s", fx.Name, got, want)
				}
			case fx.Errors != nil:
				want := append([]string{}, fx.Errors...)
				sort.Strings(want)
				if codes == nil {
					t.Fatalf("%s: expected errors %v but build succeeded", fx.Name, want)
				}
				if !reflect.DeepEqual(codes, want) {
					t.Fatalf("%s: error mismatch\n  got:  %v\n  want: %v", fx.Name, codes, want)
				}
			default:
				t.Fatalf("%s: fixture has neither expected nor errors", fx.Name)
			}
		})
	}
}

func runFixture(t *testing.T, fx *fixture) (string, []string) {
	t.Helper()
	var wf *agentworkflows.Workflow
	if len(fx.From) > 0 {
		loaded, err := agentworkflows.FromJSON(fx.From)
		if err != nil {
			t.Fatal(err)
		}
		wf = loaded
	} else {
		var opts agentworkflows.WorkflowOptions
		if fx.Workflow != nil {
			opts.PreventSubagentLoops = fx.Workflow.PreventSubagentLoops
		}
		wf = agentworkflows.New(&opts)
	}
	run(t, wf, fx.Ops, wf)
	out, err := wf.ToJSONString(nil)
	if err != nil {
		var verr *agentworkflows.ValidationError
		if !errors.As(err, &verr) {
			t.Fatal(err)
		}
		return "", verr.Codes()
	}
	return out, nil
}

// run executes ops starting at first (the workflow, or an insert/result branch) and returns the
// last op's node.
func run(t *testing.T, wf *agentworkflows.Workflow, ops []fixtureOp, first any) *agentworkflows.Node {
	t.Helper()
	previous := first
	var result *agentworkflows.Node
	for _, op := range ops {
		op := op
		switch op.Op {
		case "update":
			wf.Node(op.Node).Update(options(t, op))
			continue
		case "remove":
			wf.Remove(op.Node)
			continue
		case "onResult":
			node, ok := previous.(*agentworkflows.Node)
			if op.At != nil {
				node, ok = wf.Node(*op.At), true
			}
			if !ok {
				t.Fatal("onResult needs a node")
			}
			var branches agentworkflows.ResultBranches
			if op.Success != nil {
				branches.Success = func(b *agentworkflows.Branch) { run(t, wf, *op.Success, b) }
			}
			if op.Failure != nil {
				branches.Failure = func(b *agentworkflows.Branch) { run(t, wf, *op.Failure, b) }
			}
			node.OnResult(branches)
			continue
		case "insertBetween":
			wf.InsertBetween(op.Source, op.Target, func(b *agentworkflows.Branch) *agentworkflows.Node {
				return run(t, wf, op.Ops, b)
			})
			continue
		}

		at := previous
		if op.At != nil {
			if *op.At == "$start" {
				at = wf
			} else {
				at = wf.Node(*op.At)
			}
		}
		switch {
		case len(op.When) > 0:
			var whenOpts *agentworkflows.WhenOptions
			if len(op.Back) > 0 {
				back := condition(t, op.Back)
				whenOpts = &agentworkflows.WhenOptions{Back: &back}
			}
			switch base := at.(type) {
			case *agentworkflows.Workflow:
				at = base.When(condition(t, op.When), whenOpts)
			case *agentworkflows.Node:
				at = base.When(condition(t, op.When), whenOpts)
			default:
				t.Fatalf("when() needs a workflow or node, not %T", at)
			}
		case op.OnSuccess:
			at = at.(*agentworkflows.Node).OnSuccess(nil)
		case op.OnFailure:
			at = at.(*agentworkflows.Node).OnFailure(nil)
		}

		chain := at.(agentworkflows.Chain)
		switch op.Op {
		case "converse":
			result = chain.Converse(op.ID, op.Prompt, options(t, op))
		case "tool":
			result = chain.Tool(op.ID, op.ToolIDs, options(t, op))
		case "transfer":
			result = chain.Transfer(op.ID, destination(t, op.Destination), options(t, op))
		case "end":
			result = chain.End(options(t, op))
		case "goTo":
			var goOpts *agentworkflows.GoToOptions
			if len(op.Back) > 0 {
				back := condition(t, op.Back)
				goOpts = &agentworkflows.GoToOptions{Back: &back}
			}
			result = chain.GoTo(op.Target, goOpts)
		default:
			t.Fatalf("unknown op %q", op.Op)
		}
		previous = result
	}
	return result
}

func condition(t *testing.T, raw json.RawMessage) agentworkflows.Condition {
	t.Helper()
	var text string
	if json.Unmarshal(raw, &text) == nil {
		return agentworkflows.ConditionLLM(text, nil)
	}
	var obj struct {
		LLM   *string         `json:"llm"`
		Expr  json.RawMessage `json:"expr"`
		Label string          `json:"label"`
	}
	if err := json.Unmarshal(raw, &obj); err != nil {
		t.Fatal(err)
	}
	opts := &agentworkflows.ConditionOptions{Label: obj.Label}
	if obj.LLM != nil {
		return agentworkflows.ConditionLLM(*obj.LLM, opts)
	}
	expr, err := agentworkflows.ExprFromModel(obj.Expr)
	if err != nil {
		t.Fatal(err)
	}
	return agentworkflows.ConditionExpr(expr, opts)
}

func destination(t *testing.T, raw json.RawMessage) agentworkflows.Destination {
	t.Helper()
	var d struct {
		Phone         *string `json:"phone"`
		PhoneVariable *string `json:"phoneVariable"`
		SIP           *string `json:"sip"`
		SIPVariable   *string `json:"sipVariable"`
		Agent         string  `json:"agent"`
		NodeID        string  `json:"nodeId"`
	}
	if err := json.Unmarshal(raw, &d); err != nil {
		t.Fatal(err)
	}
	switch {
	case d.Phone != nil:
		return agentworkflows.DestinationPhone(*d.Phone)
	case d.PhoneVariable != nil:
		return agentworkflows.DestinationPhoneVariable(*d.PhoneVariable)
	case d.SIP != nil:
		return agentworkflows.DestinationSIP(*d.SIP)
	case d.SIPVariable != nil:
		return agentworkflows.DestinationSIPVariable(*d.SIPVariable)
	}
	return agentworkflows.DestinationAgent(d.Agent, &agentworkflows.AgentDestinationOptions{NodeID: d.NodeID})
}

// options converts a fixture's canonical camelCase option keys to the Options struct, whose json
// tags are the snake_case names. Nested values are already in wire format.
func options(t *testing.T, op fixtureOp) *agentworkflows.Options {
	t.Helper()
	if op.Options == nil {
		return nil
	}
	snake := map[string]json.RawMessage{}
	var dest *agentworkflows.Destination
	for key, value := range op.Options {
		if key == "destination" {
			d := destination(t, value)
			dest = &d
			continue
		}
		snake[camelToSnake(key)] = value
	}
	data, err := json.Marshal(snake)
	if err != nil {
		t.Fatal(err)
	}
	var opts agentworkflows.Options
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // an option the struct doesn't model must fail loudly, not vanish
	if err := dec.Decode(&opts); err != nil {
		t.Fatal(err)
	}
	opts.Destination = dest
	return &opts
}

func camelToSnake(key string) string {
	var b strings.Builder
	for _, r := range key {
		if unicode.IsUpper(r) {
			b.WriteByte('_')
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
