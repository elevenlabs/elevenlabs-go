package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SipTransferInputConfig struct {
	Type                      string
	SipColdConferenceTransfer *SipColdConferenceTransferConfigInput
	SipColdReferTransfer      *SipColdReferTransferConfigInput
}

func (s *SipTransferInputConfig) GetSipColdConferenceTransfer() *SipColdConferenceTransferConfigInput {
	if s == nil {
		return nil
	}
	if s.SipColdConferenceTransfer == nil {
		return nil
	}
	return s.SipColdConferenceTransfer
}

func (s *SipTransferInputConfig) GetSipColdReferTransfer() *SipColdReferTransferConfigInput {
	if s == nil {
		return nil
	}
	if s.SipColdReferTransfer == nil {
		return nil
	}
	return s.SipColdReferTransfer
}

func (s *SipTransferInputConfig) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SipTransferInputConfig) UnmarshalJSON(data []byte) error {
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
	case "sip_cold_conference_transfer":
		value := new(SipColdConferenceTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SipColdConferenceTransfer = value
	case "sip_cold_refer_transfer":
		value := new(SipColdReferTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SipColdReferTransfer = value
	}
	return nil
}

func (s SipTransferInputConfig) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.SipColdConferenceTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(s.SipColdConferenceTransfer, "type", "sip_cold_conference_transfer")
	}
	if s.SipColdReferTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(s.SipColdReferTransfer, "type", "sip_cold_refer_transfer")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SipTransferInputConfigVisitor interface {
	VisitSipColdConferenceTransfer(*SipColdConferenceTransferConfigInput) error
	VisitSipColdReferTransfer(*SipColdReferTransferConfigInput) error
}

func (s *SipTransferInputConfig) Accept(visitor SipTransferInputConfigVisitor) error {
	if s.SipColdConferenceTransfer != nil {
		return visitor.VisitSipColdConferenceTransfer(s.SipColdConferenceTransfer)
	}
	if s.SipColdReferTransfer != nil {
		return visitor.VisitSipColdReferTransfer(s.SipColdReferTransfer)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SipTransferInputConfig) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.SipColdConferenceTransfer != nil {
		fields = append(fields, "sip_cold_conference_transfer")
	}
	if s.SipColdReferTransfer != nil {
		fields = append(fields, "sip_cold_refer_transfer")
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
