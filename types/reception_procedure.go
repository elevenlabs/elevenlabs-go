package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A standard operating procedure (SOP) the receptionist can run on calls.
type ReceptionProcedure struct {
	Name    StrippedStr `json:"name" xml:"name" required:"true"`
	Trigger StrippedStr `json:"trigger" xml:"trigger" required:"true"`
	// Step-by-step instructions, as markdown.
	Body *string `json:"body,omitempty" xml:"body,omitempty" maxLength:"40000"`
	// ID of the procedure.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Set when the procedure was copied from an integration's template; null for procedures you create.
	SourceKey       *string `json:"source_key,omitempty" xml:"source_key,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionProcedure) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionProcedure) GetTrigger() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Trigger
}

func (r *ReceptionProcedure) GetBody() (value *string) {
	if r == nil {
		return
	}
	return r.Body
}

func (r *ReceptionProcedure) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionProcedure) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionProcedure) GetSourceKey() (value *string) {
	if r == nil {
		return
	}
	return r.SourceKey
}

func (r *ReceptionProcedure) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionProcedure) SetTrigger(trigger StrippedStr) {
	r.Trigger = trigger
}

func (r *ReceptionProcedure) SetBody(body *string) {
	r.Body = body
}

func (r *ReceptionProcedure) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionProcedure) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionProcedure) SetSourceKey(sourceKey *string) {
	r.SourceKey = sourceKey
}

func (r *ReceptionProcedure) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionProcedure) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionProcedure to string"
	}
	return string(jsonData)
}

func (r *ReceptionProcedure) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionProcedure) MarshalJSON() ([]byte, error) {
	type embed ReceptionProcedure
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
