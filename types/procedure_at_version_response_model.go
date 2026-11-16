package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureAtVersionResponseModel struct {
	// Procedure ID
	ProcedureId string `json:"procedure_id" xml:"procedure_id" required:"true"`
	// Version ID of a version of the procedure. None for a procedure never versioned.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
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

func (p *ProcedureAtVersionResponseModel) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureAtVersionResponseModel) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *ProcedureAtVersionResponseModel) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProcedureAtVersionResponseModel) GetContent() (value string) {
	if p == nil {
		return
	}
	return p.Content
}

func (p *ProcedureAtVersionResponseModel) GetType() (value *ProcedureType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProcedureAtVersionResponseModel) GetTrigger() (value *string) {
	if p == nil {
		return
	}
	return p.Trigger
}

func (p *ProcedureAtVersionResponseModel) GetFolderParentId() (value *string) {
	if p == nil {
		return
	}
	return p.FolderParentId
}

func (p *ProcedureAtVersionResponseModel) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureAtVersionResponseModel) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *ProcedureAtVersionResponseModel) SetName(name string) {
	p.Name = name
}

func (p *ProcedureAtVersionResponseModel) SetContent(content string) {
	p.Content = content
}

func (p *ProcedureAtVersionResponseModel) SetType(type_ *ProcedureType) {
	p.Type = type_
}

func (p *ProcedureAtVersionResponseModel) SetTrigger(trigger *string) {
	p.Trigger = trigger
}

func (p *ProcedureAtVersionResponseModel) SetFolderParentId(folderParentId *string) {
	p.FolderParentId = folderParentId
}

func (p *ProcedureAtVersionResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureAtVersionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureAtVersionResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProcedureAtVersionResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *ProcedureAtVersionResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProcedureAtVersionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
