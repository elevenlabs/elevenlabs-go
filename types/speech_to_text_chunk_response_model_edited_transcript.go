package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SpeechToTextChunkResponseModelEditedTranscript struct {
	Kind       string
	Error      *TranscriptEditError
	Transcript *EditedTranscriptResponse
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) GetError() *TranscriptEditError {
	if s == nil {
		return nil
	}
	if s.Error == nil {
		return nil
	}
	return s.Error
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) GetTranscript() *EditedTranscriptResponse {
	if s == nil {
		return nil
	}
	if s.Transcript == nil {
		return nil
	}
	return s.Transcript
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) GetKind() string {
	if s == nil {
		return ""
	}
	return s.Kind
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", s)
	}
	switch unmarshaler.Kind {
	case "error":
		value := new(TranscriptEditError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Error = value
	case "transcript":
		value := new(EditedTranscriptResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Transcript = value
	}
	return nil
}

func (s SpeechToTextChunkResponseModelEditedTranscript) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Error != nil {
		return internal.MarshalJSONWithDiscriminant(s.Error, "kind", "error")
	}
	if s.Transcript != nil {
		return internal.MarshalJSONWithDiscriminant(s.Transcript, "kind", "transcript")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SpeechToTextChunkResponseModelEditedTranscriptVisitor interface {
	VisitError(*TranscriptEditError) error
	VisitTranscript(*EditedTranscriptResponse) error
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) Accept(visitor SpeechToTextChunkResponseModelEditedTranscriptVisitor) error {
	if s.Error != nil {
		return visitor.VisitError(s.Error)
	}
	if s.Transcript != nil {
		return visitor.VisitTranscript(s.Transcript)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SpeechToTextChunkResponseModelEditedTranscript) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Error != nil {
		fields = append(fields, "error")
	}
	if s.Transcript != nil {
		fields = append(fields, "transcript")
	}
	if len(fields) == 0 {
		if s.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Kind)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Kind != "" {
		field := fields[0]
		if s.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Kind,
				s,
			)
		}
	}
	return nil
}
