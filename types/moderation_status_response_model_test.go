package types

import (
	"testing"
)

func TestModerationStatusResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_in_probation": true,
  "enterprise_check_nogo_voice": true,
  "enterprise_check_block_nogo_voice": true,
  "never_live_moderate": true,
  "nogo_voice_similar_voice_upload_count": 123,
  "enterprise_background_moderation_enabled": true,
  "is_ivc_captcha_required": true,
  "safety_status": "appeal_approved",
  "warning_status": "warning",
  "on_watchlist": true
}`
	assertRoundTrip(t, wire, new(ModerationStatusResponseModel))
}
