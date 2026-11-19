package agentworkflows

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/elevenlabs/elevenlabs-go/types"
)

// validate runs the checks that need the finished graph. Problems found while chaining are
// already in g.issues.
func validate(g *graph) []Issue {
	issues := append([]Issue{}, g.issues...)
	errorf := func(code IssueCode, nodeID, format string, args ...any) {
		issues = append(issues, Issue{Code: code, Severity: SeverityError, Message: fmt.Sprintf(format, args...), NodeID: nodeID})
	}
	warnf := func(code IssueCode, nodeID, format string, args ...any) {
		issues = append(issues, Issue{Code: code, Severity: SeverityWarning, Message: fmt.Sprintf(format, args...), NodeID: nodeID})
	}

	merged := reverseEdgeMerges(g)
	type pair struct{ a, b string }
	connected := map[pair]bool{}
	for i, e := range g.edges {
		if _, ok := merged[i]; ok {
			continue
		}
		p := pair{e.source, e.target}
		if p.b < p.a {
			p = pair{e.target, e.source}
		}
		if connected[p] {
			errorf(DuplicateEdge, e.source, "Edge %q → %q connects two nodes that already have an edge between them, and the API allows only one. Give the existing edge a different condition, or remove one", e.source, e.target)
		}
		connected[p] = true
	}

	endsWithEdges := map[string]bool{}
	// checkTransition applies the routing rules to one transition along an edge: the forward one
	// (source → target) or the reverse one a backward condition stores (target → source).
	checkTransition := func(from, to string, condition any) {
		name := fmt.Sprintf("%q → %q", from, to)
		fromType := g.nodeType(from)
		conditionType := ""
		if m, ok := condition.(map[string]any); ok {
			conditionType, _ = m["type"].(string)
		}
		if to == g.startID {
			errorf(EdgeIntoStart, from, "Edge %s points into the start node; go to the first node after it instead", name)
		}
		if fromType == "end" && !endsWithEdges[from] {
			endsWithEdges[from] = true
			errorf(EndHasOutgoingEdges, from, "End node %q can't have outgoing edges", from)
		}
		if fromType == "tool" && conditionType != "result" {
			errorf(ToolEdgeNeedsResult, from, "Edge %s leaves a tool node, so it must use OnSuccess() or OnFailure()", name)
		}
		if g.hasNode(from) && fromType != "" && fromType != "tool" && conditionType == "result" {
			errorf(ResultEdgeNeedsTool, from, "Edge %s uses OnSuccess()/OnFailure(), which only applies to tool nodes", name)
		}
	}

	for _, e := range g.edges {
		name := fmt.Sprintf("%q → %q", e.source, e.target)
		if e.source == e.target {
			errorf(SelfLoop, e.source, "Edge %s connects a node to itself, which the API doesn't allow", name)
		}
		for _, id := range []string{e.source, e.target} {
			if !g.hasNode(id) {
				errorf(UnknownNode, id, "Edge %s refers to unknown node %q", name, id)
			}
		}
		// The forward rules apply only when there is a forward transition: a loaded edge can be
		// backward-only, and then source → target doesn't exist.
		if e.forward != nil {
			checkTransition(e.source, e.target, e.forward)
		}
		// A backward condition is a second transition, target → source, stored on the same edge.
		// It obeys the same rules as any edge leaving the target.
		if e.backward != nil {
			checkTransition(e.target, e.source, e.backward)
		}
	}

	for _, id := range g.nodeOrder {
		nodeType := g.nodeType(id)
		if !modelKnowsNodeType(nodeType) {
			shown := nodeType
			if shown == "" {
				shown = "None"
			}
			warnf(UnknownNodeType, id, "Node %q has type %q, which this SDK version doesn't model. Build keeps it as given but can't type it, and agents.Create/agents.Update would reject it. Upgrade the SDK, or send ToJSON's output directly", id, shown)
		}
	}

	depths, _ := depthsFrom(g)
	for _, id := range g.nodeOrder {
		if _, ok := depths[id]; !ok {
			warnf(UnreachableNode, id, "Node %q can't be reached from the start node", id)
		}
	}
	return issues
}

// reverseEdgeMerges finds the edges the API stores as the backward condition of an earlier edge
// between the same two nodes: index of the later edge → index of the earlier one.
func reverseEdgeMerges(g *graph) map[int]int {
	merges := map[int]int{}
	taken := map[int]bool{}
	for j, later := range g.edges {
		if later.source == later.target || later.backward != nil || later.forward == nil {
			continue
		}
		for i := 0; i < j; i++ {
			earlier := g.edges[i]
			_, mergedAway := merges[i]
			if !taken[i] && !mergedAway &&
				earlier.source == later.target && earlier.target == later.source && earlier.backward == nil {
				merges[j] = i
				taken[i] = true
				break
			}
		}
	}
	return merges
}

// depthsFrom is the breadth-first depth of every node reachable from the start, with the order
// the search visited them in. Edges are followed forward in edge-list order; when that search
// ends, the first edge with a backward condition from a visited target to an unvisited source
// continues it from that source, one column further.
func depthsFrom(g *graph) (map[string]int, []string) {
	depths := map[string]int{}
	var order []string
	visit := func(start string, depth int) {
		depths[start] = depth
		order = append(order, start)
		queue := []string{start}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, e := range g.edges {
				if e.source != current || e.forward == nil || !g.hasNode(e.target) {
					continue
				}
				if _, seen := depths[e.target]; seen {
					continue
				}
				depths[e.target] = depths[current] + 1
				order = append(order, e.target)
				queue = append(queue, e.target)
			}
		}
	}
	visit(g.startID, 0)
	for {
		var next *edge
		for _, e := range g.edges {
			if e.backward == nil || !g.hasNode(e.source) {
				continue
			}
			_, targetSeen := depths[e.target]
			_, sourceSeen := depths[e.source]
			if targetSeen && !sourceSeen {
				next = e
				break
			}
		}
		if next == nil {
			break
		}
		visit(next.source, depths[next.target]+1)
	}
	return depths, order
}

var (
	knownNodeTypesOnce sync.Once
	knownNodeTypes     map[string]bool
)

// modelKnowsNodeType reports whether this SDK version's generated request type has a member for
// nodeType. The set is read from the generated union itself, so a regeneration that adds a node
// type is picked up without changing the builder.
func modelKnowsNodeType(nodeType string) bool {
	knownNodeTypesOnce.Do(func() { knownNodeTypes = generatedNodeTypes() })
	return knownNodeTypes[nodeType]
}

// generatedNodeTypes lists the discriminants of types.AgentWorkflowRequestModelNodesValue: each
// member is a pointer-to-struct field, and marshaling the union with only that member set writes
// the member's "type".
func generatedNodeTypes() map[string]bool {
	known := map[string]bool{}
	unionType := reflect.TypeOf(types.AgentWorkflowRequestModelNodesValue{})
	for i := 0; i < unionType.NumField(); i++ {
		field := unionType.Field(i)
		if !field.IsExported() || field.Type.Kind() != reflect.Ptr || field.Type.Elem().Kind() != reflect.Struct {
			continue
		}
		union := reflect.New(unionType).Elem()
		union.Field(i).Set(reflect.New(field.Type.Elem()))
		data, err := json.Marshal(union.Interface())
		if err != nil {
			continue
		}
		var probe struct {
			Type string `json:"type"`
		}
		if json.Unmarshal(data, &probe) == nil && probe.Type != "" {
			known[probe.Type] = true
		}
	}
	return known
}
