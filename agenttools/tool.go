package agenttools

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// ToolKind says what kind of tool a Tool is.
type ToolKind string

const (
	// ToolKindWebhook: the server calls an HTTP endpoint.
	ToolKindWebhook ToolKind = "webhook"
	// ToolKindClient: the agent sends an event to the user's client.
	ToolKindClient ToolKind = "client"
	// ToolKindSystem: a built-in tool, such as ending or transferring the call.
	ToolKindSystem ToolKind = "system"
	// ToolKindOther: a loaded tool of a type the builder has no constructor for (such as mcp). It
	// round-trips through ToJSON but can't be edited.
	ToolKindOther ToolKind = "other"
)

// Tool is an agent tool written as code: a *WebhookTool, *ClientTool, *SystemTool or *OtherTool.
// Create one with Webhook, Client or a System* constructor, or load one with FromJSON /
// FromResponse.
//
// Chaining never fails: every problem is recorded and reported at once by Build, BuildConfig,
// ToJSON and ToJSONString as a *ValidationError. A tool is not safe for concurrent use.
type Tool interface {
	// Name is the tool's name, as the LLM sees it.
	Name() string
	// Kind is the kind of tool.
	Kind() ToolKind
	// AsWebhook is the tool as a *WebhookTool, or nil if it's another kind.
	AsWebhook() *WebhookTool
	// AsClient is the tool as a *ClientTool, or nil if it's another kind.
	AsClient() *ClientTool
	// AsSystem is the tool as a *SystemTool, or nil if it's another kind.
	AsSystem() *SystemTool
	// Build validates the tool and returns the request model client.Agents.Tools.Create
	// and Update accept. opts may be nil.
	Build(opts *BuildOptions) (*types.ToolRequestModel, error)
	// BuildConfig validates the tool and returns only its config. For a system tool, its System
	// field is the agent's prompt.built_in_tools entry. Webhook and client tools go to tools.create
	// and are attached to an agent by id in prompt.tool_ids. opts may be nil.
	BuildConfig(opts *BuildOptions) (*types.ToolRequestModelToolConfig, error)
	// ToJSON validates like Build and returns the tool config in API wire format (the CLI's
	// tool_configs/*.json shape, with its "type"). Numbers are json.Number values written as
	// JavaScript writes them.
	ToJSON() (map[string]any, error)
	// ToJSONString validates like Build and returns the canonical file format: ToJSON's output
	// with keys sorted at every level, 2-space indentation and a trailing newline, byte-identical
	// to every other ElevenLabs SDK's output for the same program.
	ToJSONString() (string, error)

	impl() toolImpl
}

// toolImpl is what each kind provides to the shared build pipeline.
type toolImpl interface {
	base() *core
	// validate covers the builder's rules, checked by ToJSON too.
	validate() []Issue
	// modelIssues covers loaded input that ToJSON writes back as loaded but the generated request
	// model can't take, so only Build reports it.
	modelIssues() []Issue
	// toWire is the tool config in wire format, with its type (unsorted).
	toWire() map[string]any
}

// core is the state every kind shares.
type core struct {
	// config holds the tool config fields the kind doesn't keep in its own fields.
	config map[string]any
	// issues are problems found while chaining.
	issues []Issue
}

func (c *core) base() *core { return c }

func (c *core) name() string {
	name, _ := c.config["name"].(string)
	return name
}

// setOption sets an option that lives on the tool config itself.
func (c *core) setOption(key string, value any) {
	if key != "dynamic_variable_placeholders" {
		c.config[key] = value
		return
	}
	variables, ok := c.config["dynamic_variables"].(map[string]any)
	if !ok {
		variables = map[string]any{}
	}
	variables[key] = value
	c.config["dynamic_variables"] = variables
}

// numberOption reads a numeric option from the config.
func (c *core) numberOption(key string) (float64, bool) {
	n, ok := c.config[key].(json.Number)
	if !ok {
		return 0, false
	}
	f, err := strconv.ParseFloat(string(n), 64)
	return f, err == nil
}

