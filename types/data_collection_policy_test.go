package types

import (
	"testing"
)

func TestDataCollectionPolicyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "callback_time": "value",
  "redirect_phone_number": "value",
  "stop": "value",
  "mappings": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(DataCollectionPolicy))
}
