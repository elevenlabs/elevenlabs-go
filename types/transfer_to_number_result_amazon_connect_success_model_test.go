package types

import (
	"testing"
)

func TestTransferToNumberResultAmazonConnectSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_amazon_connect_success",
  "status": "success",
  "reason": "value",
  "client_message": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultAmazonConnectSuccessModel))
}

func TestTransferToNumberResultAmazonConnectSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_number_amazon_connect_success",
  "status": "success",
  "reason": "value",
  "client_message": "value",
  "note": "value"
}`
	assertRoundTrip(t, wire, new(TransferToNumberResultAmazonConnectSuccessModel))
}
