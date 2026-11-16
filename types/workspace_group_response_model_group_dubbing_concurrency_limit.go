package types

import (
	"encoding/json"
	"fmt"
)

type WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit struct {
	Integer                int
	UnlimitedStringLiteral string
	typ                    string
}

func (w *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) GetInteger() int {
	if w == nil {
		return 0
	}
	return w.Integer
}

func (w *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) GetUnlimitedStringLiteral() string {
	if w == nil {
		return ""
	}
	return w.UnlimitedStringLiteral
}

func (w *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) UnmarshalJSON(data []byte) error {
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

func (w WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if w.typ == "Integer" || w.Integer != 0 {
		return json.Marshal(w.Integer)
	}
	if w.typ == "UnlimitedStringLiteral" || w.UnlimitedStringLiteral != "" {
		return json.Marshal(w.UnlimitedStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WorkspaceGroupResponseModelGroupDubbingConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitUnlimitedStringLiteral(string) error
}

func (w *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) Accept(visitor WorkspaceGroupResponseModelGroupDubbingConcurrencyLimitVisitor) error {
	if w.typ == "Integer" || w.Integer != 0 {
		return visitor.VisitInteger(w.Integer)
	}
	if w.typ == "UnlimitedStringLiteral" || w.UnlimitedStringLiteral != "" {
		return visitor.VisitUnlimitedStringLiteral(w.UnlimitedStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}

func NewWorkspaceGroupResponseModelGroupDubbingConcurrencyLimitWithUnlimitedStringLiteral() *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit {
	return &WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit{typ: "UnlimitedStringLiteral", UnlimitedStringLiteral: "unlimited"}
}
