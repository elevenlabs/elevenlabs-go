package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type StreamEnvelopeEvent struct {
	Type                       string
	AgentChatResponsePart      *AgentChatResponsePartClientEvent
	AgentReasoningResponsePart *AgentReasoningResponsePartClientEvent
	AgentResponse              *AgentResponseClientEvent
	AgentToolRequest           *AgentToolRequestStreamedEvent
	AgentToolResponse          *AgentToolResponseStreamedEvent
	ClientToolCall             *ClientToolCallEvent
	StreamEnded                *StreamEndedEvent
}

func (s *StreamEnvelopeEvent) GetAgentChatResponsePart() *AgentChatResponsePartClientEvent {
	if s == nil {
		return nil
	}
	if s.AgentChatResponsePart == nil {
		return nil
	}
	return s.AgentChatResponsePart
}

func (s *StreamEnvelopeEvent) GetAgentReasoningResponsePart() *AgentReasoningResponsePartClientEvent {
	if s == nil {
		return nil
	}
	if s.AgentReasoningResponsePart == nil {
		return nil
	}
	return s.AgentReasoningResponsePart
}

func (s *StreamEnvelopeEvent) GetAgentResponse() *AgentResponseClientEvent {
	if s == nil {
		return nil
	}
	if s.AgentResponse == nil {
		return nil
	}
	return s.AgentResponse
}

func (s *StreamEnvelopeEvent) GetAgentToolRequest() *AgentToolRequestStreamedEvent {
	if s == nil {
		return nil
	}
	if s.AgentToolRequest == nil {
		return nil
	}
	return s.AgentToolRequest
}

func (s *StreamEnvelopeEvent) GetAgentToolResponse() *AgentToolResponseStreamedEvent {
	if s == nil {
		return nil
	}
	if s.AgentToolResponse == nil {
		return nil
	}
	return s.AgentToolResponse
}

func (s *StreamEnvelopeEvent) GetClientToolCall() *ClientToolCallEvent {
	if s == nil {
		return nil
	}
	if s.ClientToolCall == nil {
		return nil
	}
	return s.ClientToolCall
}

func (s *StreamEnvelopeEvent) GetStreamEnded() *StreamEndedEvent {
	if s == nil {
		return nil
	}
	if s.StreamEnded == nil {
		return nil
	}
	return s.StreamEnded
}

func (s *StreamEnvelopeEvent) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *StreamEnvelopeEvent) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", s)
	}
	switch unmarshaler.Type {
	case "agent_chat_response_part":
		value := new(AgentChatResponsePartClientEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentChatResponsePart = value
	case "agent_reasoning_response_part":
		value := new(AgentReasoningResponsePartClientEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentReasoningResponsePart = value
	case "agent_response":
		value := new(AgentResponseClientEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentResponse = value
	case "agent_tool_request":
		value := new(AgentToolRequestStreamedEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentToolRequest = value
	case "agent_tool_response":
		value := new(AgentToolResponseStreamedEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentToolResponse = value
	case "client_tool_call":
		value := new(ClientToolCallEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ClientToolCall = value
	case "stream_ended":
		value := new(StreamEndedEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.StreamEnded = value
	}
	return nil
}

func (s StreamEnvelopeEvent) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.AgentChatResponsePart != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentChatResponsePart, "type", "agent_chat_response_part")
	}
	if s.AgentReasoningResponsePart != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentReasoningResponsePart, "type", "agent_reasoning_response_part")
	}
	if s.AgentResponse != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentResponse, "type", "agent_response")
	}
	if s.AgentToolRequest != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentToolRequest, "type", "agent_tool_request")
	}
	if s.AgentToolResponse != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentToolResponse, "type", "agent_tool_response")
	}
	if s.ClientToolCall != nil {
		return internal.MarshalJSONWithDiscriminant(s.ClientToolCall, "type", "client_tool_call")
	}
	if s.StreamEnded != nil {
		return internal.MarshalJSONWithDiscriminant(s.StreamEnded, "type", "stream_ended")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type StreamEnvelopeEventVisitor interface {
	VisitAgentChatResponsePart(*AgentChatResponsePartClientEvent) error
	VisitAgentReasoningResponsePart(*AgentReasoningResponsePartClientEvent) error
	VisitAgentResponse(*AgentResponseClientEvent) error
	VisitAgentToolRequest(*AgentToolRequestStreamedEvent) error
	VisitAgentToolResponse(*AgentToolResponseStreamedEvent) error
	VisitClientToolCall(*ClientToolCallEvent) error
	VisitStreamEnded(*StreamEndedEvent) error
}

func (s *StreamEnvelopeEvent) Accept(visitor StreamEnvelopeEventVisitor) error {
	if s.AgentChatResponsePart != nil {
		return visitor.VisitAgentChatResponsePart(s.AgentChatResponsePart)
	}
	if s.AgentReasoningResponsePart != nil {
		return visitor.VisitAgentReasoningResponsePart(s.AgentReasoningResponsePart)
	}
	if s.AgentResponse != nil {
		return visitor.VisitAgentResponse(s.AgentResponse)
	}
	if s.AgentToolRequest != nil {
		return visitor.VisitAgentToolRequest(s.AgentToolRequest)
	}
	if s.AgentToolResponse != nil {
		return visitor.VisitAgentToolResponse(s.AgentToolResponse)
	}
	if s.ClientToolCall != nil {
		return visitor.VisitClientToolCall(s.ClientToolCall)
	}
	if s.StreamEnded != nil {
		return visitor.VisitStreamEnded(s.StreamEnded)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *StreamEnvelopeEvent) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.AgentChatResponsePart != nil {
		fields = append(fields, "agent_chat_response_part")
	}
	if s.AgentReasoningResponsePart != nil {
		fields = append(fields, "agent_reasoning_response_part")
	}
	if s.AgentResponse != nil {
		fields = append(fields, "agent_response")
	}
	if s.AgentToolRequest != nil {
		fields = append(fields, "agent_tool_request")
	}
	if s.AgentToolResponse != nil {
		fields = append(fields, "agent_tool_response")
	}
	if s.ClientToolCall != nil {
		fields = append(fields, "client_tool_call")
	}
	if s.StreamEnded != nil {
		fields = append(fields, "stream_ended")
	}
	if len(fields) == 0 {
		if s.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Type)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Type != "" {
		field := fields[0]
		if s.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Type,
				s,
			)
		}
	}
	return nil
}
