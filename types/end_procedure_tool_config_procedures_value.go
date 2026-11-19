package types

import (
	"encoding/json"
	"fmt"
)

type EndProcedureToolConfigProceduresValue struct {
	ProcedureVersionRef *ProcedureVersionRef
	ProcedureDraftRef   *ProcedureDraftRef
	typ                 string
}

func (e *EndProcedureToolConfigProceduresValue) GetProcedureVersionRef() *ProcedureVersionRef {
	if e == nil {
		return nil
	}
	if e.ProcedureVersionRef == nil {
		return nil
	}
	return e.ProcedureVersionRef
}

func (e *EndProcedureToolConfigProceduresValue) GetProcedureDraftRef() *ProcedureDraftRef {
	if e == nil {
		return nil
	}
	if e.ProcedureDraftRef == nil {
		return nil
	}
	return e.ProcedureDraftRef
}

func (e *EndProcedureToolConfigProceduresValue) UnmarshalJSON(data []byte) error {
	var valueProcedureVersionRef *ProcedureVersionRef
	if err := json.Unmarshal(data, &valueProcedureVersionRef); err == nil {
		e.typ = "ProcedureVersionRef"
		e.ProcedureVersionRef = valueProcedureVersionRef
		return nil
	}
	var valueProcedureDraftRef *ProcedureDraftRef
	if err := json.Unmarshal(data, &valueProcedureDraftRef); err == nil {
		e.typ = "ProcedureDraftRef"
		e.ProcedureDraftRef = valueProcedureDraftRef
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EndProcedureToolConfigProceduresValue) MarshalJSON() ([]byte, error) {
	if e.typ == "ProcedureVersionRef" || e.ProcedureVersionRef != nil {
		return json.Marshal(e.ProcedureVersionRef)
	}
	if e.typ == "ProcedureDraftRef" || e.ProcedureDraftRef != nil {
		return json.Marshal(e.ProcedureDraftRef)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EndProcedureToolConfigProceduresValueVisitor interface {
	VisitProcedureVersionRef(*ProcedureVersionRef) error
	VisitProcedureDraftRef(*ProcedureDraftRef) error
}

func (e *EndProcedureToolConfigProceduresValue) Accept(visitor EndProcedureToolConfigProceduresValueVisitor) error {
	if e.typ == "ProcedureVersionRef" || e.ProcedureVersionRef != nil {
		return visitor.VisitProcedureVersionRef(e.ProcedureVersionRef)
	}
	if e.typ == "ProcedureDraftRef" || e.ProcedureDraftRef != nil {
		return visitor.VisitProcedureDraftRef(e.ProcedureDraftRef)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}
