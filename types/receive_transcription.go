package types

import (
	"encoding/json"
	"fmt"
)

type ReceiveTranscription struct {
	SessionStarted                       *SessionStarted
	PartialTranscript                    *PartialTranscript
	CommittedTranscript                  *CommittedTranscript
	DiarizedTranscript                   *DiarizedTranscript
	SpeakersMerged                       *SpeakersMerged
	SpeakerNames                         *SpeakerNames
	CommittedTranscriptWithTimestamps    *CommittedTranscriptWithTimestamps
	CommittedTranscriptEntities          *CommittedTranscriptEntities
	RedactedTranscript                   *RedactedTranscript
	EditedTranscript                     *EditedTranscript
	ScribeWarning                        *ScribeWarning
	ScribeError                          *ScribeError
	ScribeAuthError                      *ScribeAuthError
	ScribeQuotaExceededError             *ScribeQuotaExceededError
	ScribeThrottledError                 *ScribeThrottledError
	ScribeUnacceptedTermsError           *ScribeUnacceptedTermsError
	ScribeRateLimitedError               *ScribeRateLimitedError
	ScribeQueueOverflowError             *ScribeQueueOverflowError
	ScribeResourceExhaustedError         *ScribeResourceExhaustedError
	ScribeSessionTimeLimitExceededError  *ScribeSessionTimeLimitExceededError
	ScribeInputError                     *ScribeInputError
	ScribeInvalidRequestError            *ScribeInvalidRequestError
	ScribeChunkSizeExceededError         *ScribeChunkSizeExceededError
	ScribeInsufficientAudioActivityError *ScribeInsufficientAudioActivityError
	ScribeTranscriberError               *ScribeTranscriberError
	typ                                  string
}

func (r *ReceiveTranscription) GetSessionStarted() *SessionStarted {
	if r == nil {
		return nil
	}
	if r.SessionStarted == nil {
		return nil
	}
	return r.SessionStarted
}

func (r *ReceiveTranscription) GetPartialTranscript() *PartialTranscript {
	if r == nil {
		return nil
	}
	if r.PartialTranscript == nil {
		return nil
	}
	return r.PartialTranscript
}

func (r *ReceiveTranscription) GetCommittedTranscript() *CommittedTranscript {
	if r == nil {
		return nil
	}
	if r.CommittedTranscript == nil {
		return nil
	}
	return r.CommittedTranscript
}

func (r *ReceiveTranscription) GetDiarizedTranscript() *DiarizedTranscript {
	if r == nil {
		return nil
	}
	if r.DiarizedTranscript == nil {
		return nil
	}
	return r.DiarizedTranscript
}

func (r *ReceiveTranscription) GetSpeakersMerged() *SpeakersMerged {
	if r == nil {
		return nil
	}
	if r.SpeakersMerged == nil {
		return nil
	}
	return r.SpeakersMerged
}

func (r *ReceiveTranscription) GetSpeakerNames() *SpeakerNames {
	if r == nil {
		return nil
	}
	if r.SpeakerNames == nil {
		return nil
	}
	return r.SpeakerNames
}

func (r *ReceiveTranscription) GetCommittedTranscriptWithTimestamps() *CommittedTranscriptWithTimestamps {
	if r == nil {
		return nil
	}
	if r.CommittedTranscriptWithTimestamps == nil {
		return nil
	}
	return r.CommittedTranscriptWithTimestamps
}

func (r *ReceiveTranscription) GetCommittedTranscriptEntities() *CommittedTranscriptEntities {
	if r == nil {
		return nil
	}
	if r.CommittedTranscriptEntities == nil {
		return nil
	}
	return r.CommittedTranscriptEntities
}

func (r *ReceiveTranscription) GetRedactedTranscript() *RedactedTranscript {
	if r == nil {
		return nil
	}
	if r.RedactedTranscript == nil {
		return nil
	}
	return r.RedactedTranscript
}

func (r *ReceiveTranscription) GetEditedTranscript() *EditedTranscript {
	if r == nil {
		return nil
	}
	if r.EditedTranscript == nil {
		return nil
	}
	return r.EditedTranscript
}

func (r *ReceiveTranscription) GetScribeWarning() *ScribeWarning {
	if r == nil {
		return nil
	}
	if r.ScribeWarning == nil {
		return nil
	}
	return r.ScribeWarning
}

