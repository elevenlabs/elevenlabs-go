package types

import (
	"encoding/json"
	"fmt"
)

// Maximum number of credits that the credit limit can be exceeded by. Managed by the workspace admin. `"unlimited"` means no cap, `0` means usage-based billing is disabled.
type SubscriptionMaxCreditLimitExtension struct {
	Integer                int
	UnlimitedStringLiteral string
	typ                    string
}

func (subscriptionMaxCreditLimitExtension *SubscriptionMaxCreditLimitExtension) GetInteger() int {
	if subscriptionMaxCreditLimitExtension == nil {
		return 0
	}
	return subscriptionMaxCreditLimitExtension.Integer
}

func (subscriptionMaxCreditLimitExtension *SubscriptionMaxCreditLimitExtension) GetUnlimitedStringLiteral() string {
	if subscriptionMaxCreditLimitExtension == nil {
		return ""
	}
	return subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral
}

func (subscriptionMaxCreditLimitExtension *SubscriptionMaxCreditLimitExtension) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		subscriptionMaxCreditLimitExtension.typ = "Integer"
		subscriptionMaxCreditLimitExtension.Integer = valueInteger
		return nil
	}
	var valueUnlimitedStringLiteral string
	if err := json.Unmarshal(data, &valueUnlimitedStringLiteral); err == nil {
		subscriptionMaxCreditLimitExtension.typ = "UnlimitedStringLiteral"
		subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral = valueUnlimitedStringLiteral
		if subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral != "unlimited" {
			return fmt.Errorf("unexpected value for literal on type %T; expected unlimited got %v", subscriptionMaxCreditLimitExtension, subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, subscriptionMaxCreditLimitExtension)
}

func (subscriptionMaxCreditLimitExtension SubscriptionMaxCreditLimitExtension) MarshalJSON() ([]byte, error) {
	if subscriptionMaxCreditLimitExtension.typ == "Integer" || subscriptionMaxCreditLimitExtension.Integer != 0 {
		return json.Marshal(subscriptionMaxCreditLimitExtension.Integer)
	}
	if subscriptionMaxCreditLimitExtension.typ == "UnlimitedStringLiteral" || subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral != "" {
		return json.Marshal(subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", subscriptionMaxCreditLimitExtension)
}

type SubscriptionMaxCreditLimitExtensionVisitor interface {
	VisitInteger(int) error
	VisitUnlimitedStringLiteral(string) error
}

func (subscriptionMaxCreditLimitExtension *SubscriptionMaxCreditLimitExtension) Accept(visitor SubscriptionMaxCreditLimitExtensionVisitor) error {
	if subscriptionMaxCreditLimitExtension.typ == "Integer" || subscriptionMaxCreditLimitExtension.Integer != 0 {
		return visitor.VisitInteger(subscriptionMaxCreditLimitExtension.Integer)
	}
	if subscriptionMaxCreditLimitExtension.typ == "UnlimitedStringLiteral" || subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral != "" {
		return visitor.VisitUnlimitedStringLiteral(subscriptionMaxCreditLimitExtension.UnlimitedStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", subscriptionMaxCreditLimitExtension)
}

func NewSubscriptionMaxCreditLimitExtensionWithUnlimitedStringLiteral() *SubscriptionMaxCreditLimitExtension {
	return &SubscriptionMaxCreditLimitExtension{typ: "UnlimitedStringLiteral", UnlimitedStringLiteral: "unlimited"}
}