// FromJSON loads a tool stored in API wire format (snake_case keys), for example ToJSONString's
// output or a CLI tool_configs/*.json file. data is JSON text (a string, []byte or
// json.RawMessage) or a parsed value (such as a map[string]any or a generated model), holding a
// bare tool config or a request or response with a "tool_config" key. It returns a tool of the
// matching kind; a type the builder has no constructor for loads as an *OtherTool, kept exactly as
// loaded.
//
// Document order decides the order of loaded properties, so pass the JSON text where you have it;
// a parsed Go map has no order and is read in sorted key order. Loaded fields the builder doesn't
// model are kept as they are, and so are fields of the wrong shape, which Build reports as
// invalid_model.
func FromJSON(data any) (Tool, error) {
	tree, err := toTree(data)
	if err != nil {
		return nil, err
	}
	top, ok := tree.(*object)
	if !ok {
		return nil, errNotObject("tool")
	}
	if config, has := top.get("tool_config"); has {
		if obj, isObject := config.(*object); isObject {
			top = obj
		}
	}
	switch top.vals["type"] {
	case "webhook", "client", "system":
	default:
		return &OtherTool{raw: plain(top).(map[string]any)}, nil
	}
	config := stripNulls(top).(*object)
	switch config.vals["type"] {
	case "webhook":
		return loadWebhook(config), nil
	case "client":
		return loadClient(config), nil
	}
	return loadSystem(config), nil
}

// FromResponse loads the result of tools.get: a *types.ToolResponseModel, or its ToolConfig.
func FromResponse(response any) (Tool, error) {
	switch v := response.(type) {
	case *types.ToolResponseModel:
		if v == nil || v.ToolConfig == nil {
			return nil, errors.New("agenttools: the tool response has no tool_config")
		}
		return FromJSON(v.ToolConfig)
	case types.ToolResponseModel:
		return FromResponse(&v)
	case *types.ToolResponseModelToolConfig:
		if v == nil {
			return nil, errors.New("agenttools: the tool config is nil")
		}
		return FromJSON(v)
	case types.ToolResponseModelToolConfig:
		return FromJSON(v)
	}
	return nil, fmt.Errorf("agenttools: FromResponse takes a get-tool response or its tool config, not %T", response)
}

// configOf splits a loaded tool config into its plain fields, without "type" and the keys the kind
// keeps itself.
func configOf(o *object, skip ...string) map[string]any {
	config := map[string]any{}
	for _, key := range o.keys {
		if key == "type" || contains(skip, key) {
			continue
		}
		config[key] = plain(o.vals[key])
	}
	return config
}

// checked runs validation: errors stop it, and warnings go to onWarning.
func checked(t toolImpl, onWarning func(Issue)) (map[string]any, error) {
	c := t.base()
	issues := append(append([]Issue{}, c.issues...), t.validate()...)
	var errs []Issue
	for _, issue := range issues {
		if issue.Severity == SeverityError {
			errs = append(errs, issue)
		}
	}
	if len(errs) > 0 {
		return nil, &ValidationError{ToolName: c.name(), Issues: errs}
	}
	if onWarning != nil {
		for _, issue := range issues {
			if issue.Severity == SeverityWarning {
				onWarning(issue)
			}
		}
	}
	return t.toWire(), nil
}

func toJSON(t toolImpl) (map[string]any, error) {
	wire, err := checked(t, nil)
	if err != nil {
		return nil, err
	}
	canonical, err := wirejson.Canonical(wire)
	if err != nil {
		return nil, err
	}
	return canonical.(map[string]any), nil
}

func toJSONString(t toolImpl) (string, error) {
	wire, err := checked(t, nil)
	if err != nil {
		return "", err
	}
	return wirejson.String(wire)
}

func build(t toolImpl, opts *BuildOptions) (*types.ToolRequestModel, error) {
	config, err := buildConfig(t, opts)
	if err != nil {
		return nil, err
	}
	return &types.ToolRequestModel{ToolConfig: config}, nil
}

func buildConfig(t toolImpl, opts *BuildOptions) (*types.ToolRequestModelToolConfig, error) {
	var onWarning func(Issue)
	if opts != nil {
		onWarning = opts.OnWarning
	}
	wire, err := checked(t, onWarning)
	if err != nil {
		return nil, err
	}
	name := t.base().name()
	if issues := t.modelIssues(); len(issues) > 0 {
		return nil, &ValidationError{ToolName: name, Issues: issues}
	}
	return toModel(name, wire)
}

