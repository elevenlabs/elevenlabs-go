package agenttools_test

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// load is FromJSON that fails the test on error.
func load(t *testing.T, text string) agenttools.Tool {
	t.Helper()
	tool, err := agenttools.FromJSON(text)
	if err != nil {
		t.Fatal(err)
	}
	return tool
}

func jsonString(t *testing.T, tool agenttools.Tool) string {
	t.Helper()
	out, err := tool.ToJSONString()
	if err != nil {
		t.Fatalf("ToJSONString: %v", err)
	}
	return out
}

// compact is canonical JSON text without whitespace, for short comparisons.
func compact(t *testing.T, text string) string {
	t.Helper()
	return mustString(t, value(t, json.RawMessage(text)))
}

func codes(t *testing.T, err error) []string {
	t.Helper()
	var verr *agenttools.ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("want a *ValidationError, got %v", err)
	}
	return verr.Codes()
}

func wantCodes(t *testing.T, err error, want ...string) {
	t.Helper()
	if got := codes(t, err); !reflect.DeepEqual(got, want) {
		t.Fatalf("codes = %v, want %v (%v)", got, want, err)
	}
}

func TestExample(t *testing.T) {
	getOrder := agenttools.Webhook("get_order", "Look up an order by the number the caller gives").
		Get("https://api.example.com/orders/{order_id}").
		PathParam("order_id", agenttools.String().Describe("The order number, e.g. A-1042")).
		QueryParam("include", agenttools.String().Constant("items")).
		Header("Authorization", agenttools.HeaderSecret("sec_orders_api")).
		Header("X-Caller", agenttools.HeaderVariable("system__caller_id")).
		Assign("order_status", "order.status", nil).
		Options(&agenttools.ToolOptions{
			ResponseTimeoutSecs: ptr(10),
			PreToolSpeech:       ptr(types.PreToolSpeechModeForce),
		})
	request, err := getOrder.Build(nil)
	if err != nil {
		t.Fatal(err)
	}
	if request.ToolConfig.Webhook == nil || request.ToolConfig.Webhook.ApiSchema.Url != "https://api.example.com/orders/{order_id}" {
		t.Fatalf("unexpected request: %v", request)
	}
	fx := loadFixtures(t)["01-docs-example.json"]
	if got, want := jsonString(t, getOrder), canonicalText(t, fx.Expected); got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}

	transfer := agenttools.SystemTransferToNumber(nil).
		Transfer("Caller asks for billing", agenttools.DestinationPhone("+15551234567"), nil).
		Transfer("Caller asks for a manager", agenttools.DestinationSIPVariable("manager_sip"),
			&agenttools.TransferOptions{TransferType: ptr(types.TransferTypeEnum("sip_refer"))})
	config, err := transfer.BuildConfig(nil)
	if err != nil {
		t.Fatal(err)
	}
	if config.System == nil || len(config.System.Params.TransferToNumber.Transfers) != 2 {
		t.Fatalf("want a built-in transfer_to_number entry with 2 transfers, got %v", config)
	}
}

func ptr[T any](v T) *T { return &v }

