package types

import (
	"encoding/json"
	"fmt"
)

type StartProcedureToolConfigProceduresValue struct {
	ProcedureVersionRef *ProcedureVersionRef
	ProcedureDraftRef   *ProcedureDraftRef
	typ                 string
}

func (s *StartProcedureToolConfigProceduresValue) GetProcedureVersionRef() *ProcedureVersionRef {
	if s == nil {
		return nil
	}
	if s.ProcedureVersionRef == nil {
		return nil
	}
	return s.ProcedureVersionRef
}

func (s *StartProcedureToolConfigProceduresValue) GetProcedureDraftRef() *ProcedureDraftRef {
	if s == nil {
		return nil
	}
	if s.ProcedureDraftRef == nil {
		return nil
	}
	return s.ProcedureDraftRef
}

func (s *StartProcedureToolConfigProceduresValue) UnmarshalJSON(data []byte) error {
	var valueProcedureVersionRef *ProcedureVersionRef
	if err := json.Unmarshal(data, &valueProcedureVersionRef); err == nil {
		s.typ = "ProcedureVersionRef"
		s.ProcedureVersionRef = valueProcedureVersionRef
		return nil
	}
	var valueProcedureDraftRef *ProcedureDraftRef
	if err := json.Unmarshal(data, &valueProcedureDraftRef); err == nil {
		s.typ = "ProcedureDraftRef"
		s.ProcedureDraftRef = valueProcedureDraftRef
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s StartProcedureToolConfigProceduresValue) MarshalJSON() ([]byte, error) {
	if s.typ == "ProcedureVersionRef" || s.ProcedureVersionRef != nil {
		return json.Marshal(s.ProcedureVersionRef)
	}
	if s.typ == "ProcedureDraftRef" || s.ProcedureDraftRef != nil {
		return json.Marshal(s.ProcedureDraftRef)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type StartProcedureToolConfigProceduresValueVisitor interface {
	VisitProcedureVersionRef(*ProcedureVersionRef) error
	VisitProcedureDraftRef(*ProcedureDraftRef) error
}

func (s *StartProcedureToolConfigProceduresValue) Accept(visitor StartProcedureToolConfigProceduresValueVisitor) error {
	if s.typ == "ProcedureVersionRef" || s.ProcedureVersionRef != nil {
		return visitor.VisitProcedureVersionRef(s.ProcedureVersionRef)
	}
	if s.typ == "ProcedureDraftRef" || s.ProcedureDraftRef != nil {
		return visitor.VisitProcedureDraftRef(s.ProcedureDraftRef)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
