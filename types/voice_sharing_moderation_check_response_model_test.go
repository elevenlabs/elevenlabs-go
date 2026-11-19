package types

import (
	"testing"
)

func TestVoiceSharingModerationCheckResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "date_checked_unix": 123,
  "name_value": "value",
  "name_check": true,
  "description_value": "value",
  "description_check": true,
  "sample_ids": [
    "value"
  ],
  "sample_checks": [
    9.99
  ],
  "captcha_ids": [
    "value"
  ],
  "captcha_checks": [
    9.99
  ]
}`
	assertRoundTrip(t, wire, new(VoiceSharingModerationCheckResponseModel))
}
