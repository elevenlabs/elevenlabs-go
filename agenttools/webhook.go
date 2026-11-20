package agenttools

import (
	"fmt"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/types"
)

// WebhookTool is a tool the server runs by calling an HTTP endpoint. Create one with Webhook. Every
// method returns the tool, so a tool is written as one chain.
type WebhookTool struct {
	core
	// api holds the api_schema fields not kept below: url, method, auth, content type and so on,
	// and loaded fields of the wrong shape, written back as loaded.
	api map[string]any
	// apiMalformed is a loaded api_schema that isn't an object.
	apiMalformed any
	pathParams   []property
	// query is an object, unless a loaded config declared another type (reported by Build).
	query Schema
	// queryTypeLoaded is whether a loaded query_params_schema declared a `type`. The API's
	// query_params_schema has none, so the one the builder writes is left out otherwise.
	queryTypeLoaded bool
	// queryUntouched is whether the loaded query_params_schema hasn't been edited, so it is written
	// back whatever it holds.
	queryUntouched bool
	body           Schema
	response       Schema
	headers        []headerEntry
	// malformed lists the loaded api_schema keys of the wrong shape, kept in api.
	malformed []string
}

type headerEntry struct {
	name  string
	value any
}

// Webhook is a tool the server runs by calling an HTTP endpoint. Set the method and URL with Get,
// Post, Put, Patch or Delete.
func Webhook(name, description string) *WebhookTool {
	return &WebhookTool{
		core: core{config: map[string]any{"name": name, "description": description}},
		api:  map[string]any{},
	}
}

func loadWebhook(o *object) *WebhookTool {
	t := &WebhookTool{core: core{config: configOf(o, "api_schema")}, api: map[string]any{}}
	value, ok := o.get("api_schema")
	if !ok {
		return t
	}
	api, isObject := value.(*object)
	if !isObject {
		t.apiMalformed = plain(value)
		return t
	}
	for _, key := range api.keys {
		value := api.vals[key]
		switch key {
		case "path_params_schema":
			if params, ok := schemaMap(value); ok {
				for _, name := range params.keys {
					child := params.vals[name].(*object)
					t.pathParams = append(t.pathParams, property{name: name, schema: loadSchema(child, false, kindOf(child))})
				}
				continue
			}
		case "query_params_schema":
			if query, ok := value.(*object); ok {
				t.query = containerSchema(query)
				_, t.queryTypeLoaded = query.get("type")
				t.queryUntouched = true
				continue
			}
		case "request_body_schema", "response_body_schema":
			if schema, ok := value.(*object); ok {
				if key == "request_body_schema" {
					t.body = containerSchema(schema)
				} else {
					t.response = containerSchema(schema)
				}
				continue
			}
		case "request_headers":
			if headers, ok := value.(*object); ok {
				// Header values of a kind Header doesn't cover are kept exactly as loaded.
				for _, name := range headers.keys {
					t.headers = append(t.headers, headerEntry{name: name, value: plain(headers.vals[name])})
				}
				continue
			}
		default:
			t.api[key] = plain(value)
			continue
		}
		// A modeled field of the wrong shape is kept as loaded; Build reports it.
		t.api[key] = plain(value)
		t.malformed = append(t.malformed, key)
	}
	return t
}

// schemaMap reads a map of names to schemas, reporting false if a value isn't an object.
func schemaMap(value any) (*object, bool) {
	names, ok := value.(*object)
	if !ok {
		return nil, false
	}
	for _, name := range names.keys {
		if _, isObject := names.vals[name].(*object); !isObject {
			return nil, false
		}
	}
	return names, true
}

func (t *WebhookTool) impl() toolImpl { return t }

// Name is the tool's name, as the LLM sees it.
func (t *WebhookTool) Name() string { return t.name() }

// Kind is ToolKindWebhook.
func (t *WebhookTool) Kind() ToolKind { return ToolKindWebhook }

// AsWebhook is the tool itself.
func (t *WebhookTool) AsWebhook() *WebhookTool { return t }

// AsClient is nil: the tool isn't a client tool.
func (t *WebhookTool) AsClient() *ClientTool { return nil }

