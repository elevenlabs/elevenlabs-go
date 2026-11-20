package agentbuilder

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Options are the settings of New and Update. Every field is optional: nil means "not set", and a
// non-nil empty slice sets an empty list.
type Options struct {
	// Name renames the agent (Update; New takes the name as its own argument).
	Name *string

	// Prompt is the agent's instructions: conversation_config.agent.prompt.prompt.
	Prompt *string
	// FirstMessage is what the agent says first: conversation_config.agent.first_message.
	FirstMessage *string
	// Language is conversation_config.agent.language, such as "en".
	Language *string
	// LLM is conversation_config.agent.prompt.llm.
	LLM *types.Llm
	// Temperature is conversation_config.agent.prompt.temperature.
	Temperature *float64
	// VoiceID is conversation_config.tts.voice_id.
	VoiceID *string
	// TTSModel is conversation_config.tts.model_id.
	TTSModel *types.TtsConversationalModel
	// Tags are the agent's tags.
	Tags []string

	// Tools are the tools the agent can use anywhere. Webhook and client tools are created or
	// updated by Push and attached by id in prompt.tool_ids; system tools go into
	// prompt.built_in_tools. Replaces the agent's tools on Update.
	Tools []agenttools.Tool
	// ToolIDs are the ids of existing tools to attach as they are, before the managed tools. Push
	// doesn't create or change them. Replaces the agent's tool ids on Update.
	ToolIDs []string
	// Workflow is the agent's workflow. Its tool nodes and subagents can take tools written as
	// code (agentworkflows.Options.Tools), which Push manages too.
	Workflow *agentworkflows.Workflow

	// ConversationConfig holds any other conversation settings, merged under the typed options
	// above. Setting a field here that a typed option also sets is reported as option_conflict.
	// Each Update merges into what earlier ones set.
	ConversationConfig *types.ConversationalConfig
	// PlatformSettings holds platform settings (privacy, evaluation, widget...), merged the same
	// way.
	PlatformSettings *types.AgentPlatformSettingsRequestModel
}

// LoadOptions are the optional settings of FromJSON.
type LoadOptions struct {
	// ToolsByID are the tools the config refers to, by id. Each matching id becomes a tool the
	// agent manages; the others stay as ToolIDs.
	ToolsByID map[string]agenttools.Tool
}

// ValidateOptions are the optional settings of Validate.
type ValidateOptions struct {
	// OnWarning is called once per warning, such as missing_prompt.
	OnWarning func(Issue)
}

// BuildOptions are the optional settings of ToJSON and Build.
type BuildOptions struct {
	// ToolIDs maps the name of each webhook and client tool the agent and its workflow use to its
	// id. Push passes these after creating or updating the tools. Every such tool needs one; a
	// missing one is reported as unresolved_tool.
	ToolIDs map[string]string
	// OnWarning is called once per warning, such as missing_prompt.
	OnWarning func(Issue)
}

// typedOption is one typed option: its name, as conflicts report it, and the piece of the
// generated ConversationalConfig it sets. The pieces are written against the generated types, so
// a wrong or renamed field fails to compile; the API-format path is read off the marshaled piece.
type typedOption struct {
	name  string
	piece func(o *Options) *types.ConversationalConfig
}

