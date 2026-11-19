package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureRefResponseModel struct {
	// Procedure ID
	ProcedureId string `json:"procedure_id" xml:"procedure_id" required:"true"`
	// Version ID of a version of the procedure. None for a procedure never versioned.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// Procedure name
	Name *string        `json:"name,omitempty" xml:"name,omitempty"`
	Type *ProcedureType `json:"type,omitempty" xml:"type,omitempty"`
	// When the agent should use this procedure. Empty string means this is a sub-procedure that should only start when another procedure references it.
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
	// Tool IDs referenced in the procedure content
	ReferencedToolIds []string `json:"referenced_tool_ids,omitempty" xml:"referenced_tool_ids,omitempty"`
	// Knowledge base IDs referenced in the procedure content
	ReferencedKbIds []string `json:"referenced_kb_ids,omitempty" xml:"referenced_kb_ids,omitempty"`
	// Procedure IDs referenced in the procedure content
	ReferencedProcedureIds []string `json:"referenced_procedure_ids,omitempty" xml:"referenced_procedure_ids,omitempty"`
	// Dynamic variable names used in the procedure content
	ReferencedDynamicVariables []string `json:"referenced_dynamic_variables,omitempty" xml:"referenced_dynamic_variables,omitempty"`
	// Dynamic variable names written by Update state steps
	WrittenDynamicVariables []string `json:"written_dynamic_variables,omitempty" xml:"written_dynamic_variables,omitempty"`
	// Procedure ID of the folder this procedure is placed in. None means root.
	FolderParentId  *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProcedureRefResponseModel) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureRefResponseModel) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *ProcedureRefResponseModel) GetName() (value *string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProcedureRefResponseModel) GetType() (value *ProcedureType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProcedureRefResponseModel) GetTrigger() (value *string) {
	if p == nil {
		return
	}
	return p.Trigger
}

func (p *ProcedureRefResponseModel) GetReferencedToolIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedToolIds
}

func (p *ProcedureRefResponseModel) GetReferencedKbIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedKbIds
}

func (p *ProcedureRefResponseModel) GetReferencedProcedureIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedProcedureIds
}

func (p *ProcedureRefResponseModel) GetReferencedDynamicVariables() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedDynamicVariables
}

func (p *ProcedureRefResponseModel) GetWrittenDynamicVariables() (value []string) {
	if p == nil {
		return
	}
	return p.WrittenDynamicVariables
}

func (p *ProcedureRefResponseModel) GetFolderParentId() (value *string) {
	if p == nil {
		return
	}
	return p.FolderParentId
}

func (p *ProcedureRefResponseModel) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureRefResponseModel) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *ProcedureRefResponseModel) SetName(name *string) {
	p.Name = name
}

func (p *ProcedureRefResponseModel) SetType(type_ *ProcedureType) {
	p.Type = type_
}

func (p *ProcedureRefResponseModel) SetTrigger(trigger *string) {
	p.Trigger = trigger
}

func (p *ProcedureRefResponseModel) SetReferencedToolIds(referencedToolIds []string) {
	p.ReferencedToolIds = referencedToolIds
}

func (p *ProcedureRefResponseModel) SetReferencedKbIds(referencedKbIds []string) {
	p.ReferencedKbIds = referencedKbIds
}

func (p *ProcedureRefResponseModel) SetReferencedProcedureIds(referencedProcedureIds []string) {
	p.ReferencedProcedureIds = referencedProcedureIds
}

func (p *ProcedureRefResponseModel) SetReferencedDynamicVariables(referencedDynamicVariables []string) {
	p.ReferencedDynamicVariables = referencedDynamicVariables
}

func (p *ProcedureRefResponseModel) SetWrittenDynamicVariables(writtenDynamicVariables []string) {
	p.WrittenDynamicVariables = writtenDynamicVariables
}

func (p *ProcedureRefResponseModel) SetFolderParentId(folderParentId *string) {
	p.FolderParentId = folderParentId
}

func (p *ProcedureRefResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureRefResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureRefResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProcedureRefResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *ProcedureRefResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProcedureRefResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