func (r *ReceiveTranscription) GetScribeError() *ScribeError {
	if r == nil {
		return nil
	}
	if r.ScribeError == nil {
		return nil
	}
	return r.ScribeError
}

func (r *ReceiveTranscription) GetScribeAuthError() *ScribeAuthError {
	if r == nil {
		return nil
	}
	if r.ScribeAuthError == nil {
		return nil
	}
	return r.ScribeAuthError
}

func (r *ReceiveTranscription) GetScribeQuotaExceededError() *ScribeQuotaExceededError {
	if r == nil {
		return nil
	}
	if r.ScribeQuotaExceededError == nil {
		return nil
	}
	return r.ScribeQuotaExceededError
}

func (r *ReceiveTranscription) GetScribeThrottledError() *ScribeThrottledError {
	if r == nil {
		return nil
	}
	if r.ScribeThrottledError == nil {
		return nil
	}
	return r.ScribeThrottledError
}

func (r *ReceiveTranscription) GetScribeUnacceptedTermsError() *ScribeUnacceptedTermsError {
	if r == nil {
		return nil
	}
	if r.ScribeUnacceptedTermsError == nil {
		return nil
	}
	return r.ScribeUnacceptedTermsError
}

func (r *ReceiveTranscription) GetScribeRateLimitedError() *ScribeRateLimitedError {
	if r == nil {
		return nil
	}
	if r.ScribeRateLimitedError == nil {
		return nil
	}
	return r.ScribeRateLimitedError
}

func (r *ReceiveTranscription) GetScribeQueueOverflowError() *ScribeQueueOverflowError {
	if r == nil {
		return nil
	}
	if r.ScribeQueueOverflowError == nil {
		return nil
	}
	return r.ScribeQueueOverflowError
}

func (r *ReceiveTranscription) GetScribeResourceExhaustedError() *ScribeResourceExhaustedError {
	if r == nil {
		return nil
	}
	if r.ScribeResourceExhaustedError == nil {
		return nil
	}
	return r.ScribeResourceExhaustedError
}

func (r *ReceiveTranscription) GetScribeSessionTimeLimitExceededError() *ScribeSessionTimeLimitExceededError {
	if r == nil {
		return nil
	}
	if r.ScribeSessionTimeLimitExceededError == nil {
		return nil
	}
	return r.ScribeSessionTimeLimitExceededError
}

func (r *ReceiveTranscription) GetScribeInputError() *ScribeInputError {
	if r == nil {
		return nil
	}
	if r.ScribeInputError == nil {
		return nil
	}
	return r.ScribeInputError
}

func (r *ReceiveTranscription) GetScribeInvalidRequestError() *ScribeInvalidRequestError {
	if r == nil {
		return nil
	}
	if r.ScribeInvalidRequestError == nil {
		return nil
	}
	return r.ScribeInvalidRequestError
}

func (r *ReceiveTranscription) GetScribeChunkSizeExceededError() *ScribeChunkSizeExceededError {
	if r == nil {
		return nil
	}
	if r.ScribeChunkSizeExceededError == nil {
		return nil
	}
	return r.ScribeChunkSizeExceededError
}

func (r *ReceiveTranscription) GetScribeInsufficientAudioActivityError() *ScribeInsufficientAudioActivityError {
	if r == nil {
		return nil
	}
	if r.ScribeInsufficientAudioActivityError == nil {
		return nil
	}
	return r.ScribeInsufficientAudioActivityError
}

func (r *ReceiveTranscription) GetScribeTranscriberError() *ScribeTranscriberError {
	if r == nil {
		return nil
	}
	if r.ScribeTranscriberError == nil {
		return nil
	}
	return r.ScribeTranscriberError
}

