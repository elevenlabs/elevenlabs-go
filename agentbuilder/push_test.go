package agentbuilder_test

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agentbuilder"
	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/core"
	"github.com/elevenlabs/elevenlabs-go/option"
)

var ctx = context.Background()

// supportAgent is a support agent whose workflow's lookup node runs tools[0], and whose own tools
// are end_call plus the rest of tools.
func supportAgent(tools ...agenttools.Tool) *agentbuilder.Agent {
	if len(tools) == 0 {
		tools = []agenttools.Tool{getOrder("Look up an order by its number"), showOrder()}
	}
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
		Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{tools[0]}}).
		OnSuccess(nil).
		Converse("resolve", "Explain the order", nil)
	return agentbuilder.New("Support line", &agentbuilder.Options{
		Prompt:   ptr("You help with orders."),
		Tools:    append([]agenttools.Tool{agenttools.SystemEndCall()}, tools[1:]...),
		Workflow: wf,
	})
}

func toolJSON(t *testing.T, tool agenttools.Tool) map[string]any {
	t.Helper()
	wire, err := tool.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	return wire
}

func fromFiles(t *testing.T, files map[string]string) *agentbuilder.Project {
	t.Helper()
	project, err := agentbuilder.ProjectFromFiles(files)
	if err != nil {
		t.Fatal(err)
	}
	return project
}

