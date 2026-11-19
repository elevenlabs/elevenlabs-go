package types

import (
	"testing"
)

func TestRawBusinessInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "company_name": "value",
  "description": "value",
  "business_phone": "value",
  "locations": [
    {
      "name": "value",
      "address": "value",
      "operating_hours": {
        "key": {}
      }
    }
  ],
  "business_country": "value",
  "business_state": "value",
  "suggested_currency": "value",
  "suggested_language": "value"
}`
	assertRoundTrip(t, wire, new(RawBusinessInfo))
}