// toModel decodes validated wire output into the generated tool config. Validation covers the
// builder's rules, not every loaded field's shape, so a loaded field of the wrong shape can still
// fail here; that is reported as invalid_model.
//
// The generated model must also send every constant exactly as ToJSON writes it. It doesn't for a
// nullable literal's null constant (it reads null as unset), for an empty object or array constant
// (which it leaves out), or for numbers inside object and array constants that a float64 can't
// hold, and validation can't foresee every such conversion; so every constant is compared with the
// model's own JSON, and one that would be changed or dropped is reported as
// constant_type_mismatch rather than sent altered.
func toModel(name string, wire map[string]any) (*types.ToolRequestModelToolConfig, error) {
	canonical, err := wirejson.Canonical(wire)
	if err != nil {
		return nil, modelError(name, err)
	}
	data, err := json.Marshal(canonical)
	if err != nil {
		return nil, modelError(name, err)
	}
	var config types.ToolRequestModelToolConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, modelError(name, err)
	}
	// The generated query params model has no omitempty on `properties`, so a loaded empty
	// query_params_schema ({}) would be sent as "properties": null. Send no properties as an
	// empty object instead, which is what ToJSON's {} means.
	if webhook := config.Webhook; webhook != nil && webhook.ApiSchema != nil {
		if query := webhook.ApiSchema.QueryParamsSchema; query != nil && query.Properties == nil {
			query.Properties = map[string]*types.QueryParamsJsonSchemaInputPropertiesValue{}
		}
	}
	wanted := constants(canonical)
	sentData, err := json.Marshal(config)
	if err != nil {
		// A null constant leaves the generated constant union empty, which can't be marshaled.
		var issues []Issue
		for _, c := range wanted {
			if c.value == nil {
				issues = append(issues, nullConstantIssue(c.display))
			}
		}
		if len(issues) > 0 {
			return nil, &ValidationError{ToolName: name, Issues: issues}
		}
		return nil, modelError(name, err)
	}
	if len(wanted) > 0 {
		sentTree, err := wirejson.Decode(sentData)
		if err == nil {
			sentTree, err = wirejson.Canonical(sentTree)
		}
		if err != nil {
			return nil, modelError(name, err)
		}
		sent := map[string]any{}
		for _, c := range constants(sentTree) {
			sent[c.key] = c.value
		}
		var issues []Issue
		for _, c := range wanted {
			got, ok := sent[c.key]
			switch {
			case ok && reflect.DeepEqual(got, c.value):
				continue
			case c.value == nil:
				issues = append(issues, nullConstantIssue(c.display))
			case ok:
				issues = append(issues, errorIssue(ConstantTypeMismatch, c.display, fmt.Sprintf("%q has the constant %s, which this SDK's generated request model would send as %s. Use a constant it can represent exactly", c.display, jsonText(c.value), jsonText(got))))
			default:
				issues = append(issues, errorIssue(ConstantTypeMismatch, c.display, fmt.Sprintf("%q has the constant %s, which this SDK's generated request model would drop. Send the ToJSON output directly, or use another constant", c.display, jsonText(c.value))))
			}
		}
		if len(issues) > 0 {
			return nil, &ValidationError{ToolName: name, Issues: issues}
		}
	}
	return &config, nil
}

func modelError(name string, err error) error {
	return &ValidationError{ToolName: name, Issues: []Issue{errorIssue(InvalidModel, "", fmt.Sprintf("The tool doesn't match the generated tool config model: %v. A loaded field likely has the wrong shape; fix it in the source JSON", err))}}
}

func nullConstantIssue(path string) Issue {
	return errorIssue(ConstantTypeMismatch, path, fmt.Sprintf("%q has the constant null, which this SDK's generated request model can't send (it reads null as unset). Send the ToJSON output directly, or use another constant", path))
}

func jsonText(value any) string {
	text, err := wirejson.String(value)
	if err != nil {
		return fmt.Sprint(value)
	}
	return strings.Join(strings.Fields(text), " ")
}

// constantAt is one schema constant in a tool config: where it is (key, unique; display, as
// validation names it) and its canonical value.
type constantAt struct {
	key, display string
	value        any
}

