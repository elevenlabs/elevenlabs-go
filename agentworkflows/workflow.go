package agentworkflows

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Workflow is an agent workflow written as code. The workflow itself acts as the start node:
// calling a verb or When on it branches from the start.
//
// A Workflow and its handles are not safe for concurrent use.
type Workflow struct {
	g *graph
}

// New returns an empty workflow holding only the start node, with id "start_node" (the id the API
// requires). opts may be nil.
func New(opts *WorkflowOptions) *Workflow {
	g := newGraph()
	g.addNode(g.startID, map[string]any{"type": "start"})
	if opts != nil && opts.PreventSubagentLoops != nil {
		value := *opts.PreventSubagentLoops
		g.preventLoop = &value
	}
	return &Workflow{g: g}
}

// FromJSON loads a workflow stored as API wire-format JSON (snake_case keys), for example
// ToJSONString's output or a CLI agent_configs/*.json file. data is JSON text (a string, []byte
// or json.RawMessage) or a parsed value (such as a map[string]any), holding either the workflow
// itself or a whole agent config with a "workflow" key. Request- and response-shaped workflows are
// both accepted; unknown keys are kept as they are.
//
// Document order matters: it decides node order and the order of edges no node lists, so pass the
// JSON text where you have it. A parsed Go map has no order, and is read in sorted key order.
func FromJSON(data any) (*Workflow, error) {
	var raw []byte
	switch v := data.(type) {
	case string:
		raw = []byte(v)
	case []byte:
		raw = v
	case json.RawMessage:
		raw = v
	default:
		encoded, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("agentworkflows: can't convert %T to JSON: %w", data, err)
		}
		raw = encoded
	}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return nil, fmt.Errorf("agentworkflows: workflow JSON must be an object: %w", err)
	}
	if top == nil {
		return nil, errors.New("agentworkflows: workflow JSON must be an object, not null")
	}
	if workflow, ok := top["workflow"]; ok {
		raw = workflow
		if string(raw) == "null" {
			raw = []byte("{}")
		}
	}
	wf := New(nil)
	if err := wf.g.load(raw); err != nil {
		return nil, err
	}
	return wf, nil
}

// FromAgent loads the workflow of an existing agent so it can be edited. agent is the result of
// agents.Get (*types.GetAgentResponseModel), its Workflow field (*types.AgentWorkflowResponseModel)
// or a request model (*types.AgentWorkflowRequestModel). Node ids, editor positions and edge order
// are kept.
//
// The generated types hold nodes and edges in Go maps, which have no order, so they are read in
// sorted id order. The generated types also can't hold node types this SDK version doesn't
// model; for such a workflow FromAgent returns an error, and FromJSON on the raw JSON keeps them.
func FromAgent(agent any) (*Workflow, error) {
	var workflow any
	switch v := agent.(type) {
	case *types.GetAgentResponseModel:
		if v == nil || v.Workflow == nil {
			return New(nil), nil
		}
		workflow = v.Workflow
	case types.GetAgentResponseModel:
		if v.Workflow == nil {
			return New(nil), nil
		}
		workflow = v.Workflow
	case *types.AgentWorkflowResponseModel, types.AgentWorkflowResponseModel,
		*types.AgentWorkflowRequestModel, types.AgentWorkflowRequestModel:
		workflow = v
	default:
		return nil, fmt.Errorf("agentworkflows: FromAgent takes a get-agent response or its workflow, not %T", agent)
	}
	// A nil workflow pointer is an agent without a workflow, as a nil response is.
	if rv := reflect.ValueOf(workflow); rv.Kind() == reflect.Ptr && rv.IsNil() {
		return New(nil), nil
	}
	tree, err := workflowModelTree(workflow)
	if err != nil {
		return nil, err
	}
	// The response model writes every unset field as null; load reads those as absent.
	raw, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}
	wf := New(nil)
	if err := wf.g.load(raw); err != nil {
		return nil, err
	}
	return wf, nil
}

