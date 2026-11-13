package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ReceiveTranslateStreamMessage struct {
	MessageType       string
	Audio             *TranslateAudio
	AudioOutputError  *TranslateAudioOutputError
	AuthError         *TranslateAuthError
	Error             *TranslateError
	FinalTranscript   *TranslateFinalTranscript
	InputError        *TranslateInputError
	PartialTranscript *TranslatePartialTranscript
	PaymentIssue      *TranslatePaymentIssueError
	QueueOverflow     *TranslateQueueOverflowError
	QuotaExceeded     *TranslateQuotaExceededError
	RateLimited       *TranslateRateLimitedError
	SessionStarted    *TranslateSessionStarted
	Status            *TranslateStatus
	TranscriberError  *TranslateTranscriberError
	Translation       *TranslateTranslation
	TranslationError  *TranslateTranslationError
	UnacceptedTerms   *TranslateUnacceptedTermsError
}

func (r *ReceiveTranslateStreamMessage) GetAudio() *TranslateAudio {
	if r == nil {
		return nil
	}
	if r.Audio == nil {
		return nil
	}
	return r.Audio
}

func (r *ReceiveTranslateStreamMessage) GetAudioOutputError() *TranslateAudioOutputError {
	if r == nil {
		return nil
	}
	if r.AudioOutputError == nil {
		return nil
	}
	return r.AudioOutputError
}

func (r *ReceiveTranslateStreamMessage) GetAuthError() *TranslateAuthError {
	if r == nil {
		return nil
	}
	if r.AuthError == nil {
		return nil
	}
	return r.AuthError
}

func (r *ReceiveTranslateStreamMessage) GetError() *TranslateError {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveTranslateStreamMessage) GetFinalTranscript() *TranslateFinalTranscript {
	if r == nil {
		return nil
	}
	if r.FinalTranscript == nil {
		return nil
	}
	return r.FinalTranscript
}

func (r *ReceiveTranslateStreamMessage) GetInputError() *TranslateInputError {
	if r == nil {
		return nil
	}
	if r.InputError == nil {
		return nil
	}
	return r.InputError
}

func (r *ReceiveTranslateStreamMessage) GetPartialTranscript() *TranslatePartialTranscript {
	if r == nil {
		return nil
	}
	if r.PartialTranscript == nil {
		return nil
	}
	return r.PartialTranscript
}

func (r *ReceiveTranslateStreamMessage) GetPaymentIssue() *TranslatePaymentIssueError {
	if r == nil {
		return nil
	}
	if r.PaymentIssue == nil {
		return nil
	}
	return r.PaymentIssue
}

func (r *ReceiveTranslateStreamMessage) GetQueueOverflow() *TranslateQueueOverflowError {
	if r == nil {
		return nil
	}
	if r.QueueOverflow == nil {
		return nil
	}
	return r.QueueOverflow
}

func (r *ReceiveTranslateStreamMessage) GetQuotaExceeded() *TranslateQuotaExceededError {
	if r == nil {
		return nil
	}
	if r.QuotaExceeded == nil {
		return nil
	}
	return r.QuotaExceeded
}

func (r *ReceiveTranslateStreamMessage) GetRateLimited() *TranslateRateLimitedError {
	if r == nil {
		return nil
	}
	if r.RateLimited == nil {
		return nil
	}
	return r.RateLimited
}

func (r *ReceiveTranslateStreamMessage) GetSessionStarted() *TranslateSessionStarted {
	if r == nil {
		return nil
	}
	if r.SessionStarted == nil {
		return nil
	}
	return r.SessionStarted
}

func (r *ReceiveTranslateStreamMessage) GetStatus() *TranslateStatus {
	if r == nil {
		return nil
	}
	if r.Status == nil {
		return nil
	}
	return r.Status
}

func (r *ReceiveTranslateStreamMessage) GetTranscriberError() *TranslateTranscriberError {
	if r == nil {
		return nil
	}
	if r.TranscriberError == nil {
		return nil
	}
	return r.TranscriberError
}

func (r *ReceiveTranslateStreamMessage) GetTranslation() *TranslateTranslation {
	if r == nil {
		return nil
	}
	if r.Translation == nil {
		return nil
	}
	return r.Translation
}

func (r *ReceiveTranslateStreamMessage) GetTranslationError() *TranslateTranslationError {
	if r == nil {
		return nil
	}
	if r.TranslationError == nil {
		return nil
	}
	return r.TranslationError
}

func (r *ReceiveTranslateStreamMessage) GetUnacceptedTerms() *TranslateUnacceptedTermsError {
	if r == nil {
		return nil
	}
	if r.UnacceptedTerms == nil {
		return nil
	}
	return r.UnacceptedTerms
}