func filesOf(t *testing.T, project *agentbuilder.Project) map[string]string {
	t.Helper()
	files, err := project.ToFiles()
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func keysOf(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func mustPush(t *testing.T, f *fakeServer, agent *agentbuilder.Agent, opts *agentbuilder.PushOptions) *agentbuilder.PushResult {
	t.Helper()
	result, err := f.client().Agents.Push(ctx, agent, opts)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func TestPushCreatesToolsThenTheAgent(t *testing.T) {
	f := newFakeServer(t)
	result := mustPush(t, f, supportAgent(), nil)

	if got := f.writes(); !reflect.DeepEqual(got, []string{"POST /v1/convai/tools", "POST /v1/convai/tools", "POST /v1/convai/agents/create"}) {
		t.Fatalf("writes = %v", got)
	}
	want := &agentbuilder.PushResult{
		AgentID: "agent_new_3",
		Action:  agentbuilder.AgentCreated,
		Tools: map[string]agentbuilder.ToolResult{
			"show_order": {ID: "tool_new_1", Action: agentbuilder.ToolCreated},
			"get_order":  {ID: "tool_new_2", Action: agentbuilder.ToolCreated},
		},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("result = %+v", result)
	}
	sent := f.agent("agent_new_3")
	if got := at(sent, "conversation_config", "agent", "prompt", "tool_ids"); !reflect.DeepEqual(got, []any{"tool_new_1"}) {
		t.Fatalf("tool_ids = %v", got)
	}
	if got := at(sent, "conversation_config", "agent", "prompt", "built_in_tools", "end_call", "name"); got != "end_call" {
		t.Fatalf("end_call = %v", got)
	}
	if got := at(sent, "workflow", "nodes", "lookup", "tools"); !reflect.DeepEqual(got, []any{map[string]any{"tool_id": "tool_new_2"}}) {
		t.Fatalf("lookup tools = %v", got)
	}
}

func TestPushUpdatesOnlyWhatChanged(t *testing.T) {
	f := newFakeServer(t)
	first := mustPush(t, f, supportAgent(), nil)
	// Another agent's tool with the same name is never touched.
	f.addTool("tool_other", toolJSON(t, getOrder("Someone else's")), "agent_other")
	f.resetCalls()

	again := mustPush(t, f, supportAgent(getOrder("Find an order"), showOrder()), &agentbuilder.PushOptions{AgentID: first.AgentID})

	if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/tools/tool_new_2", "PATCH /v1/convai/agents/" + first.AgentID}) {
		t.Fatalf("writes = %v", got)
	}
	if got := f.callList()[0].query; got.Get("used_by_agent_id") != first.AgentID || got.Get("page_size") != "100" || len(got) != 2 {
		t.Fatalf("list query = %v", got)
	}
	want := map[string]agentbuilder.ToolResult{
		"show_order": {ID: "tool_new_1", Action: agentbuilder.ToolUnchanged},
		"get_order":  {ID: "tool_new_2", Action: agentbuilder.ToolUpdated},
	}
	if !reflect.DeepEqual(again.Tools, want) || again.Action != agentbuilder.AgentUpdated {
		t.Fatalf("result = %+v", again)
	}
	if got := at(f.toolConfig("tool_other"), "description"); got != "Someone else's" {
		t.Fatalf("the other agent's tool changed: %v", got)
	}
}

func TestPushReportsDetachedTools(t *testing.T) {
	f := newFakeServer(t)
	first := mustPush(t, f, supportAgent(), nil)
	agent := supportAgent().Update(&agentbuilder.Options{Tools: []agenttools.Tool{}})
	result := mustPush(t, f, agent, &agentbuilder.PushOptions{AgentID: first.AgentID})
	if got := result.Tools["show_order"]; got != (agentbuilder.ToolResult{ID: "tool_new_1", Action: agentbuilder.ToolDetached}) {
		t.Fatalf("show_order = %+v", got)
	}
	if f.toolConfig("tool_new_1") == nil {
		t.Fatal("a detached tool must not be deleted")
	}
}

func TestPushToABranchMatchesTheBranchsTools(t *testing.T) {
	f := newFakeServer(t)
	first := mustPush(t, f, supportAgent(), nil)
	// The branch uses its own copy of get_order, and no show_order.
	f.addTool("tool_branch_order", toolJSON(t, getOrder("Branch lookup")))
	branch := toTree(t, f.agent(first.AgentID))
	at(branch, "workflow", "nodes", "lookup").(tree)["tools"] = []any{tree{"tool_id": "tool_branch_order"}}
	at(branch, "conversation_config", "agent", "prompt").(tree)["tool_ids"] = []any{}
	f.mu.Lock()
	f.branches[first.AgentID+"/agtbrch_1"] = branch
	f.mu.Unlock()
	f.resetCalls()

	result := mustPush(t, f, supportAgent(getOrder("Changed lookup")), &agentbuilder.PushOptions{AgentID: first.AgentID, BranchID: "agtbrch_1", VersionDescription: "Branch change"})
	if got := result.Tools["get_order"]; got != (agentbuilder.ToolResult{ID: "tool_branch_order", Action: agentbuilder.ToolUpdated}) {
		t.Fatalf("get_order = %+v", got)
	}
	// A main-branch tool the branch never used isn't detached.
	if _, ok := result.Tools["show_order"]; ok {
		t.Fatalf("show_order = %+v", result.Tools["show_order"])
	}
	if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/tools/tool_branch_order", "PATCH /v1/convai/agents/" + first.AgentID}) {
		t.Fatalf("writes = %v", got)
	}
	calls := f.callList()
	for _, c := range calls {
		if c.query.Has("used_by_agent_id") {
			t.Fatalf("a branch push must not list the main branch's tools: %v", c)
		}
	}
	last := calls[len(calls)-1]
	if last.query.Get("branch_id") != "agtbrch_1" || last.body["version_description"] != "Branch change" {
		t.Fatalf("agent update = %+v", last)
	}
}

func TestPushValidatesBeforeAnyRequest(t *testing.T) {
	f := newFakeServer(t)
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Tools: []agenttools.Tool{agenttools.Webhook("broken", "No URL")}})
	_, err := f.client().Agents.Push(ctx, agent, nil)
	var verr *agentbuilder.ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("err = %v", err)
	}
	if len(f.callList()) != 0 {
		t.Fatalf("calls = %v", f.callList())
	}
}