// workflowModelTree converts a generated workflow model (request or response) to a wire tree,
// node by node and edge by edge. A generated edge decoded from a null condition holds an empty
// union that can't be marshaled again; such a condition is read as absent.
func workflowModelTree(model any) (map[string]any, error) {
	v := reflect.Indirect(reflect.ValueOf(model))
	out := map[string]any{}
	if field := v.FieldByName("PreventSubagentLoops"); field.IsValid() {
		value, err := wirejson.Normalize(field.Interface())
		if err != nil {
			return nil, err
		}
		out["prevent_subagent_loops"] = value
	}
	nodes := map[string]any{}
	for iter := v.FieldByName("Nodes").MapRange(); iter.Next(); {
		id := iter.Key().String()
		value, err := wirejson.Normalize(iter.Value().Interface())
		if err != nil {
			return nil, fmt.Errorf("agentworkflows: can't read node %q of the agent's workflow (it may be a node type this SDK version doesn't model; load the raw JSON with FromJSON instead): %w", id, err)
		}
		nodes[id] = value
	}
	edges := map[string]any{}
	for iter := v.FieldByName("Edges").MapRange(); iter.Next(); {
		id := iter.Key().String()
		edge := iter.Value()
		if edge.Kind() == reflect.Ptr {
			if edge.IsNil() {
				continue
			}
			edge = edge.Elem()
		}
		copied := reflect.New(edge.Type())
		copied.Elem().Set(edge)
		for _, name := range []string{"ForwardCondition", "BackwardCondition"} {
			condition := copied.Elem().FieldByName(name)
			if condition.IsValid() && condition.Kind() == reflect.Ptr && !condition.IsNil() {
				if _, err := json.Marshal(condition.Interface()); err != nil {
					condition.Set(reflect.Zero(condition.Type()))
				}
			}
		}
		value, err := wirejson.Normalize(copied.Interface())
		if err != nil {
			return nil, fmt.Errorf("agentworkflows: can't read edge %q of the agent's workflow: %w", id, err)
		}
		edges[id] = value
	}
	out["nodes"] = nodes
	out["edges"] = edges
	return out, nil
}

// StartID is the start node's id: "start_node", unless the workflow was loaded with another.
func (w *Workflow) StartID() string { return w.g.startID }

// NodeIDs are the ids of every node, in insertion order.
func (w *Workflow) NodeIDs() []string { return append([]string{}, w.g.nodeOrder...) }

// Node returns a handle to an existing node, for editing it or going to it. An unknown id is
// reported by Build as unknown_node.
func (w *Workflow) Node(id string) *Node {
	if !w.g.hasNode(id) {
		w.g.issue(UnknownNode, fmt.Sprintf("There is no node %q", id), id)
	}
	return &Node{g: w.g, id: id}
}

// Converse adds a subagent node reached unconditionally from the start. See Chain.
func (w *Workflow) Converse(id, prompt string, opts *Options) *Node {
	return w.g.converse(unconditional(w.g.startID), id, prompt, opts)
}

// Tool adds a tool node reached unconditionally from the start. See Chain.
func (w *Workflow) Tool(id string, toolIDs []string, opts *Options) *Node {
	return w.g.tool(unconditional(w.g.startID), id, toolIDs, opts)
}

// Transfer adds a transfer node reached unconditionally from the start. See Chain.
func (w *Workflow) Transfer(id string, destination Destination, opts *Options) *Node {
	return w.g.transfer(unconditional(w.g.startID), id, destination, opts)
}

// End adds (or reuses) an end node reached unconditionally from the start. See Chain.
func (w *Workflow) End(opts *Options) *Node {
	return w.g.end(unconditional(w.g.startID), opts)
}

// GoTo connects the start node unconditionally to target. See Chain.
func (w *Workflow) GoTo(target string, opts *GoToOptions) *Node {
	return w.g.goTo(unconditional(w.g.startID), target, opts)
}

// When opens a branch from the start node, followed when condition holds.
func (w *Workflow) When(condition Condition, opts *WhenOptions) *Branch {
	return w.g.branch(w.g.startID, condition, whenBack(opts))
}

// Remove removes a node and every edge that touches it. Removing the start node is reported by
// Build as cannot_remove_start.
func (w *Workflow) Remove(id string) *Workflow {
	g := w.g
	if id == g.startID {
		g.issue(CannotRemoveStart, "The start node can't be removed", id)
		return w
	}
	if g.hasNode(id) {
		delete(g.nodes, id)
		order := g.nodeOrder[:0:0]
		for _, nodeID := range g.nodeOrder {
			if nodeID != id {
				order = append(order, nodeID)
			}
		}
		g.nodeOrder = order
	}
	edges := g.edges[:0:0]
	for _, e := range g.edges {
		if e.source != id && e.target != id {
			edges = append(edges, e)
		}
	}
	g.edges = edges
	return w
}

