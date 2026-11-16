package types

import (
	"testing"
)

func TestBatchFailureResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "failure",
  "error_code": 123,
  "error_status": "value",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(BatchFailureResponseModel))
}

func TestBatchFailureResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "failure",
  "error_code": 123,
  "error_status": "value",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(BatchFailureResponseModel))
}

func TestBatchFailureResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "failure",
  "error_code": 123,
  "error_status": "value",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(BatchFailureResponseModel))
}

func TestBatchFailureResponseModelRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "failure",
  "error_code": 123,
  "error_status": "value",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(BatchFailureResponseModel))
}
