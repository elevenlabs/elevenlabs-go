package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A stored procedure plus whether the receptionist can actually run it.
type ReceptionProcedureResponse struct {
	Name    StrippedStr `json:"name" xml:"name" required:"true"`
	Trigger StrippedStr `json:"trigger" xml:"trigger" required:"true"`
	// Step-by-step instructions, as markdown.
	Body *string `json:"body,omitempty" xml:"body,omitempty" maxLength:"40000"`
	// ID of the procedure.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Set when the procedure was copied from an integration's template; null for procedures you create.
	SourceKey *string `json:"source_key,omitempty" xml:"source_key,omitempty"`
	// Tool API names the steps point at that the receptionist does not have. While non-empty the procedure is left off calls.
	MissingTools    []string `json:"missing_tools,omitempty" xml:"missing_tools,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionProcedureResponse) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionProcedureResponse) GetTrigger() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Trigger
}

func (r *ReceptionProcedureResponse) GetBody() (value *string) {
	if r == nil {
		return
	}
	return r.Body
}

func (r *ReceptionProcedureResponse) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionProcedureResponse) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionProcedureResponse) GetSourceKey() (value *string) {
	if r == nil {
		return
	}
	return r.SourceKey
}

func (r *ReceptionProcedureResponse) GetMissingTools() (value []string) {
	if r == nil {
		return
	}
	return r.MissingTools
}

func (r *ReceptionProcedureResponse) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionProcedureResponse) SetTrigger(trigger StrippedStr) {
	r.Trigger = trigger
}

func (r *ReceptionProcedureResponse) SetBody(body *string) {
	r.Body = body
}

func (r *ReceptionProcedureResponse) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionProcedureResponse) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionProcedureResponse) SetSourceKey(sourceKey *string) {
	r.SourceKey = sourceKey
}

func (r *ReceptionProcedureResponse) SetMissingTools(missingTools []string) {
	r.MissingTools = missingTools
}

func (r *ReceptionProcedureResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionProcedureResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionProcedureResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionProcedureResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionProcedureResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionProcedureResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