// AsSystem is nil: the tool isn't a system tool.
func (t *WebhookTool) AsSystem() *SystemTool { return nil }

// Build validates the tool and returns the request model tools.create and tools.update accept.
// opts may be nil.
func (t *WebhookTool) Build(opts *BuildOptions) (*types.ToolRequestModel, error) {
	return build(t, opts)
}

// BuildConfig validates the tool and returns only its config (Webhook set).
func (t *WebhookTool) BuildConfig(opts *BuildOptions) (*types.ToolRequestModelToolConfig, error) {
	return buildConfig(t, opts)
}

// ToJSON validates like Build and returns the tool config in API wire format.
func (t *WebhookTool) ToJSON() (map[string]any, error) { return toJSON(t) }

// ToJSONString validates like Build and returns the canonical file format.
func (t *WebhookTool) ToJSONString() (string, error) { return toJSONString(t) }

// Options merges options for webhook tools (see ToolOptions). An option that doesn't apply is
// reported by Build as invalid_option.
func (t *WebhookTool) Options(opts *ToolOptions) *WebhookTool {
	for _, entry := range acceptedOptions(opts.entries(), webhookOptions, "webhook", &t.issues) {
		if entry.key == "content_type" || entry.key == "response_filter" {
			t.editAPI()[entry.key] = entry.value
		} else {
			t.setOption(entry.key, entry.value)
		}
	}
	return t
}

// Get sends a GET to url. {name} placeholders in the URL are filled from path params. A later call
// to Get, Post, Put, Patch or Delete replaces the method and URL.
func (t *WebhookTool) Get(url string) *WebhookTool { return t.request("GET", url) }

// Post sends a POST to url. It needs a Body.
func (t *WebhookTool) Post(url string) *WebhookTool { return t.request("POST", url) }

// Put sends a PUT to url. It needs a Body.
func (t *WebhookTool) Put(url string) *WebhookTool { return t.request("PUT", url) }

// Patch sends a PATCH to url. It needs a Body.
func (t *WebhookTool) Patch(url string) *WebhookTool { return t.request("PATCH", url) }

// Delete sends a DELETE to url. It may have a Body or not.
func (t *WebhookTool) Delete(url string) *WebhookTool { return t.request("DELETE", url) }

func (t *WebhookTool) request(method, url string) *WebhookTool {
	api := t.editAPI()
	api["method"] = method
	api["url"] = url
	return t
}

// editAPI is api, for an edit: a loaded api_schema that isn't an object is replaced.
func (t *WebhookTool) editAPI() map[string]any {
	t.apiMalformed = nil
	return t.api
}

// PathParam sets the path param name, filling the {name} placeholder in the URL. Literal schemas
// only. A path param is always required, whatever Optional says. A param that exists is replaced
// in place.
func (t *WebhookTool) PathParam(name string, schema Schema) *WebhookTool {
	if schema == nil {
		schema = LiteralSchema{}
	}
	for i, p := range t.pathParams {
		if p.name == name {
			t.pathParams[i].schema = schema
			return t
		}
	}
	t.pathParams = append(t.pathParams, property{name: name, schema: schema})
	return t
}

// QueryParam sets the query param name, required unless the schema is Optional. Literal schemas
// only. A loaded query params schema that isn't an object is kept as loaded, and Build reports it.
func (t *WebhookTool) QueryParam(name string, schema Schema) *WebhookTool {
	switch query := t.query.(type) {
	case nil:
		t.query = Object().Property(name, schema)
	case ObjectSchema:
		t.query = query.Property(name, schema)
	default:
		return t
	}
	t.queryUntouched = false
	return t
}

// Body sets the request body schema. It must be an object.
func (t *WebhookTool) Body(schema Schema) *WebhookTool {
	t.body = schema
	return t
}

// Response sets the response body schema. It must be an object.
func (t *WebhookTool) Response(schema Schema) *WebhookTool {
	t.response = schema
	return t
}

// Header sets the request header name. A header that exists is replaced in place.
func (t *WebhookTool) Header(name string, value Header) *WebhookTool {
	for i, h := range t.headers {
		if h.name == name {
			t.headers[i].value = value.Wire()
			return t
		}
	}
	t.headers = append(t.headers, headerEntry{name: name, value: value.Wire()})
	return t
}

