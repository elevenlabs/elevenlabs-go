package agentbuilder_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/client"
	"github.com/elevenlabs/elevenlabs-go/option"
)

type tree = map[string]any

// call is one request the fake server got.
type call struct {
	method, path string
	query        url.Values
	body         tree
}

type fakeTool struct {
	config tree
	usedBy []string
}

// fakeServer is an in-memory stand-in for the agents and tools endpoints Push and Pull use,
// modelled on the reference SDK's tests/custom/agents/agent/fakeServer.ts.
type fakeServer struct {
	t      *testing.T
	server *httptest.Server
	mu     sync.Mutex
	// tools and toolOrder hold the tools, listed in creation order.
	tools     map[string]*fakeTool
	toolOrder []string
	agents    map[string]tree
	// branches holds branch configs by "agentID/branchID". Only the main branch counts for
	// used_by_agent_id, as in the API.
	branches map[string]tree
	calls    []call
	// failOn fails every request whose "METHOD path" matches.
	failOn *regexp.Regexp
	// pageSize, when set, pages the tools list, with the next index as the cursor.
	pageSize int
	next     int
}

func newFakeServer(t *testing.T) *fakeServer {
	f := &fakeServer{t: t, tools: map[string]*fakeTool{}, agents: map[string]tree{}, branches: map[string]tree{}, next: 1}
	f.server = httptest.NewServer(http.HandlerFunc(f.handle))
	t.Cleanup(f.server.Close)
	return f
}

func (f *fakeServer) client() *client.Client {
	key := "test"
	return client.NewClient(option.WithBaseURL(f.server.URL), option.WithApiKey(&key))
}

func (f *fakeServer) addTool(id string, config any, usedBy ...string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.tools[id]; !ok {
		f.toolOrder = append(f.toolOrder, id)
	}
	f.tools[id] = &fakeTool{config: toTree(f.t, config), usedBy: usedBy}
}

func (f *fakeServer) addAgent(id string, config any) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.agents[id] = toTree(f.t, config)
}

func (f *fakeServer) resetCalls() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = nil
}

func (f *fakeServer) callList() []call {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]call{}, f.calls...)
}

// writes are the calls that change something, as "METHOD path".
func (f *fakeServer) writes() []string {
	var out []string
	for _, c := range f.callList() {
		if c.method != http.MethodGet {
			out = append(out, c.method+" "+c.path)
		}
	}
	return out
}

func (f *fakeServer) toolConfig(id string) tree {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.tools[id].config
}

func (f *fakeServer) agent(id string) tree {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.agents[id]
}

