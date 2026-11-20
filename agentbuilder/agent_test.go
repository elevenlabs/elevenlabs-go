package agentbuilder_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agentbuilder"
	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// cliProject is the shared cross-SDK project fixture, vendored from elevenlabs-js
// (docs/agents-as-code/fixtures/project/cli-graph).
var cliProject = filepath.Join("testdata", "project", "cli-graph")

func ptr[T any](v T) *T { return &v }

func readFixture(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(cliProject, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func fixtureTool(t *testing.T, path string) agenttools.Tool {
	t.Helper()
	tool, err := agenttools.FromJSON(readFixture(t, path))
	if err != nil {
		t.Fatal(err)
	}
	return tool
}

// sorted is JSON text in canonical form, keys sorted at every level, for comparing configs.
func sorted(t *testing.T, value any) string {
	t.Helper()
	if text, ok := value.(string); ok {
		tree, err := wirejson.Decode([]byte(text))
		if err != nil {
			t.Fatal(err)
		}
		value = tree
	}
	out, err := wirejson.String(value)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func getOrder(description string) *agenttools.WebhookTool {
	return agenttools.Webhook("get_order", description).
		Get("https://api.example.com/orders/{order_id}").
		PathParam("order_id", agenttools.String().Describe("The order number"))
}

func showOrder() *agenttools.ClientTool {
	return agenttools.Client("show_order", "Show the order card").
		Param("order_id", agenttools.String().FromVariable("order_id"))
}

type issue struct{ code, path string }

func issuesOf(t *testing.T, err error) []issue {
	t.Helper()
	var verr *agentbuilder.ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected a *ValidationError, got %v", err)
	}
	out := make([]issue, len(verr.Issues))
	for i, is := range verr.Issues {
		out[i] = issue{string(is.Code), is.Path}
	}
	return out
}

// at reads a path of a wire tree.
func at(value any, path ...string) any {
	for _, key := range path {
		obj, _ := value.(map[string]any)
		value = obj[key]
	}
	return value
}

func TestTypedOptionsToolsAndSystemTools(t *testing.T) {
	agent := agentbuilder.New("Support line", &agentbuilder.Options{
		Prompt:       ptr("You are a support agent."),
		FirstMessage: ptr("Hi!"),
		Language:     ptr("en"),
		LLM:          types.LlmGemini20Flash.Ptr(),
		Temperature:  ptr(0.5),
		VoiceID:      ptr("voice_1"),
		TTSModel:     types.TtsConversationalModelElevenFlashV25.Ptr(),
		Tags:         []string{"support"},
		Tools:        []agenttools.Tool{agenttools.SystemEndCall(), getOrder("Look up an order by its number")},
		ToolIDs:      []string{"tool_existing"},
	})
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}})
	if err != nil {
		t.Fatal(err)
	}
	want := `{
  "conversation_config": {
    "agent": {
      "first_message": "Hi!",
      "language": "en",
      "prompt": {
        "built_in_tools": {
          "end_call": {
            "name": "end_call",
            "params": {
              "system_tool_type": "end_call"
            },
            "type": "system"
          }
        },
        "llm": "gemini-2.0-flash",
        "prompt": "You are a support agent.",
        "temperature": 0.5,
        "tool_ids": [
          "tool_existing",
          "tool_1"
        ]
      }
    },
    "tts": {
      "model_id": "eleven_flash_v2_5",
      "voice_id": "voice_1"
    }
  },
  "name": "Support line",
  "platform_settings": {},
  "tags": [
    "support"
  ]
}
`
	if got := sorted(t, config); got != want {
		t.Fatalf("ToJSON:\n%s", got)
	}

	body, err := agent.Build(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}})
	if err != nil {
		t.Fatal(err)
	}
	if body.GetName() == nil || *body.Name != "Support line" {
		t.Fatalf("Build name = %v", body.Name)
	}
	prompt := body.ConversationConfig.GetAgent().GetPrompt()
	if !reflect.DeepEqual(prompt.GetToolIds(), []string{"tool_existing", "tool_1"}) {
		t.Fatalf("Build tool_ids = %v", prompt.GetToolIds())
	}
	if prompt.GetBuiltInTools().GetEndCall() == nil {
		t.Fatal("Build has no end_call built-in tool")
	}
}

