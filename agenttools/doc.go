// Package agenttools writes ElevenLabs agent tools as code: webhook, client and system tools.
//
// It is a hand-written layer on top of the generated client (every file is listed in
// .fernignore). Build returns the generated *types.ToolRequestModel that
// client.Agents.Tools.Create and Update accept, and ToJSONString returns canonical API
// wire-format JSON (the CLI's tool_configs/*.json format) for storing in git, byte-identical to
// the output of every other ElevenLabs SDK for the same program. The vocabulary and output rules
// are shared by all SDKs and checked by the same conformance fixtures
// (docs/agents-as-code/tool-spec.md in elevenlabs-js).
//
//	getOrder := agenttools.Webhook("get_order", "Look up an order by the number the caller gives").
//		Get("https://api.example.com/orders/{order_id}").
//		PathParam("order_id", agenttools.String().Describe("The order number, e.g. A-1042")).
//		QueryParam("include", agenttools.String().Constant("items")).
//		Header("Authorization", agenttools.HeaderSecret("sec_orders_api")).
//		Assign("order_status", "order.status", nil)
//
//	request, err := getOrder.Build(nil) // *types.ToolRequestModel for tools.create
//	text, err := getOrder.ToJSONString() // for tool_configs/get_order.json
//
// A tool starts with Webhook, Client or a System* constructor (SystemEndCall,
// SystemTransferToNumber and so on), or with FromJSON / FromResponse to edit an existing one.
// Params and bodies are Schema values built with String, Integer, Number, Bool, Object and Array.
// Transfer destinations are the workflow builder's (agentworkflows.Destination), re-exported here.
//
// Chaining never fails. Every problem (a missing URL, a literal without a value source, an option
// that doesn't apply to the tool's kind) is recorded and reported at once by Build, BuildConfig,
// ToJSON and ToJSONString as a *ValidationError.
//
// Loaded fields the builder doesn't model are kept exactly as loaded, and so are tools of a type
// it has no constructor for (such as mcp), which load as *OtherTool. Build reports invalid_model
// for loaded input the generated request model can't hold, and constant_type_mismatch for a
// constant it would send changed; ToJSON's output can be sent directly instead.
package agenttools
