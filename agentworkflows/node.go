package agentworkflows

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// Chain is implemented by *Workflow, *Node and *Branch: the verbs that add a node, connect it to
// the current point and return its handle, plus GoTo.
type Chain interface {
	// Converse adds a subagent (override_agent) node: the agent holds a conversation with prompt
	// added to its instructions until an outgoing condition holds. Its label defaults to id.
	Converse(id, prompt string, opts *Options) *Node
	// Tool adds a tool node that runs toolIDs in parallel. Route its outcome with OnSuccess,
	// OnFailure or OnResult.
	Tool(id string, toolIDs []string, opts *Options) *Node
	// Transfer adds a transfer node: phone_number for phone and SIP destinations, standalone_agent
	// for agent destinations.
	Transfer(id string, destination Destination, opts *Options) *Node
	// End adds an end node. opts.ID defaults to "end"; an id that already belongs to an end node
	// reuses that node.
	End(opts *Options) *Node
	// GoTo adds no node. It moves the conversation to target (a node id, which may be created
	// later) with a new edge, or, when target already has an edge to this point, as that edge's
	// backward condition.
	GoTo(target string, opts *GoToOptions) *Node
}

var (
	_ Chain = (*Workflow)(nil)
	_ Chain = (*Node)(nil)
	_ Chain = (*Branch)(nil)
)

// origin is where the next edge starts and which conditions it carries.
type origin struct {
	source   string
	forward  any
	backward any
}

func (g *graph) connect(from origin, target string, back *Condition) *Node {
	backward := wirejson.Clone(from.backward)
	if back != nil {
		backward = back.Wire()
	}
	g.edges = append(g.edges, &edge{
		uid:      g.uid(),
		source:   from.source,
		target:   target,
		forward:  wirejson.Clone(from.forward),
		backward: backward,
	})
	return &Node{g: g, id: target, parent: from.source, hasParent: true}
}

func (g *graph) converse(from origin, id, prompt string, opts *Options) *Node {
	base := map[string]any{"type": "override_agent", "label": id, "additional_prompt": prompt}
	g.addNode(id, g.applyOptions(id, base, opts.resolve(g, id)))
	return g.connect(from, id, nil)
}

func (g *graph) tool(from origin, id string, toolIDs []string, opts *Options) *Node {
	resolved := opts.resolve(g, id)
	resolved.toolIDs = append([]string{}, toolIDs...)
	resolved.hasToolIDs = true
	g.addNode(id, g.applyOptions(id, map[string]any{"type": "tool"}, resolved))
	return g.connect(from, id, nil)
}

func (g *graph) transfer(from origin, id string, destination Destination, opts *Options) *Node {
	g.addNode(id, g.applyOptions(id, destination.Wire(), opts.resolve(g, id)))
	return g.connect(from, id, nil)
}

func (g *graph) end(from origin, opts *Options) *Node {
	id := "end"
	var rest *Options
	if opts != nil {
		if opts.ID != "" {
			id = opts.ID
		}
		copied := *opts
		copied.ID = ""
		rest = &copied
	}
	g.addNode(id, g.applyOptions(id, map[string]any{"type": "end"}, rest.resolve(g, id)))
	return g.connect(from, id, nil)
}

func (g *graph) goTo(from origin, target string, opts *GoToOptions) *Node {
	var back *Condition
	if opts != nil {
		back = opts.Back
	}
	return g.connect(from, target, back)
}

func (g *graph) branch(source string, condition Condition, back *Condition) *Branch {
	from := origin{source: source, forward: condition.Wire()}
	if back != nil {
		from.backward = back.Wire()
	}
	return &Branch{g: g, from: from}
}

func unconditional(source string) origin {
	return origin{source: source, forward: ConditionAlways().Wire()}
}

// Branch is a pending edge out of a node, opened by When, OnSuccess or OnFailure. It only has the
// node verbs and GoTo.
type Branch struct {
	g    *graph
	from origin
}

// Converse adds a subagent node at the end of this branch. See Chain.
func (b *Branch) Converse(id, prompt string, opts *Options) *Node {
	return b.g.converse(b.from, id, prompt, opts)
}

// Tool adds a tool node at the end of this branch. See Chain.
func (b *Branch) Tool(id string, toolIDs []string, opts *Options) *Node {
	return b.g.tool(b.from, id, toolIDs, opts)
}

// Transfer adds a transfer node at the end of this branch. See Chain.
func (b *Branch) Transfer(id string, destination Destination, opts *Options) *Node {
	return b.g.transfer(b.from, id, destination, opts)
}

