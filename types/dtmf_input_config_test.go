package types

import (
	"testing"
)

func TestDtmfInputConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dtmf_input_timeout": 9.99,
  "hash_terminator": true,
  "redact_input": true
}`
	assertRoundTrip(t, wire, new(DtmfInputConfig))
}

func TestDtmfInputConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dtmf_input_timeout": 9.99,
  "hash_terminator": true,
  "redact_input": true
}`
	assertRoundTrip(t, wire, new(DtmfInputConfig))
}
