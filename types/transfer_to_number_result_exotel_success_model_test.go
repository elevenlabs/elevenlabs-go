package types

import (
	"testing"
)

func TestTransferToNumberResultExotelSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_exotel_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "agent_message": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultExotelSuccessModel))
}

func TestTransferToNumberResultExotelSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_exotel_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "agent_message": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultExotelSuccessModel))
}
