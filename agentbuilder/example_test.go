package agentbuilder_test

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/agentbuilder"
	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
)

func Example() {
	getOrder := agenttools.Webhook("get_order", "Look up an order by its number").
		Get("https://api.example.com/orders/{order_id}").
		PathParam("order_id", agenttools.String().Describe("The order number"))

	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Find out what the caller needs", nil).
		When(agentworkflows.ConditionLLM("Caller gave an order number", nil), nil).
		Tool("lookup", nil, &agentworkflows.Options{Tools: []agentworkflows.Tool{getOrder}})

	prompt := "You are a support agent for Acme."
	agent := agentbuilder.New("Support line", &agentbuilder.Options{
		Prompt:   &prompt,
		Tools:    []agenttools.Tool{agenttools.SystemEndCall()},
		Workflow: wf,
	})

	// Every check that needs no ids, without an API key.
	if err := agent.Validate(nil); err != nil {
		fmt.Println(err)
		return
	}
	// The tools client.Agents.Push creates or updates.
	for _, tool := range agent.Tools() {
		fmt.Println(tool.Name())
	}
	// The config Push sends, once it knows the tools' ids.
	config, err := agent.ToJSON(&agentbuilder.BuildOptions{ToolIDs: map[string]string{"get_order": "tool_123"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	nodes := config["workflow"].(map[string]any)["nodes"].(map[string]any)
	fmt.Println(nodes["lookup"].(map[string]any)["tools"])
	// Output:
	// get_order
	// [map[tool_id:tool_123]]
}
