package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A question the receptionist asks before booking or ordering this item.
type ReceptionIntakeQuestion struct {
	// Stable identifier referenced by stored answers. Survives label edits so history stays attributable to the right question.
	Id    *string     `json:"id,omitempty" xml:"id,omitempty"`
	Label StrippedStr `json:"label" xml:"label" required:"true"`
	// Answer shape for a per-item intake question.
	Kind *IntakeQuestionKind `json:"kind,omitempty" xml:"kind,omitempty"`
	// Allowed answers for `kind='choice'`. Must be empty for the other kinds.
	Options []StrippedStr `json:"options,omitempty" xml:"options,omitempty" maxItems:"20"`
	// Enforced on the booking page. On calls the receptionist asks, but still books if the caller won't answer.
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// Retired questions are deactivated rather than deleted, so answers already stored against them keep resolving.
	IsActive        *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionIntakeQuestion) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionIntakeQuestion) GetLabel() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Label
}

func (r *ReceptionIntakeQuestion) GetKind() (value *IntakeQuestionKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ReceptionIntakeQuestion) GetOptions() (value []StrippedStr) {
	if r == nil {
		return
	}
	return r.Options
}

func (r *ReceptionIntakeQuestion) GetRequired() (value *bool) {
	if r == nil {
		return
	}
	return r.Required
}

func (r *ReceptionIntakeQuestion) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionIntakeQuestion) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionIntakeQuestion) SetLabel(label StrippedStr) {
	r.Label = label
}

func (r *ReceptionIntakeQuestion) SetKind(kind *IntakeQuestionKind) {
	r.Kind = kind
}

func (r *ReceptionIntakeQuestion) SetOptions(options []StrippedStr) {
	r.Options = options
}

func (r *ReceptionIntakeQuestion) SetRequired(required *bool) {
	r.Required = required
}

func (r *ReceptionIntakeQuestion) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionIntakeQuestion) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionIntakeQuestion) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionIntakeQuestion to string"
	}
	return string(jsonData)
}

func (r *ReceptionIntakeQuestion) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionIntakeQuestion
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionIntakeQuestion(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionIntakeQuestion) MarshalJSON() ([]byte, error) {
	type embed ReceptionIntakeQuestion
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