func toTree(t *testing.T, value any) tree {
	t.Helper()
	data, ok := value.([]byte)
	if !ok {
		if s, isString := value.(string); isString {
			data = []byte(s)
		} else {
			var err error
			if data, err = json.Marshal(value); err != nil {
				t.Fatal(err)
			}
		}
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var out tree
	if err := dec.Decode(&out); err != nil {
		t.Fatal(err)
	}
	return out
}

func (f *fakeServer) handle(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var body tree
	if r.Body != nil {
		dec := json.NewDecoder(r.Body)
		dec.UseNumber()
		_ = dec.Decode(&body)
	}
	if r.Header.Get("xi-api-key") != "test" {
		reply(w, 401, tree{"detail": "no api key"})
		return
	}
	f.calls = append(f.calls, call{method: r.Method, path: r.URL.Path, query: r.URL.Query(), body: body})
	if f.failOn != nil && f.failOn.MatchString(r.Method+" "+r.URL.Path) {
		// A 4xx, so the generated retrier doesn't wait before giving up.
		reply(w, 400, tree{"detail": "boom"})
		return
	}
	parts := strings.Split(r.URL.Path, "/") // /v1/convai/<kind>/<id>
	kind, id := parts[3], ""
	if len(parts) > 4 {
		id = parts[4]
	}
	switch kind {
	case "tools":
		switch {
		case r.Method == http.MethodGet && id == "":
			agentID := r.URL.Query().Get("used_by_agent_id")
			tools := []any{}
			for _, toolID := range f.toolOrder {
				tool := f.tools[toolID]
				if agentID == "" || contains(tool.usedBy, agentID) {
					tools = append(tools, toolResponse(toolID, tool.config))
				}
			}
			page := tree{"tools": tools, "has_more": false}
			if f.pageSize > 0 {
				start, _ := strconv.Atoi(r.URL.Query().Get("cursor"))
				end := start + f.pageSize
				if end < len(tools) {
					page = tree{"tools": tools[start:end], "has_more": true, "next_cursor": strconv.Itoa(end)}
				} else {
					page = tree{"tools": tools[start:], "has_more": false, "next_cursor": nil}
				}
			}
			reply(w, 200, page)
		case r.Method == http.MethodPost:
			toolID := fmt.Sprintf("tool_new_%d", f.next)
			f.next++
			config, _ := body["tool_config"].(tree)
			f.tools[toolID] = &fakeTool{config: config}
			f.toolOrder = append(f.toolOrder, toolID)
			reply(w, 200, toolResponse(toolID, config))
		default:
			tool, ok := f.tools[id]
			if !ok {
				reply(w, 404, tree{"detail": "not found"})
				return
			}
			if r.Method == http.MethodPatch {
				tool.config, _ = body["tool_config"].(tree)
			}
			reply(w, 200, toolResponse(id, tool.config))
		}
	case "agents":
		if r.Method == http.MethodPost && id == "create" {
			agentID := fmt.Sprintf("agent_new_%d", f.next)
			f.next++
			f.storeAgent(agentID, body)
			reply(w, 200, tree{"agent_id": agentID})
			return
		}
		if branchID := r.URL.Query().Get("branch_id"); branchID != "" {
			key := id + "/" + branchID
			if r.Method == http.MethodPatch {
				f.branches[key] = merged(f.branches[key], body)
			}
			branch, ok := f.branches[key]
			if !ok {
				reply(w, 404, tree{"detail": "not found"})
				return
			}
			reply(w, 200, agentResponse(id, branch))
			return
		}
		if r.Method == http.MethodPatch {
			f.storeAgent(id, merged(f.agents[id], body))
			reply(w, 200, agentResponse(id, f.agents[id]))
			return
		}
		agent, ok := f.agents[id]
		if !ok {
			reply(w, 404, tree{"detail": "not found"})
			return
		}
		reply(w, 200, agentResponse(id, agent))
	default:
		f.t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		reply(w, 400, tree{"detail": "unexpected"})
	}
}

// storeAgent stores an agent's config and records which tools its main branch uses.
func (f *fakeServer) storeAgent(agentID string, config tree) {
	f.agents[agentID] = config
	var ids []string
	if list, ok := at(config, "conversation_config", "agent", "prompt", "tool_ids").([]any); ok {
		for _, id := range list {
			ids = append(ids, id.(string))
		}
	}
	nodes, _ := at(config, "workflow", "nodes").(tree)
	for _, value := range nodes {
		node := value.(tree)
		tools, _ := node["tools"].([]any)
		for _, locator := range tools {
			ids = append(ids, locator.(tree)["tool_id"].(string))
		}
		extra, _ := node["additional_tool_ids"].([]any)
		for _, id := range extra {
			ids = append(ids, id.(string))
		}
	}
	for toolID, tool := range f.tools {
		var usedBy []string
		for _, a := range tool.usedBy {
			if a != agentID {
				usedBy = append(usedBy, a)
			}
		}
		if contains(ids, toolID) {
			usedBy = append(usedBy, agentID)
		}
		tool.usedBy = usedBy
	}
}

func merged(base, update tree) tree {
	out := tree{}
	for k, v := range base {
		out[k] = v
	}
	for k, v := range update {
		out[k] = v
	}
	return out
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func toolResponse(id string, config tree) tree {
	return tree{
		"id":          id,
		"tool_config": config,
		"access_info": tree{"is_creator": true, "creator_name": "a", "creator_email": "a@example.com", "role": "admin"},
		"usage_stats": tree{"avg_latency_secs": 0.1},
	}
}

// withServerDefaults adds the defaults the API fills into a stored workflow.
func withServerDefaults(workflow tree) tree {
	nodes := tree{}
	for id, value := range at(workflow, "nodes").(tree) {
		node := value.(tree)
		if node["type"] == "override_agent" {
			node = merged(tree{
				"conversation_config":          tree{},
				"additional_knowledge_base":    []any{},
				"additional_tool_ids":          []any{},
				"entry_behavior":               "generate_immediately",
				"allow_interruptions_on_entry": true,
			}, node)
		}
		nodes[id] = node
	}
	edges := tree{}
	if all, ok := at(workflow, "edges").(tree); ok {
		for id, edge := range all {
			edges[id] = merged(tree{"reroutes": []any{}}, edge.(tree))
		}
	}
	return merged(tree{"prevent_subagent_loops": false}, merged(workflow, tree{"nodes": nodes, "edges": edges}))
}

func agentResponse(id string, config tree) tree {
	out := merged(tree{
		"agent_id":             id,
		"metadata":             tree{"created_at_unix_secs": 1, "updated_at_unix_secs": 1},
		"created_at_unix_secs": 1,
		"updated_at_unix_secs": 1,
		"version_id":           "ver_1",
	}, config)
	if workflow, ok := config["workflow"].(tree); ok {
		out["workflow"] = withServerDefaults(workflow)
	}
	return out
}

func reply(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
