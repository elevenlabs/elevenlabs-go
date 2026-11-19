package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The pending change for the user.
type SubscriptionPendingChange struct {
	Kind         string
	Cancellation *PendingCancellationResponseModel
	Change       *PendingSubscriptionSwitchResponseModel
}

func (s *SubscriptionPendingChange) GetCancellation() *PendingCancellationResponseModel {
	if s == nil {
		return nil
	}
	if s.Cancellation == nil {
		return nil
	}
	return s.Cancellation
}

func (s *SubscriptionPendingChange) GetChange() *PendingSubscriptionSwitchResponseModel {
	if s == nil {
		return nil
	}
	if s.Change == nil {
		return nil
	}
	return s.Change
}

func (s *SubscriptionPendingChange) GetKind() string {
	if s == nil {
		return ""
	}
	return s.Kind
}

func (s *SubscriptionPendingChange) UnmarshalJSON(data []byte) error {
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
	case "cancellation":
		value := new(PendingCancellationResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Cancellation = value
	case "change":
		value := new(PendingSubscriptionSwitchResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Change = value
	}
	return nil
}

func (s SubscriptionPendingChange) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Cancellation != nil {
		return internal.MarshalJSONWithDiscriminant(s.Cancellation, "kind", "cancellation")
	}
	if s.Change != nil {
		return internal.MarshalJSONWithDiscriminant(s.Change, "kind", "change")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SubscriptionPendingChangeVisitor interface {
	VisitCancellation(*PendingCancellationResponseModel) error
	VisitChange(*PendingSubscriptionSwitchResponseModel) error
}

func (s *SubscriptionPendingChange) Accept(visitor SubscriptionPendingChangeVisitor) error {
	if s.Cancellation != nil {
		return visitor.VisitCancellation(s.Cancellation)
	}
	if s.Change != nil {
		return visitor.VisitChange(s.Change)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SubscriptionPendingChange) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Cancellation != nil {
		fields = append(fields, "cancellation")
	}
	if s.Change != nil {
		fields = append(fields, "change")
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
