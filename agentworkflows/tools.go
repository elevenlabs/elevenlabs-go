package agentworkflows

import (
	"fmt"
	"reflect"
	"strings"
)

// Tool is a tool written as code, such as an agenttools.Tool, that a tool node or a subagent uses
// before it exists on the server. Pass it in Options.Tools.
//
// While the workflow is being built, the tool stands under the placeholder "<tool:NAME>". Build,
// ToJSON and ToJSONString swap in its id from BuildOptions.ToolIDs (by tool name), as
// agentbuilder's push does after creating the tools; a tool without one is reported as
// unresolved_tool. Tools are told apart by identity: two different values with the same name are
// reported as duplicate_tool_name.
type Tool interface {
	// Name is the tool's name, as the LLM sees it.
	Name() string
}

// toolPlaceholder is the placeholder a Tool stands under in the graph until Build swaps in its id.
func toolPlaceholder(name string) string { return "<tool:" + name + ">" }

// displayTool is a tool id as shown in messages: the name, for a Tool's placeholder.
func displayTool(id string) string {
	if strings.HasPrefix(id, "<tool:") && strings.HasSuffix(id, ">") {
		return id[len("<tool:") : len(id)-1]
	}
	return id
}

// sameTool reports whether a and b are the same tool value: the same pointer, for the tools the
// agenttools package returns. Values of types that can't be compared are never the same.
func sameTool(a, b Tool) (same bool) {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if reflect.TypeOf(a) != reflect.TypeOf(b) || !reflect.TypeOf(a).Comparable() {
		return false
	}
	defer func() {
		if recover() != nil {
			same = false
		}
	}()
	return a == b
}

// refTools returns ids followed by the placeholders of tools, remembering each Tool. Two different
// tools with one name conflict if this node uses both, or if another node still uses the other; a
// tool left behind by a removed or edited node doesn't count.
func (g *graph) refTools(nodeID string, ids []string, tools []Tool) []string {
	out := append([]string{}, ids...)
	inThisNode := map[string]Tool{}
	for _, tool := range tools {
		if tool == nil {
			g.issue(InvalidOption, fmt.Sprintf("Node %q has a nil tool", nodeID), nodeID)
			continue
		}
		key := toolPlaceholder(tool.Name())
		clash, ok := inThisNode[key]
		if !ok {
			if known, isKnown := g.toolRefs[key]; isKnown && g.usesPlaceholder(key, nodeID) {
				clash, ok = known, true
			}
		}
		inThisNode[key] = tool
		if ok && !sameTool(clash, tool) {
			g.issues = append(g.issues, Issue{
				Code:     DuplicateToolName,
				Severity: SeverityError,
				Message:  fmt.Sprintf("Node %q uses a different tool also named %q", nodeID, tool.Name()),
				NodeID:   nodeID,
			})
		}
		if _, isKnown := g.toolRefs[key]; !isKnown {
			g.toolRefOrder = append(g.toolRefOrder, key)
		}
		g.toolRefs[key] = tool
		out = append(out, key)
	}
	return out
}

// usesPlaceholder reports whether a node other than exceptNodeID uses key.
func (g *graph) usesPlaceholder(key, exceptNodeID string) bool {
	for _, id := range g.nodeOrder {
		if id == exceptNodeID {
			continue
		}
		for _, toolID := range nodeToolIDs(g.nodes[id]) {
			if toolID == key {
				return true
			}
		}
	}
	return false
}

// nodeToolIDs are the tool ids (or placeholders) a node uses: a tool node's tools, or a
// subagent's additional tools.
func nodeToolIDs(model map[string]any) []string {
	var ids []string
	tools, _ := model["tools"].([]any)
	for _, tool := range tools {
		if locator, ok := tool.(map[string]any); ok {
			if id, ok := locator["tool_id"].(string); ok {
				ids = append(ids, id)
			}
		}
	}
	extra, _ := model["additional_tool_ids"].([]any)
	for _, id := range extra {
		if s, ok := id.(string); ok {
			ids = append(ids, s)
		}
	}
	return ids
}