var typedOptions = []typedOption{
	{"Prompt", func(o *Options) *types.ConversationalConfig {
		if o.Prompt == nil {
			return nil
		}
		return &types.ConversationalConfig{Agent: &types.AgentConfig{Prompt: &types.PromptAgentApiModelInput{Prompt: o.Prompt}}}
	}},
	{"FirstMessage", func(o *Options) *types.ConversationalConfig {
		if o.FirstMessage == nil {
			return nil
		}
		return &types.ConversationalConfig{Agent: &types.AgentConfig{FirstMessage: o.FirstMessage}}
	}},
	{"Language", func(o *Options) *types.ConversationalConfig {
		if o.Language == nil {
			return nil
		}
		return &types.ConversationalConfig{Agent: &types.AgentConfig{Language: o.Language}}
	}},
	{"LLM", func(o *Options) *types.ConversationalConfig {
		if o.LLM == nil {
			return nil
		}
		return &types.ConversationalConfig{Agent: &types.AgentConfig{Prompt: &types.PromptAgentApiModelInput{Llm: o.LLM}}}
	}},
	{"Temperature", func(o *Options) *types.ConversationalConfig {
		if o.Temperature == nil {
			return nil
		}
		return &types.ConversationalConfig{Agent: &types.AgentConfig{Prompt: &types.PromptAgentApiModelInput{Temperature: o.Temperature}}}
	}},
	{"VoiceID", func(o *Options) *types.ConversationalConfig {
		if o.VoiceID == nil {
			return nil
		}
		return &types.ConversationalConfig{Tts: &types.TtsConversationalConfigInput{VoiceId: o.VoiceID}}
	}},
	{"TTSModel", func(o *Options) *types.ConversationalConfig {
		if o.TTSModel == nil {
			return nil
		}
		return &types.ConversationalConfig{Tts: &types.TtsConversationalConfigInput{ModelId: o.TTSModel}}
	}},
}

// typedWire returns the API-format path a typed option's piece sets, and its value there.
func typedWire(piece *types.ConversationalConfig) ([]string, any, error) {
	wire, err := wirejson.Normalize(piece)
	if err != nil {
		return nil, nil, err
	}
	path := []string{"conversation_config"}
	for {
		obj, ok := wire.(map[string]any)
		if !ok || len(obj) != 1 {
			break
		}
		for key, value := range obj {
			path = append(path, key)
			wire = value
		}
	}
	return path, wire, nil
}

var promptPath = func() []string {
	prompt := ""
	path, _, err := typedWire(typedOptions[0].piece(&Options{Prompt: &prompt}))
	if err != nil {
		panic(err)
	}
	return path
}()

// loadedSystemTool is a system tool loaded from a config: its config as loaded, sent as it is
// while the tool is unchanged, and its ToJSONString then.
type loadedSystemTool struct {
	raw      any
	baseline string
}

// Agent is an agent written as code: its settings, its tools and its workflow in one definition.
// Push it with client.Agents.Push, which creates or updates the tools and the agent together.
//
// Chaining never fails: every problem is reported at once by Validate, ToJSON and Build as a
// *ValidationError. An Agent holds no client, so it can be validated without an API key and
// pushed to several workspaces. It is not safe for concurrent use.
type Agent struct {
	name  string
	typed Options
	tags  []any
	// hasTags is whether tags were set or loaded.
	hasTags  bool
	tools    []agenttools.Tool
	toolIDs  []string
	workflow *agentworkflows.Workflow
	// rawWorkflow is a loaded workflow nobody has read yet, kept in API format so it round-trips
	// exactly; rawWorkflowJSON is its text, and rawNodeOrder its node ids in document order.
	rawWorkflow     map[string]any
	rawWorkflowJSON []byte
	rawNodeOrder    []string
	toolsByID       map[string]agenttools.Tool
	// loadedToolIDs is the loaded tool_ids, so a loaded agent keeps its tool order.
	loadedToolIDs []string
	// toolListGiven is whether Tools or ToolIDs was given (even empty), or tool_ids loaded, so the
	// agent sets tool_ids rather than leaving it alone.
	toolListGiven bool
	loadedSystem  map[*agenttools.SystemTool]loadedSystemTool
	// base holds the settings loaded from a config, in API format. Raw and typed options are
	// applied on top.
	base                  map[string]any
	rawConversationConfig []*types.ConversationalConfig
	rawPlatformSettings   []*types.AgentPlatformSettingsRequestModel
}

// New returns an agent named name. opts may be nil.
//
//	agent := agentbuilder.New("Support line", &agentbuilder.Options{
//		Prompt:       ptr("You are a support agent for Acme."),
//		FirstMessage: ptr("Hi, how can I help?"),
//		Tools:        []agenttools.Tool{agenttools.SystemEndCall(), getOrder},
//		Workflow:     wf,
//	})
func New(name string, opts *Options) *Agent {
	a := &Agent{
		name:         name,
		toolsByID:    map[string]agenttools.Tool{},
		loadedSystem: map[*agenttools.SystemTool]loadedSystemTool{},
		base:         map[string]any{},
	}
	return a.Update(opts)
}