// InsertBetween replaces the edge sourceID → targetID with new nodes. insert receives a branch
// from sourceID that carries the removed edge's forward condition, and returns the last new node,
// which is then connected to targetID (with OnSuccess if it is a tool node, otherwise
// unconditionally). InsertBetween returns that node, or nil when there is no such edge, the edge
// also has a backward condition, or insert returns nil; Build reports why.
func (w *Workflow) InsertBetween(sourceID, targetID string, insert func(from *Branch) *Node) *Node {
	g := w.g
	index := -1
	for i, e := range g.edges {
		if e.source == sourceID && e.target == targetID {
			index = i
			break
		}
	}
	if index < 0 {
		g.issue(UnknownEdge, fmt.Sprintf("There is no edge from %q to %q to insert into", sourceID, targetID), sourceID)
		return nil
	}
	original := g.edges[index]
	if original.backward != nil {
		// Splitting a bidirectional edge can't keep the target's reverse route.
		g.issue(CannotInsertBidirectional, fmt.Sprintf("Edge %q → %q has a backward condition; inserting into a bidirectional edge would drop its reverse route. Remove the backward condition first", sourceID, targetID), sourceID)
		return nil
	}

	g.edges = append(g.edges[:index:index], g.edges[index+1:]...)
	forward := original.forward
	if forward == nil {
		forward = ConditionAlways().Wire()
	}
	before := len(g.edges)
	var last *Node
	if insert != nil {
		last = insert(&Branch{g: g, from: origin{source: sourceID, forward: forward}})
	}

	// Put the first new edge that leaves the source back where the original was, so the source's
	// edge order is kept.
	for i := before; i < len(g.edges); i++ {
		if g.edges[i].source != sourceID {
			continue
		}
		moved := g.edges[i]
		g.edges = append(g.edges[:i:i], g.edges[i+1:]...)
		g.edges = append(g.edges[:index:index], append([]*edge{moved}, g.edges[index:]...)...)
		if loaded, ok := g.loadedOrder[sourceID]; ok {
			for j, uid := range loaded {
				if uid == original.uid {
					loaded[j] = moved.uid
					break
				}
			}
		}
		break
	}

	if last == nil {
		g.issue(MissingInsertResult, fmt.Sprintf("InsertBetween(%q, %q) must return the last inserted node", sourceID, targetID), sourceID)
		return nil
	}
	if last.Type() == "tool" {
		last.OnSuccess(nil).GoTo(targetID, nil)
	} else {
		last.GoTo(targetID, nil)
	}
	return last
}

// Build validates the workflow and returns the generated request model that agents.Create and
// agents.Update accept. Every error is reported at once, in a *ValidationError. opts may be nil;
// opts.OnWarning is called once per warning.
func (w *Workflow) Build(opts *BuildOptions) (*types.AgentWorkflowRequestModel, error) {
	var onWarning func(Issue)
	if opts != nil {
		onWarning = opts.OnWarning
	}
	wire, err := w.toWire(onWarning)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(wire)
	if err == nil {
		var model types.AgentWorkflowRequestModel
		if err = json.Unmarshal(data, &model); err == nil {
			return &model, nil
		}
	}
	return nil, &ValidationError{Issues: []Issue{{
		Code:     InvalidModel,
		Severity: SeverityError,
		Message:  fmt.Sprintf("The workflow doesn't match the agent workflow request model: %v. A loaded field likely has the wrong shape; fix it in the source JSON", err),
	}}}
}

// ToJSON validates like Build and returns the workflow in API wire format: snake_case keys, as
// the REST API and the CLI's agent_configs/*.json use them. Numbers are json.Number values
// written as JavaScript writes them.
func (w *Workflow) ToJSON() (map[string]any, error) {
	wire, err := w.toWire(nil)
	if err != nil {
		return nil, err
	}
	canonical, err := wirejson.Canonical(wire)
	if err != nil {
		return nil, err
	}
	return canonical.(map[string]any), nil
}