// End adds (or reuses) an end node at the end of this branch. See Chain.
func (b *Branch) End(opts *Options) *Node {
	return b.g.end(b.from, opts)
}

// GoTo connects this branch to target. See Chain.
func (b *Branch) GoTo(target string, opts *GoToOptions) *Node {
	return b.g.goTo(b.from, target, opts)
}

// Node is a handle to one node in a workflow. Handles are cheap references: all state lives in
// the Workflow.
type Node struct {
	g         *graph
	id        string
	parent    string
	hasParent bool
}

// ID is the node's id.
func (n *Node) ID() string { return n.id }

// Type is the node's API type, for example "override_agent" or "tool", or "" if the node doesn't
// exist.
func (n *Node) Type() string { return n.g.nodeType(n.id) }

// Converse adds a subagent node reached unconditionally from this node. See Chain.
func (n *Node) Converse(id, prompt string, opts *Options) *Node {
	return n.g.converse(unconditional(n.id), id, prompt, opts)
}

// Tool adds a tool node reached unconditionally from this node. See Chain.
func (n *Node) Tool(id string, toolIDs []string, opts *Options) *Node {
	return n.g.tool(unconditional(n.id), id, toolIDs, opts)
}

// Transfer adds a transfer node reached unconditionally from this node. See Chain.
func (n *Node) Transfer(id string, destination Destination, opts *Options) *Node {
	return n.g.transfer(unconditional(n.id), id, destination, opts)
}

// End adds (or reuses) an end node reached unconditionally from this node. See Chain.
func (n *Node) End(opts *Options) *Node {
	return n.g.end(unconditional(n.id), opts)
}

// GoTo connects this node unconditionally to target. See Chain.
func (n *Node) GoTo(target string, opts *GoToOptions) *Node {
	return n.g.goTo(unconditional(n.id), target, opts)
}

// When opens a branch followed when condition holds.
func (n *Node) When(condition Condition, opts *WhenOptions) *Branch {
	return n.g.branch(n.id, condition, whenBack(opts))
}

// OnSuccess opens the branch followed when every tool in this tool node succeeded.
func (n *Node) OnSuccess(opts *ResultOptions) *Branch {
	return n.g.branch(n.id, resultCondition(true, opts), nil)
}

// OnFailure opens the branch followed when any tool in this tool node failed.
func (n *Node) OnFailure(opts *ResultOptions) *Branch {
	return n.g.branch(n.id, resultCondition(false, opts), nil)
}

// OnResult routes both outcomes of this tool node: it calls branches.Success with OnSuccess's
// branch, then branches.Failure with OnFailure's, always in that order. Either may be nil. It
// returns this node.
func (n *Node) OnResult(branches ResultBranches) *Node {
	if branches.Success != nil {
		branches.Success(n.OnSuccess(nil))
	}
	if branches.Failure != nil {
		branches.Failure(n.OnFailure(nil))
	}
	return n
}

// Up returns a handle to the node this handle was chained from: the source of the edge that the
// verb or GoTo that returned it created. For a handle from Workflow.Node, it is the source of the
// first edge into the node other than a self-loop. For the start node, or a node nothing leads
// to, it is the start node. Up adds nothing to the graph; it lets a whole workflow be written as
// one chain.
func (n *Node) Up() *Node {
	parent := n.g.startID
	if n.hasParent {
		parent = n.parent
	} else {
		for _, e := range n.g.edges {
			if e.target == n.id && e.source != n.id {
				parent = e.source
				break
			}
		}
	}
	return &Node{g: n.g, id: parent}
}

// Update changes this node's options. Only options that apply to the node's kind may be set
// (plus Prompt for a subagent, ToolIDs for a subagent or tool node, and Destination for a
// transfer); Build reports the rest. It returns this node.
func (n *Node) Update(opts *Options) *Node {
	model, ok := n.g.nodes[n.id]
	if !ok {
		n.g.issue(UnknownNode, fmt.Sprintf("Can't update unknown node %q", n.id), n.id)
		return n
	}
	n.g.nodes[n.id] = n.g.applyOptions(n.id, model, opts.resolve(n.g, n.id))
	return n
}

func whenBack(opts *WhenOptions) *Condition {
	if opts == nil {
		return nil
	}
	return opts.Back
}

func resultCondition(successful bool, opts *ResultOptions) Condition {
	label := ""
	if opts != nil {
		label = opts.Label
	}
	return newCondition(map[string]any{"type": "result", "successful": successful}, label)
}
