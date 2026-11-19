package types

import (
	"testing"
)

func TestSubscriptionExtrasResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "concurrency": 123,
  "convai_concurrency": 123,
  "enterprise_music_concurrency": 123,
  "music_finetunes_total_limit": 123,
  "music_finetunes_monthly_limit": 123,
  "music_finetunes_concurrency_limit": 123,
  "convai_chars_per_minute": 123,
  "convai_asr_chars_per_minute": 123,
  "force_logging_disabled": true,
  "can_request_manual_pro_voice_verification": true,
  "can_bypass_voice_captcha": true,
  "moderation": {
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
  },
  "unused_characters_rolled_over_from_previous_period": 123,
  "overused_characters_rolled_over_from_previous_period": 123,
  "usage": {
    "rollover_credits_quota": 123,
    "subscription_cycle_credits_quota": 123,
    "manually_gifted_credits_quota": 123,
    "payg_credits_quota": 123,
    "rollover_credits_used": 123,
    "subscription_cycle_credits_used": 123,
    "manually_gifted_credits_used": 123,
    "payg_credits_used": 123,
    "paid_usage_based_credits_used": 123,
    "actual_reported_credits": 123
  }
}`
	assertRoundTrip(t, wire, new(SubscriptionExtrasResponseModel))
}
