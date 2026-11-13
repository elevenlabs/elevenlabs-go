package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendUpstreamMessage struct {
	Type          string
	AgentResponse *AgentResponse
	Pong          *Pong
}

func (s *SendUpstreamMessage) GetAgentResponse() *AgentResponse {
	if s == nil {
		return nil
	}
	if s.AgentResponse == nil {
		return nil
	}
	return s.AgentResponse
}

func (s *SendUpstreamMessage) GetPong() *Pong {
	if s == nil {
		return nil
	}
	if s.Pong == nil {
		return nil
	}
	return s.Pong
}

func (s *SendUpstreamMessage) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SendUpstreamMessage) UnmarshalJSON(data []byte) error {
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
	case "agent_response":
		value := new(AgentResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.AgentResponse = value
	case "pong":
		value := new(Pong)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Pong = value
	}
	return nil
}

func (s SendUpstreamMessage) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.AgentResponse != nil {
		return internal.MarshalJSONWithDiscriminant(s.AgentResponse, "type", "agent_response")
	}
	if s.Pong != nil {
		return internal.MarshalJSONWithDiscriminant(s.Pong, "type", "pong")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SendUpstreamMessageVisitor interface {
	VisitAgentResponse(*AgentResponse) error
	VisitPong(*Pong) error
}

func (s *SendUpstreamMessage) Accept(visitor SendUpstreamMessageVisitor) error {
	if s.AgentResponse != nil {
		return visitor.VisitAgentResponse(s.AgentResponse)
	}
	if s.Pong != nil {
		return visitor.VisitPong(s.Pong)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SendUpstreamMessage) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.AgentResponse != nil {
		fields = append(fields, "agent_response")
	}
	if s.Pong != nil {
		fields = append(fields, "pong")
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
