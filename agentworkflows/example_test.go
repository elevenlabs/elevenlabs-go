package agentworkflows_test

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
)

func Example() {
	wf := agentworkflows.New(nil)
	triage := wf.Converse("triage", "Find out what the caller needs", nil)
	triage.When(agentworkflows.ConditionLLM("Caller wants a human", nil), nil).
		Transfer("human", agentworkflows.DestinationPhone("+15551234567"), nil)

	s, err := wf.ToJSONString(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(s)
	// Output:
	// {
	//   "edges": {
	//     "start_node__triage": {
	//       "forward_condition": {
	//         "type": "unconditional"
	//       },
	//       "source": "start_node",
	//       "target": "triage"
	//     },
	//     "triage__human": {
	//       "forward_condition": {
	//         "condition": "Caller wants a human",
	//         "type": "llm"
	//       },
	//       "source": "triage",
	//       "target": "human"
	//     }
	//   },
	//   "nodes": {
	//     "human": {
	//       "edge_order": [],
	//       "position": {
	//         "x": 600,
	//         "y": 0
	//       },
	//       "transfer_destination": {
	//         "phone_number": "+15551234567",
	//         "type": "phone"
	//       },
	//       "type": "phone_number"
	//     },
	//     "start_node": {
	//       "edge_order": [
	//         "start_node__triage"
	//       ],
	//       "position": {
	//         "x": 0,
	//         "y": 0
	//       },
	//       "type": "start"
	//     },
	//     "triage": {
	//       "additional_prompt": "Find out what the caller needs",
	//       "edge_order": [
	//         "triage__human"
	//       ],
	//       "label": "triage",
	//       "position": {
	//         "x": 300,
	//         "y": 0
	//       },
	//       "type": "override_agent"
	//     }
	//   }
	// }
}

// Errors are recorded while chaining and reported together.
func ExampleWorkflow_Build() {
	wf := agentworkflows.New(nil)
	wf.Converse("triage", "Triage", nil)
	wf.Converse("triage", "Duplicate id", nil)
	wf.Node("triage").GoTo("does_not_exist", nil)

	_, err := wf.Build(nil)
	fmt.Println(err.(*agentworkflows.ValidationError).Codes())
	// Output: [duplicate_edge duplicate_node_id unknown_node]
}
