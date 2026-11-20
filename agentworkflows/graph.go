package agentworkflows

import (
	"encoding/json"
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// defaultStartID is the start node id the API requires.
const defaultStartID = "start_node"

// edge is an edge in the graph. Edges are compared by identity (uid): two edges with the same
// fields are still two edges.
type edge struct {
	uid    int
	source string
	target string
	// forward and backward are wire-format conditions, or nil when the edge has none. A loaded
	// JSON null means "no condition", the same as an absent key.
	forward  any
	backward any
	// id is set for edges loaded from an existing workflow; others get one at build time.
	id    string
	hasID bool
	// extra holds a loaded edge's other fields, such as reroutes, written back as they were.
	extra map[string]any
}

// graph is the mutable state shared by a Workflow and every handle created from it. Nodes and
// edges are kept in API wire format.
type graph struct {
	nodes       map[string]map[string]any
	nodeOrder   []string
	edges       []*edge
	issues      []Issue
	startID     string
	loadedOrder map[string][]int // each loaded node's edge_order, as edge uids
	preventLoop *bool
	nextUID     int
	// toolRefs are the Tool values nodes reference, by placeholder, and toolRefOrder their
	// placeholders in the order they were first used.
	toolRefs     map[string]Tool
	toolRefOrder []string
}

func newGraph() *graph {
	return &graph{
		nodes:       map[string]map[string]any{},
		startID:     defaultStartID,
		loadedOrder: map[string][]int{},
		toolRefs:    map[string]Tool{},
	}
}

func (g *graph) uid() int {
	g.nextUID++
	return g.nextUID
}

func (g *graph) issue(code IssueCode, message, nodeID string) {
	g.issues = append(g.issues, Issue{Code: code, Severity: SeverityError, Message: message, NodeID: nodeID})
}

// nodeType is a node's "type", or "" when the node doesn't exist or has none.
func (g *graph) nodeType(id string) string {
	model, ok := g.nodes[id]
	if !ok {
		return ""
	}
	t, _ := model["type"].(string)
	return t
}

func (g *graph) hasNode(id string) bool {
	_, ok := g.nodes[id]
	return ok
}

func (g *graph) addNode(id string, model map[string]any) {
	existing, ok := g.nodes[id]
	if !ok {
		g.nodes[id] = model
		g.nodeOrder = append(g.nodeOrder, id)
		return
	}
	if existing["type"] == "end" && model["type"] == "end" {
		return
	}
	g.issue(DuplicateNodeID, fmt.Sprintf("Node id %q is used by more than one node", id), id)
}

func nodeKind(nodeType string) string {
	switch nodeType {
	case "start", "end", "tool":
		return nodeType
	case "override_agent":
		return "subagent"
	case "phone_number":
		return "phone/SIP transfer"
	case "standalone_agent":
		return "agent transfer"
	case "":
		return `"None"`
	}
	return fmt.Sprintf("%q", nodeType)
}

// apiField is the API field a canonical option maps to for a node type, or "" if it doesn't apply.
func apiField(nodeType, key string) string {
	if key == "position" {
		// Common to every node kind, including those with no field map of their own (such as the
		// generated say and update_state types).
		return "position"
	}
	switch nodeType {
	case "override_agent":
		switch key {
		case "prompt":
			return "additional_prompt"
		case "label", "conversation_config", "entry_behavior":
			return key
		case "tool_ids":
			return "additional_tool_ids"
		case "knowledge_base":
			return "additional_knowledge_base"
		}
	case "phone_number":
		switch key {
		case "transfer_type", "custom_sip_headers", "sip_refer_play_dialtone", "uui", "post_dial_digits":
			return key
		}
	case "standalone_agent":
		switch key {
		case "delay_ms", "transfer_message", "enable_transferred_agent_first_message",
			"preserve_client_tts_overrides", "preserve_voice_settings":
			return key
		}
	}
	return ""
}

func (g *graph) invalidOption(id, option, nodeType string) {
	g.issue(InvalidOption, fmt.Sprintf("Option %q doesn't apply to %s node %q", option, nodeKind(nodeType), id), id)
}

func isTransferType(nodeType string) bool {
	return nodeType == destinationNodeTypePhoneNumber || nodeType == destinationNodeTypeAgent
}

// applyOptions applies resolved options to a node model, recording invalid_option for options
// that don't apply. It returns the updated model.
func (g *graph) applyOptions(id string, model map[string]any, opts resolvedOptions) map[string]any {
	node := model
	rest := append([]option{}, opts.rest...)

	if opts.destination != nil {
		nodeType, _ := node["type"].(string)
		if isTransferType(nodeType) {
			dest := opts.destination.Wire()
			if dest["type"] != nodeType {
				// A destination of the other transfer kind changes the node's type; only its
				// position is kept.
				position, hadPosition := node["position"]
				node = dest
				if hadPosition {
					node["position"] = position
				}
			} else {
				if nodeType == destinationNodeTypeAgent {
					// Start at the new agent's default entry point unless the destination says
					// otherwise.
					delete(node, "agent_id")
					delete(node, "node_id")
					delete(node, "branch_id")
				}
				for key, value := range dest {
					node[key] = value
				}
			}
		} else {
			g.invalidOption(id, "destination", nodeType)
		}
	}

	// Tool values stand under their placeholders, and schema overrides for one may be keyed by its
	// name.
	var refs []string
	if opts.hasToolIDs {
		refs = g.refTools(id, opts.toolIDs, opts.tools)
	}
	if opts.hasOverrides {
		overrides := make(map[string]any, len(opts.schemaOverrides))
		for key, value := range opts.schemaOverrides {
			if _, isTool := g.toolRefs[toolPlaceholder(key)]; isTool {
				key = toolPlaceholder(key)
			}
			overrides[key] = value
		}
		opts.schemaOverrides = overrides
	}

	nodeType, _ := node["type"].(string)
	if nodeType == "tool" {
		existing, _ := node["tools"].([]any)
		var ids []string
		if opts.hasToolIDs {
			ids = refs
		} else {
			for _, tool := range existing {
				if m, ok := tool.(map[string]any); ok {
					if toolID, ok := m["tool_id"].(string); ok {
						ids = append(ids, toolID)
					}
				}
			}
		}
		if opts.hasOverrides {
			runs := map[string]bool{}
			for _, toolID := range ids {
				runs[toolID] = true
			}
			for _, toolID := range wirejson.SortKeys(keysOf(opts.schemaOverrides)) {
				if !runs[toolID] {
					g.issue(InvalidOption, fmt.Sprintf("Node %q has schema overrides for tool %q, which it doesn't run", id, displayTool(toolID)), id)
				}
			}
		}
		tools := make([]any, 0, len(ids))
		for _, toolID := range ids {
			locator := map[string]any{"tool_id": toolID}
			if override, ok := opts.schemaOverrides[toolID]; ok && opts.hasOverrides {
				locator["schema_overrides"] = override
			} else if loaded := findTool(existing, toolID); loaded != nil {
				if override, ok := loaded["schema_overrides"]; ok {
					locator["schema_overrides"] = override
				}
			}
			tools = append(tools, locator)
		}
		node["tools"] = tools
	} else {
		if opts.hasOverrides {
			g.invalidOption(id, "schema_overrides", nodeType)
		}
		if opts.hasToolIDs {
			values := make([]any, len(refs))
			for i, toolID := range refs {
				values[i] = toolID
			}
			rest = append(rest, option{key: "tool_ids", value: values})
		}
	}

	for _, o := range rest {
		field := apiField(nodeType, o.key)
		if field == "" {
			g.invalidOption(id, o.key, nodeType)
			continue
		}
		node[field] = o.value
	}
	return node
}

func findTool(tools []any, toolID string) map[string]any {
	for _, tool := range tools {
		if m, ok := tool.(map[string]any); ok && m["tool_id"] == toolID {
			return m
		}
	}
	return nil
}

func keysOf(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// load replaces the graph's contents with an existing wire-format workflow (raw JSON object
// text), keeping ids, positions and edge order. Document order (as JavaScript reads it) decides
// the node order and the order of edges no node lists.
func (g *graph) load(raw []byte) error {
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return fmt.Errorf("agentworkflows: workflow JSON must be an object: %w", err)
	}
	g.nodes = map[string]map[string]any{}
	g.nodeOrder = nil
	g.edges = nil
	g.issues = nil
	g.loadedOrder = map[string][]int{}
	g.preventLoop = nil
	g.toolRefs = map[string]Tool{}
	g.toolRefOrder = nil

	if value, ok := top["prevent_subagent_loops"]; ok {
		var b bool
		if json.Unmarshal(value, &b) == nil && string(value) != "null" {
			g.preventLoop = &b
		}
	}

	nodeIDs, nodes, err := orderedObject(top["nodes"])
	if err != nil {
		return fmt.Errorf("agentworkflows: workflow nodes: %w", err)
	}
	edgeIDs, edges, err := orderedObject(top["edges"])
	if err != nil {
		return fmt.Errorf("agentworkflows: workflow edges: %w", err)
	}

	edgeOrders := map[string][]string{}
	for _, id := range nodeIDs {
		node, _ := nodes[id].(map[string]any)
		model := map[string]any{}
		// A node type the generated types model has its null fields read as absent, as the
		// reference SDK's model parser reads them, so a response-shaped workflow (every unset field
		// null) loads like a request-shaped one. Other node types are kept exactly as loaded.
		nodeType, _ := node["type"].(string)
		known := modelKnowsNodeType(nodeType)
		for key, value := range node {
			if key == "edge_order" {
				continue
			}
			if known {
				if value == nil {
					continue
				}
				value = withoutNulls(value)
			}
			model[key] = value
		}
		if order, ok := node["edge_order"].([]any); ok {
			for _, edgeID := range order {
				if s, ok := edgeID.(string); ok {
					edgeOrders[id] = append(edgeOrders[id], s)
				}
			}
		}
		g.nodes[id] = model
		g.nodeOrder = append(g.nodeOrder, id)
	}

	// The start node is the first node, in document order, whose type is "start".
	g.startID = defaultStartID
	for _, id := range nodeIDs {
		if g.nodeType(id) == "start" {
			g.startID = id
			break
		}
	}
	if !g.hasNode(g.startID) {
		g.nodes[g.startID] = map[string]any{"type": "start"}
		g.nodeOrder = append(g.nodeOrder, g.startID)
	}

	// Rebuild the edge list in each node's edge_order (so build reproduces it), then any edges no
	// node lists, in document order.
	placed := map[string]int{}
	place := func(edgeID string) {
		if _, done := placed[edgeID]; done {
			return
		}
		value, ok := edges[edgeID]
		if !ok {
			return
		}
		fields, _ := value.(map[string]any)
		e := &edge{uid: g.uid(), id: edgeID, hasID: true, extra: map[string]any{}}
		for key, v := range fields {
			switch key {
			case "source":
				e.source, _ = v.(string)
			case "target":
				e.target, _ = v.(string)
			case "forward_condition":
				e.forward = withoutNulls(v)
			case "backward_condition":
				e.backward = withoutNulls(v)
			default:
				if v != nil {
					e.extra[key] = withoutNulls(v)
				}
			}
		}
		placed[edgeID] = e.uid
		g.edges = append(g.edges, e)
	}
	for _, id := range nodeIDs {
		for _, edgeID := range edgeOrders[id] {
			place(edgeID)
		}
		var uids []int
		for _, edgeID := range edgeOrders[id] {
			if uid, ok := placed[edgeID]; ok {
				uids = append(uids, uid)
			}
		}
		g.loadedOrder[id] = uids
	}
	for _, edgeID := range edgeIDs {
		place(edgeID)
	}
	return nil
}

// withoutNulls drops null-valued object fields at every level, except a constant override's
// constant_value, where null is data.
func withoutNulls(value any) any {
	switch v := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(v))
		for key, item := range v {
			if item == nil {
				if key == "constant_value" {
					out[key] = nil
				}
				continue
			}
			out[key] = withoutNulls(item)
		}
		return out
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = withoutNulls(item)
		}
		return out
	}
	return value
}

// orderedObject decodes a JSON object into its keys, in JavaScript's order, and values. Absent or
// null input is an empty object.
func orderedObject(raw json.RawMessage) ([]string, map[string]any, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, map[string]any{}, nil
	}
	value, err := wirejson.Decode(raw)
	if err != nil {
		return nil, nil, err
	}
	obj, ok := value.(map[string]any)
	if !ok {
		return nil, nil, fmt.Errorf("expected an object")
	}
	keys, err := wirejson.ObjectKeys(raw)
	if err != nil {
		return nil, nil, err
	}
	return keys, obj, nil
}
