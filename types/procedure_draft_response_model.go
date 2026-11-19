package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureDraftResponseModel struct {
	// Procedure ID
	ProcedureId string `json:"procedure_id" xml:"procedure_id" required:"true"`
	// Procedure name
	Name string `json:"name" xml:"name" required:"true" maxLength:"200"`
	// Procedure content
	Content string         `json:"content" xml:"content" required:"true" maxLength:"50000"`
	Type    *ProcedureType `json:"type,omitempty" xml:"type,omitempty"`
	// When the agent should use this procedure. Empty string means this is a sub-procedure that should only start when another procedure references it.
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
	// Procedure ID of the folder this procedure is placed in. None means root.
	FolderParentId  *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProcedureDraftResponseModel) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureDraftResponseModel) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProcedureDraftResponseModel) GetContent() (value string) {
	if p == nil {
		return
	}
	return p.Content
}

func (p *ProcedureDraftResponseModel) GetType() (value *ProcedureType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProcedureDraftResponseModel) GetTrigger() (value *string) {
	if p == nil {
		return
	}
	return p.Trigger
}

func (p *ProcedureDraftResponseModel) GetFolderParentId() (value *string) {
	if p == nil {
		return
	}
	return p.FolderParentId
}

func (p *ProcedureDraftResponseModel) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureDraftResponseModel) SetName(name string) {
	p.Name = name
}

func (p *ProcedureDraftResponseModel) SetContent(content string) {
	p.Content = content
}

func (p *ProcedureDraftResponseModel) SetType(type_ *ProcedureType) {
	p.Type = type_
}

func (p *ProcedureDraftResponseModel) SetTrigger(trigger *string) {
	p.Trigger = trigger
}

func (p *ProcedureDraftResponseModel) SetFolderParentId(folderParentId *string) {
	p.FolderParentId = folderParentId
}

func (p *ProcedureDraftResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureDraftResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureDraftResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProcedureDraftResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProcedureDraftResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProcedureDraftResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