// Auth authenticates requests with an auth connection.
func (t *WebhookTool) Auth(connection AuthConnection) *WebhookTool {
	t.editAPI()["auth_connection"] = connection.Wire()
	return t
}

// Assign stores the value at valuePath in the response (dot notation, such as "order.status") in
// the dynamic variable variable. opts may be nil.
func (t *WebhookTool) Assign(variable, valuePath string, opts *AssignOptions) *WebhookTool {
	entry := map[string]any{"source": "response", "dynamic_variable": variable, "value_path": valuePath}
	if opts != nil && opts.Sanitize != nil {
		entry["sanitize"] = *opts.Sanitize
	}
	if opts != nil && opts.PreserveNativeType != nil {
		entry["preserve_native_type"] = *opts.PreserveNativeType
	}
	assignments, _ := t.config["assignments"].([]any)
	t.config["assignments"] = append(append([]any(nil), assignments...), entry)
	return t
}

// RemovePathParam removes the path param name. Removing one that doesn't exist does nothing.
func (t *WebhookTool) RemovePathParam(name string) *WebhookTool {
	var kept []property
	for _, p := range t.pathParams {
		if p.name != name {
			kept = append(kept, p)
		}
	}
	t.pathParams = kept
	return t
}

// RemoveQueryParam removes the query param name. Removing one that doesn't exist does nothing.
// Removing the last one leaves query_params_schema out, unless it has loaded fields of its own.
func (t *WebhookTool) RemoveQueryParam(name string) *WebhookTool {
	if query, ok := t.query.(ObjectSchema); ok && query.PropertySchema(name) != nil {
		t.query = query.Without(name)
		t.queryUntouched = false
	}
	return t
}

// RemoveHeader removes the request header name. Removing one that doesn't exist does nothing.
func (t *WebhookTool) RemoveHeader(name string) *WebhookTool {
	var kept []headerEntry
	for _, h := range t.headers {
		if h.name != name {
			kept = append(kept, h)
		}
	}
	t.headers = kept
	return t
}

// RemoveAssignment removes every assignment to variable. Removing the last one leaves assignments
// out; removing one that doesn't exist does nothing, so a loaded empty list round-trips.
func (t *WebhookTool) RemoveAssignment(variable string) *WebhookTool {
	assignments, ok := t.config["assignments"].([]any)
	if !ok {
		return t
	}
	var kept []any
	for _, a := range assignments {
		if m, isMap := a.(map[string]any); isMap && m["dynamic_variable"] == variable {
			continue
		}
		kept = append(kept, a)
	}
	switch {
	case len(kept) == len(assignments):
	case len(kept) == 0:
		delete(t.config, "assignments")
	default:
		t.config["assignments"] = kept
	}
	return t
}

// BodySchema is the request body schema, or nil. Edit it and set it again with Body.
func (t *WebhookTool) BodySchema() Schema { return t.body }

// ResponseSchema is the response body schema, or nil. Edit it and set it again with Response.
func (t *WebhookTool) ResponseSchema() Schema { return t.response }

