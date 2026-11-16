package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProcedureValidationError struct {
	// JSON path to the error, e.g. 'trigger', 'steps[0].instruction'
	Path string `json:"path" xml:"path" required:"true"`
	// Human-readable error message
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (p *ProcedureValidationError) GetPath() (value string) {
	if p == nil {
		return
	}
	return p.Path
}

func (p *ProcedureValidationError) GetMessage() (value string) {
	if p == nil {
		return
	}
	return p.Message
}

func (p *ProcedureValidationError) SetPath(path string) {
	p.Path = path
}

func (p *ProcedureValidationError) SetMessage(message string) {
	p.Message = message
}

func (p *ProcedureValidationError) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProcedureValidationError) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProcedureValidationError to string"
	}
	return string(jsonData)
}

func (p *ProcedureValidationError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProcedureValidationError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProcedureValidationError(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProcedureValidationError) MarshalJSON() ([]byte, error) {
	type embed ProcedureValidationError
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