func TestRawConfigMergesUnderTypedOptions(t *testing.T) {
	agent := agentbuilder.New("A", &agentbuilder.Options{
		Prompt: ptr("p"),
		ConversationConfig: &types.ConversationalConfig{
			Turn:  &types.TurnConfig{TurnTimeout: ptr(7.0)},
			Agent: &types.AgentConfig{DisableFirstMessageInterruptions: ptr(true)},
		},
		PlatformSettings: &types.AgentPlatformSettingsRequestModel{Privacy: &types.PrivacyConfigInput{RecordVoice: ptr(false)}},
	})
	agent.Update(&agentbuilder.Options{ConversationConfig: &types.ConversationalConfig{Turn: &types.TurnConfig{SilenceEndCallTimeout: ptr(20.0)}}})
	config, err := agent.ToJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := sorted(t, config["conversation_config"]); got != sorted(t, `{
		"agent": {"disable_first_message_interruptions": true, "prompt": {"prompt": "p"}},
		"turn": {"silence_end_call_timeout": 20, "turn_timeout": 7}
	}`) {
		t.Fatalf("conversation_config = %s", got)
	}
	if got := sorted(t, config["platform_settings"]); got != sorted(t, `{"privacy": {"record_voice": false}}`) {
		t.Fatalf("platform_settings = %s", got)
	}
}

func TestWorkflowToolsGetTheirIDs(t *testing.T) {
	lookup, show := getOrder("Look up an order by its number"), showOrder()
	override := &types.WorkflowToolLocatorSchemaOverridesValue{}
	if err := json.Unmarshal([]byte(`{"source": "constant", "constant_value": "1"}`), override); err != nil {
		t.Fatal(err)
	}
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
		Tool("lookup", nil, &agentworkflows.Options{
			Tools:           []agentworkflows.Tool{lookup},
			SchemaOverrides: map[string]map[string]*types.WorkflowToolLocatorSchemaOverridesValue{"get_order": {"path_params.order_id": override}},
		}).
		OnSuccess(nil).
		Converse("resolve", "Explain the order", &agentworkflows.Options{Tools: []agentworkflows.Tool{show}, ToolIDs: []string{"tool_raw"}})
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Workflow: wf})

	if got := agent.Tools(); !reflect.DeepEqual(got, []agenttools.Tool{lookup, show}) {
		t.Fatalf("Tools() = %v", got)
	}
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_g", "show_order": "tool_s"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := sorted(t, at(config, "workflow", "nodes", "lookup", "tools")); got != sorted(t, `[{"schema_overrides": {"path_params.order_id": {"constant_value": "1", "source": "constant"}}, "tool_id": "tool_g"}]`) {
		t.Fatalf("lookup tools = %s", got)
	}
	if got := at(config, "workflow", "nodes", "resolve", "additional_tool_ids"); !reflect.DeepEqual(got, []any{"tool_raw", "tool_s"}) {
		t.Fatalf("resolve additional_tool_ids = %v", got)
	}
	_, err = agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_g"}})
	if got := issuesOf(t, err); !reflect.DeepEqual(got, []issue{{"unresolved_tool", "tools.show_order"}}) {
		t.Fatalf("issues = %v", got)
	}
}

func TestValidateNeedsNoIDs(t *testing.T) {
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p"), Tools: []agenttools.Tool{getOrder("d")}})
	if err := agent.Validate(nil); err != nil {
		t.Fatal(err)
	}
	if _, err := agent.ToJSON(nil); err == nil {
		t.Fatal("ToJSON without ids must report unresolved_tool")
	}
}