func TestPushReportsWhatItWroteBeforeAFailure(t *testing.T) {
	f := newFakeServer(t)
	f.failOn = regexp.MustCompile(`^POST /v1/convai/agents/create$`)
	// Request options reach every request: here, no retries.
	_, err := f.client().Agents.Push(ctx, supportAgent(), nil, option.WithoutRetries())
	var perr *agentbuilder.PushError
	if !errors.As(err, &perr) {
		t.Fatalf("err = %v", err)
	}
	if !strings.HasPrefix(perr.Error(), `Couldn't create agent "Support line"`) {
		t.Fatalf("message = %q", perr.Error())
	}
	want := agentbuilder.PushResult{Tools: map[string]agentbuilder.ToolResult{
		"show_order": {ID: "tool_new_1", Action: agentbuilder.ToolCreated},
		"get_order":  {ID: "tool_new_2", Action: agentbuilder.ToolCreated},
	}}
	if !reflect.DeepEqual(perr.Result, want) {
		t.Fatalf("result = %+v", perr.Result)
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 400 {
		t.Fatalf("cause = %v", perr.Cause)
	}
}

func TestPushRefusesAgentIDWithAProject(t *testing.T) {
	f := newFakeServer(t)
	_, err := f.client().Agents.Push(ctx, supportAgent(), &agentbuilder.PushOptions{AgentID: "agent_1", Project: fromFiles(t, nil)})
	if err == nil || !strings.Contains(err.Error(), "not both") {
		t.Fatalf("err = %v", err)
	}
}

func TestPushWithAProjectRecordsIDsAndUpdatesByThem(t *testing.T) {
	f := newFakeServer(t)
	project := fromFiles(t, nil)
	mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{Project: project})

	files := filesOf(t, project)
	if got := keysOf(files); !reflect.DeepEqual(got, []string{"agent_configs/Support-line.json", "agents.json", "tool_configs/get_order.json", "tool_configs/show_order.json", "tools.json"}) {
		t.Fatalf("files = %v", got)
	}
	if got := files["agents.json"]; got != "{\n    \"agents\": [\n        {\n            \"config\": \"agent_configs/Support-line.json\",\n            \"id\": \"agent_new_3\"\n        }\n    ]\n}" {
		t.Fatalf("agents.json = %q", got)
	}
	if got := files["tools.json"]; got != `{
    "tools": [
        {
            "type": "client",
            "config": "tool_configs/show_order.json",
            "id": "tool_new_1"
        },
        {
            "type": "webhook",
            "config": "tool_configs/get_order.json",
            "id": "tool_new_2"
        }
    ]
}` {
		t.Fatalf("tools.json = %s", got)
	}
	if got := at(toTree(t, files["agent_configs/Support-line.json"]), "conversation_config", "agent", "prompt", "tool_ids"); !reflect.DeepEqual(got, []any{"tool_new_1"}) {
		t.Fatalf("tool_ids = %v", got)
	}

	// A fresh run: new objects, matched to the project's entries by name.
	f.resetCalls()
	reloaded := fromFiles(t, files)
	result := mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{Project: reloaded})
	if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/agents/agent_new_3"}) {
		t.Fatalf("writes = %v", got)
	}
	if result.Tools["get_order"].Action != agentbuilder.ToolUnchanged {
		t.Fatalf("get_order = %+v", result.Tools["get_order"])
	}
	// Only the version the update reported is new, recorded like the CLI does.
	if got := reloaded.ChangedFiles(); !reflect.DeepEqual(got, []string{"agents.json"}) {
		t.Fatalf("changed = %v", got)
	}
	if got := at(toTree(t, filesOf(t, reloaded)["agents.json"])["agents"].([]any)[0], "version_id"); got != "ver_1" {
		t.Fatalf("version_id = %v", got)
	}
}

