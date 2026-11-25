package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ReceiveSubtitlingStreamMessage struct {
	MessageType       string
	AuthError         *SubtitlingAuthError
	CommittedSubtitle *SubtitlingCommittedSubtitle
	Error             *SubtitlingError
	FinalTranscript   *SubtitlingFinalTranscript
	InputError        *SubtitlingInputError
	PartialSubtitle   *SubtitlingPartialSubtitle
	PartialTranscript *SubtitlingPartialTranscript
	PaymentIssue      *SubtitlingPaymentIssueError
	QueueOverflow     *SubtitlingQueueOverflowError
	QuotaExceeded     *SubtitlingQuotaExceededError
	RateLimited       *SubtitlingRateLimitedError
	SessionStarted    *SubtitlingSessionStarted
	Status            *SubtitlingStatus
	TranscriberError  *SubtitlingTranscriberError
	TranslationError  *SubtitlingTranslationError
	UnacceptedTerms   *SubtitlingUnacceptedTermsError
}

func (r *ReceiveSubtitlingStreamMessage) GetAuthError() *SubtitlingAuthError {
	if r == nil {
		return nil
	}
	if r.AuthError == nil {
		return nil
	}
	return r.AuthError
}

func (r *ReceiveSubtitlingStreamMessage) GetCommittedSubtitle() *SubtitlingCommittedSubtitle {
	if r == nil {
		return nil
	}
	if r.CommittedSubtitle == nil {
		return nil
	}
	return r.CommittedSubtitle
}

func (r *ReceiveSubtitlingStreamMessage) GetError() *SubtitlingError {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveSubtitlingStreamMessage) GetFinalTranscript() *SubtitlingFinalTranscript {
	if r == nil {
		return nil
	}
	if r.FinalTranscript == nil {
		return nil
	}
	return r.FinalTranscript
}

func (r *ReceiveSubtitlingStreamMessage) GetInputError() *SubtitlingInputError {
	if r == nil {
		return nil
	}
	if r.InputError == nil {
		return nil
	}
	return r.InputError
}

func (r *ReceiveSubtitlingStreamMessage) GetPartialSubtitle() *SubtitlingPartialSubtitle {
	if r == nil {
		return nil
	}
	if r.PartialSubtitle == nil {
		return nil
	}
	return r.PartialSubtitle
}

func (r *ReceiveSubtitlingStreamMessage) GetPartialTranscript() *SubtitlingPartialTranscript {
	if r == nil {
		return nil
	}
	if r.PartialTranscript == nil {
		return nil
	}
	return r.PartialTranscript
}

func (r *ReceiveSubtitlingStreamMessage) GetPaymentIssue() *SubtitlingPaymentIssueError {
	if r == nil {
		return nil
	}
	if r.PaymentIssue == nil {
		return nil
	}
	return r.PaymentIssue
}

func (r *ReceiveSubtitlingStreamMessage) GetQueueOverflow() *SubtitlingQueueOverflowError {
	if r == nil {
		return nil
	}
	if r.QueueOverflow == nil {
		return nil
	}
	return r.QueueOverflow
}

func (r *ReceiveSubtitlingStreamMessage) GetQuotaExceeded() *SubtitlingQuotaExceededError {
	if r == nil {
		return nil
	}
	if r.QuotaExceeded == nil {
		return nil
	}
	return r.QuotaExceeded
}

func (r *ReceiveSubtitlingStreamMessage) GetRateLimited() *SubtitlingRateLimitedError {
	if r == nil {
		return nil
	}
	if r.RateLimited == nil {
		return nil
	}
	return r.RateLimited
}

func (r *ReceiveSubtitlingStreamMessage) GetSessionStarted() *SubtitlingSessionStarted {
	if r == nil {
		return nil
	}
	if r.SessionStarted == nil {
		return nil
	}
	return r.SessionStarted
}

func (r *ReceiveSubtitlingStreamMessage) GetStatus() *SubtitlingStatus {
	if r == nil {
		return nil
	}
	if r.Status == nil {
		return nil
	}
	return r.Status
}

func (r *ReceiveSubtitlingStreamMessage) GetTranscriberError() *SubtitlingTranscriberError {
	if r == nil {
		return nil
	}
	if r.TranscriberError == nil {
		return nil
	}
	return r.TranscriberError
}

func (r *ReceiveSubtitlingStreamMessage) GetTranslationError() *SubtitlingTranslationError {
	if r == nil {
		return nil
	}
	if r.TranslationError == nil {
		return nil
	}
	return r.TranslationError
}

