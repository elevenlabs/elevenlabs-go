package types

import (
	"fmt"
)

type TerminalDisposition string

const (
	TerminalDispositionCampaignEnded            TerminalDisposition = "campaign_ended"
	TerminalDispositionRemoved                  TerminalDisposition = "removed"
	TerminalDispositionDoNotCall                TerminalDisposition = "do_not_call"
	TerminalDispositionCompleted                TerminalDisposition = "completed"
	TerminalDispositionCancelled                TerminalDisposition = "cancelled"
	TerminalDispositionSuppressed               TerminalDisposition = "suppressed"
	TerminalDispositionAttemptCeilingReached    TerminalDisposition = "attempt_ceiling_reached"
	TerminalDispositionCallbackCeilingReached   TerminalDisposition = "callback_ceiling_reached"
	TerminalDispositionRedirectCeilingReached   TerminalDisposition = "redirect_ceiling_reached"
	TerminalDispositionPermissionNotGranted     TerminalDisposition = "permission_not_granted"
	TerminalDispositionPermissionDeclined       TerminalDisposition = "permission_declined"
	TerminalDispositionPermissionRequestBlocked TerminalDisposition = "permission_request_blocked"
	TerminalDispositionPermissionExpired        TerminalDisposition = "permission_expired"
	TerminalDispositionCallbackTimeInvalid      TerminalDisposition = "callback_time_invalid"
	TerminalDispositionRedirectNumberInvalid    TerminalDisposition = "redirect_number_invalid"
	TerminalDispositionCouldntProcess           TerminalDisposition = "couldnt_process"
)

func NewTerminalDispositionFromString(s string) (TerminalDisposition, error) {
	switch s {
	case "campaign_ended":
		return TerminalDispositionCampaignEnded, nil
	case "removed":
		return TerminalDispositionRemoved, nil
	case "do_not_call":
		return TerminalDispositionDoNotCall, nil
	case "completed":
		return TerminalDispositionCompleted, nil
	case "cancelled":
		return TerminalDispositionCancelled, nil
	case "suppressed":
		return TerminalDispositionSuppressed, nil
	case "attempt_ceiling_reached":
		return TerminalDispositionAttemptCeilingReached, nil
	case "callback_ceiling_reached":
		return TerminalDispositionCallbackCeilingReached, nil
	case "redirect_ceiling_reached":
		return TerminalDispositionRedirectCeilingReached, nil
	case "permission_not_granted":
		return TerminalDispositionPermissionNotGranted, nil
	case "permission_declined":
		return TerminalDispositionPermissionDeclined, nil
	case "permission_request_blocked":
		return TerminalDispositionPermissionRequestBlocked, nil
	case "permission_expired":
		return TerminalDispositionPermissionExpired, nil
	case "callback_time_invalid":
		return TerminalDispositionCallbackTimeInvalid, nil
	case "redirect_number_invalid":
		return TerminalDispositionRedirectNumberInvalid, nil
	case "couldnt_process":
		return TerminalDispositionCouldntProcess, nil
	}
	var t TerminalDisposition
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TerminalDisposition) Ptr() *TerminalDisposition {
	return &e
}
