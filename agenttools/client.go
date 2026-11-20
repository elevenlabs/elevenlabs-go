package agenttools

import (
	"github.com/elevenlabs/elevenlabs-go/types"
)

// ClientTool is a tool that sends an event to the user's client. Create one with Client.
type ClientTool struct {
	core
	// parameters is an object, unless a loaded config declared another type (reported by Build).
	// Loaded parameters that aren't a JSON object stay in config, written back as loaded.
	parameters Schema
	// malformedParameters is whether config holds loaded parameters that aren't a JSON object.
	malformedParameters bool
}

// Client is a tool that sends an event to the user's client.
func Client(name, description string) *ClientTool {
	return &ClientTool{core: core{config: map[string]any{"name": name, "description": description}}}
}

func loadClient(o *object) *ClientTool {
	t := &ClientTool{core: core{config: configOf(o, "parameters")}}
	if value, ok := o.get("parameters"); ok {
		if params, isObject := value.(*object); isObject {
			t.parameters = containerSchema(params)
		} else {
			t.config["parameters"] = plain(value)
			t.malformedParameters = true
		}
	}
	return t
}

func (t *ClientTool) impl() toolImpl { return t }

// Name is the tool's name, as the LLM sees it.
func (t *ClientTool) Name() string { return t.name() }

// Kind is ToolKindClient.
func (t *ClientTool) Kind() ToolKind { return ToolKindClient }

// AsWebhook is nil: the tool isn't a webhook tool.
func (t *ClientTool) AsWebhook() *WebhookTool { return nil }

// AsClient is the tool itself.
func (t *ClientTool) AsClient() *ClientTool { return t }

// AsSystem is nil: the tool isn't a system tool.
func (t *ClientTool) AsSystem() *SystemTool { return nil }

// Build validates the tool and returns the request model tools.create and tools.update accept.
// opts may be nil.
func (t *ClientTool) Build(opts *BuildOptions) (*types.ToolRequestModel, error) {
	return build(t, opts)
}

// BuildConfig validates the tool and returns only its config (Client set).
func (t *ClientTool) BuildConfig(opts *BuildOptions) (*types.ToolRequestModelToolConfig, error) {
	return buildConfig(t, opts)
}

// ToJSON validates like Build and returns the tool config in API wire format.
func (t *ClientTool) ToJSON() (map[string]any, error) { return toJSON(t) }

// ToJSONString validates like Build and returns the canonical file format.
func (t *ClientTool) ToJSONString() (string, error) { return toJSONString(t) }

// Options merges options for client tools (see ToolOptions). An option that doesn't apply is
// reported by Build as invalid_option.
func (t *ClientTool) Options(opts *ToolOptions) *ClientTool {
	for _, entry := range acceptedOptions(opts.entries(), clientOptions, "client", &t.issues) {
		t.setOption(entry.key, entry.value)
	}
	return t
}

// Param sets the param name, or replaces it in place: parameters.properties[name], and
// parameters.required unless the schema is Optional. Any schema kind. Loaded params that aren't an
// object are left as they are, and Build reports them.
func (t *ClientTool) Param(name string, schema Schema) *ClientTool {
	switch parameters := t.parameters.(type) {
	case nil:
		if !t.malformedParameters {
			t.parameters = Object().Property(name, schema)
		}
	case ObjectSchema:
		t.parameters = parameters.Property(name, schema)
	}
	return t
}

// RemoveParam removes the param name. Removing one that doesn't exist does nothing; removing the
// last one leaves parameters out.
func (t *ClientTool) RemoveParam(name string) *ClientTool {
	if parameters, ok := t.parameters.(ObjectSchema); ok && parameters.PropertySchema(name) != nil {
		remaining := parameters.Without(name)
		if len(remaining.PropertyNames()) == 0 {
			t.parameters = nil
		} else {
			t.parameters = remaining
		}
	}
	return t
}

// ParametersSchema is the params as an object schema, for reading. It is the zero ObjectSchema
// (no properties) if there are none; ok is false if loaded params aren't an object.
func (t *ClientTool) ParametersSchema() (schema ObjectSchema, ok bool) {
	switch parameters := t.parameters.(type) {
	case nil:
		return Object(), !t.malformedParameters
	case ObjectSchema:
		return parameters, true
	}
	return ObjectSchema{}, false
}

func (t *ClientTool) validate() []Issue {
	var issues []Issue
	if t.parameters != nil {
		if t.parameters.Kind() != SchemaKindObject {
			issues = append(issues, errorIssue(BodyNotObject, "params", "The params schema must be an object"))
		}
		// Checks the params object's own RequireAnyOf/RequireAllOf and every param.
		checkSchema(t.parameters, "params", false, &issues)
	}
	timeout, hasTimeout := t.numberOption("response_timeout_secs")
	if hasTimeout && timeout != -1 && (timeout < 1 || timeout > 120) {
		issues = append(issues, errorIssue(InvalidOption, "response_timeout_secs", "A client tool's response_timeout_secs must be 1–120, or -1"))
	}
	if hasTimeout && timeout == -1 && t.config["expects_response"] != true {
		issues = append(issues, errorIssue(InvalidOption, "response_timeout_secs", "response_timeout_secs: -1 needs expects_response: true"))
	}
	return issues
}

func (t *ClientTool) modelIssues() []Issue {
	var issues []Issue
	if t.malformedParameters {
		issues = append(issues, errorIssue(InvalidModel, "params", "The loaded parameters aren't an object, which the generated tool config model can't take. Fix it in the source JSON"))
	}
	if t.parameters != nil {
		checkModelShape(t.parameters, "params", &issues)
	}
	return issues
}

func (t *ClientTool) toWire() map[string]any {
	wire := cloneMap(t.config)
	wire["type"] = "client"
	if t.parameters != nil {
		wire["parameters"] = t.parameters.state().toWire()
	}
	return wire
}
