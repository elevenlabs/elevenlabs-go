package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StartProcedureToolResultErrorModel struct {
	Status          StartProcedureToolErrorStatus `json:"status" xml:"status" required:"true"`
	ProcedureId     *string                       `json:"procedure_id,omitempty" xml:"procedure_id,omitempty"`
	Message         string                        `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StartProcedureToolResultErrorModel) GetStatus() (value StartProcedureToolErrorStatus) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *StartProcedureToolResultErrorModel) GetProcedureId() (value *string) {
	if s == nil {
		return
	}
	return s.ProcedureId
}

func (s *StartProcedureToolResultErrorModel) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *StartProcedureToolResultErrorModel) SetStatus(status StartProcedureToolErrorStatus) {
	s.Status = status
}

func (s *StartProcedureToolResultErrorModel) SetProcedureId(procedureId *string) {
	s.ProcedureId = procedureId
}

func (s *StartProcedureToolResultErrorModel) SetMessage(message string) {
	s.Message = message
}

func (s *StartProcedureToolResultErrorModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StartProcedureToolResultErrorModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StartProcedureToolResultErrorModel to string"
	}
	return string(jsonData)
}

func (s *StartProcedureToolResultErrorModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "result_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StartProcedureToolResultErrorModel) MarshalJSON() ([]byte, error) {
	type embed StartProcedureToolResultErrorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
