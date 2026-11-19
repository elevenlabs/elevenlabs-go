package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Scrape-reachable subset of an asset (resource that can be booked).
type ReceptionAssetScrapable struct {
	// Group memberships for this item. Each entry carries the group label and the item's 0-based rank within that group.
	Groups []*ReceptionGroupMembership `json:"groups,omitempty" xml:"groups,omitempty" maxItems:"20"`
	// Display position among items with no group; `null` sorts last. Change it with the reorder endpoint.
	UngroupedRank   *int         `json:"ungrouped_rank,omitempty" xml:"ungrouped_rank,omitempty" min:"0"`
	Name            StrippedStr  `json:"name" xml:"name" required:"true"`
	Description     *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionAssetScrapable) GetGroups() (value []*ReceptionGroupMembership) {
	if r == nil {
		return
	}
	return r.Groups
}

func (r *ReceptionAssetScrapable) GetUngroupedRank() (value *int) {
	if r == nil {
		return
	}
	return r.UngroupedRank
}

func (r *ReceptionAssetScrapable) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionAssetScrapable) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionAssetScrapable) SetGroups(groups []*ReceptionGroupMembership) {
	r.Groups = groups
}

func (r *ReceptionAssetScrapable) SetUngroupedRank(ungroupedRank *int) {
	r.UngroupedRank = ungroupedRank
}

func (r *ReceptionAssetScrapable) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionAssetScrapable) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionAssetScrapable) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionAssetScrapable) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionAssetScrapable to string"
	}
	return string(jsonData)
}

func (r *ReceptionAssetScrapable) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionAssetScrapable) MarshalJSON() ([]byte, error) {
	type embed ReceptionAssetScrapable
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
