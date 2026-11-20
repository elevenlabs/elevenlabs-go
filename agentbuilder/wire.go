package agentbuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// Helpers for agent configs in API wire format: plain JSON trees (map[string]any, []any,
// json.Number...) as the workflow and tool builders keep them.

var (
	toolIDsPath      = []string{"conversation_config", "agent", "prompt", "tool_ids"}
	builtInToolsPath = []string{"conversation_config", "agent", "prompt", "built_in_tools"}
)

// rawJSON is data as JSON text: text as it is, anything else marshaled.
func rawJSON(data any) ([]byte, error) {
	switch v := data.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	case json.RawMessage:
		return v, nil
	}
	encoded, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("agentbuilder: can't convert %T to JSON: %w", data, err)
	}
	return encoded, nil
}

// decodeObject parses JSON text that must be an object.
func decodeObject(data []byte, what string) (map[string]any, error) {
	tree, err := wirejson.Decode(bytes.TrimSpace(data))
	if err != nil {
		return nil, fmt.Errorf("agentbuilder: %s isn't valid JSON: %w", what, err)
	}
	obj, ok := tree.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("agentbuilder: %s must be a JSON object", what)
	}
	return obj, nil
}

// field returns the raw JSON text of one field of a JSON object, or nil.
func field(data []byte, key string) json.RawMessage {
	var fields map[string]json.RawMessage
	if json.Unmarshal(data, &fields) != nil {
		return nil
	}
	return fields[key]
}

func isObject(value any) bool {
	_, ok := value.(map[string]any)
	return ok
}

// getPath returns the value at path, and whether every step is there (a null counts as there).
func getPath(target any, path []string) (any, bool) {
	current := target
	for _, key := range path {
		obj, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		if current, ok = obj[key]; !ok {
			return nil, false
		}
	}
	return current, true
}

// isSet reports whether path holds a value other than null.
func isSet(target any, path []string) bool {
	value, ok := getPath(target, path)
	return ok && value != nil
}

// setPath sets the value at path, creating objects along the way.
func setPath(target map[string]any, path []string, value any) {
	current := target
	for _, key := range path[:len(path)-1] {
		next, ok := current[key].(map[string]any)
		if !ok {
			next = map[string]any{}
			current[key] = next
		}
		current = next
	}
	current[path[len(path)-1]] = value
}

// deletePath removes the key at path, if it's there.
func deletePath(target map[string]any, path []string) {
	current := target
	for _, key := range path[:len(path)-1] {
		next, ok := current[key].(map[string]any)
		if !ok {
			return
		}
		current = next
	}
	delete(current, path[len(path)-1])
}

// deepMerge merges source into target: objects merge key by key, anything else (arrays included)
// replaces. Values are copied, so later edits to source don't reach target.
func deepMerge(target, source map[string]any) {
	for key, value := range source {
		existing, isObj := target[key].(map[string]any)
		incoming, incomingObj := value.(map[string]any)
		if isObj && incomingObj {
			deepMerge(existing, incoming)
			continue
		}
		target[key] = wirejson.Clone(value)
	}
}

// isSubset reports whether every value in local is in remote at the same place. Extra keys in
// remote (server defaults) are fine, and a null in local matches anything. Numbers compare by
// value, since JSON can't tell 1 from 1.0.
func isSubset(local, remote any) bool {
	switch l := local.(type) {
	case []any:
		r, ok := remote.([]any)
		if !ok || len(l) != len(r) {
			return false
		}
		for i := range l {
			if !isSubset(l[i], r[i]) {
				return false
			}
		}
		return true
	case map[string]any:
		r, ok := remote.(map[string]any)
		if !ok {
			return false
		}
		for key, value := range l {
			if value != nil && !isSubset(value, r[key]) {
				return false
			}
		}
		return true
	case json.Number:
		r, ok := remote.(json.Number)
		if !ok {
			return false
		}
		if l == r {
			return true
		}
		a, errA := strconv.ParseFloat(string(l), 64)
		b, errB := strconv.ParseFloat(string(r), 64)
		return errA == nil && errB == nil && a == b
	}
	return local == remote
}

// workflowToolIDs are the tool ids a wire-format workflow's nodes use, in node order: each tool
// node's tools, and each subagent's additional_tool_ids. order is the node ids in document
// order; nodes it doesn't list follow in sorted order.
func workflowToolIDs(workflow any, order []string) []string {
	wf, _ := workflow.(map[string]any)
	nodes, _ := wf["nodes"].(map[string]any)
	var ids []string
	for _, nodeID := range orderedKeys(nodes, order) {
		node, _ := nodes[nodeID].(map[string]any)
		tools, _ := node["tools"].([]any)
		for _, tool := range tools {
			if locator, ok := tool.(map[string]any); ok {
				if id, ok := locator["tool_id"].(string); ok {
					ids = append(ids, id)
				}
			}
		}
		extra, _ := node["additional_tool_ids"].([]any)
		for _, id := range extra {
			if s, ok := id.(string); ok {
				ids = append(ids, s)
			}
		}
	}
	return ids
}