func (r *ReceiveSubtitlingStreamMessage) GetUnacceptedTerms() *SubtitlingUnacceptedTermsError {
	if r == nil {
		return nil
	}
	if r.UnacceptedTerms == nil {
		return nil
	}
	return r.UnacceptedTerms
}

func (r *ReceiveSubtitlingStreamMessage) GetMessageType() string {
	if r == nil {
		return ""
	}
	return r.MessageType
}

func (r *ReceiveSubtitlingStreamMessage) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		MessageType string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.MessageType = unmarshaler.MessageType
	if unmarshaler.MessageType == "" {
		return fmt.Errorf("%T did not include discriminant message_type", r)
	}
	switch unmarshaler.MessageType {
	case "auth_error":
		value := new(SubtitlingAuthError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.AuthError = value
	case "committed_subtitle":
		value := new(SubtitlingCommittedSubtitle)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.CommittedSubtitle = value
	case "error":
		value := new(SubtitlingError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "final_transcript":
		value := new(SubtitlingFinalTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.FinalTranscript = value
	case "input_error":
		value := new(SubtitlingInputError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.InputError = value
	case "partial_subtitle":
		value := new(SubtitlingPartialSubtitle)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PartialSubtitle = value
	case "partial_transcript":
		value := new(SubtitlingPartialTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PartialTranscript = value
	case "payment_issue":
		value := new(SubtitlingPaymentIssueError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PaymentIssue = value
	case "queue_overflow":
		value := new(SubtitlingQueueOverflowError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.QueueOverflow = value
	case "quota_exceeded":
		value := new(SubtitlingQuotaExceededError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.QuotaExceeded = value
	case "rate_limited":
		value := new(SubtitlingRateLimitedError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.RateLimited = value
	case "session_started":
		value := new(SubtitlingSessionStarted)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.SessionStarted = value
	case "status":
		value := new(SubtitlingStatus)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Status = value
	case "transcriber_error":
		value := new(SubtitlingTranscriberError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranscriberError = value
	case "translation_error":
		value := new(SubtitlingTranslationError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranslationError = value
	case "unaccepted_terms":
		value := new(SubtitlingUnacceptedTermsError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.UnacceptedTerms = value
	}
	return nil
}

func (r ReceiveSubtitlingStreamMessage) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.AuthError != nil {
		return internal.MarshalJSONWithDiscriminant(r.AuthError, "message_type", "auth_error")
	}
	if r.CommittedSubtitle != nil {
		return internal.MarshalJSONWithDiscriminant(r.CommittedSubtitle, "message_type", "committed_subtitle")
	}
	if r.Error != nil {
		return internal.MarshalJSONWithDiscriminant(r.Error, "message_type", "error")
	}
	if r.FinalTranscript != nil {
		return internal.MarshalJSONWithDiscriminant(r.FinalTranscript, "message_type", "final_transcript")
	}
	if r.InputError != nil {
		return internal.MarshalJSONWithDiscriminant(r.InputError, "message_type", "input_error")
	}
	if r.PartialSubtitle != nil {
		return internal.MarshalJSONWithDiscriminant(r.PartialSubtitle, "message_type", "partial_subtitle")
	}
	if r.PartialTranscript != nil {
		return internal.MarshalJSONWithDiscriminant(r.PartialTranscript, "message_type", "partial_transcript")
	}
	if r.PaymentIssue != nil {
		return internal.MarshalJSONWithDiscriminant(r.PaymentIssue, "message_type", "payment_issue")
	}
	if r.QueueOverflow != nil {
		return internal.MarshalJSONWithDiscriminant(r.QueueOverflow, "message_type", "queue_overflow")
	}
	if r.QuotaExceeded != nil {
		return internal.MarshalJSONWithDiscriminant(r.QuotaExceeded, "message_type", "quota_exceeded")
	}
	if r.RateLimited != nil {
		return internal.MarshalJSONWithDiscriminant(r.RateLimited, "message_type", "rate_limited")
	}
	if r.SessionStarted != nil {
		return internal.MarshalJSONWithDiscriminant(r.SessionStarted, "message_type", "session_started")
	}
	if r.Status != nil {
		return internal.MarshalJSONWithDiscriminant(r.Status, "message_type", "status")
	}
	if r.TranscriberError != nil {
		return internal.MarshalJSONWithDiscriminant(r.TranscriberError, "message_type", "transcriber_error")
	}
	if r.TranslationError != nil {
		return internal.MarshalJSONWithDiscriminant(r.TranslationError, "message_type", "translation_error")
	}
	if r.UnacceptedTerms != nil {
		return internal.MarshalJSONWithDiscriminant(r.UnacceptedTerms, "message_type", "unaccepted_terms")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type ReceiveSubtitlingStreamMessageVisitor interface {
	VisitAuthError(*SubtitlingAuthError) error
	VisitCommittedSubtitle(*SubtitlingCommittedSubtitle) error
	VisitError(*SubtitlingError) error
	VisitFinalTranscript(*SubtitlingFinalTranscript) error
	VisitInputError(*SubtitlingInputError) error
	VisitPartialSubtitle(*SubtitlingPartialSubtitle) error
	VisitPartialTranscript(*SubtitlingPartialTranscript) error
	VisitPaymentIssue(*SubtitlingPaymentIssueError) error
	VisitQueueOverflow(*SubtitlingQueueOverflowError) error
	VisitQuotaExceeded(*SubtitlingQuotaExceededError) error
	VisitRateLimited(*SubtitlingRateLimitedError) error
	VisitSessionStarted(*SubtitlingSessionStarted) error
	VisitStatus(*SubtitlingStatus) error
	VisitTranscriberError(*SubtitlingTranscriberError) error
	VisitTranslationError(*SubtitlingTranslationError) error
	VisitUnacceptedTerms(*SubtitlingUnacceptedTermsError) error
}

func (r *ReceiveSubtitlingStreamMessage) Accept(visitor ReceiveSubtitlingStreamMessageVisitor) error {
	if r.AuthError != nil {
		return visitor.VisitAuthError(r.AuthError)
	}
	if r.CommittedSubtitle != nil {
		return visitor.VisitCommittedSubtitle(r.CommittedSubtitle)
	}
	if r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	if r.FinalTranscript != nil {
		return visitor.VisitFinalTranscript(r.FinalTranscript)
	}
	if r.InputError != nil {
		return visitor.VisitInputError(r.InputError)
	}
	if r.PartialSubtitle != nil {
		return visitor.VisitPartialSubtitle(r.PartialSubtitle)
	}
	if r.PartialTranscript != nil {
		return visitor.VisitPartialTranscript(r.PartialTranscript)
	}
	if r.PaymentIssue != nil {
		return visitor.VisitPaymentIssue(r.PaymentIssue)
	}
	if r.QueueOverflow != nil {
		return visitor.VisitQueueOverflow(r.QueueOverflow)
	}
	if r.QuotaExceeded != nil {
		return visitor.VisitQuotaExceeded(r.QuotaExceeded)
	}
	if r.RateLimited != nil {
		return visitor.VisitRateLimited(r.RateLimited)
	}
	if r.SessionStarted != nil {
		return visitor.VisitSessionStarted(r.SessionStarted)
	}
	if r.Status != nil {
		return visitor.VisitStatus(r.Status)
	}
	if r.TranscriberError != nil {
		return visitor.VisitTranscriberError(r.TranscriberError)
	}
	if r.TranslationError != nil {
		return visitor.VisitTranslationError(r.TranslationError)
	}
	if r.UnacceptedTerms != nil {
		return visitor.VisitUnacceptedTerms(r.UnacceptedTerms)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *ReceiveSubtitlingStreamMessage) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.AuthError != nil {
		fields = append(fields, "auth_error")
	}
	if r.CommittedSubtitle != nil {
		fields = append(fields, "committed_subtitle")
	}
	if r.Error != nil {
		fields = append(fields, "error")
	}
	if r.FinalTranscript != nil {
		fields = append(fields, "final_transcript")
	}
	if r.InputError != nil {
		fields = append(fields, "input_error")
	}
	if r.PartialSubtitle != nil {
		fields = append(fields, "partial_subtitle")
	}
	if r.PartialTranscript != nil {
		fields = append(fields, "partial_transcript")
	}
	if r.PaymentIssue != nil {
		fields = append(fields, "payment_issue")
	}
	if r.QueueOverflow != nil {
		fields = append(fields, "queue_overflow")
	}
	if r.QuotaExceeded != nil {
		fields = append(fields, "quota_exceeded")
	}
	if r.RateLimited != nil {
		fields = append(fields, "rate_limited")
	}
	if r.SessionStarted != nil {
		fields = append(fields, "session_started")
	}
	if r.Status != nil {
		fields = append(fields, "status")
	}
	if r.TranscriberError != nil {
		fields = append(fields, "transcriber_error")
	}
	if r.TranslationError != nil {
		fields = append(fields, "translation_error")
	}
	if r.UnacceptedTerms != nil {
		fields = append(fields, "unaccepted_terms")
	}
	if len(fields) == 0 {
		if r.MessageType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", r, r.MessageType)
		}
		return fmt.Errorf("type %T is empty", r)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", r, fields)
	}
	if r.MessageType != "" {
		field := fields[0]
		if r.MessageType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				r,
				r.MessageType,
				r,
			)
		}
	}
	return nil
}
