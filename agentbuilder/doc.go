// Package agentbuilder writes ElevenLabs agents as code: an agent's settings, its tools and its
// workflow in one definition, pushed with client.Agents.Push and pulled with client.Agents.Pull,
// and stored in the ElevenLabs CLI's project layout with Project.
//
// It is a hand-written layer on top of the generated client (every file is listed in
// .fernignore), after the workflow builder (agentworkflows) and the tools builder (agenttools).
// The vocabulary, the config it assembles, the steps of push and pull and the project format are
// shared by every ElevenLabs SDK and checked by the same fixtures (docs/agents-as-code/agent-spec.md
// in elevenlabs-js).
//
//	getOrder := agenttools.Webhook("get_order", "Look up an order by its number").
//		Get("https://api.example.com/orders/{order_id}").
//		PathParam("order_id", agenttools.String().Describe("The order number"))
//
//	wf := agentworkflows.New(nil)
//	wf.Converse("triage", "Find out what the caller needs", nil).
//		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
//		Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{getOrder}})
//
//	agent := agentbuilder.New("Support line", &agentbuilder.Options{
//		Prompt:       ptr("You are a support agent for Acme."),
//		FirstMessage: ptr("Hi, how can I help?"),
//		Tools:        []agenttools.Tool{agenttools.SystemEndCall()},
//		Workflow:     wf,
//	})
//	result, err := client.Agents.Push(ctx, agent, nil) // creates get_order, then the agent
//
// Push creates or updates every webhook and client tool the agent and its workflow use, then the
// agent with their ids filled in. Pass PushOptions.AgentID to update an existing agent, or a
// Project, which records the ids in agents.json and tools.json so the next push updates the same
// agent and tools. Pull loads an existing agent, with its tools as agenttools tools.
//
// An Agent holds no client: Validate runs every check that needs no ids, so CI can check an agent
// without an API key, and one Agent can be pushed to several workspaces. Chaining never fails;
// every problem is reported at once as a *ValidationError.
//
// Loaded settings the SDK doesn't model are kept and sent as they are: Push sends ToJSON's
// API-format config rather than the generated request model, which can't hold them.
package agentbuilder
