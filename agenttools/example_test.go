package agenttools_test

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/types"
)

func Example() {
	refund := agenttools.Webhook("issue_refund", "Refund an order the caller asked to return").
		Post("https://api.example.com/refunds").
		Body(agenttools.Object().
			Property("order_id", agenttools.String().FromVariable("order_id")).
			Property("reason", agenttools.String().Describe("Why the caller wants a refund").Choices("damaged", "late", "other")).
			Property("note", agenttools.String().Describe("Anything else the caller said").Optional()))

	s, err := refund.ToJSONString()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(s)
	// Output:
	// {
	//   "api_schema": {
	//     "method": "POST",
	//     "request_body_schema": {
	//       "properties": {
	//         "note": {
	//           "description": "Anything else the caller said",
	//           "type": "string"
	//         },
	//         "order_id": {
	//           "dynamic_variable": "order_id",
	//           "type": "string"
	//         },
	//         "reason": {
	//           "description": "Why the caller wants a refund",
	//           "enum": [
	//             "damaged",
	//             "late",
	//             "other"
	//           ],
	//           "type": "string"
	//         }
	//       },
	//       "required": [
	//         "order_id",
	//         "reason"
	//       ],
	//       "type": "object"
	//     },
	//     "url": "https://api.example.com/refunds"
	//   },
	//   "description": "Refund an order the caller asked to return",
	//   "name": "issue_refund",
	//   "type": "webhook"
	// }
}

func ExampleSystemTransferToNumber() {
	transfer := agenttools.SystemTransferToNumber(nil).
		Transfer("Caller asks for billing", agenttools.DestinationPhone("+15551234567"), nil).
		Transfer("Caller asks for a manager", agenttools.DestinationSIPVariable("manager_sip"),
			&agenttools.TransferOptions{TransferType: types.TransferTypeEnum("sip_refer").Ptr()})

	// The entry for the agent's prompt.built_in_tools.transfer_to_number.
	config, err := transfer.BuildConfig(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(config.System.Params.TransferToNumber.Transfers))
	// Output: 2
}

func ExampleFromJSON() {
	tool, err := agenttools.FromJSON(`{
	  "type": "client",
	  "name": "show_store_map",
	  "description": "Show the caller a map of the nearest store",
	  "parameters": {"type": "object", "properties": {"store_id": {"type": "string", "description": "The store to show"}}, "required": ["store_id"]}
	}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := tool.AsClient().
		Param("zoom", agenttools.Integer().Describe("Zoom level, 1-20").Optional()).
		Options(&agenttools.ToolOptions{ExpectsResponse: new(bool)})
	params, _ := client.ParametersSchema()
	fmt.Println(params.PropertyNames())
	_, err = client.ToJSONString()
	fmt.Println(err)
	// Output:
	// [store_id zoom]
	// <nil>
}

func ExampleValidationError() {
	_, err := agenttools.Webhook("broken", "Broken").
		Post("https://api.example.com/{id}").
		ToJSONString()
	fmt.Println(err.(*agenttools.ValidationError).Codes())
	// Output: [missing_body missing_path_param]
}
