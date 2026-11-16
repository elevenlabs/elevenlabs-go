package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CompileProceduresValidationErrorResponseModel struct {
	// Validation errors keyed by procedure ID.
	Errors          map[string][]*ProcedureValidationError `json:"errors" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CompileProceduresValidationErrorResponseModel) GetErrors() (value map[string][]*ProcedureValidationError) {
	if c == nil {
		return
	}
	return c.Errors
}

func (c *CompileProceduresValidationErrorResponseModel) SetErrors(errors map[string][]*ProcedureValidationError) {
	c.Errors = errors
}

func (c *CompileProceduresValidationErrorResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CompileProceduresValidationErrorResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CompileProceduresValidationErrorResponseModel to string"
	}
	return string(jsonData)
}

func (c *CompileProceduresValidationErrorResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler CompileProceduresValidationErrorResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CompileProceduresValidationErrorResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CompileProceduresValidationErrorResponseModel) MarshalJSON() ([]byte, error) {
	type embed CompileProceduresValidationErrorResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