// constants lists every schema constant in a tool config (wire format, or the generated model's
// JSON). Free-form JSON values (constants, dynamic variable placeholders, headers) aren't schemas,
// so they aren't searched, and property and path param names are walked as names, so a param
// named "constant_value" isn't read as a constant.
func constants(value any) []constantAt {
	var out []constantAt
	var walk func(value any, key []string, display []string)
	walk = func(value any, key []string, display []string) {
		switch v := value.(type) {
		case map[string]any:
			for _, k := range wirejson.SortKeys(mapKeys(v)) {
				item := v[k]
				switch k {
				case "constant_value":
					out = append(out, constantAt{key: strings.Join(key, "\x00"), display: strings.Join(display, ""), value: item})
				case "dynamic_variable_placeholders", "request_headers":
				case "properties", "path_params_schema":
					names, _ := item.(map[string]any)
					for _, name := range wirejson.SortKeys(mapKeys(names)) {
						label := name
						if k == "path_params_schema" {
							label = "path_params." + name
						}
						walk(names[name], append(key, "k:"+k, "n:"+name), append(display, dot(display, label)))
					}
				default:
					walk(item, append(key, "k:"+k), append(display, displayStep(display, k)))
				}
			}
		case []any:
			for i, item := range v {
				walk(item, append(key, "i:"+strconv.Itoa(i)), append(display, fmt.Sprintf("[%d]", i)))
			}
		}
	}
	walk(value, nil, nil)
	return out
}

func mapKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// displayStep names a config key in a constant's location, as validation names schemas: "body",
// "query", "params", "items" as "[]", and wrapper keys left out.
func displayStep(display []string, key string) string {
	switch key {
	case "api_schema":
		return ""
	case "items":
		return "[]"
	case "query_params_schema":
		return dot(display, "query")
	case "request_body_schema":
		return dot(display, "body")
	case "response_body_schema":
		return dot(display, "response")
	case "parameters":
		return dot(display, "params")
	}
	return dot(display, key)
}

func dot(display []string, label string) string {
	if strings.Join(display, "") == "" {
		return label
	}
	return "." + label
}

// OtherTool is a loaded tool of a type the builder has no constructor for, such as mcp. It is kept
// exactly as loaded, and Build warns with unknown_tool_type.
type OtherTool struct {
	raw    map[string]any
	issues []Issue
}

func (t *OtherTool) impl() toolImpl         { return t }
func (t *OtherTool) base() *core            { return &core{config: t.raw, issues: t.issues} }
func (t *OtherTool) modelIssues() []Issue   { return nil }
func (t *OtherTool) toWire() map[string]any { return cloneMap(t.raw) }

// Name is the tool's name, as the LLM sees it.
func (t *OtherTool) Name() string { return t.base().name() }

// Kind is ToolKindOther.
func (t *OtherTool) Kind() ToolKind { return ToolKindOther }

// ToolType is the loaded "type", such as "mcp".
func (t *OtherTool) ToolType() string {
	toolType, _ := t.raw["type"].(string)
	return toolType
}

// AsWebhook is nil: the tool isn't a webhook tool.
func (t *OtherTool) AsWebhook() *WebhookTool { return nil }

// AsClient is nil: the tool isn't a client tool.
func (t *OtherTool) AsClient() *ClientTool { return nil }

// AsSystem is nil: the tool isn't a system tool.
func (t *OtherTool) AsSystem() *SystemTool { return nil }

// Options applies to no field of a tool the builder can't edit: every option set is reported by
// Build as invalid_option.
func (t *OtherTool) Options(opts *ToolOptions) *OtherTool {
	acceptedOptions(opts.entries(), nil, "other", &t.issues)
	return t
}

func (t *OtherTool) validate() []Issue {
	return []Issue{{
		Code:     UnknownToolType,
		Severity: SeverityWarning,
		Message:  fmt.Sprintf("Tool %q has type %q, which this SDK can't edit. It round-trips through ToJSON, but this SDK's generated request model may not hold it; send ToJSON's output directly", t.Name(), t.ToolType()),
	}}
}

// Build validates the tool (warning unknown_tool_type) and returns it as the generated request
// model, if the generated types model its type; otherwise it reports invalid_model.
func (t *OtherTool) Build(opts *BuildOptions) (*types.ToolRequestModel, error) { return build(t, opts) }

// BuildConfig is Build's tool config.
func (t *OtherTool) BuildConfig(opts *BuildOptions) (*types.ToolRequestModelToolConfig, error) {
	return buildConfig(t, opts)
}

// ToJSON returns the loaded config in API wire format.
func (t *OtherTool) ToJSON() (map[string]any, error) { return toJSON(t) }

// ToJSONString returns the loaded config in the canonical file format.
func (t *OtherTool) ToJSONString() (string, error) { return toJSONString(t) }