func (t *WebhookTool) validate() []Issue {
	var issues []Issue
	add := func(code IssueCode, path, message string) {
		issues = append(issues, errorIssue(code, path, message))
	}
	// The API reads an empty URL as unset.
	url, _ := t.api["url"].(string)
	if url == "" {
		add(MissingURL, "", "Set the method and URL with Get, Post, Put, Patch or Delete")
	}
	// As the API does: env var templates aren't placeholders, and the auth connection fills
	// auth_resolved_params.
	names := placeholders(stripEnvTemplates(url))
	authResolved, _ := stringList(t.api["auth_resolved_params"])
	for _, name := range names {
		if !t.hasPathParam(name) && !contains(authResolved, name) {
			add(MissingPathParam, "path_params."+name, fmt.Sprintf("The URL has {%s} but no PathParam(%q)", name, name))
		}
	}
	for _, p := range t.pathParams {
		path := "path_params." + p.name
		switch {
		case url != "" && !contains(names, p.name):
			add(UnknownPathParam, path, fmt.Sprintf("PathParam(%q) has no {%s} in the URL", p.name, p.name))
		case contains(authResolved, p.name):
			add(UnknownPathParam, path, fmt.Sprintf("{%s} is filled by the auth connection, not a PathParam", p.name))
		}
		if !p.schema.Kind().IsLiteral() {
			add(ParamNotLiteral, path, fmt.Sprintf("Path param %q must be a string, integer, number or bool", p.name))
		}
		// A path param is always required (the URL needs it), whatever Optional says: the API
		// schema has no `required` list for path params to leave it out of.
		if p.schema.state().omitted {
			add(OmittedRequired, path, fmt.Sprintf("Path param %q can't use Omit: the URL always needs it", p.name))
		}
		checkSchema(p.schema, path, false, &issues)
	}
	if t.query != nil {
		if t.query.Kind() != SchemaKindObject {
			add(BodyNotObject, "query", "The query params schema must be an object")
			checkSchema(t.query, "query", false, &issues)
		} else {
			query := t.query.(ObjectSchema)
			for _, name := range query.PropertyNames() {
				schema := query.PropertySchema(name)
				path := "query." + name
				if !schema.Kind().IsLiteral() {
					add(ParamNotLiteral, path, fmt.Sprintf("Query param %q must be a string, integer, number or bool", name))
				}
				checkSchema(schema, path, true, &issues)
			}
		}
	}
	for _, b := range []struct {
		label  string
		schema Schema
	}{{"body", t.body}, {"response", t.response}} {
		if b.schema == nil {
			continue
		}
		if b.schema.Kind() != SchemaKindObject {
			add(BodyNotObject, b.label, fmt.Sprintf("The %s schema must be an object", b.label))
		}
		checkSchema(b.schema, b.label, false, &issues)
	}
	method, _ := t.api["method"].(string)
	if method == "GET" && t.body != nil {
		add(BodyNotAllowed, "body", "A GET request can't have a body")
	}
	if (method == "POST" || method == "PUT" || method == "PATCH") && t.body == nil {
		add(MissingBody, "body", fmt.Sprintf("A %s request needs a body", method))
	}

	firstIn := map[string]string{}
	locations := []struct {
		label string
		names []string
	}{{"path_params", t.pathParamNames()}, {"query", propertyNames(t.query)}, {"body", propertyNames(t.body)}}
	for _, location := range locations {
		for _, name := range location.names {
			if first, seen := firstIn[name]; seen {
				add(DuplicateParam, location.label+"."+name, fmt.Sprintf("%q is in both %s and %s; param names must be unique across path, query and body", name, first, location.label))
				continue
			}
			firstIn[name] = location.label
		}
	}
	seenHeaders := map[string]string{}
	for _, h := range t.headers {
		lower := strings.ToLower(h.name)
		if other, seen := seenHeaders[lower]; seen {
			add(DuplicateHeader, "headers."+h.name, fmt.Sprintf("Headers %q and %q differ only in case", other, h.name))
		}
		seenHeaders[lower] = h.name
	}
	if timeout, ok := t.numberOption("response_timeout_secs"); ok && (timeout < 5 || timeout > 300) {
		add(InvalidOption, "response_timeout_secs", "A webhook tool's response_timeout_secs must be 5–300")
	}
	domains, _ := t.config["follow_redirects_allowed_domains"].([]any)
	if t.config["follow_redirects"] == true && len(domains) == 0 {
		add(InvalidOption, "follow_redirects", "follow_redirects needs follow_redirects_allowed_domains")
	}
	return issues
}

