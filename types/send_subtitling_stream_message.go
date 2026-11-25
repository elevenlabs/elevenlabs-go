package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendSubtitlingStreamMessage struct {
	MessageType     string
	EndOfStream     *SubtitlingEndOfStream
	InputAudioChunk *SubtitlingInputAudioChunk
}

func (s *SendSubtitlingStreamMessage) GetEndOfStream() *SubtitlingEndOfStream {
	if s == nil {
		return nil
	}
	if s.EndOfStream == nil {
		return nil
	}
	return s.EndOfStream
}

func (s *SendSubtitlingStreamMessage) GetInputAudioChunk() *SubtitlingInputAudioChunk {
	if s == nil {
		return nil
	}
	if s.InputAudioChunk == nil {
		return nil
	}
	return s.InputAudioChunk
}

func (s *SendSubtitlingStreamMessage) GetMessageType() string {
	if s == nil {
		return ""
	}
	return s.MessageType
}

func (s *SendSubtitlingStreamMessage) UnmarshalJSON(data []byte) error {
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
		value := new(SubtitlingEndOfStream)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.EndOfStream = value
	case "input_audio_chunk":
		value := new(SubtitlingInputAudioChunk)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.InputAudioChunk = value
	}
	return nil
}

func (s SendSubtitlingStreamMessage) MarshalJSON() ([]byte, error) {
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

type SendSubtitlingStreamMessageVisitor interface {
	VisitEndOfStream(*SubtitlingEndOfStream) error
	VisitInputAudioChunk(*SubtitlingInputAudioChunk) error
}

func (s *SendSubtitlingStreamMessage) Accept(visitor SendSubtitlingStreamMessageVisitor) error {
	if s.EndOfStream != nil {
		return visitor.VisitEndOfStream(s.EndOfStream)
	}
	if s.InputAudioChunk != nil {
		return visitor.VisitInputAudioChunk(s.InputAudioChunk)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SendSubtitlingStreamMessage) validate() error {
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