func TestEveryProblemAtOnce(t *testing.T) {
	wf := agentworkflows.New(nil)
	wf.Tool("hang_up", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{agenttools.SystemEndCall()}})
	agent := agentbuilder.New("", &agentbuilder.Options{
		Prompt:             ptr("p"),
		Tools:              []agenttools.Tool{getOrder("d"), getOrder("d"), agenttools.Webhook("broken", "No URL")},
		Workflow:           wf,
		ConversationConfig: &types.ConversationalConfig{Agent: &types.AgentConfig{Prompt: &types.PromptAgentApiModelInput{Prompt: ptr("raw prompt")}}},
	})
	got := issuesOf(t, agent.Validate(nil))
	want := []issue{
		{"missing_name", ""},
		{"option_conflict", "conversation_config.agent.prompt.prompt"},
		{"duplicate_tool_name", "tools.get_order"},
		{"invalid_tool", "tools.broken"},
		{"system_tool_in_workflow", "workflow.tools.end_call"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("issues = %v", got)
	}
}

func TestMissingPromptIsAWarning(t *testing.T) {
	var warnings []string
	err := agentbuilder.New("A", nil).Validate(&agentbuilder.ValidateOptions{OnWarning: func(issue agentbuilder.Issue) {
		warnings = append(warnings, string(issue.Code))
	}})
	if err != nil || !reflect.DeepEqual(warnings, []string{"missing_prompt"}) {
		t.Fatalf("err = %v, warnings = %v", err, warnings)
	}
}

func TestAddToolAttachesTheTool(t *testing.T) {
	agent := agentbuilder.New("A", &agentbuilder.Options{Prompt: ptr("p")}).AddTool(getOrder("d"))
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_1"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := at(config, "conversation_config", "agent", "prompt", "tool_ids"); !reflect.DeepEqual(got, []any{"tool_1"}) {
		t.Fatalf("tool_ids = %v", got)
	}
}

// Fixture check: agent_configs/support.json round-trips through FromJSON and ToJSON to its sorted
// form, keeping fields the SDK doesn't model and the tool order.
func TestFixtureSupportConfigRoundTrips(t *testing.T) {
	raw := readFixture(t, "agent_configs/support.json")
	lookup := fixtureTool(t, "tool_configs/lookup_order.json")
	agent, err := agentbuilder.FromJSON(raw, &agentbuilder.LoadOptions{ToolsByID: map[string]agenttools.Tool{"tool_lookup": lookup}})
	if err != nil {
		t.Fatal(err)
	}
	if agent.Name() != "Support" {
		t.Fatalf("Name() = %q", agent.Name())
	}
	if got := agent.Tools(); !reflect.DeepEqual(got, []agenttools.Tool{lookup}) {
		t.Fatalf("Tools() = %v", got)
	}
	if got := agent.ToolIDs(); !reflect.DeepEqual(got, []string{"tool_from_dashboard"}) {
		t.Fatalf("ToolIDs() = %v", got)
	}
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"lookup_order": "tool_lookup"}})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := sorted(t, config), sorted(t, raw); got != want {
		t.Fatalf("round trip:\n%s\nwant:\n%s", got, want)
	}
}

func TestLoadedWorkflowBecomesAWorkflowOnFirstRead(t *testing.T) {
	raw := readFixture(t, "agent_configs/support.beta.json")
	refund := fixtureTool(t, "tool_configs/issue_refund.json")
	agent, err := agentbuilder.FromJSON(raw, &agentbuilder.LoadOptions{ToolsByID: map[string]agenttools.Tool{"tool_refund": refund}})
	if err != nil {
		t.Fatal(err)
	}
	// Read from the raw workflow, before it's materialized.
	if got := agent.Tools(); !reflect.DeepEqual(got, []agenttools.Tool{refund}) {
		t.Fatalf("Tools() = %v", got)
	}
	unedited, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"issue_refund": "tool_refund"}})
	if err != nil {
		t.Fatal(err)
	}
	var source map[string]json.RawMessage
	if err := json.Unmarshal([]byte(raw), &source); err != nil {
		t.Fatal(err)
	}
	if got, want := sorted(t, unedited["workflow"]), sorted(t, string(source["workflow"])); got != want {
		t.Fatalf("unedited workflow:\n%s\nwant:\n%s", got, want)
	}

	agent.Workflow().Node("refunds").Update(&agentworkflows.Options{Prompt: ptr("Handle refunds")})
	edited, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"issue_refund": "tool_new"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := at(edited, "workflow", "nodes", "refunds", "additional_tool_ids"); !reflect.DeepEqual(got, []any{"tool_new"}) {
		t.Fatalf("additional_tool_ids = %v", got)
	}
	if got := at(edited, "workflow", "nodes", "refunds", "additional_prompt"); got != "Handle refunds" {
		t.Fatalf("additional_prompt = %v", got)
	}
}