func (r *ReceiveTranscription) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("message_type", "session_started") {
		matchedLiteral = true
		var valueSessionStartedLiteral *SessionStarted
		if err := json.Unmarshal(data, &valueSessionStartedLiteral); err == nil {
			r.typ = "SessionStarted"
			r.SessionStarted = valueSessionStartedLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "partial_transcript") {
		matchedLiteral = true
		var valuePartialTranscriptLiteral *PartialTranscript
		if err := json.Unmarshal(data, &valuePartialTranscriptLiteral); err == nil {
			r.typ = "PartialTranscript"
			r.PartialTranscript = valuePartialTranscriptLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "committed_transcript") {
		matchedLiteral = true
		var valueCommittedTranscriptLiteral *CommittedTranscript
		if err := json.Unmarshal(data, &valueCommittedTranscriptLiteral); err == nil {
			r.typ = "CommittedTranscript"
			r.CommittedTranscript = valueCommittedTranscriptLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "diarized_transcript") {
		matchedLiteral = true
		var valueDiarizedTranscriptLiteral *DiarizedTranscript
		if err := json.Unmarshal(data, &valueDiarizedTranscriptLiteral); err == nil {
			r.typ = "DiarizedTranscript"
			r.DiarizedTranscript = valueDiarizedTranscriptLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "speakers_merged") {
		matchedLiteral = true
		var valueSpeakersMergedLiteral *SpeakersMerged
		if err := json.Unmarshal(data, &valueSpeakersMergedLiteral); err == nil {
			r.typ = "SpeakersMerged"
			r.SpeakersMerged = valueSpeakersMergedLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "speaker_names") {
		matchedLiteral = true
		var valueSpeakerNamesLiteral *SpeakerNames
		if err := json.Unmarshal(data, &valueSpeakerNamesLiteral); err == nil {
			r.typ = "SpeakerNames"
			r.SpeakerNames = valueSpeakerNamesLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "committed_transcript_with_timestamps") {
		matchedLiteral = true
		var valueCommittedTranscriptWithTimestampsLiteral *CommittedTranscriptWithTimestamps
		if err := json.Unmarshal(data, &valueCommittedTranscriptWithTimestampsLiteral); err == nil {
			r.typ = "CommittedTranscriptWithTimestamps"
			r.CommittedTranscriptWithTimestamps = valueCommittedTranscriptWithTimestampsLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "committed_transcript_entities") {
		matchedLiteral = true
		var valueCommittedTranscriptEntitiesLiteral *CommittedTranscriptEntities
		if err := json.Unmarshal(data, &valueCommittedTranscriptEntitiesLiteral); err == nil {
			r.typ = "CommittedTranscriptEntities"
			r.CommittedTranscriptEntities = valueCommittedTranscriptEntitiesLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "redacted_transcript") {
		matchedLiteral = true
		var valueRedactedTranscriptLiteral *RedactedTranscript
		if err := json.Unmarshal(data, &valueRedactedTranscriptLiteral); err == nil {
			r.typ = "RedactedTranscript"
			r.RedactedTranscript = valueRedactedTranscriptLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "warning") {
		matchedLiteral = true
		var valueScribeWarningLiteral *ScribeWarning
		if err := json.Unmarshal(data, &valueScribeWarningLiteral); err == nil {
			r.typ = "ScribeWarning"
			r.ScribeWarning = valueScribeWarningLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "error") {
		matchedLiteral = true
		var valueScribeErrorLiteral *ScribeError
		if err := json.Unmarshal(data, &valueScribeErrorLiteral); err == nil {
			r.typ = "ScribeError"
			r.ScribeError = valueScribeErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "auth_error") {
		matchedLiteral = true
		var valueScribeAuthErrorLiteral *ScribeAuthError
		if err := json.Unmarshal(data, &valueScribeAuthErrorLiteral); err == nil {
			r.typ = "ScribeAuthError"
			r.ScribeAuthError = valueScribeAuthErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "quota_exceeded") {
		matchedLiteral = true
		var valueScribeQuotaExceededErrorLiteral *ScribeQuotaExceededError
		if err := json.Unmarshal(data, &valueScribeQuotaExceededErrorLiteral); err == nil {
			r.typ = "ScribeQuotaExceededError"
			r.ScribeQuotaExceededError = valueScribeQuotaExceededErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "commit_throttled") {
		matchedLiteral = true
		var valueScribeThrottledErrorLiteral *ScribeThrottledError
		if err := json.Unmarshal(data, &valueScribeThrottledErrorLiteral); err == nil {
			r.typ = "ScribeThrottledError"
			r.ScribeThrottledError = valueScribeThrottledErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "unaccepted_terms") {
		matchedLiteral = true
		var valueScribeUnacceptedTermsErrorLiteral *ScribeUnacceptedTermsError
		if err := json.Unmarshal(data, &valueScribeUnacceptedTermsErrorLiteral); err == nil {
			r.typ = "ScribeUnacceptedTermsError"
			r.ScribeUnacceptedTermsError = valueScribeUnacceptedTermsErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "rate_limited") {
		matchedLiteral = true
		var valueScribeRateLimitedErrorLiteral *ScribeRateLimitedError
		if err := json.Unmarshal(data, &valueScribeRateLimitedErrorLiteral); err == nil {
			r.typ = "ScribeRateLimitedError"
			r.ScribeRateLimitedError = valueScribeRateLimitedErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "queue_overflow") {
		matchedLiteral = true
		var valueScribeQueueOverflowErrorLiteral *ScribeQueueOverflowError
		if err := json.Unmarshal(data, &valueScribeQueueOverflowErrorLiteral); err == nil {
			r.typ = "ScribeQueueOverflowError"
			r.ScribeQueueOverflowError = valueScribeQueueOverflowErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "resource_exhausted") {
		matchedLiteral = true
		var valueScribeResourceExhaustedErrorLiteral *ScribeResourceExhaustedError
		if err := json.Unmarshal(data, &valueScribeResourceExhaustedErrorLiteral); err == nil {
			r.typ = "ScribeResourceExhaustedError"
			r.ScribeResourceExhaustedError = valueScribeResourceExhaustedErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "session_time_limit_exceeded") {
		matchedLiteral = true
		var valueScribeSessionTimeLimitExceededErrorLiteral *ScribeSessionTimeLimitExceededError
		if err := json.Unmarshal(data, &valueScribeSessionTimeLimitExceededErrorLiteral); err == nil {
			r.typ = "ScribeSessionTimeLimitExceededError"
			r.ScribeSessionTimeLimitExceededError = valueScribeSessionTimeLimitExceededErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "input_error") {
		matchedLiteral = true
		var valueScribeInputErrorLiteral *ScribeInputError
		if err := json.Unmarshal(data, &valueScribeInputErrorLiteral); err == nil {
			r.typ = "ScribeInputError"
			r.ScribeInputError = valueScribeInputErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "invalid_request") {
		matchedLiteral = true
		var valueScribeInvalidRequestErrorLiteral *ScribeInvalidRequestError
		if err := json.Unmarshal(data, &valueScribeInvalidRequestErrorLiteral); err == nil {
			r.typ = "ScribeInvalidRequestError"
			r.ScribeInvalidRequestError = valueScribeInvalidRequestErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "chunk_size_exceeded") {
		matchedLiteral = true
		var valueScribeChunkSizeExceededErrorLiteral *ScribeChunkSizeExceededError
		if err := json.Unmarshal(data, &valueScribeChunkSizeExceededErrorLiteral); err == nil {
			r.typ = "ScribeChunkSizeExceededError"
			r.ScribeChunkSizeExceededError = valueScribeChunkSizeExceededErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "insufficient_audio_activity") {
		matchedLiteral = true
		var valueScribeInsufficientAudioActivityErrorLiteral *ScribeInsufficientAudioActivityError
		if err := json.Unmarshal(data, &valueScribeInsufficientAudioActivityErrorLiteral); err == nil {
			r.typ = "ScribeInsufficientAudioActivityError"
			r.ScribeInsufficientAudioActivityError = valueScribeInsufficientAudioActivityErrorLiteral
			return nil
		}
	}
	if hasLiteral("message_type", "transcriber_error") {
		matchedLiteral = true
		var valueScribeTranscriberErrorLiteral *ScribeTranscriberError
		if err := json.Unmarshal(data, &valueScribeTranscriberErrorLiteral); err == nil {
			r.typ = "ScribeTranscriberError"
			r.ScribeTranscriberError = valueScribeTranscriberErrorLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
	}
	var valueSessionStarted *SessionStarted
	if err := json.Unmarshal(data, &valueSessionStarted); err == nil {
		r.typ = "SessionStarted"
		r.SessionStarted = valueSessionStarted
		return nil
	}
	var valuePartialTranscript *PartialTranscript
	if err := json.Unmarshal(data, &valuePartialTranscript); err == nil {
		r.typ = "PartialTranscript"
		r.PartialTranscript = valuePartialTranscript
		return nil
	}
	var valueCommittedTranscript *CommittedTranscript
	if err := json.Unmarshal(data, &valueCommittedTranscript); err == nil {
		r.typ = "CommittedTranscript"
		r.CommittedTranscript = valueCommittedTranscript
		return nil
	}
	var valueDiarizedTranscript *DiarizedTranscript
	if err := json.Unmarshal(data, &valueDiarizedTranscript); err == nil {
		r.typ = "DiarizedTranscript"
		r.DiarizedTranscript = valueDiarizedTranscript
		return nil
	}
	var valueSpeakersMerged *SpeakersMerged
	if err := json.Unmarshal(data, &valueSpeakersMerged); err == nil {
		r.typ = "SpeakersMerged"
		r.SpeakersMerged = valueSpeakersMerged
		return nil
	}
	var valueSpeakerNames *SpeakerNames
	if err := json.Unmarshal(data, &valueSpeakerNames); err == nil {
		r.typ = "SpeakerNames"
		r.SpeakerNames = valueSpeakerNames
		return nil
	}
	var valueCommittedTranscriptWithTimestamps *CommittedTranscriptWithTimestamps
	if err := json.Unmarshal(data, &valueCommittedTranscriptWithTimestamps); err == nil {
		r.typ = "CommittedTranscriptWithTimestamps"
		r.CommittedTranscriptWithTimestamps = valueCommittedTranscriptWithTimestamps
		return nil
	}
	var valueCommittedTranscriptEntities *CommittedTranscriptEntities
	if err := json.Unmarshal(data, &valueCommittedTranscriptEntities); err == nil {
		r.typ = "CommittedTranscriptEntities"
		r.CommittedTranscriptEntities = valueCommittedTranscriptEntities
		return nil
	}
	var valueRedactedTranscript *RedactedTranscript
	if err := json.Unmarshal(data, &valueRedactedTranscript); err == nil {
		r.typ = "RedactedTranscript"
		r.RedactedTranscript = valueRedactedTranscript
		return nil
	}
	var valueEditedTranscript *EditedTranscript
	if err := json.Unmarshal(data, &valueEditedTranscript); err == nil {
		r.typ = "EditedTranscript"
		r.EditedTranscript = valueEditedTranscript
		return nil
	}
	var valueScribeWarning *ScribeWarning
	if err := json.Unmarshal(data, &valueScribeWarning); err == nil {
		r.typ = "ScribeWarning"
		r.ScribeWarning = valueScribeWarning
		return nil
	}
	var valueScribeError *ScribeError
	if err := json.Unmarshal(data, &valueScribeError); err == nil {
		r.typ = "ScribeError"
		r.ScribeError = valueScribeError
		return nil
	}
	var valueScribeAuthError *ScribeAuthError
	if err := json.Unmarshal(data, &valueScribeAuthError); err == nil {
		r.typ = "ScribeAuthError"
		r.ScribeAuthError = valueScribeAuthError
		return nil
	}
	var valueScribeQuotaExceededError *ScribeQuotaExceededError
	if err := json.Unmarshal(data, &valueScribeQuotaExceededError); err == nil {
		r.typ = "ScribeQuotaExceededError"
		r.ScribeQuotaExceededError = valueScribeQuotaExceededError
		return nil
	}
	var valueScribeThrottledError *ScribeThrottledError
	if err := json.Unmarshal(data, &valueScribeThrottledError); err == nil {
		r.typ = "ScribeThrottledError"
		r.ScribeThrottledError = valueScribeThrottledError
		return nil
	}
	var valueScribeUnacceptedTermsError *ScribeUnacceptedTermsError
	if err := json.Unmarshal(data, &valueScribeUnacceptedTermsError); err == nil {
		r.typ = "ScribeUnacceptedTermsError"
		r.ScribeUnacceptedTermsError = valueScribeUnacceptedTermsError
		return nil
	}
	var valueScribeRateLimitedError *ScribeRateLimitedError
	if err := json.Unmarshal(data, &valueScribeRateLimitedError); err == nil {
		r.typ = "ScribeRateLimitedError"
		r.ScribeRateLimitedError = valueScribeRateLimitedError
		return nil
	}
	var valueScribeQueueOverflowError *ScribeQueueOverflowError
	if err := json.Unmarshal(data, &valueScribeQueueOverflowError); err == nil {
		r.typ = "ScribeQueueOverflowError"
		r.ScribeQueueOverflowError = valueScribeQueueOverflowError
		return nil
	}
	var valueScribeResourceExhaustedError *ScribeResourceExhaustedError
	if err := json.Unmarshal(data, &valueScribeResourceExhaustedError); err == nil {
		r.typ = "ScribeResourceExhaustedError"
		r.ScribeResourceExhaustedError = valueScribeResourceExhaustedError
		return nil
	}
	var valueScribeSessionTimeLimitExceededError *ScribeSessionTimeLimitExceededError
	if err := json.Unmarshal(data, &valueScribeSessionTimeLimitExceededError); err == nil {
		r.typ = "ScribeSessionTimeLimitExceededError"
		r.ScribeSessionTimeLimitExceededError = valueScribeSessionTimeLimitExceededError
		return nil
	}
	var valueScribeInputError *ScribeInputError
	if err := json.Unmarshal(data, &valueScribeInputError); err == nil {
		r.typ = "ScribeInputError"
		r.ScribeInputError = valueScribeInputError
		return nil
	}
	var valueScribeInvalidRequestError *ScribeInvalidRequestError
	if err := json.Unmarshal(data, &valueScribeInvalidRequestError); err == nil {
		r.typ = "ScribeInvalidRequestError"
		r.ScribeInvalidRequestError = valueScribeInvalidRequestError
		return nil
	}
	var valueScribeChunkSizeExceededError *ScribeChunkSizeExceededError
	if err := json.Unmarshal(data, &valueScribeChunkSizeExceededError); err == nil {
		r.typ = "ScribeChunkSizeExceededError"
		r.ScribeChunkSizeExceededError = valueScribeChunkSizeExceededError
		return nil
	}
	var valueScribeInsufficientAudioActivityError *ScribeInsufficientAudioActivityError
	if err := json.Unmarshal(data, &valueScribeInsufficientAudioActivityError); err == nil {
		r.typ = "ScribeInsufficientAudioActivityError"
		r.ScribeInsufficientAudioActivityError = valueScribeInsufficientAudioActivityError
		return nil
	}
	var valueScribeTranscriberError *ScribeTranscriberError
	if err := json.Unmarshal(data, &valueScribeTranscriberError); err == nil {
		r.typ = "ScribeTranscriberError"
		r.ScribeTranscriberError = valueScribeTranscriberError
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r ReceiveTranscription) MarshalJSON() ([]byte, error) {
	if r.typ == "SessionStarted" || r.SessionStarted != nil {
		return json.Marshal(r.SessionStarted)
	}
	if r.typ == "PartialTranscript" || r.PartialTranscript != nil {
		return json.Marshal(r.PartialTranscript)
	}
	if r.typ == "CommittedTranscript" || r.CommittedTranscript != nil {
		return json.Marshal(r.CommittedTranscript)
	}
	if r.typ == "DiarizedTranscript" || r.DiarizedTranscript != nil {
		return json.Marshal(r.DiarizedTranscript)
	}
	if r.typ == "SpeakersMerged" || r.SpeakersMerged != nil {
		return json.Marshal(r.SpeakersMerged)
	}
	if r.typ == "SpeakerNames" || r.SpeakerNames != nil {
		return json.Marshal(r.SpeakerNames)
	}
	if r.typ == "CommittedTranscriptWithTimestamps" || r.CommittedTranscriptWithTimestamps != nil {
		return json.Marshal(r.CommittedTranscriptWithTimestamps)
	}
	if r.typ == "CommittedTranscriptEntities" || r.CommittedTranscriptEntities != nil {
		return json.Marshal(r.CommittedTranscriptEntities)
	}
	if r.typ == "RedactedTranscript" || r.RedactedTranscript != nil {
		return json.Marshal(r.RedactedTranscript)
	}
	if r.typ == "EditedTranscript" || r.EditedTranscript != nil {
		return json.Marshal(r.EditedTranscript)
	}
	if r.typ == "ScribeWarning" || r.ScribeWarning != nil {
		return json.Marshal(r.ScribeWarning)
	}
	if r.typ == "ScribeError" || r.ScribeError != nil {
		return json.Marshal(r.ScribeError)
	}
	if r.typ == "ScribeAuthError" || r.ScribeAuthError != nil {
		return json.Marshal(r.ScribeAuthError)
	}
	if r.typ == "ScribeQuotaExceededError" || r.ScribeQuotaExceededError != nil {
		return json.Marshal(r.ScribeQuotaExceededError)
	}
	if r.typ == "ScribeThrottledError" || r.ScribeThrottledError != nil {
		return json.Marshal(r.ScribeThrottledError)
	}
	if r.typ == "ScribeUnacceptedTermsError" || r.ScribeUnacceptedTermsError != nil {
		return json.Marshal(r.ScribeUnacceptedTermsError)
	}
	if r.typ == "ScribeRateLimitedError" || r.ScribeRateLimitedError != nil {
		return json.Marshal(r.ScribeRateLimitedError)
	}
	if r.typ == "ScribeQueueOverflowError" || r.ScribeQueueOverflowError != nil {
		return json.Marshal(r.ScribeQueueOverflowError)
	}
	if r.typ == "ScribeResourceExhaustedError" || r.ScribeResourceExhaustedError != nil {
		return json.Marshal(r.ScribeResourceExhaustedError)
	}
	if r.typ == "ScribeSessionTimeLimitExceededError" || r.ScribeSessionTimeLimitExceededError != nil {
		return json.Marshal(r.ScribeSessionTimeLimitExceededError)
	}
	if r.typ == "ScribeInputError" || r.ScribeInputError != nil {
		return json.Marshal(r.ScribeInputError)
	}
	if r.typ == "ScribeInvalidRequestError" || r.ScribeInvalidRequestError != nil {
		return json.Marshal(r.ScribeInvalidRequestError)
	}
	if r.typ == "ScribeChunkSizeExceededError" || r.ScribeChunkSizeExceededError != nil {
		return json.Marshal(r.ScribeChunkSizeExceededError)
	}
	if r.typ == "ScribeInsufficientAudioActivityError" || r.ScribeInsufficientAudioActivityError != nil {
		return json.Marshal(r.ScribeInsufficientAudioActivityError)
	}
	if r.typ == "ScribeTranscriberError" || r.ScribeTranscriberError != nil {
		return json.Marshal(r.ScribeTranscriberError)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", r)
}

type ReceiveTranscriptionVisitor interface {
	VisitSessionStarted(*SessionStarted) error
	VisitPartialTranscript(*PartialTranscript) error
	VisitCommittedTranscript(*CommittedTranscript) error
	VisitDiarizedTranscript(*DiarizedTranscript) error
	VisitSpeakersMerged(*SpeakersMerged) error
	VisitSpeakerNames(*SpeakerNames) error
	VisitCommittedTranscriptWithTimestamps(*CommittedTranscriptWithTimestamps) error
	VisitCommittedTranscriptEntities(*CommittedTranscriptEntities) error
	VisitRedactedTranscript(*RedactedTranscript) error
	VisitEditedTranscript(*EditedTranscript) error
	VisitScribeWarning(*ScribeWarning) error
	VisitScribeError(*ScribeError) error
	VisitScribeAuthError(*ScribeAuthError) error
	VisitScribeQuotaExceededError(*ScribeQuotaExceededError) error
	VisitScribeThrottledError(*ScribeThrottledError) error
	VisitScribeUnacceptedTermsError(*ScribeUnacceptedTermsError) error
	VisitScribeRateLimitedError(*ScribeRateLimitedError) error
	VisitScribeQueueOverflowError(*ScribeQueueOverflowError) error
	VisitScribeResourceExhaustedError(*ScribeResourceExhaustedError) error
	VisitScribeSessionTimeLimitExceededError(*ScribeSessionTimeLimitExceededError) error
	VisitScribeInputError(*ScribeInputError) error
	VisitScribeInvalidRequestError(*ScribeInvalidRequestError) error
	VisitScribeChunkSizeExceededError(*ScribeChunkSizeExceededError) error
	VisitScribeInsufficientAudioActivityError(*ScribeInsufficientAudioActivityError) error
	VisitScribeTranscriberError(*ScribeTranscriberError) error
}

func (r *ReceiveTranscription) Accept(visitor ReceiveTranscriptionVisitor) error {
	if r.typ == "SessionStarted" || r.SessionStarted != nil {
		return visitor.VisitSessionStarted(r.SessionStarted)
	}
	if r.typ == "PartialTranscript" || r.PartialTranscript != nil {
		return visitor.VisitPartialTranscript(r.PartialTranscript)
	}
	if r.typ == "CommittedTranscript" || r.CommittedTranscript != nil {
		return visitor.VisitCommittedTranscript(r.CommittedTranscript)
	}
	if r.typ == "DiarizedTranscript" || r.DiarizedTranscript != nil {
		return visitor.VisitDiarizedTranscript(r.DiarizedTranscript)
	}
	if r.typ == "SpeakersMerged" || r.SpeakersMerged != nil {
		return visitor.VisitSpeakersMerged(r.SpeakersMerged)
	}
	if r.typ == "SpeakerNames" || r.SpeakerNames != nil {
		return visitor.VisitSpeakerNames(r.SpeakerNames)
	}
	if r.typ == "CommittedTranscriptWithTimestamps" || r.CommittedTranscriptWithTimestamps != nil {
		return visitor.VisitCommittedTranscriptWithTimestamps(r.CommittedTranscriptWithTimestamps)
	}
	if r.typ == "CommittedTranscriptEntities" || r.CommittedTranscriptEntities != nil {
		return visitor.VisitCommittedTranscriptEntities(r.CommittedTranscriptEntities)
	}
	if r.typ == "RedactedTranscript" || r.RedactedTranscript != nil {
		return visitor.VisitRedactedTranscript(r.RedactedTranscript)
	}
	if r.typ == "EditedTranscript" || r.EditedTranscript != nil {
		return visitor.VisitEditedTranscript(r.EditedTranscript)
	}
	if r.typ == "ScribeWarning" || r.ScribeWarning != nil {
		return visitor.VisitScribeWarning(r.ScribeWarning)
	}
	if r.typ == "ScribeError" || r.ScribeError != nil {
		return visitor.VisitScribeError(r.ScribeError)
	}
	if r.typ == "ScribeAuthError" || r.ScribeAuthError != nil {
		return visitor.VisitScribeAuthError(r.ScribeAuthError)
	}
	if r.typ == "ScribeQuotaExceededError" || r.ScribeQuotaExceededError != nil {
		return visitor.VisitScribeQuotaExceededError(r.ScribeQuotaExceededError)
	}
	if r.typ == "ScribeThrottledError" || r.ScribeThrottledError != nil {
		return visitor.VisitScribeThrottledError(r.ScribeThrottledError)
	}
	if r.typ == "ScribeUnacceptedTermsError" || r.ScribeUnacceptedTermsError != nil {
		return visitor.VisitScribeUnacceptedTermsError(r.ScribeUnacceptedTermsError)
	}
	if r.typ == "ScribeRateLimitedError" || r.ScribeRateLimitedError != nil {
		return visitor.VisitScribeRateLimitedError(r.ScribeRateLimitedError)
	}
	if r.typ == "ScribeQueueOverflowError" || r.ScribeQueueOverflowError != nil {
		return visitor.VisitScribeQueueOverflowError(r.ScribeQueueOverflowError)
	}
	if r.typ == "ScribeResourceExhaustedError" || r.ScribeResourceExhaustedError != nil {
		return visitor.VisitScribeResourceExhaustedError(r.ScribeResourceExhaustedError)
	}
	if r.typ == "ScribeSessionTimeLimitExceededError" || r.ScribeSessionTimeLimitExceededError != nil {
		return visitor.VisitScribeSessionTimeLimitExceededError(r.ScribeSessionTimeLimitExceededError)
	}
	if r.typ == "ScribeInputError" || r.ScribeInputError != nil {
		return visitor.VisitScribeInputError(r.ScribeInputError)
	}
	if r.typ == "ScribeInvalidRequestError" || r.ScribeInvalidRequestError != nil {
		return visitor.VisitScribeInvalidRequestError(r.ScribeInvalidRequestError)
	}
	if r.typ == "ScribeChunkSizeExceededError" || r.ScribeChunkSizeExceededError != nil {
		return visitor.VisitScribeChunkSizeExceededError(r.ScribeChunkSizeExceededError)
	}
	if r.typ == "ScribeInsufficientAudioActivityError" || r.ScribeInsufficientAudioActivityError != nil {
		return visitor.VisitScribeInsufficientAudioActivityError(r.ScribeInsufficientAudioActivityError)
	}
	if r.typ == "ScribeTranscriberError" || r.ScribeTranscriberError != nil {
		return visitor.VisitScribeTranscriberError(r.ScribeTranscriberError)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", r)
}
