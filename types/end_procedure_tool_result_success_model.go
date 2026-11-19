package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EndProcedureToolResultSuccessModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	ProcedureId     string  `json:"procedure_id" xml:"procedure_id" required:"true"`
	ProcedureName   string  `json:"procedure_name" xml:"procedure_name" required:"true"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EndProcedureToolResultSuccessModel) GetStatus() (value *string) {
	if e == nil {
		return
	}
	return e.Status
}

func (e *EndProcedureToolResultSuccessModel) GetProcedureId() (value string) {
	if e == nil {
		return
	}
	return e.ProcedureId
}

func (e *EndProcedureToolResultSuccessModel) GetProcedureName() (value string) {
	if e == nil {
		return
	}
	return e.ProcedureName
}

func (e *EndProcedureToolResultSuccessModel) GetMessage() (value *string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *EndProcedureToolResultSuccessModel) SetStatus(status *string) {
	e.Status = status
}

func (e *EndProcedureToolResultSuccessModel) SetProcedureId(procedureId string) {
	e.ProcedureId = procedureId
}

func (e *EndProcedureToolResultSuccessModel) SetProcedureName(procedureName string) {
	e.ProcedureName = procedureName
}

func (e *EndProcedureToolResultSuccessModel) SetMessage(message *string) {
	e.Message = message
}

func (e *EndProcedureToolResultSuccessModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndProcedureToolResultSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndProcedureToolResultSuccessModel to string"
	}
	return string(jsonData)
}

func (e *EndProcedureToolResultSuccessModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EndProcedureToolResultSuccessModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EndProcedureToolResultSuccessModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "result_type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndProcedureToolResultSuccessModel) MarshalJSON() ([]byte, error) {
	type embed EndProcedureToolResultSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
