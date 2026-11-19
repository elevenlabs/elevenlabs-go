package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ReceiveDubbingStreamMessage struct {
	MessageType         string
	Audio               *DubbingAudio
	AudioOutputError    *DubbingAudioOutputError
	AuthError           *DubbingAuthError
	CommittedTranscript *DubbingCommittedTranscript
	Error               *DubbingErrorPayload
	InputError          *DubbingInputError
	InvalidRequest      *DubbingInvalidRequestError
	PartialTranscript   *DubbingPartialTranscript
	QueueOverflow       *DubbingQueueOverflowError
	RateLimited         *DubbingRateLimitedError
	SessionStarted      *DubbingSessionStarted
	Status              *DubbingStatus
	TranscriberError    *DubbingTranscriberError
	Translation         *DubbingTranslation
	TranslationError    *DubbingTranslationError
}

func (r *ReceiveDubbingStreamMessage) GetAudio() *DubbingAudio {
	if r == nil {
		return nil
	}
	if r.Audio == nil {
		return nil
	}
	return r.Audio
}

func (r *ReceiveDubbingStreamMessage) GetAudioOutputError() *DubbingAudioOutputError {
	if r == nil {
		return nil
	}
	if r.AudioOutputError == nil {
		return nil
	}
	return r.AudioOutputError
}

func (r *ReceiveDubbingStreamMessage) GetAuthError() *DubbingAuthError {
	if r == nil {
		return nil
	}
	if r.AuthError == nil {
		return nil
	}
	return r.AuthError
}

func (r *ReceiveDubbingStreamMessage) GetCommittedTranscript() *DubbingCommittedTranscript {
	if r == nil {
		return nil
	}
	if r.CommittedTranscript == nil {
		return nil
	}
	return r.CommittedTranscript
}

func (r *ReceiveDubbingStreamMessage) GetError() *DubbingErrorPayload {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveDubbingStreamMessage) GetInputError() *DubbingInputError {
	if r == nil {
		return nil
	}
	if r.InputError == nil {
		return nil
	}
	return r.InputError
}

func (r *ReceiveDubbingStreamMessage) GetInvalidRequest() *DubbingInvalidRequestError {
	if r == nil {
		return nil
	}
	if r.InvalidRequest == nil {
		return nil
	}
	return r.InvalidRequest
}

func (r *ReceiveDubbingStreamMessage) GetPartialTranscript() *DubbingPartialTranscript {
	if r == nil {
		return nil
	}
	if r.PartialTranscript == nil {
		return nil
	}
	return r.PartialTranscript
}

func (r *ReceiveDubbingStreamMessage) GetQueueOverflow() *DubbingQueueOverflowError {
	if r == nil {
		return nil
	}
	if r.QueueOverflow == nil {
		return nil
	}
	return r.QueueOverflow
}

func (r *ReceiveDubbingStreamMessage) GetRateLimited() *DubbingRateLimitedError {
	if r == nil {
		return nil
	}
	if r.RateLimited == nil {
		return nil
	}
	return r.RateLimited
}

func (r *ReceiveDubbingStreamMessage) GetSessionStarted() *DubbingSessionStarted {
	if r == nil {
		return nil
	}
	if r.SessionStarted == nil {
		return nil
	}
	return r.SessionStarted
}

func (r *ReceiveDubbingStreamMessage) GetStatus() *DubbingStatus {
	if r == nil {
		return nil
	}
	if r.Status == nil {
		return nil
	}
	return r.Status
}

func (r *ReceiveDubbingStreamMessage) GetTranscriberError() *DubbingTranscriberError {
	if r == nil {
		return nil
	}
	if r.TranscriberError == nil {
		return nil
	}
	return r.TranscriberError
}

func (r *ReceiveDubbingStreamMessage) GetTranslation() *DubbingTranslation {
	if r == nil {
		return nil
	}
	if r.Translation == nil {
		return nil
	}
	return r.Translation
}

func (r *ReceiveDubbingStreamMessage) GetTranslationError() *DubbingTranslationError {
	if r == nil {
		return nil
	}
	if r.TranslationError == nil {
		return nil
	}
	return r.TranslationError
}

func (r *ReceiveDubbingStreamMessage) GetMessageType() string {
	if r == nil {
		return ""
	}
	return r.MessageType
}