func (t *WebhookTool) modelIssues() []Issue {
	var issues []Issue
	if t.apiMalformed != nil {
		issues = append(issues, errorIssue(InvalidModel, "api_schema", "The loaded api_schema isn't an object. Fix it in the source JSON"))
	}
	for _, key := range t.malformed {
		// The field is written back as loaded unless the builder has since set its own.
		replaced := map[string]bool{
			"path_params_schema":   len(t.pathParams) > 0,
			"query_params_schema":  t.query != nil,
			"request_body_schema":  t.body != nil,
			"response_body_schema": t.response != nil,
			"request_headers":      len(t.headers) > 0,
		}[key]
		if !replaced {
			issues = append(issues, errorIssue(InvalidModel, key, fmt.Sprintf("The loaded api_schema.%s has the wrong shape, which the generated tool config model can't take. Fix it in the source JSON", key)))
		}
	}
	for _, p := range t.pathParams {
		checkModelShape(p.schema, "path_params."+p.name, &issues)
	}
	for _, s := range []struct {
		label  string
		schema Schema
	}{{"query", t.query}, {"body", t.body}, {"response", t.response}} {
		if s.schema != nil {
			checkModelShape(s.schema, s.label, &issues)
		}
	}
	return issues
}

func (t *WebhookTool) toWire() map[string]any {
	wire := cloneMap(t.config)
	wire["type"] = "webhook"
	if t.apiMalformed != nil {
		wire["api_schema"] = t.apiMalformed
		return wire
	}
	api := cloneMap(t.api)
	if len(t.pathParams) > 0 {
		params := make(map[string]any, len(t.pathParams))
		for _, p := range t.pathParams {
			params[p.name] = p.schema.state().toWire()
		}
		api["path_params_schema"] = params
	}
	if t.query != nil {
		// The query params schema keeps every loaded field. Its wire form has no `type`, so the one
		// the builder adds is left out unless the loaded schema declared it. Emptied by edits, it is
		// left out: a `type` kept from the loaded schema isn't content.
		delete(api, "query_params_schema")
		query := t.query.state().toWire()
		if !t.queryTypeLoaded {
			delete(query, "type")
		}
		_, typed := query["type"]
		if t.queryUntouched || len(query) > 1 || (len(query) == 1 && !typed) {
			api["query_params_schema"] = query
		}
	}
	if t.body != nil {
		api["request_body_schema"] = t.body.state().toWire()
	}
	if t.response != nil {
		api["response_body_schema"] = t.response.state().toWire()
	}
	if len(t.headers) > 0 {
		headers := make(map[string]any, len(t.headers))
		for _, h := range t.headers {
			headers[h.name] = h.value
		}
		api["request_headers"] = headers
	}
	wire["api_schema"] = api
	return wire
}

func (t *WebhookTool) hasPathParam(name string) bool {
	for _, p := range t.pathParams {
		if p.name == name {
			return true
		}
	}
	return false
}

func (t *WebhookTool) pathParamNames() []string {
	names := make([]string, len(t.pathParams))
	for i, p := range t.pathParams {
		names[i] = p.name
	}
	return names
}

func propertyNames(schema Schema) []string {
	if object, ok := schema.(ObjectSchema); ok {
		return object.PropertyNames()
	}
	return nil
}

// stripEnvTemplates replaces each environment variable template ({{system__env_<label>}}, label
// matching [a-z][a-z0-9_]*) with "env", so its braces aren't read as a path placeholder.
func stripEnvTemplates(url string) string {
	const prefix = "{{system__env_"
	var b strings.Builder
	for i := 0; i < len(url); {
		if strings.HasPrefix(url[i:], prefix) {
			start := i + len(prefix)
			if start < len(url) && url[start] >= 'a' && url[start] <= 'z' {
				end := start + 1
				for end < len(url) && (url[end] >= 'a' && url[end] <= 'z' || url[end] >= '0' && url[end] <= '9' || url[end] == '_') {
					end++
				}
				if strings.HasPrefix(url[end:], "}}") {
					b.WriteString("env")
					i = end + 2
					continue
				}
			}
		}
		b.WriteByte(url[i])
		i++
	}
	return b.String()
}

// placeholders are the {name} placeholders in a URL, in order: name is one or more characters
// other than { and }, as the API matches them.
func placeholders(url string) []string {
	var names []string
	for i := 0; i < len(url); i++ {
		if url[i] != '{' {
			continue
		}
		j := i + 1
		for j < len(url) && url[j] != '{' && url[j] != '}' {
			j++
		}
		if j < len(url) && url[j] == '}' && j > i+1 {
			names = append(names, url[i+1:j])
			i = j
		}
	}
	return names
}
