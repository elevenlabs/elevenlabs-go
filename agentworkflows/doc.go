// Package agentworkflows writes ElevenLabs agent workflows as code.
//
// It is a hand-written layer on top of the generated client (every file is listed in
// .fernignore). Build returns the generated *types.AgentWorkflowRequestModel that
// client.Agents.Create and client.Agents.Update accept in their Workflow field, and ToJSONString
// returns canonical API wire-format JSON for storing in git, byte-identical to the output of every
// other ElevenLabs SDK for the same program. The vocabulary and output rules are shared by all SDKs and checked by
// the same conformance fixtures (docs/agents-as-code/workflow-spec.md in elevenlabs-js).
//
// A workflow starts with New (or FromJSON / FromAgent to edit an existing one). The workflow acts
// as the start node. Each verb (Converse, Tool, Transfer, End) adds a node, connects it to the
// current point and returns its *Node, so linear flows chain. When, OnSuccess and OnFailure open a
// *Branch with a condition; GoTo connects to an existing node, for loops and merges.
//
//	wf := agentworkflows.New(nil)
//	triage := wf.Converse("triage", "Find out what the caller needs", nil)
//	triage.When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
//		Tool("lookup_order", []string{"tool_get_order"}, nil).
//		OnResult(agentworkflows.ResultBranches{
//			Success: func(ok *agentworkflows.Branch) {
//				ok.Converse("resolve", "Explain the order status", nil).End(nil)
//			},
//			Failure: func(failed *agentworkflows.Branch) { failed.GoTo("triage", nil) },
//		})
//	triage.When(agentworkflows.ConditionLLM("Caller wants a human", nil), nil).
//		Transfer("human", agentworkflows.DestinationPhone("+15551234567"), nil)
//
//	workflow, err := wf.Build(nil) // *types.AgentWorkflowRequestModel
//
// Chaining never fails. Every problem (duplicate ids, unknown targets, options that don't apply
// to a node's kind, invalid routing) is recorded and reported at once by Build, ToJSON and
// ToJSONString as a *ValidationError.
//
// Node types the builder has no verb for (such as say, or types newer than this SDK version) are
// kept exactly as loaded. A type the generated types don't model gets an unknown_node_type
// warning from Build, since agents.Create can't send it; send ToJSON's output directly instead.
package agentworkflows