func (r *ReceiveDubbingStreamMessage) UnmarshalJSON(data []byte) error {
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
		value := new(DubbingAudio)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Audio = value
	case "audio_output_error":
		value := new(DubbingAudioOutputError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.AudioOutputError = value
	case "auth_error":
		value := new(DubbingAuthError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.AuthError = value
	case "committed_transcript":
		value := new(DubbingCommittedTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.CommittedTranscript = value
	case "error":
		value := new(DubbingErrorPayload)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "input_error":
		value := new(DubbingInputError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.InputError = value
	case "invalid_request":
		value := new(DubbingInvalidRequestError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.InvalidRequest = value
	case "partial_transcript":
		value := new(DubbingPartialTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.PartialTranscript = value
	case "queue_overflow":
		value := new(DubbingQueueOverflowError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.QueueOverflow = value
	case "rate_limited":
		value := new(DubbingRateLimitedError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.RateLimited = value
	case "session_started":
		value := new(DubbingSessionStarted)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.SessionStarted = value
	case "status":
		value := new(DubbingStatus)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Status = value
	case "transcriber_error":
		value := new(DubbingTranscriberError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranscriberError = value
	case "translation":
		value := new(DubbingTranslation)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Translation = value
	case "translation_error":
		value := new(DubbingTranslationError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TranslationError = value
	}
	return nil
}

func (r ReceiveDubbingStreamMessage) MarshalJSON() ([]byte, error) {
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
	if r.CommittedTranscript != nil {
		return internal.MarshalJSONWithDiscriminant(r.CommittedTranscript, "message_type", "committed_transcript")
	}
	if r.Error != nil {
		return internal.MarshalJSONWithDiscriminant(r.Error, "message_type", "error")
	}
	if r.InputError != nil {
		return internal.MarshalJSONWithDiscriminant(r.InputError, "message_type", "input_error")
	}
	if r.InvalidRequest != nil {
		return internal.MarshalJSONWithDiscriminant(r.InvalidRequest, "message_type", "invalid_request")
	}
	if r.PartialTranscript != nil {
		return internal.MarshalJSONWithDiscriminant(r.PartialTranscript, "message_type", "partial_transcript")
	}
	if r.QueueOverflow != nil {
		return internal.MarshalJSONWithDiscriminant(r.QueueOverflow, "message_type", "queue_overflow")
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
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type ReceiveDubbingStreamMessageVisitor interface {
	VisitAudio(*DubbingAudio) error
	VisitAudioOutputError(*DubbingAudioOutputError) error
	VisitAuthError(*DubbingAuthError) error
	VisitCommittedTranscript(*DubbingCommittedTranscript) error
	VisitError(*DubbingErrorPayload) error
	VisitInputError(*DubbingInputError) error
	VisitInvalidRequest(*DubbingInvalidRequestError) error
	VisitPartialTranscript(*DubbingPartialTranscript) error
	VisitQueueOverflow(*DubbingQueueOverflowError) error
	VisitRateLimited(*DubbingRateLimitedError) error
	VisitSessionStarted(*DubbingSessionStarted) error
	VisitStatus(*DubbingStatus) error
	VisitTranscriberError(*DubbingTranscriberError) error
	VisitTranslation(*DubbingTranslation) error
	VisitTranslationError(*DubbingTranslationError) error
}

func (r *ReceiveDubbingStreamMessage) Accept(visitor ReceiveDubbingStreamMessageVisitor) error {
	if r.Audio != nil {
		return visitor.VisitAudio(r.Audio)
	}
	if r.AudioOutputError != nil {
		return visitor.VisitAudioOutputError(r.AudioOutputError)
	}
	if r.AuthError != nil {
		return visitor.VisitAuthError(r.AuthError)
	}
	if r.CommittedTranscript != nil {
		return visitor.VisitCommittedTranscript(r.CommittedTranscript)
	}
	if r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	if r.InputError != nil {
		return visitor.VisitInputError(r.InputError)
	}
	if r.InvalidRequest != nil {
		return visitor.VisitInvalidRequest(r.InvalidRequest)
	}
	if r.PartialTranscript != nil {
		return visitor.VisitPartialTranscript(r.PartialTranscript)
	}
	if r.QueueOverflow != nil {
		return visitor.VisitQueueOverflow(r.QueueOverflow)
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
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *ReceiveDubbingStreamMessage) validate() error {
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
	if r.CommittedTranscript != nil {
		fields = append(fields, "committed_transcript")
	}
	if r.Error != nil {
		fields = append(fields, "error")
	}
	if r.InputError != nil {
		fields = append(fields, "input_error")
	}
	if r.InvalidRequest != nil {
		fields = append(fields, "invalid_request")
	}
	if r.PartialTranscript != nil {
		fields = append(fields, "partial_transcript")
	}
	if r.QueueOverflow != nil {
		fields = append(fields, "queue_overflow")
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