// Shared rule 1: removing the last entry leaves the field out; removing a name that doesn't exist
// does nothing; loaded empty values round-trip unchanged.
func TestRemovingTheLastEntryLeavesTheFieldOut(t *testing.T) {
	tool := agenttools.Webhook("t", "T").Get("https://x/").
		QueryParam("q", agenttools.String().Describe("Q")).
		Assign("a", "a", nil).
		RemoveQueryParam("q").
		RemoveAssignment("a")
	if got, want := jsonString(t, tool), compact(t, `{"api_schema":{"method":"GET","url":"https://x/"},"description":"T","name":"t","type":"webhook"}`); got != want {
		t.Fatalf("got %s", got)
	}

	client := agenttools.Client("c", "C").Param("p", agenttools.String().Describe("P")).RemoveParam("p")
	if got, want := jsonString(t, client), compact(t, `{"description":"C","name":"c","type":"client"}`); got != want {
		t.Fatalf("got %s", got)
	}

	loaded := `{"type":"webhook","name":"t","description":"T","assignments":[],
		"api_schema":{"method":"GET","url":"https://x/","query_params_schema":{"properties":{}}}}`
	tool = load(t, loaded).AsWebhook().RemoveAssignment("missing").RemoveQueryParam("missing")
	if got, want := jsonString(t, tool), compact(t, loaded); got != want {
		t.Fatalf("loaded empty values changed:\n%s", got)
	}

	// A leftover loaded `type` is not content; another loaded field is.
	tool = load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"GET","url":"https://x/",
		"query_params_schema":{"type":"object","properties":{"q":{"type":"string","description":"Q"}},"required":["q"]}}}`).
		AsWebhook().RemoveQueryParam("q")
	if _, has := apiSchema(t, tool)["query_params_schema"]; has {
		t.Fatal("an emptied query schema with only a loaded type must be left out")
	}
	tool = load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"GET","url":"https://x/",
		"query_params_schema":{"description":"kept","properties":{"q":{"type":"string","description":"Q"}}}}}`).
		AsWebhook().RemoveQueryParam("q")
	if got := apiSchema(t, tool)["query_params_schema"]; !reflect.DeepEqual(got, map[string]any{"description": "kept"}) {
		t.Fatalf("query_params_schema = %v", got)
	}
}

func apiSchema(t *testing.T, tool agenttools.Tool) map[string]any {
	t.Helper()
	out, err := tool.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	return out["api_schema"].(map[string]any)
}

// Shared rule 2: an empty URL is a missing one.
func TestEmptyURLIsMissing(t *testing.T) {
	_, err := agenttools.Webhook("t", "T").Get("").ToJSON()
	wantCodes(t, err, "missing_url")
}

// Shared rule 3: a loaded container schema without a type is an object; a declared non-object
// type is kept and reported.
func TestLoadedContainerSchemas(t *testing.T) {
	tool := load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"POST","url":"https://x/",
		"request_body_schema":{},"response_body_schema":{"properties":{"a":{"type":"string","description":"A"}}},
		"query_params_schema":{}}}`)
	api := apiSchema(t, tool)
	if got := api["request_body_schema"]; !reflect.DeepEqual(got, map[string]any{"type": "object"}) {
		t.Errorf("body = %v", got)
	}
	if got := api["response_body_schema"].(map[string]any)["type"]; got != "object" {
		t.Errorf("response type = %v", got)
	}
	if got := api["query_params_schema"]; !reflect.DeepEqual(got, map[string]any{}) {
		t.Errorf("query = %v, want {} (no synthetic type)", got)
	}
	client := load(t, `{"type":"client","name":"c","description":"C","parameters":{}}`)
	if got, _ := client.ToJSON(); !reflect.DeepEqual(got["parameters"], map[string]any{"type": "object"}) {
		t.Errorf("parameters = %v", got["parameters"])
	}

	bad := load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"POST","url":"https://x/",
		"request_body_schema":{"type":"string","description":"B"},"response_body_schema":{"type":"array","items":{"type":"string","description":"I"}},
		"query_params_schema":{"type":"string","description":"Q"}}}`)
	_, err := bad.ToJSON()
	wantCodes(t, err, "body_not_object", "body_not_object", "body_not_object")
	var paths []string
	for _, issue := range err.(*agenttools.ValidationError).Issues {
		paths = append(paths, issue.Path)
	}
	if want := []string{"query", "body", "response"}; !reflect.DeepEqual(paths, want) {
		t.Errorf("paths = %v, want %v", paths, want)
	}
	_, err = load(t, `{"type":"client","name":"c","description":"C","parameters":{"type":"string","description":"P"}}`).ToJSON()
	wantCodes(t, err, "body_not_object")
	if path := err.(*agenttools.ValidationError).Issues[0].Path; path != "params" {
		t.Errorf("path = %q", path)
	}
}

