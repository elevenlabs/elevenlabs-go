package types

import (
	"testing"
)

func TestTransferToNumberResultErrorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_error",
  "status": "error",
  "error": "value",
  "details": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultErrorModel))
}

func TestTransferToNumberResultErrorModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_error",
  "status": "error",
  "error": "value",
  "details": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultErrorModel))
}
