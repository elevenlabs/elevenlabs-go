package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// How widely a group is shared, besides its owner and workspace admins (who always have access).
type RealtimeConversationGroupSharingSummary struct {
	// Distinct people besides the owner with access, whether shared directly or through a workspace group they belong to.
	PeopleCount int `json:"people_count" xml:"people_count" required:"true"`
	// Whether every workspace member has access to the group.
	SharedWithEveryone *bool `json:"shared_with_everyone,omitempty" xml:"shared_with_everyone,omitempty"`
	// The preset role every workspace member has, when shared with everyone. None for a workspace custom role.
	EveryoneRole    *RealtimeConversationGroupSharingSummaryEveryoneRole `json:"everyone_role,omitempty" xml:"everyone_role,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RealtimeConversationGroupSharingSummary) GetPeopleCount() (value int) {
	if r == nil {
		return
	}
	return r.PeopleCount
}

func (r *RealtimeConversationGroupSharingSummary) GetSharedWithEveryone() (value *bool) {
	if r == nil {
		return
	}
	return r.SharedWithEveryone
}

func (r *RealtimeConversationGroupSharingSummary) GetEveryoneRole() (value *RealtimeConversationGroupSharingSummaryEveryoneRole) {
	if r == nil {
		return
	}
	return r.EveryoneRole
}

func (r *RealtimeConversationGroupSharingSummary) SetPeopleCount(peopleCount int) {
	r.PeopleCount = peopleCount
}

func (r *RealtimeConversationGroupSharingSummary) SetSharedWithEveryone(sharedWithEveryone *bool) {
	r.SharedWithEveryone = sharedWithEveryone
}

func (r *RealtimeConversationGroupSharingSummary) SetEveryoneRole(everyoneRole *RealtimeConversationGroupSharingSummaryEveryoneRole) {
	r.EveryoneRole = everyoneRole
}

func (r *RealtimeConversationGroupSharingSummary) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RealtimeConversationGroupSharingSummary) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RealtimeConversationGroupSharingSummary to string"
	}
	return string(jsonData)
}

func (r *RealtimeConversationGroupSharingSummary) UnmarshalJSON(data []byte) error {
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

func (r *RealtimeConversationGroupSharingSummary) MarshalJSON() ([]byte, error) {
	type embed RealtimeConversationGroupSharingSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
