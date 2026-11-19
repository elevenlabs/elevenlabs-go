package types

import (
	"testing"
)

func TestTransferToNumberResultSipSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_sip_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultSipSuccessModel))
}

func TestTransferToNumberResultSipSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_sip_success",
  "status": "success",
  "transfer_number": "value",
  "reason": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultSipSuccessModel))
}