// FromJSON loads an agent config in API format (snake_case keys): a CLI agent_configs/*.json
// file, or an agent's ToJSON. data is JSON text (a string, []byte or json.RawMessage) or a parsed
// value. Tool ids found in opts.ToolsByID become tools the agent manages; the others stay as
// ToolIDs. Loaded system tools become *agenttools.SystemTool values. Fields this SDK doesn't
// model are kept, so an unedited agent round-trips exactly. opts may be nil.
//
// Document order decides the order of a loaded workflow's nodes, so pass the JSON text where you
// have it; a parsed Go map is read in sorted key order.
func FromJSON(data any, opts *LoadOptions) (*Agent, error) {
	text, err := rawJSON(data)
	if err != nil {
		return nil, err
	}
	parsed, err := decodeObject(text, "the agent config")
	if err != nil {
		return nil, err
	}
	a := New("", nil)
	if opts != nil {
		for id, tool := range opts.ToolsByID {
			if tool != nil {
				a.toolsByID[id] = tool
			}
		}
	}
	a.name, _ = parsed["name"].(string)
	if tags, ok := parsed["tags"].([]any); ok {
		a.tags, a.hasTags = tags, true
	}
	if workflow, ok := parsed["workflow"].(map[string]any); ok {
		workflowJSON := field(text, "workflow")
		if _, err := agentworkflows.FromJSON(workflowJSON); err != nil {
			return nil, err
		}
		a.rawWorkflow, a.rawWorkflowJSON, a.rawNodeOrder = workflow, workflowJSON, nodeOrder(workflowJSON)
	}
	delete(parsed, "name")
	delete(parsed, "tags")
	delete(parsed, "workflow")
	a.base = parsed

	if ids, ok := getPath(parsed, toolIDsPath); ok {
		if list, isList := ids.([]any); isList {
			deletePath(parsed, toolIDsPath)
			a.loadedToolIDs = stringList(list)
			for _, id := range a.loadedToolIDs {
				if tool, managed := a.toolsByID[id]; managed {
					a.tools = append(a.tools, tool)
				} else {
					a.toolIDs = append(a.toolIDs, id)
				}
			}
			a.toolListGiven = true
		}
	}

	// System tools become SystemTool values, so replacing Tools replaces them like on an agent
	// built in code.
	if value, ok := getPath(parsed, builtInToolsPath); ok {
		if builtIn, isObj := value.(map[string]any); isObj {
			keys := make([]string, 0, len(builtIn))
			for key := range builtIn {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			for _, key := range keys {
				raw, isObj := builtIn[key].(map[string]any)
				if !isObj {
					continue
				}
				input := map[string]any{"type": "system"}
				for k, v := range raw {
					input[k] = v
				}
				tool, err := agenttools.FromJSON(input)
				if err != nil || tool.Kind() != agenttools.ToolKindSystem || tool.AsSystem().SystemToolType() != key {
					continue // kept in the config as it is
				}
				a.tools = append(a.tools, tool)
				a.loadedSystem[tool.AsSystem()] = loadedSystemTool{raw: wirejson.Clone(raw), baseline: toolText(tool)}
				delete(builtIn, key)
			}
			if len(builtIn) == 0 {
				deletePath(parsed, builtInToolsPath)
			}
		}
	}
	return a, nil
}

// Name is the agent's name.
func (a *Agent) Name() string { return a.name }

// Workflow is the agent's workflow, or nil if it has none. A loaded workflow is turned into an
// *agentworkflows.Workflow the first time it's read here, with its tool ids swapped for the tools
// the agent knows (see agentworkflows.FromJSONWithTools), so it can be edited. Until then it is
// sent exactly as loaded.
func (a *Agent) Workflow() *agentworkflows.Workflow {
	if a.workflow == nil && a.rawWorkflow != nil {
		byID := make(map[string]agentworkflows.Tool, len(a.toolsByID))
		for id, tool := range a.toolsByID {
			byID[id] = tool
		}
		// FromJSON checked that the workflow loads.
		if wf, err := agentworkflows.FromJSONWithTools(a.rawWorkflowJSON, byID); err == nil {
			a.workflow = wf
			a.rawWorkflow, a.rawWorkflowJSON, a.rawNodeOrder = nil, nil, nil
		}
	}
	return a.workflow
}

// Update changes the agent and returns it. Typed options, Name, Tags, Tools, ToolIDs and Workflow
// replace what was there; ConversationConfig and PlatformSettings merge. opts may be nil.
func (a *Agent) Update(opts *Options) *Agent {
	if opts == nil {
		return a
	}
	if opts.Name != nil {
		a.name = *opts.Name
	}
	if opts.Prompt != nil {
		a.typed.Prompt = opts.Prompt
	}
	if opts.FirstMessage != nil {
		a.typed.FirstMessage = opts.FirstMessage
	}
	if opts.Language != nil {
		a.typed.Language = opts.Language
	}
	if opts.LLM != nil {
		a.typed.LLM = opts.LLM
	}
	if opts.Temperature != nil {
		a.typed.Temperature = opts.Temperature
	}
	if opts.VoiceID != nil {
		a.typed.VoiceID = opts.VoiceID
	}
	if opts.TTSModel != nil {
		a.typed.TTSModel = opts.TTSModel
	}
	if opts.Tags != nil {
		a.tags, a.hasTags = anyList(opts.Tags), true
	}
	if opts.Tools != nil {
		a.tools = append([]agenttools.Tool{}, opts.Tools...)
		a.toolListGiven = true
	}
	if opts.ToolIDs != nil {
		a.toolIDs = append([]string{}, opts.ToolIDs...)
		a.toolListGiven = true
	}
	if opts.Workflow != nil {
		a.workflow = opts.Workflow
		a.rawWorkflow, a.rawWorkflowJSON, a.rawNodeOrder = nil, nil, nil
	}
	if opts.ConversationConfig != nil {
		a.rawConversationConfig = append(a.rawConversationConfig, opts.ConversationConfig)
	}
	if opts.PlatformSettings != nil {
		a.rawPlatformSettings = append(a.rawPlatformSettings, opts.PlatformSettings)
	}
	return a
}

// AddTool adds a tool the agent can use anywhere, and returns the agent.
func (a *Agent) AddTool(tool agenttools.Tool) *Agent {
	a.tools = append(a.tools, tool)
	a.toolListGiven = true
	return a
}

// Tools lists every webhook and client tool the agent and its workflow use, each once, in the
// order they were first used: the tools Push creates or updates. System tools aren't listed;
// they live inside the agent's config. Nor are tools of other kinds (such as mcp): they stay
// attached by the id the agent was loaded with, as Pull leaves them.
func (a *Agent) Tools() []agenttools.Tool {
	var out []agenttools.Tool
	for _, tool := range append(append([]agenttools.Tool{}, a.tools...), a.workflowTools()...) {
		if tool == nil || !managedTool(tool) || containsTool(out, tool) {
			continue
		}
		out = append(out, tool)
	}
	return out
}

// managedTool reports whether Push creates or updates tool: a webhook or client tool.
func managedTool(tool agenttools.Tool) bool {
	kind := tool.Kind()
	return kind == agenttools.ToolKindWebhook || kind == agenttools.ToolKindClient
}

// knownID is the id the agent was loaded with for tool (LoadOptions.ToolsByID; a project passes
// its tools.json ids this way). Tools of kinds Push doesn't manage stay attached by it.
func (a *Agent) knownID(tool agenttools.Tool) (string, bool) {
	var ids []string
	for id, t := range a.toolsByID {
		if sameTool(t, tool) {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		return "", false
	}
	sort.Strings(ids)
	return ids[0], true
}

// toolTypeOf is the type a tool's config gives, such as "mcp".
func toolTypeOf(tool agenttools.Tool) string {
	if other, ok := tool.(*agenttools.OtherTool); ok {
		return other.ToolType()
	}
	return string(tool.Kind())
}

// workflowTools are the agenttools tools the workflow's nodes use: the built workflow's Tools, or
// for a loaded one not read yet, the known tools its ids stand for.
func (a *Agent) workflowTools() []agenttools.Tool {
	var out []agenttools.Tool
	if a.workflow != nil {
		for _, tool := range a.workflow.Tools() {
			if t, ok := tool.(agenttools.Tool); ok {
				out = append(out, t)
			}
		}
		return out
	}
	for _, id := range workflowToolIDs(a.rawWorkflow, a.rawNodeOrder) {
		if tool, ok := a.toolsByID[id]; ok {
			out = append(out, tool)
		}
	}
	return out
}

// ToolIDs are the ids of existing tools attached as they are, which Push doesn't manage.
func (a *Agent) ToolIDs() []string { return append([]string{}, a.toolIDs...) }

// Validate checks everything that doesn't need tool ids: the tools, the workflow, the settings,
// and that no two tools share a name. CI can run it without an API key. It returns a
// *ValidationError listing every problem, if the agent is invalid. opts may be nil.
func (a *Agent) Validate(opts *ValidateOptions) error {
	var onWarning func(Issue)
	if opts != nil {
		onWarning = opts.OnWarning
	}
	_, err := a.assemble(nil, false, onWarning)
	return err
}

// ToJSON validates the agent and returns its config in API format (snake_case keys; Go maps have
// no order, and the canonical writers sort them): the shape of a CLI agent_configs/*.json file,
// and the body Push sends. It needs the id of every tool Push manages, in opts.ToolIDs. Numbers
// are json.Number values written as JavaScript writes them. opts may be nil.
func (a *Agent) ToJSON(opts *BuildOptions) (map[string]any, error) {
	var toolIDs map[string]string
	var onWarning func(Issue)
	if opts != nil {
		toolIDs, onWarning = opts.ToolIDs, opts.OnWarning
	}
	if toolIDs == nil {
		toolIDs = map[string]string{}
	}
	return a.assemble(toolIDs, true, onWarning)
}

// Build validates the agent like ToJSON and returns its config as the generated request model
// client.Agents.Create takes (client.Agents.Update takes the same fields). The generated types
// can't hold fields this SDK doesn't model, so they are dropped here; Push sends ToJSON's output,
// which keeps them. opts may be nil.
func (a *Agent) Build(opts *BuildOptions) (*types.BodyCreateAgentV1ConvaiAgentsCreatePost, error) {
	wire, err := a.ToJSON(opts)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(wire)
	if err == nil {
		var body types.BodyCreateAgentV1ConvaiAgentsCreatePost
		if err = json.Unmarshal(data, &body); err == nil {
			return &body, nil
		}
	}
	return nil, &ValidationError{AgentName: a.name, Issues: []Issue{{
		Code:     InvalidModel,
		Severity: SeverityError,
		Message:  fmt.Sprintf("The agent config doesn't match the generated create-agent request model: %v. A loaded field likely has the wrong shape; fix it in the source JSON, or send ToJSON's output", err),
	}}}
}

// assemble builds the API-format config. Without resolve (validation only), tools stand under
// placeholder ids and missing ids aren't reported.
func (a *Agent) assemble(toolIDs map[string]string, resolve bool, onWarning func(Issue)) (map[string]any, error) {
	var issues []Issue
	config := wirejson.Clone(a.base).(map[string]any)
	idOf := func(tool agentworkflows.Tool) string {
		// A kind Push doesn't manage keeps the id it was loaded with; a missing one is reported as
		// invalid_tool below.
		if t, ok := tool.(agenttools.Tool); ok && t != nil && !managedTool(t) {
			if id, known := a.knownID(t); known {
				return id
			}
		} else if id, ok := toolIDs[tool.Name()]; ok {
			return id
		}
		return "<tool:" + tool.Name() + ">"
	}

	if strings.TrimSpace(a.name) == "" {
		issues = append(issues, Issue{Code: MissingName, Severity: SeverityError, Message: "The agent needs a name"})
	}

	// Raw options first, so the typed ones win and can be checked against them.
	raw := map[string]any{}
	for _, settings := range a.rawConversationConfig {
		if wire, err := wirejson.Normalize(settings); err != nil {
			issues = append(issues, invalidConfig("ConversationConfig", err))
		} else if obj, ok := wire.(map[string]any); ok {
			deepMerge(raw, map[string]any{"conversation_config": obj})
		}
	}
	for _, settings := range a.rawPlatformSettings {
		if wire, err := wirejson.Normalize(settings); err != nil {
			issues = append(issues, invalidConfig("PlatformSettings", err))
		} else if obj, ok := wire.(map[string]any); ok {
			deepMerge(raw, map[string]any{"platform_settings": obj})
		}
	}
	deepMerge(config, raw)

	for _, option := range typedOptions {
		piece := option.piece(&a.typed)
		if piece == nil {
			continue
		}
		path, wire, err := typedWire(piece)
		if err != nil {
			issues = append(issues, invalidConfig(option.name, err))
			continue
		}
		if isSet(raw, path) {
			issues = append(issues, conflict(option.name, path))
		}
		setPath(config, path, wire)
	}

	// Tools: each one validates, and no two different tools share a name.
	var workflowTools []agentworkflows.Tool
	if a.workflow != nil {
		workflowTools = a.workflow.Tools()
	}
	everyTool := append([]agenttools.Tool{}, a.tools...)
	everyTool = append(everyTool, a.workflowTools()...)
	byName := map[string]agenttools.Tool{}
	for _, tool := range everyTool {
		if tool == nil {
			issues = append(issues, Issue{Code: InvalidTool, Severity: SeverityError, Message: "A tool is nil", Path: "tools"})
			continue
		}
		other, seen := byName[tool.Name()]
		if seen && !sameTool(other, tool) {
			issues = append(issues, Issue{Code: DuplicateToolName, Severity: SeverityError, Message: fmt.Sprintf("Two different tools are named %q", tool.Name()), Path: "tools." + tool.Name()})
		}
		if !seen {
			byName[tool.Name()] = tool
			if _, err := tool.ToJSON(); err != nil {
				issues = append(issues, Issue{Code: InvalidTool, Severity: SeverityError, Message: err.Error(), Path: "tools." + tool.Name()})
			}
		}
		// Push only creates and updates webhook and client tools; any other kind needs an id.
		if tool.Kind() != agenttools.ToolKindSystem && !managedTool(tool) {
			if _, known := a.knownID(tool); !known {
				issues = append(issues, Issue{Code: InvalidTool, Severity: SeverityError, Message: fmt.Sprintf("Tool %q has type %q, which push doesn't create or update: load it with its id or attach it with ToolIDs", tool.Name(), toolTypeOf(tool)), Path: "tools." + tool.Name()})
			}
		}
	}

	// Webhook and client tools by their pushed ids, other kinds by the ids they were loaded with.
	var attached []string
	var system []*agenttools.SystemTool
	for _, tool := range a.tools {
		if tool == nil {
			continue
		}
		if st := tool.AsSystem(); st != nil {
			system = append(system, st)
		} else {
			attached = append(attached, idOf(tool))
		}
	}
	if a.toolListGiven {
		if isSet(raw, toolIDsPath) {
			issues = append(issues, conflict("Tools", toolIDsPath))
		}
		ids := append(append([]string{}, a.toolIDs...), attached...)
		setPath(config, toolIDsPath, anyList(keepOrder(ids, a.loadedToolIDs)))
	}
	for _, tool := range system {
		if tool.SystemToolType() == "" {
			continue // reported as invalid_tool
		}
		path := append(append([]string{}, builtInToolsPath...), tool.SystemToolType())
		if isSet(raw, path) {
			issues = append(issues, conflict("Tools", path))
		}
		if loaded, ok := a.loadedSystem[tool]; ok && loaded.baseline != "" && toolText(tool) == loaded.baseline {
			setPath(config, path, wirejson.Clone(loaded.raw))
		} else if wire, err := tool.ToJSON(); err == nil {
			setPath(config, path, wire)
		}
	}

	// Workflow: its own checks, plus its nodes can only use webhook and client tools.
	if a.workflow != nil {
		ids := map[string]string{}
		for _, tool := range workflowTools {
			t, ok := tool.(agenttools.Tool)
			if !ok {
				issues = append(issues, Issue{Code: InvalidTool, Severity: SeverityError, Message: fmt.Sprintf("Workflow tool %q is a %T, not an agenttools.Tool, so it can't be pushed", tool.Name(), tool), Path: "workflow.tools." + tool.Name()})
				ids[tool.Name()] = idOf(tool)
				continue
			}
			if t.Kind() == agenttools.ToolKindSystem {
				issues = append(issues, Issue{Code: SystemToolInWorkflow, Severity: SeverityError, Message: fmt.Sprintf("System tool %q can't be used in a workflow node; add it to the agent's tools", t.Name()), Path: "workflow.tools." + t.Name()})
			}
			ids[t.Name()] = idOf(t)
		}
		if wire, err := a.workflow.ToJSON(&agentworkflows.BuildOptions{ToolIDs: ids}); err != nil {
			issues = append(issues, Issue{Code: InvalidWorkflow, Severity: SeverityError, Message: err.Error(), Path: "workflow"})
		} else {
			config["workflow"] = wire
		}
	} else if a.rawWorkflow != nil {
		// Check the loaded graph without replacing the raw copy, which keeps fields this SDK
		// doesn't model.
		wf, err := agentworkflows.FromJSON(a.rawWorkflowJSON)
		if err == nil {
			_, err = wf.ToJSON(nil)
		}
		if err != nil {
			issues = append(issues, Issue{Code: InvalidWorkflow, Severity: SeverityError, Message: err.Error(), Path: "workflow"})
		}
		// Its managed tools may have new ids, as when pushed as a new agent.
		workflow := wirejson.Clone(a.rawWorkflow).(map[string]any)
		mapWorkflowToolIDs(workflow, func(id string) string {
			tool, ok := a.toolsByID[id]
			if !ok || tool == nil || !managedTool(tool) {
				return id
			}
			if newID, ok := toolIDs[tool.Name()]; ok {
				return newID
			}
			return id
		})
		config["workflow"] = workflow
	}

	if resolve {
		for _, tool := range a.Tools() {
			if _, ok := toolIDs[tool.Name()]; !ok {
				issues = append(issues, Issue{Code: UnresolvedTool, Severity: SeverityError, Message: fmt.Sprintf("Tool %q has no id yet. Push the agent with client.Agents.Push, or pass BuildOptions.ToolIDs", tool.Name()), Path: "tools." + tool.Name()})
			}
		}
	}

	if prompt, _ := getPath(config, promptPath); !isString(prompt) {
		issues = append(issues, Issue{Code: MissingPrompt, Severity: SeverityWarning, Message: "The agent needs a prompt", Path: strings.Join(promptPath, ".")})
	}

	issues = dedupe(issues)
	var errs []Issue
	for _, issue := range issues {
		if issue.Severity == SeverityError {
			errs = append(errs, issue)
		}
	}
	if len(errs) > 0 {
		return nil, &ValidationError{AgentName: a.name, Issues: errs}
	}
	if onWarning != nil {
		for _, issue := range issues {
			if issue.Severity == SeverityWarning {
				onWarning(issue)
			}
		}
	}

	config["name"] = a.name
	if config["conversation_config"] == nil {
		config["conversation_config"] = map[string]any{}
	}
	if config["platform_settings"] == nil {
		config["platform_settings"] = map[string]any{}
	}
	if a.hasTags {
		config["tags"] = wirejson.Clone(a.tags)
	} else {
		config["tags"] = []any{}
	}
	canonical, err := wirejson.Canonical(config)
	if err != nil {
		return nil, err
	}
	return canonical.(map[string]any), nil
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}

func invalidConfig(option string, err error) Issue {
	return Issue{Code: InvalidConfig, Severity: SeverityError, Message: fmt.Sprintf("Invalid %s: %v", option, err), Path: option}
}

func conflict(option string, path []string) Issue {
	joined := strings.Join(path, ".")
	return Issue{Code: OptionConflict, Severity: SeverityError, Message: fmt.Sprintf("%s and the raw config both set %s; set it in one place", option, joined), Path: joined}
}

func dedupe(issues []Issue) []Issue {
	seen := map[Issue]bool{}
	var out []Issue
	for _, issue := range issues {
		if !seen[issue] {
			seen[issue] = true
			out = append(out, issue)
		}
	}
	return out
}
