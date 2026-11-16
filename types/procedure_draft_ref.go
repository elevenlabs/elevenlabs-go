package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureDraftRef struct {
	// Procedure ID
	ProcedureId     string `json:"procedure_id" xml:"procedure_id" required:"true"`
	VersionId       any    `json:"version_id,omitempty" xml:"version_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *ProcedureDraftRef) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureDraftRef) GetVersionId() (value any) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *ProcedureDraftRef) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureDraftRef) SetVersionId(versionId any) {
	p.VersionId = versionId
}

func (p *ProcedureDraftRef) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureDraftRef) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureDraftRef to string"
	}
	return string(jsonData)
}

func (p *ProcedureDraftRef) UnmarshalJSON(data []byte) error {
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

func (p *ProcedureDraftRef) MarshalJSON() ([]byte, error) {
	type embed ProcedureDraftRef
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