func TestPushSendsALoadedAgentsUnmodeledSettings(t *testing.T) {
	f := newFakeServer(t)
	f.addTool("tool_lookup", `{"type": "client", "name": "lookup", "description": "old"}`)
	f.addAgent("agent_1", `{}`)
	project := fromFiles(t, map[string]string{
		"agents.json":          `{"agents": [{"config": "agent_configs/a.json", "id": "agent_1"}]}`,
		"tools.json":           `{"tools": [{"type": "client", "config": "tool_configs/t.json", "id": "tool_lookup"}]}`,
		"agent_configs/a.json": `{"name": "A", "conversation_config": {"agent": {"prompt": {"prompt": "p", "tool_ids": ["tool_lookup"]}}, "future_setting": {"on": true}}, "platform_settings": {}, "tags": []}`,
		"tool_configs/t.json":  `{"type": "client", "name": "lookup", "description": "d", "future_field": 1}`,
	})
	agent, err := project.Agent("A")
	if err != nil {
		t.Fatal(err)
	}
	agent.Update(&agentbuilder.Options{FirstMessage: ptr("Hello")})
	mustPush(t, f, agent, &agentbuilder.PushOptions{Project: project})

	var toolPatch, agentPatch *call
	for _, c := range f.callList() {
		c := c
		if c.method == "PATCH" && strings.Contains(c.path, "/tools/") {
			toolPatch = &c
		}
		if c.method == "PATCH" && strings.Contains(c.path, "/agents/") {
			agentPatch = &c
		}
	}
	if toolPatch == nil || at(toolPatch.body, "tool_config", "future_field") != json.Number("1") {
		t.Fatalf("tool patch = %+v", toolPatch)
	}
	if agentPatch == nil || !reflect.DeepEqual(at(agentPatch.body, "conversation_config", "future_setting"), tree{"on": true}) {
		t.Fatalf("agent patch = %+v", agentPatch)
	}
	if got := at(agentPatch.body, "conversation_config", "agent", "first_message"); got != "Hello" {
		t.Fatalf("first_message = %v", got)
	}
	changed := project.ChangedFiles()
	sort.Strings(changed)
	if !reflect.DeepEqual(changed, []string{"agent_configs/a.json", "agents.json", "tool_configs/t.json", "tools.json"}) {
		t.Fatalf("changed = %v", changed)
	}
}

func TestPullLoadsAnAgentAndWritesItIntoAProject(t *testing.T) {
	f := newFakeServer(t)
	first := mustPush(t, f, supportAgent(), nil)

	project := fromFiles(t, nil)
	agent, err := f.client().Agents.Pull(ctx, first.AgentID, &agentbuilder.PullOptions{Project: project})
	if err != nil {
		t.Fatal(err)
	}
	if agent.Name() != "Support line" {
		t.Fatalf("Name() = %q", agent.Name())
	}
	var names []string
	for _, tool := range agent.Tools() {
		names = append(names, tool.Name())
	}
	if !reflect.DeepEqual(names, []string{"show_order", "get_order"}) {
		t.Fatalf("Tools() = %v", names)
	}
	if id, _ := project.AgentID(agent); id != first.AgentID {
		t.Fatalf("AgentID = %q", id)
	}
	files := filesOf(t, project)
	if got := keysOf(files); !reflect.DeepEqual(got, []string{"agent_configs/Support-line.json", "agents.json", "tool_configs/get_order.json", "tool_configs/show_order.json", "tools.json"}) {
		t.Fatalf("files = %v", got)
	}
	var config map[string]json.RawMessage
	if err := json.Unmarshal([]byte(files["agent_configs/Support-line.json"]), &config); err != nil {
		t.Fatal(err)
	}
	var keys []string
	for key := range config {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"conversation_config", "name", "platform_settings", "tags", "workflow"}) {
		t.Fatalf("pulled config keys = %v", keys)
	}

	// Pushing the pulled agent back changes no tools.
	f.resetCalls()
	result := mustPush(t, f, agent, &agentbuilder.PushOptions{Project: project})
	if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/agents/" + first.AgentID}) {
		t.Fatalf("writes = %v", got)
	}
	for name, tool := range result.Tools {
		if tool.Action != agentbuilder.ToolUnchanged {
			t.Fatalf("%s = %+v", name, tool)
		}
	}
	if len(result.Tools) != 2 {
		t.Fatalf("tools = %+v", result.Tools)
	}
}

