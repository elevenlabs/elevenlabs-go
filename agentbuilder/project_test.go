package agentbuilder_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agentbuilder"
	"github.com/elevenlabs/elevenlabs-go/agenttools"
)

// readTree reads every file under dir, keyed by slash-separated relative path.
func readTree(t *testing.T, dir string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(rel)] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func names[T interface{ Name() string }](items []T) []string {
	out := make([]string, len(items))
	for i, item := range items {
		out[i] = item.Name()
	}
	return out
}

func projectError(t *testing.T, err error, pattern string) {
	t.Helper()
	var perr *agentbuilder.ProjectError
	if !errors.As(err, &perr) {
		t.Fatalf("expected a *ProjectError matching %q, got %v", pattern, err)
	}
	if !regexp.MustCompile(pattern).MatchString(perr.Error()) {
		t.Fatalf("error %q doesn't match %q", perr.Error(), pattern)
	}
}

// Fixture check: the CLI project loads with the same agents, tools, ids and attached ids.
func TestFixtureProjectLoads(t *testing.T) {
	project := fromFiles(t, readTree(t, cliProject))
	if got := names(project.Agents()); !reflect.DeepEqual(got, []string{"Support", "Billing"}) {
		t.Fatalf("Agents() = %v", got)
	}
	if got := names(project.Tools()); !reflect.DeepEqual(got, []string{"lookup_order", "issue_refund"}) {
		t.Fatalf("Tools() = %v", got)
	}
	support, err := project.Agent("Support")
	if err != nil {
		t.Fatal(err)
	}
	if again, _ := project.Agent("Support"); again != support {
		t.Fatal("Agent(name) must return the same object every time")
	}
	if got := names(support.Tools()); !reflect.DeepEqual(got, []string{"lookup_order"}) {
		t.Fatalf("support Tools() = %v", got)
	}
	if got := support.ToolIDs(); !reflect.DeepEqual(got, []string{"tool_from_dashboard"}) {
		t.Fatalf("support ToolIDs() = %v", got)
	}
	if id, _ := project.AgentID(support); id != "agent_support" {
		t.Fatalf("AgentID = %q", id)
	}
	if id, _ := project.ToolID(support.Tools()[0]); id != "tool_lookup" {
		t.Fatalf("ToolID = %q", id)
	}
	billing, _ := project.Agent("Billing")
	if got := names(billing.Tools()); !reflect.DeepEqual(got, []string{"issue_refund"}) {
		t.Fatalf("billing Tools() = %v", got)
	}
	if billing.Tools()[0] != project.Tools()[1] {
		t.Fatal("agents share the project's tool objects")
	}
}

// Fixture check: every file comes back byte-identical when nothing was pushed, branch configs
// included.
func TestFixtureProjectGivesBackEveryFileUnchanged(t *testing.T) {
	files := readTree(t, cliProject)
	got := filesOf(t, fromFiles(t, files))
	if !reflect.DeepEqual(got, files) {
		t.Fatalf("ToFiles() = %v\nwant %v", keysOf(got), keysOf(files))
	}
}

func TestProjectReportsBadFiles(t *testing.T) {
	index := func(config string) string { return `{"agents": [{"config": "` + config + `"}]}` }
	_, err := agentbuilder.ProjectFromFiles(map[string]string{"agents.json": index("agent_configs/a.yaml")})
	projectError(t, err, `only JSON config files are supported yet`)
	_, err = agentbuilder.ProjectFromFiles(map[string]string{"agents.json": index("agent_configs/a.toml")})
	projectError(t, err, `only JSON config files are supported yet`)
	_, err = agentbuilder.ProjectFromFiles(map[string]string{"agents.json": index("agent_configs/a.json")})
	projectError(t, err, `points at agent_configs/a.json, which doesn't exist`)
	_, err = agentbuilder.ProjectFromFiles(map[string]string{"agents.json": "{"})
	projectError(t, err, `agents.json isn't valid JSON`)
	_, err = agentbuilder.ProjectFromFiles(map[string]string{"agents.json": index("agent_configs/a.json"), "agent_configs/a.json": "{"})
	projectError(t, err, `agent_configs/a.json isn't valid JSON`)
	_, err = fromFiles(t, nil).Agent("Nope")
	projectError(t, err, `No agent named "Nope"`)
}