func TestReplacingToolsReplacesLoadedSystemTools(t *testing.T) {
	agent, err := agentbuilder.FromJSON(readFixture(t, "agent_configs/support.json"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(agent.Tools()) != 0 {
		t.Fatalf("system tools aren't managed by push: %v", agent.Tools())
	}
	agent.Update(&agentbuilder.Options{Tools: []agenttools.Tool{}})
	config, err := agent.ToJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	prompt := at(config, "conversation_config", "agent", "prompt")
	if at(prompt, "built_in_tools") != nil {
		t.Fatalf("built_in_tools = %v", at(prompt, "built_in_tools"))
	}
	// Ids it doesn't manage stay.
	if got := at(prompt, "tool_ids"); !reflect.DeepEqual(got, []any{"tool_lookup", "tool_from_dashboard"}) {
		t.Fatalf("tool_ids = %v", got)
	}
}

func TestEditedLoadedSystemToolIsRebuilt(t *testing.T) {
	agent, err := agentbuilder.FromJSON(readFixture(t, "agent_configs/support.json"), nil)
	if err != nil {
		t.Fatal(err)
	}
	agent.Update(&agentbuilder.Options{Tools: []agenttools.Tool{agenttools.SystemEndCall()}})
	config, err := agent.ToJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	builtIn := at(config, "conversation_config", "agent", "prompt", "built_in_tools").(map[string]any)
	if _, ok := builtIn["end_call"]; !ok || len(builtIn) != 1 {
		t.Fatalf("built_in_tools = %v", builtIn)
	}
}

func TestEmptyToolListCountsAsSet(t *testing.T) {
	agent := agentbuilder.New("A", &agentbuilder.Options{
		Prompt:             ptr("p"),
		ToolIDs:            []string{},
		ConversationConfig: &types.ConversationalConfig{Agent: &types.AgentConfig{Prompt: &types.PromptAgentApiModelInput{ToolIds: []string{"tool_existing"}}}},
	})
	if got := issuesOf(t, agent.Validate(nil)); !reflect.DeepEqual(got, []issue{{"option_conflict", "conversation_config.agent.prompt.tool_ids"}}) {
		t.Fatalf("issues = %v", got)
	}
}

func TestLoadedWorkflowIsValidatedWithoutBeingMaterialized(t *testing.T) {
	raw := readFixture(t, "agent_configs/support.beta.json")
	broken := strings.Replace(raw, `"e1": {`, `"broken": { "source": "start_node", "target": "nowhere" }, "e1": {`, 1)
	agent, err := agentbuilder.FromJSON(broken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := issuesOf(t, agent.Validate(nil)); !reflect.DeepEqual(got, []issue{{"invalid_workflow", "workflow"}}) {
		t.Fatalf("issues = %v", got)
	}
}

func TestNewToolIDsReachAnUnreadLoadedWorkflow(t *testing.T) {
	refund := fixtureTool(t, "tool_configs/issue_refund.json")
	agent, err := agentbuilder.FromJSON(readFixture(t, "agent_configs/support.beta.json"), &agentbuilder.LoadOptions{ToolsByID: map[string]agenttools.Tool{"tool_refund": refund}})
	if err != nil {
		t.Fatal(err)
	}
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"issue_refund": "tool_clone"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := at(config, "workflow", "nodes", "refunds", "additional_tool_ids"); !reflect.DeepEqual(got, []any{"tool_clone"}) {
		t.Fatalf("additional_tool_ids = %v", got)
	}
}

func TestUpdateRenames(t *testing.T) {
	agent := agentbuilder.New("A", nil).Update(&agentbuilder.Options{Name: ptr("B")})
	if agent.Name() != "B" {
		t.Fatalf("Name() = %q", agent.Name())
	}
}
