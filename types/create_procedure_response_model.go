package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateProcedureResponseModel struct {
	// ID of the created procedure
	ProcedureId     string `json:"procedure_id" xml:"procedure_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreateProcedureResponseModel) GetProcedureId() (value string) {
	if c == nil {
		return
	}
	return c.ProcedureId
}

func (c *CreateProcedureResponseModel) SetProcedureId(procedureId string) {
	c.ProcedureId = procedureId
}

func (c *CreateProcedureResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateProcedureResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateProcedureResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateProcedureResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateProcedureResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateProcedureResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateProcedureResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateProcedureResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
