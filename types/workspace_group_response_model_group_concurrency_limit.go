package types

import (
	"encoding/json"
	"fmt"
)

type WorkspaceGroupResponseModelGroupConcurrencyLimit struct {
	Integer                int
	UnlimitedStringLiteral string
	typ                    string
}

func (w *WorkspaceGroupResponseModelGroupConcurrencyLimit) GetInteger() int {
	if w == nil {
		return 0
	}
	return w.Integer
}

func (w *WorkspaceGroupResponseModelGroupConcurrencyLimit) GetUnlimitedStringLiteral() string {
	if w == nil {
		return ""
	}
	return w.UnlimitedStringLiteral
}

func (w *WorkspaceGroupResponseModelGroupConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		w.typ = "Integer"
		w.Integer = valueInteger
		return nil
	}
	var valueUnlimitedStringLiteral string
	if err := json.Unmarshal(data, &valueUnlimitedStringLiteral); err == nil {
		w.typ = "UnlimitedStringLiteral"
		w.UnlimitedStringLiteral = valueUnlimitedStringLiteral
		if w.UnlimitedStringLiteral != "unlimited" {
			return fmt.Errorf("unexpected value for literal on type %T; expected unlimited got %v", w, w.UnlimitedStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
}

func (w WorkspaceGroupResponseModelGroupConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if w.typ == "Integer" || w.Integer != 0 {
		return json.Marshal(w.Integer)
	}
	if w.typ == "UnlimitedStringLiteral" || w.UnlimitedStringLiteral != "" {
		return json.Marshal(w.UnlimitedStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WorkspaceGroupResponseModelGroupConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitUnlimitedStringLiteral(string) error
}

func (w *WorkspaceGroupResponseModelGroupConcurrencyLimit) Accept(visitor WorkspaceGroupResponseModelGroupConcurrencyLimitVisitor) error {
	if w.typ == "Integer" || w.Integer != 0 {
		return visitor.VisitInteger(w.Integer)
	}
	if w.typ == "UnlimitedStringLiteral" || w.UnlimitedStringLiteral != "" {
		return visitor.VisitUnlimitedStringLiteral(w.UnlimitedStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}

func NewWorkspaceGroupResponseModelGroupConcurrencyLimitWithUnlimitedStringLiteral() *WorkspaceGroupResponseModelGroupConcurrencyLimit {
	return &WorkspaceGroupResponseModelGroupConcurrencyLimit{typ: "UnlimitedStringLiteral", UnlimitedStringLiteral: "unlimited"}
}
