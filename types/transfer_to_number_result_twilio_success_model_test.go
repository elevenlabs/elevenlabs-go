package types

import (
	"testing"
)

func TestTransferToNumberResultTwilioSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_twilio_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "client_message": "value",
  "agent_message": "value",
  "conference_name": "value",
  "post_dial_digits": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultTwilioSuccessModel))
}

func TestTransferToNumberResultTwilioSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_twilio_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "client_message": "value",
  "agent_message": "value",
  "conference_name": "value",
  "post_dial_digits": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultTwilioSuccessModel))
}