func TestPushUpdatesAProjectToolWhoseOnlyChangeIsUnmodeled(t *testing.T) {
	f := newFakeServer(t)
	f.addTool("tool_1", `{"type": "client", "name": "ping", "description": "d", "future_field": 1}`, "agent_1")
	f.addAgent("agent_1", `{}`)
	project := fromFiles(t, map[string]string{
		"agents.json":          `{"agents": [{"config": "agent_configs/a.json", "id": "agent_1"}]}`,
		"tools.json":           `{"tools": [{"type": "client", "config": "tool_configs/t.json", "id": "tool_1"}]}`,
		"agent_configs/a.json": `{"name": "A", "conversation_config": {"agent": {"prompt": {"prompt": "p", "tool_ids": ["tool_1"]}}}}`,
		"tool_configs/t.json":  `{"type": "client", "name": "ping", "description": "d", "future_field": 2}`,
	})
	agent, _ := project.Agent("A")
	result := mustPush(t, f, agent, &agentbuilder.PushOptions{Project: project})
	if result.Tools["ping"].Action != agentbuilder.ToolUpdated {
		t.Fatalf("ping = %+v", result.Tools["ping"])
	}
	if got := at(f.toolConfig("tool_1"), "future_field"); got != json.Number("2") {
		t.Fatalf("future_field = %v", got)
	}
}

func TestPushBringsTheProjectUpToDateWhenTheServerMatches(t *testing.T) {
	f := newFakeServer(t)
	project := fromFiles(t, nil)
	mustPush(t, f, supportAgent(getOrder("v2"), showOrder()), &agentbuilder.PushOptions{Project: project})
	files := filesOf(t, project)
	old, err := getOrder("v1").ToJSONString()
	if err != nil {
		t.Fatal(err)
	}
	files["tool_configs/get_order.json"] = old
	stale := fromFiles(t, files)
	result := mustPush(t, f, supportAgent(getOrder("v2"), showOrder()), &agentbuilder.PushOptions{Project: stale})
	if result.Tools["get_order"].Action != agentbuilder.ToolUnchanged {
		t.Fatalf("get_order = %+v", result.Tools["get_order"])
	}
	if got := at(toTree(t, filesOf(t, stale)["tool_configs/get_order.json"]), "description"); got != "v2" {
		t.Fatalf("description = %v", got)
	}
}

func TestPushDoesntDetachAToolTheWorkflowStillUses(t *testing.T) {
	f := newFakeServer(t)
	f.addTool("tool_raw", toolJSON(t, showOrder()))
	wf := agentworkflows.New(nil)
	wf.Tool("run", []string{"tool_raw"}, nil)
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Workflow: wf})
	first := mustPush(t, f, agent, nil)
	again := mustPush(t, f, agent, &agentbuilder.PushOptions{AgentID: first.AgentID})
	if len(again.Tools) != 0 {
		t.Fatalf("tools = %+v", again.Tools)
	}
}

func TestPushKeepsUnmodeledFieldsOfAToolRebuiltInCode(t *testing.T) {
	f := newFakeServer(t)
	project := fromFiles(t, nil)
	mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{Project: project})
	files := filesOf(t, project)
	withExtra := toTree(t, files["tool_configs/get_order.json"])
	withExtra["future_field"] = json.Number("1")
	text, _ := json.MarshalIndent(withExtra, "", "    ")
	files["tool_configs/get_order.json"] = string(text)
	f.mu.Lock()
	f.tools["tool_new_2"].config = toTree(t, withExtra)
	f.mu.Unlock()

	reloaded := fromFiles(t, files)
	result := mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{Project: reloaded})
	if result.Tools["get_order"].Action != agentbuilder.ToolUnchanged {
		t.Fatalf("get_order = %+v", result.Tools["get_order"])
	}
	if got := at(toTree(t, filesOf(t, reloaded)["tool_configs/get_order.json"]), "future_field"); got != json.Number("1") {
		t.Fatalf("future_field = %v", got)
	}
}

