package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StartProcedureToolResultSuccessModel struct {
	Status                      *string `json:"status,omitempty" xml:"status,omitempty"`
	ProcedureId                 string  `json:"procedure_id" xml:"procedure_id" required:"true"`
	ProcedureName               string  `json:"procedure_name" xml:"procedure_name" required:"true"`
	ProcedureEntryWorkflowNode  *string `json:"procedure_entry_workflow_node,omitempty" xml:"procedure_entry_workflow_node,omitempty"`
	ProcedureReturnWorkflowNode *string `json:"procedure_return_workflow_node,omitempty" xml:"procedure_return_workflow_node,omitempty"`
	Message                     *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties             map[string]interface{}
}

func (s *StartProcedureToolResultSuccessModel) GetStatus() (value *string) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *StartProcedureToolResultSuccessModel) GetProcedureId() (value string) {
	if s == nil {
		return
	}
	return s.ProcedureId
}

func (s *StartProcedureToolResultSuccessModel) GetProcedureName() (value string) {
	if s == nil {
		return
	}
	return s.ProcedureName
}

func (s *StartProcedureToolResultSuccessModel) GetProcedureEntryWorkflowNode() (value *string) {
	if s == nil {
		return
	}
	return s.ProcedureEntryWorkflowNode
}

func (s *StartProcedureToolResultSuccessModel) GetProcedureReturnWorkflowNode() (value *string) {
	if s == nil {
		return
	}
	return s.ProcedureReturnWorkflowNode
}

func (s *StartProcedureToolResultSuccessModel) GetMessage() (value *string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *StartProcedureToolResultSuccessModel) SetStatus(status *string) {
	s.Status = status
}

func (s *StartProcedureToolResultSuccessModel) SetProcedureId(procedureId string) {
	s.ProcedureId = procedureId
}

func (s *StartProcedureToolResultSuccessModel) SetProcedureName(procedureName string) {
	s.ProcedureName = procedureName
}

func (s *StartProcedureToolResultSuccessModel) SetProcedureEntryWorkflowNode(procedureEntryWorkflowNode *string) {
	s.ProcedureEntryWorkflowNode = procedureEntryWorkflowNode
}

func (s *StartProcedureToolResultSuccessModel) SetProcedureReturnWorkflowNode(procedureReturnWorkflowNode *string) {
	s.ProcedureReturnWorkflowNode = procedureReturnWorkflowNode
}

func (s *StartProcedureToolResultSuccessModel) SetMessage(message *string) {
	s.Message = message
}

func (s *StartProcedureToolResultSuccessModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StartProcedureToolResultSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StartProcedureToolResultSuccessModel to string"
	}
	return string(jsonData)
}

func (s *StartProcedureToolResultSuccessModel) UnmarshalJSON(data []byte) error {
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

func (s *StartProcedureToolResultSuccessModel) MarshalJSON() ([]byte, error) {
	type embed StartProcedureToolResultSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
