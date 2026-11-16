package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendDubbingStreamMessage struct {
	MessageType     string
	EndOfStream     *DubbingEndOfStream
	InputAudioChunk *DubbingInputAudioChunk
}

func (s *SendDubbingStreamMessage) GetEndOfStream() *DubbingEndOfStream {
	if s == nil {
		return nil
	}
	if s.EndOfStream == nil {
		return nil
	}
	return s.EndOfStream
}

func (s *SendDubbingStreamMessage) GetInputAudioChunk() *DubbingInputAudioChunk {
	if s == nil {
		return nil
	}
	if s.InputAudioChunk == nil {
		return nil
	}
	return s.InputAudioChunk
}

func (s *SendDubbingStreamMessage) GetMessageType() string {
	if s == nil {
		return ""
	}
	return s.MessageType
}

func (s *SendDubbingStreamMessage) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		MessageType string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.MessageType = unmarshaler.MessageType
	if unmarshaler.MessageType == "" {
		return fmt.Errorf("%T did not include discriminant message_type", s)
	}
	switch unmarshaler.MessageType {
	case "end_of_stream":
		value := new(DubbingEndOfStream)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.EndOfStream = value
	case "input_audio_chunk":
		value := new(DubbingInputAudioChunk)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.InputAudioChunk = value
	}
	return nil
}

func (s SendDubbingStreamMessage) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.EndOfStream != nil {
		return internal.MarshalJSONWithDiscriminant(s.EndOfStream, "message_type", "end_of_stream")
	}
	if s.InputAudioChunk != nil {
		return internal.MarshalJSONWithDiscriminant(s.InputAudioChunk, "message_type", "input_audio_chunk")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SendDubbingStreamMessageVisitor interface {
	VisitEndOfStream(*DubbingEndOfStream) error
	VisitInputAudioChunk(*DubbingInputAudioChunk) error
}

func (s *SendDubbingStreamMessage) Accept(visitor SendDubbingStreamMessageVisitor) error {
	if s.EndOfStream != nil {
		return visitor.VisitEndOfStream(s.EndOfStream)
	}
	if s.InputAudioChunk != nil {
		return visitor.VisitInputAudioChunk(s.InputAudioChunk)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SendDubbingStreamMessage) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.EndOfStream != nil {
		fields = append(fields, "end_of_stream")
	}
	if s.InputAudioChunk != nil {
		fields = append(fields, "input_audio_chunk")
	}
	if len(fields) == 0 {
		if s.MessageType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.MessageType)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.MessageType != "" {
		field := fields[0]
		if s.MessageType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.MessageType,
				s,
			)
		}
	}
	return nil
}
