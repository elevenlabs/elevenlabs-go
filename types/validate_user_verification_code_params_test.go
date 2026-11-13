package types

import (
	"testing"
)

func TestValidateUserVerificationCodeParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "validate_user_verification_code"
}`
	assertRoundTrip(t, wire, new(ValidateUserVerificationCodeParams))
}