// mapNodeToolIDs returns a copy of a tool node or subagent with each tool id passed through fn.
// Other nodes are returned as they are.
func mapNodeToolIDs(model map[string]any, fn func(string) string) map[string]any {
	switch model["type"] {
	case "tool":
		tools, ok := model["tools"].([]any)
		if !ok {
			return model
		}
		out := shallowCopy(model)
		mapped := make([]any, len(tools))
		for i, tool := range tools {
			locator, isMap := tool.(map[string]any)
			id, isString := locator["tool_id"].(string)
			if !isMap || !isString {
				mapped[i] = tool
				continue
			}
			copied := shallowCopy(locator)
			copied["tool_id"] = fn(id)
			mapped[i] = copied
		}
		out["tools"] = mapped
		return out
	case "override_agent":
		extra, ok := model["additional_tool_ids"].([]any)
		if !ok {
			return model
		}
		out := shallowCopy(model)
		mapped := make([]any, len(extra))
		for i, id := range extra {
			if s, isString := id.(string); isString {
				mapped[i] = fn(s)
			} else {
				mapped[i] = id
			}
		}
		out["additional_tool_ids"] = mapped
		return out
	}
	return model
}

func shallowCopy(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for key, value := range m {
		out[key] = value
	}
	return out
}

// Tools lists the Tool values the workflow's nodes use, each once, in the order they were first
// used. A tool only a removed or edited node used isn't listed.
func (w *Workflow) Tools() []Tool {
	used := map[string]bool{}
	for _, id := range w.g.nodeOrder {
		for _, toolID := range nodeToolIDs(w.g.nodes[id]) {
			used[toolID] = true
		}
	}
	var tools []Tool
	for _, key := range w.g.toolRefOrder {
		if used[key] {
			tools = append(tools, w.g.toolRefs[key])
		}
	}
	return tools
}

// FromJSONWithTools loads a workflow like FromJSON, then swaps each tool id found in toolsByID for
// that Tool, so the loaded workflow refers to it as a built one would: Tools lists it, and Build
// takes its id from BuildOptions.ToolIDs. Ids missing from toolsByID stay as they are.
func FromJSONWithTools(data any, toolsByID map[string]Tool) (*Workflow, error) {
	wf, err := FromJSON(data)
	if err != nil {
		return nil, err
	}
	g := wf.g
	swap := func(id string) string {
		tool, ok := toolsByID[id]
		if !ok || tool == nil {
			return id
		}
		key := toolPlaceholder(tool.Name())
		if _, known := g.toolRefs[key]; !known {
			g.toolRefOrder = append(g.toolRefOrder, key)
		}
		g.toolRefs[key] = tool
		return key
	}
	for _, id := range g.nodeOrder {
		g.nodes[id] = mapNodeToolIDs(g.nodes[id], swap)
	}
	return wf, nil
}

// unresolvedTools reports each node that uses a Tool with no id in toolIDs.
func unresolvedTools(g *graph, toolIDs map[string]string) []Issue {
	var issues []Issue
	for _, nodeID := range g.nodeOrder {
		for _, id := range nodeToolIDs(g.nodes[nodeID]) {
			tool, ok := g.toolRefs[id]
			if !ok {
				continue
			}
			if _, has := toolIDs[tool.Name()]; !has {
				issues = append(issues, Issue{
					Code:     UnresolvedTool,
					Severity: SeverityError,
					Message:  fmt.Sprintf("Node %q uses tool %q, which has no id yet. Push the agent with agentbuilder's Push, or pass BuildOptions.ToolIDs", nodeID, tool.Name()),
					NodeID:   nodeID,
				})
			}
		}
	}
	return issues
}

// resolveTools returns a copy of a node model with each Tool's placeholder replaced by its id.
func resolveTools(g *graph, model map[string]any, toolIDs map[string]string) map[string]any {
	if len(g.toolRefs) == 0 {
		return model
	}
	return mapNodeToolIDs(model, func(id string) string {
		if tool, ok := g.toolRefs[id]; ok {
			return toolIDs[tool.Name()]
		}
		return id
	})
}