// mapWorkflowToolIDs replaces the tool ids in a wire-format workflow's tool nodes and subagents,
// in place.
func mapWorkflowToolIDs(workflow map[string]any, fn func(string) string) {
	nodes, _ := workflow["nodes"].(map[string]any)
	for _, value := range nodes {
		node, ok := value.(map[string]any)
		if !ok {
			continue
		}
		if tools, ok := node["tools"].([]any); ok {
			for _, tool := range tools {
				if locator, ok := tool.(map[string]any); ok {
					if id, ok := locator["tool_id"].(string); ok {
						locator["tool_id"] = fn(id)
					}
				}
			}
		}
		if extra, ok := node["additional_tool_ids"].([]any); ok {
			for i, id := range extra {
				if s, ok := id.(string); ok {
					extra[i] = fn(s)
				}
			}
		}
	}
}

// nodeOrder is the document order of a JSON workflow's node ids, or nil.
func nodeOrder(workflowJSON []byte) []string {
	nodes := field(workflowJSON, "nodes")
	if len(nodes) == 0 {
		return nil
	}
	keys, err := wirejson.ObjectKeys(nodes)
	if err != nil {
		return nil
	}
	return keys
}

// orderedKeys lists m's keys: those in order first, as there, then the rest sorted.
func orderedKeys(m map[string]any, order []string) []string {
	var keys []string
	seen := map[string]bool{}
	for _, key := range order {
		if _, ok := m[key]; ok && !seen[key] {
			seen[key] = true
			keys = append(keys, key)
		}
	}
	var rest []string
	for key := range m {
		if !seen[key] {
			rest = append(rest, key)
		}
	}
	sort.Strings(rest)
	return append(keys, rest...)
}

// stringList reads a JSON array of strings; other items are skipped.
func stringList(value any) []string {
	items, _ := value.([]any)
	var out []string
	for _, item := range items {
		if s, ok := item.(string); ok {
			out = append(out, s)
		}
	}
	return out
}

func anyList(ids []string) []any {
	out := make([]any, len(ids))
	for i, id := range ids {
		out[i] = id
	}
	return out
}

// unique returns ids without repeats, in first-seen order.
func unique(ids []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, id := range ids {
		if !seen[id] {
			seen[id] = true
			out = append(out, id)
		}
	}
	return out
}

// keepOrder returns ids with the ones in loaded first, in that order, then the rest as they are.
func keepOrder(ids, loaded []string) []string {
	rank := func(id string) int {
		for i, l := range loaded {
			if l == id {
				return i
			}
		}
		return len(loaded)
	}
	out := append([]string{}, ids...)
	sort.SliceStable(out, func(i, j int) bool { return rank(out[i]) < rank(out[j]) })
	return out
}

// sameTool reports whether a and b are the same tool: the same pointer, for the tools the
// agenttools package returns.
func sameTool(a, b agenttools.Tool) (same bool) {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	defer func() {
		if recover() != nil {
			same = false
		}
	}()
	return a == b
}

// containsTool reports whether tools holds tool itself.
func containsTool(tools []agenttools.Tool, tool agenttools.Tool) bool {
	for _, t := range tools {
		if sameTool(t, tool) {
			return true
		}
	}
	return false
}

// toolDefinition is a tool's definition as this SDK's generated model holds it, as text to
// compare: its BuildConfig, marshaled, which leaves out fields the generated types don't declare.
// The tool builder keeps such fields from a loaded config, but a tool written in code can't have
// them, so they don't count when telling whether two definitions are the same tool. ok is false
// if the tool is invalid.
func toolDefinition(tool agenttools.Tool) (definition string, ok bool) {
	config, err := tool.BuildConfig(nil)
	if err != nil {
		return "", false
	}
	data, err := json.Marshal(config)
	if err != nil {
		return "", false
	}
	tree, err := wirejson.Decode(data)
	if err != nil {
		return "", false
	}
	text, err := wirejson.String(tree)
	if err != nil {
		return "", false
	}
	return text, true
}

// toolText is a tool's ToJSONString, or "" if it's invalid.
func toolText(tool agenttools.Tool) string {
	text, err := tool.ToJSONString()
	if err != nil {
		return ""
	}
	return text
}
