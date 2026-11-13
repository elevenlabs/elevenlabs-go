package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureVersionRef struct {
	// Procedure ID
	ProcedureId string `json:"procedure_id" xml:"procedure_id" required:"true"`
	// Version ID of the procedure version.
	VersionId       string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties map[string]interface{}
}

func (p *ProcedureVersionRef) GetProcedureId() (value string) {
	if p == nil {
		return
	}
	return p.ProcedureId
}

func (p *ProcedureVersionRef) GetVersionId() (value string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *ProcedureVersionRef) SetProcedureId(procedureId string) {
	p.ProcedureId = procedureId
}

func (p *ProcedureVersionRef) SetVersionId(versionId string) {
	p.VersionId = versionId
}

func (p *ProcedureVersionRef) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureVersionRef) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureVersionRef to string"
	}
	return string(jsonData)
}

func (p *ProcedureVersionRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProcedureVersionRef
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProcedureVersionRef(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProcedureVersionRef) MarshalJSON() ([]byte, error) {
	type embed ProcedureVersionRef
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