// ToJSONString validates like Build and returns the canonical file format: ToJSON's output with
// keys sorted at every level, 2-space indentation and a trailing newline. It is byte-identical to
// every other ElevenLabs SDK's output for the same program, and FromJSON(s).ToJSONString() == s.
func (w *Workflow) ToJSONString() (string, error) {
	wire, err := w.toWire(nil)
	if err != nil {
		return "", err
	}
	return wirejson.String(wire)
}

func (w *Workflow) toWire(onWarning func(Issue)) (map[string]any, error) {
	g := w.g
	issues := validate(g)
	var errs []Issue
	for _, issue := range issues {
		if issue.Severity == SeverityError {
			errs = append(errs, issue)
		}
	}
	if len(errs) > 0 {
		return nil, &ValidationError{Issues: errs}
	}
	if onWarning != nil {
		for _, issue := range issues {
			if issue.Severity == SeverityWarning {
				onWarning(issue)
			}
		}
	}

	merges := reverseEdgeMerges(g)
	var kept []*edge
	for i, e := range g.edges {
		if _, merged := merges[i]; !merged {
			kept = append(kept, e)
		}
	}
	idOf := assignEdgeIDs(kept)
	positions := layout(g)

	outgoing := map[string][]int{}
	edgesWire := map[string]any{}
	for i, e := range g.edges {
		if into, merged := merges[i]; merged {
			// Sent as the backward condition of the earlier edge the other way.
			reverse := g.edges[into]
			outgoing[e.source] = append(outgoing[e.source], reverse.uid)
			if obj, ok := edgesWire[idOf[reverse.uid]].(map[string]any); ok && e.forward != nil {
				obj["backward_condition"] = wirejson.Clone(e.forward)
			}
			continue
		}
		// An edge is in a node's edge_order only for a transition that leaves the node: the source
		// has one when there is a forward condition, the target when there is a backward one.
		if e.forward != nil {
			outgoing[e.source] = append(outgoing[e.source], e.uid)
		}
		if e.backward != nil {
			outgoing[e.target] = append(outgoing[e.target], e.uid)
		}
		obj := cloneMap(e.extra)
		if obj == nil {
			obj = map[string]any{}
		}
		obj["source"] = e.source
		obj["target"] = e.target
		if e.forward != nil {
			obj["forward_condition"] = wirejson.Clone(e.forward)
		}
		if e.backward != nil {
			obj["backward_condition"] = wirejson.Clone(e.backward)
		}
		edgesWire[idOf[e.uid]] = obj
	}

	nodesWire := map[string]any{}
	for _, id := range g.nodeOrder {
		obj := cloneMap(g.nodes[id])
		obj["position"] = wirejson.Clone(positions[id])
		uids := edgeOrder(g.loadedOrder[id], outgoing[id])
		order := make([]any, len(uids))
		for i, uid := range uids {
			order[i] = idOf[uid]
		}
		obj["edge_order"] = order
		nodesWire[id] = obj
	}

	workflow := map[string]any{"nodes": nodesWire, "edges": edgesWire}
	if g.preventLoop != nil {
		workflow["prevent_subagent_loops"] = *g.preventLoop
	}
	return workflow, nil
}

// edgeOrder is a loaded node's order for the edges it still has, then any others in edge-list
// order.
func edgeOrder(loaded, current []int) []int {
	inCurrent := map[int]bool{}
	for _, uid := range current {
		inCurrent[uid] = true
	}
	var out []int
	inKept := map[int]bool{}
	for _, uid := range loaded {
		if inCurrent[uid] {
			out = append(out, uid)
			inKept[uid] = true
		}
	}
	for _, uid := range current {
		if !inKept[uid] {
			out = append(out, uid)
		}
	}
	return out
}

// assignEdgeIDs keeps loaded edges' ids. A new edge gets "{source}__{target}", or "__2", "__3"
// and so on after it when that id is taken, in edge-list order.
func assignEdgeIDs(edges []*edge) map[int]string {
	taken := map[string]bool{}
	for _, e := range edges {
		if e.hasID {
			taken[e.id] = true
		}
	}
	ids := map[int]string{}
	for _, e := range edges {
		if e.hasID {
			ids[e.uid] = e.id
			continue
		}
		base := e.source + "__" + e.target
		id := base
		for n := 2; taken[id]; n++ {
			id = fmt.Sprintf("%s__%d", base, n)
		}
		taken[id] = true
		ids[e.uid] = id
	}
	return ids
}
