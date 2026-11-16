package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Reference to a tool for unit test evaluation.
type ReferencedToolCommonModel struct {
	// The ID of the tool
	Id string `json:"id" xml:"id" required:"true"`
	// The type of the tool
	Type            ReferencedToolCommonModelType `json:"type" xml:"type" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ReferencedToolCommonModel) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReferencedToolCommonModel) GetType() (value ReferencedToolCommonModelType) {
	if r == nil {
		return
	}
	return r.Type
}

func (r *ReferencedToolCommonModel) SetId(id string) {
	r.Id = id
}

func (r *ReferencedToolCommonModel) SetType(type_ ReferencedToolCommonModelType) {
	r.Type = type_
}

func (r *ReferencedToolCommonModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReferencedToolCommonModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReferencedToolCommonModel to string"
	}
	return string(jsonData)
}

func (r *ReferencedToolCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReferencedToolCommonModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReferencedToolCommonModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReferencedToolCommonModel) MarshalJSON() ([]byte, error) {
	type embed ReferencedToolCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
