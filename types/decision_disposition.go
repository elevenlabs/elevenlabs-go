package types

import (
	"fmt"
)

// How a settled call's decision left its lead.
type DecisionDisposition string

const (
	DecisionDispositionCompleted                DecisionDisposition = "completed"
	DecisionDispositionRetryQueued              DecisionDisposition = "retry_queued"
	DecisionDispositionCallbackQueued           DecisionDisposition = "callback_queued"
	DecisionDispositionRedirectQueued           DecisionDisposition = "redirect_queued"
	DecisionDispositionAttention                DecisionDisposition = "attention"
	DecisionDispositionCancelled                DecisionDisposition = "cancelled"
	DecisionDispositionDoNotCall                DecisionDisposition = "do_not_call"
	DecisionDispositionSuppressed               DecisionDisposition = "suppressed"
	DecisionDispositionWindowSuspended          DecisionDisposition = "window_suspended"
	DecisionDispositionAttemptCeilingReached    DecisionDisposition = "attempt_ceiling_reached"
	DecisionDispositionCallbackCeilingReached   DecisionDisposition = "callback_ceiling_reached"
	DecisionDispositionRedirectCeilingReached   DecisionDisposition = "redirect_ceiling_reached"
	DecisionDispositionCallbackTimeInvalid      DecisionDisposition = "callback_time_invalid"
	DecisionDispositionRedirectNumberInvalid    DecisionDisposition = "redirect_number_invalid"
	DecisionDispositionCouldntProcess           DecisionDisposition = "couldnt_process"
	DecisionDispositionSuppressedDoNotCall      DecisionDisposition = "suppressed_do_not_call"
	DecisionDispositionSuppressedAfterStop      DecisionDisposition = "suppressed_after_stop"
	DecisionDispositionPermissionNotGranted     DecisionDisposition = "permission_not_granted"
	DecisionDispositionPermissionDeclined       DecisionDisposition = "permission_declined"
	DecisionDispositionPermissionRequestBlocked DecisionDisposition = "permission_request_blocked"
	DecisionDispositionPermissionExpired        DecisionDisposition = "permission_expired"
)

func NewDecisionDispositionFromString(s string) (DecisionDisposition, error) {
	switch s {
	case "completed":
		return DecisionDispositionCompleted, nil
	case "retry_queued":
		return DecisionDispositionRetryQueued, nil
	case "callback_queued":
		return DecisionDispositionCallbackQueued, nil
	case "redirect_queued":
		return DecisionDispositionRedirectQueued, nil
	case "attention":
		return DecisionDispositionAttention, nil
	case "cancelled":
		return DecisionDispositionCancelled, nil
	case "do_not_call":
		return DecisionDispositionDoNotCall, nil
	case "suppressed":
		return DecisionDispositionSuppressed, nil
	case "window_suspended":
		return DecisionDispositionWindowSuspended, nil
	case "attempt_ceiling_reached":
		return DecisionDispositionAttemptCeilingReached, nil
	case "callback_ceiling_reached":
		return DecisionDispositionCallbackCeilingReached, nil
	case "redirect_ceiling_reached":
		return DecisionDispositionRedirectCeilingReached, nil
	case "callback_time_invalid":
		return DecisionDispositionCallbackTimeInvalid, nil
	case "redirect_number_invalid":
		return DecisionDispositionRedirectNumberInvalid, nil
	case "couldnt_process":
		return DecisionDispositionCouldntProcess, nil
	case "suppressed_do_not_call":
		return DecisionDispositionSuppressedDoNotCall, nil
	case "suppressed_after_stop":
		return DecisionDispositionSuppressedAfterStop, nil
	case "permission_not_granted":
		return DecisionDispositionPermissionNotGranted, nil
	case "permission_declined":
		return DecisionDispositionPermissionDeclined, nil
	case "permission_request_blocked":
		return DecisionDispositionPermissionRequestBlocked, nil
	case "permission_expired":
		return DecisionDispositionPermissionExpired, nil
	}
	var t DecisionDisposition
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DecisionDisposition) Ptr() *DecisionDisposition {
	return &e
}
