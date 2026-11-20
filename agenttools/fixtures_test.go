package agenttools_test

// Cross-SDK conformance: runs every shared tool fixture (vendored from elevenlabs-js
// docs/agents-as-code/fixtures/tools into testdata/fixtures) through the Go builder and compares
// ToJSONString() byte for byte, or the sorted error codes. Mirrors the JS runner in
// tests/custom/agents/tools/fixtureRunner.ts; see the tool spec, "Fixtures".

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

const fixtureCount = 23

type fixture struct {
	Name string `json:"name"`
	Tool *struct {
		Webhook       []string        `json:"webhook"`
		Client        []string        `json:"client"`
		System        string          `json:"system"`
		SystemOptions json.RawMessage `json:"systemOptions"`
	} `json:"tool"`
	From     json.RawMessage `json:"from"`
	Schema   []fixtureOp     `json:"schema"`
	Ops      []fixtureOp     `json:"ops"`
	Expected json.RawMessage `json:"expected"`
	Errors   []string        `json:"errors"`
}

type fixtureOp struct {
	Op   string            `json:"op"`
	Args []json.RawMessage `json:"args"`
}

func loadFixtures(t *testing.T) map[string]*fixture {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join("testdata", "fixtures", "*.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != fixtureCount {
		t.Fatalf("expected %d fixtures, found %d", fixtureCount, len(paths))
	}
	out := map[string]*fixture{}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var fx fixture
		if err := json.Unmarshal(data, &fx); err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		out[filepath.Base(path)] = &fx
	}
	return out
}

