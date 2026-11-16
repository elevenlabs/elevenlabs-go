package types

import (
	"encoding/json"
	"fmt"
)

// Maximum number of credits that the credit limit can be exceeded by. Managed by the workspace admin. `"unlimited"` means no cap, `0` means usage-based billing is disabled.
type SubscriptionResponseMaxCreditLimitExtension struct {
	Integer                int
	UnlimitedStringLiteral string
	typ                    string
}

func (s *SubscriptionResponseMaxCreditLimitExtension) GetInteger() int {
	if s == nil {
		return 0
	}
	return s.Integer
}

func (s *SubscriptionResponseMaxCreditLimitExtension) GetUnlimitedStringLiteral() string {
	if s == nil {
		return ""
	}
	return s.UnlimitedStringLiteral
}

func (s *SubscriptionResponseMaxCreditLimitExtension) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		s.typ = "Integer"
		s.Integer = valueInteger
		return nil
	}
	var valueUnlimitedStringLiteral string
	if err := json.Unmarshal(data, &valueUnlimitedStringLiteral); err == nil {
		s.typ = "UnlimitedStringLiteral"
		s.UnlimitedStringLiteral = valueUnlimitedStringLiteral
		if s.UnlimitedStringLiteral != "unlimited" {
			return fmt.Errorf("unexpected value for literal on type %T; expected unlimited got %v", s, s.UnlimitedStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SubscriptionResponseMaxCreditLimitExtension) MarshalJSON() ([]byte, error) {
	if s.typ == "Integer" || s.Integer != 0 {
		return json.Marshal(s.Integer)
	}
	if s.typ == "UnlimitedStringLiteral" || s.UnlimitedStringLiteral != "" {
		return json.Marshal(s.UnlimitedStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SubscriptionResponseMaxCreditLimitExtensionVisitor interface {
	VisitInteger(int) error
	VisitUnlimitedStringLiteral(string) error
}

func (s *SubscriptionResponseMaxCreditLimitExtension) Accept(visitor SubscriptionResponseMaxCreditLimitExtensionVisitor) error {
	if s.typ == "Integer" || s.Integer != 0 {
		return visitor.VisitInteger(s.Integer)
	}
	if s.typ == "UnlimitedStringLiteral" || s.UnlimitedStringLiteral != "" {
		return visitor.VisitUnlimitedStringLiteral(s.UnlimitedStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}

func NewSubscriptionResponseMaxCreditLimitExtensionWithUnlimitedStringLiteral() *SubscriptionResponseMaxCreditLimitExtension {
	return &SubscriptionResponseMaxCreditLimitExtension{typ: "UnlimitedStringLiteral", UnlimitedStringLiteral: "unlimited"}
}