func TestPushFollowsTheListCursor(t *testing.T) {
	f := newFakeServer(t)
	first := mustPush(t, f, supportAgent(), nil)
	f.mu.Lock()
	f.pageSize = 1
	f.mu.Unlock()
	f.resetCalls()
	result := mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{AgentID: first.AgentID})
	var cursors []string
	for _, c := range f.callList() {
		if c.path == "/v1/convai/tools" {
			cursors = append(cursors, c.query.Get("cursor"))
		}
	}
	if !reflect.DeepEqual(cursors, []string{"", "1"}) {
		t.Fatalf("list cursors = %q", cursors)
	}
	if result.Tools["show_order"].Action != agentbuilder.ToolUnchanged || result.Tools["get_order"].Action != agentbuilder.ToolUnchanged {
		t.Fatalf("result = %+v", result)
	}
	if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/agents/" + first.AgentID}) {
		t.Fatalf("writes = %v", got)
	}
}

func TestRootClientWiresTheAgentsWrapper(t *testing.T) {
	f := newFakeServer(t)
	c := f.client()
	var _ *agentbuilder.AgentsClient = c.Agents
	f.addAgent("agent_1", `{"name": "A", "conversation_config": {}}`)
	// The generated methods are still there.
	response, err := c.Agents.Get(ctx, "agent_1", nil)
	if err != nil || response.AgentId != "agent_1" {
		t.Fatalf("generated Get: %+v %v", response, err)
	}
}

const mcpToolJSON = `{"type": "mcp", "name": "docs_search", "description": "Search the docs", "mcp_server_id": "mcp_1"}`

func mcpTool(t *testing.T) agenttools.Tool {
	t.Helper()
	tool, err := agenttools.FromJSON(mcpToolJSON)
	if err != nil {
		t.Fatal(err)
	}
	return tool
}

// A tool Push doesn't manage needs the id it was loaded with: given as a new object, it's reported
// before any request, whether the agent or a workflow node uses it.
func TestPushRejectsAnMCPToolWithoutAnID(t *testing.T) {
	f := newFakeServer(t)
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p")}).AddTool(mcpTool(t))
	if got := issuesOf(t, agent.Validate(nil)); !reflect.DeepEqual(got, []issue{{"invalid_tool", "tools.docs_search"}}) {
		t.Fatalf("issues = %v", got)
	}
	_, err := f.client().Agents.Push(ctx, agent, nil)
	if got := issuesOf(t, err); !reflect.DeepEqual(got, []issue{{"invalid_tool", "tools.docs_search"}}) {
		t.Fatalf("issues = %v", got)
	}
	var verr *agentbuilder.ValidationError
	errors.As(err, &verr)
	if want := `Tool "docs_search" has type "mcp", which push doesn't create or update: load it with its id or attach it with ToolIDs`; verr.Issues[0].Message != want {
		t.Fatalf("message = %q", verr.Issues[0].Message)
	}

	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		Tool("search", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{mcpTool(t)}})
	inWorkflow := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Workflow: wf})
	if got := issuesOf(t, inWorkflow.Validate(nil)); !reflect.DeepEqual(got, []issue{{"invalid_tool", "tools.docs_search"}}) {
		t.Fatalf("issues = %v", got)
	}
	if len(f.callList()) != 0 {
		t.Fatalf("calls = %v", f.callList())
	}
}