func sortedNames(fixtures map[string]*fixture) []string {
	names := make([]string, 0, len(fixtures))
	for name := range fixtures {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func TestFixtures(t *testing.T) {
	fixtures := loadFixtures(t)
	for _, file := range sortedNames(fixtures) {
		fx := fixtures[file]
		t.Run(file, func(t *testing.T) {
			got, codes := runFixture(t, fx)
			switch {
			case len(fx.Expected) > 0:
				want := canonicalText(t, fx.Expected)
				if codes != nil {
					t.Fatalf("%s: expected JSON but build failed: %v", fx.Name, codes)
				}
				if got != want {
					t.Fatalf("%s: JSON mismatch\n--- got ---\n%s\n--- want ---\n%s", fx.Name, got, want)
				}
			case fx.Errors != nil:
				want := append([]string{}, fx.Errors...)
				sort.Strings(want)
				if codes == nil {
					t.Fatalf("%s: expected errors %v but build succeeded:\n%s", fx.Name, want, got)
				}
				if !reflect.DeepEqual(codes, want) {
					t.Fatalf("%s: error mismatch\n  got:  %v\n  want: %v", fx.Name, codes, want)
				}
			default:
				t.Fatalf("%s: fixture has neither expected nor errors", fx.Name)
			}
		})
	}
}

// TestFixtureOutputsRoundTrip: every expected output loads back and writes byte-identically.
func TestFixtureOutputsRoundTrip(t *testing.T) {
	fixtures := loadFixtures(t)
	for _, file := range sortedNames(fixtures) {
		fx := fixtures[file]
		if len(fx.Expected) == 0 {
			continue
		}
		want := canonicalText(t, fx.Expected)
		var got string
		if fx.Schema != nil {
			schema, err := agenttools.SchemaFromJSON(fx.Expected)
			if err != nil {
				t.Fatal(err)
			}
			got = mustString(t, schema.ToJSON())
		} else {
			tool, err := agenttools.FromJSON(fx.Expected)
			if err != nil {
				t.Fatal(err)
			}
			if got, err = tool.ToJSONString(); err != nil {
				t.Fatalf("%s: %v", file, err)
			}
		}
		if got != want {
			t.Errorf("%s: round trip changed the output\n--- got ---\n%s\n--- want ---\n%s", file, got, want)
		}
	}
}

// TestFixtureOutputsBuild: every tool fixture that writes JSON also builds into the generated request
// model, and the model's JSON is ToJSON's output, except for empty lists, which the model leaves out
// (the API reads them as unset). The mcp tool is the exception: the generated request union has no
// variant for it.
func TestFixtureOutputsBuild(t *testing.T) {
	fixtures := loadFixtures(t)
	for _, file := range sortedNames(fixtures) {
		fx := fixtures[file]
		if len(fx.Expected) == 0 || fx.Schema != nil {
			continue
		}
		tool, err := agenttools.FromJSON(fx.Expected)
		if err != nil {
			t.Fatal(err)
		}
		var warnings []string
		request, err := tool.Build(&agenttools.BuildOptions{OnWarning: func(issue agenttools.Issue) {
			warnings = append(warnings, string(issue.Code))
		}})
		if tool.Kind() == agenttools.ToolKindOther {
			var verr *agenttools.ValidationError
			if !errors.As(err, &verr) || !reflect.DeepEqual(verr.Codes(), []string{"invalid_model"}) {
				t.Errorf("%s: want invalid_model for an mcp tool, got %v", file, err)
			}
			if !reflect.DeepEqual(warnings, []string{"unknown_tool_type"}) {
				t.Errorf("%s: want an unknown_tool_type warning, got %v", file, warnings)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: Build: %v", file, err)
			continue
		}
		data, err := json.Marshal(request.ToolConfig)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := withoutEmpty(t, data), withoutEmpty(t, fx.Expected); got != want {
			t.Errorf("%s: the built model's JSON differs from ToJSON\n--- model ---\n%s\n--- ToJSON ---\n%s", file, got, want)
		}
	}
}

func runFixture(t *testing.T, fx *fixture) (string, []string) {
	t.Helper()
	if fx.Schema != nil {
		return mustString(t, buildSchema(t, fx.Schema).ToJSON()), nil
	}
	var tool agenttools.Tool
	switch {
	case len(fx.From) > 0:
		loaded, err := agenttools.FromJSON(fx.From)
		if err != nil {
			t.Fatal(err)
		}
		tool = loaded
	case fx.Tool != nil && fx.Tool.Webhook != nil:
		tool = agenttools.Webhook(fx.Tool.Webhook[0], fx.Tool.Webhook[1])
	case fx.Tool != nil && fx.Tool.Client != nil:
		tool = agenttools.Client(fx.Tool.Client[0], fx.Tool.Client[1])
	case fx.Tool != nil:
		tool = systemTool(t, fx.Tool.System, fx.Tool.SystemOptions)
	default:
		t.Fatalf("%s: fixture has no tool", fx.Name)
	}
	for _, op := range fx.Ops {
		apply(t, tool, op)
	}
	out, err := tool.ToJSONString()
	if err != nil {
		var verr *agenttools.ValidationError
		if !errors.As(err, &verr) {
			t.Fatal(err)
		}
		return "", verr.Codes()
	}
	return out, nil
}

func apply(t *testing.T, tool agenttools.Tool, op fixtureOp) {
	t.Helper()
	if op.Op == "options" {
		var opts agenttools.ToolOptions
		decodeOptions(t, arg(op, 0), &opts)
		switch v := tool.(type) {
		case *agenttools.WebhookTool:
			v.Options(&opts)
		case *agenttools.ClientTool:
			v.Options(&opts)
		case *agenttools.SystemTool:
			v.Options(&opts)
		case *agenttools.OtherTool:
			v.Options(&opts)
		}
		return
	}
	switch v := tool.(type) {
	case *agenttools.WebhookTool:
		switch op.Op {
		case "get":
			v.Get(str(t, op, 0))
		case "post":
			v.Post(str(t, op, 0))
		case "put":
			v.Put(str(t, op, 0))
		case "patch":
			v.Patch(str(t, op, 0))
		case "delete":
			v.Delete(str(t, op, 0))
		case "pathParam":
			v.PathParam(str(t, op, 0), schemaArg(t, arg(op, 1)))
		case "queryParam":
			v.QueryParam(str(t, op, 0), schemaArg(t, arg(op, 1)))
		case "body":
			v.Body(schemaArg(t, arg(op, 0)))
		case "response":
			v.Response(schemaArg(t, arg(op, 0)))
		case "header":
			v.Header(str(t, op, 0), header(t, arg(op, 1)))
		case "auth":
			v.Auth(auth(t, arg(op, 0)))
		case "assign":
			var opts agenttools.AssignOptions
			decodeOptions(t, arg(op, 2), &opts)
			v.Assign(str(t, op, 0), str(t, op, 1), &opts)
		case "removePathParam":
			v.RemovePathParam(str(t, op, 0))
		case "removeQueryParam":
			v.RemoveQueryParam(str(t, op, 0))
		case "removeHeader":
			v.RemoveHeader(str(t, op, 0))
		case "removeAssignment":
			v.RemoveAssignment(str(t, op, 0))
		default:
			t.Fatalf("unknown webhook op %q", op.Op)
		}
	case *agenttools.ClientTool:
		switch op.Op {
		case "param":
			v.Param(str(t, op, 0), schemaArg(t, arg(op, 1)))
		case "removeParam":
			v.RemoveParam(str(t, op, 0))
		default:
			t.Fatalf("unknown client op %q", op.Op)
		}
	case *agenttools.SystemTool:
		if op.Op != "transfer" {
			t.Fatalf("unknown system op %q", op.Op)
		}
		var opts agenttools.TransferOptions
		decodeOptions(t, arg(op, 2), &opts)
		v.Transfer(str(t, op, 0), destination(t, arg(op, 1)), &opts)
	default:
		t.Fatalf("op %q on a tool of an unknown type", op.Op)
	}
}

func systemTool(t *testing.T, systemToolType string, raw json.RawMessage) *agenttools.SystemTool {
	t.Helper()
	switch systemToolType {
	case "end_call":
		return agenttools.SystemEndCall()
	case "skip_turn":
		var opts agenttools.SkipTurnOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemSkipTurn(&opts)
	case "language_detection":
		var opts agenttools.LanguageDetectionOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemLanguageDetection(&opts)
	case "play_keypad_touch_tone":
		var opts agenttools.PlayKeypadTouchToneOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemPlayKeypadTouchTone(&opts)
	case "voicemail_detection":
		var opts agenttools.VoicemailDetectionOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemVoicemailDetection(&opts)
	case "knowledge_base":
		var opts agenttools.KnowledgeBaseOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemKnowledgeBase(&opts)
	case "transfer_to_number":
		var opts agenttools.TransferToNumberOptions
		decodeOptions(t, raw, &opts)
		return agenttools.SystemTransferToNumber(&opts)
	case "transfer_to_agent":
		return agenttools.SystemTransferToAgent()
	}
	t.Fatalf("unknown system tool %q", systemToolType)
	return nil
}

// buildSchema runs schema ops: the first is a constructor, each later one a method called on the
// previous result.
func buildSchema(t *testing.T, ops []fixtureOp) agenttools.Schema {
	t.Helper()
	first := ops[0]
	var schema agenttools.Schema
	switch first.Op {
	case "string":
		schema = agenttools.String()
	case "integer":
		schema = agenttools.Integer()
	case "number":
		schema = agenttools.Number()
	case "bool":
		schema = agenttools.Bool()
	case "object":
		schema = agenttools.Object()
	case "array":
		schema = agenttools.Array(schemaArg(t, arg(first, 0)))
	case "fromJSON":
		loaded, err := agenttools.SchemaFromJSON(arg(first, 0))
		if err != nil {
			t.Fatal(err)
		}
		schema = loaded
	default:
		t.Fatalf("unknown schema constructor %q", first.Op)
	}
	for _, op := range ops[1:] {
		schema = schemaOp(t, schema, op)
	}
	return schema
}

func schemaOp(t *testing.T, schema agenttools.Schema, op fixtureOp) agenttools.Schema {
	t.Helper()
	switch s := schema.(type) {
	case agenttools.LiteralSchema:
		switch op.Op {
		case "describe":
			return s.Describe(str(t, op, 0))
		case "fromVariable":
			return s.FromVariable(str(t, op, 0))
		case "constant":
			return s.Constant(value(t, arg(op, 0)))
		case "omit":
			return s.Omit()
		case "optional":
			return s.Optional()
		case "systemProvided":
			return s.SystemProvided()
		case "choices":
			return s.Choices(strs(t, arg(op, 0))...)
		case "allowedFrom":
			return s.AllowedFrom(str(t, op, 0))
		}
	case agenttools.ObjectSchema:
		switch op.Op {
		case "describe":
			return s.Describe(str(t, op, 0))
		case "fromVariable":
			return s.FromVariable(str(t, op, 0))
		case "constant":
			return s.Constant(value(t, arg(op, 0)))
		case "omit":
			return s.Omit()
		case "optional":
			return s.Optional()
		case "property":
			return s.Property(str(t, op, 0), schemaArg(t, arg(op, 1)))
		case "without":
			return s.Without(str(t, op, 0))
		case "requireAnyOf":
			return s.RequireAnyOf(strs(t, arg(op, 0))...)
		case "requireAllOf":
			return s.RequireAllOf(strs(t, arg(op, 0))...)
		}
	case agenttools.ArraySchema:
		switch op.Op {
		case "describe":
			return s.Describe(str(t, op, 0))
		case "fromVariable":
			return s.FromVariable(str(t, op, 0))
		case "constant":
			return s.Constant(value(t, arg(op, 0)))
		case "omit":
			return s.Omit()
		case "optional":
			return s.Optional()
		}
	}
	t.Fatalf("schema op %q doesn't apply to a %s schema", op.Op, schema.Kind())
	return nil
}

// schemaArg is a schema in wire format, or {"build": [ops]} for one built with Schema methods
// (needed for Optional, which has no wire form of its own).
func schemaArg(t *testing.T, raw json.RawMessage) agenttools.Schema {
	t.Helper()
	var built struct {
		Build []fixtureOp `json:"build"`
	}
	if err := json.Unmarshal(raw, &built); err == nil && built.Build != nil {
		return buildSchema(t, built.Build)
	}
	schema, err := agenttools.SchemaFromJSON(raw)
	if err != nil {
		t.Fatal(err)
	}
	return schema
}

func header(t *testing.T, raw json.RawMessage) agenttools.Header {
	t.Helper()
	var h struct {
		Value    *string `json:"value"`
		Secret   *string `json:"secret"`
		Variable *string `json:"variable"`
		EnvVar   string  `json:"envVar"`
	}
	if err := json.Unmarshal(raw, &h); err != nil {
		t.Fatal(err)
	}
	switch {
	case h.Value != nil:
		return agenttools.HeaderValue(*h.Value)
	case h.Secret != nil:
		return agenttools.HeaderSecret(*h.Secret)
	case h.Variable != nil:
		return agenttools.HeaderVariable(*h.Variable)
	}
	return agenttools.HeaderEnvVar(h.EnvVar)
}

func auth(t *testing.T, raw json.RawMessage) agenttools.AuthConnection {
	t.Helper()
	var a struct {
		ID     *string `json:"id"`
		EnvVar string  `json:"envVar"`
	}
	if err := json.Unmarshal(raw, &a); err != nil {
		t.Fatal(err)
	}
	if a.ID != nil {
		return agenttools.AuthConnectionID(*a.ID)
	}
	return agenttools.AuthConnectionEnvVar(a.EnvVar)
}

func destination(t *testing.T, raw json.RawMessage) agenttools.Destination {
	t.Helper()
	var d struct {
		Phone         *string `json:"phone"`
		PhoneVariable *string `json:"phoneVariable"`
		SIP           *string `json:"sip"`
		SIPVariable   *string `json:"sipVariable"`
		Agent         string  `json:"agent"`
		NodeID        string  `json:"nodeId"`
	}
	if err := json.Unmarshal(raw, &d); err != nil {
		t.Fatal(err)
	}
	switch {
	case d.Phone != nil:
		return agenttools.DestinationPhone(*d.Phone)
	case d.PhoneVariable != nil:
		return agenttools.DestinationPhoneVariable(*d.PhoneVariable)
	case d.SIP != nil:
		return agenttools.DestinationSIP(*d.SIP)
	case d.SIPVariable != nil:
		return agenttools.DestinationSIPVariable(*d.SIPVariable)
	}
	return agenttools.DestinationAgent(d.Agent, &agenttools.AgentDestinationOptions{NodeID: d.NodeID})
}

// decodeOptions converts a fixture's camelCase option keys to snake_case, the options structs'
// json tags, and decodes them. An option the struct doesn't model fails loudly instead of
// vanishing. A missing argument leaves the struct empty.
func decodeOptions(t *testing.T, raw json.RawMessage, into any) {
	t.Helper()
	if len(raw) == 0 {
		return
	}
	var camel map[string]json.RawMessage
	if err := json.Unmarshal(raw, &camel); err != nil {
		t.Fatal(err)
	}
	snake := map[string]json.RawMessage{}
	for key, value := range camel {
		snake[camelToSnake(key)] = value
	}
	data, err := json.Marshal(snake)
	if err != nil {
		t.Fatal(err)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(into); err != nil {
		t.Fatal(err)
	}
}

func camelToSnake(key string) string {
	var b strings.Builder
	for _, r := range key {
		if unicode.IsUpper(r) {
			b.WriteByte('_')
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func arg(op fixtureOp, i int) json.RawMessage {
	if i >= len(op.Args) {
		return nil
	}
	return op.Args[i]
}

func str(t *testing.T, op fixtureOp, i int) string {
	t.Helper()
	var s string
	if err := json.Unmarshal(arg(op, i), &s); err != nil {
		t.Fatalf("%s: argument %d must be a string: %v", op.Op, i, err)
	}
	return s
}

func strs(t *testing.T, raw json.RawMessage) []string {
	t.Helper()
	var list []string
	if err := json.Unmarshal(raw, &list); err != nil {
		t.Fatal(err)
	}
	return list
}

// value is a JSON argument as a wire value (numbers keep their text).
func value(t *testing.T, raw json.RawMessage) any {
	t.Helper()
	v, err := wirejson.Decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

// withoutEmpty is canonicalText with object fields holding [] left out.
func withoutEmpty(t *testing.T, raw []byte) string {
	t.Helper()
	v, err := wirejson.Decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	var drop func(any) any
	drop = func(v any) any {
		switch v := v.(type) {
		case map[string]any:
			for key, item := range v {
				if list, ok := item.([]any); ok && len(list) == 0 {
					delete(v, key)
					continue
				}
				v[key] = drop(item)
			}
		case []any:
			for i, item := range v {
				v[i] = drop(item)
			}
		}
		return v
	}
	return mustString(t, drop(v))
}

func canonicalText(t *testing.T, raw []byte) string {
	t.Helper()
	v, err := wirejson.Decode(raw)
	if err != nil {
		t.Fatal(err)
	}
	return mustString(t, v)
}

func mustString(t *testing.T, v any) string {
	t.Helper()
	s, err := wirejson.String(v)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
