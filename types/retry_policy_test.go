package types

import (
	"testing"
)

func TestRetryPolicyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "retry_delay_seconds": 123,
  "max_short_retries": 30,
  "max_total_calls": 30,
  "max_callbacks": 30,
  "max_redirects": 30,
  "interval_placement": "start",
  "rotate_step_seconds": 123,
  "max_consecutive_failures_per_number": 30
}`
	assertRoundTrip(t, wire, new(RetryPolicy))
}

func TestRetryPolicyRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "retry_delay_seconds": 123,
  "max_short_retries": 30,
  "max_total_calls": 30,
  "max_callbacks": 30,
  "max_redirects": 30,
  "interval_placement": "start",
  "rotate_step_seconds": 123,
  "max_consecutive_failures_per_number": 30
}`
	assertRoundTrip(t, wire, new(RetryPolicy))
}