func TestProjectDetectsAmbiguousNames(t *testing.T) {
	project := fromFiles(t, map[string]string{
		"tools.json":          `{"tools": [{"type": "client", "config": "tool_configs/a.json"}, {"type": "client", "config": "tool_configs/b.json"}]}`,
		"tool_configs/a.json": `{"type": "client", "name": "ping", "description": "a"}`,
		"tool_configs/b.json": `{"type": "client", "name": "ping", "description": "b"}`,
	})
	_, err := project.ToolID(agenttools.Client("ping", "c"))
	projectError(t, err, `2 entries in tools.json are named "ping"`)
	// The loaded object itself is unambiguous.
	if _, err := project.ToolID(project.Tools()[0]); err != nil {
		t.Fatal(err)
	}
	f := newFakeServer(t)
	_, err = f.client().Agents.Push(ctx, agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Tools: []agenttools.Tool{agenttools.Client("ping", "c")}}), &agentbuilder.PushOptions{Project: project})
	projectError(t, err, `named "ping"`)
	if len(f.callList()) != 0 {
		t.Fatalf("calls = %v", f.callList())
	}
}

func TestOpenProjectAndPushWriteTheCLILayoutOwnerOnly(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "agent_configs"), 0o755); err != nil {
		t.Fatal(err)
	}
	// A file the project doesn't list.
	if err := os.WriteFile(filepath.Join(dir, "agent_configs", "A.json"), []byte("not ours"), 0o644); err != nil {
		t.Fatal(err)
	}
	project, err := agentbuilder.OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	f := newFakeServer(t)
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Tools: []agenttools.Tool{agenttools.Client("ping", "Ping the app")}})
	mustPush(t, f, agent, &agentbuilder.PushOptions{Project: project})

	files := readTree(t, dir)
	if got := keysOf(files); !reflect.DeepEqual(got, []string{"agent_configs/A-1.json", "agent_configs/A.json", "agents.json", "tool_configs/ping.json", "tools.json"}) {
		t.Fatalf("files = %v", got)
	}
	if files["agent_configs/A.json"] != "not ours" {
		t.Fatalf("a file the project doesn't list was overwritten: %q", files["agent_configs/A.json"])
	}
	if got := files["tool_configs/ping.json"]; got != "{\n    \"description\": \"Ping the app\",\n    \"name\": \"ping\",\n    \"type\": \"client\"\n}" {
		t.Fatalf("ping.json = %q", got)
	}
	if runtime.GOOS != "windows" {
		for _, path := range []string{"agents.json", "tools.json", "agent_configs/A-1.json", "tool_configs/ping.json"} {
			info, err := os.Stat(filepath.Join(dir, filepath.FromSlash(path)))
			if err != nil {
				t.Fatal(err)
			}
			if info.Mode().Perm() != 0o600 {
				t.Fatalf("%s mode = %v", path, info.Mode().Perm())
			}
		}
	}
	if len(project.ChangedFiles()) != 0 {
		t.Fatalf("a saved project has nothing pending: %v", project.ChangedFiles())
	}

	reopened, err := agentbuilder.OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := reopened.Agent("A")
	if err != nil {
		t.Fatal(err)
	}
	got, _ := reopened.AgentID(loaded)
	want, _ := project.AgentID(agent)
	if got != want || got == "" {
		t.Fatalf("reopened id %q, want %q", got, want)
	}
}