// Push manages webhook and client tools only: an MCP tool, like Pull leaves it, stays attached by
// the id it was loaded with and is never created or updated.
func TestPushLeavesAnMCPToolAttachedByID(t *testing.T) {
	check := func(t *testing.T, f *fakeServer, agent *agentbuilder.Agent, opts *agentbuilder.PushOptions) {
		t.Helper()
		if got := agent.Tools(); len(got) != 0 {
			t.Fatalf("Tools() = %v", names(got))
		}
		result := mustPush(t, f, agent, opts)
		if got := f.writes(); !reflect.DeepEqual(got, []string{"PATCH /v1/convai/agents/agent_1"}) {
			t.Fatalf("writes = %v", got)
		}
		if _, ok := result.Tools["docs_search"]; ok {
			t.Fatalf("result.Tools = %+v", result.Tools)
		}
		sent := f.agent("agent_1")
		if got := at(sent, "conversation_config", "agent", "prompt", "tool_ids"); !reflect.DeepEqual(got, []any{"tool_mcp"}) {
			t.Fatalf("tool_ids = %v", got)
		}
		if got := f.toolConfig("tool_mcp"); !reflect.DeepEqual(got, toTree(t, mcpToolJSON)) {
			t.Fatalf("mcp tool = %v", got)
		}
	}
	config := `{"name": "A", "conversation_config": {"agent": {"prompt": {"prompt": "p", "tool_ids": ["tool_mcp"]}}}, "platform_settings": {}, "tags": []}`

	t.Run("loaded with its id", func(t *testing.T) {
		f := newFakeServer(t)
		f.addTool("tool_mcp", mcpToolJSON, "agent_1")
		f.addAgent("agent_1", `{}`)
		agent, err := agentbuilder.FromJSON(config, &agentbuilder.LoadOptions{ToolsByID: map[string]agenttools.Tool{"tool_mcp": mcpTool(t)}})
		if err != nil {
			t.Fatal(err)
		}
		check(t, f, agent, &agentbuilder.PushOptions{AgentID: "agent_1"})
	})

	t.Run("from a project's tools.json", func(t *testing.T) {
		f := newFakeServer(t)
		f.addTool("tool_mcp", mcpToolJSON, "agent_1")
		f.addAgent("agent_1", `{}`)
		project := fromFiles(t, map[string]string{
			"agents.json":            `{"agents": [{"config": "agent_configs/a.json", "id": "agent_1"}]}`,
			"tools.json":             `{"tools": [{"type": "mcp", "config": "tool_configs/docs.json", "id": "tool_mcp"}]}`,
			"agent_configs/a.json":   config,
			"tool_configs/docs.json": mcpToolJSON,
		})
		agent, err := project.Agent("A")
		if err != nil {
			t.Fatal(err)
		}
		check(t, f, agent, &agentbuilder.PushOptions{Project: project})
	})

	t.Run("in a workflow node", func(t *testing.T) {
		f := newFakeServer(t)
		f.addTool("tool_mcp", mcpToolJSON, "agent_1")
		f.addAgent("agent_1", `{}`)
		mcp := mcpTool(t)
		agent, err := agentbuilder.FromJSON(config, &agentbuilder.LoadOptions{ToolsByID: map[string]agenttools.Tool{"tool_mcp": mcp}})
		if err != nil {
			t.Fatal(err)
		}
		wf := agentworkflows.New(nil)
		wf.Converse("triage", "Find out what the caller needs", nil).
			Tool("search", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{mcp}})
		agent.Update(&agentbuilder.Options{Workflow: wf})
		check(t, f, agent, &agentbuilder.PushOptions{AgentID: "agent_1"})
		if got := at(f.agent("agent_1"), "workflow", "nodes", "search", "tools"); !reflect.DeepEqual(got, []any{map[string]any{"tool_id": "tool_mcp"}}) {
			t.Fatalf("search tools = %v", got)
		}
	})
}
