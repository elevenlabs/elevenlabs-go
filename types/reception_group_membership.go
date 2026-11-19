package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Membership of an item in a single user-defined group, with its display rank inside that group.
type ReceptionGroupMembership struct {
	Label StrippedStr `json:"label" xml:"label" required:"true"`
	// 0-based display position within this group. Change it with the reorder endpoint; send back the value you read to keep the order, a membership sent without it goes to position 0.
	Rank            *int `json:"rank,omitempty" xml:"rank,omitempty" min:"0"`
	extraProperties map[string]interface{}
}

func (r *ReceptionGroupMembership) GetLabel() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Label
}

func (r *ReceptionGroupMembership) GetRank() (value *int) {
	if r == nil {
		return
	}
	return r.Rank
}

func (r *ReceptionGroupMembership) SetLabel(label StrippedStr) {
	r.Label = label
}

func (r *ReceptionGroupMembership) SetRank(rank *int) {
	r.Rank = rank
}

func (r *ReceptionGroupMembership) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionGroupMembership) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionGroupMembership to string"
	}
	return string(jsonData)
}

func (r *ReceptionGroupMembership) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionGroupMembership
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionGroupMembership(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionGroupMembership) MarshalJSON() ([]byte, error) {
	type embed ReceptionGroupMembership
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