func TestOpenProjectOnAnEmptyDirectory(t *testing.T) {
	project, err := agentbuilder.OpenProject(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if len(project.Agents()) != 0 || len(project.Tools()) != 0 {
		t.Fatal("an empty directory is an empty project")
	}
	if paths, err := project.Save(); err != nil || len(paths) != 0 {
		t.Fatalf("Save() = %v, %v", paths, err)
	}
	if _, err := fromFiles(t, nil).Save(); err == nil {
		t.Fatal("a project from files has no directory to save to")
	}
}

func TestPushSavesTheProjectAfterAPartialFailure(t *testing.T) {
	dir := t.TempDir()
	project, err := agentbuilder.OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	f := newFakeServer(t)
	f.failOn = regexp.MustCompile(`^POST /v1/convai/agents/create$`)
	_, err = f.client().Agents.Push(ctx, supportAgent(), &agentbuilder.PushOptions{Project: project})
	var perr *agentbuilder.PushError
	if !errors.As(err, &perr) || len(perr.Result.Tools) != 2 {
		t.Fatalf("err = %v", err)
	}
	// The tools it created are on disk, so the next push updates them instead of making copies.
	files := readTree(t, dir)
	if got := keysOf(files); !reflect.DeepEqual(got, []string{"tool_configs/get_order.json", "tool_configs/show_order.json", "tools.json"}) {
		t.Fatalf("files = %v", got)
	}
	f.failOn = nil
	reopened, err := agentbuilder.OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	f.resetCalls()
	result := mustPush(t, f, supportAgent(), &agentbuilder.PushOptions{Project: reopened})
	if got := f.writes(); !reflect.DeepEqual(got, []string{"POST /v1/convai/agents/create"}) {
		t.Fatalf("writes = %v", got)
	}
	if result.Tools["get_order"].Action != agentbuilder.ToolUnchanged {
		t.Fatalf("get_order = %+v", result.Tools["get_order"])
	}
}

func TestProjectRefusesPathsOutsideIt(t *testing.T) {
	outside := t.TempDir()
	if err := os.WriteFile(filepath.Join(outside, "a.json"), []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}

	escaping := t.TempDir()
	if err := os.WriteFile(filepath.Join(escaping, "agents.json"), []byte(`{"agents": [{"config": "../a.json"}]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := agentbuilder.OpenProject(escaping)
	projectError(t, err, `outside the project`)

	if runtime.GOOS == "windows" {
		t.Skip("symlinks need privileges on Windows")
	}
	linked := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(linked, "agent_configs")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(linked, "agents.json"), []byte(`{"agents": [{"config": "agent_configs/a.json"}]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err = agentbuilder.OpenProject(linked)
	projectError(t, err, `through a symlink`)
}

func TestProjectNeverWritesThroughASymlink(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlinks need privileges on Windows")
	}
	dir := t.TempDir()
	target := filepath.Join(dir, "elsewhere.json")
	if err := os.WriteFile(target, []byte(`{"agents": []}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, filepath.Join(dir, "agents.json")); err != nil {
		t.Fatal(err)
	}
	project, err := agentbuilder.OpenProject(dir)
	if err != nil {
		t.Fatal(err)
	}
	f := newFakeServer(t)
	_, err = f.client().Agents.Push(ctx, agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p")}), &agentbuilder.PushOptions{Project: project})
	projectError(t, err, `Refusing to write agents.json: it's a symlink`)
	if data, _ := os.ReadFile(target); string(data) != `{"agents": []}` {
		t.Fatalf("the symlink's target was written: %q", data)
	}
}

func TestBranchWithoutAFileIsRejectedBeforeAnyRequest(t *testing.T) {
	project := fromFiles(t, readTree(t, cliProject))
	f := newFakeServer(t)
	billing, _ := project.Agent("Billing")
	_, err := f.client().Agents.Push(ctx, billing, &agentbuilder.PushOptions{Project: project, BranchID: "agtbrch_x"})
	projectError(t, err, `has no config file`)
	_, err = f.client().Agents.Pull(ctx, "agent_support", &agentbuilder.PullOptions{Project: project, BranchID: "agtbrch_x"})
	projectError(t, err, `has no config file`)
	if len(f.callList()) != 0 {
		t.Fatalf("calls = %v", f.callList())
	}
}

func TestRecordingABranchKeepsTheMainAgent(t *testing.T) {
	project := fromFiles(t, readTree(t, cliProject))
	main, _ := project.Agent("Support")
	f := newFakeServer(t)
	// The server has the branch; the project's main entry is untouched.
	beta := toTree(t, readFixture(t, "agent_configs/support.beta.json"))
	f.addAgent("agent_support", readFixture(t, "agent_configs/support.json"))
	f.mu.Lock()
	f.branches["agent_support/agtbrch_beta"] = beta
	f.mu.Unlock()
	f.addTool("tool_lookup", readFixture(t, "tool_configs/lookup_order.json"))
	f.addTool("tool_refund", readFixture(t, "tool_configs/issue_refund.json"))

	pulled, err := f.client().Agents.Pull(ctx, "agent_support", &agentbuilder.PullOptions{Project: project, BranchID: "agtbrch_beta"})
	if err != nil {
		t.Fatal(err)
	}
	if again, _ := project.Agent("Support"); again != main || pulled == main {
		t.Fatal("recording a branch must keep the main branch's agent")
	}
	changed := project.ChangedFiles()
	if !containsString(changed, "agent_configs/support.beta.json") || containsString(changed, "agent_configs/support.json") {
		t.Fatalf("changed = %v", changed)
	}
	// The version is recorded on the branch's entry.
	agents := toTree(t, filesOf(t, project)["agents.json"])["agents"].([]any)
	if got := at(agents[0], "branches", "beta", "version_id"); got != "ver_1" {
		t.Fatalf("branch version_id = %v", got)
	}
	if got := at(agents[0], "version_id"); got != nil {
		t.Fatalf("main version_id = %v", got)
	}
}

func TestRecordingAToolKeepsItsTypeInStep(t *testing.T) {
	project := fromFiles(t, readTree(t, cliProject))
	f := newFakeServer(t)
	// The server's lookup_order is now a client tool; pulling records it by id.
	f.addTool("tool_lookup", `{"type": "client", "name": "lookup_order", "description": "Now a client tool"}`)
	f.addTool("tool_from_dashboard", `{"type": "mcp", "name": "dashboard"}`)
	f.addTool("tool_refund", readFixture(t, "tool_configs/issue_refund.json"))
	f.addAgent("agent_support", readFixture(t, "agent_configs/support.json"))
	if _, err := f.client().Agents.Pull(ctx, "agent_support", &agentbuilder.PullOptions{Project: project}); err != nil {
		t.Fatal(err)
	}
	tools := toTree(t, filesOf(t, project)["tools.json"])["tools"].([]any)
	if len(tools) != 2 || at(tools[0], "type") != "client" || at(tools[0], "id") != "tool_lookup" {
		t.Fatalf("tools.json = %v", tools)
	}
}

func TestAPendingWriteIsDroppedOnceAFileIsBack(t *testing.T) {
	loaded := "{\n    \"conversation_config\": {},\n    \"name\": \"A\",\n    \"platform_settings\": {},\n    \"tags\": []\n}"
	files := map[string]string{
		"agents.json":          "{\n    \"agents\": [\n        {\n            \"config\": \"agent_configs/a.json\",\n            \"id\": \"agent_1\"\n        }\n    ]\n}",
		"agent_configs/a.json": loaded,
	}
	project := fromFiles(t, files)
	agent, _ := project.Agent("A")
	f := newFakeServer(t)
	f.addAgent("agent_1", `{}`)

	agent.Update(&agentbuilder.Options{Tags: []string{"x"}})
	if _, err := f.client().Agents.Push(ctx, agent, &agentbuilder.PushOptions{Project: project}); err != nil {
		t.Fatal(err)
	}
	if got := project.ChangedFiles(); !containsString(got, "agent_configs/a.json") {
		t.Fatalf("changed = %v", got)
	}
	agent.Update(&agentbuilder.Options{Tags: []string{}})
	if _, err := f.client().Agents.Push(ctx, agent, &agentbuilder.PushOptions{Project: project}); err != nil {
		t.Fatal(err)
	}
	// Only agents.json, for the version the update reported.
	if got := project.ChangedFiles(); !reflect.DeepEqual(got, []string{"agents.json"}) {
		t.Fatalf("changed = %v", got)
	}
	if got := filesOf(t, project)["agent_configs/a.json"]; got != loaded {
		t.Fatalf("a.json = %q", got)
	}
	if got := filesOf(t, project)["tools.json"]; got != "{\n    \"tools\": []\n}" {
		t.Fatalf("tools.json = %q", got)
	}
}

func containsString(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
