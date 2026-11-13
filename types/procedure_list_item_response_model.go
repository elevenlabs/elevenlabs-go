package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureListItemResponseModel struct {
	// Procedure ID
	ProcedureId string `json:"procedure_id" xml:"procedure_id" required:"true"`
	// Version ID of a version of the procedure. None for a procedure never versioned.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// Procedure name
	Name *string        `json:"name,omitempty" xml:"name,omitempty"`
	Type *ProcedureType `json:"type,omitempty" xml:"type,omitempty"`
	// When the agent should use this procedure. Empty string means this is a sub-procedure that should only start when another procedure references it.
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
	// True when the procedure has unpublished draft changes on this branch (a newly created or edited procedure not yet published). When true, the name, type, and trigger reflect that draft.
	HasDraft bool `json:"has_draft" xml:"has_draft" required:"true"`
	// Tool IDs referenced in the procedure content
	ReferencedToolIds []string `json:"referenced_tool_ids,omitempty" xml:"referenced_tool_ids,omitempty"`
	// Knowledge base IDs referenced in the procedure content
	ReferencedKbIds []string `json:"referenced_kb_ids,omitempty" xml:"referenced_kb_ids,omitempty"`
	// Procedure IDs referenced in the procedure content
	ReferencedProcedureIds []string `json:"referenced_procedure_ids,omitempty" xml:"referenced_procedure_ids,omitempty"`
	// Dynamic variable names used in the procedure content
	ReferencedDynamicVariables []string `json:"referenced_dynamic_variables,omitempty" xml:"referenced_dynamic_variables,omitempty"`
	// Procedure ID of the folder this procedure is placed in. None means root.
	FolderParentId  *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProcedureListItemResponseModel) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureListItemResponseModel) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *ProcedureListItemResponseModel) GetName() (value *string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *ProcedureListItemResponseModel) GetType() (value *ProcedureType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProcedureListItemResponseModel) GetTrigger() (value *string) {
	if p == nil {
		return
	}
	return p.Trigger
}

func (p *ProcedureListItemResponseModel) GetHasDraft() (value bool) {
	if p == nil {
		return
	}
	return p.HasDraft
}

func (p *ProcedureListItemResponseModel) GetReferencedToolIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedToolIds
}

func (p *ProcedureListItemResponseModel) GetReferencedKbIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedKbIds
}

func (p *ProcedureListItemResponseModel) GetReferencedProcedureIds() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedProcedureIds
}

func (p *ProcedureListItemResponseModel) GetReferencedDynamicVariables() (value []string) {
	if p == nil {
		return
	}
	return p.ReferencedDynamicVariables
}

func (p *ProcedureListItemResponseModel) GetFolderParentId() (value *string) {
	if p == nil {
		return
	}
	return p.FolderParentId
}

func (p *ProcedureListItemResponseModel) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureListItemResponseModel) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *ProcedureListItemResponseModel) SetName(name *string) {
	p.Name = name
}

func (p *ProcedureListItemResponseModel) SetType(type_ *ProcedureType) {
	p.Type = type_
}

func (p *ProcedureListItemResponseModel) SetTrigger(trigger *string) {
	p.Trigger = trigger
}

func (p *ProcedureListItemResponseModel) SetHasDraft(hasDraft bool) {
	p.HasDraft = hasDraft
}

func (p *ProcedureListItemResponseModel) SetReferencedToolIds(referencedToolIds []string) {
	p.ReferencedToolIds = referencedToolIds
}

func (p *ProcedureListItemResponseModel) SetReferencedKbIds(referencedKbIds []string) {
	p.ReferencedKbIds = referencedKbIds
}

func (p *ProcedureListItemResponseModel) SetReferencedProcedureIds(referencedProcedureIds []string) {
	p.ReferencedProcedureIds = referencedProcedureIds
}

func (p *ProcedureListItemResponseModel) SetReferencedDynamicVariables(referencedDynamicVariables []string) {
	p.ReferencedDynamicVariables = referencedDynamicVariables
}

func (p *ProcedureListItemResponseModel) SetFolderParentId(folderParentId *string) {
	p.FolderParentId = folderParentId
}

func (p *ProcedureListItemResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureListItemResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureListItemResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProcedureListItemResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *ProcedureListItemResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProcedureListItemResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
