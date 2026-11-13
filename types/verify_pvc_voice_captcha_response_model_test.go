package types

import (
	"testing"
)

func TestVerifyPvcVoiceCaptchaResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(VerifyPvcVoiceCaptchaResponseModel))
}
