package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EndProcedureToolResultErrorModel struct {
	Status          EndProcedureToolErrorStatus `json:"status" xml:"status" required:"true"`
	ProcedureId     *string                     `json:"procedure_id,omitempty" xml:"procedure_id,omitempty"`
	Message         string                      `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EndProcedureToolResultErrorModel) GetStatus() (value EndProcedureToolErrorStatus) {
	if e == nil {
		return
	}
	return e.Status
}

func (e *EndProcedureToolResultErrorModel) GetProcedureId() (value *string) {
	if e == nil {
		return
	}
	return e.ProcedureId
}

func (e *EndProcedureToolResultErrorModel) GetMessage() (value string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *EndProcedureToolResultErrorModel) SetStatus(status EndProcedureToolErrorStatus) {
	e.Status = status
}

func (e *EndProcedureToolResultErrorModel) SetProcedureId(procedureId *string) {
	e.ProcedureId = procedureId
}

func (e *EndProcedureToolResultErrorModel) SetMessage(message string) {
	e.Message = message
}

func (e *EndProcedureToolResultErrorModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndProcedureToolResultErrorModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndProcedureToolResultErrorModel to string"
	}
	return string(jsonData)
}

func (e *EndProcedureToolResultErrorModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "result_type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndProcedureToolResultErrorModel) MarshalJSON() ([]byte, error) {
	type embed EndProcedureToolResultErrorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