func (r *ReceiveTranslateStreamMessage) GetMessageType() string {
	if r == nil {
		return ""
	}
	return r.MessageType
}

func (r *ReceiveTranslateStreamMessage) UnmarshalJSON(data []byte) error {
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
	case "audio":
		value := new(TranslateAudio)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Audio = value
	case "audio_output_error":
		value := new(TranslateAudioOutputError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.AudioOutputError = value
	case "auth_error":
		value := new(TranslateAuthError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.AuthError = value
	case "error":
		value := new(TranslateError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "final_transcript":
		value := new(TranslateFinalTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.FinalTranscript = value
	case "input_error":
		value := new(TranslateInputError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.InputError = value
	case "partial_transcript":
		value := new(TranslatePartialTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PartialTranscript = value
	case "payment_issue":
		value := new(TranslatePaymentIssueError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PaymentIssue = value
	case "queue_overflow":
		value := new(TranslateQueueOverflowError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.QueueOverflow = value
	case "quota_exceeded":
		value := new(TranslateQuotaExceededError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.QuotaExceeded = value
	case "rate_limited":
		value := new(TranslateRateLimitedError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.RateLimited = value
	case "session_started":
		value := new(TranslateSessionStarted)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.SessionStarted = value
	case "status":
		value := new(TranslateStatus)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Status = value
	case "transcriber_error":
		value := new(TranslateTranscriberError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranscriberError = value
	case "translation":
		value := new(TranslateTranslation)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Translation = value
	case "translation_error":
		value := new(TranslateTranslationError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranslationError = value
	case "unaccepted_terms":
		value := new(TranslateUnacceptedTermsError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.UnacceptedTerms = value
	}
	return nil
}

func (r ReceiveTranslateStreamMessage) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.Audio != nil {
		return internal.MarshalJSONWithDiscriminant(r.Audio, "message_type", "audio")
	}
	if r.AudioOutputError != nil {
		return internal.MarshalJSONWithDiscriminant(r.AudioOutputError, "message_type", "audio_output_error")
	}
	if r.AuthError != nil {
		return internal.MarshalJSONWithDiscriminant(r.AuthError, "message_type", "auth_error")
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
	if r.Translation != nil {
		return internal.MarshalJSONWithDiscriminant(r.Translation, "message_type", "translation")
	}
	if r.TranslationError != nil {
		return internal.MarshalJSONWithDiscriminant(r.TranslationError, "message_type", "translation_error")
	}
	if r.UnacceptedTerms != nil {
		return internal.MarshalJSONWithDiscriminant(r.UnacceptedTerms, "message_type", "unaccepted_terms")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type ReceiveTranslateStreamMessageVisitor interface {
	VisitAudio(*TranslateAudio) error
	VisitAudioOutputError(*TranslateAudioOutputError) error
	VisitAuthError(*TranslateAuthError) error
	VisitError(*TranslateError) error
	VisitFinalTranscript(*TranslateFinalTranscript) error
	VisitInputError(*TranslateInputError) error
	VisitPartialTranscript(*TranslatePartialTranscript) error
	VisitPaymentIssue(*TranslatePaymentIssueError) error
	VisitQueueOverflow(*TranslateQueueOverflowError) error
	VisitQuotaExceeded(*TranslateQuotaExceededError) error
	VisitRateLimited(*TranslateRateLimitedError) error
	VisitSessionStarted(*TranslateSessionStarted) error
	VisitStatus(*TranslateStatus) error
	VisitTranscriberError(*TranslateTranscriberError) error
	VisitTranslation(*TranslateTranslation) error
	VisitTranslationError(*TranslateTranslationError) error
	VisitUnacceptedTerms(*TranslateUnacceptedTermsError) error
}

func (r *ReceiveTranslateStreamMessage) Accept(visitor ReceiveTranslateStreamMessageVisitor) error {
	if r.Audio != nil {
		return visitor.VisitAudio(r.Audio)
	}
	if r.AudioOutputError != nil {
		return visitor.VisitAudioOutputError(r.AudioOutputError)
	}
	if r.AuthError != nil {
		return visitor.VisitAuthError(r.AuthError)
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
	if r.Translation != nil {
		return visitor.VisitTranslation(r.Translation)
	}
	if r.TranslationError != nil {
		return visitor.VisitTranslationError(r.TranslationError)
	}
	if r.UnacceptedTerms != nil {
		return visitor.VisitUnacceptedTerms(r.UnacceptedTerms)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *ReceiveTranslateStreamMessage) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.Audio != nil {
		fields = append(fields, "audio")
	}
	if r.AudioOutputError != nil {
		fields = append(fields, "audio_output_error")
	}
	if r.AuthError != nil {
		fields = append(fields, "auth_error")
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
	if r.Translation != nil {
		fields = append(fields, "translation")
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