// Shared rule 4: a null constant on a nullable literal survives load; the generated model can't
// carry it, so Build reports constant_type_mismatch.
func TestNullConstantOnNullableLiteral(t *testing.T) {
	text := `{"type":"client","name":"c","description":"C","parameters":{"type":"object",
		"properties":{"n":{"type":["integer","null"],"constant_value":null},"s":{"type":"string","description":"S","constant_value":null}},
		"required":["n","s"]}}`
	tool := load(t, text)
	out := jsonString(t, tool)
	if !strings.Contains(out, `"constant_value": null`) {
		t.Fatalf("the null constant was dropped:\n%s", out)
	}
	if strings.Count(out, "constant_value") != 1 {
		t.Fatalf("a null constant on a non-nullable literal must read as unset:\n%s", out)
	}
	_, err := tool.Build(nil)
	wantCodes(t, err, "constant_type_mismatch")
	if path := err.(*agenttools.ValidationError).Issues[0].Path; path != "params.n" {
		t.Errorf("path = %q", path)
	}
}

// Shared rule 5: loaded query_params_schema, allowed_values and required_constraints keep their
// unmodeled fields; edits are overlaid.
func TestLoadedObjectsKeepUnmodeledFields(t *testing.T) {
	tool := load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"GET","url":"https://x/",
		"query_params_schema":{"x_note":"kept","properties":{"q":{"type":"string","description":"Q"}}}}}`).
		AsWebhook().QueryParam("r", agenttools.String().FromVariable("r"))
	query := apiSchema(t, tool)["query_params_schema"].(map[string]any)
	if query["x_note"] != "kept" || !reflect.DeepEqual(query["required"], []any{"r"}) {
		t.Errorf("query = %v", query)
	}

	schema, err := agenttools.SchemaFromJSON(`{"type":"object",
		"properties":{"a":{"type":"string","description":"A","allowed_values":{"dynamic_variable":"ids","x":1}}},
		"required_constraints":{"any_of":[{"required":["a"],"note":"n"}],"x":true}}`)
	if err != nil {
		t.Fatal(err)
	}
	object := schema.(agenttools.ObjectSchema)
	a := object.PropertySchema("a").(agenttools.LiteralSchema).AllowedFrom("other_ids")
	object = object.Property("a", a).RequireAllOf("a")
	got := mustString(t, object.ToJSON())
	want := compact(t, `{"type":"object",
		"properties":{"a":{"type":"string","description":"A","allowed_values":{"dynamic_variable":"other_ids","x":1}}},
		"required_constraints":{"any_of":[{"required":["a"],"note":"n"}],"all_of":[{"required":["a"]}],"x":true}}`)
	if got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}
}

func TestConstantsTheModelWouldChangeAreRejected(t *testing.T) {
	param := func(schema agenttools.Schema) agenttools.Tool {
		return agenttools.Client("c", "C").Param("p", schema)
	}
	cases := map[string]agenttools.Schema{
		"integer beyond int64":         agenttools.Integer().Constant(json.Number("9223372036854775808")),
		"negative beyond int64":        agenttools.Integer().Constant(json.Number("-9223372036854775809")),
		"fraction with an exponent":    agenttools.Integer().Constant(json.Number("12.5e-1")),
		"fraction beyond float64":      agenttools.Integer().Constant(json.Number("1.0000000000000000001")),
		"non-finite number":            agenttools.Number().Constant(json.Number("1e400")),
		"Go infinity":                  agenttools.Number().Constant(math.Inf(1)),
		"non-finite inside an object":  agenttools.Object().Constant(map[string]any{"x": json.Number("-1e999")}),
		"empty string":                 agenttools.String().Constant(""),
		"integer constant for a float": agenttools.Integer().Constant(1.5),
	}
	for name, schema := range cases {
		_, err := param(schema).ToJSON()
		if got := codes(t, err); !reflect.DeepEqual(got, []string{"constant_type_mismatch"}) {
			t.Errorf("%s: codes = %v", name, got)
		}
	}

	// Values the model can send exactly are kept, written as JavaScript writes them.
	tool := param(agenttools.Object().
		Property("big", agenttools.Integer().Constant(json.Number("9007199254740993"))).
		Property("exp", agenttools.Integer().Constant(json.Number("1.5e1"))).
		Property("max", agenttools.Integer().Constant(int64(math.MaxInt64))).
		Property("small", agenttools.Number().Constant(1e-7)).
		Property("large", agenttools.Number().Constant(1e21)).
		Property("whole", agenttools.Number().Constant(json.Number("300.0"))))
	out, err := tool.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	p := out["parameters"].(map[string]any)["properties"].(map[string]any)["p"]
	properties := p.(map[string]any)["properties"].(map[string]any)
	for name, want := range map[string]string{
		"big": "9007199254740993", "exp": "15", "max": "9223372036854775807",
		"small": "1e-7", "large": "1e+21", "whole": "300",
	} {
		if got := properties[name].(map[string]any)["constant_value"]; got != json.Number(want) {
			t.Errorf("%s = %v, want %s", name, got, want)
		}
	}
	if _, err := tool.Build(nil); err != nil {
		t.Errorf("Build: %v", err)
	}

	// Inside object and array constants, the safety net catches what the model would change: an
	// empty object (left out) and an integer a float64 can't hold.
	for name, schema := range map[string]agenttools.Schema{
		"empty object":      agenttools.Object().Constant(map[string]any{}),
		"empty array":       agenttools.Array(agenttools.String().Describe("S")).Constant([]any{}),
		"big nested number": agenttools.Object().Constant(map[string]any{"id": json.Number("9007199254740993")}),
	} {
		tool := param(schema)
		if _, err := tool.ToJSON(); err != nil {
			t.Errorf("%s: ToJSON: %v", name, err)
		}
		_, err := tool.Build(nil)
		if got := codes(t, err); !reflect.DeepEqual(got, []string{"constant_type_mismatch"}) {
			t.Errorf("%s: Build codes = %v", name, got)
		}
	}
	nested := param(agenttools.Object().Constant(map[string]any{"a": map[string]any{}, "b": []any{}}))
	if got := jsonString(t, nested); !strings.Contains(got, `"a": {}`) || !strings.Contains(got, `"b": []`) {
		t.Errorf("empty values inside a constant must stay as they are:\n%s", got)
	}
	if _, err := nested.Build(nil); err != nil {
		t.Errorf("Build: %v", err)
	}
}

func TestNumericPropertyNamesStayAnObject(t *testing.T) {
	schema, err := agenttools.SchemaFromJSON(`{"type":"object","properties":{
		"b":{"type":"string","description":"B"},"1":{"type":"string","description":"One"},"0":{"type":"string","description":"Zero"}}}`)
	if err != nil {
		t.Fatal(err)
	}
	object := schema.(agenttools.ObjectSchema)
	if got := object.PropertyNames(); !reflect.DeepEqual(got, []string{"0", "1", "b"}) {
		t.Errorf("names = %v, want JavaScript order", got)
	}
	object = object.
		Property("b", agenttools.String().Describe("B")).
		Property("1", agenttools.String().Describe("One")).
		Property("0", agenttools.String().Describe("Zero"))
	got := mustString(t, object.ToJSON())
	want := compact(t, `{"type":"object","properties":{"0":{"type":"string","description":"Zero"},"1":{"type":"string","description":"One"},
		"b":{"type":"string","description":"B"}},"required":["0","1","b"]}`)
	if got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}
}

func TestPathParamsAreAlwaysRequired(t *testing.T) {
	tool := agenttools.Webhook("t", "T").Get("https://x/{id}").
		PathParam("id", agenttools.String().Omit().Optional())
	_, err := tool.ToJSON()
	wantCodes(t, err, "omitted_required")
}

func TestMalformedLoadedFieldsAreKeptAndReported(t *testing.T) {
	for name, text := range map[string]string{
		"is_system_provided": `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{"p":{"type":"string","description":"P","is_system_provided":"yes"}}}}`,
		"properties":         `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{"p":42}}}`,
		"parameters":         `{"type":"client","name":"c","description":"C","parameters":42}`,
		"enum":               `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{"p":{"type":"string","description":"P","enum":"x"}}}}`,
		"items":              `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{"p":{"type":"array","items":"x"}}}}`,
		"required_constraints": `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{"p":{"type":"string","description":"P"}},
			"required_constraints":{"any_of":"p"}}}`,
		"path_params_schema":       `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"GET","url":"https://x/","path_params_schema":{"id":"x"}}}`,
		"params missing":           `{"type":"system","name":"end_call"}`,
		"params not object":        `{"type":"system","name":"end_call","params":[]}`,
		"no system_tool_type":      `{"type":"system","name":"end_call","params":{"x":1}}`,
		"unknown system_tool_type": `{"type":"system","name":"x","params":{"system_tool_type":"future_tool"}}`,
	} {
		tool := load(t, text)
		got, err := tool.ToJSONString()
		if err != nil {
			t.Errorf("%s: ToJSON: %v", name, err)
			continue
		}
		if got != compact(t, text) {
			t.Errorf("%s: not kept as loaded:\n%s", name, got)
		}
		_, err = tool.Build(nil)
		if codes := codes(t, err); len(codes) == 0 || codes[0] != "invalid_model" {
			t.Errorf("%s: Build codes = %v", name, codes)
		}
	}
}

func TestPropertyNamedLikeAField(t *testing.T) {
	text := `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{
		"constant_value":{"type":"string","description":"A param named constant_value","dynamic_variable":null},
		"properties":{"type":"integer","constant_value":3}},"required":["constant_value","properties"]}}`
	tool := load(t, text)
	got := jsonString(t, tool)
	want := compact(t, `{"type":"client","name":"c","description":"C","parameters":{"type":"object","properties":{
		"constant_value":{"type":"string","description":"A param named constant_value"},
		"properties":{"type":"integer","constant_value":3}},"required":["constant_value","properties"]}}`)
	if got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}
	if _, err := tool.Build(nil); err != nil {
		t.Fatalf("Build: %v", err)
	}
}

func TestSchemasAreValues(t *testing.T) {
	base := agenttools.Object().Property("a", agenttools.String().Describe("A"))
	withB := base.Property("b", agenttools.String().Describe("B"))
	without := withB.Without("a")
	if !reflect.DeepEqual(base.PropertyNames(), []string{"a"}) ||
		!reflect.DeepEqual(withB.PropertyNames(), []string{"a", "b"}) ||
		!reflect.DeepEqual(without.PropertyNames(), []string{"b"}) {
		t.Fatal("a schema method changed the schema it was called on")
	}
	s := agenttools.String()
	_ = s.Describe("x")
	if _, has := s.ToJSON()["description"]; has {
		t.Fatal("Describe changed the schema it was called on")
	}
}

func TestOptionsOfTheWrongKind(t *testing.T) {
	_, err := agenttools.SystemEndCall().Options(&agenttools.ToolOptions{ExpectsResponse: ptr(true)}).ToJSON()
	wantCodes(t, err, "invalid_option")
	_, err = agenttools.SystemTransferToAgent().
		Transfer("x", agenttools.DestinationAgent("agent_1", nil), &agenttools.TransferOptions{TransferType: ptr(types.TransferTypeEnum("conference"))}).
		ToJSON()
	wantCodes(t, err, "invalid_option")
	_, err = agenttools.SystemEndCall().Transfer("x", agenttools.DestinationPhone("+1"), nil).ToJSON()
	wantCodes(t, err, "invalid_option")
	_, err = agenttools.SystemTransferToNumber(nil).Transfer("x", agenttools.Destination{}, nil).ToJSON()
	wantCodes(t, err, "invalid_destination", "no_transfers")
}

func TestFromResponse(t *testing.T) {
	fx := loadFixtures(t)["09-round-trip.json"]
	var response types.ToolResponseModel
	data, err := json.Marshal(map[string]any{
		"id": "tool_1", "tool_config": json.RawMessage(fx.From),
		"access_info": map[string]any{"is_creator": true, "creator_name": "x", "creator_email": "x@example.com", "role": "admin"},
		"usage_stats": map[string]any{"avg_latency_secs": 0},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatal(err)
	}
	tool, err := agenttools.FromResponse(&response)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := jsonString(t, tool), canonicalText(t, fx.Expected); got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}
	if _, err := agenttools.FromResponse(42); err == nil {
		t.Fatal("want an error for a value that isn't a response")
	}
}

func TestLoadersRejectNonObjects(t *testing.T) {
	for _, data := range []any{`[]`, `"x"`, `{`, 42} {
		if _, err := agenttools.FromJSON(data); err == nil {
			t.Errorf("FromJSON(%v): want an error", data)
		}
		if _, err := agenttools.SchemaFromJSON(data); err == nil {
			t.Errorf("SchemaFromJSON(%v): want an error", data)
		}
	}
	// A request wrapper is unwrapped.
	tool := load(t, `{"tool_config":{"type":"system","name":"end_call","params":{"system_tool_type":"end_call"}}}`)
	if tool.Kind() != agenttools.ToolKindSystem || tool.AsSystem().SystemToolType() != "end_call" {
		t.Fatalf("got %v", tool.Kind())
	}
}

// A loaded empty query_params_schema is written as {} by ToJSON; Build must not send it as
// "properties": null, which the generated model's missing omitempty would otherwise produce.
func TestLoadedEmptyQuerySchemaBuildsWithEmptyProperties(t *testing.T) {
	tool := load(t, `{"type":"webhook","name":"t","description":"T","api_schema":{"method":"GET","url":"https://x/",
		"query_params_schema":{}}}`)
	if got := apiSchema(t, tool)["query_params_schema"]; !reflect.DeepEqual(got, map[string]any{}) {
		t.Fatalf("ToJSON query = %v, want {}", got)
	}
	request, err := tool.Build(nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(request.ToolConfig)
	if err != nil {
		t.Fatal(err)
	}
	var sent struct {
		APISchema struct {
			Query map[string]json.RawMessage `json:"query_params_schema"`
		} `json:"api_schema"`
	}
	if err := json.Unmarshal(data, &sent); err != nil {
		t.Fatal(err)
	}
	if got := string(sent.APISchema.Query["properties"]); got != "{}" {
		t.Fatalf("Build sent query_params_schema.properties = %s, want {} (%s)", got, data)
	}
}

const transferRule = `{"condition":"Caller asks for the front desk","provider_configs":[{"config":{"transfer_destination":{"destination":"+15551234567","type":"phone"},"type":"twilio_cold_blind_transfer"},"type":"twilio"}]}`

func loadedTransfer(t *testing.T, systemToolType string) agenttools.Tool {
	t.Helper()
	return load(t, `{"type":"system","name":"`+systemToolType+`","params":{"system_tool_type":"`+systemToolType+
		`","transfers":[],"transfer_rules":[`+transferRule+`]}}`)
}

func TestTransferRulesDoNotSatisfyAnAgentTransfer(t *testing.T) {
	_, err := loadedTransfer(t, "transfer_to_agent").Build(nil)
	wantCodes(t, err, "no_transfers")
}

func TestBuildSendsLoadedTransferRules(t *testing.T) {
	request, err := loadedTransfer(t, "transfer_to_number").Build(nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	var sent struct {
		ToolConfig struct {
			Params struct {
				TransferRules []any `json:"transfer_rules"`
			} `json:"params"`
		} `json:"tool_config"`
	}
	if err := json.Unmarshal(data, &sent); err != nil {
		t.Fatal(err)
	}
	var want []any
	if err := json.Unmarshal([]byte(`[`+transferRule+`]`), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(sent.ToolConfig.Params.TransferRules, want) {
		t.Fatalf("sent transfer_rules = %v, want %v\n%s", sent.ToolConfig.Params.TransferRules, want, data)
	}
}
